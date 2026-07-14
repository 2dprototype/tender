package stdlib

import (
	"bytes"
	"archive/tar"
	"io"
	"github.com/2dprototype/tender"
)

var tarModule = map[string]tender.Object{
	"writer": &tender.NativeFunction{Name: "writer", Value: tarNewWriter},
	"reader": &tender.NativeFunction{Name: "reader", Value: tarNewReader},
}

func tarNewWriter(args ...tender.Object) (tender.Object, error) {
	if len(args) != 0 {
		return nil, tender.ErrWrongNumArguments
	}

	var tarBuffer bytes.Buffer
	tarWriter := tar.NewWriter(&tarBuffer)

	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"create": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 {
						return nil, tender.ErrWrongNumArguments
					}
					filename, _ := tender.ToString(args[0])
					content, _ := tender.ToByteSlice(args[1])

					header := &tar.Header{
						Name: filename,
						Size: int64(len(content)),
					}

					if err := tarWriter.WriteHeader(header); err != nil {
						return wrapError(err), nil
					}

					if _, err := tarWriter.Write(content); err != nil {
						return wrapError(err), nil
					}

					return nil, nil
				},
			},
			"bytes": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					return &tender.Bytes{Value: tarBuffer.Bytes()}, nil
				},
			},
			"close": &tender.NativeFunction{
				Value: FuncARE(tarWriter.Close),
			},
		},
	}, nil
}

func tarNewReader(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	data, _ := tender.ToByteSlice(args[0])
	tarReader := tar.NewReader(bytes.NewReader(data))

	files := make([]tender.Object, 0)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return wrapError(err), nil
		}

		content, err := io.ReadAll(tarReader)
		if err != nil {
			return wrapError(err), nil
		}

		files = append(files, &tender.ImmutableMap{
			Value: map[string]tender.Object{
				"name": &tender.String{Value: header.Name},
				"mode": &tender.Int{Value: int64(header.Mode)},
				"size": &tender.Int{Value: header.Size},
				"data": &tender.Bytes{Value: content},
			},
		})
	}

	return &tender.Array{Value: files}, nil
}
