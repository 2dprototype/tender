//go:build glu

package stdlib

import (
	"unsafe"

	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/v/glu"
)

var gluModule = map[string]tender.Object{
	// ==================== Constants ====================
	// TessCallback constants
	"TESS_BEGIN_DATA":     &tender.Int{Value: int64(glu.TESS_BEGIN_DATA)},
	"TESS_VERTEX_DATA":    &tender.Int{Value: int64(glu.TESS_VERTEX_DATA)},
	"TESS_END_DATA":       &tender.Int{Value: int64(glu.TESS_END_DATA)},
	"TESS_ERROR_DATA":     &tender.Int{Value: int64(glu.TESS_ERROR_DATA)},
	"TESS_EDGE_FLAG_DATA": &tender.Int{Value: int64(glu.TESS_EDGE_FLAG_DATA)},
	"TESS_COMBINE_DATA":   &tender.Int{Value: int64(glu.TESS_COMBINE_DATA)},

	// TessProperty constants
	"TESS_WINDING_RULE":  &tender.Int{Value: int64(glu.TESS_WINDING_RULE)},
	"TESS_BOUNDARY_ONLY": &tender.Int{Value: int64(glu.TESS_BOUNDARY_ONLY)},
	"TESS_TOLERANCE":     &tender.Int{Value: int64(glu.TESS_TOLERANCE)},

	// TessWinding constants
	"TESS_WINDING_ODD":         &tender.Int{Value: int64(glu.TESS_WINDING_ODD)},
	"TESS_WINDING_NONZERO":     &tender.Int{Value: int64(glu.TESS_WINDING_NONZERO)},
	"TESS_WINDING_POSITIVE":    &tender.Int{Value: int64(glu.TESS_WINDING_POSITIVE)},
	"TESS_WINDING_NEGATIVE":    &tender.Int{Value: int64(glu.TESS_WINDING_NEGATIVE)},
	"TESS_WINDING_ABS_GEQ_TWO": &tender.Int{Value: int64(glu.TESS_WINDING_ABS_GEQ_TWO)},

	// ==================== Core GLU Functions ====================

	"build_2d_mipmaps": &tender.BuiltinFunction{
		Name: "build_2d_mipmaps",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 7 {
				return nil, tender.ErrInvalidArgCount
			}
			target, okT := tender.ToUint32(args[0])
			if !okT {
				return nil, tender.ErrInvalidArgCount
			}
			internalFormat, okIF := tender.ToInt(args[1])
			if !okIF {
				return nil, tender.ErrInvalidArgCount
			}
			width, okW := tender.ToInt(args[2])
			if !okW {
				return nil, tender.ErrInvalidArgCount
			}
			height, okH := tender.ToInt(args[3])
			if !okH {
				return nil, tender.ErrInvalidArgCount
			}
			format, okF := tender.ToUint32(args[4])
			if !okF {
				return nil, tender.ErrInvalidArgCount
			}
			typ, okTy := tender.ToUint32(args[5])
			if !okTy {
				return nil, tender.ErrInvalidArgCount
			}
			var data interface{}
			if args[6] != tender.NullValue {
				bytes, ok := tender.ToByteSlice(args[6])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				data = bytes
			}
			result := glu.Build2DMipmaps(target, internalFormat, width, height, format, typ, data)
			return &tender.Int{Value: int64(result)}, nil
		},
	},

	"look_at": &tender.BuiltinFunction{
		Name: "look_at",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 9 {
				return nil, tender.ErrInvalidArgCount
			}
			eyeX, ok1 := tender.ToFloat64(args[0])
			eyeY, ok2 := tender.ToFloat64(args[1])
			eyeZ, ok3 := tender.ToFloat64(args[2])
			centerX, ok4 := tender.ToFloat64(args[3])
			centerY, ok5 := tender.ToFloat64(args[4])
			centerZ, ok6 := tender.ToFloat64(args[5])
			upX, ok7 := tender.ToFloat64(args[6])
			upY, ok8 := tender.ToFloat64(args[7])
			upZ, ok9 := tender.ToFloat64(args[8])
			if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 || !ok6 || !ok7 || !ok8 || !ok9 {
				return nil, tender.ErrInvalidArgCount
			}
			glu.LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ)
			return tender.NullValue, nil
		},
	},

	"perspective": &tender.BuiltinFunction{
		Name: "perspective",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			fovy, ok1 := tender.ToFloat64(args[0])
			aspect, ok2 := tender.ToFloat64(args[1])
			zNear, ok3 := tender.ToFloat64(args[2])
			zFar, ok4 := tender.ToFloat64(args[3])
			if !ok1 || !ok2 || !ok3 || !ok4 {
				return nil, tender.ErrInvalidArgCount
			}
			glu.Perspective(fovy, aspect, zNear, zFar)
			return tender.NullValue, nil
		},
	},

	"project": &tender.BuiltinFunction{
		Name: "project",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 38 {
				return nil, tender.ErrInvalidArgCount
			}
			projX, ok1 := tender.ToFloat64(args[0])
			projY, ok2 := tender.ToFloat64(args[1])
			projZ, ok3 := tender.ToFloat64(args[2])
			if !ok1 || !ok2 || !ok3 {
				return nil, tender.ErrInvalidArgCount
			}
			var model [16]float64
			for i := 0; i < 16; i++ {
				val, ok := tender.ToFloat64(args[3+i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				model[i] = val
			}
			var proj [16]float64
			for i := 0; i < 16; i++ {
				val, ok := tender.ToFloat64(args[19+i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				proj[i] = val
			}
			var view [4]int32
			for i := 0; i < 4; i++ {
				val, ok := tender.ToInt32(args[35+i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				view[i] = val
			}
			winX, winY, winZ := glu.Project(projX, projY, projZ, &model, &proj, &view)
			return &tender.Array{
				Value: []tender.Object{
					&tender.Float{Value: winX},
					&tender.Float{Value: winY},
					&tender.Float{Value: winZ},
				},
			}, nil
		},
	},

	"unproject": &tender.BuiltinFunction{
		Name: "unproject",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 38 {
				return nil, tender.ErrInvalidArgCount
			}
			winX, ok1 := tender.ToFloat64(args[0])
			winY, ok2 := tender.ToFloat64(args[1])
			winZ, ok3 := tender.ToFloat64(args[2])
			if !ok1 || !ok2 || !ok3 {
				return nil, tender.ErrInvalidArgCount
			}
			var model [16]float64
			for i := 0; i < 16; i++ {
				val, ok := tender.ToFloat64(args[3+i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				model[i] = val
			}
			var proj [16]float64
			for i := 0; i < 16; i++ {
				val, ok := tender.ToFloat64(args[19+i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				proj[i] = val
			}
			var view [4]int32
			for i := 0; i < 4; i++ {
				val, ok := tender.ToInt32(args[35+i])
				if !ok {
					return nil, tender.ErrInvalidArgCount
				}
				view[i] = val
			}
			objX, objY, objZ := glu.UnProject(winX, winY, winZ, &model, &proj, &view)
			return &tender.Array{
				Value: []tender.Object{
					&tender.Float{Value: objX},
					&tender.Float{Value: objY},
					&tender.Float{Value: objZ},
				},
			}, nil
		},
	},

	"error_string": &tender.BuiltinFunction{
		Name: "error_string",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			errCode, ok := tender.ToUint32(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgCount
			}
			str, err := glu.ErrorString(errCode)
			if err != nil {
				return wrapError(err), nil
			}
			return &tender.String{Value: str}, nil
		},
	},

	// ==================== Quadric Functions ====================

	"new_quadric": &tender.BuiltinFunction{
		Name: "new_quadric",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			q := glu.NewQuadric()
			return &tender.Int{Value: int64(uintptr(q))}, nil
		},
	},

	"sphere": &tender.BuiltinFunction{
		Name: "sphere",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			qPtr, okQ := tender.ToUint64(args[0])
			if !okQ {
				return nil, tender.ErrInvalidArgCount
			}
			radius, okR := tender.ToFloat32(args[1])
			if !okR {
				return nil, tender.ErrInvalidArgCount
			}
			slices, okS := tender.ToInt(args[2])
			if !okS {
				return nil, tender.ErrInvalidArgCount
			}
			stacks, okSt := tender.ToInt(args[3])
			if !okSt {
				return nil, tender.ErrInvalidArgCount
			}
			q := unsafe.Pointer(uintptr(qPtr))
			glu.Sphere(q, radius, slices, stacks)
			return tender.NullValue, nil
		},
	},

	"cylinder": &tender.BuiltinFunction{
		Name: "cylinder",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 6 {
				return nil, tender.ErrInvalidArgCount
			}
			qPtr, okQ := tender.ToUint64(args[0])
			if !okQ {
				return nil, tender.ErrInvalidArgCount
			}
			base, okB := tender.ToFloat32(args[1])
			if !okB {
				return nil, tender.ErrInvalidArgCount
			}
			top, okT := tender.ToFloat32(args[2])
			if !okT {
				return nil, tender.ErrInvalidArgCount
			}
			height, okH := tender.ToFloat32(args[3])
			if !okH {
				return nil, tender.ErrInvalidArgCount
			}
			slices, okS := tender.ToInt(args[4])
			if !okS {
				return nil, tender.ErrInvalidArgCount
			}
			stacks, okSt := tender.ToInt(args[5])
			if !okSt {
				return nil, tender.ErrInvalidArgCount
			}
			q := unsafe.Pointer(uintptr(qPtr))
			glu.Cylinder(q, base, top, height, slices, stacks)
			return tender.NullValue, nil
		},
	},

	"disk": &tender.BuiltinFunction{
		Name: "disk",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 5 {
				return nil, tender.ErrInvalidArgCount
			}
			qPtr, okQ := tender.ToUint64(args[0])
			if !okQ {
				return nil, tender.ErrInvalidArgCount
			}
			inner, okI := tender.ToFloat32(args[1])
			if !okI {
				return nil, tender.ErrInvalidArgCount
			}
			outer, okO := tender.ToFloat32(args[2])
			if !okO {
				return nil, tender.ErrInvalidArgCount
			}
			slices, okS := tender.ToInt(args[3])
			if !okS {
				return nil, tender.ErrInvalidArgCount
			}
			loops, okL := tender.ToInt(args[4])
			if !okL {
				return nil, tender.ErrInvalidArgCount
			}
			q := unsafe.Pointer(uintptr(qPtr))
			glu.Disk(q, inner, outer, slices, loops)
			return tender.NullValue, nil
		},
	},

	"partial_disk": &tender.BuiltinFunction{
		Name: "partial_disk",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 7 {
				return nil, tender.ErrInvalidArgCount
			}
			qPtr, okQ := tender.ToUint64(args[0])
			if !okQ {
				return nil, tender.ErrInvalidArgCount
			}
			inner, okI := tender.ToFloat32(args[1])
			if !okI {
				return nil, tender.ErrInvalidArgCount
			}
			outer, okO := tender.ToFloat32(args[2])
			if !okO {
				return nil, tender.ErrInvalidArgCount
			}
			slices, okS := tender.ToInt(args[3])
			if !okS {
				return nil, tender.ErrInvalidArgCount
			}
			loops, okL := tender.ToInt(args[4])
			if !okL {
				return nil, tender.ErrInvalidArgCount
			}
			startAngle, okSA := tender.ToFloat32(args[5])
			if !okSA {
				return nil, tender.ErrInvalidArgCount
			}
			sweepAngle, okSW := tender.ToFloat32(args[6])
			if !okSW {
				return nil, tender.ErrInvalidArgCount
			}
			q := unsafe.Pointer(uintptr(qPtr))
			glu.PartialDisk(q, inner, outer, slices, loops, startAngle, sweepAngle)
			return tender.NullValue, nil
		},
	},
}