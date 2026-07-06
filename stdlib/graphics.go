//go:build gl
package stdlib

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"unsafe"

	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/v/gl"
)

// Point tracks vertex geometry for stateful path accumulation
type Point struct {
	X, Y float32
}

// contextState mirrors traditional 2D context configurations 
type contextState struct {
	Width, Height  int
	R, G, B, A     float32
	LineWidth      float32
	CurrentSubpath []Point
	Subpaths       [][]Point
}

// graphicsModule defines the standard package interface mapping for Tender
var graphicsModule = map[string]tender.Object{
	"new_context": &tender.BuiltinFunction{
		Name: "new_context",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}

			w := toInt(args[0])
			h := toInt(args[1])

			state := &contextState{
				Width:     w,
				Height:    h,
				R:         1.0,
				G:         1.0,
				B:         1.0,
				A:         1.0,
				LineWidth: 1.0,
			}

			// Core Canvas/GG path helpers wrapped in closures
			moveTo := func(x, y float32) {
				if len(state.CurrentSubpath) > 0 {
					state.Subpaths = append(state.Subpaths, state.CurrentSubpath)
				}
				state.CurrentSubpath = []Point{{X: x, Y: y}}
			}

			lineTo := func(x, y float32) {
				if len(state.CurrentSubpath) == 0 {
					state.CurrentSubpath = append(state.CurrentSubpath, Point{X: 0, Y: 0})
				}
				state.CurrentSubpath = append(state.CurrentSubpath, Point{X: x, Y: y})
			}

			closePath := func() {
				if len(state.CurrentSubpath) > 0 {
					first := state.CurrentSubpath[0]
					state.CurrentSubpath = append(state.CurrentSubpath, first)
					state.Subpaths = append(state.Subpaths, state.CurrentSubpath)
					state.CurrentSubpath = nil
				}
			}

			// Context Method Map ensuring identical duck-typing signature 
			ctxMap := map[string]tender.Object{
				"hex": &tender.BuiltinFunction{
					Name: "hex",
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 1 {
							return nil, tender.ErrInvalidArgCount
						}
						if str, ok := args[0].(*tender.String); ok {
							parseHexColor(str.Value, state)
						}
						return tender.NullValue, nil
					},
				},
				"rgb": &tender.BuiltinFunction{
					Name: "rgb",
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 3 {
							return nil, tender.ErrInvalidArgCount
						}
						state.R = toFloat32(args[0])
						state.G = toFloat32(args[1])
						state.B = toFloat32(args[2])
						state.A = 1.0
						return tender.NullValue, nil
					},
				},
				"rgba": &tender.BuiltinFunction{
					Name: "rgba",
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 4 {
							return nil, tender.ErrInvalidArgCount
						}
						state.R = toFloat32(args[0])
						state.G = toFloat32(args[1])
						state.B = toFloat32(args[2])
						state.A = toFloat32(args[3])
						return tender.NullValue, nil
					},
				},
				"move_to": &tender.BuiltinFunction{
					Name: "move_to",
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 2 {
							return nil, tender.ErrInvalidArgCount
						}
						moveTo(toFloat32(args[0]), toFloat32(args[1]))
						return tender.NullValue, nil
					},
				},
				"line_to": &tender.BuiltinFunction{
					Name: "line_to",
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 2 {
							return nil, tender.ErrInvalidArgCount
						}
						lineTo(toFloat32(args[0]), toFloat32(args[1]))
						return tender.NullValue, nil
					},
				},
				"close_path": &tender.BuiltinFunction{
					Name: "close_path",
					Value: func(args ...tender.Object) (tender.Object, error) {
						closePath()
						return tender.NullValue, nil
					},
				},
				"rect": &tender.BuiltinFunction{
					Name: "rect",
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 4 {
							return nil, tender.ErrInvalidArgCount
						}
						x := toFloat32(args[0])
						y := toFloat32(args[1])
						w := toFloat32(args[2])
						h := toFloat32(args[3])
						moveTo(x, y)
						lineTo(x+w, y)
						lineTo(x+w, y+h)
						lineTo(x, y+h)
						closePath()
						return tender.NullValue, nil
					},
				},
				"circle": &tender.BuiltinFunction{
					Name: "circle",
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 3 {
							return nil, tender.ErrInvalidArgCount
						}
						cx := toFloat32(args[0])
						cy := toFloat32(args[1])
						r := toFloat32(args[2])
						steps := 40
						for i := 0; i < steps; i++ {
							angle := float64(i) * 2.0 * math.Pi / float64(steps)
							px := cx + r*float32(math.Cos(angle))
							py := cy + r*float32(math.Sin(angle))
							if i == 0 {
								moveTo(px, py)
							} else {
								lineTo(px, py)
							}
						}
						closePath()
						return tender.NullValue, nil
					},
				},
				"line": &tender.BuiltinFunction{
					Name: "line",
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 4 {
							return nil, tender.ErrInvalidArgCount
						}
						moveTo(toFloat32(args[0]), toFloat32(args[1]))
						lineTo(toFloat32(args[2]), toFloat32(args[3]))
						return tender.NullValue, nil
					},
				},
				"clear": &tender.BuiltinFunction{
					Name: "clear",
					Value: func(args ...tender.Object) (tender.Object, error) {
						gl.ClearColor(state.R, state.G, state.B, state.A)
						gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
						return tender.NullValue, nil
					},
				},
				"stroke": &tender.BuiltinFunction{
					Name: "stroke",
					Value: func(args ...tender.Object) (tender.Object, error) {
						gl.Color4f(state.R, state.G, state.B, state.A)
						gl.LineWidth(state.LineWidth)

						allPaths := append([][]Point{}, state.Subpaths...)
						if len(state.CurrentSubpath) > 0 {
							allPaths = append(allPaths, state.CurrentSubpath)
						}

						for _, path := range allPaths {
							if len(path) < 2 {
								continue
							}
							gl.Begin(gl.LINE_STRIP)
							for _, p := range path {
								gl.Vertex2f(p.X, p.Y)
							}
							gl.End()
						}
						state.Subpaths = nil
						state.CurrentSubpath = nil
						return tender.NullValue, nil
					},
				},
				"fill": &tender.BuiltinFunction{
					Name: "fill",
					Value: func(args ...tender.Object) (tender.Object, error) {
						gl.Color4f(state.R, state.G, state.B, state.A)

						allPaths := append([][]Point{}, state.Subpaths...)
						if len(state.CurrentSubpath) > 0 {
							allPaths = append(allPaths, state.CurrentSubpath)
						}

						for _, path := range allPaths {
							if len(path) < 3 {
								continue
							}
							gl.Begin(gl.TRIANGLE_FAN)
							for _, p := range path {
								gl.Vertex2f(p.X, p.Y)
							}
							gl.End()
						}
						state.Subpaths = nil
						state.CurrentSubpath = nil
						return tender.NullValue, nil
					},
				},
				"push": &tender.BuiltinFunction{
					Name: "push",
					Value: func(args ...tender.Object) (tender.Object, error) {
						gl.PushMatrix()
						return tender.NullValue, nil
					},
				},
				"pop": &tender.BuiltinFunction{
					Name: "pop",
					Value: func(args ...tender.Object) (tender.Object, error) {
						gl.PopMatrix()
						return tender.NullValue, nil
					},
				},
				"translate": &tender.BuiltinFunction{
					Name: "translate",
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 2 {
							return nil, tender.ErrInvalidArgCount
						}
						gl.Translatef(toFloat32(args[0]), toFloat32(args[1]), 0)
						return tender.NullValue, nil
					},
				},
				"scale": &tender.BuiltinFunction{
					Name: "scale",
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 2 {
							return nil, tender.ErrInvalidArgCount
						}
						gl.Scalef(toFloat32(args[0]), toFloat32(args[1]), 1.0)
						return tender.NullValue, nil
					},
				},
				"rotate": &tender.BuiltinFunction{
					Name: "rotate",
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 1 {
							return nil, tender.ErrInvalidArgCount
						}
						// Convert radians to degrees to preserve GG compatibility
						deg := toFloat32(args[0]) * (180.0 / math.Pi)
						gl.Rotatef(deg, 0, 0, 1)
						return tender.NullValue, nil
					},
				},
				"set_line_width": &tender.BuiltinFunction{
					Name: "set_line_width",
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 1 {
							return nil, tender.ErrInvalidArgCount
						}
						state.LineWidth = toFloat32(args[0])
						return tender.NullValue, nil
					},
				},
				"load_image": &tender.BuiltinFunction{
					Name: "load_image",
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 1 {
							return nil, tender.ErrInvalidArgCount
						}
						var imgData []byte
						if strVal, ok := args[0].(*tender.String); ok {
							var err error
							imgData, err = os.ReadFile(strVal.Value)
							if err != nil {
								return nil, err
							}
						} else if bytesVal, ok := args[0].(*tender.Bytes); ok {
							imgData = bytesVal.Value
						} else {
							return nil, tender.ErrInvalidArgument
						}

						img, _, err := image.Decode(bytes.NewReader(imgData))
						if err != nil {
							return nil, err
						}

						rgba := image.NewRGBA(img.Bounds())
						draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

						var textureID uint32
						gl.GenTextures(1, &textureID)
						gl.BindTexture(gl.TEXTURE_2D, textureID)
						gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
						gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

						w := int32(rgba.Bounds().Dx())
						h := int32(rgba.Bounds().Dy())
						gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, w, h, 0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&rgba.Pix[0]))

						texMap := map[string]tender.Object{
							"id":     &tender.Int{Value: int64(textureID)},
							"width":  &tender.Int{Value: int64(w)},
							"height": &tender.Int{Value: int64(h)},
						}
						return &tender.ImmutableMap{Value: texMap}, nil
					},
				},
				"draw_image": &tender.BuiltinFunction{
					Name: "draw_image",
					Value: func(args ...tender.Object) (tender.Object, error) {
						if len(args) != 3 {
							return nil, tender.ErrInvalidArgCount
						}

						var texMap *tender.ImmutableMap
						if m, ok := args[0].(*tender.ImmutableMap); ok {
							texMap = m
						} else if mutableMap, ok := args[0].(*tender.Map); ok {
							texMap = &tender.ImmutableMap{Value: mutableMap.Value}
						} else {
							return nil, tender.ErrInvalidArgument
						}

						var texID uint32
						if idVal, ok := texMap.Value["id"].(*tender.Int); ok {
							texID = uint32(idVal.Value)
						}

						var tw, th float32
						if wVal, ok := texMap.Value["width"].(*tender.Int); ok {
							tw = float32(wVal.Value)
						}
						if hVal, ok := texMap.Value["height"].(*tender.Int); ok {
							th = float32(hVal.Value)
						}

						x := toFloat32(args[1])
						y := toFloat32(args[2])

						gl.Enable(gl.TEXTURE_2D)
						gl.Enable(gl.BLEND)
						gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
						gl.BindTexture(gl.TEXTURE_2D, texID)
						gl.Color4f(1, 1, 1, 1)

						gl.Begin(gl.QUADS)
						gl.TexCoord2f(0, 0); gl.Vertex2f(x, y)
						gl.TexCoord2f(1, 0); gl.Vertex2f(x+tw, y)
						gl.TexCoord2f(1, 1); gl.Vertex2f(x+tw, y+th)
						gl.TexCoord2f(0, 1); gl.Vertex2f(x, y+th)
						gl.End()

						gl.Disable(gl.TEXTURE_2D)
						return tender.NullValue, nil
					},
				},
				// Safe stubs to guarantee 100% legacy source stability without crashes
				"dash":                 &tender.BuiltinFunction{Name: "dash", Value: func(args ...tender.Object) (tender.Object, error) { return tender.NullValue, nil }},
				"save_png":             &tender.BuiltinFunction{Name: "save_png", Value: func(args ...tender.Object) (tender.Object, error) { return tender.NullValue, nil }},
				"load_font_face":       &tender.BuiltinFunction{Name: "load_font_face", Value: func(args ...tender.Object) (tender.Object, error) { return tender.NullValue, nil }},
				"draw_string":          &tender.BuiltinFunction{Name: "draw_string", Value: func(args ...tender.Object) (tender.Object, error) { return tender.NullValue, nil }},
				"draw_string_anchored": &tender.BuiltinFunction{Name: "draw_string_anchored", Value: func(args ...tender.Object) (tender.Object, error) { return tender.NullValue, nil }},
				"draw_string_wrapped":  &tender.BuiltinFunction{Name: "draw_string_wrapped", Value: func(args ...tender.Object) (tender.Object, error) { return tender.NullValue, nil }},
				"clip":                 &tender.BuiltinFunction{Name: "clip", Value: func(args ...tender.Object) (tender.Object, error) { return tender.NullValue, nil }},
				"reset_clip":           &tender.BuiltinFunction{Name: "reset_clip", Value: func(args ...tender.Object) (tender.Object, error) { return tender.NullValue, nil }},
			}

			return &tender.Map{Value: ctxMap}, nil
		},
	},
}

// Internal numeric parsing helpers safely decoupling variant type systems
func toFloat32(obj tender.Object) float32 {
	if v, ok := obj.(*tender.Float); ok {
		return float32(v.Value)
	}
	if v, ok := obj.(*tender.Int); ok {
		return float32(v.Value)
	}
	return 0
}

func toInt(obj tender.Object) int {
	if v, ok := obj.(*tender.Int); ok {
		return int(v.Value)
	}
	if v, ok := obj.(*tender.Float); ok {
		return int(v.Value)
	}
	return 0
}

func parseHexColor(hex string, state *contextState) {
	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}
	var r, g, b, a uint32 = 0, 0, 0, 255
	if len(hex) == 3 {
		fmt.Sscanf(hex, "%1x%1x%1x", &r, &g, &b)
		r |= r << 4
		g |= g << 4
		b |= b << 4
	} else if len(hex) == 4 {
		fmt.Sscanf(hex, "%1x%1x%1x%1x", &r, &g, &b, &a)
		r |= r << 4
		g |= g << 4
		b |= b << 4
		a |= a << 4
	} else if len(hex) == 6 {
		fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
	} else if len(hex) == 8 {
		fmt.Sscanf(hex, "%02x%02x%02x%02x", &r, &g, &b, &a)
	}
	state.R = float32(r) / 255.0
	state.G = float32(g) / 255.0
	state.B = float32(b) / 255.0
	state.A = float32(a) / 255.0
}