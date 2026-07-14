
package stdlib

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/sha3"

	"github.com/2dprototype/tender"
)

var cryptoHMACModule = &tender.ImmutableMap{
	Value: map[string]tender.Object{
		"md5":      &tender.NativeFunction{Name: "md5", Value: hmacMD5},
		"sha1":     &tender.NativeFunction{Name: "sha1", Value: hmacSHA1},
		"sha256":   &tender.NativeFunction{Name: "sha256", Value: hmacSHA256},
		"sha384":   &tender.NativeFunction{Name: "sha384", Value: hmacSHA384},
		"sha512":   &tender.NativeFunction{Name: "sha512", Value: hmacSHA512},
		"sha3_256": &tender.NativeFunction{Name: "sha3_256", Value: hmacSHA3_256},
		"sha3_512": &tender.NativeFunction{Name: "sha3_512", Value: hmacSHA3_512},
	},
}

func hmacMD5(args ...tender.Object) (ret tender.Object, err error) {
	return hmacHash(args, "md5")
}

func hmacSHA1(args ...tender.Object) (ret tender.Object, err error) {
	return hmacHash(args, "sha1")
}

func hmacSHA256(args ...tender.Object) (ret tender.Object, err error) {
	return hmacHash(args, "sha256")
}

func hmacSHA384(args ...tender.Object) (ret tender.Object, err error) {
	return hmacHash(args, "sha384")
}

func hmacSHA512(args ...tender.Object) (ret tender.Object, err error) {
	return hmacHash(args, "sha512")
}

func hmacSHA3_256(args ...tender.Object) (ret tender.Object, err error) {
	return hmacHash(args, "sha3_256")
}

func hmacSHA3_512(args ...tender.Object) (ret tender.Object, err error) {
	return hmacHash(args, "sha3_512")
}

func hmacHash(args []tender.Object, algorithm string) (tender.Object, error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}

	key, _ := tender.ToByteSlice(args[0])
	data, _ := tender.ToByteSlice(args[1])

	var hash []byte
	switch strings.ToLower(algorithm) {
	case "md5":
		mac := hmac.New(md5.New, key)
		mac.Write(data)
		hash = mac.Sum(nil)
	case "sha1":
		mac := hmac.New(sha1.New, key)
		mac.Write(data)
		hash = mac.Sum(nil)
	case "sha256":
		mac := hmac.New(sha256.New, key)
		mac.Write(data)
		hash = mac.Sum(nil)
	case "sha384":
		mac := hmac.New(sha512.New384, key)
		mac.Write(data)
		hash = mac.Sum(nil)
	case "sha512":
		mac := hmac.New(sha512.New, key)
		mac.Write(data)
		hash = mac.Sum(nil)
	case "sha3_256":
		mac := hmac.New(sha3.New256, key)
		mac.Write(data)
		hash = mac.Sum(nil)
	case "sha3_512":
		mac := hmac.New(sha3.New512, key)
		mac.Write(data)
		hash = mac.Sum(nil)
	default:
		// return nil, &tender.Error{Value: &tender.String{Value: "unsupported HMAC algorithm: " + algorithm}}
		return nil, nil
	}

	return &tender.String{Value: hex.EncodeToString(hash)}, nil
}