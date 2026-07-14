package stdlib

import (
	"github.com/2dprototype/tender"
	"bytes"
	"image"
	_ "image/jpeg"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	_ "image/png"
	"github.com/2dprototype/tender/v/gg"
)

var canvasModule = map[string]tender.Object{
	"new_context": &tender.NativeFunction{Name: "new_context", Value: ggNewContext},
	"load_image": &tender.NativeFunction{Name:  "load_image", Value: imageLoad},	
	"radians": &tender.NativeFunction{Name: "radians", Value: FuncAFRF(gg.Radians)},
	"degrees": &tender.NativeFunction{Name: "degrees", Value: FuncAFRF(gg.Degrees)},
	"load_font": &tender.NativeFunction{
		Name: "load_font",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrWrongNumArguments
			}
			name, _ := tender.ToString(args[0])
			size, _ := tender.ToFloat64(args[1])
			src := args[2]

			if path, ok := tender.ToString(src); ok {
				err := gg.LoadFont(name, size, tender.ResolvePath(path))
				if err != nil {
					return wrapError(err), nil
				}
			} else if data, ok := tender.ToByteSlice(src); ok {
				err := gg.Font(name, size, data)
				if err != nil {
					return wrapError(err), nil
				}
			} else {
				return nil, tender.ErrInvalidArgumentType{Name: "src", Expected: "string or bytes"}
			}
			return &tender.Null{}, nil
		},
	},
	"load_fontdata": &tender.NativeFunction{
		Name: "load_fontdata",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrWrongNumArguments
			}
			name, _ := tender.ToString(args[0])
			size, _ := tender.ToFloat64(args[1])
			data, ok := tender.ToByteSlice(args[2])
			if !ok {
				return nil, tender.ErrInvalidArgumentType{Name: "font_data", Expected: "bytes"}
			}
			err := gg.Font(name, size, data)
			if err != nil {
				return wrapError(err), nil
			}
			return &tender.Null{}, nil
		},
	},
}


func ggNewContext(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}
	width, _ := tender.ToInt(args[0])
	height, _ := tender.ToInt(args[1])
	dc := gg.NewContext(width, height)
	return makeGGContext(dc), nil
}

func makeGGContext(ctx *gg.Context) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"drawimage": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 3 {
						return nil, tender.ErrWrongNumArguments
					}
					imageBytes, _ := tender.ToByteSlice(args[0])
					ix, _ := tender.ToInt(args[1])
					iy, _ := tender.ToInt(args[2])
					img, _, err := image.Decode(bytes.NewReader(imageBytes))
					if err != nil {
						return wrapError(err), nil
					}
					ctx.DrawImage(img, ix, iy)
					return nil, nil
				},
			},	
			"drawimage_anchored": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 5 {
						return nil, tender.ErrWrongNumArguments
					}
					imageBytes, _ := tender.ToByteSlice(args[0])
					ix, _ := tender.ToInt(args[1])
					iy, _ := tender.ToInt(args[2])	
					fx, _ := tender.ToFloat64(args[3])
					fy, _ := tender.ToFloat64(args[4])
					img, _, err := image.Decode(bytes.NewReader(imageBytes))
					if err != nil {
						return wrapError(err), nil
					}
					ctx.DrawImageAnchored(img, ix, iy, fx, fy)
					return nil, nil
				},
				},	
			"save_png": &tender.NativeFunction{
				Value: FuncASRE(ctx.SavePNG),
			},	
			"point": &tender.NativeFunction{
				Value: FuncAFFFR(ctx.DrawPoint),
			},	
			"line": &tender.NativeFunction{
				Value: FuncAFFFFR(ctx.DrawLine),
			},	
			"rect": &tender.NativeFunction{
				Value: FuncAFFFFR(ctx.DrawRectangle),
			},
			"polygon": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 5 {
						return nil, tender.ErrWrongNumArguments
					}
					i0, _ := tender.ToInt(args[0])
					f1, _ := tender.ToFloat64(args[1])
					f2, _ := tender.ToFloat64(args[2])
					f3, _ := tender.ToFloat64(args[3])
					f4, _ := tender.ToFloat64(args[4])
					ctx.DrawRegularPolygon(i0, f1, f2, f3, f4)
					return nil, nil
				},
			},	
			"roundrect": &tender.NativeFunction{
				Value: FuncAFFFFFR(ctx.DrawRoundedRectangle),
			},
			"circle": &tender.NativeFunction{
				Value: FuncAFFFR(ctx.DrawCircle),
			},	
			"arc": &tender.NativeFunction{
				Value: FuncAFFFFFR(ctx.DrawArc),
			},
			"ellipse": &tender.NativeFunction{
				Value: FuncAFFFFR(ctx.DrawEllipse),
			},
			"ellipsearc": &tender.NativeFunction{
				Value: FuncAFFFFFFR(ctx.DrawEllipticalArc),
			},
			"set_pixel": &tender.NativeFunction{
				Name:  "set_pixel",
				Value: FuncAIIR(ctx.SetPixel),
			},	
			"rgb": &tender.NativeFunction{
				Value: FuncAFFFR(ctx.SetRGB),
			},
			"rgba": &tender.NativeFunction{
				Value: FuncAFFFFR(ctx.SetRGBA),
			},	
			"rgba255": &tender.NativeFunction{
				Value: FuncAIIIIR(ctx.SetRGBA255),
			},	
			"rgb255": &tender.NativeFunction{
				Value: FuncAIIIR(ctx.SetRGB255),
			},
			"hex": &tender.NativeFunction{
				Value: FuncASR(ctx.SetHexColor),
			},
			"linewidth": &tender.NativeFunction{
				Value: FuncAFR(ctx.SetLineWidth),
			},	
			"dashoffset": &tender.NativeFunction{
				Value: FuncAFR(ctx.SetDashOffset),
			},
			"dash": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) < 1 {
						return nil, tender.ErrWrongNumArguments
					}
					elements := make([]float64, len(args))
					for i, arg := range args {
						s, _ := tender.ToFloat64(arg)
						elements[i] = s
					}
					ctx.SetDash(elements...)
					return &tender.Null{}, nil
				},
			},	
			"move_to": &tender.NativeFunction{
				Value: FuncAFFR(ctx.MoveTo),
			},	
			"line_to": &tender.NativeFunction{
				Value: FuncAFFR(ctx.LineTo),
			},	
			"quadratic_to": &tender.NativeFunction{
				Value: FuncAFFFFR(ctx.QuadraticTo),
			},	
			"cubic_to": &tender.NativeFunction{
				Value: FuncAFFFFFFR(ctx.CubicTo),
			},
			"closepath": &tender.NativeFunction{
				Value: FuncAR(ctx.ClosePath),
			},	
			"clearpath": &tender.NativeFunction{
				Value: FuncAR(ctx.ClearPath),
			},	
			"newsubpath": &tender.NativeFunction{
				Value: FuncAR(ctx.NewSubPath),
			},	
			"clear": &tender.NativeFunction{
				Value: FuncAR(ctx.Clear),
			},
			"stroke": &tender.NativeFunction{
				Value: FuncAR(ctx.Stroke),
			},	
			"fill": &tender.NativeFunction{
				Value: FuncAR(ctx.Fill),
			},		
			"stroke_preserve": &tender.NativeFunction{
				Value: FuncAR(ctx.StrokePreserve),
			},	
			"fill_preserve": &tender.NativeFunction{
				Value: FuncAR(ctx.FillPreserve),
			},	
			"text": &tender.NativeFunction{
				Value: FuncASFFR(ctx.DrawString),
			},	
			"text_anchored": &tender.NativeFunction{
				Value: FuncASFFFFR(ctx.DrawStringAnchored),
			},	
			"measure_text": &tender.NativeFunction{
				Value: FuncASRFF(ctx.MeasureString),
			},	
			"measure_multiline_text": &tender.NativeFunction{
				Value: FuncASFRFF(ctx.MeasureMultilineString),
			},	
			"load_fontface": &tender.NativeFunction{
				Value: FuncASFRE(ctx.LoadFontFace),
			},	
			"fontface": &tender.NativeFunction{
				Value: FuncAYFRE(ctx.FontFace),
			},	
			"fontheight": &tender.NativeFunction{
				Value: FuncARF(ctx.FontHeight),
			},	
			"set_font": &tender.NativeFunction{
				Value: FuncASRE(ctx.SetFont),
			},
			"identity": &tender.NativeFunction{
				Name:  "identity",
				Value: FuncAR(ctx.Identity),
			},	
			"translate": &tender.NativeFunction{
				Value: FuncAFFR(ctx.Translate),
			},	
			"scale": &tender.NativeFunction{
				Value: FuncAFFR(ctx.Scale),
			},	
			"rotate": &tender.NativeFunction{
				Value: FuncAFR(ctx.Rotate),
			},	
			"shear": &tender.NativeFunction{
				Value: FuncAFFR(ctx.Shear),
			},
			"scaleabout": &tender.NativeFunction{
				Value: FuncAFFFFR(ctx.ScaleAbout),
			},	
			"rotateabout": &tender.NativeFunction{
				Value: FuncAFFFR(ctx.RotateAbout),
			},
			"shearabout": &tender.NativeFunction{
				Value: FuncAFFFFR(ctx.ShearAbout),
			},	
			"transform_point": &tender.NativeFunction{
				Value: FuncAFFRFF(ctx.TransformPoint),
			},
			"invertmask": &tender.NativeFunction{
				Value: FuncAR(ctx.InvertMask),
			},	
			"inverty": &tender.NativeFunction{
				Value: FuncAR(ctx.InvertY),
			},	
			"push": &tender.NativeFunction{
				Value: FuncAR(ctx.Push),
			},	
			"pop": &tender.NativeFunction{
				Value: FuncAR(ctx.Pop),
			},	
			"clip": &tender.NativeFunction{
				Value: FuncAR(ctx.Clip),
			},		
			"clip_preserve": &tender.NativeFunction{
				Value: FuncAR(ctx.ClipPreserve),
			},	
			"resetclip": &tender.NativeFunction{
				Value: FuncAR(ctx.ResetClip),
			},
			"height": &tender.NativeFunction{
				Value: FuncARI(ctx.Height),
			},	
			"width": &tender.NativeFunction{
				Value: FuncARI(ctx.Width),
			},	
			"wordwrap": &tender.NativeFunction{
				Value: FuncASFRSs(ctx.WordWrap),
			},
			"image": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					return makeImage(ctx.Image()), nil
				},
			},	
		},
	}
}

