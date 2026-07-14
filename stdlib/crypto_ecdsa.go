package stdlib

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"math/big"

	"github.com/2dprototype/tender"
)

var cryptoECDSAModule = &tender.ImmutableMap{
	Value: map[string]tender.Object{
		"generate_key": &tender.NativeFunction{Name: "generate_key", Value: ecdsaGenerateKey},
		"sign":         &tender.NativeFunction{Name: "sign", Value: ecdsaSign},
		"verify":       &tender.NativeFunction{Name: "verify", Value: ecdsaVerify},
		"export_key":   &tender.NativeFunction{Name: "export_key", Value: ecdsaExportKey},
		"import_key":   &tender.NativeFunction{Name: "import_key", Value: ecdsaImportKey},
	},
}

func ecdsaGenerateKey(args ...tender.Object) (ret tender.Object, err error) {
	curveName := "p256"
	if len(args) > 0 {
		curveNameArg, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "curve",
				Expected: "string",
				Found:    args[0].TypeName(),
			}
		}
		curveName = curveNameArg
	}

	var curve elliptic.Curve
	switch curveName {
	case "p224":
		curve = elliptic.P224()
	case "p256":
		curve = elliptic.P256()
	case "p384":
		curve = elliptic.P384()
	case "p521":
		curve = elliptic.P521()
	default:
		return wrapError(errors.New("unsupported curve: " + curveName)), nil
	}

	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return wrapError(err), nil
	}

	// Return both private and public key as a map
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return wrapError(err), nil
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return wrapError(err), nil
	}

	keyPair := map[string]tender.Object{
		"private": &tender.Bytes{Value: privateKeyBytes},
		"public":  &tender.Bytes{Value: publicKeyBytes},
	}

	return &tender.Map{Value: keyPair}, nil
}

func ecdsaSign(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) < 2 {
		return nil, tender.ErrWrongNumArguments
	}

	data, _ := tender.ToByteSlice(args[0])
	privateKeyBytes, _ := tender.ToByteSlice(args[1])

	privateKey, err := x509.ParseECPrivateKey(privateKeyBytes)
	if err != nil {
		// Try parsing as PEM
		block, _ := pem.Decode(privateKeyBytes)
		if block != nil {
			if block.Type == "EC PRIVATE KEY" {
				privateKey, err = x509.ParseECPrivateKey(block.Bytes)
				if err != nil {
					return wrapError(err), nil
				}
			} else {
				return wrapError(errors.New("unsupported PEM type: " + block.Type)), nil
			}
		} else {
			return wrapError(err), nil
		}
	}

	hashed := sha256.Sum256(data)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashed[:])
	if err != nil {
		return wrapError(err), nil
	}

	// Encode signature as simple concatenation of r and s (more reliable than ASN.1)
	signature := make([]byte, 64) // 32 bytes for r, 32 bytes for s
	rBytes := r.Bytes()
	sBytes := s.Bytes()
	
	// Copy r and s to fixed-size positions
	copy(signature[32-len(rBytes):32], rBytes)
	copy(signature[64-len(sBytes):64], sBytes)

	return &tender.Bytes{Value: signature}, nil
}

func ecdsaVerify(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) < 3 {
		return nil, tender.ErrWrongNumArguments
	}

	data, _ := tender.ToByteSlice(args[0])
	signature, _ := tender.ToByteSlice(args[1])
	publicKeyBytes, _ := tender.ToByteSlice(args[2])

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		// Try parsing as PEM
		block, _ := pem.Decode(publicKeyBytes)
		if block != nil {
			if block.Type == "PUBLIC KEY" {
				publicKey, err = x509.ParsePKIXPublicKey(block.Bytes)
				if err != nil {
					return wrapError(err), nil
				}
			} else {
				return wrapError(errors.New("unsupported PEM type: " + block.Type)), nil
			}
		} else {
			return wrapError(err), nil
		}
	}

	ecdsaPublicKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return wrapError(errors.New("not an ECDSA public key")), nil
	}

	// Parse signature as simple concatenation (64 bytes total)
	if len(signature) != 64 {
		return wrapError(errors.New("signature must be 64 bytes")), nil
	}

	r := new(big.Int).SetBytes(signature[:32])
	s := new(big.Int).SetBytes(signature[32:])

	hashed := sha256.Sum256(data)
	valid := ecdsa.Verify(ecdsaPublicKey, hashed[:], r, s)
	if valid {
		return tender.TrueValue, nil
	} else {
		return tender.FalseValue, nil
	}
}

func ecdsaExportKey(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) < 2 {
		return nil, tender.ErrWrongNumArguments
	}

	keyBytes, _ := tender.ToByteSlice(args[0])
	keyType, ok := tender.ToString(args[1])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "keyType",
			Expected: "string",
			Found:    args[1].TypeName(),
		}
	}

	var pemBlock *pem.Block
	switch keyType {
	case "private":
		privateKey, err := x509.ParseECPrivateKey(keyBytes)
		if err != nil {
			return wrapError(err), nil
		}
		privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
		if err != nil {
			return wrapError(err), nil
		}
		pemBlock = &pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: privateKeyBytes,
		}
	case "public":
		// If we have private key bytes, extract public key from it
		privateKey, err := x509.ParseECPrivateKey(keyBytes)
		if err == nil {
			publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
			if err != nil {
				return wrapError(err), nil
			}
			pemBlock = &pem.Block{
				Type:  "PUBLIC KEY",
				Bytes: publicKeyBytes,
			}
		} else {
			// Try to parse as public key directly
			publicKey, err := x509.ParsePKIXPublicKey(keyBytes)
			if err != nil {
				return wrapError(err), nil
			}
			publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
			if err != nil {
				return wrapError(err), nil
			}
			pemBlock = &pem.Block{
				Type:  "PUBLIC KEY",
				Bytes: publicKeyBytes,
			}
		}
	default:
		return wrapError(errors.New("keyType must be 'private' or 'public'")), nil
	}

	pemData := pem.EncodeToMemory(pemBlock)
	return &tender.Bytes{Value: pemData}, nil
}

func ecdsaImportKey(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) < 2 {
		return nil, tender.ErrWrongNumArguments
	}

	pemData, _ := tender.ToByteSlice(args[0])
	keyType, ok := tender.ToString(args[1])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "keyType",
			Expected: "string",
			Found:    args[1].TypeName(),
		}
	}

	pemBlock, _ := pem.Decode(pemData)
	if pemBlock == nil {
		return wrapError(errors.New("invalid PEM data")), nil
	}

	switch keyType {
	case "private":
		privateKey, err := x509.ParseECPrivateKey(pemBlock.Bytes)
		if err != nil {
			return wrapError(err), nil
		}
		privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
		if err != nil {
			return wrapError(err), nil
		}
		return &tender.Bytes{Value: privateKeyBytes}, nil
	case "public":
		publicKey, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
		if err != nil {
			return wrapError(err), nil
		}
		publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
		if err != nil {
			return wrapError(err), nil
		}
		return &tender.Bytes{Value: publicKeyBytes}, nil
	default:
		return wrapError(errors.New("keyType must be 'private' or 'public'")), nil
	}
}
