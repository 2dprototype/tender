package stdlib

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"

	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/v/gg"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

var canvasModule = map[string]tender.Object{
	"new_context": &tender.NativeFunction{Name: "new_context", Value: ggNewContext},
	"load_image":  &tender.NativeFunction{Name: "load_image", Value: imageLoad},
	"radians":     &tender.NativeFunction{Name: "radians", Value: FuncAFRF(gg.Radians)},
	"degrees":     &tender.NativeFunction{Name: "degrees", Value: FuncAFRF(gg.Degrees)},
	"load_font": &tender.NativeFunction{
		Name: "load_font",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) == 2 {
				size, _ := tender.ToFloat64(args[1])
				src := args[0]
				if path, ok := tender.ToString(src); ok {
					err := gg.LoadFont("default", size, tender.ResolvePath(path))
					if err != nil {
						return wrapError(err), nil
					}
				} else if data, ok := tender.ToByteSlice(src); ok {
					err := gg.Font("default", size, data)
					if err != nil {
						return wrapError(err), nil
					}
				} else {
					return nil, tender.ErrInvalidArgumentType{Name: "src", Expected: "string or bytes"}
				}
				return tender.NullValue, nil
			}
			if len(args) == 3 {
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
				return tender.NullValue, nil
			}
			return nil, tender.ErrWrongNumArguments
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
			return tender.NullValue, nil
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

func decodeImageArg(obj tender.Object) (image.Image, error) {
	b, err := ToFileData(obj)
	if err != nil {
		return nil, fmt.Errorf("invalid image argument: expected bytes or file path")
	}
	var img image.Image
	img, _, err = image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("invalid image format")
	}
	return img, err
}

func drawEllipticalRoundedRect(dc *gg.Context, x, y, w, h, rx, ry float64) {
	if rx <= 0 || ry <= 0 {
		dc.DrawRectangle(x, y, w, h)
		return
	}
	x0, x1, x2, x3 := x, x+rx, x+w-rx, x+w
	y0, y1, y2, y3 := y, y+ry, y+h-ry, y+h
	dc.NewSubPath()
	dc.MoveTo(x1, y0)
	dc.LineTo(x2, y0)
	dc.DrawEllipticalArc(x2, y1, rx, ry, gg.Radians(270), gg.Radians(360))
	dc.LineTo(x3, y2)
	dc.DrawEllipticalArc(x2, y2, rx, ry, gg.Radians(0), gg.Radians(90))
	dc.LineTo(x1, y3)
	dc.DrawEllipticalArc(x1, y2, rx, ry, gg.Radians(90), gg.Radians(180))
	dc.LineTo(x0, y1)
	dc.DrawEllipticalArc(x1, y1, rx, ry, gg.Radians(180), gg.Radians(270))
	dc.ClosePath()
}

func makeGGContext(ctx *gg.Context) *tender.ImmutableMap {
	var currentFontData []byte
	var currentFontPath string

	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"draw_image": &tender.NativeFunction{
				Name: "draw_image",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 3 {
						return nil, tender.ErrWrongNumArguments
					}
					img, err := decodeImageArg(args[0])
					if err != nil {
						return wrapError(err), nil
					}
					ix, _ := tender.ToInt(args[1])
					iy, _ := tender.ToInt(args[2])
					ctx.DrawImage(img, ix, iy)
					return tender.NullValue, nil
				},
			},
			"draw_image_anchored": &tender.NativeFunction{
				Name: "draw_image_anchored",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 5 {
						return nil, tender.ErrWrongNumArguments
					}
					img, err := decodeImageArg(args[0])
					if err != nil {
						return wrapError(err), nil
					}
					ix, _ := tender.ToInt(args[1])
					iy, _ := tender.ToInt(args[2])
					fx, _ := tender.ToFloat64(args[3])
					fy, _ := tender.ToFloat64(args[4])
					ctx.DrawImageAnchored(img, ix, iy, fx, fy)
					return tender.NullValue, nil
				},
			},
			"draw_image_rect": &tender.NativeFunction{
				Name: "draw_image_rect",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 9 {
						return nil, tender.ErrWrongNumArguments
					}
					img, err := decodeImageArg(args[0])
					if err != nil {
						return wrapError(err), nil
					}
					sx, _ := tender.ToInt(args[1])
					sy, _ := tender.ToInt(args[2])
					sw, _ := tender.ToInt(args[3])
					sh, _ := tender.ToInt(args[4])
					dx, _ := tender.ToFloat64(args[5])
					dy, _ := tender.ToFloat64(args[6])
					dw, _ := tender.ToFloat64(args[7])
					dh, _ := tender.ToFloat64(args[8])

					type subImager interface {
						SubImage(r image.Rectangle) image.Image
					}
					var cropped image.Image
					if si, ok := img.(subImager); ok {
						cropped = si.SubImage(image.Rect(sx, sy, sx+sw, sy+sh))
					} else {
						rgba := image.NewRGBA(image.Rect(0, 0, sw, sh))
						draw.Draw(rgba, rgba.Bounds(), img, image.Point{X: sx, Y: sy}, draw.Src)
						cropped = rgba
					}

					ctx.Push()
					ctx.Translate(dx, dy)
					if float64(sw) > 0 && float64(sh) > 0 {
						ctx.Scale(dw/float64(sw), dh/float64(sh))
					}
					ctx.DrawImage(cropped, 0, 0)
					ctx.Pop()
					return tender.NullValue, nil
				},
			},
			"save_png": &tender.NativeFunction{
				Name:  "save_png",
				Value: FuncASRE(ctx.SavePNG),
			},
			"point": &tender.NativeFunction{
				Name:  "point",
				Value: FuncAFFFR(ctx.DrawPoint),
			},
			"line": &tender.NativeFunction{
				Name:  "line",
				Value: FuncAFFFFR(ctx.DrawLine),
			},
			"rect": &tender.NativeFunction{
				Name:  "rect",
				Value: FuncAFFFFR(ctx.DrawRectangle),
			},
			"polygon": &tender.NativeFunction{
				Name: "polygon",
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
					return tender.NullValue, nil
				},
			},
			"round_rect": &tender.NativeFunction{
				Name: "round_rect",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 5 && len(args) != 6 {
						return nil, tender.ErrWrongNumArguments
					}
					x, _ := tender.ToFloat64(args[0])
					y, _ := tender.ToFloat64(args[1])
					w, _ := tender.ToFloat64(args[2])
					h, _ := tender.ToFloat64(args[3])
					rx, _ := tender.ToFloat64(args[4])
					ry := rx
					if len(args) == 6 {
						ry, _ = tender.ToFloat64(args[5])
					}
					if rx == ry {
						ctx.DrawRoundedRectangle(x, y, w, h, rx)
					} else {
						drawEllipticalRoundedRect(ctx, x, y, w, h, rx, ry)
					}
					return tender.NullValue, nil
				},
			},
			"circle": &tender.NativeFunction{
				Name:  "circle",
				Value: FuncAFFFR(ctx.DrawCircle),
			},
			"ellipse": &tender.NativeFunction{
				Name:  "ellipse",
				Value: FuncAFFFFR(ctx.DrawEllipse),
			},
			"arc": &tender.NativeFunction{
				Name:  "arc",
				Value: FuncAFFFFFR(ctx.DrawArc),
			},
			"elliptical_arc": &tender.NativeFunction{
				Name:  "elliptical_arc",
				Value: FuncAFFFFFFR(ctx.DrawEllipticalArc),
			},
			"set_pixel": &tender.NativeFunction{
				Name:  "set_pixel",
				Value: FuncAIIR(ctx.SetPixel),
			},
			"rgb": &tender.NativeFunction{
				Name:  "rgb",
				Value: FuncAFFFR(ctx.SetRGB),
			},
			"rgba": &tender.NativeFunction{
				Name:  "rgba",
				Value: FuncAFFFFR(ctx.SetRGBA),
			},
			"rgba255": &tender.NativeFunction{
				Name:  "rgba255",
				Value: FuncAIIIIR(ctx.SetRGBA255),
			},
			"rgb255": &tender.NativeFunction{
				Name:  "rgb255",
				Value: FuncAIIIR(ctx.SetRGB255),
			},
			"hex": &tender.NativeFunction{
				Name:  "hex",
				Value: FuncASR(ctx.SetHexColor),
			},
			"line_width": &tender.NativeFunction{
				Name:  "line_width",
				Value: FuncAFR(ctx.SetLineWidth),
			},
			"dashoffset": &tender.NativeFunction{
				Name:  "dashoffset",
				Value: FuncAFR(ctx.SetDashOffset),
			},
			"dash": &tender.NativeFunction{
				Name: "dash",
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
					return tender.NullValue, nil
				},
			},
			"move_to": &tender.NativeFunction{
				Name:  "move_to",
				Value: FuncAFFR(ctx.MoveTo),
			},
			"line_to": &tender.NativeFunction{
				Name:  "line_to",
				Value: FuncAFFR(ctx.LineTo),
			},
			"quadratic_to": &tender.NativeFunction{
				Name:  "quadratic_to",
				Value: FuncAFFFFR(ctx.QuadraticTo),
			},
			"cubic_to": &tender.NativeFunction{
				Name:  "cubic_to",
				Value: FuncAFFFFFFR(ctx.CubicTo),
			},
			"close_path": &tender.NativeFunction{
				Name:  "close_path",
				Value: FuncAR(ctx.ClosePath),
			},
			"clear_path": &tender.NativeFunction{
				Name:  "clear_path",
				Value: FuncAR(ctx.ClearPath),
			},
			"new_subpath": &tender.NativeFunction{
				Name:  "new_subpath",
				Value: FuncAR(ctx.NewSubPath),
			},
			"clear": &tender.NativeFunction{
				Name: "clear",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) == 0 {
						ctx.Clear()
						return tender.NullValue, nil
					}
					ctx.Push()
					defer ctx.Pop()
					if len(args) == 1 {
						if hexStr, ok := tender.ToString(args[0]); ok {
							ctx.SetHexColor(hexStr)
						}
					} else if len(args) == 3 {
						r, _ := tender.ToFloat64(args[0])
						g, _ := tender.ToFloat64(args[1])
						b, _ := tender.ToFloat64(args[2])
						ctx.SetRGB(r, g, b)
					} else if len(args) == 4 {
						r, _ := tender.ToFloat64(args[0])
						g, _ := tender.ToFloat64(args[1])
						b, _ := tender.ToFloat64(args[2])
						a, _ := tender.ToFloat64(args[3])
						ctx.SetRGBA(r, g, b, a)
					}
					ctx.Clear()
					return tender.NullValue, nil
				},
			},
			"stroke": &tender.NativeFunction{
				Name:  "stroke",
				Value: FuncAR(ctx.Stroke),
			},
			"fill": &tender.NativeFunction{
				Name:  "fill",
				Value: FuncAR(ctx.Fill),
			},
			"stroke_preserve": &tender.NativeFunction{
				Name:  "stroke_preserve",
				Value: FuncAR(ctx.StrokePreserve),
			},
			"fill_preserve": &tender.NativeFunction{
				Name:  "fill_preserve",
				Value: FuncAR(ctx.FillPreserve),
			},
			"text": &tender.NativeFunction{
				Name:  "text",
				Value: FuncASFFR(ctx.DrawString),
			},
			"text_anchored": &tender.NativeFunction{
				Name:  "text_anchored",
				Value: FuncASFFFFR(ctx.DrawStringAnchored),
			},
			"text_wrapped": &tender.NativeFunction{
				Name: "text_wrapped",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) < 4 {
						return nil, tender.ErrWrongNumArguments
					}
					text, ok := tender.ToString(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{Name: "text", Expected: "string"}
					}
					x, _ := tender.ToFloat64(args[1])
					y, _ := tender.ToFloat64(args[2])
					maxWidth, _ := tender.ToFloat64(args[3])
					lineSpacing := 1.5
					if len(args) >= 5 {
						lineSpacing, _ = tender.ToFloat64(args[4])
					}
					ctx.DrawStringWrapped(text, x, y, 0, 0, maxWidth, lineSpacing, gg.AlignLeft)
					return tender.NullValue, nil
				},
			},
			"measure_text": &tender.NativeFunction{
				Name: "measure_text",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					s, ok := tender.ToString(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{Name: "text", Expected: "string"}
					}
					w, h := ctx.MeasureString(s)
					return &tender.ImmutableMap{
						Value: map[string]tender.Object{
							"width":  &tender.Float{Value: w},
							"height": &tender.Float{Value: h},
						},
					}, nil
				},
			},
			"measure_multiline_text": &tender.NativeFunction{
				Name:  "measure_multiline_text",
				Value: FuncASFRFF(ctx.MeasureMultilineString),
			},
			"load_font": &tender.NativeFunction{
				Name: "load_font",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 {
						return nil, tender.ErrWrongNumArguments
					}
					size, _ := tender.ToFloat64(args[1])
					if path, ok := tender.ToString(args[0]); ok {
						resolved := tender.ResolvePath(path)
						err := ctx.LoadFontFace(resolved, size)
						if err != nil {
							return wrapError(err), nil
						}
						currentFontPath = resolved
						currentFontData = nil
					} else if data, ok := tender.ToByteSlice(args[0]); ok {
						err := ctx.FontFace(data, size)
						if err != nil {
							return wrapError(err), nil
						}
						currentFontData = data
						currentFontPath = ""
					} else {
						return nil, tender.ErrInvalidArgumentType{Name: "font", Expected: "string or bytes"}
					}
					return tender.NullValue, nil
				},
			},
			"set_font_size": &tender.NativeFunction{
				Name: "set_font_size",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					size, _ := tender.ToFloat64(args[0])
					if currentFontPath != "" {
						_ = ctx.LoadFontFace(currentFontPath, size)
					} else if len(currentFontData) > 0 {
						_ = ctx.FontFace(currentFontData, size)
					}
					return tender.NullValue, nil
				},
			},
			"load_fontface": &tender.NativeFunction{
				Name:  "load_fontface",
				Value: FuncASFRE(ctx.LoadFontFace),
			},
			"fontface": &tender.NativeFunction{
				Name:  "fontface",
				Value: FuncAYFRE(ctx.FontFace),
			},
			"font_height": &tender.NativeFunction{
				Name:  "font_height",
				Value: FuncARF(ctx.FontHeight),
			},
			"set_font": &tender.NativeFunction{
				Name:  "set_font",
				Value: FuncASRE(ctx.SetFont),
			},
			"identity": &tender.NativeFunction{
				Name:  "identity",
				Value: FuncAR(ctx.Identity),
			},
			"translate": &tender.NativeFunction{
				Name:  "translate",
				Value: FuncAFFR(ctx.Translate),
			},
			"scale": &tender.NativeFunction{
				Name:  "scale",
				Value: FuncAFFR(ctx.Scale),
			},
			"rotate": &tender.NativeFunction{
				Name:  "rotate",
				Value: FuncAFR(ctx.Rotate),
			},
			"shear": &tender.NativeFunction{
				Name:  "shear",
				Value: FuncAFFR(ctx.Shear),
			},
			"scale_about": &tender.NativeFunction{
				Name:  "scale_about",
				Value: FuncAFFFFR(ctx.ScaleAbout),
			},
			"rotate_about": &tender.NativeFunction{
				Name:  "rotate_about",
				Value: FuncAFFFR(ctx.RotateAbout),
			},
			"shear_about": &tender.NativeFunction{
				Name:  "shear_about",
				Value: FuncAFFFFR(ctx.ShearAbout),
			},
			"transform_point": &tender.NativeFunction{
				Name:  "transform_point",
				Value: FuncAFFRFF(ctx.TransformPoint),
			},
			"invertmask": &tender.NativeFunction{
				Name:  "invertmask",
				Value: FuncAR(ctx.InvertMask),
			},
			"inverty": &tender.NativeFunction{
				Name:  "inverty",
				Value: FuncAR(ctx.InvertY),
			},
			"push": &tender.NativeFunction{
				Name:  "push",
				Value: FuncAR(ctx.Push),
			},
			"pop": &tender.NativeFunction{
				Name:  "pop",
				Value: FuncAR(ctx.Pop),
			},
			"clip": &tender.NativeFunction{
				Name:  "clip",
				Value: FuncAR(ctx.Clip),
			},
			"clip_preserve": &tender.NativeFunction{
				Name:  "clip_preserve",
				Value: FuncAR(ctx.ClipPreserve),
			},
			"reset_clip": &tender.NativeFunction{
				Name:  "reset_clip",
				Value: FuncAR(ctx.ResetClip),
			},
			"height": &tender.NativeFunction{
				Name:  "height",
				Value: FuncARI(ctx.Height),
			},
			"width": &tender.NativeFunction{
				Name:  "width",
				Value: FuncARI(ctx.Width),
			},
			"wordwrap": &tender.NativeFunction{
				Name:  "wordwrap",
				Value: FuncASFRSs(ctx.WordWrap),
			},
			"get_image": &tender.NativeFunction{
				Name: "get_image",
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
