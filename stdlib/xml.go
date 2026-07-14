package stdlib

import (
	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/stdlib/xml"
)

var xmlModule = map[string]tender.Object{
	"decode": &tender.NativeFunction{
		Name:  "decode",
		Value: xmlDecode,
	},
	"encode": &tender.NativeFunction{
		Name:  "encode",
		Value: xmlEncode,
	},
	"escape": &tender.NativeFunction{
		Name:  "escape",
		Value: xmlEscape,
	},
	"unescape": &tender.NativeFunction{
		Name:  "unescape",
		Value: xmlUnescape,
	},
}

func xmlDecode(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *tender.Bytes:
		v, err := xml.Decode(o.Value)
		if err != nil {
			return &tender.Error{
				Value: &tender.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	case *tender.String:
		v, err := xml.Decode([]byte(o.Value))
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

func xmlEncode(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	b, err := xml.Encode(args[0])
	if err != nil {
		return &tender.Error{Value: &tender.String{Value: err.Error()}}, nil
	}

	return &tender.Bytes{Value: b}, nil
}

func xmlEscape(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	str, ok := tender.ToString(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	escaped := xml.EscapeString(str)
	return &tender.String{Value: escaped}, nil
}

func xmlUnescape(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	str, ok := tender.ToString(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	unescaped := xml.UnescapeString(str)
	return &tender.String{Value: unescaped}, nil
}