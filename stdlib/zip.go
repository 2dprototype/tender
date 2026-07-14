package stdlib

import (
	"bytes"
	"archive/zip"
	"io"
	"github.com/2dprototype/tender"
)

var zipModule = map[string]tender.Object{
	"writer":       &tender.NativeFunction{Name: "writer", Value: zipNewWriter},
	"reader":       &tender.NativeFunction{Name: "reader", Value: zipNewReader},
}

func zipNewWriter(args ...tender.Object) (tender.Object, error) {
	if len(args) != 0 {
		return nil, tender.ErrWrongNumArguments
	}
	
	var zipBuffer bytes.Buffer
	zipWriter := zip.NewWriter(&zipBuffer)

	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"create": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 {
						return nil, tender.ErrWrongNumArguments
					}
					filename, _ := tender.ToString(args[0])
					content, _ := tender.ToByteSlice(args[1])
					f, err := zipWriter.Create(filename)
					if err != nil {
						return wrapError(err), nil
					}
					
					// header := &zip.FileHeader{
						// Name: "JJ",
						// Comment: "MyMYYYYY",
					// }
					
					// zipWriter.CreateHeader(header)

					_, err = f.Write(content)
					if err != nil {
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
					return &tender.Bytes{Value: zipBuffer.Bytes()}, nil
				},
			},
			"close": &tender.NativeFunction{
				Value: FuncARE(zipWriter.Close),
			},	
			"flush": &tender.NativeFunction{
				Value: FuncARE(zipWriter.Flush),
			},
			"set_comment": &tender.NativeFunction{
				Value: FuncASRE(zipWriter.SetComment),
			},
			"set_offset": &tender.NativeFunction{
				Value: FuncAI64R(zipWriter.SetOffset),
			},
		},
	}, nil
}

func zipNewReader(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	data, _ := tender.ToByteSlice(args[0])
	reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return wrapError(err), nil
	}
	
	files := make([]tender.Object, len(reader.File))
	
	for i, file := range reader.File {
		files[i] = makeZipFile(file)
	}

	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"files" : &tender.Array{Value: files},
			"comment" : &tender.String{Value: reader.Comment},
		},
	}, nil
}


func makeZipFile(file *zip.File) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"name" : &tender.String{Value: file.Name},
			"comment" : &tender.String{Value: file.Comment},
			"non_utf8" : tender.FromBool(file.NonUTF8),
			"creator_version" : &tender.Int{Value: int64(file.CreatorVersion)},
			"reader_version" : &tender.Int{Value: int64(file.ReaderVersion)},
			"method" : &tender.Int{Value: int64(file.Method)},
			"modified" : &tender.Time{Value: file.Modified},
			"modified_time" : &tender.Int{Value: int64(file.ModifiedTime)},
			"modified_date" : &tender.Int{Value: int64(file.ModifiedDate)},
			"crc32" : &tender.Int{Value: int64(file.CRC32)},
			"compressed_size" : &tender.Int{Value: int64(file.CompressedSize64)},
			"uncompressed_size" : &tender.Int{Value: int64(file.UncompressedSize64)},
			"extra" : &tender.Bytes{Value: file.Extra},
			"read": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					fileReader, err := file.Open()
					if err != nil {
						return wrapError(err), nil
					}
					defer fileReader.Close()
					content, err := io.ReadAll(fileReader)
					if err != nil {
						return wrapError(err), nil
					}
					return &tender.Bytes{Value: content}, nil
				},
			},
		},
	}
}

