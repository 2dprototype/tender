package tender

import (
	"fmt"
	"strings"
)

// --- Array Methods ---
var arrayMethods = map[string]CallableFunc{
	"push": func(args ...Object) (Object, error) {
		arr := args[0].(*Array)
		arr.Value = append(arr.Value, args[1:]...)
		return &Int{Value: int64(len(arr.Value))}, nil
	},
	"pop": func(args ...Object) (Object, error) {
		arr := args[0].(*Array)
		if len(arr.Value) == 0 {
			return NullValue, nil
		}
		last := arr.Value[len(arr.Value)-1]
		arr.Value = arr.Value[:len(arr.Value)-1]
		return last, nil
	},
	"shift": func(args ...Object) (Object, error) {
		arr := args[0].(*Array)
		if len(arr.Value) == 0 {
			return NullValue, nil
		}
		first := arr.Value[0]
		arr.Value = arr.Value[1:]
		return first, nil
	},
	"unshift": func(args ...Object) (Object, error) {
		arr := args[0].(*Array)
		arr.Value = append(args[1:], arr.Value...)
		return &Int{Value: int64(len(arr.Value))}, nil
	},
	"join": func(args ...Object) (Object, error) {
		arr := args[0].(*Array)
		sep := ""
		if len(args) > 1 {
			s, ok := ToString(args[1])
			if !ok {
				return nil, fmt.Errorf("join() separator must be a string")
			}
			sep = s
		}
		var sParts []string
		for _, item := range arr.Value {
			sParts = append(sParts, item.String())
		}
		return &String{Value: strings.Join(sParts, sep)}, nil
	},
}

// --- ImmutableArray Methods ---
var immutableArrayMethods = map[string]CallableFunc{
	"join": func(args ...Object) (Object, error) {
		arr := args[0].(*ImmutableArray)
		sep := ""
		if len(args) > 1 {
			s, ok := ToString(args[1])
			if !ok {
				return nil, fmt.Errorf("join() separator must be a string")
			}
			sep = s
		}
		var sParts []string
		for _, item := range arr.Value {
			sParts = append(sParts, item.String())
		}
		return &String{Value: strings.Join(sParts, sep)}, nil
	},
}
