package stdlib

import (
	"bytes"
	"encoding/gob"
	"github.com/2dprototype/tender"
)


var gobModule = map[string]tender.Object{
	"encode": &tender.NativeFunction{Name: "encode", Value: gobEncode},
	"decode": &tender.NativeFunction{Name: "decode", Value: gobDecode},
}

type Gob struct {
	Value tender.Object
}

func gobEncode(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
    var buffer bytes.Buffer
    enc := gob.NewEncoder(&buffer)
	enc.Encode(Gob{Value: args[0]})
	return &tender.Bytes{Value: buffer.Bytes()}, nil
}

func gobDecode(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	var decoded Gob
	b, _ := tender.ToByteSlice(args[0])
	byteBuffer := bytes.NewReader(b)
	decoder := gob.NewDecoder(byteBuffer)
	decoder.Decode(&decoded)
	return decoded.Value, nil
}