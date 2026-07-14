package stdlib

import (
	"bytes"
	"compress/gzip"
	"github.com/2dprototype/tender"
)

var gzipModule = map[string]tender.Object{
	"compress":   &tender.NativeFunction{Name: "compress", Value: gzipCompress},
	"decompress": &tender.NativeFunction{Name: "decompress", Value: gzipDecompress},
}

func gzipCompress(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	input, _ := tender.ToByteSlice(args[0])

	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	_, err = writer.Write(input)
	if err != nil {
		return wrapError(err), nil
	}
	err = writer.Close()
	if err != nil {
		return wrapError(err), nil
	}

	return &tender.Bytes{Value: buf.Bytes()}, nil
}

func gzipDecompress(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	input, _ := tender.ToByteSlice(args[0])
	reader, err := gzip.NewReader(bytes.NewReader(input))
	if err != nil {
		return wrapError(err), nil
	}
	defer reader.Close()

	var decompressed bytes.Buffer
	_, err = decompressed.ReadFrom(reader)
	if err != nil {
		return nil, nil
	}

	return &tender.Bytes{Value: decompressed.Bytes()}, nil
}
