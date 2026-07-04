package gl

import (
	"errors"

	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/v/gl"
)

var (
	ErrInvalidArgCount     = errors.New("invalid number of arguments")
	ErrInvalidArgumentType = errors.New("invalid argument type")
)

// Helper functions to safely extract underlying Go types from Tender Objects
func getFloat64(obj tender.Object) (float64, bool) {
	if f, ok := obj.(*tender.Float); ok {
		return f.Value, true
	}
	if i, ok := obj.(*tender.Int); ok {
		return float64(i.Value), true
	}
	return 0, false
}

func getInt64(obj tender.Object) (int64, bool) {
	if i, ok := obj.(*tender.Int); ok {
		return i.Value, true
	}
	return 0, false
}

func getInt(obj tender.Object) (int, bool) {
	if i, ok := obj.(*tender.Int); ok {
		return int(i.Value), true
	}
	return 0, false
}

func getUint32(obj tender.Object) (uint32, bool) {
	if i, ok := obj.(*tender.Int); ok {
		return uint32(i.Value), true
	}
	return 0, false
}

func getFloat32(obj tender.Object) (float32, bool) {
	if f, ok := obj.(*tender.Float); ok {
		return float32(f.Value), true
	}
	if i, ok := obj.(*tender.Int); ok {
		return float32(i.Value), true
	}
	return 0, false
}

// Module exposes the OpenGL API to Tender scripts
var Module = &tender.BuiltinModule{
	Attrs: map[string]tender.Object{
		// ==================== Matrix Mode Constants ====================
		"MODELVIEW":  &tender.Int{Value: int64(gl.MODELVIEW)},
		"PROJECTION": &tender.Int{Value: int64(gl.PROJECTION)},
		"TEXTURE":    &tender.Int{Value: int64(gl.TEXTURE)},

		// ==================== Clear Buffer Bits ====================
		"COLOR_BUFFER_BIT":   &tender.Int{Value: int64(gl.COLOR_BUFFER_BIT)},
		"DEPTH_BUFFER_BIT":   &tender.Int{Value: int64(gl.DEPTH_BUFFER_BIT)},
		"ACCUM_BUFFER_BIT":   &tender.Int{Value: int64(gl.ACCUM_BUFFER_BIT)},
		"STENCIL_BUFFER_BIT": &tender.Int{Value: int64(gl.STENCIL_BUFFER_BIT)},

		// ==================== Primitive Types ====================
		"POINTS":         &tender.Int{Value: int64(gl.POINTS)},
		"LINES":          &tender.Int{Value: int64(gl.LINES)},
		"LINE_LOOP":      &tender.Int{Value: int64(gl.LINE_LOOP)},
		"LINE_STRIP":     &tender.Int{Value: int64(gl.LINE_STRIP)},
		"TRIANGLES":      &tender.Int{Value: int64(gl.TRIANGLES)},
		"TRIANGLE_STRIP": &tender.Int{Value: int64(gl.TRIANGLE_STRIP)},
		"TRIANGLE_FAN":   &tender.Int{Value: int64(gl.TRIANGLE_FAN)},
		"QUADS":          &tender.Int{Value: int64(gl.QUADS)},
		"QUAD_STRIP":     &tender.Int{Value: int64(gl.QUAD_STRIP)},
		"POLYGON":        &tender.Int{Value: int64(gl.POLYGON)},

		// ==================== Shading Models ====================
		"FLAT":   &tender.Int{Value: int64(gl.FLAT)},
		"SMOOTH": &tender.Int{Value: int64(gl.SMOOTH)},

		// ==================== Enabling Features ====================
		"BLEND":        &tender.Int{Value: int64(gl.BLEND)},
		"DEPTH_TEST":   &tender.Int{Value: int64(gl.DEPTH_TEST)},
		"CULL_FACE":    &tender.Int{Value: int64(gl.CULL_FACE)},
		"LIGHTING":     &tender.Int{Value: int64(gl.LIGHTING)},
		"LIGHT0":       &tender.Int{Value: int64(gl.LIGHT0)},
		"LIGHT1":       &tender.Int{Value: int64(gl.LIGHT1)},
		"LIGHT2":       &tender.Int{Value: int64(gl.LIGHT2)},
		"LIGHT3":       &tender.Int{Value: int64(gl.LIGHT3)},
		"LIGHT4":       &tender.Int{Value: int64(gl.LIGHT4)},
		"LIGHT5":       &tender.Int{Value: int64(gl.LIGHT5)},
		"LIGHT6":       &tender.Int{Value: int64(gl.LIGHT6)},
		"LIGHT7":       &tender.Int{Value: int64(gl.LIGHT7)},
		"TEXTURE_2D":   &tender.Int{Value: int64(gl.TEXTURE_2D)},
		"FOG":          &tender.Int{Value: int64(gl.FOG)},
		"SCISSOR_TEST": &tender.Int{Value: int64(gl.SCISSOR_TEST)},
		"STENCIL_TEST": &tender.Int{Value: int64(gl.STENCIL_TEST)},
		"ALPHA_TEST":   &tender.Int{Value: int64(gl.ALPHA_TEST)},
		"NORMALIZE":    &tender.Int{Value: int64(gl.NORMALIZE)},
		"COLOR_MATERIAL": &tender.Int{Value: int64(gl.COLOR_MATERIAL)},

		// ==================== Blend Factors ====================
		"ZERO":                  &tender.Int{Value: int64(gl.ZERO)},
		"ONE":                   &tender.Int{Value: int64(gl.ONE)},
		"SRC_COLOR":             &tender.Int{Value: int64(gl.SRC_COLOR)},
		"ONE_MINUS_SRC_COLOR":   &tender.Int{Value: int64(gl.ONE_MINUS_SRC_COLOR)},
		"DST_COLOR":             &tender.Int{Value: int64(gl.DST_COLOR)},
		"ONE_MINUS_DST_COLOR":   &tender.Int{Value: int64(gl.ONE_MINUS_DST_COLOR)},
		"SRC_ALPHA":             &tender.Int{Value: int64(gl.SRC_ALPHA)},
		"ONE_MINUS_SRC_ALPHA":   &tender.Int{Value: int64(gl.ONE_MINUS_SRC_ALPHA)},
		"DST_ALPHA":             &tender.Int{Value: int64(gl.DST_ALPHA)},
		"ONE_MINUS_DST_ALPHA":   &tender.Int{Value: int64(gl.ONE_MINUS_DST_ALPHA)},
		"SRC_ALPHA_SATURATE":    &tender.Int{Value: int64(gl.SRC_ALPHA_SATURATE)},

		// ==================== Cull Face Modes ====================
		"FRONT":          &tender.Int{Value: int64(gl.FRONT)},
		"BACK":           &tender.Int{Value: int64(gl.BACK)},
		"FRONT_AND_BACK": &tender.Int{Value: int64(gl.FRONT_AND_BACK)},

		// ==================== Polygon Modes ====================
		"POINT": &tender.Int{Value: int64(gl.POINT)},
		"LINE":  &tender.Int{Value: int64(gl.LINE)},
		"FILL":  &tender.Int{Value: int64(gl.FILL)},

		// ==================== Depth Function ====================
		"NEVER":    &tender.Int{Value: int64(gl.NEVER)},
		"LESS":     &tender.Int{Value: int64(gl.LESS)},
		"EQUAL":    &tender.Int{Value: int64(gl.EQUAL)},
		"LEQUAL":   &tender.Int{Value: int64(gl.LEQUAL)},
		"GREATER":  &tender.Int{Value: int64(gl.GREATER)},
		"NOTEQUAL": &tender.Int{Value: int64(gl.NOTEQUAL)},
		"GEQUAL":   &tender.Int{Value: int64(gl.GEQUAL)},
		"ALWAYS":   &tender.Int{Value: int64(gl.ALWAYS)},

		// ==================== Alpha Function ====================
		// "FUNC_NEVER":    &tender.Int{Value: int64(gl.FUNC_NEVER)},
		// "FUNC_LESS":     &tender.Int{Value: int64(gl.FUNC_LESS)},
		// "FUNC_EQUAL":    &tender.Int{Value: int64(gl.FUNC_EQUAL)},
		// "FUNC_LEQUAL":   &tender.Int{Value: int64(gl.FUNC_LEQUAL)},
		// "FUNC_GREATER":  &tender.Int{Value: int64(gl.FUNC_GREATER)},
		// "FUNC_NOTEQUAL": &tender.Int{Value: int64(gl.FUNC_NOTEQUAL)},
		// "FUNC_GEQUAL":   &tender.Int{Value: int64(gl.FUNC_GEQUAL)},
		// "FUNC_ALWAYS":   &tender.Int{Value: int64(gl.FUNC_ALWAYS)},

		// ==================== Hint Targets ====================
		"PERSPECTIVE_CORRECTION_HINT": &tender.Int{Value: int64(gl.PERSPECTIVE_CORRECTION_HINT)},
		"POINT_SMOOTH_HINT":           &tender.Int{Value: int64(gl.POINT_SMOOTH_HINT)},
		"LINE_SMOOTH_HINT":            &tender.Int{Value: int64(gl.LINE_SMOOTH_HINT)},
		"POLYGON_SMOOTH_HINT":         &tender.Int{Value: int64(gl.POLYGON_SMOOTH_HINT)},
		"FOG_HINT":                    &tender.Int{Value: int64(gl.FOG_HINT)},

		// ==================== Hint Modes ====================
		"DONT_CARE": &tender.Int{Value: int64(gl.DONT_CARE)},
		"FASTEST":   &tender.Int{Value: int64(gl.FASTEST)},
		"NICEST":    &tender.Int{Value: int64(gl.NICEST)},

		// ==================== Error Codes ====================
		"NO_ERROR":          &tender.Int{Value: int64(gl.NO_ERROR)},
		"INVALID_ENUM":      &tender.Int{Value: int64(gl.INVALID_ENUM)},
		"INVALID_VALUE":     &tender.Int{Value: int64(gl.INVALID_VALUE)},
		"INVALID_OPERATION": &tender.Int{Value: int64(gl.INVALID_OPERATION)},
		"STACK_OVERFLOW":    &tender.Int{Value: int64(gl.STACK_OVERFLOW)},
		"STACK_UNDERFLOW":   &tender.Int{Value: int64(gl.STACK_UNDERFLOW)},
		"OUT_OF_MEMORY":     &tender.Int{Value: int64(gl.OUT_OF_MEMORY)},
		

        // ==================== Initialization ====================
        "init": &tender.BuiltinFunction{
            Name: "init",
            Value: func(args ...tender.Object) (tender.Object, error) {
                if len(args) != 0 {
                    return nil, ErrInvalidArgCount
                }
                err := gl.Init()
                if err != nil {
                    return nil, err
                }
                return tender.NullValue, nil
            },
        },

		// ==================== Matrix Operations ====================
		"matrixMode": &tender.BuiltinFunction{
			Name: "matrixMode",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				mode, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				gl.MatrixMode(mode)
				return tender.NullValue, nil
			},
		},

		"loadIdentity": &tender.BuiltinFunction{
			Name: "loadIdentity",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				gl.LoadIdentity()
				return tender.NullValue, nil
			},
		},

		"pushMatrix": &tender.BuiltinFunction{
			Name: "pushMatrix",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				gl.PushMatrix()
				return tender.NullValue, nil
			},
		},

		"popMatrix": &tender.BuiltinFunction{
			Name: "popMatrix",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				gl.PopMatrix()
				return tender.NullValue, nil
			},
		},

		"multMatrixf": &tender.BuiltinFunction{
			Name: "multMatrixf",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 16 {
					return nil, ErrInvalidArgCount
				}
				m := make([]float32, 16)
				for i := 0; i < 16; i++ {
					val, ok := getFloat32(args[i])
					if !ok {
						return nil, ErrInvalidArgumentType
					}
					m[i] = val
				}
				gl.MultMatrixf(&m[0])
				return tender.NullValue, nil
			},
		},

		"loadMatrixf": &tender.BuiltinFunction{
			Name: "loadMatrixf",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 16 {
					return nil, ErrInvalidArgCount
				}
				m := make([]float32, 16)
				for i := 0; i < 16; i++ {
					val, ok := getFloat32(args[i])
					if !ok {
						return nil, ErrInvalidArgumentType
					}
					m[i] = val
				}
				gl.LoadMatrixf(&m[0])
				return tender.NullValue, nil
			},
		},

		// ==================== Transformation Functions ====================
		"translatef": &tender.BuiltinFunction{
			Name: "translatef",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 3 {
					return nil, ErrInvalidArgCount
				}
				x, okX := getFloat32(args[0])
				y, okY := getFloat32(args[1])
				z, okZ := getFloat32(args[2])
				if !okX || !okY || !okZ {
					return nil, ErrInvalidArgumentType
				}
				gl.Translatef(x, y, z)
				return tender.NullValue, nil
			},
		},

		"rotatef": &tender.BuiltinFunction{
			Name: "rotatef",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, ErrInvalidArgCount
				}
				angle, okA := getFloat32(args[0])
				x, okX := getFloat32(args[1])
				y, okY := getFloat32(args[2])
				z, okZ := getFloat32(args[3])
				if !okA || !okX || !okY || !okZ {
					return nil, ErrInvalidArgumentType
				}
				gl.Rotatef(angle, x, y, z)
				return tender.NullValue, nil
			},
		},

		"scalef": &tender.BuiltinFunction{
			Name: "scalef",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 3 {
					return nil, ErrInvalidArgCount
				}
				x, okX := getFloat32(args[0])
				y, okY := getFloat32(args[1])
				z, okZ := getFloat32(args[2])
				if !okX || !okY || !okZ {
					return nil, ErrInvalidArgumentType
				}
				gl.Scalef(x, y, z)
				return tender.NullValue, nil
			},
		},

		// ==================== Viewport and Clipping ====================
		"viewport": &tender.BuiltinFunction{
			Name: "viewport",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, ErrInvalidArgCount
				}
				x, okX := tender.ToInt32(args[0])
				y, okY := tender.ToInt32(args[1])
				w, okW := tender.ToInt32(args[2])
				h, okH := tender.ToInt32(args[3])
				if !okX || !okY || !okW || !okH {
					return nil, ErrInvalidArgumentType
				}
				gl.Viewport(x, y, w, h)
				return tender.NullValue, nil
			},
		},

		"scissor": &tender.BuiltinFunction{
			Name: "scissor",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, ErrInvalidArgCount
				}
				x, okX := tender.ToInt32(args[0])
				y, okY := tender.ToInt32(args[1])
				w, okW := tender.ToInt32(args[2])
				h, okH := tender.ToInt32(args[3])
				if !okX || !okY || !okW || !okH {
					return nil, ErrInvalidArgumentType
				}
				gl.Scissor(x, y, w, h)
				return tender.NullValue, nil
			},
		},

		// ==================== Clearing ====================
		"clear": &tender.BuiltinFunction{
			Name: "clear",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				mask, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				gl.Clear(mask)
				return tender.NullValue, nil
			},
		},

		"clearColor": &tender.BuiltinFunction{
			Name: "clearColor",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, ErrInvalidArgCount
				}
				r, okR := getFloat32(args[0])
				g, okG := getFloat32(args[1])
				b, okB := getFloat32(args[2])
				a, okA := getFloat32(args[3])
				if !okR || !okG || !okB || !okA {
					return nil, ErrInvalidArgumentType
				}
				gl.ClearColor(r, g, b, a)
				return tender.NullValue, nil
			},
		},

		"clearDepth": &tender.BuiltinFunction{
			Name: "clearDepth",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				depth, ok := getFloat64(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				gl.ClearDepth(float64(depth))
				return tender.NullValue, nil
			},
		},

		"clearStencil": &tender.BuiltinFunction{
			Name: "clearStencil",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				s, ok := tender.ToInt32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				gl.ClearStencil(s)
				return tender.NullValue, nil
			},
		},

		"clearAccum": &tender.BuiltinFunction{
			Name: "clearAccum",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, ErrInvalidArgCount
				}
				r, okR := getFloat32(args[0])
				g, okG := getFloat32(args[1])
				b, okB := getFloat32(args[2])
				a, okA := getFloat32(args[3])
				if !okR || !okG || !okB || !okA {
					return nil, ErrInvalidArgumentType
				}
				gl.ClearAccum(r, g, b, a)
				return tender.NullValue, nil
			},
		},

		// ==================== Primitives ====================
		"begin": &tender.BuiltinFunction{
			Name: "begin",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				mode, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				gl.Begin(mode)
				return tender.NullValue, nil
			},
		},

		"end": &tender.BuiltinFunction{
			Name: "end",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				gl.End()
				return tender.NullValue, nil
			},
		},

		"vertex2f": &tender.BuiltinFunction{
			Name: "vertex2f",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				x, okX := getFloat32(args[0])
				y, okY := getFloat32(args[1])
				if !okX || !okY {
					return nil, ErrInvalidArgumentType
				}
				gl.Vertex2f(x, y)
				return tender.NullValue, nil
			},
		},

		"vertex2d": &tender.BuiltinFunction{
			Name: "vertex2d",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				x, okX := getFloat64(args[0])
				y, okY := getFloat64(args[1])
				if !okX || !okY {
					return nil, ErrInvalidArgumentType
				}
				gl.Vertex2d(x, y)
				return tender.NullValue, nil
			},
		},

		"vertex2i": &tender.BuiltinFunction{
			Name: "vertex2i",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				x, okX := tender.ToInt32(args[0])
				y, okY := tender.ToInt32(args[1])
				if !okX || !okY {
					return nil, ErrInvalidArgumentType
				}
				gl.Vertex2i(x, y)
				return tender.NullValue, nil
			},
		},

		"vertex3f": &tender.BuiltinFunction{
			Name: "vertex3f",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 3 {
					return nil, ErrInvalidArgCount
				}
				x, okX := getFloat32(args[0])
				y, okY := getFloat32(args[1])
				z, okZ := getFloat32(args[2])
				if !okX || !okY || !okZ {
					return nil, ErrInvalidArgumentType
				}
				gl.Vertex3f(x, y, z)
				return tender.NullValue, nil
			},
		},

		"vertex3d": &tender.BuiltinFunction{
			Name: "vertex3d",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 3 {
					return nil, ErrInvalidArgCount
				}
				x, okX := getFloat64(args[0])
				y, okY := getFloat64(args[1])
				z, okZ := getFloat64(args[2])
				if !okX || !okY || !okZ {
					return nil, ErrInvalidArgumentType
				}
				gl.Vertex3d(x, y, z)
				return tender.NullValue, nil
			},
		},

		"vertex3i": &tender.BuiltinFunction{
			Name: "vertex3i",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 3 {
					return nil, ErrInvalidArgCount
				}
				x, okX := tender.ToInt32(args[0])
				y, okY := tender.ToInt32(args[1])
				z, okZ := tender.ToInt32(args[2])
				if !okX || !okY || !okZ {
					return nil, ErrInvalidArgumentType
				}
				gl.Vertex3i(x, y, z)
				return tender.NullValue, nil
			},
		},

		"vertex4f": &tender.BuiltinFunction{
			Name: "vertex4f",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, ErrInvalidArgCount
				}
				x, okX := getFloat32(args[0])
				y, okY := getFloat32(args[1])
				z, okZ := getFloat32(args[2])
				w, okW := getFloat32(args[3])
				if !okX || !okY || !okZ || !okW {
					return nil, ErrInvalidArgumentType
				}
				gl.Vertex4f(x, y, z, w)
				return tender.NullValue, nil
			},
		},

		"vertex4d": &tender.BuiltinFunction{
			Name: "vertex4d",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, ErrInvalidArgCount
				}
				x, okX := getFloat64(args[0])
				y, okY := getFloat64(args[1])
				z, okZ := getFloat64(args[2])
				w, okW := getFloat64(args[3])
				if !okX || !okY || !okZ || !okW {
					return nil, ErrInvalidArgumentType
				}
				gl.Vertex4d(x, y, z, w)
				return tender.NullValue, nil
			},
		},

		// ==================== Colors ====================
		"color3f": &tender.BuiltinFunction{
			Name: "color3f",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 3 {
					return nil, ErrInvalidArgCount
				}
				r, okR := getFloat32(args[0])
				g, okG := getFloat32(args[1])
				b, okB := getFloat32(args[2])
				if !okR || !okG || !okB {
					return nil, ErrInvalidArgumentType
				}
				gl.Color3f(r, g, b)
				return tender.NullValue, nil
			},
		},

		"color3d": &tender.BuiltinFunction{
			Name: "color3d",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 3 {
					return nil, ErrInvalidArgCount
				}
				r, okR := getFloat64(args[0])
				g, okG := getFloat64(args[1])
				b, okB := getFloat64(args[2])
				if !okR || !okG || !okB {
					return nil, ErrInvalidArgumentType
				}
				gl.Color3d(r, g, b)
				return tender.NullValue, nil
			},
		},

		"color3ub": &tender.BuiltinFunction{
			Name: "color3ub",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 3 {
					return nil, ErrInvalidArgCount
				}
				r, okR := getInt(args[0])
				g, okG := getInt(args[1])
				b, okB := getInt(args[2])
				if !okR || !okG || !okB {
					return nil, ErrInvalidArgumentType
				}
				gl.Color3ub(uint8(r), uint8(g), uint8(b))
				return tender.NullValue, nil
			},
		},

		"color4f": &tender.BuiltinFunction{
			Name: "color4f",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, ErrInvalidArgCount
				}
				r, okR := getFloat32(args[0])
				g, okG := getFloat32(args[1])
				b, okB := getFloat32(args[2])
				a, okA := getFloat32(args[3])
				if !okR || !okG || !okB || !okA {
					return nil, ErrInvalidArgumentType
				}
				gl.Color4f(r, g, b, a)
				return tender.NullValue, nil
			},
		},

		"color4d": &tender.BuiltinFunction{
			Name: "color4d",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, ErrInvalidArgCount
				}
				r, okR := getFloat64(args[0])
				g, okG := getFloat64(args[1])
				b, okB := getFloat64(args[2])
				a, okA := getFloat64(args[3])
				if !okR || !okG || !okB || !okA {
					return nil, ErrInvalidArgumentType
				}
				gl.Color4d(r, g, b, a)
				return tender.NullValue, nil
			},
		},

		"color4ub": &tender.BuiltinFunction{
			Name: "color4ub",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, ErrInvalidArgCount
				}
				r, okR := getInt(args[0])
				g, okG := getInt(args[1])
				b, okB := getInt(args[2])
				a, okA := getInt(args[3])
				if !okR || !okG || !okB || !okA {
					return nil, ErrInvalidArgumentType
				}
				gl.Color4ub(uint8(r), uint8(g), uint8(b), uint8(a))
				return tender.NullValue, nil
			},
		},

		// ==================== Normals ====================
		"normal3f": &tender.BuiltinFunction{
			Name: "normal3f",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 3 {
					return nil, ErrInvalidArgCount
				}
				nx, okX := getFloat32(args[0])
				ny, okY := getFloat32(args[1])
				nz, okZ := getFloat32(args[2])
				if !okX || !okY || !okZ {
					return nil, ErrInvalidArgumentType
				}
				gl.Normal3f(nx, ny, nz)
				return tender.NullValue, nil
			},
		},

		"normal3d": &tender.BuiltinFunction{
			Name: "normal3d",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 3 {
					return nil, ErrInvalidArgCount
				}
				nx, okX := getFloat64(args[0])
				ny, okY := getFloat64(args[1])
				nz, okZ := getFloat64(args[2])
				if !okX || !okY || !okZ {
					return nil, ErrInvalidArgumentType
				}
				gl.Normal3d(nx, ny, nz)
				return tender.NullValue, nil
			},
		},

		// ==================== Texture Coordinates ====================
		"texCoord2f": &tender.BuiltinFunction{
			Name: "texCoord2f",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				s, okS := getFloat32(args[0])
				t, okT := getFloat32(args[1])
				if !okS || !okT {
					return nil, ErrInvalidArgumentType
				}
				gl.TexCoord2f(s, t)
				return tender.NullValue, nil
			},
		},

		"texCoord2d": &tender.BuiltinFunction{
			Name: "texCoord2d",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				s, okS := getFloat64(args[0])
				t, okT := getFloat64(args[1])
				if !okS || !okT {
					return nil, ErrInvalidArgumentType
				}
				gl.TexCoord2d(s, t)
				return tender.NullValue, nil
			},
		},

		// ==================== Enabling/Disabling ====================
		"enable": &tender.BuiltinFunction{
			Name: "enable",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				cap, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				gl.Enable(cap)
				return tender.NullValue, nil
			},
		},

		"disable": &tender.BuiltinFunction{
			Name: "disable",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				cap, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				gl.Disable(cap)
				return tender.NullValue, nil
			},
		},

		"isEnabled": &tender.BuiltinFunction{
			Name: "isEnabled",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				cap, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				if gl.IsEnabled(cap) {
					return tender.TrueValue, nil
				}
				return tender.FalseValue, nil
			},
		},

		// ==================== Blend Functions ====================
		"blendFunc": &tender.BuiltinFunction{
			Name: "blendFunc",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				sfactor, okS := getUint32(args[0])
				dfactor, okD := getUint32(args[1])
				if !okS || !okD {
					return nil, ErrInvalidArgumentType
				}
				gl.BlendFunc(sfactor, dfactor)
				return tender.NullValue, nil
			},
		},

		// ==================== Depth Functions ====================
		"depthFunc": &tender.BuiltinFunction{
			Name: "depthFunc",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				fn, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				gl.DepthFunc(fn)
				return tender.NullValue, nil
			},
		},

		"depthMask": &tender.BuiltinFunction{
			Name: "depthMask",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				flag := args[0] == tender.TrueValue
				if i, ok := getInt64(args[0]); ok && i != 0 {
					flag = true
				}
				gl.DepthMask(flag)
				return tender.NullValue, nil
			},
		},

		"depthRange": &tender.BuiltinFunction{
			Name: "depthRange",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				near, okN := getFloat64(args[0])
				far, okF := getFloat64(args[1])
				if !okN || !okF {
					return nil, ErrInvalidArgumentType
				}
				gl.DepthRange(near, far)
				return tender.NullValue, nil
			},
		},

		// ==================== Alpha Function ====================
		"alphaFunc": &tender.BuiltinFunction{
			Name: "alphaFunc",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				fn, okF := getUint32(args[0])
				ref, okR := getFloat32(args[1])
				if !okF || !okR {
					return nil, ErrInvalidArgumentType
				}
				gl.AlphaFunc(fn, ref)
				return tender.NullValue, nil
			},
		},

		// ==================== Culling ====================
		"cullFace": &tender.BuiltinFunction{
			Name: "cullFace",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				mode, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				gl.CullFace(mode)
				return tender.NullValue, nil
			},
		},

		"frontFace": &tender.BuiltinFunction{
			Name: "frontFace",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				mode, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				gl.FrontFace(mode)
				return tender.NullValue, nil
			},
		},

		// ==================== Polygon Modes ====================
		"polygonMode": &tender.BuiltinFunction{
			Name: "polygonMode",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				face, okF := getUint32(args[0])
				mode, okM := getUint32(args[1])
				if !okF || !okM {
					return nil, ErrInvalidArgumentType
				}
				gl.PolygonMode(face, mode)
				return tender.NullValue, nil
			},
		},

		"polygonOffset": &tender.BuiltinFunction{
			Name: "polygonOffset",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				factor, okF := getFloat32(args[0])
				units, okU := getFloat32(args[1])
				if !okF || !okU {
					return nil, ErrInvalidArgumentType
				}
				gl.PolygonOffset(factor, units)
				return tender.NullValue, nil
			},
		},

		// ==================== Shading ====================
		"shadeModel": &tender.BuiltinFunction{
			Name: "shadeModel",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				mode, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				gl.ShadeModel(mode)
				return tender.NullValue, nil
			},
		},

		// ==================== Hints ====================
		"hint": &tender.BuiltinFunction{
			Name: "hint",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				target, okT := getUint32(args[0])
				mode, okM := getUint32(args[1])
				if !okT || !okM {
					return nil, ErrInvalidArgumentType
				}
				gl.Hint(target, mode)
				return tender.NullValue, nil
			},
		},

		// ==================== Flush and Finish ====================
		"flush": &tender.BuiltinFunction{
			Name: "flush",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				gl.Flush()
				return tender.NullValue, nil
			},
		},

		"finish": &tender.BuiltinFunction{
			Name: "finish",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				gl.Finish()
				return tender.NullValue, nil
			},
		},

		// ==================== Error Handling ====================
		"getError": &tender.BuiltinFunction{
			Name: "getError",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				err := gl.GetError()
				return &tender.Int{Value: int64(err)}, nil
			},
		},

		// ==================== Get Functions ====================
		"getBooleanv": &tender.BuiltinFunction{
			Name: "getBooleanv",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				pname, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				var val bool
				gl.GetBooleanv(pname, &val)
				if val {
					return tender.TrueValue, nil
				}
				return tender.FalseValue, nil
			},
		},

		"getIntegerv": &tender.BuiltinFunction{
			Name: "getIntegerv",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				pname, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				var val int32
				gl.GetIntegerv(pname, &val)
				return &tender.Int{Value: int64(val)}, nil
			},
		},

		"getFloatv": &tender.BuiltinFunction{
			Name: "getFloatv",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				pname, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				var val float32
				gl.GetFloatv(pname, &val)
				return &tender.Float{Value: float64(val)}, nil
			},
		},

		"getDoublev": &tender.BuiltinFunction{
			Name: "getDoublev",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				pname, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				var val float64
				gl.GetDoublev(pname, &val)
				return &tender.Float{Value: val}, nil
			},
		},
		// ==================== Accumulation Buffer ====================
		"accum": &tender.BuiltinFunction{
			Name: "accum",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				op, okO := getUint32(args[0])
				value, okV := getFloat32(args[1])
				if !okO || !okV {
					return nil, ErrInvalidArgumentType
				}
				gl.Accum(op, value)
				return tender.NullValue, nil
			},
		},

		// ==================== Render Mode ====================
		"renderMode": &tender.BuiltinFunction{
			Name: "renderMode",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				mode, ok := getUint32(args[0])
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				result := gl.RenderMode(mode)
				return &tender.Int{Value: int64(result)}, nil
			},
		},
	},
}