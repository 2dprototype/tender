package stdlib

import (
	"fmt"

	"github.com/2dprototype/tender"
)

var fmtModule = map[string]tender.Object{
	"fprint":   &tender.NativeFunction{Name: "fprint",  Value: fmtFprint},
	"fprintln": &tender.NativeFunction{Name: "fprint",  Value: fmtFprintln},
	"print":    &tender.NativeFunction{Name: "print",   Value: fmtPrint},
	"printf":   &tender.NativeFunction{Name: "printf",  Value: fmtPrintf},
	"println":  &tender.NativeFunction{Name: "println", Value: fmtPrintln},
	"sprintf":  &tender.NativeFunction{Name: "sprintf", Value: fmtSprintf},
	"scanln":   &tender.NativeFunction{Name: "scanln",  Value: fmtScanln},
}

func fmtFprintln(args ...tender.Object) (ret tender.Object, err error) {
	numArgs := len(args)
	if numArgs < 2 {
		return nil, tender.ErrWrongNumArguments
	}
	
	writer, ok := args[0].(*IOWriter)
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "io.writer",
			Found:    args[0].TypeName(),
		}
	}
	str := ""
	for i, arg := range args[1:] {
		s, _ := tender.ToStringFormated(arg)
		str += s
		if i < len(args) - 1 {
			str += " "
		}
	}
	fmt.Fprintln(writer.Value, str)
	return nil, nil
}

func fmtFprint(args ...tender.Object) (ret tender.Object, err error) {
	numArgs := len(args)
	if numArgs < 2 {
		return nil, tender.ErrWrongNumArguments
	}
	
	writer, ok := args[0].(*IOWriter)
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "io.writer",
			Found:    args[0].TypeName(),
		}
	}
	str := ""
	for i, arg := range args[1:] {
		s, _ := tender.ToStringFormated(arg)
		str += s
		if i < len(args) - 1 {
			str += " "
		}
	}
	fmt.Fprint(writer.Value, str)
	return nil, nil
}

func fmtPrint(args ...tender.Object) (ret tender.Object, err error) {
	str := ""
	for i, arg := range args {
		s, _ := tender.ToStringFormated(arg)
		str += s
		if i < len(args) - 1 {
			str += " "
		}
	}
	
	fmt.Print(str)
	return nil, nil
	// i, err := fmt.Print(str)
	// if err != nil {
		// return wrapError(err), nil
	// }
	// return &tender.Int{Value: int64(i)}, nil
}

func fmtPrintln(args ...tender.Object) (ret tender.Object, err error) {
	str := ""
	for i, arg := range args {
		s, _ := tender.ToStringFormated(arg)
		str += s
		if i < len(args) - 1 {
			str += " "
		}
	}
	
	fmt.Println(str)
	return nil, nil
}

func fmtPrintf(args ...tender.Object) (ret tender.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, tender.ErrWrongNumArguments
	}

	format, ok := args[0].(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		fmt.Print(format)
		return nil, nil
	}

	s, err := tender.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	fmt.Print(s)
	return nil, nil
}

func fmtSprintf(args ...tender.Object) (ret tender.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, tender.ErrWrongNumArguments
	}

	format, ok := args[0].(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		// okay to return 'format' directly as String is immutable
		return format, nil
	}
	s, err := tender.Format(format.Value, args[1:]...)
	if err != nil {
		return wrapError(err), nil
	}
	return &tender.String{Value: s}, nil
}

func fmtScanln(args ...tender.Object) (ret tender.Object, err error) {
	var input string
	_, err = fmt.Scanln(&input)
	if err != nil {
		return wrapError(err), nil
	}
	return &tender.String{Value: input}, nil
}


// func getPrintArgs(args ...tender.Object) ([]interface{}, error) {
	// var printArgs []interface{}
	// l := 0
	// for _, arg := range args {
		// s, _ := tender.ToString(arg)
		// slen := len(s)
		// // make sure length does not exceed the limit
		// if l+slen > tender.MaxStringLen {
			// return nil, tender.ErrStringLimit
		// }
		// l += slen
		// printArgs = append(printArgs, s)
	// }
	// return printArgs, nil
// }
