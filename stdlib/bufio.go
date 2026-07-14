package stdlib

import (
	"bufio"
	"os"
	"github.com/2dprototype/tender"
)

var bufioModule = map[string]tender.Object{
	"readline": &tender.NativeFunction{Name: "readline", Value: bufioReadline},
	"readstring": &tender.NativeFunction{Name: "readstring", Value: bufioReadString},
	"readbytes": &tender.NativeFunction{Name: "readbytes", Value: bufioReadBytes},
}

func bufioReadline(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 0 {
		return nil, tender.ErrWrongNumArguments
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return wrapError(err), nil
	}

	// Remove the trailing newline character
	input = input[:len(input)-1]

	return &tender.String{Value: input}, nil
}

func bufioReadString(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(args[0].(*tender.String).Value[0])
	if err != nil {
		return wrapError(err), nil
	}

	// Remove the delimiter character
	input = input[:len(input)-1]

	return &tender.String{Value: input}, nil
}

func bufioReadBytes(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	reader := bufio.NewReader(os.Stdin)
	// Read a specified number of bytes
	data := make([]byte, args[0].(*tender.Int).Value)
	_, readErr := reader.Read(data)
	if readErr != nil {
		return wrapError(err), nil
	}

	return &tender.Bytes{Value: data}, nil
}