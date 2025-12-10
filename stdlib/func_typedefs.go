package stdlib

import (
	"fmt"
	"time"

	"github.com/2dprototype/tender"
)

// FuncAR transform a function of 'func()' signature into CallableFunc type.
func FuncAR(fn func()) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 0 {
			return nil, tender.ErrWrongNumArguments
		}
		fn()
		return tender.NullValue, nil
	}
}

// FuncARI transform a function of 'func() int' signature into CallableFunc
// type.
func FuncARI(fn func() int) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 0 {
			return nil, tender.ErrWrongNumArguments
		}
		return &tender.Int{Value: int64(fn())}, nil
	}
}

// FuncARI64 transform a function of 'func() int64' signature into CallableFunc
// type.
func FuncARI64(fn func() int64) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 0 {
			return nil, tender.ErrWrongNumArguments
		}
		return &tender.Int{Value: fn()}, nil
	}
}


func FuncARI64E(fn func() (int64, error)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 0 {
			return nil, tender.ErrWrongNumArguments
		}
		res, err := fn()
		if err != nil {
			return wrapError(err), nil
		}
		return &tender.Int{Value: res}, nil
	}
}

// FuncAI64RI64 transform a function of 'func(int64) int64' signature into
// CallableFunc type.
func FuncAI64RI64(fn func(int64) int64) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}

		i1, ok := tender.ToInt64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		return &tender.Int{Value: fn(i1)}, nil
	}
}

// FuncAI64R transform a function of 'func(int64)' signature into CallableFunc
// type.
func FuncAI64R(fn func(int64)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}

		i1, ok := tender.ToInt64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		fn(i1)
		return tender.NullValue, nil
	}
}

// FuncARB transform a function of 'func() bool' signature into CallableFunc
// type.
func FuncARB(fn func() bool) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 0 {
			return nil, tender.ErrWrongNumArguments
		}
		if fn() {
			return tender.TrueValue, nil
		}
		return tender.FalseValue, nil
	}
}


func FuncABR(fn func(bool)) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		b, ok := tender.ToBool(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "bool",
				Found:    args[0].TypeName(),
			}
		}
		fn(b)
		return tender.NullValue, nil
	}
}

// FuncARE transform a function of 'func() error' signature into CallableFunc
// type.
func FuncARE(fn func() error) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 0 {
			return nil, tender.ErrWrongNumArguments
		}
		return wrapError(fn()), nil
	}
}

// FuncARS transform a function of 'func() string' signature into CallableFunc
// type.
func FuncARS(fn func() string) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 0 {
			return nil, tender.ErrWrongNumArguments
		}
		s := fn()
		if len(s) > tender.MaxStringLen {
			return nil, tender.ErrStringLimit
		}
		return &tender.String{Value: s}, nil
	}
}

// FuncARSE transform a function of 'func() (string, error)' signature into
// CallableFunc type.
func FuncARSE(fn func() (string, error)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 0 {
			return nil, tender.ErrWrongNumArguments
		}
		res, err := fn()
		if err != nil {
			return wrapError(err), nil
		}
		if len(res) > tender.MaxStringLen {
			return nil, tender.ErrStringLimit
		}
		return &tender.String{Value: res}, nil
	}
}

// FuncARYE transform a function of 'func() ([]byte, error)' signature into
// CallableFunc type.
func FuncARYE(fn func() ([]byte, error)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 0 {
			return nil, tender.ErrWrongNumArguments
		}
		res, err := fn()
		if err != nil {
			return wrapError(err), nil
		}
		if len(res) > tender.MaxBytesLen {
			return nil, tender.ErrBytesLimit
		}
		return &tender.Bytes{Value: res}, nil
	}
}

// FuncARF transform a function of 'func() float64' signature into CallableFunc
// type.
func FuncARF(fn func() float64) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 0 {
			return nil, tender.ErrWrongNumArguments
		}
		return &tender.Float{Value: fn()}, nil
	}
}

func FuncARu8(fn func() uint8) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 0 {
			return nil, tender.ErrWrongNumArguments
		}
		return &tender.Int{Value: int64(fn())}, nil
	}
}

// FuncARSs transform a function of 'func() []string' signature into
// CallableFunc type.
func FuncARSs(fn func() []string) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 0 {
			return nil, tender.ErrWrongNumArguments
		}
		arr := &tender.Array{}
		for _, elem := range fn() {
			if len(elem) > tender.MaxStringLen {
				return nil, tender.ErrStringLimit
			}
			arr.Value = append(arr.Value, &tender.String{Value: elem})
		}
		return arr, nil
	}
}


// FuncARSs transform a function of 'func() []string' signature into
// CallableFunc type.
func FuncASFRSs(fn func(string, float64) []string) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		s, _ := tender.ToString(args[0])
		f, _ := tender.ToFloat64(args[1])
		arr := &tender.Array{}
		for _, elem := range fn(s, f) {
			if len(elem) > tender.MaxStringLen {
				return nil, tender.ErrStringLimit
			}
			arr.Value = append(arr.Value, &tender.String{Value: elem})
		}
		return arr, nil
	}
}


// FuncARIsE transform a function of 'func() ([]int, error)' signature into
// CallableFunc type.
func FuncARIsE(fn func() ([]int, error)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 0 {
			return nil, tender.ErrWrongNumArguments
		}
		res, err := fn()
		if err != nil {
			return wrapError(err), nil
		}
		arr := &tender.Array{}
		for _, v := range res {
			arr.Value = append(arr.Value, &tender.Int{Value: int64(v)})
		}
		return arr, nil
	}
}

// FuncAIRIs transform a function of 'func(int) []int' signature into
// CallableFunc type.
func FuncAIRIs(fn func(int) []int) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		i1, ok := tender.ToInt(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res := fn(i1)
		arr := &tender.Array{}
		for _, v := range res {
			arr.Value = append(arr.Value, &tender.Int{Value: int64(v)})
		}
		return arr, nil
	}
}

// FuncAFRF transform a function of 'func(float64) float64' signature into
// CallableFunc type.
func FuncAFRF(fn func(float64) float64) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		f1, ok := tender.ToFloat64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		return &tender.Float{Value: fn(f1)}, nil
	}
}


// FuncAIR transform a function of 'func(uint)' signature into CallableFunc type.
func FuncAuIR(fn func(uint)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		i1, ok := tender.ToUint(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "uint(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		fn(i1)
		return tender.NullValue, nil
	}
}
// FuncAIR transform a function of 'func(int)' signature into CallableFunc type.
func FuncAIR(fn func(int)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		i1, ok := tender.ToInt(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		fn(i1)
		return tender.NullValue, nil
	}
}

// FuncAIRF transform a function of 'func(int) float64' signature into
// CallableFunc type.
func FuncAIRF(fn func(int) float64) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		i1, ok := tender.ToInt(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		return &tender.Float{Value: fn(i1)}, nil
	}
}

// FuncAFRI transform a function of 'func(float64) int' signature into
// CallableFunc type.
func FuncAFRI(fn func(float64) int) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		f1, ok := tender.ToFloat64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		return &tender.Int{Value: int64(fn(f1))}, nil
	}
}

// FuncAFFRF transform a function of 'func(float64, float64) float64' signature
// into CallableFunc type.
func FuncAFFRF(fn func(float64, float64) float64) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		f1, ok := tender.ToFloat64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		f2, ok := tender.ToFloat64(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "float(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		return &tender.Float{Value: fn(f1, f2)}, nil
	}
}

// FuncAFFRFF transform a function of 'func(float64, float64) (float64, float64)' signature
// into CallableFunc type.
func FuncAFFRFF(fn func(float64, float64) (float64, float64)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		f1, ok := tender.ToFloat64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		f2, ok := tender.ToFloat64(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "float(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		v1, v2 := fn(f1, f2)
		arr := &tender.Array{
			Value: []tender.Object{
				&tender.Float{Value: v1},
				&tender.Float{Value: v2},
			},
		}
		return arr, nil
	}
}


func FuncASRFF(fn func(string) (float64, float64)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		v1, v2 := fn(s1)
		arr := &tender.Array{
			Value: []tender.Object{
				&tender.Float{Value: v1},
				&tender.Float{Value: v2},
			},
		}
		return arr, nil
	}
}


func FuncASFRFF(fn func(string, float64) (float64, float64)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		f1, ok := tender.ToFloat64(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "float(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		v1, v2 := fn(s1, f1)
		arr := &tender.Array{
			Value: []tender.Object{
				&tender.Float{Value: v1},
				&tender.Float{Value: v2},
			},
		}
		return arr, nil
	}
}



// FuncAFR transform a function of 'func(float64)' signature
// into CallableFunc type.
func FuncAFR(fn func(float64)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		f1, ok := tender.ToFloat64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		
		fn(f1)
		
		return &tender.Null{}, nil
	}
}

// FuncAFFR transform a function of 'func(float64, float64)' signature
// into CallableFunc type.
func FuncAFFR(fn func(float64, float64)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		f1, ok := tender.ToFloat64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		f2, ok := tender.ToFloat64(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "float(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		
		fn(f1, f2)
		
		return &tender.Null{}, nil
	}
}


// FuncAFFFR transform a function of 'func(float64, float64, float64)' signature
// into CallableFunc type.
func FuncAFFFR(fn func(float64, float64, float64)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 3 {
			return nil, tender.ErrWrongNumArguments
		}
		f1, ok := tender.ToFloat64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		f2, ok := tender.ToFloat64(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "float(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		f3, ok := tender.ToFloat64(args[2])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "float(compatible)",
				Found:    args[2].TypeName(),
			}
		}
		
		fn(f1, f2, f3)
		return &tender.Null{}, nil
	}
}

func FuncAIIIR(fn func(int, int, int)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 3 {
			return nil, tender.ErrWrongNumArguments
		}
		i1, ok := tender.ToInt(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := tender.ToInt(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		i3, ok := tender.ToInt(args[2])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int(compatible)",
				Found:    args[2].TypeName(),
			}
		}
		
		fn(i1, i2, i3)
		return &tender.Null{}, nil
	}
}

func FuncAIIIIR(fn func(int, int, int, int)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 4 {
			return nil, tender.ErrWrongNumArguments
		}
		i1, ok := tender.ToInt(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := tender.ToInt(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		i3, ok := tender.ToInt(args[2])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int(compatible)",
				Found:    args[2].TypeName(),
			}
		}
		i4, ok := tender.ToInt(args[3])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "fourth",
				Expected: "int(compatible)",
				Found:    args[3].TypeName(),
			}
		}
		
		fn(i1, i2, i3, i4)
		return &tender.Null{}, nil
	}
}



// FuncAFFFFR transform a function of 'func(float64, float64, float64, float64)' signature
// into CallableFunc type.
func FuncAFFFFR(fn func(float64, float64, float64, float64)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 4 {
			return nil, tender.ErrWrongNumArguments
		}
		f1, ok := tender.ToFloat64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		f2, ok := tender.ToFloat64(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "float(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		f3, ok := tender.ToFloat64(args[2])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "float(compatible)",
				Found:    args[2].TypeName(),
			}
		}
		f4, ok := tender.ToFloat64(args[3])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "fourth",
				Expected: "float(compatible)",
				Found:    args[3].TypeName(),
			}
		}
		
		fn(f1, f2, f3, f4)
		return &tender.Null{}, nil
	}
}


// FuncAFFFFFR transform a function of 'func(float64, float64, float64, float64, float64)' signature
// into CallableFunc type.
func FuncAFFFFFR(fn func(float64, float64, float64, float64, float64)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 5 {
			return nil, tender.ErrWrongNumArguments
		}
		f1, ok := tender.ToFloat64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		f2, ok := tender.ToFloat64(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "float(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		f3, ok := tender.ToFloat64(args[2])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "float(compatible)",
				Found:    args[2].TypeName(),
			}
		}
		f4, ok := tender.ToFloat64(args[3])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "fourth",
				Expected: "float(compatible)",
				Found:    args[3].TypeName(),
			}
		}
		
		f5, ok := tender.ToFloat64(args[4])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "fifth",
				Expected: "float(compatible)",
				Found:    args[4].TypeName(),
			}
		}
		
		fn(f1, f2, f3, f4, f5)
		return &tender.Null{}, nil
	}
}


// FuncAFFFFFFR transform a function of 'func(float64, float64, float64, float64, float64, float64)' signature
// into CallableFunc type.
func FuncAFFFFFFR(fn func(float64, float64, float64, float64, float64, float64)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 6 {
			return nil, tender.ErrWrongNumArguments
		}
		f1, ok := tender.ToFloat64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		f2, ok := tender.ToFloat64(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "float(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		f3, ok := tender.ToFloat64(args[2])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "float(compatible)",
				Found:    args[2].TypeName(),
			}
		}
		f4, ok := tender.ToFloat64(args[3])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "fourth",
				Expected: "float(compatible)",
				Found:    args[3].TypeName(),
			}
		}
		f5, ok := tender.ToFloat64(args[4])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "fifth",
				Expected: "float(compatible)",
				Found:    args[4].TypeName(),
			}
		}		
		f6, ok := tender.ToFloat64(args[5])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "sixth",
				Expected: "float(compatible)",
				Found:    args[5].TypeName(),
			}
		}
		
		fn(f1, f2, f3, f4, f5, f6)
		return &tender.Null{}, nil
	}
}


// FuncAIFRF transform a function of 'func(int, float64) float64' signature
// into CallableFunc type.
func FuncAIFRF(fn func(int, float64) float64) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		i1, ok := tender.ToInt(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		f2, ok := tender.ToFloat64(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "float(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		return &tender.Float{Value: fn(i1, f2)}, nil
	}
}

// FuncAFIRF transform a function of 'func(float64, int) float64' signature
// into CallableFunc type.
func FuncAFIRF(fn func(float64, int) float64) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		f1, ok := tender.ToFloat64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := tender.ToInt(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		return &tender.Float{Value: fn(f1, i2)}, nil
	}
}

// FuncAFIRB transform a function of 'func(float64, int) bool' signature
// into CallableFunc type.
func FuncAFIRB(fn func(float64, int) bool) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		f1, ok := tender.ToFloat64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := tender.ToInt(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		if fn(f1, i2) {
			return tender.TrueValue, nil
		}
		return tender.FalseValue, nil
	}
}

// FuncAFRB transform a function of 'func(float64) bool' signature
// into CallableFunc type.
func FuncAFRB(fn func(float64) bool) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		f1, ok := tender.ToFloat64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "float(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		if fn(f1) {
			return tender.TrueValue, nil
		}
		return tender.FalseValue, nil
	}
}


// FuncASFFR transform a function of 'func(string, float, float)' signature into
// CallableFunc type. User function will return 'true' if underlying native
// function returns nil.
func FuncASFFR(fn func(string, float64, float64)) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 3 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		f1, ok := tender.ToFloat64(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "float(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		f2, ok := tender.ToFloat64(args[2])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "float(compatible)",
				Found:    args[2].TypeName(),
			}
		}
		fn(s1, f1, f2)
		return &tender.Null{}, nil
	}
}


func FuncASFFFFR(fn func(string, float64, float64, float64, float64)) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 5 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		f1, ok := tender.ToFloat64(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "float(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		f2, ok := tender.ToFloat64(args[2])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "float(compatible)",
				Found:    args[2].TypeName(),
			}
		}
		f3, ok := tender.ToFloat64(args[3])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "fourth",
				Expected: "float(compatible)",
				Found:    args[3].TypeName(),
			}
		}
		
		f4, ok := tender.ToFloat64(args[4])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "fifth",
				Expected: "float(compatible)",
				Found:    args[4].TypeName(),
			}
		}
		fn(s1, f1, f2, f3, f4)
		return &tender.Null{}, nil
	}
}

// FuncASRS transform a function of 'func(string) string' signature into
// CallableFunc type. User function will return 'true' if underlying native
// function returns nil.
func FuncASRS(fn func(string) string) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s := fn(s1)
		if len(s) > tender.MaxStringLen {
			return nil, tender.ErrStringLimit
		}
		return &tender.String{Value: s}, nil
	}
}


func FuncASRB(fn func(string) bool) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		if fn(s1) {
			return tender.TrueValue, nil
		} else {
			return tender.FalseValue, nil
		}
	}
}

// FuncASRSs transform a function of 'func(string) []string' signature into
// CallableFunc type.
func FuncASRSs(fn func(string) []string) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res := fn(s1)
		arr := &tender.Array{}
		for _, elem := range res {
			if len(elem) > tender.MaxStringLen {
				return nil, tender.ErrStringLimit
			}
			arr.Value = append(arr.Value, &tender.String{Value: elem})
		}
		return arr, nil
	}
}

// FuncASR transform a function of 'func(string)' signature
// native function returns nil.
func FuncASR(fn func(string) ) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		fn(s1)
		return &tender.Null{}, nil
	}
}

// FuncASRSE transform a function of 'func(string) (string, error)' signature
// into CallableFunc type. User function will return 'true' if underlying
// native function returns nil.
func FuncASRSE(fn func(string) (string, error)) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res, err := fn(s1)
		if err != nil {
			return wrapError(err), nil
		}
		if len(res) > tender.MaxStringLen {
			return nil, tender.ErrStringLimit
		}
		return &tender.String{Value: res}, nil
	}
}

// FuncASRE transform a function of 'func(string) error' signature into
// CallableFunc type. User function will return 'true' if underlying native
// function returns nil.
func FuncASRE(fn func(string) error) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		return wrapError(fn(s1)), nil
	}
}

func FuncATRE(fn func(time.Time) error) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, _ := tender.ToTime(args[0])
		return wrapError(fn(s1)), nil
	}
}

func FuncAYFRE(fn func([]byte, float64) error) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		b1, _ := tender.ToByteSlice(args[0])
		f1, _ := tender.ToFloat64(args[1])
		return wrapError(fn(b1, f1)), nil
	}
}

func FuncASFRE(fn func(string, float64) error) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		f1, ok := tender.ToFloat64(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "float(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		return wrapError(fn(s1, f1)), nil
	}
}

// FuncASSRE transform a function of 'func(string, string) error' signature
// into CallableFunc type. User function will return 'true' if underlying
// native function returns nil.
func FuncASSRE(fn func(string, string) error) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s2, ok := tender.ToString(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "string(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		return wrapError(fn(s1, s2)), nil
	}
}

// FuncASSRSs transform a function of 'func(string, string) []string'
// signature into CallableFunc type.
func FuncASSRSs(fn func(string, string) []string) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s2, ok := tender.ToString(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		arr := &tender.Array{}
		for _, res := range fn(s1, s2) {
			if len(res) > tender.MaxStringLen {
				return nil, tender.ErrStringLimit
			}
			arr.Value = append(arr.Value, &tender.String{Value: res})
		}
		return arr, nil
	}
}

// FuncASSIRSs transform a function of 'func(string, string, int) []string'
// signature into CallableFunc type.
func FuncASSIRSs(fn func(string, string, int) []string) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 3 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s2, ok := tender.ToString(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "string(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		i3, ok := tender.ToInt(args[2])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int(compatible)",
				Found:    args[2].TypeName(),
			}
		}
		arr := &tender.Array{}
		for _, res := range fn(s1, s2, i3) {
			if len(res) > tender.MaxStringLen {
				return nil, tender.ErrStringLimit
			}
			arr.Value = append(arr.Value, &tender.String{Value: res})
		}
		return arr, nil
	}
}

// FuncASSRI transform a function of 'func(string, string) int' signature into
// CallableFunc type.
func FuncASSRI(fn func(string, string) int) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s2, ok := tender.ToString(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		return &tender.Int{Value: int64(fn(s1, s2))}, nil
	}
}

// FuncASSRS transform a function of 'func(string, string) string' signature
// into CallableFunc type.
func FuncASSRS(fn func(string, string) string) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s2, ok := tender.ToString(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "string(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		s := fn(s1, s2)
		if len(s) > tender.MaxStringLen {
			return nil, tender.ErrStringLimit
		}
		return &tender.String{Value: s}, nil
	}
}

// FuncASSRB transform a function of 'func(string, string) bool' signature
// into CallableFunc type.
func FuncASSRB(fn func(string, string) bool) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s2, ok := tender.ToString(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "string(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		if fn(s1, s2) {
			return tender.TrueValue, nil
		}
		return tender.FalseValue, nil
	}
}

// FuncASsSRS transform a function of 'func([]string, string) string' signature
// into CallableFunc type.
func FuncASsSRS(fn func([]string, string) string) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		var ss1 []string
		switch arg0 := args[0].(type) {
		case *tender.Array:
			for idx, a := range arg0.Value {
				as, ok := tender.ToString(a)
				if !ok {
					return nil, tender.ErrInvalidArgumentType{
						Name:     fmt.Sprintf("first[%d]", idx),
						Expected: "string(compatible)",
						Found:    a.TypeName(),
					}
				}
				ss1 = append(ss1, as)
			}
		case *tender.ImmutableArray:
			for idx, a := range arg0.Value {
				as, ok := tender.ToString(a)
				if !ok {
					return nil, tender.ErrInvalidArgumentType{
						Name:     fmt.Sprintf("first[%d]", idx),
						Expected: "string(compatible)",
						Found:    a.TypeName(),
					}
				}
				ss1 = append(ss1, as)
			}
		default:
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "array",
				Found:    args[0].TypeName(),
			}
		}
		s2, ok := tender.ToString(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "string(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		s := fn(ss1, s2)
		if len(s) > tender.MaxStringLen {
			return nil, tender.ErrStringLimit
		}
		return &tender.String{Value: s}, nil
	}
}

// FuncASI64RE transform a function of 'func(string, int64) error' signature
// into CallableFunc type.
func FuncASI64RE(fn func(string, int64) error) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := tender.ToInt64(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		return wrapError(fn(s1, i2)), nil
	}
}


// FuncAIIR transform a function of 'func(int, int)' signature
// into CallableFunc type.
func FuncAIIR(fn func(int, int)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		i1, ok := tender.ToInt(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := tender.ToInt(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		
		fn(i1, i2)
		
		return  &tender.Null{}, nil
	}
}


// FuncAIIRE transform a function of 'func(int, int) error' signature
// into CallableFunc type.
func FuncAIIRE(fn func(int, int) error) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		i1, ok := tender.ToInt(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := tender.ToInt(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		return wrapError(fn(i1, i2)), nil
	}
}

func FuncAI64IRI64E(fn func(int64, int) (int64, error)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		i1, ok := tender.ToInt64(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := tender.ToInt(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		
		i3, err := fn(i1, i2)
		
		if err != nil {
			return wrapError(err), nil
		}
		
		return &tender.Int{Value: i3}, nil
	}
}

// FuncASIRS transform a function of 'func(string, int) string' signature
// into CallableFunc type.
func FuncASIRS(fn func(string, int) string) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := tender.ToInt(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		s := fn(s1, i2)
		if len(s) > tender.MaxStringLen {
			return nil, tender.ErrStringLimit
		}
		return &tender.String{Value: s}, nil
	}
}

// FuncASIIRE transform a function of 'func(string, int, int) error' signature
// into CallableFunc type.
func FuncASIIRE(fn func(string, int, int) error) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 3 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i2, ok := tender.ToInt(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		i3, ok := tender.ToInt(args[2])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int(compatible)",
				Found:    args[2].TypeName(),
			}
		}
		return wrapError(fn(s1, i2, i3)), nil
	}
}

// FuncAYRIE transform a function of 'func([]byte) (int, error)' signature
// into CallableFunc type.
func FuncAYRIE(fn func([]byte) (int, error)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		y1, ok := tender.ToByteSlice(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "bytes(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res, err := fn(y1)
		if err != nil {
			return wrapError(err), nil
		}
		return &tender.Int{Value: int64(res)}, nil
	}
}


func FuncAYI64RIE(fn func([]byte, int64) (int, error)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 2 {
			return nil, tender.ErrWrongNumArguments
		}
		y1, ok := tender.ToByteSlice(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "bytes(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		i1, ok := tender.ToInt64(args[1])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int(compatible)",
				Found:    args[1].TypeName(),
			}
		}
		res, err := fn(y1, i1)
		if err != nil {
			return wrapError(err), nil
		}
		return &tender.Int{Value: int64(res)}, nil
	}
}


// FuncAYRS transform a function of 'func([]byte) string' signature into
// CallableFunc type.
func FuncAYRS(fn func([]byte) string) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		y1, ok := tender.ToByteSlice(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "bytes(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res := fn(y1)
		return &tender.String{Value: res}, nil
	}
}

// FuncASRIE transform a function of 'func(string) (int, error)' signature
// into CallableFunc type.
func FuncASRIE(fn func(string) (int, error)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res, err := fn(s1)
		if err != nil {
			return wrapError(err), nil
		}
		return &tender.Int{Value: int64(res)}, nil
	}
}

// FuncASRI transform a function of 'func(string) (int)' signature
// into CallableFunc type.
func FuncASRI(fn func(string) int) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res := fn(s1)
		return &tender.Int{Value: int64(res)}, nil
	}
}

// FuncASRYE transform a function of 'func(string) ([]byte, error)' signature
// into CallableFunc type.
func FuncASRYE(fn func(string) ([]byte, error)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		s1, ok := tender.ToString(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "string(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res, err := fn(s1)
		if err != nil {
			return wrapError(err), nil
		}
		if len(res) > tender.MaxBytesLen {
			return nil, tender.ErrBytesLimit
		}
		return &tender.Bytes{Value: res}, nil
	}
}

// FuncAIRSsE transform a function of 'func(int) ([]string, error)' signature
// into CallableFunc type.
func FuncAIRSsE(fn func(int) ([]string, error)) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		i1, ok := tender.ToInt(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		res, err := fn(i1)
		if err != nil {
			return wrapError(err), nil
		}
		arr := &tender.Array{}
		for _, r := range res {
			if len(r) > tender.MaxStringLen {
				return nil, tender.ErrStringLimit
			}
			arr.Value = append(arr.Value, &tender.String{Value: r})
		}
		return arr, nil
	}
}

// FuncAIRS transform a function of 'func(int) string' signature into
// CallableFunc type.
func FuncAIRS(fn func(int) string) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 1 {
			return nil, tender.ErrWrongNumArguments
		}
		i1, ok := tender.ToInt(args[0])
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "int(compatible)",
				Found:    args[0].TypeName(),
			}
		}
		s := fn(i1)
		if len(s) > tender.MaxStringLen {
			return nil, tender.ErrStringLimit
		}
		return &tender.String{Value: s}, nil
	}
}


func FuncAf64iiR(fn func(float64, int, int) ) tender.CallableFunc {
	return func(args ...tender.Object) (ret tender.Object, err error) {
		if len(args) != 3 {
			return nil, tender.ErrWrongNumArguments
		}
		f1, _ := tender.ToFloat64(args[0])
		i1, _ := tender.ToInt(args[1])
		i2, _ := tender.ToInt(args[2])
		fn(f1, i1, i2)
		return nil, nil
	}
}
