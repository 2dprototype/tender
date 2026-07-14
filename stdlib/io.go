package stdlib

import (
	"io"
	"io/ioutil"
	"io/fs"
	"github.com/2dprototype/tender"
)

var ioModule = map[string]tender.Object{
	// "writer":    &IOWriter{},
	// "reader":    &IOReader{},
	"readfile":  &tender.NativeFunction{Value: ioReadFile},
	"writefile": &tender.NativeFunction{Value: ioWriteFile},
	"read_all":  &tender.NativeFunction{Value: ioReadAll},
	"read_full": &tender.NativeFunction{Value: ioReadFull},
}


type IOWriter struct {
    tender.ObjectImpl
    Value io.Writer 
}

// TypeName returns the name of the type.
func (p *IOWriter) TypeName() string {
    return "io.writer"
}

func (p *IOWriter) String() string {
    return "<io.writer>"
}

func (p *IOWriter) Copy() tender.Object {
    return &IOWriter{Value: p.Value}
}

func (o *IOWriter) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String) 
	if ok {
		if strIdx.Value == "write" {
			res = &tender.NativeFunction{
				Value: FuncAYRIE(o.Value.Write),
			}
		} 
	}
	return
}


type IOReader struct {
    tender.ObjectImpl
    Value io.Reader
}

func (p *IOReader) TypeName() string {
    return "io.reader"
}

func (p *IOReader) String() string {
    return "<io.reader>"
}

func (p *IOReader) Copy() tender.Object {
    return &IOReader{Value: p.Value}
}

func (o *IOReader) IndexGet(index tender.Object) (res tender.Object, err error) {
	strIdx, ok := index.(*tender.String) 
	if ok {
		if strIdx.Value == "read" {
			res = &tender.NativeFunction{
				Value: FuncAYRIE(o.Value.Read),
			}
		} 
	}
	return
}


func ioReadFile(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	
	filePath, _ := tender.ToString(args[0])
	
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return wrapError(err), nil
	}
	
	return &tender.String{Value: string(data)}, nil
}

func ioWriteFile(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 && len(args) != 3 {
		return nil, tender.ErrWrongNumArguments
	}
	filePath, _ := tender.ToString(args[0])
	content, _ := tender.ToByteSlice(args[1])
	mode := 0644
	if len(args) == 3 {
		mode, _ = tender.ToInt(args[2])
	}
	err = ioutil.WriteFile(filePath, content, fs.FileMode(mode))
	if err != nil {
		return wrapError(err), nil
	}
	return tender.NullValue, nil
}

func ioReadAll(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	reader, ok := args[0].(*IOReader)
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "io.reader",
			Found:    args[0].TypeName(),
		}
	}
	
	data, err := io.ReadAll(reader.Value)
	if err != nil {
		return wrapError(err), nil
	}
	
	return &tender.Bytes{Value: data}, nil
}

func ioReadFull(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}

	reader, ok := args[0].(*IOReader)
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "io.reader",
			Found:    args[0].TypeName(),
		}
	}
	
	buf, _ := tender.ToByteSlice(args[1])
	
	i, err := io.ReadFull(reader.Value, buf)
	if err != nil {
		return wrapError(err), nil
	}
	
	return &tender.Int{Value: int64(i)}, nil
}

