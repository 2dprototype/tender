package stdlib

import (
	"github.com/2dprototype/tender"
	"bytes"
	"image"
	"image/draw"
	"image/color"
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
	
	"fmt"
)

var canvasModule = map[string]tender.Object{
	"new_context": &tender.UserFunction{Name: "new_context", Value: ggNewContext},
	"load_image": &tender.UserFunction{Name:  "load_image", Value: imageLoad},	
	"radians": &tender.UserFunction{Name: "radians", Value: FuncAFRF(gg.Radians)},
	"degrees": &tender.UserFunction{Name: "degrees", Value: FuncAFRF(gg.Degrees)},
	"new_window" : &tender.BuiltinFunction{
		Name: "new_window",
		NeedVMObj: true,
		Value: canvasNewWindow,
	},	
	
	//new
	"new_linear_gradient": &tender.UserFunction{
		Name:  "new_linear_gradient",
		Value: canvasNewLinearGradient,
	},
	"new_radial_gradient": &tender.UserFunction{
		Name:  "new_radial_gradient",
		Value: canvasNewRadialGradient,
	},
	"new_conic_gradient": &tender.UserFunction{
		Name:  "new_conic_gradient",
		Value: canvasNewConicGradient,
	},
	"new_solid_pattern": &tender.UserFunction{
		Name:  "new_solid_pattern",
		Value: canvasNewSolidPattern,
	},
	"new_surface_pattern": &tender.UserFunction{
		Name:  "new_surface_pattern",
		Value: canvasNewSurfacePattern,
	},
}




func canvasNewLinearGradient(args ...tender.Object) (tender.Object, error) {
	if len(args) != 4 {
		return nil, tender.ErrWrongNumArguments
	}
	x0, _ := tender.ToFloat64(args[0])
	y0, _ := tender.ToFloat64(args[1])
	x1, _ := tender.ToFloat64(args[2])
	y1, _ := tender.ToFloat64(args[3])
	grad := gg.NewLinearGradient(x0, y0, x1, y1)
	return makeGradient(grad), nil
}

func canvasNewRadialGradient(args ...tender.Object) (tender.Object, error) {
	if len(args) != 6 {
		return nil, tender.ErrWrongNumArguments
	}
	x0, _ := tender.ToFloat64(args[0])
	y0, _ := tender.ToFloat64(args[1])
	r0, _ := tender.ToFloat64(args[2])
	x1, _ := tender.ToFloat64(args[3])
	y1, _ := tender.ToFloat64(args[4])
	r1, _ := tender.ToFloat64(args[5])
	grad := gg.NewRadialGradient(x0, y0, r0, x1, y1, r1)
	return makeGradient(grad), nil
}

func canvasNewConicGradient(args ...tender.Object) (tender.Object, error) {
	if len(args) != 3 {
		return nil, tender.ErrWrongNumArguments
	}
	cx, _ := tender.ToFloat64(args[0])
	cy, _ := tender.ToFloat64(args[1])
	deg, _ := tender.ToFloat64(args[2])
	grad := gg.NewConicGradient(cx, cy, deg)
	return makeGradient(grad), nil
}

func canvasNewSolidPattern(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	c, err := toColor(args[0])
	if err != nil {
		return nil, err
	}
	pat := gg.NewSolidPattern(c)
	return makePattern(pat), nil
}

func canvasNewSurfacePattern(args ...tender.Object) (tender.Object, error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}
	// Arg 0: image bytes or image object? Ideally we accept bytes like drawimage
	// But NewSurfacePattern needs image.Image.
	// We can reuse imageLoad logic or decode bytes.
	// Let's assume bytes for consistency with drawimage, OR a wrapped image object if we have one.
	// For now let's support bytes.
	imgBytes, _ := tender.ToByteSlice(args[0])
	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return wrapError(err), nil
	}

	opVal, _ := tender.ToInt(args[1])
	op := gg.RepeatOp(opVal)

	pat := gg.NewSurfacePattern(img, op)
	return makePattern(pat), nil
}

func makeGradient(grad gg.Gradient) *CanvasPattern {
	return &CanvasPattern{Value: grad}
}

func makePattern(pat gg.Pattern) *CanvasPattern {
	return &CanvasPattern{Value: pat}
}

type CanvasPattern struct {
	tender.ObjectImpl
	Value gg.Pattern
}

func (o *CanvasPattern) TypeName() string { return "canvas-pattern" }

func (o *CanvasPattern) String() string { return "<canvas-pattern>" }

func (o *CanvasPattern) Copy() tender.Object { return &CanvasPattern{Value: o.Value} }

func (o *CanvasPattern) Equals(x tender.Object) bool { return o == x }

func (o *CanvasPattern) IndexGet(index tender.Object) (tender.Object, error) {
	s, ok := tender.ToString(index)
	if !ok {
		return nil, nil
	}
	switch s {
	case "color_at":
		return &tender.UserFunction{
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrWrongNumArguments
				}
				x, _ := tender.ToInt(args[0])
				y, _ := tender.ToInt(args[1])
				c := o.Value.ColorAt(x, y)
				r, g, b, a := c.RGBA()
				return &tender.Array{
					Value: []tender.Object{
						&tender.Int{Value: int64(r >> 8)},
						&tender.Int{Value: int64(g >> 8)},
						&tender.Int{Value: int64(b >> 8)},
						&tender.Int{Value: int64(a >> 8)},
					},
				}, nil
			},
		}, nil
	case "add_color_stop":
		if grad, ok := o.Value.(gg.Gradient); ok {
			return &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 {
						return nil, tender.ErrWrongNumArguments
					}
					off, _ := tender.ToFloat64(args[0])
					c, err := toColor(args[1])
					if err != nil {
						return nil, err
					}
					grad.AddColorStop(off, c)
					return &tender.Null{}, nil
				},
			}, nil
		}
	}
	return nil, nil
}

func toColor(obj tender.Object) (color.Color, error) {
	// Try string (hex/name)
	if s, ok := tender.ToString(obj); ok {
		// Basic hex support
		if len(s) > 0 && s[0] == '#' {
			// use gg's hex parsing or simple implementation
			var r, g, b, a uint8 = 0, 0, 0, 255
			// A simple parser for #RRGGBB or #RGB or #RRGGBBAA
			hex := s[1:]
			if len(hex) == 3 {
				fmt.Sscanf(hex, "%1x%1x%1x", &r, &g, &b)
				r |= r << 4
				g |= g << 4
				b |= b << 4
			} else if len(hex) == 6 {
				fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
			} else if len(hex) == 8 {
				fmt.Sscanf(hex, "%02x%02x%02x%02x", &r, &g, &b, &a)
			} else {
				return nil, fmt.Errorf("invalid hex color format")
			}
			return color.RGBA{R: r, G: g, B: b, A: a}, nil
		}
		// If named colors support is needed, map names to colors.
		// For now fail.
		return nil, fmt.Errorf("unknown color string format")
	}
	
	// Support array [r, g, b, a] (0-255)
	// if arr, ok := obj.(*tender.Array); ok {
		// if len(arr.Value) >= 3 {
			// r, _ := tender.ToInt(arr.Value[0])
			// g, _ := tender.ToInt(arr.Value[1])
			// b, _ := tender.ToInt(arr.Value[2])
			// a := 255
			// if len(arr.Value) > 3 {
				// av, _ := tender.ToInt(arr.Value[3])
				// a = av
			// }
			// return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}, nil
		// }
	// }

	return color.Black, nil
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
			"linecap": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					c, _ := tender.ToInt(args[0])
					ctx.SetLineCap(gg.LineCap(c))
					return nil, nil
				},
			},
			"linecap_round": &tender.UserFunction{Value: FuncAR(ctx.SetLineCapRound)},
			"linecap_butt": &tender.UserFunction{Value: FuncAR(ctx.SetLineCapButt)},
			"linecap_square": &tender.UserFunction{Value: FuncAR(ctx.SetLineCapSquare)},
			"linejoin": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					c, _ := tender.ToInt(args[0])
					ctx.SetLineJoin(gg.LineJoin(c))
					return nil, nil
				},
			},
			"linejoin_round": &tender.UserFunction{Value: FuncAR(ctx.SetLineJoinRound)},
			"linejoin_bevel": &tender.UserFunction{Value: FuncAR(ctx.SetLineJoinBevel)},
			"fillrule": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					c, _ := tender.ToInt(args[0])
					ctx.SetFillRule(gg.FillRule(c))
					return nil, nil
				},
			},
			"fillrule_winding": &tender.UserFunction{Value: FuncAR(ctx.SetFillRuleWinding)},
			"fillrule_even_odd": &tender.UserFunction{Value: FuncAR(ctx.SetFillRuleEvenOdd)},
			"fillstyle": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					if cp, ok := args[0].(*CanvasPattern); ok {
						ctx.SetFillStyle(cp.Value)
						return nil, nil
					}
					return nil, tender.ErrInvalidArgumentType{Expected: "canvas-pattern"}
				},
			},
			"strokestyle": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					if cp, ok := args[0].(*CanvasPattern); ok {
						ctx.SetStrokeStyle(cp.Value)
						return nil, nil
					}
					return nil, tender.ErrInvalidArgumentType{Expected: "canvas-pattern"}
				},
			},
			"set_color": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					c, err := toColor(args[0])
					if err != nil {
						return nil, err
					}
					ctx.SetColor(c)
					return nil, nil
				},
			},
		},
	}
}
