package tender

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"strings"
	
	"math/big"
)

var (
	// MaxStringLen is the maximum byte-length for string value. Note this
	// limit applies to all compiler/VM instances in the process.
	MaxStringLen = 2147483647

	// MaxBytesLen is the maximum length for bytes value. Note this limit
	// applies to all compiler/VM instances in the process.
	MaxBytesLen = 2147483647
)

const (
	// GlobalsSize is the maximum number of global variables for a VM.
	GlobalsSize = 10240

	// StackSize is the maximum stack size for a VM.
	StackSize = 20480

	// MaxFrames is the maximum number of function frames for a VM.
	MaxFrames = 10240

	// SourceFileExtDefault is the default extension for source files.
	SourceFileExtDefault = ".td"
)

// CallableFunc is a function signature for the callable functions.
type CallableFunc = func(args ...Object) (ret Object, err error)

// CountObjects returns the number of objects that a given object o contains.
// For scalar value types, it will always be 1. For compound value types,
// this will include its elements and all of their elements recursively.
func CountObjects(o Object) (c int) {
	c = 1
	switch o := o.(type) {
	case *Array:
		for _, v := range o.Value {
			c += CountObjects(v)
		}
	case *ImmutableArray:
		for _, v := range o.Value {
			c += CountObjects(v)
		}
	case *Map:
		for _, v := range o.Value {
			c += CountObjects(v)
		}
	case *ImmutableMap:
		for _, v := range o.Value {
			c += CountObjects(v)
		}
	case *Error:
		c += CountObjects(o.Value)
	case *Struct:
		for _, v := range o.Fields {
			c += CountObjects(v)
		}
	}
	return
}


// func ToString(o Object) string {
	// if str, isStr := o.(*String); isStr {
		// return str.Value
	// } else if byt, isByt := o.(*Bytes); isByt {
		// return string(byt.Value)
	// } else {
		// return o.String()
	// }
// }

// ToString will try to convert object o to string value.
func ToString(o Object) (v string, ok bool) {
	ok = true
	if str, isStr := o.(*String); isStr {
		v = str.Value
	} else if byt, isByt := o.(*Bytes); isByt {
		v = string(byt.Value)
	} else {
		v = o.String()
	}
	return
}


// ToString will try to convert object o to formated string value.
func ToStringFormated(o Object) (v string, ok bool) {
	ok = true
	if str, isStr := o.(*String); isStr {
		v = str.Value
	} else {
		v = o.String()
	}
	return
}


func ToByte(o Object) (v byte, ok bool) {
	switch o := o.(type) {
	case *Int:
		v = byte(o.Value)
		ok = true
	case *Float:
		v = byte(int(o.Value))
		ok = true
	case *Char:
		v = byte(o.Value)
		ok = true
	case *Bool:
		if o == TrueValue {
			v = 1
		}
		ok = true
	case *String:
		if len(o.Value) == 1 {
			v = o.Value[0]
			ok = true
		} else if len(o.Value) > 1 {
			c, err := strconv.ParseInt(o.Value, 10, 64)
			if err == nil {
				v = byte(c)
				ok = true
			}
		}
	}
	return
}

func ToStringPretty(vm *VM, o Object) string {
	m, ok := o.(*Map)
	if ok {
		if val, ok := m.Value["__print__"]; ok {
			fn, ok := val.(*CompiledFunction)
			if ok {
				vv, _ := WrapFuncCall(vm, fn)
				s, ok := vv.(*String)
				if ok {
					return s.Value
				} else {
					return vv.String()
				}
			}
		}
	}
	var builder strings.Builder
	visited := make(map[Object]bool)
	writeObjectPretty(&builder, o, 0, visited)
	return builder.String()
}

func writeObjectPretty(builder *strings.Builder, o Object, indentLevel int, visited map[Object]bool) {
	indent := strings.Repeat("  ", indentLevel)

	// Check for cycle
	if visited[o] {
		builder.WriteString("[Circular]")
		return
	}
	visited[o] = true

	switch obj := o.(type) {
	case *ImmutableMap:
		if len(obj.Value) == 0 {
			builder.WriteString("{}")
		} else {
			builder.WriteString("{\n")
			keys := make([]string, 0, len(obj.Value))
			for k := range obj.Value {
				keys = append(keys, k)
			}
			for i, k := range keys {
				v := obj.Value[k]
				builder.WriteString(indent + "  " + k + ": ")
				writeObjectPretty(builder, v, indentLevel+1, visited)
				if i < len(keys)-1 {
					builder.WriteString(",\n")
				}
			}
			builder.WriteString("\n" + indent + "}")
		}
	case *ImmutableArray:
		if len(obj.Value) == 0 {
			builder.WriteString("[]")
		} else {
			builder.WriteString("[\n")
			for i, elem := range obj.Value {
				builder.WriteString(indent + "  ")
				writeObjectPretty(builder, elem, indentLevel+1, visited)
				if i < len(obj.Value)-1 {
					builder.WriteString(",\n")
				}
			}
			builder.WriteString("\n" + indent + "]")
		}
	case *Map:
		if len(obj.Value) == 0 {
			builder.WriteString("{}")
		} else {
			builder.WriteString("{\n")
			keys := make([]string, 0, len(obj.Value))
			for k := range obj.Value {
				keys = append(keys, k)
			}
			for i, k := range keys {
				v := obj.Value[k]
				builder.WriteString(indent + "  " + k + ": ")
				writeObjectPretty(builder, v, indentLevel+1, visited)
				if i < len(keys)-1 {
					builder.WriteString(",\n")
				}
			}
			builder.WriteString("\n" + indent + "}")
		}
	case *Array:
		if len(obj.Value) == 0 {
			builder.WriteString("[]")
		} else {
			builder.WriteString("[\n")
			for i, elem := range obj.Value {
				builder.WriteString(indent + "  ")
				writeObjectPretty(builder, elem, indentLevel+1, visited)
				if i < len(obj.Value)-1 {
					builder.WriteString(",\n")
				}
			}
			builder.WriteString("\n" + indent + "]")
		}
	case *Tuple:
		if len(obj.Value) == 0 {
			builder.WriteString("()")
		} else {
			builder.WriteString("(")
			for i, elem := range obj.Value {
				writeObjectPretty(builder, elem, indentLevel, visited)
				if i < len(obj.Value)-1 {
					builder.WriteString(", ")
				}
			}
			if len(obj.Value) == 1 {
				builder.WriteString(",")
			}
			builder.WriteString(")")
		}
	case *Struct:
		if len(obj.Fields) == 0 {
			if obj.Type.Name != "" {
				builder.WriteString(obj.Type.Name + " {}")
			} else {
				builder.WriteString("{}")
			}
		} else {
			if obj.Type.Name != "" {
				builder.WriteString(obj.Type.Name + " {\n")
			} else {
				builder.WriteString("{\n")
			}
			for i, f := range obj.Type.Fields {
				val := obj.Fields[f.Name]
				if val == nil {
					val = NullValue
				}
				builder.WriteString(indent + "  " + f.Name + ": ")
				writeObjectPretty(builder, val, indentLevel+1, visited)
				if i < len(obj.Type.Fields)-1 {
					builder.WriteString(",\n")
				}
			}
			builder.WriteString("\n" + indent + "}")
		}
	case *Char:
		builder.WriteString("'" + obj.String() + "'")
	case *String:
		builder.WriteString(`"` + obj.Value + `"`)
	case *Bool:
		builder.WriteString(obj.String())
	default:
		builder.WriteString(obj.String())
	}

	delete(visited, o)
}

func ToStringPrettyColored(vm *VM, o Object) string {
	m, ok := o.(*Map)
	if ok {
		if val, ok := m.Value["__print__"]; ok {
			fn, ok := val.(*CompiledFunction)
			if ok {
				vv, _ := WrapFuncCall(vm, fn)
				s, ok := vv.(*String)
				if ok {
					return s.Value
				} else {
					return vv.String()
				}
			}
		}
	}
	var builder strings.Builder
	visited := make(map[Object]bool)
	writeObjectPrettyColored(&builder, o, 0, visited)
	return builder.String()
}

func writeObjectPrettyColored(builder *strings.Builder, o Object, indentLevel int, visited map[Object]bool) {
	indent := strings.Repeat("  ", indentLevel)

	// Check for cycle
	if visited[o] {
		builder.WriteString("\033[31m[Circular]\033[0m")
		return
	}
	visited[o] = true

	switch obj := o.(type) {
	case *ImmutableMap:
		if len(obj.Value) == 0 {
			builder.WriteString("{}")
		} else {
			builder.WriteString("{\n")
			keys := make([]string, 0, len(obj.Value))
			for k := range obj.Value {
				keys = append(keys, k)
			}
			for i, k := range keys {
				v := obj.Value[k]
				builder.WriteString(indent + "  " + "\033[36m" + k + "\033[0m: ")
				writeObjectPrettyColored(builder, v, indentLevel+1, visited)
				if i < len(keys)-1 {
					builder.WriteString(",\n")
				}
			}
			builder.WriteString("\n" + indent + "}")
		}
	case *ImmutableArray:
		if len(obj.Value) == 0 {
			builder.WriteString("[]")
		} else {
			builder.WriteString("[\n")
			for i, elem := range obj.Value {
				builder.WriteString(indent + "  ")
				writeObjectPrettyColored(builder, elem, indentLevel+1, visited)
				if i < len(obj.Value)-1 {
					builder.WriteString(",\n")
				}
			}
			builder.WriteString("\n" + indent + "]")
		}
	case *Map:
		if len(obj.Value) == 0 {
			builder.WriteString("{}")
		} else {
			builder.WriteString("{\n")
			keys := make([]string, 0, len(obj.Value))
			for k := range obj.Value {
				keys = append(keys, k)
			}
			for i, k := range keys {
				v := obj.Value[k]
				builder.WriteString(indent + "  " + "\033[36m" + k + "\033[0m: ")
				writeObjectPrettyColored(builder, v, indentLevel+1, visited)
				if i < len(keys)-1 {
					builder.WriteString(",\n")
				}
			}
			builder.WriteString("\n" + indent + "}")
		}
	case *Array:
		if len(obj.Value) == 0 {
			builder.WriteString("[]")
		} else {
			builder.WriteString("[\n")
			for i, elem := range obj.Value {
				builder.WriteString(indent + "  ")
				writeObjectPrettyColored(builder, elem, indentLevel+1, visited)
				if i < len(obj.Value)-1 {
					builder.WriteString(",\n")
				}
			}
			builder.WriteString("\n" + indent + "]")
		}
	case *Tuple:
		if len(obj.Value) == 0 {
			builder.WriteString("()")
		} else {
			builder.WriteString("(")
			for i, elem := range obj.Value {
				writeObjectPrettyColored(builder, elem, indentLevel, visited)
				if i < len(obj.Value)-1 {
					builder.WriteString(", ")
				}
			}
			if len(obj.Value) == 1 {
				builder.WriteString(",")
			}
			builder.WriteString(")")
		}
	case *Struct:
		if len(obj.Fields) == 0 {
			if obj.Type.Name != "" {
				builder.WriteString("\033[36m" + obj.Type.Name + "\033[0m {}")
			} else {
				builder.WriteString("{}")
			}
		} else {
			if obj.Type.Name != "" {
				builder.WriteString("\033[36m" + obj.Type.Name + "\033[0m {\n")
			} else {
				builder.WriteString("{\n")
			}
			for i, f := range obj.Type.Fields {
				val := obj.Fields[f.Name]
				if val == nil {
					val = NullValue
				}
				builder.WriteString(indent + "  " + "\033[36m" + f.Name + "\033[0m: ")
				writeObjectPrettyColored(builder, val, indentLevel+1, visited)
				if i < len(obj.Type.Fields)-1 {
					builder.WriteString(",\n")
				}
			}
			builder.WriteString("\n" + indent + "}")
		}
	case *String:
		builder.WriteString("\033[32m\"" + obj.Value + "\"\033[0m")
	case *Int, *Float, *Time, *BigInt, *BigFloat, *Complex:
		builder.WriteString("\033[33m" + obj.String() + "\033[0m")
	case *Char:
		builder.WriteString("\033[33m'" + obj.String() + "'\033[0m")
	case *Bool:
		builder.WriteString("\033[35m" + obj.String() + "\033[0m")
	case *Bytes, *UserFunction, *BuiltinFunction, *CompiledFunction:
		builder.WriteString("\033[36m" + obj.String() + "\033[0m")
	case *Null:
		builder.WriteString("\033[90mnull\033[0m")
	default:
		builder.WriteString(obj.String())
	}

	delete(visited, o)
}

// ToInt will try to convert object o to int32 value.
func ToInt32(o Object) (v int32, ok bool) {
	switch o := o.(type) {
	case *Int:
		v = int32(o.Value)
		ok = true
	case *Float:
		v = int32(o.Value)
		ok = true
	case *BigInt:
		v = int32(o.Value.Int64())
		ok = true	
	case *BigFloat:
		b, _ := o.Value.Int64()
		v = int32(b)
		ok = true
	case *Char:
		v = int32(o.Value)
		ok = true
	case *Bool:
		if o == TrueValue {
			v = 1
		}
		ok = true
	case *String:
		c, err := strconv.ParseInt(o.Value, 10, 64)
		if err == nil {
			v = int32(c)
			ok = true
		}
	}
	return
}


// ToInt will try to convert object o to int value.
func ToInt(o Object) (v int, ok bool) {
	switch o := o.(type) {
	case *Int:
		v = int(o.Value)
		ok = true
	case *Float:
		v = int(o.Value)
		ok = true
	case *BigInt:
		v = int(o.Value.Int64())
		ok = true
	case *BigFloat:
		b, _ := o.Value.Int64()
		v = int(b)
		ok = true
	case *Char:
		v = int(o.Value)
		ok = true
	case *Bool:
		if o == TrueValue {
			v = 1
		}
		ok = true
	case *String:
		c, err := strconv.ParseInt(o.Value, 10, 64)
		if err == nil {
			v = int(c)
			ok = true
		}
	}
	return
}

// ToInt64 will try to convert object o to int64 value.
func ToInt64(o Object) (v int64, ok bool) {
	switch o := o.(type) {
	case *Int:
		v = o.Value
		ok = true
	case *Float:
		v = int64(o.Value)
		ok = true
	case *BigInt:
		v = o.Value.Int64()
		ok = true
	case *BigFloat:
		v, _ = o.Value.Int64()
		ok = true
	case *Char:
		v = int64(o.Value)
		ok = true
	case *Bool:
		if o == TrueValue {
			v = 1
		}
		ok = true
	case *String:
		c, err := strconv.ParseInt(o.Value, 10, 64)
		if err == nil {
			v = c
			ok = true
		}
	}
	return
}

// ToBigInt will try to convert an object o to a *big.Int value.
func ToBigInt(o Object) (v *big.Int, ok bool) {
	switch o := o.(type) {
	case *Int:
		v = big.NewInt(o.Value)
		ok = true
	case *Float:
		v = new(big.Int)
		new(big.Float).SetFloat64(o.Value).Int(v)
		ok = true
	case *BigInt:
		v = new(big.Int).Set(o.Value)
		ok = true
	case *BigFloat:
		v = new(big.Int)
		o.Value.Int(v)
		ok = true
	case *Char:
		v = big.NewInt(int64(o.Value))
		ok = true
	case *Bool:
		if o == TrueValue {
			v = big.NewInt(1)
		} else {
			v = big.NewInt(0)
		}
		ok = true
	case *String:
		v, ok = new(big.Int).SetString(o.Value, 10)
	}
	return
}

// ToBigFloat will try to convert an object o to a *big.Float value.
func ToBigFloat(o Object) (v *big.Float, ok bool) {
	switch o := o.(type) {
	case *Int:
		v = new(big.Float).SetInt(big.NewInt(o.Value))
		ok = true
	case *Float:
		v = big.NewFloat(o.Value)
		ok = true
	case *BigInt:
		v = new(big.Float).SetInt(o.Value)
		ok = true
	case *BigFloat:
		v = new(big.Float).Set(o.Value)
		ok = true
	case *Char:
		v = new(big.Float).SetInt(big.NewInt(int64(o.Value)))
		ok = true
	case *Bool:
		if o == TrueValue {
			v = big.NewFloat(1)
		} else {
			v = big.NewFloat(0)
		}
		ok = true
	case *String:
		v, ok = new(big.Float).SetString(o.Value)
	}
	return
}


// ToUint8 will try to convert object o to uint8 value.
func ToUint8(o Object) (v uint8, ok bool) {
	switch o := o.(type) {
	case *Int:
		v = uint8(o.Value)
		ok = true
	case *Float:
		v = uint8(o.Value)
		ok = true
	case *BigInt:
		v = uint8(o.Value.Int64())
		ok = true
	case *BigFloat:
		b, _ := o.Value.Int64()
		v = uint8(b)
		ok = true
	case *Char:
		v = uint8(o.Value)
		ok = true
	case *Bool:
		if o == TrueValue {
			v = 1
		}
		ok = true
	case *String:
		c, err := strconv.ParseInt(o.Value, 10, 64)
		if err == nil {
			v = uint8(c)
			ok = true
		}
	}
	return
}

// ToUint will try to convert object o to uint value.
func ToUint(o Object) (v uint, ok bool) {
	switch o := o.(type) {
	case *Int:
		v = uint(o.Value)
		ok = true
	case *Float:
		v = uint(o.Value)
		ok = true
	case *BigInt:
		v = uint(o.Value.Int64())
		ok = true
	case *BigFloat:
		b, _ := o.Value.Int64()
		v = uint(b)
		ok = true
	case *Char:
		v = uint(o.Value)
		ok = true
	case *Bool:
		if o == TrueValue {
			v = 1
		}
		ok = true
	case *String:
		c, err := strconv.ParseInt(o.Value, 10, 64)
		if err == nil {
			v = uint(c)
			ok = true
		}
	}
	return
}

// ToFloat64 will try to convert object o to float64 value.
func ToFloat64(o Object) (v float64, ok bool) {
	switch o := o.(type) {
	case *Int:
		v = float64(o.Value)
		ok = true
	case *Float:
		v = o.Value
		ok = true
	case *BigInt:
		v, _ = o.Value.Float64()
		ok = true
	case *BigFloat:
		v, _ = o.Value.Float64()
		ok = true
	case *String:
		c, err := strconv.ParseFloat(o.Value, 64)
		if err == nil {
			v = c
			ok = true
		}
	}
	return
}

// ToBool will try to convert object o to bool value.
func ToBool(o Object) (v bool, ok bool) {
	ok = true
	v = !o.IsFalsy()
	return
}

// ToRune will try to convert object o to rune value.
func ToRune(o Object) (v rune, ok bool) {
	switch o := o.(type) {
	case *Int:
		v = rune(o.Value)
		ok = true
	case *Char:
		v = o.Value
		ok = true
	}
	return
}

// ToByteSlice will try to convert object o to []byte value.
func ToByteSlice(o Object) (v []byte, ok bool) {
	switch o := o.(type) {
	case *Bytes:
		v = o.Value
		ok = true
	case *String:
		v = []byte(o.Value)
		ok = true
	case *Int:
		v = []byte(string(rune(o.Value)))
		ok = true
	case *Char:
		v = []byte(string(o.Value))
		ok = true
	case *Array:
		var bytes []byte
		for _, obj := range o.Value {
			b, ok := ToByteSlice(obj)
			if ok {
				bytes = append(bytes, b...)
			} else {
				bytes = append(bytes, 0)
			}
		}
		v = bytes
		ok = true
	}
	return
}

// ToTime will try to convert object o to time.Time value.
func ToTime(o Object) (v time.Time, ok bool) {
	switch o := o.(type) {
	case *Time:
		v = o.Value
		ok = true
	case *Int:
		v = time.Unix(o.Value, 0)
		ok = true
	}
	return
}

// ToInterface attempts to convert an object o to an interface{} value
func ToInterface(o Object) (res interface{}) {
	switch o := o.(type) {
	case *Int:
		res = o.Value
	case *String:
		res = o.Value
	case *Float:
		res = o.Value
	case *Bool:
		res = o == TrueValue
	case *Char:
		res = o.Value
	case *Bytes:
		res = o.Value
	case *Array:
		res = make([]interface{}, len(o.Value))
		for i, val := range o.Value {
			res.([]interface{})[i] = ToInterface(val)
		}
	case *ImmutableArray:
		res = make([]interface{}, len(o.Value))
		for i, val := range o.Value {
			res.([]interface{})[i] = ToInterface(val)
		}
	case *Map:
		res = make(map[string]interface{})
		for key, v := range o.Value {
			res.(map[string]interface{})[key] = ToInterface(v)
		}
	case *ImmutableMap:
		res = make(map[string]interface{})
		for key, v := range o.Value {
			res.(map[string]interface{})[key] = ToInterface(v)
		}
	case *Time:
		res = o.Value
	case *Error:
		res = errors.New(o.String())
	case *Null:
		res = nil
	case *Struct:
		resMap := make(map[string]interface{})
		for k, v := range o.Fields {
			resMap[k] = ToInterface(v)
		}
		res = resMap
	case Object:
		return o
	}
	return
}

// FromInterface will attempt to convert an interface{} v to a tender Object
func FromInterface(v interface{}) (Object, error) {
	switch v := v.(type) {
	case nil:
		return NullValue, nil
	case string:
		if len(v) > MaxStringLen {
			return nil, ErrStringLimit
		}
		return &String{Value: v}, nil
	case int64:
		return &Int{Value: v}, nil
	case int:
		return &Int{Value: int64(v)}, nil
	case bool:
		if v {
			return TrueValue, nil
		}
		return FalseValue, nil
	case rune:
		return &Char{Value: v}, nil
	case byte:
		return &Char{Value: rune(v)}, nil
	case float64:
		return &Float{Value: v}, nil
	case []byte:
		if len(v) > MaxBytesLen {
			return nil, ErrBytesLimit
		}
		return &Bytes{Value: v}, nil
	case error:
		return &Error{Value: &String{Value: v.Error()}}, nil
	case map[string]Object:
		return &Map{Value: v}, nil
	case map[string]interface{}:
		kv := make(map[string]Object)
		for vk, vv := range v {
			vo, err := FromInterface(vv)
			if err != nil {
				return nil, err
			}
			kv[vk] = vo
		}
		return &Map{Value: kv}, nil
	case []Object:
		return &Array{Value: v}, nil
	case []interface{}:
		arr := make([]Object, len(v))
		for i, e := range v {
			vo, err := FromInterface(e)
			if err != nil {
				return nil, err
			}
			arr[i] = vo
		}
		return &Array{Value: arr}, nil
	case time.Time:
		return &Time{Value: v}, nil
	case Object:
		return v, nil
	case CallableFunc:
		return &UserFunction{Value: v}, nil
	}
	return nil, fmt.Errorf("cannot convert to object: %T", v)
}


func FromBool(b bool) Object {
	if b {
		return TrueValue
	}
	return FalseValue
}