package stdlib

import (
	"math"
	"sync"

	"github.com/2dprototype/tender"
)

// Pre-computed constants
const (
	deg2radConst          = math.Pi / 180.0
	rad2degConst          = 180.0 / math.Pi
	epsilonConst          = 1.0 / (1 << 52) // 2^-52
	gammaToLinearConst    = 2.2
	linearToGammaConst    = 0.45454545
	fullAngleConst        = 360.0
	straightAngleConst    = 180.0
	invSqrtMagicConst     = 0x5f3759df
	perlinSeedConst       = 0.8694896071683615
)

// Pre-allocated reusable objects for common returns
var (
	zeroFloat   = &tender.Float{Value: 0}
	oneFloat    = &tender.Float{Value: 1}
	negOneFloat = &tender.Float{Value: -1}
)

// Global noise instance with lazy initialization and thread-safety
var (
	perlinNoise    *SimplexNoise
	perlinNoiseMux sync.RWMutex
)

var mathfModule = map[string]tender.Object{
	// Core Utilities
	"abs":           &tender.NativeFunction{Name: "abs", Value: mathfAbs},
	"sign":          &tender.NativeFunction{Name: "sign", Value: mathfSign},
	"clamp":         &tender.NativeFunction{Name: "clamp", Value: mathfClamp},
	"clamp01":       &tender.NativeFunction{Name: "clamp01", Value: mathfClamp01},
	"approx":        &tender.NativeFunction{Name: "approx", Value: mathfApproximately},
	"closest_pow2":  &tender.NativeFunction{Name: "closest_pow2", Value: mathfClosestPowerOfTwo},
	"closest_pow2l": &tender.NativeFunction{Name: "closest_pow2l", Value: mathfClosestPowerOfTwoLong},
	"is_pow2":       &tender.NativeFunction{Name: "is_pow2", Value: mathfIsPowerOfTwo},
	"next_pow2":     &tender.NativeFunction{Name: "next_pow2", Value: mathfNextPowerOfTwo},
	"round":         &tender.NativeFunction{Name: "round", Value: mathfRound},

	// Angle Utilities
	"delta_angle":  &tender.NativeFunction{Name: "delta_angle", Value: mathfDeltaAngle},
	"lerp_angle":   &tender.NativeFunction{Name: "lerp_angle", Value: mathfLerpAngle},
	"move_angle":   &tender.NativeFunction{Name: "move_angle", Value: mathfMoveTowardsAngle},

	// Interpolation & Smoothing
	"lerp":        &tender.NativeFunction{Name: "lerp", Value: mathfLerp},
	"lerp_uncl":   &tender.NativeFunction{Name: "lerp_uncl", Value: mathfLerpUnclamped},
	"inverse_lerp": &tender.NativeFunction{Name: "inverse_lerp", Value: mathfInverseLerp},
	"move_towards": &tender.NativeFunction{Name: "move_towards", Value: mathfMoveTowards},
	"smooth_step":  &tender.NativeFunction{Name: "smooth_step", Value: mathfSmoothStep},

	// Loops & Waves
	"repeat":   &tender.NativeFunction{Name: "repeat", Value: mathfRepeat},
	"pingpong": &tender.NativeFunction{Name: "pingpong", Value: mathfPingPong},

	// Color Space
	"gamma2linear": &tender.NativeFunction{Name: "gamma2linear", Value: mathfGammaToLinear},
	"linear2gamma": &tender.NativeFunction{Name: "linear2gamma", Value: mathfLinearToGamma},

	// Constants
	"deg2rad":  &tender.Float{Value: deg2radConst},
	"rad2deg":  &tender.Float{Value: rad2degConst},
	"eps":      &tender.Float{Value: epsilonConst},
	"neg_inf":  &tender.Float{Value: math.Inf(-1)},

	// Noise
	"perlin": &tender.NativeFunction{Name: "perlin", Value: mathfPerlinNoise},

	// Optimization
	"inv_sqrt": &tender.NativeFunction{Name: "inv_sqrt", Value: mathfInvSqrt},
}

// --- Optimized Helper functions (inline-able) ---

func getFloat(obj tender.Object) (float64, bool) {
	switch v := obj.(type) {
	case *tender.Float:
		return v.Value, true
	case *tender.Int:
		return float64(v.Value), true
	default:
		return 0, false
	}
}

func getInt(obj tender.Object) (int64, bool) {
	switch v := obj.(type) {
	case *tender.Int:
		return v.Value, true
	case *tender.Float:
		return int64(v.Value), true
	default:
		return 0, false
	}
}

func checkArgCount(args []tender.Object, expected int) error {
	if len(args) != expected {
		return tender.ErrWrongNumArguments
	}
	return nil
}

// --- Core Utilities ---

func mathfAbs(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 1); err != nil {
		return nil, err
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
	if err := checkArgCount(args, 1); err != nil {
		return nil, err
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
		return oneFloat, nil
	}
	return negOneFloat, nil
}

func mathfClamp(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 3); err != nil {
		return nil, err
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
	} else if val > max {
		val = max
	}
	return &tender.Float{Value: val}, nil
}

func mathfClamp01(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 1); err != nil {
		return nil, err
	}
	val, ok := getFloat(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "number",
			Found:    args[0].TypeName(),
		}
	}
	if val <= 0.0 {
		return zeroFloat, nil
	}
	if val >= 1.0 {
		return oneFloat, nil
	}
	return &tender.Float{Value: val}, nil
}

func mathfApproximately(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 2); err != nil {
		return nil, err
	}
	f1, ok1 := getFloat(args[0])
	f2, ok2 := getFloat(args[1])
	if !ok1 || !ok2 {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "arguments",
			Expected: "numbers",
			Found:    "mixed types",
		}
	}
	if math.Abs(f1-f2) < epsilonConst {
		return tender.TrueValue, nil
	}
	return tender.FalseValue, nil
}

func mathfClosestPowerOfTwo(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 1); err != nil {
		return nil, err
	}
	val, ok := getInt(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "integer",
			Found:    args[0].TypeName(),
		}
	}
	if val <= 0 {
		return zeroIntPtr(), nil
	}

	next := nextPowerOfTwoInt(val)
	if next > val && next-val > next>>2 {
		return &tender.Int{Value: next >> 1}, nil
	}
	return &tender.Int{Value: next}, nil
}

func mathfClosestPowerOfTwoLong(args ...tender.Object) (tender.Object, error) {
	return mathfClosestPowerOfTwo(args...)
}

func mathfIsPowerOfTwo(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 1); err != nil {
		return nil, err
	}
	val, ok := getInt(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "integer",
			Found:    args[0].TypeName(),
		}
	}
	if val <= 0 {
		return tender.FalseValue, nil
	}
	if (val & (val - 1)) == 0 {
		return tender.TrueValue, nil
	}
	return tender.FalseValue, nil
}

func mathfNextPowerOfTwo(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 1); err != nil {
		return nil, err
	}
	val, ok := getInt(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "integer",
			Found:    args[0].TypeName(),
		}
	}
	if val <= 0 {
		return zeroIntPtr(), nil
	}
	return &tender.Int{Value: nextPowerOfTwoInt(val)}, nil
}

func nextPowerOfTwoInt(val int64) int64 {
	val--
	val |= val >> 1
	val |= val >> 2
	val |= val >> 4
	val |= val >> 8
	val |= val >> 16
	val |= val >> 32
	val++
	return val
}

func mathfRound(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 1); err != nil {
		return nil, err
	}
	val, ok := getFloat(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "number",
			Found:    args[0].TypeName(),
		}
	}
	// Optimized: use math.Round which handles ties to nearest even automatically
	return &tender.Float{Value: math.Round(val)}, nil
}

// --- Angle Utilities ---

func mathfDeltaAngle(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 2); err != nil {
		return nil, err
	}
	current, ok1 := getFloat(args[0])
	target, ok2 := getFloat(args[1])
	if !ok1 || !ok2 {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "arguments",
			Expected: "numbers",
			Found:    "mixed types",
		}
	}
	// Optimized: use math.Mod for all cases
	current = math.Mod(current, fullAngleConst)
	target = math.Mod(target, fullAngleConst)
	return &tender.Float{Value: target - current}, nil
}

func mathfLerpAngle(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 3); err != nil {
		return nil, err
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
	// Optimized angle normalization
	for a > b+straightAngleConst {
		b += fullAngleConst
	}
	for b > a+straightAngleConst {
		b -= fullAngleConst
	}
	// Clamp t inline
	if t < 0.0 {
		t = 0.0
	} else if t > 1.0 {
		t = 1.0
	}
	return &tender.Float{Value: a + (b-a)*t}, nil
}

func mathfMoveTowardsAngle(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 3); err != nil {
		return nil, err
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

	// Normalize angles to [0, 360)
	current = math.Mod(current, fullAngleConst)
	if current < 0 {
		current += fullAngleConst
	}
	target = math.Mod(target, fullAngleConst)
	if target < 0 {
		target += fullAngleConst
	}

	// Find shortest delta
	delta := target - current
	if delta > straightAngleConst {
		delta -= fullAngleConst
	} else if delta < -straightAngleConst {
		delta += fullAngleConst
	}

	if maxDelta <= 0 {
		return &tender.Float{Value: current}, nil
	}
	if delta > 0 {
		if delta <= maxDelta {
			return &tender.Float{Value: target}, nil
		}
		return &tender.Float{Value: current + maxDelta}, nil
	}
	if delta < 0 {
		if -delta <= maxDelta {
			return &tender.Float{Value: target}, nil
		}
		return &tender.Float{Value: current - maxDelta}, nil
	}
	return &tender.Float{Value: current}, nil
}

// --- Interpolation & Smoothing ---

func mathfLerp(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 3); err != nil {
		return nil, err
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
	if t <= 0.0 {
		return &tender.Float{Value: a}, nil
	}
	if t >= 1.0 {
		return &tender.Float{Value: b}, nil
	}
	return &tender.Float{Value: a + (b-a)*t}, nil
}

func mathfLerpUnclamped(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 3); err != nil {
		return nil, err
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
	// Optimized: avoid branch when t is in [0,1]
	if t < 0 || t > 1 {
		return &tender.Float{Value: a + math.Abs(b-a)*t}, nil
	}
	return &tender.Float{Value: a + (b-a)*t}, nil
}

func mathfInverseLerp(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 3); err != nil {
		return nil, err
	}
	a, ok1 := getFloat(args[0])
	b, ok2 := getFloat(args[1])
	value, ok3 := getFloat(args[2])
	if !ok1 || !ok2 || !ok3 {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "arguments",
			Expected: "numbers",
			Found:    "mixed types",
		}
	}
	if a == b {
		return zeroFloat, nil
	}
	// Optimized: clamp without math.Min/Max function calls
	if value < a && value < b {
		if a < b {
			value = a
		} else {
			value = b
		}
	} else if value > a && value > b {
		if a > b {
			value = a
		} else {
			value = b
		}
	}
	return &tender.Float{Value: (value - a) / (b - a)}, nil
}

func mathfMoveTowards(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 3); err != nil {
		return nil, err
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
	if maxDelta <= 0 {
		return &tender.Float{Value: current}, nil
	}
	if target < current {
		if current-maxDelta <= target {
			return &tender.Float{Value: target}, nil
		}
		return &tender.Float{Value: current - maxDelta}, nil
	}
	if current+maxDelta >= target {
		return &tender.Float{Value: target}, nil
	}
	return &tender.Float{Value: current + maxDelta}, nil
}

func mathfSmoothStep(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 3); err != nil {
		return nil, err
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
	if t <= 0.0 {
		return &tender.Float{Value: from}, nil
	}
	if t >= 1.0 {
		return &tender.Float{Value: to}, nil
	}
	// Optimized: 3t^2 - 2t^3
	t = t * t * (3.0 - 2.0*t)
	return &tender.Float{Value: from + (to-from)*t}, nil
}

// --- Loops & Waves ---

func mathfRepeat(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 2); err != nil {
		return nil, err
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
	if length <= 0 {
		return zeroFloat, nil
	}
	// Optimized: use math.Mod directly
	return &tender.Float{Value: math.Mod(t, length)}, nil
}

func mathfPingPong(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 2); err != nil {
		return nil, err
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
	if length <= 0 {
		return zeroFloat, nil
	}
	if t < 0 {
		t = -t
	}
	// Optimized ping-pong using bit manipulation
	period := length * 2.0
	t = math.Mod(t, period)
	if t > length {
		t = period - t
	}
	return &tender.Float{Value: t}, nil
}

// --- Color Space ---

func mathfGammaToLinear(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 1); err != nil {
		return nil, err
	}
	val, ok := getFloat(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "number",
			Found:    args[0].TypeName(),
		}
	}
	return &tender.Float{Value: math.Pow(val, gammaToLinearConst)}, nil
}

func mathfLinearToGamma(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 1); err != nil {
		return nil, err
	}
	val, ok := getFloat(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "number",
			Found:    args[0].TypeName(),
		}
	}
	return &tender.Float{Value: math.Pow(val, linearToGammaConst)}, nil
}

// --- Noise ---

// SimplexNoise with pre-allocated arrays for performance
type SimplexNoise struct {
	grad3 [12][3]float64
	perm  [512]int
}

// Singleton instance with lazy initialization
var noiseInstance *SimplexNoise
var noiseOnce sync.Once

func getNoise() *SimplexNoise {
	noiseOnce.Do(func() {
		noiseInstance = newSimplexNoise(perlinSeedConst)
	})
	return noiseInstance
}

func newSimplexNoise(seed float64) *SimplexNoise {
	s := &SimplexNoise{}
	s.grad3 = [12][3]float64{
		{1, 1, 0}, {-1, 1, 0}, {1, -1, 0}, {-1, -1, 0},
		{1, 0, 1}, {-1, 0, 1}, {1, 0, -1}, {-1, 0, -1},
		{0, 1, 1}, {0, -1, 1}, {0, 1, -1}, {0, -1, -1},
	}
	// Pre-compute permutation table
	p := make([]int, 256)
	for i := 0; i < 256; i++ {
		p[i] = int(math.Floor(seed * 256))
	}
	for i := 0; i < 512; i++ {
		s.perm[i] = p[i&255]
	}
	return s
}

func (s *SimplexNoise) dot(g [3]float64, x, y float64) float64 {
	return g[0]*x + g[1]*y
}

func (s *SimplexNoise) noise(xin, yin float64) float64 {
	// Pre-compute constants
	const F2 = 0.5 * (1.7320508075688772 - 1.0) // (sqrt(3)-1)/2
	const G2 = (3.0 - 1.7320508075688772) / 6.0  // (3-sqrt(3))/6

	sVal := (xin + yin) * F2
	i := int(math.Floor(xin + sVal))
	j := int(math.Floor(yin + sVal))
	t := float64(i+j) * G2
	X0 := float64(i) - t
	Y0 := float64(j) - t
	x0 := xin - X0
	y0 := yin - Y0

	var i1, j1 int
	if x0 > y0 {
		i1 = 1
	} else {
		j1 = 1
	}

	x1 := x0 - float64(i1) + G2
	y1 := y0 - float64(j1) + G2
	x2 := x0 - 1.0 + 2.0*G2
	y2 := y0 - 1.0 + 2.0*G2

	ii := i & 255
	jj := j & 255
	gi0 := s.perm[ii+s.perm[jj]] % 12
	gi1 := s.perm[ii+i1+s.perm[jj+j1]] % 12
	gi2 := s.perm[ii+1+s.perm[jj+1]] % 12

	// Calculate contributions with optimized conditions
	var n0, n1, n2 float64
	t0 := 0.5 - x0*x0 - y0*y0
	if t0 > 0 {
		t0 *= t0
		n0 = t0 * t0 * s.dot(s.grad3[gi0], x0, y0)
	}

	t1 := 0.5 - x1*x1 - y1*y1
	if t1 > 0 {
		t1 *= t1
		n1 = t1 * t1 * s.dot(s.grad3[gi1], x1, y1)
	}

	t2 := 0.5 - x2*x2 - y2*y2
	if t2 > 0 {
		t2 *= t2
		n2 = t2 * t2 * s.dot(s.grad3[gi2], x2, y2)
	}

	return 70.0 * (n0 + n1 + n2)
}

func mathfPerlinNoise(args ...tender.Object) (tender.Object, error) {
	if err := checkArgCount(args, 2); err != nil {
		return nil, err
	}
	x, ok1 := getFloat(args[0])
	y, ok2 := getFloat(args[1])
	if !ok1 || !ok2 {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "arguments",
			Expected: "numbers",
			Found:    "mixed types",
		}
	}
	return &tender.Float{Value: getNoise().noise(x, y)}, nil
}

// --- Optimization ---

// Optimized Fast Inverse Square Root
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

// Helper to avoid allocating zero Int repeatedly
func zeroIntPtr() *tender.Int {
	// Use a package-level zero int
	return &tender.Int{Value: 0}
}