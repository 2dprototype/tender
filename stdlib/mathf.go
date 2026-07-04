package stdlib

import (
	"math"
	"github.com/2dprototype/tender"
)

var mathfModule = map[string]tender.Object{
	// Core Utilities
	"abs":   &tender.UserFunction{Name: "abs", Value: mathfAbs},
	"sign":  &tender.UserFunction{Name: "sign", Value: mathfSign},
	"clamp": &tender.UserFunction{Name: "clamp", Value: mathfClamp},
	"clamp01": &tender.UserFunction{Name: "clamp01", Value: mathfClamp01},

	// Interpolation & Smoothing
	"lerp":           &tender.UserFunction{Name: "lerp", Value: mathfLerp},
	"lerp_unclamped": &tender.UserFunction{Name: "lerp_unclamped", Value: mathfLerpUnclamped},
	"move_towards":   &tender.UserFunction{Name: "move_towards", Value: mathfMoveTowards},
	"smooth_step":    &tender.UserFunction{Name: "smooth_step", Value: mathfSmoothStep},

	// Loops & Waves
	"repeat":    &tender.UserFunction{Name: "repeat", Value: mathfRepeat},
	"pingpong":  &tender.UserFunction{Name: "pingpong", Value: mathfPingPong},
	"deg2rad":   &tender.Float{Value: math.Pi / 180.0},
	"rad2deg":   &tender.Float{Value: 180.0 / math.Pi},

	// Optimization Tricks
	"inv_sqrt": &tender.UserFunction{Name: "inv_sqrt", Value: mathfInvSqrt},
}

// --- Helper to extract float64 from any numeric Object ---
func getFloat(obj tender.Object) (float64, bool) {
	if f, ok := tender.ToFloat64(obj); ok {
		return f, true
	}
	return 0, false
}

// --- Core Utilities ---

func mathfAbs(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	val, ok := getFloat(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "number",
			Found:    args[0].TypeName(),
		}
	}
	return &tender.Float{Value: math.Abs(val)}, nil
}

func mathfSign(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	val, ok := getFloat(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "number",
			Found:    args[0].TypeName(),
		}
	}
	if val >= 0 {
		return &tender.Float{Value: 1.0}, nil
	}
	return &tender.Float{Value: -1.0}, nil
}

func mathfClamp(args ...tender.Object) (tender.Object, error) {
	if len(args) != 3 {
		return nil, tender.ErrWrongNumArguments
	}
	val, ok1 := getFloat(args[0])
	min, ok2 := getFloat(args[1])
	max, ok3 := getFloat(args[2])
	if !ok1 || !ok2 || !ok3 {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "arguments",
			Expected: "numbers",
			Found:    "mixed types",
		}
	}
	if val < min {
		val = min
	}
	if val > max {
		val = max
	}
	return &tender.Float{Value: val}, nil
}

func mathfClamp01(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	val, ok := getFloat(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "number",
			Found:    args[0].TypeName(),
		}
	}
	if val < 0.0 {
		val = 0.0
	}
	if val > 1.0 {
		val = 1.0
	}
	return &tender.Float{Value: val}, nil
}

// --- Interpolation & Smoothing ---

func mathfLerp(args ...tender.Object) (tender.Object, error) {
	if len(args) != 3 {
		return nil, tender.ErrWrongNumArguments
	}
	a, ok1 := getFloat(args[0])
	b, ok2 := getFloat(args[1])
	t, ok3 := getFloat(args[2])
	if !ok1 || !ok2 || !ok3 {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "arguments",
			Expected: "numbers",
			Found:    "mixed types",
		}
	}
	if t < 0.0 {
		t = 0.0
	}
	if t > 1.0 {
		t = 1.0
	}
	return &tender.Float{Value: a + (b-a)*t}, nil
}

func mathfLerpUnclamped(args ...tender.Object) (tender.Object, error) {
	if len(args) != 3 {
		return nil, tender.ErrWrongNumArguments
	}
	a, ok1 := getFloat(args[0])
	b, ok2 := getFloat(args[1])
	t, ok3 := getFloat(args[2])
	if !ok1 || !ok2 || !ok3 {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "arguments",
			Expected: "numbers",
			Found:    "mixed types",
		}
	}
	return &tender.Float{Value: a + (b-a)*t}, nil
}

func mathfMoveTowards(args ...tender.Object) (tender.Object, error) {
	if len(args) != 3 {
		return nil, tender.ErrWrongNumArguments
	}
	current, ok1 := getFloat(args[0])
	target, ok2 := getFloat(args[1])
	maxDelta, ok3 := getFloat(args[2])
	if !ok1 || !ok2 || !ok3 {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "arguments",
			Expected: "numbers",
			Found:    "mixed types",
		}
	}
	if math.Abs(target-current) <= maxDelta {
		return &tender.Float{Value: target}, nil
	}
	if target > current {
		return &tender.Float{Value: current + maxDelta}, nil
	}
	return &tender.Float{Value: current - maxDelta}, nil
}

func mathfSmoothStep(args ...tender.Object) (tender.Object, error) {
	if len(args) != 3 {
		return nil, tender.ErrWrongNumArguments
	}
	from, ok1 := getFloat(args[0])
	to, ok2 := getFloat(args[1])
	t, ok3 := getFloat(args[2])
	if !ok1 || !ok2 || !ok3 {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "arguments",
			Expected: "numbers",
			Found:    "mixed types",
		}
	}
	if t < 0.0 {
		t = 0.0
	}
	if t > 1.0 {
		t = 1.0
	}
	// Smooth Hermite interpolation: 3t^2 - 2t^3
	t = t * t * (3.0 - 2.0*t)
	return &tender.Float{Value: from + (to-from)*t}, nil
}

// --- Loops & Waves ---

func mathfRepeat(args ...tender.Object) (tender.Object, error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}
	t, ok1 := getFloat(args[0])
	length, ok2 := getFloat(args[1])
	if !ok1 || !ok2 {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "arguments",
			Expected: "numbers",
			Found:    "mixed types",
		}
	}
	val := t - math.Floor(t/length)*length
	return &tender.Float{Value: val}, nil
}

func mathfPingPong(args ...tender.Object) (tender.Object, error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}
	t, ok1 := getFloat(args[0])
	length, ok2 := getFloat(args[1])
	if !ok1 || !ok2 {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "arguments",
			Expected: "numbers",
			Found:    "mixed types",
		}
	}
	// Loop twice the length, then bounce back via absolute difference
	t = t - math.Floor(t/(length*2.0))*(length*2.0)
	val := length - math.Abs(t-length)
	return &tender.Float{Value: val}, nil
}

// --- Optimization ---

// Fast Inverse Square Root (Quake III bit-hack adapted for Go 64-bit to 32-bit compilation)
func mathfInvSqrt(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	val, ok := getFloat(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "number",
			Found:    args[0].TypeName(),
		}
	}
	x32 := float32(val)
	xhalf := float32(0.5) * x32
	i := math.Float32bits(x32)
	i = 0x5f3759df - (i >> 1) // The magic number
	x32 = math.Float32frombits(i)
	x32 = x32 * (float32(1.5) - xhalf*x32*x32) // 1st Newton-Raphson iteration
	return &tender.Float{Value: float64(x32)}, nil
}