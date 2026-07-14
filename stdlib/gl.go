//go:build gl

package stdlib

import (
	"os"
	"image"
	"image/color"
	"unsafe"
	"path/filepath"

	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/v/gl"
)

// Module exposes the OpenGL API to Tender scripts
var glModule = map[string]tender.Object{
	// ==================== Matrix Mode Constants ====================
	"MODELVIEW":  &tender.Int{Value: int64(gl.MODELVIEW)},
	"PROJECTION": &tender.Int{Value: int64(gl.PROJECTION)},

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

	// ==================== Enable/Disable Capabilities ====================
	"BLEND":          &tender.Int{Value: int64(gl.BLEND)},
	"DEPTH_TEST":     &tender.Int{Value: int64(gl.DEPTH_TEST)},
	"CULL_FACE":      &tender.Int{Value: int64(gl.CULL_FACE)},
	"LIGHTING":       &tender.Int{Value: int64(gl.LIGHTING)},
	"LIGHT0":         &tender.Int{Value: int64(gl.LIGHT0)},
	"LIGHT1":         &tender.Int{Value: int64(gl.LIGHT1)},
	"LIGHT2":         &tender.Int{Value: int64(gl.LIGHT2)},
	"LIGHT3":         &tender.Int{Value: int64(gl.LIGHT3)},
	"LIGHT4":         &tender.Int{Value: int64(gl.LIGHT4)},
	"LIGHT5":         &tender.Int{Value: int64(gl.LIGHT5)},
	"LIGHT6":         &tender.Int{Value: int64(gl.LIGHT6)},
	"LIGHT7":         &tender.Int{Value: int64(gl.LIGHT7)},
	"FOG":            &tender.Int{Value: int64(gl.FOG)},
	"SCISSOR_TEST":   &tender.Int{Value: int64(gl.SCISSOR_TEST)},
	"STENCIL_TEST":   &tender.Int{Value: int64(gl.STENCIL_TEST)},
	"ALPHA_TEST":     &tender.Int{Value: int64(gl.ALPHA_TEST)},
	"NORMALIZE":      &tender.Int{Value: int64(gl.NORMALIZE)},
	"COLOR_MATERIAL": &tender.Int{Value: int64(gl.COLOR_MATERIAL)},

	// ==================== Blend Factors ====================
	"ZERO":                     &tender.Int{Value: int64(gl.ZERO)},
	"ONE":                      &tender.Int{Value: int64(gl.ONE)},
	"SRC_COLOR":                &tender.Int{Value: int64(gl.SRC_COLOR)},
	"ONE_MINUS_SRC_COLOR":      &tender.Int{Value: int64(gl.ONE_MINUS_SRC_COLOR)},
	"DST_COLOR":                &tender.Int{Value: int64(gl.DST_COLOR)},
	"ONE_MINUS_DST_COLOR":      &tender.Int{Value: int64(gl.ONE_MINUS_DST_COLOR)},
	"SRC_ALPHA":                &tender.Int{Value: int64(gl.SRC_ALPHA)},
	"ONE_MINUS_SRC_ALPHA":      &tender.Int{Value: int64(gl.ONE_MINUS_SRC_ALPHA)},
	"DST_ALPHA":                &tender.Int{Value: int64(gl.DST_ALPHA)},
	"ONE_MINUS_DST_ALPHA":      &tender.Int{Value: int64(gl.ONE_MINUS_DST_ALPHA)},
	"SRC_ALPHA_SATURATE":       &tender.Int{Value: int64(gl.SRC_ALPHA_SATURATE)},
	"CONSTANT_COLOR":           &tender.Int{Value: int64(gl.CONSTANT_COLOR)},
	"ONE_MINUS_CONSTANT_COLOR": &tender.Int{Value: int64(gl.ONE_MINUS_CONSTANT_COLOR)},
	"CONSTANT_ALPHA":           &tender.Int{Value: int64(gl.CONSTANT_ALPHA)},
	"ONE_MINUS_CONSTANT_ALPHA": &tender.Int{Value: int64(gl.ONE_MINUS_CONSTANT_ALPHA)},

	// ==================== Blend Equations ====================
	"FUNC_ADD":              &tender.Int{Value: int64(gl.FUNC_ADD)},
	"FUNC_SUBTRACT":         &tender.Int{Value: int64(gl.FUNC_SUBTRACT)},
	"FUNC_REVERSE_SUBTRACT": &tender.Int{Value: int64(gl.FUNC_REVERSE_SUBTRACT)},
	"MIN":                   &tender.Int{Value: int64(gl.MIN)},
	"MAX":                   &tender.Int{Value: int64(gl.MAX)},

	// ==================== Depth / Alpha Functions ====================
	"NEVER":    &tender.Int{Value: int64(gl.NEVER)},
	"LESS":     &tender.Int{Value: int64(gl.LESS)},
	"EQUAL":    &tender.Int{Value: int64(gl.EQUAL)},
	"LEQUAL":   &tender.Int{Value: int64(gl.LEQUAL)},
	"GREATER":  &tender.Int{Value: int64(gl.GREATER)},
	"NOTEQUAL": &tender.Int{Value: int64(gl.NOTEQUAL)},
	"GEQUAL":   &tender.Int{Value: int64(gl.GEQUAL)},
	"ALWAYS":   &tender.Int{Value: int64(gl.ALWAYS)},

	// ==================== Cull Face Modes ====================
	"FRONT":          &tender.Int{Value: int64(gl.FRONT)},
	"BACK":           &tender.Int{Value: int64(gl.BACK)},
	"FRONT_AND_BACK": &tender.Int{Value: int64(gl.FRONT_AND_BACK)},

	// ==================== Front Face ====================
	"CW":  &tender.Int{Value: int64(gl.CW)},
	"CCW": &tender.Int{Value: int64(gl.CCW)},

	// ==================== Polygon Modes ====================
	"POINT": &tender.Int{Value: int64(gl.POINT)},
	"LINE":  &tender.Int{Value: int64(gl.LINE)},
	"FILL":  &tender.Int{Value: int64(gl.FILL)},

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

	// ==================== GetString Names ====================
	"VENDOR":     &tender.Int{Value: int64(gl.VENDOR)},
	"RENDERER":   &tender.Int{Value: int64(gl.RENDERER)},
	"VERSION":    &tender.Int{Value: int64(gl.VERSION)},
	"EXTENSIONS": &tender.Int{Value: int64(gl.EXTENSIONS)},

	// ==================== Stencil Operations ====================
	"KEEP":        &tender.Int{Value: int64(gl.KEEP)},
	"REPLACE":     &tender.Int{Value: int64(gl.REPLACE)},
	"INCR":        &tender.Int{Value: int64(gl.INCR)},
	"DECR":        &tender.Int{Value: int64(gl.DECR)},
	"INVERT":      &tender.Int{Value: int64(gl.INVERT)},
	"INCR_WRAP":   &tender.Int{Value: int64(gl.INCR_WRAP)},
	"DECR_WRAP":   &tender.Int{Value: int64(gl.DECR_WRAP)},

	// ==================== Light / Material Parameters ====================
	"AMBIENT":                 &tender.Int{Value: int64(gl.AMBIENT)},
	"DIFFUSE":                 &tender.Int{Value: int64(gl.DIFFUSE)},
	"SPECULAR":                &tender.Int{Value: int64(gl.SPECULAR)},
	"POSITION":                &tender.Int{Value: int64(gl.POSITION)},
	"SPOT_DIRECTION":          &tender.Int{Value: int64(gl.SPOT_DIRECTION)},
	"SPOT_EXPONENT":           &tender.Int{Value: int64(gl.SPOT_EXPONENT)},
	"SPOT_CUTOFF":             &tender.Int{Value: int64(gl.SPOT_CUTOFF)},
	"CONSTANT_ATTENUATION":    &tender.Int{Value: int64(gl.CONSTANT_ATTENUATION)},
	"LINEAR_ATTENUATION":      &tender.Int{Value: int64(gl.LINEAR_ATTENUATION)},
	"QUADRATIC_ATTENUATION":   &tender.Int{Value: int64(gl.QUADRATIC_ATTENUATION)},
	"SHININESS":               &tender.Int{Value: int64(gl.SHININESS)},
	"EMISSION":                &tender.Int{Value: int64(gl.EMISSION)},
	"LIGHT_MODEL_AMBIENT":     &tender.Int{Value: int64(gl.LIGHT_MODEL_AMBIENT)},
	"LIGHT_MODEL_LOCAL_VIEWER": &tender.Int{Value: int64(gl.LIGHT_MODEL_LOCAL_VIEWER)},
	"LIGHT_MODEL_TWO_SIDE":    &tender.Int{Value: int64(gl.LIGHT_MODEL_TWO_SIDE)},
	"SINGLE_COLOR":            &tender.Int{Value: int64(gl.SINGLE_COLOR)},
	"SEPARATE_SPECULAR_COLOR": &tender.Int{Value: int64(gl.SEPARATE_SPECULAR_COLOR)},

	// ==================== Fog Parameters ====================
	"FOG_COLOR":   &tender.Int{Value: int64(gl.FOG_COLOR)},
	"FOG_DENSITY": &tender.Int{Value: int64(gl.FOG_DENSITY)},
	"FOG_START":   &tender.Int{Value: int64(gl.FOG_START)},
	"FOG_END":     &tender.Int{Value: int64(gl.FOG_END)},
	"FOG_MODE":    &tender.Int{Value: int64(gl.FOG_MODE)},

	// ==================== Fog Modes ====================
	"EXP":   &tender.Int{Value: int64(gl.EXP)},
	"EXP2":  &tender.Int{Value: int64(gl.EXP2)},

	// ==================== Texture Targets ====================
	"TEXTURE_1D":                &tender.Int{Value: int64(gl.TEXTURE_1D)},
	"TEXTURE_2D":                &tender.Int{Value: int64(gl.TEXTURE_2D)},
	"TEXTURE_3D":                &tender.Int{Value: int64(gl.TEXTURE_3D)},
	"TEXTURE_CUBE_MAP":          &tender.Int{Value: int64(gl.TEXTURE_CUBE_MAP)},
	"TEXTURE_CUBE_MAP_POSITIVE_X": &tender.Int{Value: int64(gl.TEXTURE_CUBE_MAP_POSITIVE_X)},
	"TEXTURE_CUBE_MAP_NEGATIVE_X": &tender.Int{Value: int64(gl.TEXTURE_CUBE_MAP_NEGATIVE_X)},
	"TEXTURE_CUBE_MAP_POSITIVE_Y": &tender.Int{Value: int64(gl.TEXTURE_CUBE_MAP_POSITIVE_Y)},
	"TEXTURE_CUBE_MAP_NEGATIVE_Y": &tender.Int{Value: int64(gl.TEXTURE_CUBE_MAP_NEGATIVE_Y)},
	"TEXTURE_CUBE_MAP_POSITIVE_Z": &tender.Int{Value: int64(gl.TEXTURE_CUBE_MAP_POSITIVE_Z)},
	"TEXTURE_CUBE_MAP_NEGATIVE_Z": &tender.Int{Value: int64(gl.TEXTURE_CUBE_MAP_NEGATIVE_Z)},
	
	"TEXTURE0": &tender.Int{Value: int64(gl.TEXTURE0)},
	"TEXTURE1": &tender.Int{Value: int64(gl.TEXTURE1)},
	"TEXTURE2": &tender.Int{Value: int64(gl.TEXTURE2)},
	"TEXTURE3": &tender.Int{Value: int64(gl.TEXTURE3)},
	"TEXTURE4": &tender.Int{Value: int64(gl.TEXTURE4)},
	"TEXTURE5": &tender.Int{Value: int64(gl.TEXTURE5)},
	"TEXTURE6": &tender.Int{Value: int64(gl.TEXTURE6)},
	"TEXTURE7": &tender.Int{Value: int64(gl.TEXTURE7)},

	// ==================== Texture Parameters ====================
	"TEXTURE_MIN_FILTER": &tender.Int{Value: int64(gl.TEXTURE_MIN_FILTER)},
	"TEXTURE_MAG_FILTER": &tender.Int{Value: int64(gl.TEXTURE_MAG_FILTER)},
	"TEXTURE_WRAP_S":     &tender.Int{Value: int64(gl.TEXTURE_WRAP_S)},
	"TEXTURE_WRAP_T":     &tender.Int{Value: int64(gl.TEXTURE_WRAP_T)},
	"TEXTURE_BORDER_COLOR": &tender.Int{Value: int64(gl.TEXTURE_BORDER_COLOR)},
	"GENERATE_MIPMAP":    &tender.Int{Value: int64(gl.GENERATE_MIPMAP)},
	"TEXTURE_LOD_BIAS":   &tender.Int{Value: int64(gl.TEXTURE_LOD_BIAS)},

	// ==================== Texture Min/Mag Filters ====================
	"NEAREST":               &tender.Int{Value: int64(gl.NEAREST)},
	"LINEAR":                &tender.Int{Value: int64(gl.LINEAR)},
	"NEAREST_MIPMAP_NEAREST": &tender.Int{Value: int64(gl.NEAREST_MIPMAP_NEAREST)},
	"LINEAR_MIPMAP_NEAREST":  &tender.Int{Value: int64(gl.LINEAR_MIPMAP_NEAREST)},
	"NEAREST_MIPMAP_LINEAR":  &tender.Int{Value: int64(gl.NEAREST_MIPMAP_LINEAR)},
	"LINEAR_MIPMAP_LINEAR":   &tender.Int{Value: int64(gl.LINEAR_MIPMAP_LINEAR)},

	// ==================== Texture Wrap Modes ====================
	"CLAMP":           &tender.Int{Value: int64(gl.CLAMP)},
	"REPEAT":          &tender.Int{Value: int64(gl.REPEAT)},
	"CLAMP_TO_EDGE":   &tender.Int{Value: int64(gl.CLAMP_TO_EDGE)},
	"CLAMP_TO_BORDER": &tender.Int{Value: int64(gl.CLAMP_TO_BORDER)},
	"MIRRORED_REPEAT": &tender.Int{Value: int64(gl.MIRRORED_REPEAT)},

	// ==================== Texture Environment ====================
	"TEXTURE_ENV":        &tender.Int{Value: int64(gl.TEXTURE_ENV)},
	"TEXTURE_ENV_MODE":   &tender.Int{Value: int64(gl.TEXTURE_ENV_MODE)},
	"TEXTURE_ENV_COLOR":  &tender.Int{Value: int64(gl.TEXTURE_ENV_COLOR)},
	"MODULATE":           &tender.Int{Value: int64(gl.MODULATE)},
	"DECAL":              &tender.Int{Value: int64(gl.DECAL)},
	"COMBINE":            &tender.Int{Value: int64(gl.COMBINE)},
	"COMBINE_RGB":        &tender.Int{Value: int64(gl.COMBINE_RGB)},
	"COMBINE_ALPHA":      &tender.Int{Value: int64(gl.COMBINE_ALPHA)},
	"SOURCE0_RGB":        &tender.Int{Value: int64(gl.SOURCE0_RGB)},
	"SOURCE1_RGB":        &tender.Int{Value: int64(gl.SOURCE1_RGB)},
	"SOURCE2_RGB":        &tender.Int{Value: int64(gl.SOURCE2_RGB)},
	"SOURCE0_ALPHA":      &tender.Int{Value: int64(gl.SOURCE0_ALPHA)},
	"SOURCE1_ALPHA":      &tender.Int{Value: int64(gl.SOURCE1_ALPHA)},
	"SOURCE2_ALPHA":      &tender.Int{Value: int64(gl.SOURCE2_ALPHA)},
	"OPERAND0_RGB":       &tender.Int{Value: int64(gl.OPERAND0_RGB)},
	"OPERAND1_RGB":       &tender.Int{Value: int64(gl.OPERAND1_RGB)},
	"OPERAND2_RGB":       &tender.Int{Value: int64(gl.OPERAND2_RGB)},
	"OPERAND0_ALPHA":     &tender.Int{Value: int64(gl.OPERAND0_ALPHA)},
	"OPERAND1_ALPHA":     &tender.Int{Value: int64(gl.OPERAND1_ALPHA)},
	"OPERAND2_ALPHA":     &tender.Int{Value: int64(gl.OPERAND2_ALPHA)},
	"PREVIOUS":           &tender.Int{Value: int64(gl.PREVIOUS)},
	"CONSTANT":           &tender.Int{Value: int64(gl.CONSTANT)},
	"PRIMARY_COLOR":      &tender.Int{Value: int64(gl.PRIMARY_COLOR)},
	"TEXTURE":            &tender.Int{Value: int64(gl.TEXTURE)},

	// ==================== Internal Formats (common) ====================
	"ALPHA":               &tender.Int{Value: int64(gl.ALPHA)},
	"ALPHA4":              &tender.Int{Value: int64(gl.ALPHA4)},
	"ALPHA8":              &tender.Int{Value: int64(gl.ALPHA8)},
	"ALPHA12":             &tender.Int{Value: int64(gl.ALPHA12)},
	"ALPHA16":             &tender.Int{Value: int64(gl.ALPHA16)},
	"LUMINANCE":           &tender.Int{Value: int64(gl.LUMINANCE)},
	"LUMINANCE4":          &tender.Int{Value: int64(gl.LUMINANCE4)},
	"LUMINANCE8":          &tender.Int{Value: int64(gl.LUMINANCE8)},
	"LUMINANCE12":         &tender.Int{Value: int64(gl.LUMINANCE12)},
	"LUMINANCE16":         &tender.Int{Value: int64(gl.LUMINANCE16)},
	"LUMINANCE_ALPHA":     &tender.Int{Value: int64(gl.LUMINANCE_ALPHA)},
	"LUMINANCE4_ALPHA4":   &tender.Int{Value: int64(gl.LUMINANCE4_ALPHA4)},
	"LUMINANCE6_ALPHA2":   &tender.Int{Value: int64(gl.LUMINANCE6_ALPHA2)},
	"LUMINANCE8_ALPHA8":   &tender.Int{Value: int64(gl.LUMINANCE8_ALPHA8)},
	"LUMINANCE12_ALPHA4":  &tender.Int{Value: int64(gl.LUMINANCE12_ALPHA4)},
	"LUMINANCE12_ALPHA12": &tender.Int{Value: int64(gl.LUMINANCE12_ALPHA12)},
	"LUMINANCE16_ALPHA16": &tender.Int{Value: int64(gl.LUMINANCE16_ALPHA16)},
	"INTENSITY":           &tender.Int{Value: int64(gl.INTENSITY)},
	"INTENSITY4":          &tender.Int{Value: int64(gl.INTENSITY4)},
	"INTENSITY8":          &tender.Int{Value: int64(gl.INTENSITY8)},
	"INTENSITY12":         &tender.Int{Value: int64(gl.INTENSITY12)},
	"INTENSITY16":         &tender.Int{Value: int64(gl.INTENSITY16)},
	"RGB":                 &tender.Int{Value: int64(gl.RGB)},
	"RGB4":                &tender.Int{Value: int64(gl.RGB4)},
	"RGB8":                &tender.Int{Value: int64(gl.RGB8)},
	"RGB10":               &tender.Int{Value: int64(gl.RGB10)},
	"RGB12":               &tender.Int{Value: int64(gl.RGB12)},
	"RGB16":               &tender.Int{Value: int64(gl.RGB16)},
	"RGBA":                &tender.Int{Value: int64(gl.RGBA)},
	"RGBA2":               &tender.Int{Value: int64(gl.RGBA2)},
	"RGBA4":               &tender.Int{Value: int64(gl.RGBA4)},
	"RGB5_A1":             &tender.Int{Value: int64(gl.RGB5_A1)},
	"RGBA8":               &tender.Int{Value: int64(gl.RGBA8)},
	"RGB10_A2":            &tender.Int{Value: int64(gl.RGB10_A2)},
	"RGBA12":              &tender.Int{Value: int64(gl.RGBA12)},
	"RGBA16":              &tender.Int{Value: int64(gl.RGBA16)},
	"DEPTH_COMPONENT":     &tender.Int{Value: int64(gl.DEPTH_COMPONENT)},
	"DEPTH_COMPONENT16":   &tender.Int{Value: int64(gl.DEPTH_COMPONENT16)},
	"DEPTH_COMPONENT24":   &tender.Int{Value: int64(gl.DEPTH_COMPONENT24)},
	"DEPTH_COMPONENT32":   &tender.Int{Value: int64(gl.DEPTH_COMPONENT32)},
	"DEPTH_STENCIL":       &tender.Int{Value: int64(gl.DEPTH_STENCIL)},
	"DEPTH24_STENCIL8":    &tender.Int{Value: int64(gl.DEPTH24_STENCIL8)},
	"STENCIL_INDEX":       &tender.Int{Value: int64(gl.STENCIL_INDEX)},
	"STENCIL_INDEX1":      &tender.Int{Value: int64(gl.STENCIL_INDEX1)},
	"STENCIL_INDEX4":      &tender.Int{Value: int64(gl.STENCIL_INDEX4)},
	"STENCIL_INDEX8":      &tender.Int{Value: int64(gl.STENCIL_INDEX8)},
	"STENCIL_INDEX16":     &tender.Int{Value: int64(gl.STENCIL_INDEX16)},

	// ==================== Pixel Formats & Types ====================
	"RED":   &tender.Int{Value: int64(gl.RED)},
	"GREEN": &tender.Int{Value: int64(gl.GREEN)},
	"BLUE":  &tender.Int{Value: int64(gl.BLUE)},
	"BGRA":  &tender.Int{Value: int64(gl.BGRA)},
	"BGR":   &tender.Int{Value: int64(gl.BGR)},
	"UNSIGNED_BYTE":  &tender.Int{Value: int64(gl.UNSIGNED_BYTE)},
	"BYTE":           &tender.Int{Value: int64(gl.BYTE)},
	"UNSIGNED_SHORT": &tender.Int{Value: int64(gl.UNSIGNED_SHORT)},
	"SHORT":          &tender.Int{Value: int64(gl.SHORT)},
	"UNSIGNED_INT":   &tender.Int{Value: int64(gl.UNSIGNED_INT)},
	"INT":            &tender.Int{Value: int64(gl.INT)},
	"FLOAT":          &tender.Int{Value: int64(gl.FLOAT)},
	"DOUBLE":         &tender.Int{Value: int64(gl.DOUBLE)},

	// ==================== Vertex Array Client States ====================
	"VERTEX_ARRAY":       &tender.Int{Value: int64(gl.VERTEX_ARRAY)},
	"NORMAL_ARRAY":       &tender.Int{Value: int64(gl.NORMAL_ARRAY)},
	"COLOR_ARRAY":        &tender.Int{Value: int64(gl.COLOR_ARRAY)},
	"TEXTURE_COORD_ARRAY": &tender.Int{Value: int64(gl.TEXTURE_COORD_ARRAY)},

	// ==================== PixelStore Parameters ====================
	"PACK_ALIGNMENT":     &tender.Int{Value: int64(gl.PACK_ALIGNMENT)},
	"UNPACK_ALIGNMENT":   &tender.Int{Value: int64(gl.UNPACK_ALIGNMENT)},
	"PACK_ROW_LENGTH":    &tender.Int{Value: int64(gl.PACK_ROW_LENGTH)},
	"UNPACK_ROW_LENGTH":  &tender.Int{Value: int64(gl.UNPACK_ROW_LENGTH)},
	"PACK_SWAP_BYTES":    &tender.Int{Value: int64(gl.PACK_SWAP_BYTES)},
	"UNPACK_SWAP_BYTES":  &tender.Int{Value: int64(gl.UNPACK_SWAP_BYTES)},
	"PACK_LSB_FIRST":     &tender.Int{Value: int64(gl.PACK_LSB_FIRST)},
	"UNPACK_LSB_FIRST":   &tender.Int{Value: int64(gl.UNPACK_LSB_FIRST)},

	// ==================== Display List Modes ====================
	"COMPILE":             &tender.Int{Value: int64(gl.COMPILE)},
	"COMPILE_AND_EXECUTE": &tender.Int{Value: int64(gl.COMPILE_AND_EXECUTE)},

	// ==================== Accumulation Buffer Operations ====================
	"ACCUM":          &tender.Int{Value: int64(gl.ACCUM)},
	"LOAD":           &tender.Int{Value: int64(gl.LOAD)},
	"RETURN":         &tender.Int{Value: int64(gl.RETURN)},
	"MULT":           &tender.Int{Value: int64(gl.MULT)},
	"ADD":            &tender.Int{Value: int64(gl.ADD)},

	// ==================== Render Modes ====================
	"RENDER":          &tender.Int{Value: int64(gl.RENDER)},
	"FEEDBACK":        &tender.Int{Value: int64(gl.FEEDBACK)},
	"SELECT":          &tender.Int{Value: int64(gl.SELECT)},
	
	"INFO_LOG_LENGTH": &tender.Int{Value: int64(gl.INFO_LOG_LENGTH)},

	// ==================== Initialization ====================
	"init": &tender.NativeFunction{
		Name: "init",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			if err := gl.Init(); err != nil {
				return nil, err
			}
			return tender.NullValue, nil
		},
	},

	// ==================== Matrix Operations ====================
	"matrix_mode": &tender.NativeFunction{
		Name: "matrix_mode",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			mode, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.MatrixMode(mode)
			return tender.NullValue, nil
		},
	},

	"load_identity": &tender.NativeFunction{
		Name: "load_identity",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			gl.LoadIdentity()
			return tender.NullValue, nil
		},
	},

	"push_matrix": &tender.NativeFunction{
		Name: "push_matrix",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			gl.PushMatrix()
			return tender.NullValue, nil
		},
	},

	"pop_matrix": &tender.NativeFunction{
		Name: "pop_matrix",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			gl.PopMatrix()
			return tender.NullValue, nil
		},
	},

	"mult_matrixf": &tender.NativeFunction{
		Name: "mult_matrixf",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 16 {
				return nil, tender.ErrInvalidArgCount
			}
			var m [16]float32
			for i := 0; i < 16; i++ {
				val, ok := tender.ToFloat32(args[i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				m[i] = val
			}
			gl.MultMatrixf(&m[0])
			return tender.NullValue, nil
		},
	},

	"load_matrixf": &tender.NativeFunction{
		Name: "load_matrixf",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 16 {
				return nil, tender.ErrInvalidArgCount
			}
			var m [16]float32
			for i := 0; i < 16; i++ {
				val, ok := tender.ToFloat32(args[i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				m[i] = val
			}
			gl.LoadMatrixf(&m[0])
			return tender.NullValue, nil
		},
	},

	"translatef": &tender.NativeFunction{
		Name: "translatef",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := tender.ToFloat32(args[0])
			y, okY := tender.ToFloat32(args[1])
			z, okZ := tender.ToFloat32(args[2])
			if !okX || !okY || !okZ {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Translatef(x, y, z)
			return tender.NullValue, nil
		},
	},

	"rotatef": &tender.NativeFunction{
		Name: "rotatef",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			angle, okA := tender.ToFloat32(args[0])
			x, okX := tender.ToFloat32(args[1])
			y, okY := tender.ToFloat32(args[2])
			z, okZ := tender.ToFloat32(args[3])
			if !okA || !okX || !okY || !okZ {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Rotatef(angle, x, y, z)
			return tender.NullValue, nil
		},
	},

	"scalef": &tender.NativeFunction{
		Name: "scalef",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := tender.ToFloat32(args[0])
			y, okY := tender.ToFloat32(args[1])
			z, okZ := tender.ToFloat32(args[2])
			if !okX || !okY || !okZ {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Scalef(x, y, z)
			return tender.NullValue, nil
		},
	},

	// ==================== Viewport and Scissor ====================
	"viewport": &tender.NativeFunction{
		Name: "viewport",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := tender.ToInt32(args[0])
			y, okY := tender.ToInt32(args[1])
			w, okW := tender.ToInt32(args[2])
			h, okH := tender.ToInt32(args[3])
			if !okX || !okY || !okW || !okH {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Viewport(x, y, w, h)
			return tender.NullValue, nil
		},
	},

	"scissor": &tender.NativeFunction{
		Name: "scissor",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := tender.ToInt32(args[0])
			y, okY := tender.ToInt32(args[1])
			w, okW := tender.ToInt32(args[2])
			h, okH := tender.ToInt32(args[3])
			if !okX || !okY || !okW || !okH {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Scissor(x, y, w, h)
			return tender.NullValue, nil
		},
	},

	// ==================== Clearing ====================
	"clear": &tender.NativeFunction{
		Name: "clear",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			mask, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Clear(mask)
			return tender.NullValue, nil
		},
	},

	"clear_color": &tender.NativeFunction{
		Name: "clear_color",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			r, okR := tender.ToFloat32(args[0])
			g, okG := tender.ToFloat32(args[1])
			b, okB := tender.ToFloat32(args[2])
			a, okA := tender.ToFloat32(args[3])
			if !okR || !okG || !okB || !okA {
				return nil, tender.ErrInvalidArgCount
			}
			gl.ClearColor(r, g, b, a)
			return tender.NullValue, nil
		},
	},

	"clear_depth": &tender.NativeFunction{
		Name: "clear_depth",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			depth, ok := tender.ToFloat64(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.ClearDepth(depth)
			return tender.NullValue, nil
		},
	},

	"clear_stencil": &tender.NativeFunction{
		Name: "clear_stencil",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			s, ok := tender.ToInt32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.ClearStencil(s)
			return tender.NullValue, nil
		},
	},

	"clear_accum": &tender.NativeFunction{
		Name: "clear_accum",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			r, okR := tender.ToFloat32(args[0])
			g, okG := tender.ToFloat32(args[1])
			b, okB := tender.ToFloat32(args[2])
			a, okA := tender.ToFloat32(args[3])
			if !okR || !okG || !okB || !okA {
				return nil, tender.ErrInvalidArgCount
			}
			gl.ClearAccum(r, g, b, a)
			return tender.NullValue, nil
		},
	},

	// ==================== Primitives ====================
	"begin": &tender.NativeFunction{
		Name: "begin",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			mode, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Begin(mode)
			return tender.NullValue, nil
		},
	},

	"end": &tender.NativeFunction{
		Name: "end",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			gl.End()
			return tender.NullValue, nil
		},
	},

	"vertex2f": &tender.NativeFunction{
		Name: "vertex2f",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := tender.ToFloat32(args[0])
			y, okY := tender.ToFloat32(args[1])
			if !okX || !okY {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Vertex2f(x, y)
			return tender.NullValue, nil
		},
	},

	"vertex3f": &tender.NativeFunction{
		Name: "vertex3f",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := tender.ToFloat32(args[0])
			y, okY := tender.ToFloat32(args[1])
			z, okZ := tender.ToFloat32(args[2])
			if !okX || !okY || !okZ {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Vertex3f(x, y, z)
			return tender.NullValue, nil
		},
	},

	"vertex4f": &tender.NativeFunction{
		Name: "vertex4f",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := tender.ToFloat32(args[0])
			y, okY := tender.ToFloat32(args[1])
			z, okZ := tender.ToFloat32(args[2])
			w, okW := tender.ToFloat32(args[3])
			if !okX || !okY || !okZ || !okW {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Vertex4f(x, y, z, w)
			return tender.NullValue, nil
		},
	},

	"vertex2d": &tender.NativeFunction{
		Name: "vertex2d",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := tender.ToFloat64(args[0])
			y, okY := tender.ToFloat64(args[1])
			if !okX || !okY {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Vertex2d(x, y)
			return tender.NullValue, nil
		},
	},

	"vertex3d": &tender.NativeFunction{
		Name: "vertex3d",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := tender.ToFloat64(args[0])
			y, okY := tender.ToFloat64(args[1])
			z, okZ := tender.ToFloat64(args[2])
			if !okX || !okY || !okZ {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Vertex3d(x, y, z)
			return tender.NullValue, nil
		},
	},

	"vertex4d": &tender.NativeFunction{
		Name: "vertex4d",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := tender.ToFloat64(args[0])
			y, okY := tender.ToFloat64(args[1])
			z, okZ := tender.ToFloat64(args[2])
			w, okW := tender.ToFloat64(args[3])
			if !okX || !okY || !okZ || !okW {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Vertex4d(x, y, z, w)
			return tender.NullValue, nil
		},
	},

	"vertex2i": &tender.NativeFunction{
		Name: "vertex2i",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := tender.ToInt32(args[0])
			y, okY := tender.ToInt32(args[1])
			if !okX || !okY {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Vertex2i(x, y)
			return tender.NullValue, nil
		},
	},

	"vertex3i": &tender.NativeFunction{
		Name: "vertex3i",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := tender.ToInt32(args[0])
			y, okY := tender.ToInt32(args[1])
			z, okZ := tender.ToInt32(args[2])
			if !okX || !okY || !okZ {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Vertex3i(x, y, z)
			return tender.NullValue, nil
		},
	},

	// ==================== Colors ====================
	"color3f": &tender.NativeFunction{
		Name: "color3f",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			r, okR := tender.ToFloat32(args[0])
			g, okG := tender.ToFloat32(args[1])
			b, okB := tender.ToFloat32(args[2])
			if !okR || !okG || !okB {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Color3f(r, g, b)
			return tender.NullValue, nil
		},
	},

	"color4f": &tender.NativeFunction{
		Name: "color4f",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			r, okR := tender.ToFloat32(args[0])
			g, okG := tender.ToFloat32(args[1])
			b, okB := tender.ToFloat32(args[2])
			a, okA := tender.ToFloat32(args[3])
			if !okR || !okG || !okB || !okA {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Color4f(r, g, b, a)
			return tender.NullValue, nil
		},
	},

	"color3ub": &tender.NativeFunction{
		Name: "color3ub",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			r, okR := tender.ToUint8(args[0])
			g, okG := tender.ToUint8(args[1])
			b, okB := tender.ToUint8(args[2])
			if !okR || !okG || !okB {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Color3ub(r, g, b)
			return tender.NullValue, nil
		},
	},

	"color4ub": &tender.NativeFunction{
		Name: "color4ub",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			r, okR := tender.ToUint8(args[0])
			g, okG := tender.ToUint8(args[1])
			b, okB := tender.ToUint8(args[2])
			a, okA := tender.ToUint8(args[3])
			if !okR || !okG || !okB || !okA {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Color4ub(r, g, b, a)
			return tender.NullValue, nil
		},
	},

	// ==================== Normals ====================
	"normal3f": &tender.NativeFunction{
		Name: "normal3f",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			nx, okX := tender.ToFloat32(args[0])
			ny, okY := tender.ToFloat32(args[1])
			nz, okZ := tender.ToFloat32(args[2])
			if !okX || !okY || !okZ {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Normal3f(nx, ny, nz)
			return tender.NullValue, nil
		},
	},

	// ==================== Texture Coordinates ====================
	"tex_coord2f": &tender.NativeFunction{
		Name: "tex_coord2f",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			s, okS := tender.ToFloat32(args[0])
			t, okT := tender.ToFloat32(args[1])
			if !okS || !okT {
				return nil, tender.ErrInvalidArgCount
			}
			gl.TexCoord2f(s, t)
			return tender.NullValue, nil
		},
	},

	// ==================== Enable/Disable ====================
	"enable": &tender.NativeFunction{
		Name: "enable",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			cap, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Enable(cap)
			return tender.NullValue, nil
		},
	},

	"disable": &tender.NativeFunction{
		Name: "disable",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			cap, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Disable(cap)
			return tender.NullValue, nil
		},
	},

	"is_enabled": &tender.NativeFunction{
		Name: "is_enabled",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			cap, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			if gl.IsEnabled(cap) {
				return tender.TrueValue, nil
			}
			return tender.FalseValue, nil
		},
	},

	// ==================== Blending ====================
	"blend_func": &tender.NativeFunction{
		Name: "blend_func",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			sfactor, okS := tender.ToUint32(args[0])
			dfactor, okD := tender.ToUint32(args[1])
			if !okS || !okD {
				return nil, tender.ErrInvalidArgCount
			}
			gl.BlendFunc(sfactor, dfactor)
			return tender.NullValue, nil
		},
	},

	"blend_equation": &tender.NativeFunction{
		Name: "blend_equation",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			mode, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.BlendEquation(mode)
			return tender.NullValue, nil
		},
	},

	"blend_color": &tender.NativeFunction{
		Name: "blend_color",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			r, okR := tender.ToFloat32(args[0])
			g, okG := tender.ToFloat32(args[1])
			b, okB := tender.ToFloat32(args[2])
			a, okA := tender.ToFloat32(args[3])
			if !okR || !okG || !okB || !okA {
				return nil, tender.ErrInvalidArgCount
			}
			gl.BlendColor(r, g, b, a)
			return tender.NullValue, nil
		},
	},

	// ==================== Depth Test ====================
	"depth_func": &tender.NativeFunction{
		Name: "depth_func",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			fn, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.DepthFunc(fn)
			return tender.NullValue, nil
		},
	},

	"depth_mask": &tender.NativeFunction{
		Name: "depth_mask",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			flag := args[0] == tender.TrueValue
			if i, ok := tender.ToInt64(args[0]); ok && i != 0 {
				flag = true
			}
			gl.DepthMask(flag)
			return tender.NullValue, nil
		},
	},

	"depth_range": &tender.NativeFunction{
		Name: "depth_range",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			near, okN := tender.ToFloat64(args[0])
			far, okF := tender.ToFloat64(args[1])
			if !okN || !okF {
				return nil, tender.ErrInvalidArgCount
			}
			gl.DepthRange(near, far)
			return tender.NullValue, nil
		},
	},

	// ==================== Alpha Test ====================
	"alpha_func": &tender.NativeFunction{
		Name: "alpha_func",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			fn, okF := tender.ToUint32(args[0])
			ref, okR := tender.ToFloat32(args[1])
			if !okF || !okR {
				return nil, tender.ErrInvalidArgCount
			}
			gl.AlphaFunc(fn, ref)
			return tender.NullValue, nil
		},
	},

	// ==================== Culling ====================
	"cull_face": &tender.NativeFunction{
		Name: "cull_face",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			mode, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.CullFace(mode)
			return tender.NullValue, nil
		},
	},

	"front_face": &tender.NativeFunction{
		Name: "front_face",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			mode, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.FrontFace(mode)
			return tender.NullValue, nil
		},
	},

	// ==================== Polygon Modes ====================
	"polygon_mode": &tender.NativeFunction{
		Name: "polygon_mode",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			face, okF := tender.ToUint32(args[0])
			mode, okM := tender.ToUint32(args[1])
			if !okF || !okM {
				return nil, tender.ErrInvalidArgCount
			}
			gl.PolygonMode(face, mode)
			return tender.NullValue, nil
		},
	},

	"polygon_offset": &tender.NativeFunction{
		Name: "polygon_offset",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			factor, okF := tender.ToFloat32(args[0])
			units, okU := tender.ToFloat32(args[1])
			if !okF || !okU {
				return nil, tender.ErrInvalidArgCount
			}
			gl.PolygonOffset(factor, units)
			return tender.NullValue, nil
		},
	},

	// ==================== Line & Point Sizes ====================
	"line_width": &tender.NativeFunction{
		Name: "line_width",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			width, ok := tender.ToFloat32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.LineWidth(width)
			return tender.NullValue, nil
		},
	},

	"point_size": &tender.NativeFunction{
		Name: "point_size",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			size, ok := tender.ToFloat32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.PointSize(size)
			return tender.NullValue, nil
		},
	},

	// ==================== Stencil ====================
	"stencil_func": &tender.NativeFunction{
		Name: "stencil_func",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			fn, okF := tender.ToUint32(args[0])
			ref, okR := tender.ToInt32(args[1])
			mask, okM := tender.ToUint32(args[2])
			if !okF || !okR || !okM {
				return nil, tender.ErrInvalidArgCount
			}
			gl.StencilFunc(fn, ref, mask)
			return tender.NullValue, nil
		},
	},

	"stencil_mask": &tender.NativeFunction{
		Name: "stencil_mask",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			mask, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.StencilMask(mask)
			return tender.NullValue, nil
		},
	},

	"stencil_op": &tender.NativeFunction{
		Name: "stencil_op",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			fail, okF := tender.ToUint32(args[0])
			zfail, okZ := tender.ToUint32(args[1])
			zpass, okP := tender.ToUint32(args[2])
			if !okF || !okZ || !okP {
				return nil, tender.ErrInvalidArgCount
			}
			gl.StencilOp(fail, zfail, zpass)
			return tender.NullValue, nil
		},
	},

	// ==================== State Saving ====================
	"push_attrib": &tender.NativeFunction{
		Name: "push_attrib",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			mask, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.PushAttrib(mask)
			return tender.NullValue, nil
		},
	},

	"pop_attrib": &tender.NativeFunction{
		Name: "pop_attrib",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			gl.PopAttrib()
			return tender.NullValue, nil
		},
	},

	"push_client_attrib": &tender.NativeFunction{
		Name: "push_client_attrib",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			mask, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.PushClientAttrib(mask)
			return tender.NullValue, nil
		},
	},

	"pop_client_attrib": &tender.NativeFunction{
		Name: "pop_client_attrib",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			gl.PopClientAttrib()
			return tender.NullValue, nil
		},
	},

	// ==================== Shading ====================
	"shade_model": &tender.NativeFunction{
		Name: "shade_model",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			mode, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.ShadeModel(mode)
			return tender.NullValue, nil
		},
	},

	// ==================== Hints ====================
	"hint": &tender.NativeFunction{
		Name: "hint",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			target, okT := tender.ToUint32(args[0])
			mode, okM := tender.ToUint32(args[1])
			if !okT || !okM {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Hint(target, mode)
			return tender.NullValue, nil
		},
	},

	// ==================== Lighting ====================
	"lightf": &tender.NativeFunction{
		Name: "lightf",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			light, okL := tender.ToUint32(args[0])
			pname, okP := tender.ToUint32(args[1])
			param, okV := tender.ToFloat32(args[2])
			if !okL || !okP || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Lightf(light, pname, param)
			return tender.NullValue, nil
		},
	},

	"lightfv": &tender.NativeFunction{
		Name: "lightfv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			// Accept variable number of floats (4 for color/position, 1 for scalar)
			if len(args) < 3 {
				return nil, tender.ErrInvalidArgCount
			}
			light, okL := tender.ToUint32(args[0])
			if !okL {
				return nil, tender.ErrInvalidArgCount
			}
			pname, okP := tender.ToUint32(args[1])
			if !okP {
				return nil, tender.ErrInvalidArgCount
			}
			// Rest are params
			params := make([]float32, len(args)-2)
			for i := 2; i < len(args); i++ {
				val, ok := tender.ToFloat32(args[i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				params[i-2] = val
			}
			gl.Lightfv(light, pname, &params[0])
			return tender.NullValue, nil
		},
	},

	"lighti": &tender.NativeFunction{
		Name: "lighti",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			light, okL := tender.ToUint32(args[0])
			pname, okP := tender.ToUint32(args[1])
			param, okV := tender.ToInt32(args[2])
			if !okL || !okP || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Lighti(light, pname, param)
			return tender.NullValue, nil
		},
	},

	"light_modelf": &tender.NativeFunction{
		Name: "light_modelf",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			pname, okP := tender.ToUint32(args[0])
			param, okV := tender.ToFloat32(args[1])
			if !okP || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.LightModelf(pname, param)
			return tender.NullValue, nil
		},
	},

	"light_modeli": &tender.NativeFunction{
		Name: "light_modeli",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			pname, okP := tender.ToUint32(args[0])
			param, okV := tender.ToInt32(args[1])
			if !okP || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.LightModeli(pname, param)
			return tender.NullValue, nil
		},
	},

	"materialf": &tender.NativeFunction{
		Name: "materialf",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			face, okF := tender.ToUint32(args[0])
			pname, okP := tender.ToUint32(args[1])
			param, okV := tender.ToFloat32(args[2])
			if !okF || !okP || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Materialf(face, pname, param)
			return tender.NullValue, nil
		},
	},

	"materialfv": &tender.NativeFunction{
		Name: "materialfv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) < 3 {
				return nil, tender.ErrInvalidArgCount
			}
			face, okF := tender.ToUint32(args[0])
			if !okF {
				return nil, tender.ErrInvalidArgCount
			}
			pname, okP := tender.ToUint32(args[1])
			if !okP {
				return nil, tender.ErrInvalidArgCount
			}
			params := make([]float32, len(args)-2)
			for i := 2; i < len(args); i++ {
				val, ok := tender.ToFloat32(args[i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				params[i-2] = val
			}
			gl.Materialfv(face, pname, &params[0])
			return tender.NullValue, nil
		},
	},

	// ==================== Fog ====================
	"fogf": &tender.NativeFunction{
		Name: "fogf",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			pname, okP := tender.ToUint32(args[0])
			param, okV := tender.ToFloat32(args[1])
			if !okP || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Fogf(pname, param)
			return tender.NullValue, nil
		},
	},

	"fogfv": &tender.NativeFunction{
		Name: "fogfv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) < 2 {
				return nil, tender.ErrInvalidArgCount
			}
			pname, okP := tender.ToUint32(args[0])
			if !okP {
				return nil, tender.ErrInvalidArgCount
			}
			params := make([]float32, len(args)-1)
			for i := 1; i < len(args); i++ {
				val, ok := tender.ToFloat32(args[i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				params[i-1] = val
			}
			gl.Fogfv(pname, &params[0])
			return tender.NullValue, nil
		},
	},

	// ==================== Textures ====================
	"gen_textures": &tender.NativeFunction{
		Name: "gen_textures",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			n, ok := tender.ToInt32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			textures := make([]uint32, n)
			gl.GenTextures(n, &textures[0])
			// Return as array of ints
			arr := make([]tender.Object, n)
			for i, t := range textures {
				arr[i] = &tender.Int{Value: int64(t)}
			}
			return &tender.Array{Value: arr}, nil
		},
	},

	"bind_texture": &tender.NativeFunction{
		Name: "bind_texture",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			target, okT := tender.ToUint32(args[0])
			texture, okTex := tender.ToUint32(args[1])
			if !okT || !okTex {
				return nil, tender.ErrInvalidArgCount
			}
			gl.BindTexture(target, texture)
			return tender.NullValue, nil
		},
	},

	"delete_textures": &tender.NativeFunction{
		Name: "delete_textures",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			// Expect an array of ints or a single int
			if arr, ok := args[0].(*tender.Array); ok {
				textures := make([]uint32, len(arr.Value))
				for i, obj := range arr.Value {
					t, ok := tender.ToUint32(obj)
					if !ok {
						return nil, tender.ErrInvalidArgCount
					}
					textures[i] = t
				}
				gl.DeleteTextures(int32(len(textures)), &textures[0])
				return tender.NullValue, nil
			}
			// Single texture
			if t, ok := tender.ToUint32(args[0]); ok {
				gl.DeleteTextures(1, &t)
				return tender.NullValue, nil
			}
			return nil, tender.ErrInvalidArgCount
		},
	},

	"tex_image2d": &tender.NativeFunction{
		Name: "tex_image2d",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 9 {
				return nil, tender.ErrInvalidArgCount
			}
			target, _ := tender.ToUint32(args[0])
			level, _ := tender.ToInt32(args[1])
			internalFormat, _ := tender.ToInt32(args[2])
			width, _ := tender.ToInt32(args[3])
			height, _ := tender.ToInt32(args[4])
			border, _ := tender.ToInt32(args[5])
			format, _ := tender.ToUint32(args[6])
			typ, _ := tender.ToUint32(args[7])
			// Pixel data can be bytes array or null
			if args[8] == tender.NullValue {
				gl.TexImage2D(target, level, internalFormat, width, height, border, format, typ, nil)
				return tender.NullValue, nil
			}
			// Try to get bytes
			bytes, ok := tender.ToByteSlice(args[8])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			if len(bytes) == 0 {
				gl.TexImage2D(target, level, internalFormat, width, height, border, format, typ, nil)
			} else {
				gl.TexImage2D(target, level, internalFormat, width, height, border, format, typ, unsafe.Pointer(&bytes[0]))
			}
			return tender.NullValue, nil
		},
	},

	"tex_sub_image2d": &tender.NativeFunction{
		Name: "tex_sub_image2d",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 9 {
				return nil, tender.ErrInvalidArgCount
			}
			target, _ := tender.ToUint32(args[0])
			level, _ := tender.ToInt32(args[1])
			xoffset, _ := tender.ToInt32(args[2])
			yoffset, _ := tender.ToInt32(args[3])
			width, _ := tender.ToInt32(args[4])
			height, _ := tender.ToInt32(args[5])
			format, _ := tender.ToUint32(args[6])
			typ, _ := tender.ToUint32(args[7])
			// Pixel data must be bytes
			bytes, ok := tender.ToByteSlice(args[8])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			if len(bytes) == 0 {
				gl.TexSubImage2D(target, level, xoffset, yoffset, width, height, format, typ, nil)
			} else {
				gl.TexSubImage2D(target, level, xoffset, yoffset, width, height, format, typ, unsafe.Pointer(&bytes[0]))
			}
			return tender.NullValue, nil
		},
	},

	"tex_parameterf": &tender.NativeFunction{
		Name: "tex_parameterf",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			target, okT := tender.ToUint32(args[0])
			pname, okP := tender.ToUint32(args[1])
			param, okV := tender.ToFloat32(args[2])
			if !okT || !okP || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.TexParameterf(target, pname, param)
			return tender.NullValue, nil
		},
	},

	"tex_parameteri": &tender.NativeFunction{
		Name: "tex_parameteri",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			target, okT := tender.ToUint32(args[0])
			pname, okP := tender.ToUint32(args[1])
			param, okV := tender.ToInt32(args[2])
			if !okT || !okP || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.TexParameteri(target, pname, param)
			return tender.NullValue, nil
		},
	},

	"tex_parameteriv": &tender.NativeFunction{
		Name: "tex_parameteriv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) < 3 {
				return nil, tender.ErrInvalidArgCount
			}
			target, okT := tender.ToUint32(args[0])
			if !okT {
				return nil, tender.ErrInvalidArgCount
			}
			pname, okP := tender.ToUint32(args[1])
			if !okP {
				return nil, tender.ErrInvalidArgCount
			}
			params := make([]int32, len(args)-2)
			for i := 2; i < len(args); i++ {
				val, ok := tender.ToInt32(args[i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				params[i-2] = val
			}
			gl.TexParameteriv(target, pname, &params[0])
			return tender.NullValue, nil
		},
	},

	"tex_parameterfv": &tender.NativeFunction{
		Name: "tex_parameterfv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) < 3 {
				return nil, tender.ErrInvalidArgCount
			}
			target, okT := tender.ToUint32(args[0])
			if !okT {
				return nil, tender.ErrInvalidArgCount
			}
			pname, okP := tender.ToUint32(args[1])
			if !okP {
				return nil, tender.ErrInvalidArgCount
			}
			params := make([]float32, len(args)-2)
			for i := 2; i < len(args); i++ {
				val, ok := tender.ToFloat32(args[i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				params[i-2] = val
			}
			gl.TexParameterfv(target, pname, &params[0])
			return tender.NullValue, nil
		},
	},

	"tex_envf": &tender.NativeFunction{
		Name: "tex_envf",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			target, okT := tender.ToUint32(args[0])
			pname, okP := tender.ToUint32(args[1])
			param, okV := tender.ToFloat32(args[2])
			if !okT || !okP || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.TexEnvf(target, pname, param)
			return tender.NullValue, nil
		},
	},

	"tex_envi": &tender.NativeFunction{
		Name: "tex_envi",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			target, okT := tender.ToUint32(args[0])
			pname, okP := tender.ToUint32(args[1])
			param, okV := tender.ToInt32(args[2])
			if !okT || !okP || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.TexEnvi(target, pname, param)
			return tender.NullValue, nil
		},
	},

	"tex_envfv": &tender.NativeFunction{
		Name: "tex_envfv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) < 3 {
				return nil, tender.ErrInvalidArgCount
			}
			target, okT := tender.ToUint32(args[0])
			if !okT {
				return nil, tender.ErrInvalidArgCount
			}
			pname, okP := tender.ToUint32(args[1])
			if !okP {
				return nil, tender.ErrInvalidArgCount
			}
			params := make([]float32, len(args)-2)
			for i := 2; i < len(args); i++ {
				val, ok := tender.ToFloat32(args[i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				params[i-2] = val
			}
			gl.TexEnvfv(target, pname, &params[0])
			return tender.NullValue, nil
		},
	},

	"tex_enviv": &tender.NativeFunction{
		Name: "tex_enviv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) < 3 {
				return nil, tender.ErrInvalidArgCount
			}
			target, okT := tender.ToUint32(args[0])
			if !okT {
				return nil, tender.ErrInvalidArgCount
			}
			pname, okP := tender.ToUint32(args[1])
			if !okP {
				return nil, tender.ErrInvalidArgCount
			}
			params := make([]int32, len(args)-2)
			for i := 2; i < len(args); i++ {
				val, ok := tender.ToInt32(args[i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				params[i-2] = val
			}
			gl.TexEnviv(target, pname, &params[0])
			return tender.NullValue, nil
		},
	},

	"active_texture": &tender.NativeFunction{
		Name: "active_texture",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			texture, _ := tender.ToUint32(args[0])
			gl.ActiveTexture(texture)
			return tender.NullValue, nil
		},
	},

	"tex_genf": &tender.NativeFunction{
		Name: "tex_genf",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			coord, okC := tender.ToUint32(args[0])
			pname, okP := tender.ToUint32(args[1])
			param, okV := tender.ToFloat32(args[2])
			if !okC || !okP || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.TexGenf(coord, pname, param)
			return tender.NullValue, nil
		},
	},

	"tex_gend": &tender.NativeFunction{
		Name: "tex_gend",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			coord, okC := tender.ToUint32(args[0])
			pname, okP := tender.ToUint32(args[1])
			param, okV := tender.ToFloat64(args[2])
			if !okC || !okP || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.TexGend(coord, pname, param)
			return tender.NullValue, nil
		},
	},

	"tex_geni": &tender.NativeFunction{
		Name: "tex_geni",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			coord, okC := tender.ToUint32(args[0])
			pname, okP := tender.ToUint32(args[1])
			param, okV := tender.ToInt32(args[2])
			if !okC || !okP || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.TexGeni(coord, pname, param)
			return tender.NullValue, nil
		},
	},

	// ==================== Vertex Arrays ====================
	"enable_client_state": &tender.NativeFunction{
		Name: "enable_client_state",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			array, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.EnableClientState(array)
			return tender.NullValue, nil
		},
	},

	"disable_client_state": &tender.NativeFunction{
		Name: "disable_client_state",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			array, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.DisableClientState(array)
			return tender.NullValue, nil
		},
	},

	"vertex_pointer": &tender.NativeFunction{
		Name: "vertex_pointer",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			size, _ := tender.ToInt32(args[0])
			typ, _ := tender.ToUint32(args[1])
			stride, _ := tender.ToInt32(args[2])
			// The pointer must be a bytes object or null
			var ptr unsafe.Pointer
			if args[3] != tender.NullValue {
				bytes, ok := tender.ToByteSlice(args[3])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				ptr = unsafe.Pointer(&bytes[0])
			}
			gl.VertexPointer(size, typ, stride, ptr)
			return tender.NullValue, nil
		},
	},

	"normal_pointer": &tender.NativeFunction{
		Name: "normal_pointer",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			typ, _ := tender.ToUint32(args[0])
			stride, _ := tender.ToInt32(args[1])
			var ptr unsafe.Pointer
			if args[2] != tender.NullValue {
				bytes, ok := tender.ToByteSlice(args[2])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				ptr = unsafe.Pointer(&bytes[0])
			}
			gl.NormalPointer(typ, stride, ptr)
			return tender.NullValue, nil
		},
	},

	"color_pointer": &tender.NativeFunction{
		Name: "color_pointer",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			size, _ := tender.ToInt32(args[0])
			typ, _ := tender.ToUint32(args[1])
			stride, _ := tender.ToInt32(args[2])
			var ptr unsafe.Pointer
			if args[3] != tender.NullValue {
				bytes, ok := tender.ToByteSlice(args[3])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				ptr = unsafe.Pointer(&bytes[0])
			}
			gl.ColorPointer(size, typ, stride, ptr)
			return tender.NullValue, nil
		},
	},

	"tex_coord_pointer": &tender.NativeFunction{
		Name: "tex_coord_pointer",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			size, _ := tender.ToInt32(args[0])
			typ, _ := tender.ToUint32(args[1])
			stride, _ := tender.ToInt32(args[2])
			var ptr unsafe.Pointer
			if args[3] != tender.NullValue {
				bytes, ok := tender.ToByteSlice(args[3])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				ptr = unsafe.Pointer(&bytes[0])
			}
			gl.TexCoordPointer(size, typ, stride, ptr)
			return tender.NullValue, nil
		},
	},

	"draw_arrays": &tender.NativeFunction{
		Name: "draw_arrays",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			mode, okM := tender.ToUint32(args[0])
			first, okF := tender.ToInt32(args[1])
			count, okC := tender.ToInt32(args[2])
			if !okM || !okF || !okC {
				return nil, tender.ErrInvalidArgCount
			}
			gl.DrawArrays(mode, first, count)
			return tender.NullValue, nil
		},
	},

	"draw_elements": &tender.NativeFunction{
		Name: "draw_elements",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			mode, _ := tender.ToUint32(args[0])
			count, _ := tender.ToInt32(args[1])
			typ, _ := tender.ToUint32(args[2])
			var ptr unsafe.Pointer
			if args[3] != tender.NullValue {
				bytes, ok := tender.ToByteSlice(args[3])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				ptr = unsafe.Pointer(&bytes[0])
			}
			gl.DrawElements(mode, count, typ, ptr)
			return tender.NullValue, nil
		},
	},

	// ==================== Display Lists ====================
	"gen_lists": &tender.NativeFunction{
		Name: "gen_lists",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			range_, ok := tender.ToInt32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			list := gl.GenLists(range_)
			return &tender.Int{Value: int64(list)}, nil
		},
	},

	"new_list": &tender.NativeFunction{
		Name: "new_list",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			list, okL := tender.ToUint32(args[0])
			mode, okM := tender.ToUint32(args[1])
			if !okL || !okM {
				return nil, tender.ErrInvalidArgCount
			}
			gl.NewList(list, mode)
			return tender.NullValue, nil
		},
	},

	"end_list": &tender.NativeFunction{
		Name: "end_list",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			gl.EndList()
			return tender.NullValue, nil
		},
	},

	"call_list": &tender.NativeFunction{
		Name: "call_list",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			list, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.CallList(list)
			return tender.NullValue, nil
		},
	},

	"delete_lists": &tender.NativeFunction{
		Name: "delete_lists",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			list, okL := tender.ToUint32(args[0])
			range_, okR := tender.ToInt32(args[1])
			if !okL || !okR {
				return nil, tender.ErrInvalidArgCount
			}
			gl.DeleteLists(list, range_)
			return tender.NullValue, nil
		},
	},

	// ==================== Pixel Storage ====================
	"pixel_storei": &tender.NativeFunction{
		Name: "pixel_storei",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			pname, okP := tender.ToUint32(args[0])
			param, okV := tender.ToInt32(args[1])
			if !okP || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.PixelStorei(pname, param)
			return tender.NullValue, nil
		},
	},

	// ==================== Query & Sync ====================
	"get_error": &tender.NativeFunction{
		Name: "get_error",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			err := gl.GetError()
			return &tender.Int{Value: int64(err)}, nil
		},
	},

	"get_integerv": &tender.NativeFunction{
		Name: "get_integerv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			pname, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			var val int32
			gl.GetIntegerv(pname, &val)
			return &tender.Int{Value: int64(val)}, nil
		},
	},

	"get_floatv": &tender.NativeFunction{
		Name: "get_floatv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			pname, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			var val float32
			gl.GetFloatv(pname, &val)
			return &tender.Float{Value: float64(val)}, nil
		},
	},

	"get_doublev": &tender.NativeFunction{
		Name: "get_doublev",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			pname, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			var val float64
			gl.GetDoublev(pname, &val)
			return &tender.Float{Value: val}, nil
		},
	},

	"get_booleanv": &tender.NativeFunction{
		Name: "get_booleanv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			pname, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			var val bool
			gl.GetBooleanv(pname, &val)
			if val {
				return tender.TrueValue, nil
			}
			return tender.FalseValue, nil
		},
	},

	"get_string": &tender.NativeFunction{
		Name: "get_string",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			name, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			str := gl.GetString(name)
			if str == nil {
				return tender.NullValue, nil
			}
			// Convert *uint8 to Go string
			// The C string is null-terminated
			ptr := unsafe.Pointer(str)
			// Use cgo? But we can't import cgo directly; we assume gl package provides a way.
			// The gl package's GetString returns *uint8, but we can use unsafe and convert to byte slice.
			// Better to use the gl package's helper if available, but we'll do a simple conversion.
			// Since we don't have cgo, we'll assume the gl package returns *C.char, but in the Go binding it's *uint8.
			// We'll convert using a loop.
			var bytes []byte
			for i := 0; ; i++ {
				b := *(*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(i)))
				if b == 0 {
					break
				}
				bytes = append(bytes, b)
			}
			return &tender.String{Value: string(bytes)}, nil
		},
	},

	// ==================== Flush & Finish ====================
	"flush": &tender.NativeFunction{
		Name: "flush",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Flush()
			return tender.NullValue, nil
		},
	},

	"finish": &tender.NativeFunction{
		Name: "finish",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Finish()
			return tender.NullValue, nil
		},
	},

	// ==================== Accumulation Buffer ====================
	"accum": &tender.NativeFunction{
		Name: "accum",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			op, okO := tender.ToUint32(args[0])
			value, okV := tender.ToFloat32(args[1])
			if !okO || !okV {
				return nil, tender.ErrInvalidArgCount
			}
			gl.Accum(op, value)
			return tender.NullValue, nil
		},
	},

	// ==================== Render Mode ====================
	"render_mode": &tender.NativeFunction{
		Name: "render_mode",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			mode, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			result := gl.RenderMode(mode)
			return &tender.Int{Value: int64(result)}, nil
		},
	},
	// ==================== Matrix Projection ====================
	"ortho": &tender.NativeFunction{
		Name: "ortho",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 6 {
				return nil, tender.ErrInvalidArgCount
			}
			left, _ := tender.ToFloat64(args[0])
			right, _ := tender.ToFloat64(args[1])
			bottom, _ := tender.ToFloat64(args[2])
			top, _ := tender.ToFloat64(args[3])
			zNear, _ := tender.ToFloat64(args[4])
			zFar, _ := tender.ToFloat64(args[5])
			gl.Ortho(left, right, bottom, top, zNear, zFar)
			return tender.NullValue, nil
		},
	},

	"frustum": &tender.NativeFunction{
		Name: "frustum",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 6 {
				return nil, tender.ErrInvalidArgCount
			}
			left, _ := tender.ToFloat64(args[0])
			right, _ := tender.ToFloat64(args[1])
			bottom, _ := tender.ToFloat64(args[2])
			top, _ := tender.ToFloat64(args[3])
			zNear, _ := tender.ToFloat64(args[4])
			zFar, _ := tender.ToFloat64(args[5])
			gl.Frustum(left, right, bottom, top, zNear, zFar)
			return tender.NullValue, nil
		},
	},

	// ==================== Raster Position ====================
	"raster_pos2i": &tender.NativeFunction{
		Name: "raster_pos2i",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			x, _ := tender.ToInt32(args[0])
			y, _ := tender.ToInt32(args[1])
			gl.RasterPos2i(x, y)
			return tender.NullValue, nil
		},
	},

	"raster_pos2f": &tender.NativeFunction{
		Name: "raster_pos2f",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			x, _ := tender.ToFloat32(args[0])
			y, _ := tender.ToFloat32(args[1])
			gl.RasterPos2f(x, y)
			return tender.NullValue, nil
		},
	},

	"raster_pos3f": &tender.NativeFunction{
		Name: "raster_pos3f",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			x, _ := tender.ToFloat32(args[0])
			y, _ := tender.ToFloat32(args[1])
			z, _ := tender.ToFloat32(args[2])
			gl.RasterPos3f(x, y, z)
			return tender.NullValue, nil
		},
	},

	// ==================== Bitmap & Pixel Drawing ====================
	"bitmap": &tender.NativeFunction{
		Name: "bitmap",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 7 {
				return nil, tender.ErrInvalidArgCount
			}
			width, _ := tender.ToInt32(args[0])
			height, _ := tender.ToInt32(args[1])
			xorig, _ := tender.ToFloat32(args[2])
			yorig, _ := tender.ToFloat32(args[3])
			xmove, _ := tender.ToFloat32(args[4])
			ymove, _ := tender.ToFloat32(args[5])
			if args[6] == tender.NullValue {
				gl.Bitmap(width, height, xorig, yorig, xmove, ymove, nil)
			} else {
				bytes, ok := tender.ToByteSlice(args[6])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				// bytes slice may be empty; use &bytes[0] only if non‑empty
				if len(bytes) == 0 {
					gl.Bitmap(width, height, xorig, yorig, xmove, ymove, nil)
				} else {
					gl.Bitmap(width, height, xorig, yorig, xmove, ymove, &bytes[0])
				}
			}
			return tender.NullValue, nil
		},
	},

	"draw_pixels": &tender.NativeFunction{
		Name: "draw_pixels",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 5 {
				return nil, tender.ErrInvalidArgCount
			}
			width, _ := tender.ToInt32(args[0])
			height, _ := tender.ToInt32(args[1])
			format, _ := tender.ToUint32(args[2])
			typ, _ := tender.ToUint32(args[3])
			bytes, ok := tender.ToByteSlice(args[4])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.DrawPixels(width, height, format, typ, unsafe.Pointer(&bytes[0]))
			return tender.NullValue, nil
		},
	},

	// ==================== Pixel Read / Copy / Zoom ====================
	"read_pixels": &tender.NativeFunction{
		Name: "read_pixels",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 7 {
				return nil, tender.ErrInvalidArgCount
			}
			x, _ := tender.ToInt32(args[0])
			y, _ := tender.ToInt32(args[1])
			width, _ := tender.ToInt32(args[2])
			height, _ := tender.ToInt32(args[3])
			format, _ := tender.ToUint32(args[4])
			typ, _ := tender.ToUint32(args[5])
			bytes, ok := tender.ToByteSlice(args[6])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.ReadPixels(x, y, width, height, format, typ, unsafe.Pointer(&bytes[0]))
			return tender.NullValue, nil
		},
	},

	"copy_pixels": &tender.NativeFunction{
		Name: "copy_pixels",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 5 {
				return nil, tender.ErrInvalidArgCount
			}
			x, _ := tender.ToInt32(args[0])
			y, _ := tender.ToInt32(args[1])
			width, _ := tender.ToInt32(args[2])
			height, _ := tender.ToInt32(args[3])
			typ, _ := tender.ToUint32(args[4])
			gl.CopyPixels(x, y, width, height, typ)
			return tender.NullValue, nil
		},
	},

	"pixel_zoom": &tender.NativeFunction{
		Name: "pixel_zoom",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			xfactor, _ := tender.ToFloat32(args[0])
			yfactor, _ := tender.ToFloat32(args[1])
			gl.PixelZoom(xfactor, yfactor)
			return tender.NullValue, nil
		},
	},

	// ==================== Colour Mask & Logic Op ====================
	"color_mask": &tender.NativeFunction{
		Name: "color_mask",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			red, _ := tender.ToBool(args[0])
			green, _ := tender.ToBool(args[1])
			blue, _ := tender.ToBool(args[2])
			alpha, _ := tender.ToBool(args[3])
			gl.ColorMask(red, green, blue, alpha)
			return tender.NullValue, nil
		},
	},

	"logic_op": &tender.NativeFunction{
		Name: "logic_op",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			opcode, _ := tender.ToUint32(args[0])
			gl.LogicOp(opcode)
			return tender.NullValue, nil
		},
	},

	// ==================== Color Material ====================
	"color_material": &tender.NativeFunction{
		Name: "color_material",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			face, _ := tender.ToUint32(args[0])
			mode, _ := tender.ToUint32(args[1])
			gl.ColorMaterial(face, mode)
			return tender.NullValue, nil
		},
	},

	// ==================== Clipping Planes ====================
	"clip_plane": &tender.NativeFunction{
		Name: "clip_plane",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 5 {
				return nil, tender.ErrInvalidArgCount
			}
			plane, _ := tender.ToUint32(args[0])
			eq := [4]float64{}
			for i := 0; i < 4; i++ {
				val, ok := tender.ToFloat64(args[i+1])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				eq[i] = val
			}
			gl.ClipPlane(plane, &eq[0])
			return tender.NullValue, nil
		},
	},

	// ==================== 1D / 3D Textures ====================
	"tex_image1d": &tender.NativeFunction{
		Name: "tex_image1d",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 8 {
				return nil, tender.ErrInvalidArgCount
			}
			target, _ := tender.ToUint32(args[0])
			level, _ := tender.ToInt32(args[1])
			internalFormat, _ := tender.ToInt32(args[2])
			width, _ := tender.ToInt32(args[3])
			border, _ := tender.ToInt32(args[4])
			format, _ := tender.ToUint32(args[5])
			typ, _ := tender.ToUint32(args[6])
			var ptr unsafe.Pointer
			if args[7] != tender.NullValue {
				bytes, ok := tender.ToByteSlice(args[7])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				ptr = unsafe.Pointer(&bytes[0])
			}
			gl.TexImage1D(target, level, internalFormat, width, border, format, typ, ptr)
			return tender.NullValue, nil
		},
	},

	"tex_image3d": &tender.NativeFunction{
		Name: "tex_image3d",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 10 {
				return nil, tender.ErrInvalidArgCount
			}
			target, _ := tender.ToUint32(args[0])
			level, _ := tender.ToInt32(args[1])
			internalFormat, _ := tender.ToInt32(args[2])
			width, _ := tender.ToInt32(args[3])
			height, _ := tender.ToInt32(args[4])
			depth, _ := tender.ToInt32(args[5])
			border, _ := tender.ToInt32(args[6])
			format, _ := tender.ToUint32(args[7])
			typ, _ := tender.ToUint32(args[8])
			var ptr unsafe.Pointer
			if args[9] != tender.NullValue {
				bytes, ok := tender.ToByteSlice(args[9])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				ptr = unsafe.Pointer(&bytes[0])
			}
			gl.TexImage3D(target, level, internalFormat, width, height, depth, border, format, typ, ptr)
			return tender.NullValue, nil
		},
	},

	"tex_sub_image1d": &tender.NativeFunction{
		Name: "tex_sub_image1d",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 7 {
				return nil, tender.ErrInvalidArgCount
			}
			target, _ := tender.ToUint32(args[0])
			level, _ := tender.ToInt32(args[1])
			xoffset, _ := tender.ToInt32(args[2])
			width, _ := tender.ToInt32(args[3])
			format, _ := tender.ToUint32(args[4])
			typ, _ := tender.ToUint32(args[5])
			bytes, ok := tender.ToByteSlice(args[6])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.TexSubImage1D(target, level, xoffset, width, format, typ, unsafe.Pointer(&bytes[0]))
			return tender.NullValue, nil
		},
	},

	"tex_sub_image3d": &tender.NativeFunction{
		Name: "tex_sub_image3d",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 10 { // now 10 arguments: target, level, xoffset, yoffset, zoffset, width, height, depth, format, typ, data
				return nil, tender.ErrInvalidArgCount
			}
			target, _ := tender.ToUint32(args[0])
			level, _ := tender.ToInt32(args[1])
			xoffset, _ := tender.ToInt32(args[2])
			yoffset, _ := tender.ToInt32(args[3])
			zoffset, _ := tender.ToInt32(args[4])
			width, _ := tender.ToInt32(args[5])
			height, _ := tender.ToInt32(args[6])
			depth, _ := tender.ToInt32(args[7])
			format, _ := tender.ToUint32(args[8])
			typ, _ := tender.ToUint32(args[9])
			bytes, ok := tender.ToByteSlice(args[10]) // data is the 11th argument
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			var ptr unsafe.Pointer
			if len(bytes) > 0 {
				ptr = unsafe.Pointer(&bytes[0])
			}
			gl.TexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, typ, ptr)
			return tender.NullValue, nil
		},
	},
	
	// ==================== Copy Texture ====================
	"copy_tex_image2d": &tender.NativeFunction{
		Name: "copy_tex_image2d",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 8 {
				return nil, tender.ErrInvalidArgCount
			}
			target, _ := tender.ToUint32(args[0])
			level, _ := tender.ToInt32(args[1])
			internalFormat, _ := tender.ToUint32(args[2])
			x, _ := tender.ToInt32(args[3])
			y, _ := tender.ToInt32(args[4])
			width, _ := tender.ToInt32(args[5])
			height, _ := tender.ToInt32(args[6])
			border, _ := tender.ToInt32(args[7])
			gl.CopyTexImage2D(target, level, internalFormat, x, y, width, height, border)
			return tender.NullValue, nil
		},
	},

	"copy_tex_sub_image2d": &tender.NativeFunction{
		Name: "copy_tex_sub_image2d",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 8 {
				return nil, tender.ErrInvalidArgCount
			}
			target, _ := tender.ToUint32(args[0])
			level, _ := tender.ToInt32(args[1])
			xoffset, _ := tender.ToInt32(args[2])
			yoffset, _ := tender.ToInt32(args[3])
			x, _ := tender.ToInt32(args[4])
			y, _ := tender.ToInt32(args[5])
			width, _ := tender.ToInt32(args[6])
			height, _ := tender.ToInt32(args[7])
			gl.CopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height)
			return tender.NullValue, nil
		},
	},

	// ==================== Texture Queries ====================
	"get_tex_parameteriv": &tender.NativeFunction{
		Name: "get_tex_parameteriv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			target, _ := tender.ToUint32(args[0])
			pname, _ := tender.ToUint32(args[1])
			var val int32
			gl.GetTexParameteriv(target, pname, &val)
			return &tender.Int{Value: int64(val)}, nil
		},
	},

	"get_tex_parameterfv": &tender.NativeFunction{
		Name: "get_tex_parameterfv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			target, _ := tender.ToUint32(args[0])
			pname, _ := tender.ToUint32(args[1])
			var val float32
			gl.GetTexParameterfv(target, pname, &val)
			return &tender.Float{Value: float64(val)}, nil
		},
	},

	"get_tex_image": &tender.NativeFunction{
		Name: "get_tex_image",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 5 {
				return nil, tender.ErrInvalidArgCount
			}
			target, _ := tender.ToUint32(args[0])
			level, _ := tender.ToInt32(args[1])
			format, _ := tender.ToUint32(args[2])
			typ, _ := tender.ToUint32(args[3])
			bytes, ok := tender.ToByteSlice(args[4])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			gl.GetTexImage(target, level, format, typ, unsafe.Pointer(&bytes[0]))
			return tender.NullValue, nil
		},
	},
	
	// ==================== Shader Constants ====================
	"FRAGMENT_SHADER": &tender.Int{Value: int64(gl.FRAGMENT_SHADER)},
	"VERTEX_SHADER":   &tender.Int{Value: int64(gl.VERTEX_SHADER)},
	"COMPILE_STATUS":  &tender.Int{Value: int64(gl.COMPILE_STATUS)},
	"LINK_STATUS":     &tender.Int{Value: int64(gl.LINK_STATUS)},

	// ==================== Shader Lifecycle ====================
	"create_shader": &tender.NativeFunction{
		Name: "create_shader",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			shaderType, ok := tender.ToUint32(args[0])
			if !ok { return nil, tender.ErrInvalidArgument }
			shader := gl.CreateShader(shaderType)
			return &tender.Int{Value: int64(shader)}, nil
		},
	},

	"shader_source": &tender.NativeFunction{
		Name: "shader_source",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 { return nil, tender.ErrInvalidArgCount }
			shader, okS := tender.ToUint32(args[0])
			source, okSrc := args[1].(*tender.String)
			if !okS || !okSrc { return nil, tender.ErrInvalidArgument }
			
			cstrs, free := gl.Strs(source.Value + "\x00")
			defer free()
			gl.ShaderSource(shader, 1, cstrs, nil)
			return tender.NullValue, nil
		},
	},

	"compile_shader": &tender.NativeFunction{
		Name: "compile_shader",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			shader, _ := tender.ToUint32(args[0])
			gl.CompileShader(shader)
			return tender.NullValue, nil
		},
	},

	"create_program": &tender.NativeFunction{
		Name: "create_program",
		Value: func(args ...tender.Object) (tender.Object, error) {
			prog := gl.CreateProgram()
			return &tender.Int{Value: int64(prog)}, nil
		},
	},

	"attach_shader": &tender.NativeFunction{
		Name: "attach_shader",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 { return nil, tender.ErrInvalidArgCount }
			prog, _ := tender.ToUint32(args[0])
			shader, _ := tender.ToUint32(args[1])
			gl.AttachShader(prog, shader)
			return tender.NullValue, nil
		},
	},

	"link_program": &tender.NativeFunction{
		Name: "link_program",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			prog, _ := tender.ToUint32(args[0])
			gl.LinkProgram(prog)
			return tender.NullValue, nil
		},
	},

	"use_program": &tender.NativeFunction{
		Name: "use_program",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			prog, _ := tender.ToUint32(args[0])
			gl.UseProgram(prog)
			return tender.NullValue, nil
		},
	},

	"delete_shader": &tender.NativeFunction{
		Name: "delete_shader",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			shader, _ := tender.ToUint32(args[0])
			gl.DeleteShader(shader)
			return tender.NullValue, nil
		},
	},

	"delete_program": &tender.NativeFunction{
		Name: "delete_program",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			program, _ := tender.ToUint32(args[0])
			gl.DeleteProgram(program)
			return tender.NullValue, nil
		},
	},

	// ==================== Uniforms & Attributes ====================
	"get_uniform_location": &tender.NativeFunction{
		Name: "get_uniform_location",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 { return nil, tender.ErrInvalidArgCount }
			prog, _ := tender.ToUint32(args[0])
			name, _ := args[1].(*tender.String)
			loc := gl.GetUniformLocation(prog, gl.Str(name.Value+"\x00"))
			return &tender.Int{Value: int64(loc)}, nil
		},
	},

	"get_attrib_location": &tender.NativeFunction{
		Name: "get_attrib_location",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 { return nil, tender.ErrInvalidArgCount }
			program, _ := tender.ToUint32(args[0])
			name, _ := args[1].(*tender.String)
			loc := gl.GetAttribLocation(program, gl.Str(name.Value+"\x00"))
			return &tender.Int{Value: int64(loc)}, nil
		},
	},

	"uniform1f": &tender.NativeFunction{
		Name: "uniform1f",
		Value: func(args ...tender.Object) (tender.Object, error) {
			loc, _ := tender.ToInt32(args[0])
			val, _ := tender.ToFloat32(args[1])
			gl.Uniform1f(loc, val)
			return tender.NullValue, nil
		},
	},
	
	// ==================== Modern OpenGL Constants ====================
	"ARRAY_BUFFER":         &tender.Int{Value: int64(gl.ARRAY_BUFFER)},
	"STATIC_DRAW":          &tender.Int{Value: int64(gl.STATIC_DRAW)},
	"DYNAMIC_DRAW":         &tender.Int{Value: int64(gl.DYNAMIC_DRAW)},
	"FRAMEBUFFER":          &tender.Int{Value: int64(gl.FRAMEBUFFER)},
	"COLOR_ATTACHMENT0":    &tender.Int{Value: int64(gl.COLOR_ATTACHMENT0)},
	"FRAMEBUFFER_COMPLETE": &tender.Int{Value: int64(gl.FRAMEBUFFER_COMPLETE)},

	// ==================== Vertex Arrays (VAO) & Buffers (VBO) ====================
	"gen_vertex_arrays": &tender.NativeFunction{
		Name: "gen_vertex_arrays",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			n, _ := tender.ToInt32(args[0])
			vaos := make([]uint32, n)
			gl.GenVertexArrays(n, &vaos[0])
			
			arr := make([]tender.Object, n)
			for i, v := range vaos { arr[i] = &tender.Int{Value: int64(v)} }
			return &tender.Array{Value: arr}, nil
		},
	},

	"bind_vertex_array": &tender.NativeFunction{
		Name: "bind_vertex_array",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			array, _ := tender.ToUint32(args[0])
			gl.BindVertexArray(array)
			return tender.NullValue, nil
		},
	},

	"gen_buffers": &tender.NativeFunction{
		Name: "gen_buffers",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			n, _ := tender.ToInt32(args[0])
			buffers := make([]uint32, n)
			gl.GenBuffers(n, &buffers[0])
			
			arr := make([]tender.Object, n)
			for i, b := range buffers { arr[i] = &tender.Int{Value: int64(b)} }
			return &tender.Array{Value: arr}, nil
		},
	},

	"bind_buffer": &tender.NativeFunction{
		Name: "bind_buffer",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 { return nil, tender.ErrInvalidArgCount }
			target, _ := tender.ToUint32(args[0])
			buffer, _ := tender.ToUint32(args[1])
			gl.BindBuffer(target, buffer)
			return tender.NullValue, nil
		},
	},
	
	"delete_buffers": &tender.NativeFunction{
		Name: "delete_buffers",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			if arr, ok := args[0].(*tender.Array); ok {
				buffers := make([]uint32, len(arr.Value))
				for i, obj := range arr.Value {
					b, _ := tender.ToUint32(obj)
					buffers[i] = b
				}
				gl.DeleteBuffers(int32(len(buffers)), &buffers[0])
			} else if b, ok := tender.ToUint32(args[0]); ok {
				gl.DeleteBuffers(1, &b)
			}
			return tender.NullValue, nil
		},
	},

	"delete_vertex_arrays": &tender.NativeFunction{
		Name: "delete_vertex_arrays",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			if arr, ok := args[0].(*tender.Array); ok {
				vaos := make([]uint32, len(arr.Value))
				for i, obj := range arr.Value {
					v, _ := tender.ToUint32(obj)
					vaos[i] = v
				}
				gl.DeleteVertexArrays(int32(len(vaos)), &vaos[0])
			} else if v, ok := tender.ToUint32(args[0]); ok {
				gl.DeleteVertexArrays(1, &v)
			}
			return tender.NullValue, nil
		},
	},

	"buffer_data": &tender.NativeFunction{
		Name: "buffer_data",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 { return nil, tender.ErrInvalidArgCount }
			
			target, _ := tender.ToUint32(args[0])
			dataArray, ok := args[1].(*tender.Array)
			usage, _ := tender.ToUint32(args[2])
			
			if !ok || args[1] == tender.NullValue {
				gl.BufferData(target, 0, nil, usage)
				return tender.NullValue, nil
			}
			
			// Pack Tender high-level numbers into flat float32 binary data
			floats := make([]float32, len(dataArray.Value))
			for i, val := range dataArray.Value {
				if f, ok := val.(*tender.Float); ok {
					floats[i] = float32(f.Value)
				} else if integrity, ok := val.(*tender.Int); ok {
					floats[i] = float32(integrity.Value)
				}
			}
			
			size := len(floats) * 4 // 4 bytes per float32
			gl.BufferData(target, size, unsafe.Pointer(&floats[0]), usage)
			return tender.NullValue, nil
		},
	},

	"enable_vertex_attrib_array": &tender.NativeFunction{
		Name: "enable_vertex_attrib_array",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			index, _ := tender.ToUint32(args[0])
			gl.EnableVertexAttribArray(index)
			return tender.NullValue, nil
		},
	},

	"vertex_attrib_pointer": &tender.NativeFunction{
		Name: "vertex_attrib_pointer",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 6 { return nil, tender.ErrInvalidArgCount }
			index, _ := tender.ToUint32(args[0])
			size, _ := tender.ToInt32(args[1])
			xtype, _ := tender.ToUint32(args[2])
			normalized, _ := tender.ToBool(args[3])
			stride, _ := tender.ToInt32(args[4])
			offset, _ := tender.ToInt(args[5])
			
			gl.VertexAttribPointer(index, size, xtype, normalized, stride, unsafe.Pointer(uintptr(offset)))
			return tender.NullValue, nil
		},
	},

	// ==================== Framebuffers (FBO) ====================
	"gen_framebuffers": &tender.NativeFunction{
		Name: "gen_framebuffers",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			n, _ := tender.ToInt32(args[0])
			fbos := make([]uint32, n)
			gl.GenFramebuffers(n, &fbos[0])
			
			arr := make([]tender.Object, n)
			for i, f := range fbos { arr[i] = &tender.Int{Value: int64(f)} }
			return &tender.Array{Value: arr}, nil
		},
	},

	"bind_framebuffer": &tender.NativeFunction{
		Name: "bind_framebuffer",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 { return nil, tender.ErrInvalidArgCount }
			target, _ := tender.ToUint32(args[0])
			framebuffer, _ := tender.ToUint32(args[1])
			gl.BindFramebuffer(target, framebuffer)
			return tender.NullValue, nil
		},
	},

	"framebuffer_texture2d": &tender.NativeFunction{
		Name: "framebuffer_texture2d",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 5 { return nil, tender.ErrInvalidArgCount }
			target, _ := tender.ToUint32(args[0])
			attachment, _ := tender.ToUint32(args[1])
			textarget, _ := tender.ToUint32(args[2])
			texture, _ := tender.ToUint32(args[3])
			level, _ := tender.ToInt32(args[4])
			
			gl.FramebufferTexture2D(target, attachment, textarget, texture, level)
			return tender.NullValue, nil
		},
	},

	// ==================== Advanced UI Blending ====================
	"blend_func_separate": &tender.NativeFunction{
		Name: "blend_func_separate",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 { return nil, tender.ErrInvalidArgCount }
			sRGB, _ := tender.ToUint32(args[0])
			dRGB, _ := tender.ToUint32(args[1])
			sAlpha, _ := tender.ToUint32(args[2])
			dAlpha, _ := tender.ToUint32(args[3])
			gl.BlendFuncSeparate(sRGB, dRGB, sAlpha, dAlpha)
			return tender.NullValue, nil
		},
	},

	// ==================== Shader Uniforms Expansion ====================
	"uniform1i": &tender.NativeFunction{
		Name: "uniform1i",
		Value: func(args ...tender.Object) (tender.Object, error) {
			loc, _ := tender.ToInt32(args[0])
			val, _ := tender.ToInt32(args[1])
			gl.Uniform1i(loc, val)
			return tender.NullValue, nil
		},
	},

	"uniform2f": &tender.NativeFunction{
		Name: "uniform2f",
		Value: func(args ...tender.Object) (tender.Object, error) {
			loc, _ := tender.ToInt32(args[0])
			v0, _ := tender.ToFloat32(args[1])
			v1, _ := tender.ToFloat32(args[2])
			gl.Uniform2f(loc, v0, v1)
			return tender.NullValue, nil
		},
	},
	
	"uniform3f": &tender.NativeFunction{
		Name: "uniform3f",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 { return nil, tender.ErrInvalidArgCount }
			loc, _ := tender.ToInt32(args[0])
			v0, _ := tender.ToFloat32(args[1])
			v1, _ := tender.ToFloat32(args[2])
			v2, _ := tender.ToFloat32(args[3])
			gl.Uniform3f(loc, v0, v1, v2)
			return tender.NullValue, nil
		},
	},

	"uniform4f": &tender.NativeFunction{
		Name: "uniform4f",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 5 { return nil, tender.ErrInvalidArgCount }
			loc, _ := tender.ToInt32(args[0])
			v0, _ := tender.ToFloat32(args[1])
			v1, _ := tender.ToFloat32(args[2])
			v2, _ := tender.ToFloat32(args[3])
			v3, _ := tender.ToFloat32(args[4])
			gl.Uniform4f(loc, v0, v1, v2, v3)
			return tender.NullValue, nil
		},
	},

	"uniform_matrix4fv": &tender.NativeFunction{
		Name: "uniform_matrix4fv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 { return nil, tender.ErrInvalidArgCount }
			loc, _ := tender.ToInt32(args[0])
			count, _ := tender.ToInt32(args[1])
			transpose, _ := tender.ToBool(args[2])

			arr, ok := args[3].(*tender.Array)
			if !ok || len(arr.Value) != 16 { return nil, tender.ErrInvalidArgCount }
			floats := make([]float32, 16)
			for idx, val := range arr.Value {
				if f, ok := val.(*tender.Float); ok {
					floats[idx] = float32(f.Value)
				} else if intVal, ok := val.(*tender.Int); ok {
					floats[idx] = float32(intVal.Value)
				} else {
					return nil, tender.ErrInvalidArgCount
				}
			}
			gl.UniformMatrix4fv(loc, count, transpose, &floats[0])
			return tender.NullValue, nil
		},
	},
		
	// ==================== Shader Compilation Diagnostics ====================
	"get_shader_iv": &tender.NativeFunction{
		Name: "get_shader_iv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 { return nil, tender.ErrInvalidArgCount }
			shader, _ := tender.ToUint32(args[0])
			pname, _ := tender.ToUint32(args[1])
			var param int32
			gl.GetShaderiv(shader, pname, &param)
			return &tender.Int{Value: int64(param)}, nil
		},
	},

	"get_shader_info_log": &tender.NativeFunction{
		Name: "get_shader_info_log",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			shader, _ := tender.ToUint32(args[0])
			var logLength int32
			gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
			if logLength == 0 { return &tender.String{Value: ""}, nil }
			
			logBytes := make([]byte, logLength)
			gl.GetShaderInfoLog(shader, logLength, nil, &logBytes[0])
			return &tender.String{Value: string(logBytes)}, nil
		},
	},
	
	"get_program_iv": &tender.NativeFunction{
		Name: "get_program_iv",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 { return nil, tender.ErrInvalidArgCount }
			program, _ := tender.ToUint32(args[0])
			pname, _ := tender.ToUint32(args[1])
			var param int32
			gl.GetProgramiv(program, pname, &param)
			return &tender.Int{Value: int64(param)}, nil
		},
	},

	"get_program_info_log": &tender.NativeFunction{
		Name: "get_program_info_log",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			program, _ := tender.ToUint32(args[0])
			var logLength int32
			gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
			if logLength == 0 { return &tender.String{Value: ""}, nil }
			
			logBytes := make([]byte, logLength)
			gl.GetProgramInfoLog(program, logLength, nil, &logBytes[0])
			return &tender.String{Value: string(logBytes)}, nil
		},
	},
	
	// ==================== Optimized OBJ Loader & Drawer ====================
	"load_obj": &tender.NativeFunction{
		Name: "load_obj",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			path, ok := args[0].(*tender.String)
			if !ok { return nil, tender.ErrInvalidArgument }
			
			mesh, err := LoadOBJ(tender.ResolvePath(path.Value))
			if err != nil { return nil, err }
			
			// Passing nil for materials compiles the mesh exactly as it did before
			return &tender.Int{Value: int64(compileMeshToDisplayList(mesh, nil))}, nil
		},
	},

	"load_obj_with_mtl": &tender.NativeFunction{
		Name: "load_obj_with_mtl",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 { return nil, tender.ErrInvalidArgCount }
			
			objPath, ok1 := args[0].(*tender.String)
			mtlPath, ok2 := args[1].(*tender.String)
			if !ok1 || !ok2 { return nil, tender.ErrInvalidArgument }

			resolvedObjPath := tender.ResolvePath(objPath.Value)
			resolvedMtlPath := tender.ResolvePath(mtlPath.Value)
			
			materials, err := LoadMTL(resolvedMtlPath)
			if err != nil { return nil, err }

			mtlDir := filepath.Dir(resolvedMtlPath)
			for _, mat := range materials {
				if mat.DiffuseMap != "" {
					resolvedTexPath := tender.ResolvePathFromDir(mat.DiffuseMap, mtlDir)
					texID, err := internalLoadTexture(resolvedTexPath)
					if err == nil {
						mat.TextureID = texID
					}
				}
			}
			
			mesh, err := LoadOBJ(resolvedObjPath)
			if err != nil { return nil, err }
			
			displayList := compileMeshToDisplayList(mesh, materials)
			return &tender.Int{Value: int64(displayList)}, nil
		},
	},


	"parse_obj": &tender.NativeFunction{
		Name: "parse_obj",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			strData, ok := args[0].(*tender.String)
			if !ok { return nil, tender.ErrInvalidArgument }
			
			mesh, err := ParseOBJ([]byte(strData.Value))
			if err != nil { return nil, err }
			
			return &tender.Int{Value: int64(compileMeshToDisplayList(mesh, nil))}, nil
		},
	},

	"load_texture": &tender.NativeFunction{
		Name: "load_texture",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 { return nil, tender.ErrInvalidArgCount }
			
			path, ok := args[0].(*tender.String)
			if !ok { return nil, tender.ErrInvalidArgument }
			
			texID, err := internalLoadTexture(tender.ResolvePath(path.Value))
			if err != nil { return nil, err }
			
			return &tender.Int{Value: int64(texID)}, nil
		},
	},
	
	"draw_obj": &tender.NativeFunction{
		Name: "draw_obj",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			listID, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			
			// Blistering fast hardware-accelerated draw call
			gl.CallList(uint32(listID.Value))
			
			return tender.NullValue, nil
		},
	},

}

// internalLoadTexture skips script reflection and returns the raw uint32 texture ID.
func internalLoadTexture(path string) (uint32, error) {
	file, err := os.Open(path)
	if err != nil { return 0, err }
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil { return 0, err }

	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	rgba := make([]byte, w*h*4)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// Respect image bounds offsets and flip vertically for OpenGL
			srcX := bounds.Min.X + x
			srcY := bounds.Max.Y - 1 - y
			
			// NRGBAModel prevents pre-multiplied alpha from darkening your colors
			c := color.NRGBAModel.Convert(img.At(srcX, srcY)).(color.NRGBA)
			
			idx := (y*w + x) * 4
			rgba[idx]   = c.R
			rgba[idx+1] = c.G
			rgba[idx+2] = c.B
			rgba[idx+3] = c.A
		}
	}

	var texID uint32
	gl.GenTextures(1, &texID)
	gl.BindTexture(gl.TEXTURE_2D, texID)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&rgba[0]))

	return texID, nil
}