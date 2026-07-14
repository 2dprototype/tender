package tender

import "fmt"
import "sort"
import "time"
import "os"
import "unsafe"
import "github.com/2dprototype/tender/v/colorable"
import "math/big"

var builtinFuncs []*NativeFunction

// if needVMObj is true, VM will pass [VMObj, args...] to fn when calling it.
func addBuiltinFunction(name string, fn CallableFunc, needVMObj bool) {
	builtinFuncs = append(builtinFuncs, &NativeFunction{Name: name, Value: fn, NeedVMObj: needVMObj})
}

func init() {
	addBuiltinFunction("pointer", builtinPointer, true)
	addBuiltinFunction("deref", builtinDeref, false)
	addBuiltinFunction("set", builtinSet, false)
	addBuiltinFunction("is_pointer", builtinIsPointer, false)
	addBuiltinFunction("debug", builtinDebug, true)
	addBuiltinFunction("sysout", builtinSysout, false)
	addBuiltinFunction("print", builtinPrint, false)
	addBuiltinFunction("println", builtinPrintln, false)
	addBuiltinFunction("reverse", builtinReverse, false)
	addBuiltinFunction("includes", builtinIncludes, false)
	addBuiltinFunction("indexof", builtinIndexOf, false)
	addBuiltinFunction("lastindexof", builtinLastIndexOf, false)
	addBuiltinFunction("cap", builtinCap, false)
	addBuiltinFunction("len", builtinLen, false)
	addBuiltinFunction("copy", builtinCopy, false)
	addBuiltinFunction("append", builtinAppend, false)
	addBuiltinFunction("delete", builtinDelete, false)
	addBuiltinFunction("splice", builtinSplice, false)
	addBuiltinFunction("sort", builtinSort, false)
	addBuiltinFunction("rune", builtinRune, false)
	addBuiltinFunction("string", builtinString, false)
	addBuiltinFunction("int", builtinInt, false)
	addBuiltinFunction("bigint", builtinBigint, false)
	addBuiltinFunction("bool", builtinBool, false)
	addBuiltinFunction("float", builtinFloat, false)
	addBuiltinFunction("bigfloat", builtinBigFloat, false)
	addBuiltinFunction("complex", builtinComplex, false)
	addBuiltinFunction("char", builtinChar, false)
	addBuiltinFunction("bytes", builtinBytes, false)
	addBuiltinFunction("time", builtinTime, false)
	addBuiltinFunction("is_cycle", builtinIsCycle, false)
	addBuiltinFunction("is_int", builtinIsInt, false)
	addBuiltinFunction("is_float", builtinIsFloat, false)
	addBuiltinFunction("is_bigint", builtinIsBigInt, false)
	addBuiltinFunction("is_bigfloat", builtinIsBigFloat, false)
	addBuiltinFunction("is_complex", builtinIsComplex, false)
	addBuiltinFunction("is_string", builtinIsString, false)
	addBuiltinFunction("is_bool", builtinIsBool, false)
	addBuiltinFunction("is_char", builtinIsChar, false)
	addBuiltinFunction("is_bytes", builtinIsBytes, false)
	addBuiltinFunction("is_array", builtinIsArray, false)
	addBuiltinFunction("is_immutable_array", builtinIsImmutableArray, false)
	addBuiltinFunction("is_map", builtinIsMap, false)
	addBuiltinFunction("is_immutable_map", builtinIsImmutableMap, false)
	addBuiltinFunction("is_iterable", builtinIsIterable, false)
	addBuiltinFunction("is_time", builtinIsTime, false)
	addBuiltinFunction("is_error", builtinIsError, false)
	addBuiltinFunction("is_null", builtinIsNull, false)
	addBuiltinFunction("is_function", builtinIsFunction, false)
	addBuiltinFunction("is_callable", builtinIsCallable, false)
	addBuiltinFunction("typeof", builtinTypeOf, false)
	addBuiltinFunction("format", builtinFormat, false)
	addBuiltinFunction("is_tuple", builtinIsTuple, false)
	addBuiltinFunction("is_struct", builtinIsStruct, false)
	addBuiltinFunction("range",  builtinRange, false)
	addBuiltinFunction("freeze",  builtinFreeze, false)
}

// GetAllBuiltinFunctions returns all builtin function objects.
func GetAllBuiltinFunctions() []*NativeFunction {
	return append([]*NativeFunction{}, builtinFuncs...)
}



// Pointer builtins
func builtinPointer(args ...Object) (Object, error) {
    if len(args) != 2 {
        return nil, ErrWrongNumArguments
    }

    // first argument is VM object (because needVMObj = true)
    vm := args[0].(*VMObj).Value
    arg := args[1]

    // find the variable slot reference in the environment
    name, ok := vm.FindGlobalIndexByValue(arg)
    if !ok {
        // fallback to normal behavior (value only)
        slot := &arg
        return &Pointer{Slot: slot, Address: uintptr(unsafe.Pointer(slot))}, nil
    }

    slot := vm.GetGlobalSlotPointer(name)
    return &Pointer{
        Slot:    slot,
        Address: uintptr(unsafe.Pointer(slot)),
    }, nil
}


func builtinDeref(args ...Object) (Object, error) {
    if len(args) != 1 {
        return nil, ErrWrongNumArguments
    }

    p, ok := args[0].(*Pointer)
    if !ok {
        return nil, ErrInvalidArgumentType{
            Name:     "first",
            Expected: "pointer",
            Found:    args[0].TypeName(),
        }
    }
    if p.Slot == nil || *p.Slot == nil {
        return NullValue, nil
    }
    return *p.Slot, nil
}

func builtinSet(args ...Object) (Object, error) {
    if len(args) != 2 {
        return nil, ErrWrongNumArguments
    }

    p, ok := args[0].(*Pointer)
    if !ok {
        return nil, ErrInvalidArgumentType{
            Name:     "first",
            Expected: "pointer",
            Found:    args[0].TypeName(),
        }
    }
    if p.Slot == nil {
        return wrapError(fmt.Errorf("null pointer assignment")), nil
    }

    *p.Slot = args[1]
    return args[1], nil
}

func builtinIsPointer(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Pointer); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}


func builtinDebug(args ...Object) (Object, error) {
	vm := args[0].(*VMObj).Value
	args = args[1:] 
    mode, err := os.Stdout.Stat()
	if err == nil {
		if mode.Mode()&os.ModeCharDevice != 0 {
			str := ""
			for i, arg := range args {
				str += ToStringPrettyColored(vm, arg)
				if i < len(args) - 1 {
					str += " "
				}
			}
			fmt.Fprintln(colorable.NewColorableStdout(), str)
			return nil, nil
		}
	}
	str := ""
	for i, arg := range args {
		str += ToStringPretty(vm, arg)
		if i < len(args) - 1 {
			str += " "
		}
	}
	fmt.Println(str)
	return nil, nil
}

func builtinSysout(args ...Object) (Object, error) {
	str := ""
	for _, arg := range args {
		s, _ := ToStringFormated(arg)
		str += s
	}
	fmt.Fprint(colorable.NewColorableStdout(), str)
	return nil, nil
}

func builtinPrint(args ...Object) (Object, error) {
	str := ""
	for i, arg := range args {
		s, _ := ToStringFormated(arg)
		str += s
		if i < len(args) - 1 {
			str += " "
		}
	}
	fmt.Print(str)
	return nil, nil
}

func builtinPrintln(args ...Object) (Object, error) {
	str := ""
	for i, arg := range args {
		s, _ := ToStringFormated(arg)
		str += s
		if i < len(args) - 1 {
			str += " "
		}
	}
	fmt.Println(str)
	return nil, nil
}

func builtinReverse(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}

	switch container := args[0].(type) {
	case *Array:
		reversed := &Array{Value: make([]Object, len(container.Value))}
		for i, j := 0, len(container.Value)-1; i < len(container.Value); i, j = i+1, j-1 {
			reversed.Value[i] = container.Value[j]
		}
		return reversed, nil

	case *String:
		runes := []rune(container.Value)
		reversed := make([]rune, len(runes))
		for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
			reversed[i] = runes[j]
		}
		return &String{Value: string(reversed)}, nil
	case *Bytes:
		runes := []byte(container.Value)
		reversed := make([]byte, len(runes))
		for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
			reversed[i] = runes[j]
		}
		return &Bytes{Value: reversed}, nil

	default:
		return nil, nil
	}
}

func builtinLastIndexOf(args ...Object) (Object, error) {
	if len(args) != 2 {
		return nil, ErrWrongNumArguments
	}

	switch container := args[0].(type) {
		case *Array:
			element := args[1]
			for i := len(container.Value) - 1; i >= 0; i-- {
				if container.Value[i].Equals(element) {
					return &Int{Value: int64(i)}, nil
				}
			}
			return &Int{Value: -1}, nil
		case *ImmutableArray:
			element := args[1]
			for i := len(container.Value) - 1; i >= 0; i-- {
				if container.Value[i].Equals(element) {
					return &Int{Value: int64(i)}, nil
				}
			}
			return &Int{Value: -1}, nil

		case *String:
			char, ok := args[1].(*Char)
			if !ok {
				return nil, ErrInvalidArgumentType{
					Name:     "second",
					Expected: "char",
					Found:    args[1].TypeName(),
				}
			}
			str := container.Value
			chr := char.Value
			for i := len(str) - 1; i >= 0; i-- {
				if rune(str[i]) == chr {
					return &Int{Value: int64(i)}, nil
				}
			}
			return &Int{Value: -1}, nil

		default:
			return nil, ErrInvalidArgumentType{
				Name:     "first",
				Expected: "array or string",
				Found:    args[0].TypeName(),
			}
	}
}



func builtinIndexOf(args ...Object) (Object, error) {
	if len(args) != 2 {
		return nil, ErrWrongNumArguments
	}

	switch container := args[0].(type) {
		case *Array:
			for i, element := range container.Value {
				if element.Equals(args[1]) {
					return &Int{Value: int64(i)}, nil
				}
			}
			return &Int{Value: -1}, nil	
		case *ImmutableArray:
			for i, element := range container.Value {
				if element.Equals(args[1]) {
					return &Int{Value: int64(i)}, nil
				}
			}
			return &Int{Value: -1}, nil

		case *String:
			char, ok := args[1].(*Char)
			if !ok {
				return nil, ErrInvalidArgumentType{
					Name:     "second",
					Expected: "char",
					Found:    args[1].TypeName(),
				}
			}
			str := container.Value
			chr := char.Value
			for i, c := range str {
				if c == chr {
					return &Int{Value: int64(i)}, nil
				}
			}
			return &Int{Value: -1}, nil

		default:
			return nil, ErrInvalidArgumentType{
				Name:     "first",
				Expected: "array or string",
				Found:    args[0].TypeName(),
			}
	}
}


func builtinIncludes(args ...Object) (Object, error) {
	if len(args) != 2 {
		return nil, ErrWrongNumArguments
	}

	switch container := args[0].(type) {
		case *Map:
			key, ok := args[1].(*String)
			if !ok {
				return nil, ErrInvalidArgumentType{
					Name:     "second",
					Expected: "string",
					Found:    args[1].TypeName(),
				}
			}
			_, exists := container.Value[key.Value]
			if exists {
				return TrueValue, nil
			}
			return FalseValue, nil	
		case *ImmutableMap:
			key, ok := args[1].(*String)
			if !ok {
				return nil, ErrInvalidArgumentType{
					Name:     "second",
					Expected: "string",
					Found:    args[1].TypeName(),
				}
			}
			_, exists := container.Value[key.Value]
			if exists {
				return TrueValue, nil
			}
			return FalseValue, nil
		case *Array:
			for _, element := range container.Value {
				if element.Equals(args[1]) {
					return TrueValue, nil
				}
			}
			return FalseValue, nil	
		case *ImmutableArray:
			for _, element := range container.Value {
				if element.Equals(args[1]) {
					return TrueValue, nil
				}
			}
			return FalseValue, nil
		case *String:
			char, ok := args[1].(*Char)
			if !ok {
				return nil, ErrInvalidArgumentType{
					Name:     "second",
					Expected: "char",
					Found:    args[1].TypeName(),
				}
			}
			str := container.Value
			chr := char.Value
			for _, c := range str {
				if c == chr {
					return TrueValue, nil
				}
			}
			return FalseValue, nil

		default:
			return FalseValue, nil
	}
}


func builtinCap(args ...Object) (Object, error) {
    if len(args) != 1 {
        return nil, ErrWrongNumArguments
    }
	var n int
	switch v := args[0].(type) {
		case *Array:
			n = cap(v.Value)
		case *Bytes:
			n = cap(v.Value)
		default:
			return &Null{}, nil
	}
	return &Int{Value: int64(n)}, nil
}

func builtinSort(args ...Object) (Object, error) {
    if len(args) != 1 {
        return nil, ErrWrongNumArguments
    }

    switch arg := args[0].(type) {
		case *Array:
			arr := arg.Value
			sort.Slice(arr, func(i, j int) bool {
				a, b := arr[i], arr[j]
				// Try numeric comparison
				af, aok := ToFloat64(a)
				bf, bok := ToFloat64(b)
				if aok && bok {
					return af < bf
				}
				// Fallback to string comparison for non‑numeric types
				as, _ := ToString(a)
				bs, _ := ToString(b)
				return as < bs
			})
			return &Array{Value: arr}, nil

		case *String:
			runes := []rune(arg.Value)
			sort.Slice(runes, func(i, j int) bool {
				return runes[i] < runes[j]
			})
			return &String{Value: string(runes)}, nil
			
		case *Bytes:
			runes := arg.Value
			sort.Slice(runes, func(i, j int) bool {
				return runes[i] < runes[j]
			})
			return &Bytes{Value: runes}, nil
		default:
			return nil, nil
    }
}

func builtinRune(args ...Object) (Object, error) {
    if len(args) != 1 {
        return nil, ErrWrongNumArguments
    }

	switch arg := args[0].(type) {
		case *Char:
			return &Int{Value: int64(arg.Value)}, nil
		case *String:
			if len(arg.Value) == 1 {
				return &Int{Value: int64(arg.Value[0])}, nil
			} else {
				return &Int{Value: int64([]rune(arg.Value)[0])}, nil
			}
    }

    return nil, ErrInvalidArgumentType{
        Name:     "first",
        Expected: "char or string of length 1",
        Found:    args[0].TypeName(),
    }
}

func builtinTypeOf(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	return &String{Value: args[0].TypeName()}, nil
}

func builtinIsTuple(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Tuple); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsString(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*String); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsCycle(args ...Object) (Object, error) {
	
	visitedImmutableMaps := make(map[*ImmutableMap]bool)
    visitedImmutableArrays := make(map[*ImmutableArray]bool)
    visitedMaps := make(map[*Map]bool)
    visitedArrays := make(map[*Array]bool)
    traversingImmutableMaps := make(map[*ImmutableMap]bool)
    traversingImmutableArrays := make(map[*ImmutableArray]bool)
    traversingMaps := make(map[*Map]bool)
    traversingArrays := make(map[*Array]bool)
	visitedTuples    := make(map[*Tuple]bool)
	traversingTuples := make(map[*Tuple]bool)


    var dfsMap func(*Map, *Map) bool
    var dfsArray func(*Array, *Array) bool
    var dfsImmutableMap func(*ImmutableMap, *ImmutableMap) bool
    var dfsImmutableArray func(*ImmutableArray, *ImmutableArray) bool
    var dfsTuple func(*Tuple, *Tuple) bool
	
    dfsMap = func(currMap, parentMap *Map) bool {
        if traversingMaps[currMap] {
            return true // Cycle detected in map
        }
        if visitedMaps[currMap] {
            return false // Already visited, no cycle
        }
        traversingMaps[currMap] = true
        defer delete(traversingMaps, currMap)

        visitedMaps[currMap] = true
        for _, value := range currMap.Value {
            if m, ok := value.(*Map); ok {
                if dfsMap(m, currMap) {
                    return true
                }
            } else if m, ok := value.(*ImmutableMap); ok {
                if dfsImmutableMap(m, nil) {
                    return true
                }
            } else if a, ok := value.(*Array); ok {
                if dfsArray(a, nil) {
                    return true
                }
            } else if a, ok := value.(*ImmutableArray); ok {
                if dfsImmutableArray(a, nil) {
                    return true
                }
            }
        }
        return false
    }

    dfsArray = func(currArray, parentArray *Array) bool {
        if traversingArrays[currArray] {
            return true // Cycle detected in array
        }
        if visitedArrays[currArray] {
            return false // Already visited, no cycle
        }
        traversingArrays[currArray] = true
        defer delete(traversingArrays, currArray)

        visitedArrays[currArray] = true
        for _, value := range currArray.Value {
            if m, ok := value.(*Map); ok {
                if dfsMap(m, nil) {
                    return true
                }
            } else if a, ok := value.(*Array); ok {
                if dfsArray(a, currArray) {
                    return true
                }
            } else if m, ok := value.(*ImmutableMap); ok {
                if dfsImmutableMap(m, nil) {
                    return true
                }
            } else if a, ok := value.(*ImmutableArray); ok {
                if dfsImmutableArray(a, nil) {
                    return true
                }
            }
        }
        return false
    }

    dfsImmutableMap = func(currImmutableMap, parentImmutableMap *ImmutableMap) bool {
        if traversingImmutableMaps[currImmutableMap] {
            return true // Cycle detected in map
        }
        if visitedImmutableMaps[currImmutableMap] {
            return false // Already visited, no cycle
        }
        traversingImmutableMaps[currImmutableMap] = true
        defer delete(traversingImmutableMaps, currImmutableMap)

        visitedImmutableMaps[currImmutableMap] = true
        for _, value := range currImmutableMap.Value {
            if m, ok := value.(*ImmutableMap); ok {
                if dfsImmutableMap(m, currImmutableMap) {
                    return true
                }
            } else if a, ok := value.(*ImmutableArray); ok {
                if dfsImmutableArray(a, nil) {
                    return true
                }
            } else if m, ok := value.(*Map); ok {
                if dfsMap(m, nil) {
                    return true
                }
            } else if a, ok := value.(*Array); ok {
                if dfsArray(a, nil) {
                    return true
                }
            }
        }
        return false
    }

    dfsImmutableArray = func(currImmutableArray, parentImmutableArray *ImmutableArray) bool {
        if traversingImmutableArrays[currImmutableArray] {
            return true // Cycle detected in array
        }
        if visitedImmutableArrays[currImmutableArray] {
            return false // Already visited, no cycle
        }
        traversingImmutableArrays[currImmutableArray] = true
        defer delete(traversingImmutableArrays, currImmutableArray)

        visitedImmutableArrays[currImmutableArray] = true
        for _, value := range currImmutableArray.Value {
            if m, ok := value.(*ImmutableMap); ok {
                if dfsImmutableMap(m, nil) {
                    return true
                }
            } else if a, ok := value.(*ImmutableArray); ok {
                if dfsImmutableArray(a, currImmutableArray) {
                    return true
                }
            } else if m, ok := value.(*Map); ok {
                if dfsMap(m, nil) {
                    return true
                }
            } else if a, ok := value.(*Array); ok {
                if dfsArray(a, nil) {
                    return true
                }
            }
        }
        return false
    }

	dfsTuple = func(currTuple, parentTuple *Tuple) bool {
		if traversingTuples[currTuple] {
			return true // cycle detected
		}
		if visitedTuples[currTuple] {
			return false // already visited, no cycle
		}
		traversingTuples[currTuple] = true
		defer delete(traversingTuples, currTuple)

		visitedTuples[currTuple] = true
		for _, value := range currTuple.Value {
			// recursively check all contained containers
			switch v := value.(type) {
			case *Map:
				if dfsMap(v, nil) {
					return true
				}
			case *ImmutableMap:
				if dfsImmutableMap(v, nil) {
					return true
				}
			case *Array:
				if dfsArray(v, nil) {
					return true
				}
			case *ImmutableArray:
				if dfsImmutableArray(v, nil) {
					return true
				}
			case *Tuple:
				if dfsTuple(v, currTuple) {
					return true
				}
			}
		}
		return false
	}

    switch obj := args[0].(type) {
    case *Map:
        if dfsMap(obj, nil) {
            return TrueValue, nil
        }
    case *Array:
        if dfsArray(obj, nil) {
            return TrueValue, nil
        }
    case *ImmutableMap:
        if dfsImmutableMap(obj, nil) {
            return TrueValue, nil
        }
    case *ImmutableArray:
        if dfsImmutableArray(obj, nil) {
            return TrueValue, nil
        }
    case *Tuple:
        if dfsTuple(obj, nil) {
            return TrueValue, nil
        }
    default:
        return nil, ErrInvalidArgumentType{
            Name:     "first",
            Expected: "map, array, immutable-map, immutable-array, tuple",
            Found:    args[0].TypeName(),
        }
    }

    return FalseValue, nil
}

// func builtinIsCycle(args ...Object) (Object, error) {
    // if len(args) != 1 {
        // return nil, ErrWrongNumArguments
    // }

    // visitedMaps := make(map[*Map]bool)
    // visitedArrays := make(map[*Array]bool)
    // traversingMaps := make(map[*Map]bool)
    // traversingArrays := make(map[*Array]bool)

    // var dfsMap func(*Map, *Map) bool
    // var dfsArray func(*Array, *Array) bool

    // dfsMap = func(currMap, parentMap *Map) bool {
        // if traversingMaps[currMap] {
            // return true // Cycle detected in map
        // }
        // if visitedMaps[currMap] {
            // return false // Already visited, no cycle
        // }
        // traversingMaps[currMap] = true
        // defer delete(traversingMaps, currMap)

        // visitedMaps[currMap] = true
        // for _, value := range currMap.Value {
            // if m, ok := value.(*Map); ok {
                // if dfsMap(m, currMap) {
                    // return true
                // }
            // } else if a, ok := value.(*Array); ok {
                // if dfsArray(a, nil) {
                    // return true
                // }
            // }
        // }
        // return false
    // }

    // dfsArray = func(currArray, parentArray *Array) bool {
        // if traversingArrays[currArray] {
            // return true // Cycle detected in array
        // }
        // if visitedArrays[currArray] {
            // return false // Already visited, no cycle
        // }
        // traversingArrays[currArray] = true
        // defer delete(traversingArrays, currArray)

        // visitedArrays[currArray] = true
        // for _, value := range currArray.Value {
            // if m, ok := value.(*Map); ok {
                // if dfsMap(m, nil) {
                    // return true
                // }
            // } else if a, ok := value.(*Array); ok {
                // if dfsArray(a, currArray) {
                    // return true
                // }
            // }
        // }
        // return false
    // }

    // switch obj := args[0].(type) {
    // case *Map:
        // if dfsMap(obj, nil) {
            // return TrueValue, nil
        // }
    // case *Array:
        // if dfsArray(obj, nil) {
            // return TrueValue, nil
        // }
    // default:
        // return nil, ErrInvalidArgumentType{
            // Name:     "first",
            // Expected: "map or array",
            // Found:    args[0].TypeName(),
        // }
    // }

    // return FalseValue, nil
// }

func builtinIsInt(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Int); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsFloat(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Float); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsBigInt(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*BigInt); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsBigFloat(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*BigFloat); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsComplex(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Complex); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsBool(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Bool); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsChar(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Char); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsBytes(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Bytes); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsArray(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Array); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsImmutableArray(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*ImmutableArray); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsMap(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Map); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsImmutableMap(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*ImmutableMap); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsTime(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Time); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsError(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Error); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsNull(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if args[0] == NullValue {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsFunction(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	switch args[0].(type) {
	case *Function:
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsCallable(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if args[0].CanCall() {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsIterable(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if args[0].CanIterate() {
		return TrueValue, nil
	}
	return FalseValue, nil
}

// len(obj object) => int
func builtinLen(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *Array:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *ImmutableArray:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *String:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *Bytes:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *Map:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *ImmutableMap:
		return &Int{Value: int64(len(arg.Value))}, nil
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array/string/bytes/map",
			Found:    arg.TypeName(),
		}
	}
}

//range(start, stop[, step])
func builtinRange(args ...Object) (Object, error) {
	numArgs := len(args)
	if numArgs < 2 || numArgs > 3 {
		return nil, ErrWrongNumArguments
	}
	var start, stop, step *Int

	for i, arg := range args {
		v, ok := args[i].(*Int)
		if !ok {
			var name string
			switch i {
			case 0:
				name = "start"
			case 1:
				name = "stop"
			case 2:
				name = "step"
			}

			return nil, ErrInvalidArgumentType{
				Name:     name,
				Expected: "int",
				Found:    arg.TypeName(),
			}
		}
		if i == 2 && v.Value <= 0 {
			return nil, ErrInvalidRangeStep
		}
		switch i {
		case 0:
			start = v
		case 1:
			stop = v
		case 2:
			step = v
		}
	}

	if step == nil {
		step = &Int{Value: int64(1)}
	}

	return buildRange(start.Value, stop.Value, step.Value), nil
}

func buildRange(start, stop, step int64) *Array {
	array := &Array{}
	if start <= stop {
		for i := start; i < stop; i += step {
			array.Value = append(array.Value, &Int{
				Value: i,
			})
		}
	} else {
		for i := start; i > stop; i -= step {
			array.Value = append(array.Value, &Int{
				Value: i,
			})
		}
	}
	return array
}

func builtinFreeze(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	return freezeObject(args[0], make(map[Object]Object)), nil
}

// freezeObject recursively converts mutable containers to their immutable
// equivalents. memo prevents infinite recursion on self-referential structures.
func freezeObject(o Object, memo map[Object]Object) Object {
	switch v := o.(type) {
	case *Array:
		if cached, ok := memo[o]; ok {
			return cached
		}
		frozen := &ImmutableArray{Value: make([]Object, len(v.Value))}
		memo[o] = frozen
		for i, elem := range v.Value {
			frozen.Value[i] = freezeObject(elem, memo)
		}
		return frozen

	case *Map:
		if cached, ok := memo[o]; ok {
			return cached
		}
		frozen := &ImmutableMap{Value: make(map[string]Object, len(v.Value))}
		memo[o] = frozen
		for k, val := range v.Value {
			frozen.Value[k] = freezeObject(val, memo)
		}
		return frozen

	case *ImmutableArray:
		// Re-freeze elements in case they contain mutable values.
		newElems := make([]Object, len(v.Value))
		changed := false
		for i, elem := range v.Value {
			f := freezeObject(elem, memo)
			newElems[i] = f
			if f != elem {
				changed = true
			}
		}
		if !changed {
			return o
		}
		return &ImmutableArray{Value: newElems}

	case *ImmutableMap:
		newMap := make(map[string]Object, len(v.Value))
		changed := false
		for k, val := range v.Value {
			f := freezeObject(val, memo)
			newMap[k] = f
			if f != val {
				changed = true
			}
		}
		if !changed {
			return o
		}
		return &ImmutableMap{Value: newMap}

	default:
		// Primitives, strings, bytes, time, functions, errors — return as-is.
		return o
	}
}

func builtinFormat(args ...Object) (Object, error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, ErrWrongNumArguments
	}
	format, ok := args[0].(*String)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		// okay to return 'format' directly as String is immutable
		return format, nil
	}
	s, err := Format(format.Value, args[1:]...)
	if err != nil {
		return wrapError(err), nil
	}
	return &String{Value: s}, nil
}

func builtinCopy(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	return args[0].Copy(), nil
}

func builtinString(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen == 0 {
		return &String{}, nil
	}
	if argsLen != 1 {
		return nil, ErrWrongNumArguments
	}
	v, _ := ToString(args[0])
	return &String{Value: v}, nil
}

func builtinInt(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen == 0 {
		return &Int{}, nil
	}
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Int); ok {
		return args[0], nil
	}
	v, ok := ToInt64(args[0])
	if ok {
		return &Int{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return NullValue, nil
}


func builtinBigint(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen == 0 {
		return &BigInt{Value: new(big.Int)}, nil
	}
	if argsLen != 1 {
		return nil, ErrWrongNumArguments
	}
	bi, ok := ToBigInt(args[0])
	if ok {
		return &BigInt{Value: bi}, nil
	}	
	return NullValue, nil
}

func builtinFloat(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen == 0 {
		return &Float{}, nil
	}
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Float); ok {
		return args[0], nil
	}
	v, ok := ToFloat64(args[0])
	if ok {
		return &Float{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return NullValue, nil
}


func builtinBigFloat(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen == 0 {
		return &BigFloat{Value: new(big.Float)}, nil
	}
	if argsLen != 1 {
		return nil, ErrWrongNumArguments
	}
	bf, ok := ToBigFloat(args[0])
	if ok {
		return &BigFloat{Value: bf}, nil
	}	
	return NullValue, nil
}

func builtinComplex(args ...Object) (Object, error) {
	if len(args) != 2 {
		return nil, ErrWrongNumArguments
	}
	i1, _ := ToFloat64(args[0])
	i2, _ := ToFloat64(args[1])
	return &Complex{Value: complex(i1, i2)}, nil
}

func builtinBool(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Bool); ok {
		return args[0], nil
	}
	v, ok := ToBool(args[0])
	if ok {
		if v {
			return TrueValue, nil
		}
		return FalseValue, nil
	}
	return NullValue, nil
}

func builtinChar(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen == 0 {
		return &Char{}, nil
	}
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Char); ok {
		return args[0], nil
	}
	v, ok := ToRune(args[0])
	if ok {
		return &Char{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return NullValue, nil
}
func builtinBytes(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen == 0 {
		return &Bytes{}, nil
	}
	if argsLen > 1 {
		var concatenatedBytes []byte
		for _, obj := range args {
			b, ok := ToByteSlice(obj)
			if !ok {
				return nil, ErrInvalidArgumentType{
					Name:     "argument",
					Expected: "bytes",
					Found:    obj.TypeName(),
				}
			}
			concatenatedBytes = append(concatenatedBytes, b...)
		}
		if len(concatenatedBytes) > MaxBytesLen {
			return nil, ErrBytesLimit
		}
		return &Bytes{Value: concatenatedBytes}, nil
	}

	// If the first argument is not an array and there are no additional arguments,
	// handle the single argument case as before
	if n, ok := args[0].(*Int); ok {
		if n.Value > int64(MaxBytesLen) {
			return nil, ErrBytesLimit
		}
		return &Bytes{Value: make([]byte, int(n.Value))}, nil
	}
	v, ok := ToByteSlice(args[0])
	if ok {
		if len(v) > MaxBytesLen {
			return nil, ErrBytesLimit
		}
		return &Bytes{Value: v}, nil
	}
	return NullValue, nil
}


func builtinTime(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen == 0 {
		return &Time{Value: time.Now()}, nil
	}
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Time); ok {
		return args[0], nil
	}
	v, ok := ToTime(args[0])
	if ok {
		return &Time{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return NullValue, nil
}

// append(arr, items...)
func builtinAppend(args ...Object) (Object, error) {
	if len(args) < 2 {
		return nil, ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *Array:
		return &Array{Value: append(arg.Value, args[1:]...)}, nil
	case *ImmutableArray:
		return &ImmutableArray{Value: append(arg.Value, args[1:]...)}, nil
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array",
			Found:    arg.TypeName(),
		}
	}
}

// builtinDelete deletes Map keys
// usage: delete(map, "key")
// key must be a string
func builtinDelete(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen != 2 {
		return nil, ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *Map:
		if key, ok := args[1].(*String); ok {
			delete(arg.Value, key.Value)
			return NullValue, nil
		}
		return nil, ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string",
			Found:    args[1].TypeName(),
		}
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "map",
			Found:    arg.TypeName(),
		}
	}
}

// builtinSplice deletes and changes given Array, returns deleted items.
// usage:
// deleted_items := splice(array[,start[,delete_count[,item1[,item2[,...]]]])
func builtinSplice(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen == 0 {
		return nil, ErrWrongNumArguments
	}

	array, ok := args[0].(*Array)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array",
			Found:    args[0].TypeName(),
		}
	}
	arrayLen := len(array.Value)

	var startIdx int
	if argsLen > 1 {
		arg1, ok := args[1].(*Int)
		if !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int",
				Found:    args[1].TypeName(),
			}
		}
		startIdx = int(arg1.Value)
		if startIdx < 0 || startIdx > arrayLen {
			return nil, ErrIndexOutOfBounds
		}
	}

	delCount := len(array.Value)
	if argsLen > 2 {
		arg2, ok := args[2].(*Int)
		if !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int",
				Found:    args[2].TypeName(),
			}
		}
		delCount = int(arg2.Value)
		if delCount < 0 {
			return nil, ErrIndexOutOfBounds
		}
	}
	// if count of to be deleted items is bigger than expected, truncate it
	if startIdx+delCount > arrayLen {
		delCount = arrayLen - startIdx
	}
	// delete items
	endIdx := startIdx + delCount
	deleted := append([]Object{}, array.Value[startIdx:endIdx]...)

	head := array.Value[:startIdx]
	var items []Object
	if argsLen > 3 {
		items = make([]Object, 0, argsLen-3)
		for i := 3; i < argsLen; i++ {
			items = append(items, args[i])
		}
	}
	items = append(items, array.Value[endIdx:]...)
	array.Value = append(head, items...)

	// return deleted items
	return &Array{Value: deleted}, nil
}

func builtinIsStruct(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Struct); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}


// sortKeys sorts keys alphanumerically (numbers first, then letters)
func sortKeys(keys []string) {
	sort.Slice(keys, func(i, j int) bool {
		return naturalLess(keys[i], keys[j])
	})
}

// naturalLess compares two strings alphanumerically (0-9, a-z)
func naturalLess(a, b string) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if isDigit(a[i]) && isDigit(b[j]) {
			// Compare numeric values
			numA, lenA := parseNumber(a[i:])
			numB, lenB := parseNumber(b[j:])
			if numA != numB {
				return numA < numB
			}
			i += lenA
			j += lenB
		} else {
			// Compare characters case-insensitively
			ca := toLower(a[i])
			cb := toLower(b[j])
			if ca != cb {
				return ca < cb
			}
			i++
			j++
		}
	}
	return len(a) < len(b)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func parseNumber(s string) (int64, int) {
	var num int64
	i := 0
	for i < len(s) && isDigit(s[i]) {
		num = num*10 + int64(s[i]-'0')
		i++
	}
	return num, i
}
