package stdlib

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"

	"github.com/2dprototype/tender"
)

var cryptoRandomModule = &tender.ImmutableMap{
	Value: map[string]tender.Object{
		"bytes":      &tender.NativeFunction{Name: "bytes", Value: randomBytes},
		"int":        &tender.NativeFunction{Name: "int", Value: randomInt},
		"float":      &tender.NativeFunction{Name: "float", Value: randomFloat},
		"uuid":       &tender.NativeFunction{Name: "uuid", Value: randomUUID},
	},
}

func randomBytes(args ...tender.Object) (ret tender.Object, err error) {
	size := 32
	if len(args) > 0 {
		sizeArg, ok := tender.ToInt(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "size",
				Expected: "int",
				Found:    args[0].TypeName(),
			}
		}
		size = sizeArg
	}

	if size <= 0 {
		// return nil, &tender.Error{Value: &tender.String{Value: "size must be positive"}}
		return nil, nil
	}

	bytes := make([]byte, size)
	_, err = rand.Read(bytes)
	if err != nil {
		return wrapError(err), nil
	}

	return &tender.Bytes{Value: bytes}, nil
}

func randomInt(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) < 2 {
		return nil, tender.ErrWrongNumArguments
	}

	min, ok := tender.ToInt(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "min",
			Expected: "int",
			Found:    args[0].TypeName(),
		}
	}

	max, ok := tender.ToInt(args[1])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "max",
			Expected: "int",
			Found:    args[1].TypeName(),
		}
	}

	if min >= max {
		// return nil, &tender.Error{Value: &tender.String{Value: "min must be less than max"}}
		return nil, nil
	}

	var buf [8]byte
	_, err = rand.Read(buf[:])
	if err != nil {
		return wrapError(err), nil
	}

	randomUint := binary.BigEndian.Uint64(buf[:])
	rangeSize := int64(max - min + 1)
	result := min + int(randomUint%uint64(rangeSize))

	return &tender.Int{Value: int64(result)}, nil
}

func randomFloat(args ...tender.Object) (ret tender.Object, err error) {
	var buf [8]byte
	_, err = rand.Read(buf[:])
	if err != nil {
		return wrapError(err), nil
	}

	randomUint := binary.BigEndian.Uint64(buf[:])
	// Convert to float in [0, 1)
	result := float64(randomUint) / float64(1<<64)

	return &tender.Float{Value: result}, nil
}

func randomUUID(args ...tender.Object) (ret tender.Object, err error) {
	bytes := make([]byte, 16)
	_, err = rand.Read(bytes)
	if err != nil {
		return wrapError(err), nil
	}

	// Set version (4) and variant (RFC 4122)
	bytes[6] = (bytes[6] & 0x0f) | 0x40
	bytes[8] = (bytes[8] & 0x3f) | 0x80

	uuid := make([]byte, 36)
	hex.Encode(uuid[0:8], bytes[0:4])
	uuid[8] = '-'
	hex.Encode(uuid[9:13], bytes[4:6])
	uuid[13] = '-'
	hex.Encode(uuid[14:18], bytes[6:8])
	uuid[18] = '-'
	hex.Encode(uuid[19:23], bytes[8:10])
	uuid[23] = '-'
	hex.Encode(uuid[24:], bytes[10:])

	return &tender.String{Value: string(uuid)}, nil
}