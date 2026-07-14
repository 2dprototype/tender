package stdlib

import (
	"encoding/hex"

	"golang.org/x/crypto/argon2"

	"github.com/2dprototype/tender"
)

var cryptoArgon2Module = &tender.ImmutableMap{
	Value: map[string]tender.Object{
		"id":   &tender.NativeFunction{Name: "id", Value: argon2id},
		"i":    &tender.NativeFunction{Name: "i", Value: argon2i},
	},
}

func argon2id(args ...tender.Object) (ret tender.Object, err error) {
	return argon2Hash(args, argon2.IDKey)
}

func argon2i(args ...tender.Object) (ret tender.Object, err error) {
	return argon2Hash(args, argon2.Key)
}

func argon2Hash(args []tender.Object, argon2Func func([]byte, []byte, uint32, uint32, uint8, uint32) []byte) (tender.Object, error) {
	if len(args) < 2 {
		return nil, tender.ErrWrongNumArguments
	}

	password, _ := tender.ToByteSlice(args[0])
	salt, _ := tender.ToByteSlice(args[1])

	timeCost := uint32(1)
	if len(args) > 2 {
		timeCostArg, ok := tender.ToInt(args[2])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "timeCost",
				Expected: "int",
				Found:    args[2].TypeName(),
			}
		}
		timeCost = uint32(timeCostArg)
	}

	memoryCost := uint32(64 * 1024)
	if len(args) > 3 {
		memoryCostArg, ok := tender.ToInt(args[3])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "memoryCost",
				Expected: "int",
				Found:    args[3].TypeName(),
			}
		}
		memoryCost = uint32(memoryCostArg)
	}

	threads := uint8(4)
	if len(args) > 4 {
		threadsArg, ok := tender.ToInt(args[4])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "threads",
				Expected: "int",
				Found:    args[4].TypeName(),
			}
		}
		threads = uint8(threadsArg)
	}

	keyLen := uint32(32)
	if len(args) > 5 {
		keyLenArg, ok := tender.ToInt(args[5])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "keyLen",
				Expected: "int",
				Found:    args[5].TypeName(),
			}
		}
		keyLen = uint32(keyLenArg)
	}

	hash := argon2Func(password, salt, timeCost, memoryCost, threads, keyLen)
	return &tender.String{Value: hex.EncodeToString(hash)}, nil
}