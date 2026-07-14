package stdlib

import (
	"crypto/ed25519"
	"crypto/rand"

	"github.com/2dprototype/tender"
)

var cryptoEd25519Module = &tender.ImmutableMap{
	Value: map[string]tender.Object{
		"generate_key": &tender.NativeFunction{Name: "generate_key", Value: ed25519GenerateKey},
		"sign":         &tender.NativeFunction{Name: "sign", Value: ed25519Sign},
		"verify":       &tender.NativeFunction{Name: "verify", Value: ed25519Verify},
	},
}

func ed25519GenerateKey(args ...tender.Object) (ret tender.Object, err error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return wrapError(err), nil
	}

	keyPair := map[string]tender.Object{
		"public":  &tender.Bytes{Value: publicKey},
		"private": &tender.Bytes{Value: privateKey},
	}

	return &tender.Map{Value: keyPair}, nil
}

func ed25519Sign(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) < 2 {
		return nil, tender.ErrWrongNumArguments
	}

	data, _ := tender.ToByteSlice(args[0])
	privateKey, _ := tender.ToByteSlice(args[1])

	signature := ed25519.Sign(ed25519.PrivateKey(privateKey), data)
	return &tender.Bytes{Value: signature}, nil
}

func ed25519Verify(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) < 3 {
		return nil, tender.ErrWrongNumArguments
	}

	data, _ := tender.ToByteSlice(args[0])
	signature, _ := tender.ToByteSlice(args[1])
	publicKey, _ := tender.ToByteSlice(args[2])

	valid := ed25519.Verify(ed25519.PublicKey(publicKey), data, signature)

	if valid {
		return tender.TrueValue,  nil
	} else {
		return tender.FalseValue,  nil
	}
}