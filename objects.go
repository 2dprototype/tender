package tender

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"math/big"
	"math/cmplx"

	"github.com/2dprototype/tender/parser"
	"github.com/2dprototype/tender/token"
)

var (
	// TrueValue represents a true value.
	TrueValue Object = &Bool{value: true}

	// FalseValue represents a false value.
	FalseValue Object = &Bool{value: false}

	// NullValue represents an null value.
	NullValue Object = &Null{}
)

// Object represents an object in the VM.
type Object interface {
	// TypeName should return the name of the type.
	TypeName() string

	// String should return a string representation of the type's value.
	String() string

	// BinaryOp should return another object that is the result of a given
	// binary operator and a right-hand side object. If BinaryOp returns an
	// error, the VM will treat it as a run-time error.
	BinaryOp(op token.Token, rhs Object) (Object, error)

	// IsFalsy should return true if the value of the type should be considered
	// as falsy.
	IsFalsy() bool

	// Equals should return true if the value of the type should be considered
	// as equal to the value of another object.
	Equals(another Object) bool

	// Copy should return a copy of the type (and its value). Copy function
	// will be used for copy() builtin function which is expected to deep-copy
	// the values generally.
	Copy() Object

	// IndexGet should take an index Object and return a result Object or an
	// error for indexable objects. Indexable is an object that can take an
	// index and return an object. If error is returned, the runtime will treat
	// it as a run-time error and ignore returned value. If Object is not
	// indexable, ErrNotIndexable should be returned as error. If nil is
	// returned as value, it will be converted to NullToken value by the
	// runtime.
	IndexGet(index Object) (value Object, err error)

	// IndexSet should take an index Object and a value Object for index
	// assignable objects. Index assignable is an object that can take an index
	// and a value on the left-hand side of the assignment statement. If Object
	// is not index assignable, ErrNotIndexAssignable should be returned as
	// error. If an error is returned, it will be treated as a run-time error.
	IndexSet(index, value Object) error

	// Iterate should return an Iterator for the type.
	Iterate() Iterator

	// CanIterate should return whether the Object can be Iterated.
	CanIterate() bool

	// Call should take an arbitrary number of arguments and returns a return
	// value and/or an error, which the VM will consider as a run-time error.
	Call(args ...Object) (ret Object, err error)

	// CanCall should return whether the Object can be Called.
	CanCall() bool
}

// ObjectImpl represents a default Object Implementation. To defined a new
// value type, one can embed ObjectImpl in their type declarations to avoid
// implementing all non-significant methods. TypeName() and String() methods
// still need to be implemented.
type ObjectImpl struct {
}

// TypeName returns the name of the type.
func (o *ObjectImpl) TypeName() string {
	panic(ErrNotImplemented)
}

func (o *ObjectImpl) String() string {
	panic(ErrNotImplemented)
}

// BinaryOp returns another object that is the result of a given binary
// operator and a right-hand side object.
func (o *ObjectImpl) BinaryOp(_ token.Token, _ Object) (Object, error) {
	return nil, ErrInvalidOperator
}

// Copy returns a copy of the type.
func (o *ObjectImpl) Copy() Object {
	return nil
}

// IsFalsy returns true if the value of the type is falsy.
func (o *ObjectImpl) IsFalsy() bool {
	return false
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *ObjectImpl) Equals(x Object) bool {
	return o == x
}

// IndexGet returns an element at a given index.
func (o *ObjectImpl) IndexGet(_ Object) (res Object, err error) {
	return nil, ErrNotIndexable
}

// IndexSet sets an element at a given index.
func (o *ObjectImpl) IndexSet(_, _ Object) (err error) {
	return ErrNotIndexAssignable
}

// Iterate returns an iterator.
func (o *ObjectImpl) Iterate() Iterator {
	return nil
}

// CanIterate returns whether the Object can be Iterated.
func (o *ObjectImpl) CanIterate() bool {
	return false
}

// Call takes an arbitrary number of arguments and returns a return value
// and/or an error.
func (o *ObjectImpl) Call(_ ...Object) (ret Object, err error) {
	return nil, nil
}

// CanCall returns whether the Object can be Called.
func (o *ObjectImpl) CanCall() bool {
	return false
}

// Array represents an array of objects.
type Array struct {
	ObjectImpl
	Value []Object
}

// TypeName returns the name of the type.
func (o *Array) TypeName() string {
	return "array"
}

func (o *Array) String() string {
	var elements []string
	for _, e := range o.Value {
		elements = append(elements, e.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(elements, ", "))
}

// BinaryOp returns another object that is the result of a given binary
// operator and a right-hand side object.
func (o *Array) BinaryOp(op token.Token, rhs Object) (Object, error) {
	if rhs, ok := rhs.(*Array); ok {
		switch op {
		case token.Add:
			if len(rhs.Value) == 0 {
				return o, nil
			}
			return &Array{Value: append(o.Value, rhs.Value...)}, nil
		}
	}
	return nil, ErrInvalidOperator
}

// Copy returns a copy of the type.
func (o *Array) Copy() Object {
	var c []Object
	for _, elem := range o.Value {
		c = append(c, elem.Copy())
	}
	return &Array{Value: c}
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Array) IsFalsy() bool {
	return len(o.Value) == 0
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *Array) Equals(x Object) bool {
	var xVal []Object
	switch x := x.(type) {
	case *Array:
		xVal = x.Value
	case *ImmutableArray:
		xVal = x.Value
	default:
		return false
	}
	if len(o.Value) != len(xVal) {
		return false
	}
	for i, e := range o.Value {
		if !e.Equals(xVal[i]) {
			return false
		}
	}
	return true
}

// IndexGet returns an element at a given index.
// func (o *Array) IndexGet(index Object) (res Object, err error) {
	// strIdx, ok := index.(*String) 
	// if ok {
		// if strIdx.Value == "length" {
			// return &Int{Value: int64(len(o.Value))}, nil
		// } else if strIdx.Value == "push" {
			// return &BuiltinFunction{
				// Value: func(args ...Object) (Object, error) {
					// o.Value = append(o.Value, args...)
					// return o, nil
				// },
			// }, nil
		// } 
		// return nil, nil
	// }
	// intIdx, ok := index.(*Int)
	// if !ok {
		// err = ErrInvalidIndexType
		// return
	// }
	// idxVal := int(intIdx.Value)
	// if idxVal < 0 || idxVal >= len(o.Value) {
		// res = NullValue
		// return
	// }
	// res = o.Value[idxVal]
	// return
// }

func (o *Array) IndexGet(index Object) (res Object, err error) {
	strIdx, ok := index.(*String)
	if ok {
		switch strIdx.Value {
		case "length":
			return &Int{Value: int64(len(o.Value))}, nil
		case "push":
			return &BuiltinFunction{
				Value: func(args ...Object) (Object, error) {
					o.Value = append(o.Value, args...)
					return &Int{Value: int64(len(o.Value))}, nil
				},
			}, nil
		case "pop":
			return &BuiltinFunction{
				Value: func(args ...Object) (Object, error) {
					if len(o.Value) == 0 {
						return NullValue, nil
					}
					last := o.Value[len(o.Value)-1]
					o.Value = o.Value[:len(o.Value)-1]
					return last, nil
				},
			}, nil
		case "shift":
			return &BuiltinFunction{
				Value: func(args ...Object) (Object, error) {
					if len(o.Value) == 0 {
						return NullValue, nil
					}
					first := o.Value[0]
					o.Value = o.Value[1:]
					return first, nil
				},
			}, nil
		case "unshift":
			return &BuiltinFunction{
				Value: func(args ...Object) (Object, error) {
					o.Value = append(args, o.Value...)
					return &Int{Value: int64(len(o.Value))}, nil
				},
			}, nil
		// case "slice":
			// return &BuiltinFunction{
				// Value: func(args ...Object) (Object, error) {
					// if len(args) == 0 {
						// return o.Copy(), nil
					// }
					// startObj, ok := args[0].(*Int)
					// if !ok {
						// return nil, ErrInvalidIndexType
					// }
					// start := int(startObj.Value)
					// end := len(o.Value)
					// if len(args) > 1 {
						// endObj, ok := args[1].(*Int)
						// if !ok {
							// return nil, ErrInvalidIndexType
						// }
						// end = int(endObj.Value)
					// }
					// if start < 0 {
						// start = 0
					// }
					// if end > len(o.Value) {
						// end = len(o.Value)
					// }
					// return &Array{Value: append([]Object{}, o.Value[start:end]...)}, nil
				// },
			// }, nil
		default:
			return nil, nil
		}
	}

	intIdx, ok := index.(*Int)
	if !ok {
		err = ErrInvalidIndexType
		return
	}

	idxVal := int(intIdx.Value)
	if idxVal < 0 || idxVal >= len(o.Value) {
		res = NullValue
		return
	}
	res = o.Value[idxVal]
	return
}


// IndexSet sets an element at a given index.
func (o *Array) IndexSet(index, value Object) (err error) {
	intIdx, ok := ToInt(index)
	if !ok {
		err = ErrInvalidIndexType
		return
	}
	if intIdx < 0 || intIdx >= len(o.Value) {
		err = ErrIndexOutOfBounds
		return
	}
	o.Value[intIdx] = value
	return nil
}

// Iterate creates an array iterator.
func (o *Array) Iterate() Iterator {
	return &ArrayIterator{
		v: o.Value,
		l: len(o.Value),
	}
}

// CanIterate returns whether the Object can be Iterated.
func (o *Array) CanIterate() bool {
	return true
}

// Bool represents a boolean value.
type Bool struct {
	ObjectImpl

	// this is intentionally non-public to force using objects.TrueValue and
	// FalseValue always
	value bool
}

func (o *Bool) String() string {
	if o.value {
		return "true"
	}

	return "false"
}

// TypeName returns the name of the type.
func (o *Bool) TypeName() string {
	return "bool"
}

// Copy returns a copy of the type.
func (o *Bool) Copy() Object {
	return o
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Bool) IsFalsy() bool {
	return !o.value
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *Bool) Equals(x Object) bool {
	return o == x
}

// GobDecode decodes bool value from input bytes.
func (o *Bool) GobDecode(b []byte) (err error) {
	o.value = b[0] == 1
	return
}

// GobEncode encodes bool values into bytes.
func (o *Bool) GobEncode() (b []byte, err error) {
	if o.value {
		b = []byte{1}
	} else {
		b = []byte{0}
	}
	return
}

// BuiltinFunction represents a builtin function.
type BuiltinFunction struct {
	ObjectImpl
	Name      string
	Value     CallableFunc
	NeedVMObj bool
}

// TypeName returns the name of the type.
func (o *BuiltinFunction) TypeName() string {
	return "builtin-function:" + o.Name
}

func (o *BuiltinFunction) String() string {
	return "<builtin-function>"
}

// Copy returns a copy of the type.
func (o *BuiltinFunction) Copy() Object {
	return &BuiltinFunction{Value: o.Value, NeedVMObj: o.NeedVMObj}
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *BuiltinFunction) Equals(_ Object) bool {
	return false
}

// Call executes a builtin function.
func (o *BuiltinFunction) Call(args ...Object) (Object, error) {
	return o.Value(args...)
}

// CanCall returns whether the Object can be Called.
func (o *BuiltinFunction) CanCall() bool {
	return true
}

// BuiltinModule is an importable module that's written in Go.
type BuiltinModule struct {
	Attrs map[string]Object
}

// Import returns an immutable map for the module.
func (m *BuiltinModule) Import(moduleName string) (interface{}, error) {
	return m.AsImmutableMap(moduleName), nil
}

// AsImmutableMap converts builtin module into an immutable map.
func (m *BuiltinModule) AsImmutableMap(moduleName string) *ImmutableMap {
	attrs := make(map[string]Object, len(m.Attrs))
	for k, v := range m.Attrs {
		attrs[k] = v.Copy()
	}
	attrs["__module_name__"] = &String{Value: moduleName}
	return &ImmutableMap{Value: attrs}
}

// Bytes represents a byte array.
type Bytes struct {
	ObjectImpl
	Value []byte
}

func (o *Bytes) String() string {
	return fmt.Sprintf("%v", o.Value)
}

// TypeName returns the name of the type.
func (o *Bytes) TypeName() string {
	return "bytes"
}

// BinaryOp returns another object that is the result of a given binary
// operator and a right-hand side object.
func (o *Bytes) BinaryOp(op token.Token, rhs Object) (Object, error) {
	switch op {
	case token.Add:
		switch rhs := rhs.(type) {
		case *Bytes:
			if len(o.Value)+len(rhs.Value) > MaxBytesLen {
				return nil, ErrBytesLimit
			}
			return &Bytes{Value: append(o.Value, rhs.Value...)}, nil
		}
	}
	return nil, ErrInvalidOperator
}

// Copy returns a copy of the type.
func (o *Bytes) Copy() Object {
	return &Bytes{Value: append([]byte{}, o.Value...)}
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Bytes) IsFalsy() bool {
	return len(o.Value) == 0
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *Bytes) Equals(x Object) bool {
	t, ok := x.(*Bytes)
	if !ok {
		return false
	}
	return bytes.Equal(o.Value, t.Value)
}


func (o *Bytes) IndexSet(index, value Object) (err error) {
    intIdx, ok := ToInt(index)
    if !ok {
        return ErrInvalidIndexType
    }
    if intIdx < 0 || intIdx >= len(o.Value) {
        return ErrIndexOutOfBounds
    }
    
    // Ensure the value is an integer
    byteValue, ok := ToByte(value)
    if !ok {
        return ErrInvalidValueType
    }
    
    // Ensure the integer value is within byte range
    if byteValue < 0 || byteValue > 255 {
        return ErrByteValueOutOfRange
    }
    
    // Set the value at the specified index
    o.Value[intIdx] = byteValue
    return nil
}


// IndexGet returns an element (as Int) at a given index.
func (o *Bytes) IndexGet(index Object) (res Object, err error) {
	strIdx, ok := index.(*String) 
	if ok {
		if strIdx.Value == "length" {
			return &Int{Value: int64(len(o.Value))}, nil
		}
		return nil, nil
	}
	intIdx, ok := index.(*Int)
	if !ok {
		err = ErrInvalidIndexType
		return
	}
	idxVal := int(intIdx.Value)
	if idxVal < 0 || idxVal >= len(o.Value) {
		res = NullValue
		return
	}
	res = &Int{Value: int64(o.Value[idxVal])}
	return
}

// Iterate creates a bytes iterator.
func (o *Bytes) Iterate() Iterator {
	return &BytesIterator{
		v: o.Value,
		l: len(o.Value),
	}
}

// CanIterate returns whether the Object can be Iterated.
func (o *Bytes) CanIterate() bool {
	return true
}

// Char represents a character value.
type Char struct {
	ObjectImpl
	Value rune
}

func (o *Char) String() string {
	return string(o.Value)
}

// TypeName returns the name of the type.
func (o *Char) TypeName() string {
	return "char"
}

// BinaryOp returns another object that is the result of a given binary
// operator and a right-hand side object.
func (o *Char) BinaryOp(op token.Token, rhs Object) (Object, error) {
	switch rhs := rhs.(type) {
	case *Char:
		switch op {
		case token.Add:
			r := o.Value + rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Char{Value: r}, nil
		case token.Sub:
			r := o.Value - rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Char{Value: r}, nil
		case token.Less:
			if o.Value < rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.Greater:
			if o.Value > rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.LessEq:
			if o.Value <= rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.GreaterEq:
			if o.Value >= rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		}
	case *Int:
		switch op {
		case token.Add:
			r := o.Value + rune(rhs.Value)
			if r == o.Value {
				return o, nil
			}
			return &Char{Value: r}, nil
		case token.Sub:
			r := o.Value - rune(rhs.Value)
			if r == o.Value {
				return o, nil
			}
			return &Char{Value: r}, nil
		case token.Less:
			if int64(o.Value) < rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.Greater:
			if int64(o.Value) > rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.LessEq:
			if int64(o.Value) <= rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.GreaterEq:
			if int64(o.Value) >= rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		}
	}
	return nil, ErrInvalidOperator
}

// Copy returns a copy of the type.
func (o *Char) Copy() Object {
	return &Char{Value: o.Value}
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Char) IsFalsy() bool {
	return o.Value == 0
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *Char) Equals(x Object) bool {
	t, ok := x.(*Char)
	if !ok {
		return false
	}
	return o.Value == t.Value
}

// CompiledFunction represents a compiled function.
type CompiledFunction struct {
	ObjectImpl
	Instructions  []byte
	NumLocals     int // number of local variables (including function parameters)
	NumParameters int
	VarArgs       bool
	SourceMap     map[int]parser.Pos
	Free          []*ObjectPtr
}

// TypeName returns the name of the type.
func (o *CompiledFunction) TypeName() string {
	return "compiled-function"
}

func (o *CompiledFunction) String() string {
	return "<compiled-function>"
}

// Copy returns a copy of the type.
func (o *CompiledFunction) Copy() Object {
	return &CompiledFunction{
		Instructions:  append([]byte{}, o.Instructions...),
		NumLocals:     o.NumLocals,
		NumParameters: o.NumParameters,
		VarArgs:       o.VarArgs,
		SourceMap:     o.SourceMap,
		Free:          append([]*ObjectPtr{}, o.Free...), // DO NOT Copy() of elements; these are variable pointers
	}
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *CompiledFunction) Equals(_ Object) bool {
	return false
}

// SourcePos returns the source position of the instruction at ip.
func (o *CompiledFunction) SourcePos(ip int) parser.Pos {
	for ip >= 0 {
		if p, ok := o.SourceMap[ip]; ok {
			return p
		}
		ip--
	}
	return parser.NoPos
}

// CanCall returns whether the Object can be Called.
func (o *CompiledFunction) CanCall() bool {
	return true
}

// Error represents an error value.
type Error struct {
	ObjectImpl
	Value Object
}

// TypeName returns the name of the type.
func (o *Error) TypeName() string {
	return "error"
}

func (o *Error) String() string {
	if o.Value != nil {
		return fmt.Sprintf("error: %s", o.Value.String())
	}
	return "error"
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Error) IsFalsy() bool {
	return true // error is always false.
}

// Copy returns a copy of the type.
func (o *Error) Copy() Object {
	return &Error{Value: o.Value.Copy()}
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *Error) Equals(x Object) bool {
	return o == x // pointer equality
}

// IndexGet returns an element at a given index.
func (o *Error) IndexGet(index Object) (res Object, err error) {
	if strIdx, _ := ToString(index); strIdx != "value" {
		err = ErrInvalidIndexOnError
		return
	}
	res = o.Value
	return
}

// ImmutableArray represents an immutable array of objects.
type ImmutableArray struct {
	ObjectImpl
	Value []Object
}

// TypeName returns the name of the type.
func (o *ImmutableArray) TypeName() string {
	return "immutable-array"
}

func (o *ImmutableArray) String() string {
	var elements []string
	for _, e := range o.Value {
		elements = append(elements, e.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(elements, ", "))
}

// BinaryOp returns another object that is the result of a given binary
// operator and a right-hand side object.
func (o *ImmutableArray) BinaryOp(op token.Token, rhs Object) (Object, error) {
	if rhs, ok := rhs.(*ImmutableArray); ok {
		switch op {
		case token.Add:
			return &Array{Value: append(o.Value, rhs.Value...)}, nil
		}
	}
	return nil, ErrInvalidOperator
}

// Copy returns a copy of the type.
func (o *ImmutableArray) Copy() Object {
	var c []Object
	for _, elem := range o.Value {
		c = append(c, elem.Copy())
	}
	return &Array{Value: c}
}

// IsFalsy returns true if the value of the type is falsy.
func (o *ImmutableArray) IsFalsy() bool {
	return len(o.Value) == 0
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *ImmutableArray) Equals(x Object) bool {
	var xVal []Object
	switch x := x.(type) {
	case *Array:
		xVal = x.Value
	case *ImmutableArray:
		xVal = x.Value
	default:
		return false
	}
	if len(o.Value) != len(xVal) {
		return false
	}
	for i, e := range o.Value {
		if !e.Equals(xVal[i]) {
			return false
		}
	}
	return true
}

// IndexGet returns an element at a given index.
func (o *ImmutableArray) IndexGet(index Object) (res Object, err error) {
	strIdx, ok := index.(*String) 
	if ok {
		if strIdx.Value == "length" {
			return &Int{Value: int64(len(o.Value))}, nil
		}
		return nil, nil
	}
	intIdx, ok := index.(*Int)
	if !ok {
		err = ErrInvalidIndexType
		return
	}
	idxVal := int(intIdx.Value)
	if idxVal < 0 || idxVal >= len(o.Value) {
		res = NullValue
		return
	}
	res = o.Value[idxVal]
	return
}

// Iterate creates an array iterator.
func (o *ImmutableArray) Iterate() Iterator {
	return &ArrayIterator{
		v: o.Value,
		l: len(o.Value),
	}
}

// CanIterate returns whether the Object can be Iterated.
func (o *ImmutableArray) CanIterate() bool {
	return true
}

// ImmutableMap represents an immutable map object.
type ImmutableMap struct {
	ObjectImpl
	Value map[string]Object
}

// TypeName returns the name of the type.
func (o *ImmutableMap) TypeName() string {
	return "immutable-map"
}

func (o *ImmutableMap) String() string {
	var pairs []string
	for k, v := range o.Value {
		pairs = append(pairs, fmt.Sprintf("%s: %s", k, v.String()))
	}
	return fmt.Sprintf("{%s}", strings.Join(pairs, ", "))
}

// Copy returns a copy of the type.
func (o *ImmutableMap) Copy() Object {
	c := make(map[string]Object)
	for k, v := range o.Value {
		c[k] = v.Copy()
	}
	return &Map{Value: c}
}

// IsFalsy returns true if the value of the type is falsy.
func (o *ImmutableMap) IsFalsy() bool {
	return len(o.Value) == 0
}

// IndexGet returns the value for the given key.
func (o *ImmutableMap) IndexGet(index Object) (res Object, err error) {
	strIdx, ok := ToString(index)
	if !ok {
		err = ErrInvalidIndexType
		return
	}
	res, ok = o.Value[strIdx]
	if !ok {
		res = NullValue
	}
	return
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *ImmutableMap) Equals(x Object) bool {
	var xVal map[string]Object
	switch x := x.(type) {
	case *Map:
		xVal = x.Value
	case *ImmutableMap:
		xVal = x.Value
	default:
		return false
	}
	if len(o.Value) != len(xVal) {
		return false
	}
	for k, v := range o.Value {
		tv := xVal[k]
		if !v.Equals(tv) {
			return false
		}
	}
	return true
}

// Iterate creates an immutable map iterator.
func (o *ImmutableMap) Iterate() Iterator {
	var keys []string
	for k := range o.Value {
		keys = append(keys, k)
	}
	return &MapIterator{
		v: o.Value,
		k: keys,
		l: len(keys),
	}
}

// CanIterate returns whether the Object can be Iterated.
func (o *ImmutableMap) CanIterate() bool {
	return true
}


// Float represents a floating point number value.
type Float struct {
	ObjectImpl
	Value float64
}

func (o *Float) String() string {
	return strconv.FormatFloat(o.Value, 'f', -1, 64)
}

// TypeName returns the name of the type.
func (o *Float) TypeName() string {
	return "float"
}

// BinaryOp returns another object that is the result of a given binary
// operator and a right-hand side object.
func (o *Float) BinaryOp(op token.Token, rhs Object) (Object, error) {
	switch rhs := rhs.(type) {
	case *Float:
		switch op {
		case token.Add:
			r := o.Value + rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
		case token.Sub:
			r := o.Value - rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
		case token.Mul:
			r := o.Value * rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
		case token.Quo:
			if rhs.Value == 0 {
				return &Float{Value: math.Inf(1)}, nil
			}
			r := o.Value / rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
		case token.Less:
			if o.Value < rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.Greater:
			if o.Value > rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.LessEq:
			if o.Value <= rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.GreaterEq:
			if o.Value >= rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		}
	case *Int:
		switch op {
		case token.Add:
			r := o.Value + float64(rhs.Value)
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
		case token.Sub:
			r := o.Value - float64(rhs.Value)
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
		case token.Mul:
			r := o.Value * float64(rhs.Value)
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
		case token.Quo:
			if rhs.Value == 0 {
				return &Float{Value: math.Inf(1)}, nil
			}
			r := o.Value / float64(rhs.Value)
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
		case token.Less:
			if o.Value < float64(rhs.Value) {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.Greater:
			if o.Value > float64(rhs.Value) {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.LessEq:
			if o.Value <= float64(rhs.Value) {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.GreaterEq:
			if o.Value >= float64(rhs.Value) {
				return TrueValue, nil
			}
			return FalseValue, nil
		}
		case *BigInt:
			bi := new(big.Int)
			new(big.Float).SetFloat64(o.Value).Int(bi)
			return binaryOpBigInt(op, bi, rhs.Value), nil
		case *BigFloat:
			return binaryOpBigFloat(op, new(big.Float).SetFloat64(o.Value), rhs.Value), nil
		case *Complex:
			return binaryOpComplex(op, complex(o.Value, 0), rhs.Value), nil
	}
	return nil, ErrInvalidOperator
}

// Copy returns a copy of the type.
func (o *Float) Copy() Object {
	return &Float{Value: o.Value}
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Float) IsFalsy() bool {
	return math.IsNaN(o.Value)
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *Float) Equals(x Object) bool {
	t, ok := x.(*Float)
	if !ok {
		return false
	}
	return o.Value == t.Value
}

// Int represents an integer value.
type Int struct {
	ObjectImpl
	Value int64
}

func (o *Int) String() string {
	return strconv.FormatInt(o.Value, 10)
}

// TypeName returns the name of the type.
func (o *Int) TypeName() string {
	return "int"
}

// BinaryOp returns another object that is the result of a given binary
// operator and a right-hand side object.
func (o *Int) BinaryOp(op token.Token, rhs Object) (Object, error) {
	switch rhs := rhs.(type) {
	case *Int:
		switch op {
		case token.Add:
			r := o.Value + rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil
		case token.Sub:
			r := o.Value - rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil
		case token.Mul:
			r := o.Value * rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil
		case token.Quo:
			if rhs.Value == 0 {
				return &Float{Value: math.Inf(1)}, nil
			}
			r := o.Value / rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil
		case token.Rem:
			r := o.Value % rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil
		case token.And:
			r := o.Value & rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil
		case token.Or:
			r := o.Value | rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil
		case token.Xor:
			r := o.Value ^ rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil
		case token.AndNot:
			r := o.Value &^ rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil
		case token.Shl:
			r := o.Value << uint64(rhs.Value)
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil
		case token.Shr:
			r := o.Value >> uint64(rhs.Value)
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil
		case token.Less:
			if o.Value < rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.Greater:
			if o.Value > rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.LessEq:
			if o.Value <= rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.GreaterEq:
			if o.Value >= rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		}
	case *Float:
		switch op {
		case token.Add:
			return &Float{Value: float64(o.Value) + rhs.Value}, nil
		case token.Sub:
			return &Float{Value: float64(o.Value) - rhs.Value}, nil
		case token.Mul:
			return &Float{Value: float64(o.Value) * rhs.Value}, nil
		case token.Quo:
			if rhs.Value == 0 {
				return &Float{Value: math.Inf(1)}, nil
			}
			return &Float{Value: float64(o.Value) / rhs.Value}, nil
		case token.Less:
			if float64(o.Value) < rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.Greater:
			if float64(o.Value) > rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.LessEq:
			if float64(o.Value) <= rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.GreaterEq:
			if float64(o.Value) >= rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		}
	case *Char:
		switch op {
		case token.Add:
			return &Char{Value: rune(o.Value) + rhs.Value}, nil
		case token.Sub:
			return &Char{Value: rune(o.Value) - rhs.Value}, nil
		case token.Less:
			if o.Value < int64(rhs.Value) {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.Greater:
			if o.Value > int64(rhs.Value) {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.LessEq:
			if o.Value <= int64(rhs.Value) {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.GreaterEq:
			if o.Value >= int64(rhs.Value) {
				return TrueValue, nil
			}
			return FalseValue, nil
		}
		case *BigInt:
			return binaryOpBigInt(op, new(big.Int).SetInt64(o.Value), rhs.Value), nil
		case *BigFloat:
			return binaryOpBigFloat(op, new(big.Float).SetInt64(o.Value), rhs.Value), nil
		case *Complex:
			return binaryOpComplex(op, complex(float64(o.Value), 0), rhs.Value), nil
	}
	return nil, ErrInvalidOperator
}

// Copy returns a copy of the type.
func (o *Int) Copy() Object {
	return &Int{Value: o.Value}
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Int) IsFalsy() bool {
	return o.Value == 0
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *Int) Equals(x Object) bool {
	t, ok := x.(*Int)
	if !ok {
		return false
	}
	return o.Value == t.Value
}

// BigInt represents an arbitrary-precision integer value.
type BigInt struct {
    ObjectImpl
    Value *big.Int
}

func (b *BigInt) String() string {
    return b.Value.String()
}

// TypeName returns the name of the type.
func (b *BigInt) TypeName() string {
    return "bigint"
}


func binaryOpBigInt (op token.Token, lhs *big.Int, rhs *big.Int) Object {
	switch op {
        case token.Add:
            r := new(big.Int).Set(lhs)
            r.Add(r, rhs)
            return &BigInt{Value: r}
        case token.Sub:
            r := new(big.Int).Set(lhs)
            r.Sub(r, rhs)
            return &BigInt{Value: r}
        case token.Mul:
            r := new(big.Int).Set(lhs)
            r.Mul(r, rhs)
            return &BigInt{Value: r}
        case token.Quo:
			if rhs.Int64() == 0 {
				return &Float{Value: math.Inf(1)}
			}
            r := new(big.Int).Set(lhs)
            r.Div(r, rhs)
            return &BigInt{Value: r}
        case token.Rem:
            r := new(big.Int).Set(lhs)
            r.Mod(r, rhs)
            return &BigInt{Value: r}
        case token.And:
            r := new(big.Int).Set(lhs)
            r.And(r, rhs)
            return &BigInt{Value: r}
        case token.Or:
            r := new(big.Int).Set(lhs)
            r.Or(r, rhs)
            return &BigInt{Value: r}
        case token.Xor:
            r := new(big.Int).Set(lhs)
            r.Xor(r, rhs)
            return &BigInt{Value: r}
        case token.AndNot:
            r := new(big.Int).Set(lhs)
            r.AndNot(r, rhs)
            return &BigInt{Value: r}
        case token.Shl:
            r := new(big.Int).Set(lhs)
            r.Lsh(r, uint(rhs.Int64()))
            return &BigInt{Value: r}
        case token.Shr:
            r := new(big.Int).Set(lhs)
            r.Rsh(r, uint(rhs.Int64()))
            return &BigInt{Value: r}
        case token.Less:
            if lhs.Cmp(rhs) < 0 {
                return TrueValue
            }
            return FalseValue
        case token.Greater:
            if lhs.Cmp(rhs) > 0 {
                return TrueValue
            }
            return FalseValue
        case token.LessEq:
            if lhs.Cmp(rhs) <= 0 {
                return TrueValue
            }
            return FalseValue
        case token.GreaterEq:
            if lhs.Cmp(rhs) >= 0 {
                return TrueValue
            }
            return FalseValue
	}
	
	return nil
}

// BinaryOp performs binary operations with another Object.
func (b *BigInt) BinaryOp(op token.Token, rhs Object) (Object, error) {
    switch rhs := rhs.(type) {
		case *BigInt:
			return binaryOpBigInt(op, b.Value, rhs.Value), nil
		case *Float:
			return binaryOpBigInt(op, b.Value, new(big.Int).SetInt64(int64(rhs.Value))), nil
		case *Int:
			return binaryOpBigInt(op, b.Value, new(big.Int).SetInt64(rhs.Value)), nil
		case *BigFloat:
			return binaryOpBigFloat(op, new(big.Float).SetInt(b.Value), rhs.Value), nil
    }
    return nil, ErrInvalidOperator
}

// Copy returns a copy of the BigInt.
func (b *BigInt) Copy() Object {
    return &BigInt{Value: new(big.Int).Set(b.Value)}
}

// IsFalsy returns true if the value of the BigInt is falsy (i.e., zero).
func (b *BigInt) IsFalsy() bool {
    return b.Value.Sign() == 0
}

// Equals checks if the BigInt is equal to another Object.
func (b *BigInt) Equals(x Object) bool {
    t, ok := x.(*BigInt)
    if !ok {
        return false
    }
    return b.Value.Cmp(t.Value) == 0
}


// BigFloat represents an arbitrary-precision floating-point value.
type BigFloat struct {
    ObjectImpl
    Value *big.Float
}

// String returns the string representation of the BigFloat.
func (b *BigFloat) String() string {
    return b.Value.String()
}

// TypeName returns the name of the type.
func (b *BigFloat) TypeName() string {
    return "bigfloat"
}

func binaryOpBigFloat (op token.Token, lhs *big.Float, rhs *big.Float) Object {
	switch op {
        case token.Add:
			r := new(big.Float).Set(lhs)
            r.Add(r, rhs)
            return &BigFloat{Value: r}
        case token.Sub:
			r := new(big.Float).Set(lhs)
            r.Sub(r, rhs)
            return &BigFloat{Value: r}
        case token.Mul:
			r := new(big.Float).Set(lhs)
            r.Mul(r, rhs)
            return &BigFloat{Value: r}
        case token.Quo:
			rv, _ := rhs.Int64()
			if rv == 0 {
				return &Float{Value: math.Inf(1)}
			}
			r := new(big.Float).Set(lhs)
            r.Quo(r, rhs)
            return &BigFloat{Value: r}
        case token.Less:
            if lhs.Cmp(rhs) < 0 {
                return TrueValue
            }
            return FalseValue
        case token.Greater:
            if lhs.Cmp(rhs) > 0 {
                return TrueValue
            }
            return FalseValue
        case token.LessEq:
            if lhs.Cmp(rhs) <= 0 {
                return TrueValue
            }
            return FalseValue
        case token.GreaterEq:
            if lhs.Cmp(rhs) >= 0 {
                return TrueValue
            }
            return FalseValue
	}
	return nil
}

// BinaryOp performs binary operations with another Object.
func (b *BigFloat) BinaryOp(op token.Token, rhs Object) (Object, error) {
    switch rhs := rhs.(type) {
		case *BigFloat:
			return binaryOpBigFloat(op, b.Value, rhs.Value), nil
		case *Float:
			return binaryOpBigFloat(op, b.Value, new(big.Float).SetFloat64(rhs.Value)), nil
		case *Int:
			return binaryOpBigFloat(op, b.Value, new(big.Float).SetInt64(rhs.Value)), nil
		case *BigInt:
			return binaryOpBigFloat(op, b.Value, new(big.Float).SetInt(rhs.Value)), nil
    }
    return nil, ErrInvalidOperator
}

// Copy returns a copy of the BigFloat.
func (b *BigFloat) Copy() Object {
    return &BigFloat{Value: new(big.Float).Set(b.Value)}
}

// IsFalsy returns true if the value of the BigFloat is falsy (i.e., zero).
func (b *BigFloat) IsFalsy() bool {
    return b.Value.Cmp(big.NewFloat(0)) == 0
}

// Equals checks if the BigFloat is equal to another Object.
func (b *BigFloat) Equals(x Object) bool {
    t, ok := x.(*BigFloat)
    if !ok {
        return false
    }
    return b.Value.Cmp(t.Value) == 0
}

// Complex represents a complex number
type Complex struct {
	ObjectImpl
	Value complex128
}

// String returns the string representation of the Complex number.
func (c *Complex) String() string {
	if cmplx.IsNaN(c.Value) {
		return "NaN"
	} else if cmplx.IsInf(c.Value) {
		return "Inf"
	}
	if imag(c.Value) < 0 {
		return fmt.Sprintf("%g%gi", real(c.Value), imag(c.Value))
	}
	return fmt.Sprintf("%g+%gi", real(c.Value), imag(c.Value))
}


// IndexGet returns a character at a given index.
func (c *Complex) IndexGet(index Object) (res Object, err error) {
	strIdx, ok := index.(*String) 
	if ok {
		if strIdx.Value == "real" {
			return &Float{Value: real(c.Value)}, nil
		} else if strIdx.Value == "imag" {
			return &Float{Value: imag(c.Value)}, nil
		}
		return nil, nil
	}
	return nil, nil
}

// TypeName returns the name of the type.
func (c *Complex) TypeName() string {
	return "complex"
}

// BinaryOp performs binary operations with another Object.
func (c *Complex) BinaryOp(op token.Token, rhs Object) (Object, error) {
	switch rhs := rhs.(type) {
	case *Complex:
		return binaryOpComplex(op, c.Value, rhs.Value), nil
	case *Float:
		return binaryOpComplex(op, c.Value, complex(rhs.Value, 0)), nil
	case *Int:
		return binaryOpComplex(op, c.Value, complex(float64(rhs.Value), 0)), nil
	}
	return nil, ErrInvalidOperator
}

// binaryOpComplex handles binary operations for Complex numbers.
func binaryOpComplex(op token.Token, lhs, rhs complex128) Object {
	switch op {
	case token.Add:
		return &Complex{Value: lhs + rhs}
	case token.Sub:
		return &Complex{Value: lhs - rhs}
	case token.Mul:
		return &Complex{Value: lhs * rhs}
	case token.Quo:
		if rhs == 0 {
			// return &Complex{Value: cmplx.Inf()}
			return &Complex{Value: lhs / rhs}
		}
		return &Complex{Value: lhs / rhs}
	}
	return nil
} 

// Copy returns a copy of the Complex number.
func (c *Complex) Copy() Object {
	return &Complex{Value: c.Value}
}

// IsFalsy returns true if the Complex number is zero.
func (c *Complex) IsFalsy() bool {
	return c.Value == 0
}

// Equals checks if the Complex number is equal to another Object.
func (c *Complex) Equals(x Object) bool {
	t, ok := x.(*Complex)
	if !ok {
		return false
	}
	return c.Value == t.Value
}



// Map represents a map of objects.
type Map struct {
	ObjectImpl
	Value map[string]Object
}

// TypeName returns the name of the type.
func (o *Map) TypeName() string {
	return "map"
}

func (o *Map) String() string {
	var pairs []string
	for k, v := range o.Value {
		pairs = append(pairs, fmt.Sprintf("%s: %s", k, v.String()))
	}
	return fmt.Sprintf("{%s}", strings.Join(pairs, ", "))
}

// Copy returns a copy of the type.
func (o *Map) Copy() Object {
	c := make(map[string]Object)
	for k, v := range o.Value {
		c[k] = v.Copy()
	}
	return &Map{Value: c}
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Map) IsFalsy() bool {
	return len(o.Value) == 0
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *Map) Equals(x Object) bool {
	var xVal map[string]Object
	switch x := x.(type) {
	case *Map:
		xVal = x.Value
	case *ImmutableMap:
		xVal = x.Value
	default:
		return false
	}
	if len(o.Value) != len(xVal) {
		return false
	}
	for k, v := range o.Value {
		tv := xVal[k]
		if !v.Equals(tv) {
			return false
		}
	}
	return true
}

// IndexGet returns the value for the given key.
func (o *Map) IndexGet(index Object) (res Object, err error) {
	strIdx, ok := ToString(index)
	if !ok {
		err = ErrInvalidIndexType
		return
	}
	res, ok = o.Value[strIdx]
	if !ok {
		res = NullValue
	}
	return
}

// IndexSet sets the value for the given key.
func (o *Map) IndexSet(index, value Object) (err error) {
	strIdx, ok := ToString(index)
	if !ok {
		err = ErrInvalidIndexType
		return
	}
	o.Value[strIdx] = value
	return nil
}

// Iterate creates a map iterator.
func (o *Map) Iterate() Iterator {
	var keys []string
	for k := range o.Value {
		keys = append(keys, k)
	}
	return &MapIterator{
		v: o.Value,
		k: keys,
		l: len(keys),
	}
}

// CanIterate returns whether the Object can be Iterated.
func (o *Map) CanIterate() bool {
	return true
}

// ObjectPtr represents a free variable.
type ObjectPtr struct {
	ObjectImpl
	Value *Object
}

func (o *ObjectPtr) String() string {
	return "free-var"
}

// TypeName returns the name of the type.
func (o *ObjectPtr) TypeName() string {
	return "<free-var>"
}

// Copy returns a copy of the type.
func (o *ObjectPtr) Copy() Object {
	return o
}

// IsFalsy returns true if the value of the type is falsy.
func (o *ObjectPtr) IsFalsy() bool {
	return o.Value == nil
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *ObjectPtr) Equals(x Object) bool {
	return o == x
}

// String represents a string value.
type String struct {
	ObjectImpl
	Value   string
	runeStr []rune
}

// TypeName returns the name of the type.
func (o *String) TypeName() string {
	return "string"
}

func (o *String) String() string {
	return strconv.Quote(o.Value)
}

// BinaryOp returns another object that is the result of a given binary
// operator and a right-hand side object.
func (o *String) BinaryOp(op token.Token, rhs Object) (Object, error) {
	switch op {
	case token.Add:
		switch rhs := rhs.(type) {
		case *String:
			if len(o.Value)+len(rhs.Value) > MaxStringLen {
				return nil, ErrStringLimit
			}
			return &String{Value: o.Value + rhs.Value}, nil
		default:
			rhsStr := rhs.String()
			if len(o.Value)+len(rhsStr) > MaxStringLen {
				return nil, ErrStringLimit
			}
			return &String{Value: o.Value + rhsStr}, nil
		}
	case token.Less:
		switch rhs := rhs.(type) {
		case *String:
			if o.Value < rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		}
	case token.LessEq:
		switch rhs := rhs.(type) {
		case *String:
			if o.Value <= rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		}
	case token.Greater:
		switch rhs := rhs.(type) {
		case *String:
			if o.Value > rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		}
	case token.GreaterEq:
		switch rhs := rhs.(type) {
		case *String:
			if o.Value >= rhs.Value {
				return TrueValue, nil
			}
			return FalseValue, nil
		}
	}
	return nil, ErrInvalidOperator
}

// IsFalsy returns true if the value of the type is falsy.
func (o *String) IsFalsy() bool {
	return len(o.Value) == 0
}

// Copy returns a copy of the type.
func (o *String) Copy() Object {
	return &String{Value: o.Value}
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *String) Equals(x Object) bool {
	t, ok := x.(*String)
	if !ok {
		return false
	}
	return o.Value == t.Value
}

var ansiColorMap = map[string]string {
	"reset" : "\033[0m",           // Text Reset
	// Regular Colors
	"black" : "\033[0;30m",        // black
	"red" : "\033[0;31m",          // red
	"green" : "\033[0;32m",        // green
	"yellow" : "\033[0;33m",       // yellow
	"blue" : "\033[0;34m",         // blue
	"purple" : "\033[0;35m",       // purple
	"cyan" : "\033[0;36m",         // cyan
	"white" : "\033[0;37m",        // white
	// Bold
	"bblack" : "\033[1;30m",       // black
	"bred" : "\033[1;31m",         // red
	"bgreen" : "\033[1;32m",       // green
	"byellow" : "\033[1;33m",      // yellow
	"bblue" : "\033[1;34m",        // blue
	"bpurple" : "\033[1;35m",      // purple
	"bcyan" : "\033[1;36m",        // cyan
	"bwhite" : "\033[1;37m",       // white
	// Underline
	"ublack" : "\033[4;30m",       // black
	"ured" : "\033[4;31m",         // red
	"ugreen" : "\033[4;32m",       // green
	"uyellow" : "\033[4;33m",      // yellow
	"ublue" : "\033[4;34m",        // blue
	"upurple" : "\033[4;35m",      // purple
	"ucyan" : "\033[4;36m",        // cyan
	"uwhite" : "\033[4;37m",       // white
	// Background
	"on_black" : "\033[40m",       // black
	"on_red" : "\033[41m",         // red
	"on_green" : "\033[42m",       // green
	"on_yellow" : "\033[43m",      // yellow
	"on_blue" : "\033[44m",        // blue
	"on_purple" : "\033[45m",      // purple
	"on_cyan" : "\033[46m",        // cyan
	"on_white" : "\033[47m",       // white
	// High Intensty
	"iblack" : "\033[0;90m",       // black
	"ired" : "\033[0;91m",         // red
	"igreen" : "\033[0;92m",       // green
	"iyellow" : "\033[0;93m",      // yellow
	"iblue" : "\033[0;94m",        // blue
	"ipurple" : "\033[0;95m",      // purple
	"icyan" : "\033[0;96m",        // cyan
	"iwhite" : "\033[0;97m",       // white
	// Bold High Intensty
	"biblack" : "\033[1;90m",      // black
	"bired" : "\033[1;91m",        // red
	"bigreen" : "\033[1;92m",      // green
	"biyellow" : "\033[1;93m",     // yellow
	"biblue" : "\033[1;94m",       // blue
	"bipurple" : "\033[1;95m",     // purple
	"bicyan" : "\033[1;96m",       // cyan
	"biwhite" : "\033[1;97m",      // white
	// High Intensty backgrounds
	"on_iblack" : "\033[0;100m",   // black
	"on_ired" : "\033[0;101m",     // red
	"on_igreen" : "\033[0;102m",   // green
	"on_iyellow" : "\033[0;103m",  // yellow
	"on_iblue" : "\033[0;104m",    // blue
	"on_ipurple" : "\033[10;95m",  // purple
	"on_icyan" : "\033[0;106m",    // cyan
	"on_iwhite" : "\033[0;107m",   // white
}

// IndexGet returns a character at a given index.
func (o *String) IndexGet(index Object) (res Object, err error) {
	strIdx, ok := index.(*String) 
	if ok {
		if strIdx.Value == "length" {
			return &Int{Value: int64(len(o.Value))}, nil
		} else if v, ok := ansiColorMap[strIdx.Value]; ok {
			return &String{Value: (v + o.Value + "\033[0m")}, nil
		}
		return nil, nil
	}
	intIdx, ok := index.(*Int)
	if !ok {
		err = ErrInvalidIndexType
		return
	}
	idxVal := int(intIdx.Value)
	if o.runeStr == nil {
		o.runeStr = []rune(o.Value)
	}
	if idxVal < 0 || idxVal >= len(o.runeStr) {
		res = NullValue
		return
	}
	res = &Char{Value: o.runeStr[idxVal]}
	return
}

func (o *String) IndexSet(index, value Object) error {
	char, ok := value.(*Char)
	if !ok {
		return ErrInvalidIndexValueType
	}

    intIdx, ok := index.(*Int)
    if ok {
        if intIdx.Value >= 0 && intIdx.Value < int64(len(o.Value)) {
			runes := []rune(o.Value)
            runes[intIdx.Value] = char.Value
			o.Value = string(runes)
            return nil
        }

        return ErrIndexOutOfBounds
    }

    return ErrInvalidIndexType
}

// Iterate creates a string iterator.
func (o *String) Iterate() Iterator {
	if o.runeStr == nil {
		o.runeStr = []rune(o.Value)
	}
	return &StringIterator{
		v: o.runeStr,
		l: len(o.runeStr),
	}
}

// CanIterate returns whether the Object can be Iterated.
func (o *String) CanIterate() bool {
	return true
}

// Time represents a time value.
type Time struct {
	ObjectImpl
	Value time.Time
}

func (o *Time) String() string {
	// return o.Value.String()
	return o.Value.Format("02/01/2006 3:04PM")
}

// TypeName returns the name of the type.
func (o *Time) TypeName() string {
	return "time"
}

// BinaryOp returns another object that is the result of a given binary
// operator and a right-hand side object.
func (o *Time) BinaryOp(op token.Token, rhs Object) (Object, error) {
	switch rhs := rhs.(type) {
	case *Int:
		switch op {
		case token.Add: // time + int => time
			if rhs.Value == 0 {
				return o, nil
			}
			return &Time{Value: o.Value.Add(time.Duration(rhs.Value))}, nil
		case token.Sub: // time - int => time
			if rhs.Value == 0 {
				return o, nil
			}
			return &Time{Value: o.Value.Add(time.Duration(-rhs.Value))}, nil
		}
	case *Time:
		switch op {
		case token.Sub: // time - time => int (duration)
			return &Int{Value: int64(o.Value.Sub(rhs.Value))}, nil
		case token.Less: // time < time => bool
			if o.Value.Before(rhs.Value) {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.Greater:
			if o.Value.After(rhs.Value) {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.LessEq:
			if o.Value.Equal(rhs.Value) || o.Value.Before(rhs.Value) {
				return TrueValue, nil
			}
			return FalseValue, nil
		case token.GreaterEq:
			if o.Value.Equal(rhs.Value) || o.Value.After(rhs.Value) {
				return TrueValue, nil
			}
			return FalseValue, nil
		}
	}
	return nil, ErrInvalidOperator
}

// Copy returns a copy of the type.
func (o *Time) Copy() Object {
	return &Time{Value: o.Value}
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Time) IsFalsy() bool {
	return o.Value.IsZero()
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *Time) Equals(x Object) bool {
	t, ok := x.(*Time)
	if !ok {
		return false
	}
	return o.Value.Equal(t.Value)
}

// Null represents an null value.
type Null struct {
	ObjectImpl
}

// TypeName returns the name of the type.
func (o *Null) TypeName() string {
	return "null"
}

func (o *Null) String() string {
	return "null"
}

// Copy returns a copy of the type.
func (o *Null) Copy() Object {
	return o
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Null) IsFalsy() bool {
	return true
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *Null) Equals(x Object) bool {
	return o == x
}

// IndexGet returns an element at a given index.
func (o *Null) IndexGet(_ Object) (Object, error) {
	return NullValue, nil
}

// Iterate creates a map iterator.
func (o *Null) Iterate() Iterator {
	return o
}

// CanIterate returns whether the Object can be Iterated.
func (o *Null) CanIterate() bool {
	return true
}

// Next returns true if there are more elements to iterate.
func (o *Null) Next() bool {
	return false
}

// Key returns the key or index value of the current element.
func (o *Null) Key() Object {
	return o
}

// Value returns the value of the current element.
func (o *Null) Value() Object {
	return o
}

// UserFunction represents a user function.
type UserFunction struct {
	ObjectImpl
	Name       string
	Value      CallableFunc
	EncodingID string
}

// TypeName returns the name of the type.
func (o *UserFunction) TypeName() string {
	return "user-function:" + o.Name
}

func (o *UserFunction) String() string {
	return "<user-function>"
}

// Copy returns a copy of the type.
func (o *UserFunction) Copy() Object {
	return &UserFunction{Value: o.Value}
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *UserFunction) Equals(_ Object) bool {
	return false
}

// Call invokes a user function.
func (o *UserFunction) Call(args ...Object) (Object, error) {
	return o.Value(args...)
}

// CanCall returns whether the Object can be Called.
func (o *UserFunction) CanCall() bool {
	return true
}


// Pointer represents a pointer to another object.
type Pointer struct {
    ObjectImpl
    Slot    *Object // points to a variable’s memory slot (real reference)
    Address uintptr
}

func (p *Pointer) TypeName() string { return "pointer" }

func (p *Pointer) String() string {
    if p.Slot == nil || *p.Slot == nil {
        return fmt.Sprintf("&<nil>@%#x", p.Address)
    }
    return fmt.Sprintf("&%s@%#x", (*p.Slot).TypeName(), p.Address)
}

func (p *Pointer) Copy() Object {
    return &Pointer{Slot: p.Slot, Address: p.Address}
}

func (p *Pointer) IsFalsy() bool {
    return p.Slot == nil
}

func (p *Pointer) Equals(x Object) bool {
    q, ok := x.(*Pointer)
    if !ok {
        return false
    }
    return p.Address == q.Address
}


func (p *Pointer) IndexGet(index Object) (Object, error) {
	strIdx, ok := index.(*String)
	if !ok {
		return nil, ErrInvalidIndexType
	}
	switch strIdx.Value {
	case "addr":
		return &Int{Value: int64(p.Address)}, nil
	case "val":
		if *p.Slot == nil {
			return NullValue, nil
		}
		return *p.Slot, nil
	}
	return nil, nil
}