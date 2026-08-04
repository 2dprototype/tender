package tender

import (
	"math"
	"unsafe"
)

// Value represents a NaN-tagged 64-bit packed value for high-performance
// zero-allocation runtime operations in Tender.
type Value uint64

const (
	valNanMask uint64 = 0x7FF8000000000000
	valTagInt  uint64 = 0xFFF9000000000000
	valTagBool uint64 = 0xFFFA000000000000
	valTagNull uint64 = 0xFFFB000000000000
	valTagChar uint64 = 0xFFFC000000000000
	valTagPtr  uint64 = 0xFFFD000000000000
	valPtrMask uint64 = 0x0000FFFFFFFFFFFF
)

// NewFloatValue creates a Value from a float64.
func NewFloatValue(f float64) Value {
	return Value(math.Float64bits(f))
}

// NewIntValue creates a Value from a 32-bit int.
func NewIntValue(i int32) Value {
	return Value(valTagInt | uint64(uint32(i)))
}

// NewBoolValue creates a Value from a bool.
func NewBoolValue(b bool) Value {
	if b {
		return Value(valTagBool | 1)
	}
	return Value(valTagBool | 0)
}

// NullVal is the NaN-tagged representation of Null.
const NullVal = Value(valTagNull)

// IsFloat returns true if v is a float64.
func (v Value) IsFloat() bool {
	return (uint64(v) & valNanMask) != valNanMask
}

// IsInt returns true if v is an int32.
func (v Value) IsInt() bool {
	return (uint64(v) & 0xFFFF000000000000) == valTagInt
}

// IsBool returns true if v is a bool.
func (v Value) IsBool() bool {
	return (uint64(v) & 0xFFFF000000000000) == valTagBool
}

// IsNull returns true if v is null.
func (v Value) IsNull() bool {
	return uint64(v) == valTagNull
}

// Float returns the float64 value.
func (v Value) Float() float64 {
	return math.Float64frombits(uint64(v))
}

// Int returns the int32 value.
func (v Value) Int() int32 {
	return int32(uint32(v))
}

// Bool returns the bool value.
func (v Value) Bool() bool {
	return (uint64(v) & 1) != 0
}

// ToObject converts Value back to an Object interface.
func (v Value) ToObject() Object {
	if v.IsFloat() {
		return &Float{Value: v.Float()}
	}
	if v.IsInt() {
		return MakeInt(int64(v.Int()))
	}
	if v.IsBool() {
		if v.Bool() {
			return TrueValue
		}
		return FalseValue
	}
	if v.IsNull() {
		return NullValue
	}
	ptr := unsafe.Pointer(uintptr(uint64(v) & valPtrMask))
	return *(*Object)(unsafe.Pointer(&ptr))
}

// FromObject packs an Object interface into a NaN-tagged Value.
func FromObject(obj Object) Value {
	if obj == nil || obj == NullValue {
		return NullVal
	}
	switch val := obj.(type) {
	case *Int:
		if val.Value >= math.MinInt32 && val.Value <= math.MaxInt32 {
			return NewIntValue(int32(val.Value))
		}
	case *Float:
		return NewFloatValue(val.Value)
	case *Bool:
		return NewBoolValue(val.value)
	}
	ptr := *(*uintptr)(unsafe.Pointer(&obj))
	return Value(valTagPtr | (uint64(ptr) & valPtrMask))
}
