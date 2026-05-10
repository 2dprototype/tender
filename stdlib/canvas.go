package stdlib

import (
	"github.com/2dprototype/tender"
	"bytes"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	_ "image/png"
	"github.com/2dprototype/tender/v/gg"
	
	// "golang.org/x/mobile/event/lifecycle"
	// "golang.org/x/exp/shiny/driver"
	// "golang.org/x/exp/shiny/screen"
	// "github.com/oakmound/shiny/driver"
	// "golang.org/x/exp/shiny/driver/gldriver"
	"github.com/oakmound/shiny/driver"
	"github.com/oakmound/shiny/screen"
	// "golang.org/x/mobile/event/lifecycle"
	
	// "fmt"
)

var canvasModule = map[string]tender.Object{
	"new_context": &tender.UserFunction{Name: "new_context", Value: ggNewContext},
	"load_image": &tender.UserFunction{Name:  "load_image", Value: imageLoad},	
	"radians": &tender.UserFunction{Name: "radians", Value: FuncAFRF(gg.Radians)},
	"degrees": &tender.UserFunction{Name: "degrees", Value: FuncAFRF(gg.Degrees)},
	"load_font": &tender.UserFunction{
		Name: "load_font",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrWrongNumArguments
			}
			name, _ := tender.ToString(args[0])
			size, _ := tender.ToFloat64(args[1])
			src := args[2]

			if path, ok := tender.ToString(src); ok {
				err := gg.LoadFont(name, size, path)
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
	"load_fontdata": &tender.UserFunction{
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
	"new_window" : &tender.BuiltinFunction{
		Name: "new_window",
		NeedVMObj: true,
		Value: canvasNewWindow,
	},	
}



func canvasNewWindow(args ...tender.Object) (ret tender.Object, err error) {
	vm := args[0].(*tender.VMObj).Value
	args = args[1:] // the first arg is VMObj inserted by VM
	if !(len(args) == 4 || len(args) == 2) {
		return nil, tender.ErrWrongNumArguments
	}
	
	var wOpts screen.WindowGenerator
	
	if len(args) == 4 {
		width, _ := tender.ToInt(args[0])
		height, _ := tender.ToInt(args[1])
		title, _ := tender.ToString(args[2])
		wOpts = screen.WindowGenerator{
			Title:  title,
			Width:  width,
			Height: height,
		}
	} else {
		var width int = 400
		var height int = 400
		var title = ""
		var fullscreen = false
		var borderless = false
		var topMost = false
		var noScaling = false
		var x int32
		var y int32
		
		m, ok := args[0].(*tender.Map)
		if !ok {
			return nil, nil
		}
		if val, ok := m.Value["width"]; ok {
			width, _ = tender.ToInt(val)
		}
		if val, ok := m.Value["height"]; ok {
			height, _ = tender.ToInt(val)
		}	
		if val, ok := m.Value["title"]; ok {
			title, _ = tender.ToString(val)
		}		
		if val, ok := m.Value["fullscreen"]; ok {
			fullscreen, _ = tender.ToBool(val)
		}	
		if val, ok := m.Value["borderless"]; ok {
			borderless, _ = tender.ToBool(val)
		}	
		if val, ok := m.Value["top_most"]; ok {
			topMost, _ = tender.ToBool(val)
		}
		if val, ok := m.Value["no_scaling"]; ok {
			noScaling, _ = tender.ToBool(val)
		}
		if val, ok := m.Value["x"]; ok {
			x, _ = tender.ToInt32(val)
		}	
		if val, ok := m.Value["y"]; ok {
			y, _ = tender.ToInt32(val)
		}
		wOpts = screen.WindowGenerator{
			Title:  title,
			Width:  width,
			Height: height,
			Fullscreen: fullscreen,
			Borderless: borderless,
			TopMost: topMost,
			NoScaling: noScaling,
			X: x,
			Y: y,
		}
	}
	// Use the Go exp/shiny package to create a window
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(wOpts)
		
		if err != nil {
			return
		}
		
		var ctx *gg.Context
		
		wmap := &tender.ImmutableMap{
			Value: map[string]tender.Object{
				"release": &tender.UserFunction{Value: FuncAR(w.Release)},
				// "wooh": &tender.UserFunction{
					// Value: func(args ...tender.Object) (tender.Object, error) {
						// screen.Title("My Shiny Window")
						// return nil, nil
					// },
				// },		
				"new_context": &tender.UserFunction{
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 2 {
							return nil, tender.ErrWrongNumArguments
						}
						ww, _ := tender.ToInt(args[0])
						hh, _ := tender.ToInt(args[1])
						ctx = gg.NewContext(ww, hh)
						return makeGGContext(ctx), nil
					},
				},	
				"update": &tender.UserFunction{
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 2 {
							return nil, tender.ErrWrongNumArguments
						}
						ww, _ := tender.ToInt(args[0])
						hh, _ := tender.ToInt(args[1])
						screenImage, err := s.NewImage(image.Point{X: ww, Y: hh})
						if err != nil {
							return wrapError(err), nil
						}
						defer screenImage.Release()
						draw.Draw(screenImage.RGBA(), screenImage.Bounds(), ctx.Image().(*image.RGBA), image.Point{}, draw.Over)
						w.Upload(image.Point{0, 0}, screenImage, screenImage.Bounds())
						w.Publish()
						return nil, nil
					},
				},	
				"next_event": &tender.UserFunction{
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 0 {
							return nil, tender.ErrWrongNumArguments
						}
						return eventToObject(w.NextEvent()), nil
					},
				},
			},
		}
		
		defer w.Release()
		
		if len(args) == 4 {
			tender.WrapFuncCall(vm, args[3], wmap)
		} else {
			tender.WrapFuncCall(vm, args[1], wmap)	
		}
	})
	
	return nil, nil
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
			"drawimage": &tender.UserFunction{
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
			"drawimage_anchored": &tender.UserFunction{
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
			"save_png": &tender.UserFunction{
				Value: FuncASRE(ctx.SavePNG),
			},	
			"point": &tender.UserFunction{
				Value: FuncAFFFR(ctx.DrawPoint),
			},	
			"line": &tender.UserFunction{
				Value: FuncAFFFFR(ctx.DrawLine),
			},	
			"rect": &tender.UserFunction{
				Value: FuncAFFFFR(ctx.DrawRectangle),
			},
			"polygon": &tender.UserFunction{
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
			"roundrect": &tender.UserFunction{
				Value: FuncAFFFFFR(ctx.DrawRoundedRectangle),
			},
			"circle": &tender.UserFunction{
				Value: FuncAFFFR(ctx.DrawCircle),
			},	
			"arc": &tender.UserFunction{
				Value: FuncAFFFFFR(ctx.DrawArc),
			},
			"ellipse": &tender.UserFunction{
				Value: FuncAFFFFR(ctx.DrawEllipse),
			},
			"ellipsearc": &tender.UserFunction{
				Value: FuncAFFFFFFR(ctx.DrawEllipticalArc),
			},
			"set_pixel": &tender.UserFunction{
				Name:  "set_pixel",
				Value: FuncAIIR(ctx.SetPixel),
			},	
			"rgb": &tender.UserFunction{
				Value: FuncAFFFR(ctx.SetRGB),
			},
			"rgba": &tender.UserFunction{
				Value: FuncAFFFFR(ctx.SetRGBA),
			},	
			"rgba255": &tender.UserFunction{
				Value: FuncAIIIIR(ctx.SetRGBA255),
			},	
			"rgb255": &tender.UserFunction{
				Value: FuncAIIIR(ctx.SetRGB255),
			},
			"hex": &tender.UserFunction{
				Value: FuncASR(ctx.SetHexColor),
			},
			"linewidth": &tender.UserFunction{
				Value: FuncAFR(ctx.SetLineWidth),
			},	
			"dashoffset": &tender.UserFunction{
				Value: FuncAFR(ctx.SetDashOffset),
			},
			"dash": &tender.UserFunction{
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
			"move_to": &tender.UserFunction{
				Value: FuncAFFR(ctx.MoveTo),
			},	
			"line_to": &tender.UserFunction{
				Value: FuncAFFR(ctx.LineTo),
			},	
			"quadratic_to": &tender.UserFunction{
				Value: FuncAFFFFR(ctx.QuadraticTo),
			},	
			"cubic_to": &tender.UserFunction{
				Value: FuncAFFFFFFR(ctx.CubicTo),
			},
			"closepath": &tender.UserFunction{
				Value: FuncAR(ctx.ClosePath),
			},	
			"clearpath": &tender.UserFunction{
				Value: FuncAR(ctx.ClearPath),
			},	
			"newsubpath": &tender.UserFunction{
				Value: FuncAR(ctx.NewSubPath),
			},	
			"clear": &tender.UserFunction{
				Value: FuncAR(ctx.Clear),
			},
			"stroke": &tender.UserFunction{
				Value: FuncAR(ctx.Stroke),
			},	
			"fill": &tender.UserFunction{
				Value: FuncAR(ctx.Fill),
			},		
			"stroke_preserve": &tender.UserFunction{
				Value: FuncAR(ctx.StrokePreserve),
			},	
			"fill_preserve": &tender.UserFunction{
				Value: FuncAR(ctx.FillPreserve),
			},	
			"text": &tender.UserFunction{
				Value: FuncASFFR(ctx.DrawString),
			},	
			"text_anchored": &tender.UserFunction{
				Value: FuncASFFFFR(ctx.DrawStringAnchored),
			},	
			"measure_text": &tender.UserFunction{
				Value: FuncASRFF(ctx.MeasureString),
			},	
			"measure_multiline_text": &tender.UserFunction{
				Value: FuncASFRFF(ctx.MeasureMultilineString),
			},	
			"load_fontface": &tender.UserFunction{
				Value: FuncASFRE(ctx.LoadFontFace),
			},	
			"fontface": &tender.UserFunction{
				Value: FuncAYFRE(ctx.FontFace),
			},	
			"fontheight": &tender.UserFunction{
				Value: FuncARF(ctx.FontHeight),
			},	
			"set_font": &tender.UserFunction{
				Value: FuncASRE(ctx.SetFont),
			},
			"identity": &tender.UserFunction{
				Name:  "identity",
				Value: FuncAR(ctx.Identity),
			},	
			"translate": &tender.UserFunction{
				Value: FuncAFFR(ctx.Translate),
			},	
			"scale": &tender.UserFunction{
				Value: FuncAFFR(ctx.Scale),
			},	
			"rotate": &tender.UserFunction{
				Value: FuncAFR(ctx.Rotate),
			},	
			"shear": &tender.UserFunction{
				Value: FuncAFFR(ctx.Shear),
			},
			"scaleabout": &tender.UserFunction{
				Value: FuncAFFFFR(ctx.ScaleAbout),
			},	
			"rotateabout": &tender.UserFunction{
				Value: FuncAFFFR(ctx.RotateAbout),
			},
			"shearabout": &tender.UserFunction{
				Value: FuncAFFFFR(ctx.ShearAbout),
			},	
			"transform_point": &tender.UserFunction{
				Value: FuncAFFRFF(ctx.TransformPoint),
			},
			"invertmask": &tender.UserFunction{
				Value: FuncAR(ctx.InvertMask),
			},	
			"inverty": &tender.UserFunction{
				Value: FuncAR(ctx.InvertY),
			},	
			"push": &tender.UserFunction{
				Value: FuncAR(ctx.Push),
			},	
			"pop": &tender.UserFunction{
				Value: FuncAR(ctx.Pop),
			},	
			"clip": &tender.UserFunction{
				Value: FuncAR(ctx.Clip),
			},		
			"clip_preserve": &tender.UserFunction{
				Value: FuncAR(ctx.ClipPreserve),
			},	
			"resetclip": &tender.UserFunction{
				Value: FuncAR(ctx.ResetClip),
			},
			"height": &tender.UserFunction{
				Value: FuncARI(ctx.Height),
			},	
			"width": &tender.UserFunction{
				Value: FuncARI(ctx.Width),
			},	
			"wordwrap": &tender.UserFunction{
				Value: FuncASFRSs(ctx.WordWrap),
			},
			"image": &tender.UserFunction{
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

