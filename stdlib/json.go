package stdlib

import (
	"bytes"
	gojson "encoding/json"

	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/stdlib/json"
)

var jsonModule = map[string]tender.Object{
	"decode": &tender.NativeFunction{
		Name:  "decode",
		Value: jsonDecode,
	},
	"encode": &tender.NativeFunction{
		Name:  "encode",
		Value: jsonEncode,
	},
	"indent": &tender.NativeFunction{
		Name:  "encode",
		Value: jsonIndent,
	},
	"html_escape": &tender.NativeFunction{
		Name:  "html_escape",
		Value: jsonHTMLEscape,
	},
}

func jsonDecode(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *tender.Bytes:
		v, err := json.Decode(o.Value)
		if err != nil {
			return &tender.Error{
				Value: &tender.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	case *tender.String:
		v, err := json.Decode([]byte(o.Value))
		if err != nil {
			return &tender.Error{
				Value: &tender.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	default:
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonEncode(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	b, err := json.Encode(args[0])
	if err != nil {
		return &tender.Error{Value: &tender.String{Value: err.Error()}}, nil
	}

	return &tender.Bytes{Value: b}, nil
}

func jsonIndent(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 3 {
		return nil, tender.ErrWrongNumArguments
	}

	prefix, ok := tender.ToString(args[1])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "prefix",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	indent, ok := tender.ToString(args[2])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "indent",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	switch o := args[0].(type) {
	case *tender.Bytes:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, o.Value, prefix, indent)
		if err != nil {
			return &tender.Error{
				Value: &tender.String{Value: err.Error()},
			}, nil
		}
		return &tender.Bytes{Value: dst.Bytes()}, nil
	case *tender.String:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, []byte(o.Value), prefix, indent)
		if err != nil {
			return &tender.Error{
				Value: &tender.String{Value: err.Error()},
			}, nil
		}
		return &tender.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonHTMLEscape(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *tender.Bytes:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, o.Value)
		return &tender.Bytes{Value: dst.Bytes()}, nil
	case *tender.String:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, []byte(o.Value))
		return &tender.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}
