package stdlib

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512" 
    "crypto/sha1"

	"crypto/subtle"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"

	"github.com/2dprototype/tender"
)

var cryptoModule = map[string]tender.Object{
    "md5":          &tender.NativeFunction{Name: "md5", Value: cryptoMd5Hash},
    "sha1":         &tender.NativeFunction{Name: "sha1", Value: cryptoSha1Hash},
    "sha224":       &tender.NativeFunction{Name: "sha224", Value: cryptoSha224Hash},
    "sha256":       &tender.NativeFunction{Name: "sha256", Value: cryptoSha256Hash},
    "sha384":       &tender.NativeFunction{Name: "sha384", Value: cryptoSha384Hash},
    "sha512":       &tender.NativeFunction{Name: "sha512", Value: cryptoSha512Hash},
    "sha3_224":     &tender.NativeFunction{Name: "sha3_224", Value: cryptoSha3_224Hash},
    "sha3_256":     &tender.NativeFunction{Name: "sha3_256", Value: cryptoSha3_256Hash},
    "sha3_384":     &tender.NativeFunction{Name: "sha3_384", Value: cryptoSha3_384Hash},
    "sha3_512":     &tender.NativeFunction{Name: "sha3_512", Value: cryptoSha3_512Hash},
    "blake2b_256":  &tender.NativeFunction{Name: "blake2b_256", Value: cryptoBlake2b256Hash},
    "blake2b_512":  &tender.NativeFunction{Name: "blake2b_512", Value: cryptoBlake2b512Hash},
	"hmac":         cryptoHMACModule,
	"aes":          cryptoAESModule,
	"rsa":          cryptoRSAModule,
	"ecdsa":        cryptoECDSAModule,
	"ed25519":      cryptoEd25519Module,
	"random":       cryptoRandomModule,
	"pbkdf2":       &tender.NativeFunction{Name: "pbkdf2", Value: cryptoPBKDF2},
	"bcrypt":       &tender.NativeFunction{Name: "bcrypt", Value: cryptoBcrypt},
	"scrypt":       &tender.NativeFunction{Name: "scrypt", Value: cryptoScrypt},
	"argon2":       cryptoArgon2Module,
	"constant_time_compare": &tender.NativeFunction{Name: "constant_time_compare", Value: cryptoConstantTimeCompare},
}

func cryptoMd5Hash(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	input, _ := tender.ToByteSlice(args[0])
	hash := md5.Sum(input)
	hashString := hex.EncodeToString(hash[:])
	return &tender.String{Value: hashString}, nil
}

func cryptoSha1Hash(args ...tender.Object) (ret tender.Object, err error) {
    if len(args) != 1 {
        return nil, tender.ErrWrongNumArguments
    }

    input, _ := tender.ToByteSlice(args[0])
    hash := sha1.Sum(input)
    hashString := hex.EncodeToString(hash[:])
    return &tender.String{Value: hashString}, nil
}

func cryptoSha224Hash(args ...tender.Object) (ret tender.Object, err error) {
    if len(args) != 1 {
        return nil, tender.ErrWrongNumArguments
    }

    input, _ := tender.ToByteSlice(args[0])
    hash := sha256.Sum224(input)
    hashString := hex.EncodeToString(hash[:])
    return &tender.String{Value: hashString}, nil
}

func cryptoSha256Hash(args ...tender.Object) (ret tender.Object, err error) {
    if len(args) != 1 {
        return nil, tender.ErrWrongNumArguments
    }

    input, _ := tender.ToByteSlice(args[0])
    hash := sha256.Sum256(input)
    hashString := hex.EncodeToString(hash[:])
    return &tender.String{Value: hashString}, nil
}

func cryptoSha384Hash(args ...tender.Object) (ret tender.Object, err error) {
    if len(args) != 1 {
        return nil, tender.ErrWrongNumArguments
    }

    input, _ := tender.ToByteSlice(args[0])
    hash := sha512.Sum384(input)
    hashString := hex.EncodeToString(hash[:])
    return &tender.String{Value: hashString}, nil
}

func cryptoSha512Hash(args ...tender.Object) (ret tender.Object, err error) {
    if len(args) != 1 {
        return nil, tender.ErrWrongNumArguments
    }

    input, _ := tender.ToByteSlice(args[0])
    hash := sha512.Sum512(input)
    hashString := hex.EncodeToString(hash[:])
    return &tender.String{Value: hashString}, nil
}

func cryptoSha3_224Hash(args ...tender.Object) (ret tender.Object, err error) {
    if len(args) != 1 {
        return nil, tender.ErrWrongNumArguments
    }

    input, _ := tender.ToByteSlice(args[0])
    hash := sha3.Sum224(input)
    hashString := hex.EncodeToString(hash[:])
    return &tender.String{Value: hashString}, nil
}

func cryptoSha3_256Hash(args ...tender.Object) (ret tender.Object, err error) {
    if len(args) != 1 {
        return nil, tender.ErrWrongNumArguments
    }

    input, _ := tender.ToByteSlice(args[0])
    hash := sha3.Sum256(input)
    hashString := hex.EncodeToString(hash[:])
    return &tender.String{Value: hashString}, nil
}

func cryptoSha3_384Hash(args ...tender.Object) (ret tender.Object, err error) {
    if len(args) != 1 {
        return nil, tender.ErrWrongNumArguments
    }

    input, _ := tender.ToByteSlice(args[0])
    hash := sha3.Sum384(input)
    hashString := hex.EncodeToString(hash[:])
    return &tender.String{Value: hashString}, nil
}

func cryptoSha3_512Hash(args ...tender.Object) (ret tender.Object, err error) {
    if len(args) != 1 {
        return nil, tender.ErrWrongNumArguments
    }

    input, _ := tender.ToByteSlice(args[0])
    hash := sha3.Sum512(input)
    hashString := hex.EncodeToString(hash[:])
    return &tender.String{Value: hashString}, nil
}

func cryptoBlake2b256Hash(args ...tender.Object) (ret tender.Object, err error) {
    if len(args) != 1 {
        return nil, tender.ErrWrongNumArguments
    }

    input, _ := tender.ToByteSlice(args[0])
    hash := sha3.New256()
    hash.Write(input)
    hashBytes := hash.Sum(nil)
    hashString := hex.EncodeToString(hashBytes)
    return &tender.String{Value: hashString}, nil
}

func cryptoBlake2b512Hash(args ...tender.Object) (ret tender.Object, err error) {
    if len(args) != 1 {
        return nil, tender.ErrWrongNumArguments
    }

    input, _ := tender.ToByteSlice(args[0])
    hash := sha3.New512()
    hash.Write(input)
    hashBytes := hash.Sum(nil)
    hashString := hex.EncodeToString(hashBytes)
    return &tender.String{Value: hashString}, nil
}

func cryptoPBKDF2(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) < 4 {
		return nil, tender.ErrWrongNumArguments
	}

	password, _ := tender.ToByteSlice(args[0])
	salt, _ := tender.ToByteSlice(args[1])
	
	iterations, ok := tender.ToInt(args[2])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "iterations",
			Expected: "int",
			Found:    args[2].TypeName(),
		}
	}
	
	keyLen, ok := tender.ToInt(args[3])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "keyLen",
			Expected: "int",
			Found:    args[3].TypeName(),
		}
	}
	
	hashFunc := "sha256"
	if len(args) > 4 {
		hashFuncStr, ok := tender.ToString(args[4])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "hashFunc",
				Expected: "string",
				Found:    args[4].TypeName(),
			}
		}
		hashFunc = hashFuncStr
	}
	
	var key []byte
	switch strings.ToLower(hashFunc) {
	case "sha1":
		key = pbkdf2.Key(password, salt, iterations, keyLen, sha1.New)
	case "sha256":
		key = pbkdf2.Key(password, salt, iterations, keyLen, sha256.New)
	case "sha512":
		key = pbkdf2.Key(password, salt, iterations, keyLen, sha512.New)
	case "sha3_256":
		key = pbkdf2.Key(password, salt, iterations, keyLen, sha3.New256)
	case "sha3_512":
		key = pbkdf2.Key(password, salt, iterations, keyLen, sha3.New512)
	default:
		// return nil, &tender.Error{Value: &tender.String{Value: "unsupported hash function: " + hashFunc}}
		return nil, nil
	}
	
	return &tender.Bytes{Value: key}, nil
}

func cryptoBcrypt(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) < 1 {
		return nil, tender.ErrWrongNumArguments
	}

	password, _ := tender.ToByteSlice(args[0])
	
	cost := 10
	if len(args) > 1 {
		costArg, ok := tender.ToInt(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "cost",
				Expected: "int",
				Found:    args[1].TypeName(),
			}
		}
		cost = costArg
	}
	
	hash, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return wrapError(err), nil
	}
	
	return &tender.Bytes{Value: hash}, nil
}

func cryptoScrypt(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) < 3 {
		return nil, tender.ErrWrongNumArguments
	}

	password, _ := tender.ToByteSlice(args[0])
	salt, _ := tender.ToByteSlice(args[1])
	
	keyLen := 32
	if len(args) > 2 {
		keyLenArg, ok := tender.ToInt(args[2])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "keyLen",
				Expected: "int",
				Found:    args[2].TypeName(),
			}
		}
		keyLen = keyLenArg
	}
	
	N := 32768
	if len(args) > 3 {
		nArg, ok := tender.ToInt(args[3])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "N",
				Expected: "int",
				Found:    args[3].TypeName(),
			}
		}
		N = nArg
	}
	
	r := 8
	if len(args) > 4 {
		rArg, ok := tender.ToInt(args[4])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "r",
				Expected: "int",
				Found:    args[4].TypeName(),
			}
		}
		r = rArg
	}
	
	p := 1
	if len(args) > 5 {
		pArg, ok := tender.ToInt(args[5])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "p",
				Expected: "int",
				Found:    args[5].TypeName(),
			}
		}
		p = pArg
	}
	
	key, err := scrypt.Key(password, salt, N, r, p, keyLen)
	if err != nil {
		return wrapError(err), nil
	}
	
	return &tender.Bytes{Value: key}, nil
}

func cryptoConstantTimeCompare(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}

	a, _ := tender.ToByteSlice(args[0])
	b, _ := tender.ToByteSlice(args[1])
	
	result := subtle.ConstantTimeCompare(a, b)
	// return &tender.Bool{Value: result == 1}, nil
	if result == 1 {
		return tender.TrueValue,  nil
	} else {
		return tender.FalseValue,  nil
	}
}

