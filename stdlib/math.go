package stdlib

import (
	"math"

	"github.com/2dprototype/tender"
)

var mathModule = map[string]tender.Object{
	"e":       &tender.Float{Value: math.E},
	"pi":      &tender.Float{Value: math.Pi},
	"phi":     &tender.Float{Value: math.Phi},
	"sqrt2":   &tender.Float{Value: math.Sqrt2},
	"sqrtE":   &tender.Float{Value: math.SqrtE},
	"sqrtPi":  &tender.Float{Value: math.SqrtPi},
	"sqrtPhi": &tender.Float{Value: math.SqrtPhi},
	"ln2":     &tender.Float{Value: math.Ln2},
	"log2E":   &tender.Float{Value: math.Log2E},
	"ln10":    &tender.Float{Value: math.Ln10},
	"log10E":  &tender.Float{Value: math.Log10E},
	"abs": &tender.NativeFunction{
		Name:  "abs",
		Value: FuncAFRF(math.Abs),
	},
	"acos": &tender.NativeFunction{
		Name:  "acos",
		Value: FuncAFRF(math.Acos),
	},
	"acosh": &tender.NativeFunction{
		Name:  "acosh",
		Value: FuncAFRF(math.Acosh),
	},
	"asin": &tender.NativeFunction{
		Name:  "asin",
		Value: FuncAFRF(math.Asin),
	},
	"asinh": &tender.NativeFunction{
		Name:  "asinh",
		Value: FuncAFRF(math.Asinh),
	},
	"atan": &tender.NativeFunction{
		Name:  "atan",
		Value: FuncAFRF(math.Atan),
	},
	"atan2": &tender.NativeFunction{
		Name:  "atan2",
		Value: FuncAFFRF(math.Atan2),
	},
	"atanh": &tender.NativeFunction{
		Name:  "atanh",
		Value: FuncAFRF(math.Atanh),
	},
	"cbrt": &tender.NativeFunction{
		Name:  "cbrt",
		Value: FuncAFRF(math.Cbrt),
	},
	"ceil": &tender.NativeFunction{
		Name:  "ceil",
		Value: FuncAFRF(math.Ceil),
	},
	"copysign": &tender.NativeFunction{
		Name:  "copysign",
		Value: FuncAFFRF(math.Copysign),
	},
	"cos": &tender.NativeFunction{
		Name:  "cos",
		Value: FuncAFRF(math.Cos),
	},
	"cosh": &tender.NativeFunction{
		Name:  "cosh",
		Value: FuncAFRF(math.Cosh),
	},
	"dim": &tender.NativeFunction{
		Name:  "dim",
		Value: FuncAFFRF(math.Dim),
	},
	"erf": &tender.NativeFunction{
		Name:  "erf",
		Value: FuncAFRF(math.Erf),
	},
	"erfc": &tender.NativeFunction{
		Name:  "erfc",
		Value: FuncAFRF(math.Erfc),
	},
	"exp": &tender.NativeFunction{
		Name:  "exp",
		Value: FuncAFRF(math.Exp),
	},
	"exp2": &tender.NativeFunction{
		Name:  "exp2",
		Value: FuncAFRF(math.Exp2),
	},
	"expm1": &tender.NativeFunction{
		Name:  "expm1",
		Value: FuncAFRF(math.Expm1),
	},
	"floor": &tender.NativeFunction{
		Name:  "floor",
		Value: FuncAFRF(math.Floor),
	},
	"gamma": &tender.NativeFunction{
		Name:  "gamma",
		Value: FuncAFRF(math.Gamma),
	},
	"hypot": &tender.NativeFunction{
		Name:  "hypot",
		Value: FuncAFFRF(math.Hypot),
	},
	"ilogb": &tender.NativeFunction{
		Name:  "ilogb",
		Value: FuncAFRI(math.Ilogb),
	},
	"inf": &tender.NativeFunction{
		Name:  "inf",
		Value: FuncAIRF(math.Inf),
	},
	"is_inf": &tender.NativeFunction{
		Name:  "is_inf",
		Value: FuncAFIRB(math.IsInf),
	},
	"is_nan": &tender.NativeFunction{
		Name:  "is_nan",
		Value: FuncAFRB(math.IsNaN),
	},
	"j0": &tender.NativeFunction{
		Name:  "j0",
		Value: FuncAFRF(math.J0),
	},
	"j1": &tender.NativeFunction{
		Name:  "j1",
		Value: FuncAFRF(math.J1),
	},
	"jn": &tender.NativeFunction{
		Name:  "jn",
		Value: FuncAIFRF(math.Jn),
	},
	"ldexp": &tender.NativeFunction{
		Name:  "ldexp",
		Value: FuncAFIRF(math.Ldexp),
	},
	"log": &tender.NativeFunction{
		Name:  "log",
		Value: FuncAFRF(math.Log),
	},
	"log10": &tender.NativeFunction{
		Name:  "log10",
		Value: FuncAFRF(math.Log10),
	},
	"log1p": &tender.NativeFunction{
		Name:  "log1p",
		Value: FuncAFRF(math.Log1p),
	},
	"log2": &tender.NativeFunction{
		Name:  "log2",
		Value: FuncAFRF(math.Log2),
	},
	"logb": &tender.NativeFunction{
		Name:  "logb",
		Value: FuncAFRF(math.Logb),
	},
	"max": &tender.NativeFunction{
		Name:  "max",
		Value: FuncAFFRF(math.Max),
	},
	"min": &tender.NativeFunction{
		Name:  "min",
		Value: FuncAFFRF(math.Min),
	},
	"mod": &tender.NativeFunction{
		Name:  "mod",
		Value: FuncAFFRF(math.Mod),
	},
	"nan": &tender.NativeFunction{
		Name:  "nan",
		Value: FuncARF(math.NaN),
	},
	"nextafter": &tender.NativeFunction{
		Name:  "nextafter",
		Value: FuncAFFRF(math.Nextafter),
	},
	"pow": &tender.NativeFunction{
		Name:  "pow",
		Value: FuncAFFRF(math.Pow),
	},
	"pow10": &tender.NativeFunction{
		Name:  "pow10",
		Value: FuncAIRF(math.Pow10),
	},
	"remainder": &tender.NativeFunction{
		Name:  "remainder",
		Value: FuncAFFRF(math.Remainder),
	},
	"signbit": &tender.NativeFunction{
		Name:  "signbit",
		Value: FuncAFRB(math.Signbit),
	},
	"sin": &tender.NativeFunction{
		Name:  "sin",
		Value: FuncAFRF(math.Sin),
	},
	"sinh": &tender.NativeFunction{
		Name:  "sinh",
		Value: FuncAFRF(math.Sinh),
	},
	"sqrt": &tender.NativeFunction{
		Name:  "sqrt",
		Value: FuncAFRF(math.Sqrt),
	},
	"tan": &tender.NativeFunction{
		Name:  "tan",
		Value: FuncAFRF(math.Tan),
	},
	"tanh": &tender.NativeFunction{
		Name:  "tanh",
		Value: FuncAFRF(math.Tanh),
	},
	"trunc": &tender.NativeFunction{
		Name:  "trunc",
		Value: FuncAFRF(math.Trunc),
	},
	"y0": &tender.NativeFunction{
		Name:  "y0",
		Value: FuncAFRF(math.Y0),
	},
	"y1": &tender.NativeFunction{
		Name:  "y1",
		Value: FuncAFRF(math.Y1),
	},
	"yn": &tender.NativeFunction{
		Name:  "yn",
		Value: FuncAIFRF(math.Yn),
	},
}
