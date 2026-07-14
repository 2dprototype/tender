package stdlib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"github.com/2dprototype/tender"
)

var cryptoAESModule =  &tender.ImmutableMap{
	Value: map[string]tender.Object{
		"encrypt":    &tender.NativeFunction{Name: "encrypt", Value: aesEncrypt},
		"decrypt":    &tender.NativeFunction{Name: "decrypt", Value: aesDecrypt},
		"block_size": &tender.Int{Value: aes.BlockSize},
	},
}

func aesEncrypt(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}

	plaintext, _ := tender.ToByteSlice(args[0])
	key, _ := tender.ToByteSlice(args[1])

	ciphertext, err := encrypt(key, plaintext)
	if err != nil {
		return wrapError(err), nil
	}

	return &tender.Bytes{Value: ciphertext}, nil
}

func aesDecrypt(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}

	ciphertext, _ := tender.ToByteSlice(args[0])
	key, _ := tender.ToByteSlice(args[1])

	plaintext, err := decrypt(key, ciphertext)
	if err != nil {
		return wrapError(err), nil
	}

	return &tender.Bytes{Value: plaintext}, nil
}

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)

	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)

	return text, nil
}
