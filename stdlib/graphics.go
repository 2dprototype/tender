//go:build gl

package stdlib

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"strings"
	"sync"
	"unsafe"

	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/v/gl"
	"github.com/2dprototype/tender/v/glut"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
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
	FontFace       font.Face
	FontSize       float64
	Font           *truetype.Font
}

// graphicsModule defines the standard package interface mapping for Tender
var graphicsModule = map[string]tender.Object{
	"get_context": &tender.BuiltinFunction{
		Name: "get_context",
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

			ctxMap := createDrawingMethods(state)
			return &tender.Map{Value: ctxMap}, nil
		},
	},	
	"new_context": &tender.BuiltinFunction{
		Name: "new_context",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}

			w := toInt(args[0])
			h := toInt(args[1])
			
			glut.Init()
			glut.InitDisplayMode(glut.RGBA | glut.DOUBLE | glut.DEPTH | uint(0x0400) | uint(0x0800))
			glut.InitWindowSize(w, h)
			glut.CreateWindow("")
			glut.HideWindow()
			gl.Init()

			state := &contextState{
				Width:     w,
				Height:    h,
				R:         1.0,
				G:         1.0,
				B:         1.0,
				A:         1.0,
				LineWidth: 1.0,
			}

			gl.Viewport(0, 0, int32(w), int32(h))
			gl.MatrixMode(gl.PROJECTION)
			gl.LoadIdentity()
			gl.Ortho(0, float64(w), float64(h), 0, -1, 1)
			gl.MatrixMode(gl.MODELVIEW)
			gl.LoadIdentity()

			ctxMap := createDrawingMethods(state)

			ctxMap["width"] = &tender.Int{Value: int64(state.Width)}
			ctxMap["height"] = &tender.Int{Value: int64(state.Height)}

			return &tender.Map{Value: ctxMap}, nil
		},
	},
	"new_window": &tender.BuiltinFunction{
		Name:      "new_window",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			w := toInt(args[0])
			h := toInt(args[1])
			title, ok := args[2].(*tender.String)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}

			glut.Init()
			glut.InitDisplayMode(glut.RGB | glut.DOUBLE | glut.DEPTH)
			glut.InitWindowSize(w, h)
			glut.CreateWindow(title.Value)
			gl.Init()

			state := &contextState{
				Width:     w,
				Height:    h,
				R:         1.0,
				G:         1.0,
				B:         1.0,
				A:         1.0,
				LineWidth: 1.0,
			}

			setupViewport := func() {
				gl.Viewport(0, 0, int32(state.Width), int32(state.Height))
				gl.MatrixMode(gl.PROJECTION)
				gl.LoadIdentity()
				gl.Ortho(0, float64(state.Width), float64(state.Height), 0, -1, 1)
				gl.MatrixMode(gl.MODELVIEW)
				gl.LoadIdentity()
			}

			setupViewport()

			// Register default reshape function to adjust coordinate space dynamically
			glut.ReshapeFunc(func(rw, rh int) {
				state.Width = rw
				state.Height = rh
				setupViewport()
			})

			ctxMap := createDrawingMethods(state)

			// Add window specific event functions
			ctxMap["on_draw"] = &tender.BuiltinFunction{
				Name: "on_draw",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrInvalidArgCount
					}
					cb := args[0]
					if cb != tender.NullValue && !cb.CanCall() {
						return nil, tender.ErrNotCallable
					}

					glut.DisplayFunc(func() {
						setupViewport()
						// Default background color setup from draw color
						gl.ClearColor(0.15, 0.15, 0.18, 1.0)
						gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

						if cb != tender.NullValue {
							if _, err := tender.WrapFuncCall(vm, cb); err != nil {
								fmt.Println("on_draw callback error:", err)
							}
						}
						glut.SwapBuffers()
					})
					return tender.NullValue, nil
				},
			}

			ctxMap["on_update"] = &tender.BuiltinFunction{
				Name: "on_update",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrInvalidArgCount
					}
					cb := args[0]
					if cb != tender.NullValue && !cb.CanCall() {
						return nil, tender.ErrNotCallable
					}

					glut.IdleFunc(func() {
						if cb != tender.NullValue {
							if _, err := tender.WrapFuncCall(vm, cb); err != nil {
								fmt.Println("on_update callback error:", err)
							}
						}
						glut.PostRedisplay()
					})
					return tender.NullValue, nil
				},
			}

			ctxMap["on_key"] = &tender.BuiltinFunction{
				Name: "on_key",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrInvalidArgCount
					}
					cb := args[0]
					if cb != tender.NullValue && !cb.CanCall() {
						return nil, tender.ErrNotCallable
					}

					if cb == tender.NullValue {
						glut.KeyboardFunc(nil)
						glut.SpecialFunc(nil)
						return tender.NullValue, nil
					}

					glut.KeyboardFunc(func(key byte, x, y int) {
						if _, err := tender.WrapFuncCall(vm, cb,
							&tender.String{Value: string(key)},
							&tender.Int{Value: int64(x)},
							&tender.Int{Value: int64(y)},
						); err != nil {
							fmt.Println("on_key callback error:", err)
						}
					})

					glut.SpecialFunc(func(key int, x, y int) {
						if _, err := tender.WrapFuncCall(vm, cb,
							&tender.String{Value: specialKeyName(key)},
							&tender.Int{Value: int64(x)},
							&tender.Int{Value: int64(y)},
						); err != nil {
							fmt.Println("on_key (special) callback error:", err)
						}
					})

					return tender.NullValue, nil
				},
			}

			ctxMap["on_mouse"] = &tender.BuiltinFunction{
				Name: "on_mouse",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrInvalidArgCount
					}
					cb := args[0]
					if cb != tender.NullValue && !cb.CanCall() {
						return nil, tender.ErrNotCallable
					}

					if cb == tender.NullValue {
						glut.MouseFunc(nil)
						return tender.NullValue, nil
					}

					glut.MouseFunc(func(button, stateVal, x, y int) {
						if _, err := tender.WrapFuncCall(vm, cb,
							&tender.String{Value: mouseButtonName(button)},
							&tender.String{Value: mouseActionName(stateVal)},
							&tender.Int{Value: int64(x)},
							&tender.Int{Value: int64(y)},
						); err != nil {
							fmt.Println("on_mouse callback error:", err)
						}
					})
					return tender.NullValue, nil
				},
			}

			ctxMap["on_mouse_move"] = &tender.BuiltinFunction{
				Name: "on_mouse_move",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrInvalidArgCount
					}
					cb := args[0]
					if cb != tender.NullValue && !cb.CanCall() {
						return nil, tender.ErrNotCallable
					}

					if cb == tender.NullValue {
						glut.MotionFunc(nil)
						glut.PassiveMotionFunc(nil)
						return tender.NullValue, nil
					}

					moveFunc := func(x, y int) {
						if _, err := tender.WrapFuncCall(vm, cb,
							&tender.Int{Value: int64(x)},
							&tender.Int{Value: int64(y)},
						); err != nil {
							fmt.Println("on_mouse_move callback error:", err)
						}
					}
					glut.MotionFunc(moveFunc)
					glut.PassiveMotionFunc(moveFunc)
					return tender.NullValue, nil
				},
			}

			ctxMap["run"] = &tender.BuiltinFunction{
				Name: "run",
				Value: func(args ...tender.Object) (tender.Object, error) {
					glut.MainLoop()
					return tender.NullValue, nil
				},
			}

			// Add width and height getters dynamically as fields
			ctxMap["width"] = &tender.Int{Value: int64(state.Width)}
			ctxMap["height"] = &tender.Int{Value: int64(state.Height)}

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

func specialKeyName(key int) string {
	switch key {
	case glut.KEY_LEFT:
		return "left"
	case glut.KEY_RIGHT:
		return "right"
	case glut.KEY_UP:
		return "up"
	case glut.KEY_DOWN:
		return "down"
	case glut.KEY_PAGE_UP:
		return "page_up"
	case glut.KEY_PAGE_DOWN:
		return "page_down"
	case glut.KEY_HOME:
		return "home"
	case glut.KEY_END:
		return "end"
	case glut.KEY_INSERT:
		return "insert"
	case glut.KEY_F1:
		return "f1"
	case glut.KEY_F2:
		return "f2"
	case glut.KEY_F3:
		return "f3"
	case glut.KEY_F4:
		return "f4"
	case glut.KEY_F5:
		return "f5"
	case glut.KEY_F6:
		return "f6"
	case glut.KEY_F7:
		return "f7"
	case glut.KEY_F8:
		return "f8"
	case glut.KEY_F9:
		return "f9"
	case glut.KEY_F10:
		return "f10"
	case glut.KEY_F11:
		return "f11"
	case glut.KEY_F12:
		return "f12"
	default:
		return fmt.Sprintf("special_%d", key)
	}
}

func mouseButtonName(btn int) string {
	switch btn {
	case glut.LEFT_BUTTON:
		return "left"
	case glut.MIDDLE_BUTTON:
		return "middle"
	case glut.RIGHT_BUTTON:
		return "right"
	default:
		return fmt.Sprintf("button_%d", btn)
	}
}

func mouseActionName(action int) string {
	switch action {
	case glut.DOWN:
		return "down"
	case glut.UP:
		return "up"
	default:
		return fmt.Sprintf("action_%d", action)
	}
}

type textTexInfo struct {
	id     uint32
	width  int
	height int
	ascent int
}

var (
	textCacheMu sync.Mutex
	textCache   = make(map[string]textTexInfo)
)

func getOrCreateTextTexture(face font.Face, text string) textTexInfo {
	textCacheMu.Lock()
	defer textCacheMu.Unlock()

	key := fmt.Sprintf("%p_%s", face, text)
	if info, ok := textCache[key]; ok {
		return info
	}

	d := &font.Drawer{Face: face}
	ascent := int(math.Ceil(float64(face.Metrics().Ascent) / 64))
	descent := int(math.Ceil(float64(face.Metrics().Descent) / 64))
	width := int(math.Ceil(float64(d.MeasureString(text)) / 64))
	height := ascent + descent
	if width <= 0 {
		width = 1
	}
	if height <= 0 {
		height = 1
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// Initialize with transparent white
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i] = 255
		img.Pix[i+1] = 255
		img.Pix[i+2] = 255
		img.Pix[i+3] = 0
	}

	d.Dst = img
	d.Src = image.NewUniform(color.RGBA{255, 255, 255, 255})
	d.Dot = fixed.Point26_6{X: 0, Y: fixed.I(ascent)}
	d.DrawString(text)

	var textureID uint32
	gl.GenTextures(1, &textureID)
	gl.BindTexture(gl.TEXTURE_2D, textureID)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(width), int32(height), 0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&img.Pix[0]))

	info := textTexInfo{
		id:     textureID,
		width:  width,
		height: height,
		ascent: ascent,
	}
	textCache[key] = info
	return info
}

func wordWrap(face font.Face, s string, width float64) []string {
	d := &font.Drawer{Face: face}
	var lines []string
	for _, line := range strings.Split(s, "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			lines = append(lines, "")
			continue
		}
		currentLine := fields[0]
		for _, field := range fields[1:] {
			candidate := currentLine + " " + field
			w := float64(d.MeasureString(candidate) >> 6)
			if w > width {
				lines = append(lines, currentLine)
				currentLine = field
			} else {
				currentLine = candidate
			}
		}
		lines = append(lines, currentLine)
	}
	return lines
}

func drawEllipticalArc(state *contextState, x, y, rx, ry, angle1, angle2 float32) {
	const n = 16
	for i := 0; i < n; i++ {
		p1 := float32(i+0) / n
		p2 := float32(i+1) / n
		a1 := angle1 + (angle2-angle1)*p1
		a2 := angle1 + (angle2-angle1)*p2

		x0 := x + rx*float32(math.Cos(float64(a1)))
		y0 := y + ry*float32(math.Sin(float64(a1)))
		x1 := x + rx*float32(math.Cos(float64((a1+a2)/2)))
		y1 := y + ry*float32(math.Sin(float64((a1+a2)/2)))
		x2 := x + rx*float32(math.Cos(float64(a2)))
		y2 := y + ry*float32(math.Sin(float64(a2)))

		cx := 2*x1 - x0/2 - x2/2
		cy := 2*y1 - y0/2 - y2/2

		if i == 0 {
			if len(state.CurrentSubpath) > 0 {
				state.CurrentSubpath = append(state.CurrentSubpath, Point{X: x0, Y: y0})
			} else {
				state.CurrentSubpath = []Point{{X: x0, Y: y0}}
			}
		}

		var startPt Point
		if len(state.CurrentSubpath) > 0 {
			startPt = state.CurrentSubpath[len(state.CurrentSubpath)-1]
		}
		x0_pt, y0_pt := startPt.X, startPt.Y

		l := math.Hypot(float64(cx-x0_pt), float64(cy-y0_pt)) + math.Hypot(float64(x2-cx), float64(y2-cy))
		segN := int(l + 0.5)
		if segN < 4 {
			segN = 4
		}
		d := float32(segN - 1)
		for j := 1; j < segN; j++ {
			t := float32(j) / d
			u := 1.0 - t
			a := u * u
			b := 2.0 * u * t
			c := t * t
			px := a*x0_pt + b*cx + c*x2
			py := a*y0_pt + b*cy + c*y2
			state.CurrentSubpath = append(state.CurrentSubpath, Point{X: px, Y: py})
		}
	}
}

func drawRoundRect(state *contextState, x, y, w, h, rx, ry float32) {
	if rx <= 0 || ry <= 0 {
		if len(state.CurrentSubpath) > 0 {
			state.Subpaths = append(state.Subpaths, state.CurrentSubpath)
		}
		state.CurrentSubpath = []Point{
			{X: x, Y: y},
			{X: x + w, Y: y},
			{X: x + w, Y: y + h},
			{X: x, Y: y + h},
			{X: x, Y: y},
		}
		state.Subpaths = append(state.Subpaths, state.CurrentSubpath)
		state.CurrentSubpath = nil
		return
	}

	drawEllipticalArc(state, x+w-rx, y+ry, rx, ry, 1.5*math.Pi, 2.0*math.Pi)
	drawEllipticalArc(state, x+w-rx, y+h-ry, rx, ry, 0.0, 0.5*math.Pi)
	drawEllipticalArc(state, x+rx, y+h-ry, rx, ry, 0.5*math.Pi, 1.0*math.Pi)
	drawEllipticalArc(state, x+rx, y+ry, rx, ry, 1.0*math.Pi, 1.5*math.Pi)

	if len(state.CurrentSubpath) > 0 {
		first := state.CurrentSubpath[0]
		state.CurrentSubpath = append(state.CurrentSubpath, first)
		state.Subpaths = append(state.Subpaths, state.CurrentSubpath)
		state.CurrentSubpath = nil
	}
}

func captureGLImage(width, height int) image.Image {
	pix := make([]byte, width*height*4)
	gl.ReadPixels(0, 0, int32(width), int32(height), gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&pix[0]))

	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	rowSize := width * 4
	for y := 0; y < height; y++ {
		srcRow := (height - 1 - y) * rowSize
		destRow := y * rgba.Stride
		copy(rgba.Pix[destRow:destRow+rowSize], pix[srcRow:srcRow+rowSize])
	}
	return rgba
}

type imageTexInfo struct {
	id     uint32
	width  int
	height int
}

var (
	bytesImageCacheMu sync.Mutex
	bytesImageCache   = make(map[[32]byte]imageTexInfo)
)

func getOrCreateImageBytesTexture(data []byte) (imageTexInfo, error) {
	hash := sha256.Sum256(data)

	bytesImageCacheMu.Lock()
	defer bytesImageCacheMu.Unlock()

	if info, ok := bytesImageCache[hash]; ok {
		return info, nil
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return imageTexInfo{}, err
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var textureID uint32
	gl.GenTextures(1, &textureID)
	gl.BindTexture(gl.TEXTURE_2D, textureID)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	w := rgba.Bounds().Dx()
	h := rgba.Bounds().Dy()
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&rgba.Pix[0]))

	info := imageTexInfo{
		id:     textureID,
		width:  w,
		height: h,
	}
	bytesImageCache[hash] = info
	return info, nil
}

func createDrawingMethods(state *contextState) map[string]tender.Object {
	return map[string]tender.Object{
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
				x := toFloat32(args[0])
				y := toFloat32(args[1])
				if len(state.CurrentSubpath) > 0 {
					state.Subpaths = append(state.Subpaths, state.CurrentSubpath)
				}
				state.CurrentSubpath = []Point{{X: x, Y: y}}
				return tender.NullValue, nil
			},
		},
		"line_to": &tender.BuiltinFunction{
			Name: "line_to",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrInvalidArgCount
				}
				x := toFloat32(args[0])
				y := toFloat32(args[1])
				if len(state.CurrentSubpath) == 0 {
					state.CurrentSubpath = append(state.CurrentSubpath, Point{X: 0, Y: 0})
				}
				state.CurrentSubpath = append(state.CurrentSubpath, Point{X: x, Y: y})
				return tender.NullValue, nil
			},
		},
		"quadratic_to": &tender.BuiltinFunction{
			Name: "quadratic_to",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, tender.ErrInvalidArgCount
				}
				x1 := toFloat32(args[0])
				y1 := toFloat32(args[1])
				x2 := toFloat32(args[2])
				y2 := toFloat32(args[3])

				var x0, y0 float32 = 0, 0
				if len(state.CurrentSubpath) > 0 {
					last := state.CurrentSubpath[len(state.CurrentSubpath)-1]
					x0, y0 = last.X, last.Y
				} else {
					state.CurrentSubpath = append(state.CurrentSubpath, Point{X: 0, Y: 0})
				}

				l := math.Hypot(float64(x1-x0), float64(y1-y0)) + math.Hypot(float64(x2-x1), float64(y2-y1))
				n := int(l + 0.5)
				if n < 4 {
					n = 4
				}
				d := float32(n - 1)
				for i := 1; i < n; i++ {
					t := float32(i) / d
					u := 1.0 - t
					a := u * u
					b := 2.0 * u * t
					c := t * t
					px := a*x0 + b*x1 + c*x2
					py := a*y0 + b*y1 + c*y2
					state.CurrentSubpath = append(state.CurrentSubpath, Point{X: px, Y: py})
				}
				return tender.NullValue, nil
			},
		},
		"cubic_to": &tender.BuiltinFunction{
			Name: "cubic_to",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 6 {
					return nil, tender.ErrInvalidArgCount
				}
				x1 := toFloat32(args[0])
				y1 := toFloat32(args[1])
				x2 := toFloat32(args[2])
				y2 := toFloat32(args[3])
				x3 := toFloat32(args[4])
				y3 := toFloat32(args[5])

				var x0, y0 float32 = 0, 0
				if len(state.CurrentSubpath) > 0 {
					last := state.CurrentSubpath[len(state.CurrentSubpath)-1]
					x0, y0 = last.X, last.Y
				} else {
					state.CurrentSubpath = append(state.CurrentSubpath, Point{X: 0, Y: 0})
				}

				l := math.Hypot(float64(x1-x0), float64(y1-y0)) +
					math.Hypot(float64(x2-x1), float64(y2-y1)) +
					math.Hypot(float64(x3-x2), float64(y3-y2))
				n := int(l + 0.5)
				if n < 4 {
					n = 4
				}
				d := float32(n - 1)
				for i := 1; i < n; i++ {
					t := float32(i) / d
					u := 1.0 - t
					a := u * u * u
					b := 3.0 * u * u * t
					c := 3.0 * u * t * t
					dVal := t * t * t
					px := a*x0 + b*x1 + c*x2 + dVal*x3
					py := a*y0 + b*y1 + c*y2 + dVal*y3
					state.CurrentSubpath = append(state.CurrentSubpath, Point{X: px, Y: py})
				}
				return tender.NullValue, nil
			},
		},
		"arc": &tender.BuiltinFunction{
			Name: "arc",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 5 {
					return nil, tender.ErrInvalidArgCount
				}
				x := toFloat32(args[0])
				y := toFloat32(args[1])
				r := toFloat32(args[2])
				angle1 := toFloat32(args[3])
				angle2 := toFloat32(args[4])

				drawEllipticalArc(state, x, y, r, r, angle1, angle2)
				return tender.NullValue, nil
			},
		},
		"elliptical_arc": &tender.BuiltinFunction{
			Name: "elliptical_arc",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 6 {
					return nil, tender.ErrInvalidArgCount
				}
				x := toFloat32(args[0])
				y := toFloat32(args[1])
				rx := toFloat32(args[2])
				ry := toFloat32(args[3])
				angle1 := toFloat32(args[4])
				angle2 := toFloat32(args[5])

				drawEllipticalArc(state, x, y, rx, ry, angle1, angle2)
				return tender.NullValue, nil
			},
		},
		"close_path": &tender.BuiltinFunction{
			Name: "close_path",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(state.CurrentSubpath) > 0 {
					first := state.CurrentSubpath[0]
					state.CurrentSubpath = append(state.CurrentSubpath, first)
					state.Subpaths = append(state.Subpaths, state.CurrentSubpath)
					state.CurrentSubpath = nil
				}
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

				if len(state.CurrentSubpath) > 0 {
					state.Subpaths = append(state.Subpaths, state.CurrentSubpath)
				}
				state.CurrentSubpath = []Point{
					{X: x, Y: y},
					{X: x + w, Y: y},
					{X: x + w, Y: y + h},
					{X: x, Y: y + h},
					{X: x, Y: y},
				}
				state.Subpaths = append(state.Subpaths, state.CurrentSubpath)
				state.CurrentSubpath = nil
				return tender.NullValue, nil
			},
		},
		"roundrect": &tender.BuiltinFunction{
			Name: "roundrect",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 6 {
					return nil, tender.ErrInvalidArgCount
				}
				x := toFloat32(args[0])
				y := toFloat32(args[1])
				w := toFloat32(args[2])
				h := toFloat32(args[3])
				rx := toFloat32(args[4])
				ry := toFloat32(args[5])

				drawRoundRect(state, x, y, w, h, rx, ry)
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

				if len(state.CurrentSubpath) > 0 {
					state.Subpaths = append(state.Subpaths, state.CurrentSubpath)
				}

				state.CurrentSubpath = make([]Point, 0, steps+1)
				for i := 0; i <= steps; i++ {
					angle := float64(i) * 2.0 * math.Pi / float64(steps)
					px := cx + r*float32(math.Cos(angle))
					py := cy + r*float32(math.Sin(angle))
					state.CurrentSubpath = append(state.CurrentSubpath, Point{X: px, Y: py})
				}
				state.Subpaths = append(state.Subpaths, state.CurrentSubpath)
				state.CurrentSubpath = nil
				return tender.NullValue, nil
			},
		},
		"line": &tender.BuiltinFunction{
			Name: "line",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, tender.ErrInvalidArgCount
				}
				x1 := toFloat32(args[0])
				y1 := toFloat32(args[1])
				x2 := toFloat32(args[2])
				y2 := toFloat32(args[3])

				if len(state.CurrentSubpath) > 0 {
					state.Subpaths = append(state.Subpaths, state.CurrentSubpath)
				}
				state.CurrentSubpath = []Point{
					{X: x1, Y: y1},
					{X: x2, Y: y2},
				}
				state.Subpaths = append(state.Subpaths, state.CurrentSubpath)
				state.CurrentSubpath = nil
				return tender.NullValue, nil
			},
		},
		"point": &tender.BuiltinFunction{
			Name: "point",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrInvalidArgCount
				}
				x := toFloat32(args[0])
				y := toFloat32(args[1])

				gl.Color4f(state.R, state.G, state.B, state.A)
				gl.Begin(gl.POINTS)
				gl.Vertex2f(x, y)
				gl.End()
				return tender.NullValue, nil
			},
		},
		"clear": &tender.BuiltinFunction{
			Name: "clear",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) == 1 {
					if str, ok := args[0].(*tender.String); ok {
						parseHexColor(str.Value, state)
					}
				} else if len(args) == 3 {
					state.R = toFloat32(args[0])
					state.G = toFloat32(args[1])
					state.B = toFloat32(args[2])
					state.A = 1.0
				} else if len(args) == 4 {
					state.R = toFloat32(args[0])
					state.G = toFloat32(args[1])
					state.B = toFloat32(args[2])
					state.A = toFloat32(args[3])
				}
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

				if len(allPaths) > 0 {
					gl.EnableClientState(gl.VERTEX_ARRAY)
					for _, path := range allPaths {
						if len(path) < 2 {
							continue
						}
						gl.VertexPointer(2, gl.FLOAT, 0, unsafe.Pointer(&path[0]))
						gl.DrawArrays(gl.LINE_STRIP, 0, int32(len(path)))
					}
					gl.DisableClientState(gl.VERTEX_ARRAY)
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

				if len(allPaths) > 0 {
					gl.EnableClientState(gl.VERTEX_ARRAY)
					for _, path := range allPaths {
						if len(path) < 3 {
							continue
						}
						gl.VertexPointer(2, gl.FLOAT, 0, unsafe.Pointer(&path[0]))
						gl.DrawArrays(gl.TRIANGLE_FAN, 0, int32(len(path)))
					}
					gl.DisableClientState(gl.VERTEX_ARRAY)
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
					imgData, err = os.ReadFile(tender.ResolvePath(strVal.Value))
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
				gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
				gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

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

				var texID uint32
				var tw, th float32

				// Support bytes drawing (like canvas drawimage)
				if bytesObj, ok := args[0].(*tender.Bytes); ok {
					info, err := getOrCreateImageBytesTexture(bytesObj.Value)
					if err != nil {
						return nil, err
					}
					texID = info.id
					tw = float32(info.width)
					th = float32(info.height)
				} else if bytesSlice, ok := tender.ToByteSlice(args[0]); ok {
					info, err := getOrCreateImageBytesTexture(bytesSlice)
					if err != nil {
						return nil, err
					}
					texID = info.id
					tw = float32(info.width)
					th = float32(info.height)
				} else {
					var texMap *tender.ImmutableMap
					if m, ok := args[0].(*tender.ImmutableMap); ok {
						texMap = m
					} else if mutableMap, ok := args[0].(*tender.Map); ok {
						texMap = &tender.ImmutableMap{Value: mutableMap.Value}
					} else {
						return nil, tender.ErrInvalidArgument
					}

					if idVal, ok := texMap.Value["id"].(*tender.Int); ok {
						texID = uint32(idVal.Value)
					}
					if wVal, ok := texMap.Value["width"].(*tender.Int); ok {
						tw = float32(wVal.Value)
					}
					if hVal, ok := texMap.Value["height"].(*tender.Int); ok {
						th = float32(hVal.Value)
					}
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
		"draw_image_rect": &tender.BuiltinFunction{
			Name: "draw_image_rect",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 9 {
					return nil, tender.ErrInvalidArgCount
				}

				var texID uint32
				var tw, th float32

				// Support bytes drawing
				if bytesObj, ok := args[0].(*tender.Bytes); ok {
					info, err := getOrCreateImageBytesTexture(bytesObj.Value)
					if err != nil {
						return nil, err
					}
					texID = info.id
					tw = float32(info.width)
					th = float32(info.height)
				} else if bytesSlice, ok := tender.ToByteSlice(args[0]); ok {
					info, err := getOrCreateImageBytesTexture(bytesSlice)
					if err != nil {
						return nil, err
					}
					texID = info.id
					tw = float32(info.width)
					th = float32(info.height)
				} else {
					var texMap *tender.ImmutableMap
					if m, ok := args[0].(*tender.ImmutableMap); ok {
						texMap = m
					} else if mutableMap, ok := args[0].(*tender.Map); ok {
						texMap = &tender.ImmutableMap{Value: mutableMap.Value}
					} else {
						return nil, tender.ErrInvalidArgument
					}

					if idVal, ok := texMap.Value["id"].(*tender.Int); ok {
						texID = uint32(idVal.Value)
					}
					if wVal, ok := texMap.Value["width"].(*tender.Int); ok {
						tw = float32(wVal.Value)
					}
					if hVal, ok := texMap.Value["height"].(*tender.Int); ok {
						th = float32(hVal.Value)
					}
				}

				sx := toFloat32(args[1])
				sy := toFloat32(args[2])
				sw := toFloat32(args[3])
				sh := toFloat32(args[4])
				dx := toFloat32(args[5])
				dy := toFloat32(args[6])
				dw := toFloat32(args[7])
				dh := toFloat32(args[8])

				u0 := sx / tw
				v0 := sy / th
				u1 := (sx + sw) / tw
				v1 := (sy + sh) / th

				gl.Enable(gl.TEXTURE_2D)
				gl.Enable(gl.BLEND)
				gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
				gl.BindTexture(gl.TEXTURE_2D, texID)
				gl.Color4f(1, 1, 1, 1)

				gl.Begin(gl.QUADS)
				gl.TexCoord2f(u0, v0); gl.Vertex2f(dx, dy)
				gl.TexCoord2f(u1, v0); gl.Vertex2f(dx+dw, dy)
				gl.TexCoord2f(u1, v1); gl.Vertex2f(dx+dw, dy+dh)
				gl.TexCoord2f(u0, v1); gl.Vertex2f(dx, dy+dh)
				gl.End()

				gl.Disable(gl.TEXTURE_2D)
				return tender.NullValue, nil
			},
		},
		"load_font": &tender.BuiltinFunction{
			Name: "load_font",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrInvalidArgCount
				}
				path, ok := args[0].(*tender.String)
				if !ok {
					return nil, tender.ErrInvalidArgument
				}
				size := toFloat32(args[1])

				fontBytes, err := os.ReadFile(tender.ResolvePath(path.Value))
				if err != nil {
					return nil, err
				}
				f, err := truetype.Parse(fontBytes)
				if err != nil {
					return nil, err
				}
				face := truetype.NewFace(f, &truetype.Options{
					Size: float64(size),
				})
				state.Font = f
				state.FontFace = face
				state.FontSize = float64(size)

				return tender.NullValue, nil
			},
		},
		"load_font_face": &tender.BuiltinFunction{
			Name: "load_font_face",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, tender.ErrInvalidArgCount
				}
				path, ok := args[0].(*tender.String)
				if !ok {
					return nil, tender.ErrInvalidArgument
				}
				size := toFloat32(args[1])

				fontBytes, err := os.ReadFile(tender.ResolvePath(path.Value))
				if err != nil {
					return nil, err
				}
				f, err := truetype.Parse(fontBytes)
				if err != nil {
					return nil, err
				}
				face := truetype.NewFace(f, &truetype.Options{
					Size: float64(size),
				})
				state.Font = f
				state.FontFace = face
				state.FontSize = float64(size)

				return tender.NullValue, nil
			},
		},
		"set_font_size": &tender.BuiltinFunction{
			Name: "set_font_size",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrInvalidArgCount
				}
				size := toFloat32(args[0])
				state.FontSize = float64(size)

				if state.Font != nil {
					face := truetype.NewFace(state.Font, &truetype.Options{
						Size: float64(size),
					})
					state.FontFace = face
				}
				return tender.NullValue, nil
			},
		},
		"measure_string": &tender.BuiltinFunction{
			Name: "measure_string",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, tender.ErrInvalidArgCount
				}
				text := ""
				if strObj, ok := args[0].(*tender.String); ok {
					text = strObj.Value
				} else {
					text = fmt.Sprint(args[0])
				}

				face := state.FontFace
				if face == nil {
					face = basicfont.Face7x13
				}

				d := &font.Drawer{Face: face}
				w := float64(d.MeasureString(text) >> 6)

				metrics := face.Metrics()
				h := float64(metrics.Height) / 64.0
				if h == 0 {
					h = float64(metrics.Ascent + metrics.Descent) / 64.0
				}

				resMap := map[string]tender.Object{
					"width":  &tender.Float{Value: w},
					"height": &tender.Float{Value: h},
				}
				return &tender.ImmutableMap{Value: resMap}, nil
			},
		},
		"draw_string": &tender.BuiltinFunction{
			Name: "draw_string",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 3 {
					return nil, tender.ErrInvalidArgCount
				}
				x := toFloat32(args[0])
				y := toFloat32(args[1])
				text := ""
				if strObj, ok := args[2].(*tender.String); ok {
					text = strObj.Value
				} else {
					text = fmt.Sprint(args[2])
				}

				if text == "" {
					return tender.NullValue, nil
				}

				face := state.FontFace
				if face == nil {
					face = basicfont.Face7x13
				}

				info := getOrCreateTextTexture(face, text)

				gl.Enable(gl.TEXTURE_2D)
				gl.Enable(gl.BLEND)
				gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
				gl.BindTexture(gl.TEXTURE_2D, info.id)

				gl.Color4f(state.R, state.G, state.B, state.A)

				y_top := y - float32(info.ascent)
				y_bot := y_top + float32(info.height)
				x_right := x + float32(info.width)

				gl.Begin(gl.QUADS)
				gl.TexCoord2f(0, 0); gl.Vertex2f(x, y_top)
				gl.TexCoord2f(1, 0); gl.Vertex2f(x_right, y_top)
				gl.TexCoord2f(1, 1); gl.Vertex2f(x_right, y_bot)
				gl.TexCoord2f(0, 1); gl.Vertex2f(x, y_bot)
				gl.End()

				gl.Disable(gl.TEXTURE_2D)
				return tender.NullValue, nil
			},
		},
		"draw_string_anchored": &tender.BuiltinFunction{
			Name: "draw_string_anchored",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 5 {
					return nil, tender.ErrInvalidArgCount
				}
				x := toFloat32(args[0])
				y := toFloat32(args[1])
				text := ""
				if strObj, ok := args[2].(*tender.String); ok {
					text = strObj.Value
				} else {
					text = fmt.Sprint(args[2])
				}
				ax := toFloat32(args[3])
				ay := toFloat32(args[4])

				if text == "" {
					return tender.NullValue, nil
				}

				face := state.FontFace
				if face == nil {
					face = basicfont.Face7x13
				}

				info := getOrCreateTextTexture(face, text)

				metrics := face.Metrics()
				fontHeight := float32(metrics.Height) / 64.0
				if fontHeight == 0 {
					fontHeight = float32(metrics.Ascent+metrics.Descent) / 64.0
				}

				tx := x - ax*float32(info.width)
				ty := y + ay*fontHeight

				gl.Enable(gl.TEXTURE_2D)
				gl.Enable(gl.BLEND)
				gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
				gl.BindTexture(gl.TEXTURE_2D, info.id)

				gl.Color4f(state.R, state.G, state.B, state.A)

				y_top := ty - float32(info.ascent)
				y_bot := y_top + float32(info.height)
				x_right := tx + float32(info.width)

				gl.Begin(gl.QUADS)
				gl.TexCoord2f(0, 0); gl.Vertex2f(tx, y_top)
				gl.TexCoord2f(1, 0); gl.Vertex2f(x_right, y_top)
				gl.TexCoord2f(1, 1); gl.Vertex2f(x_right, y_bot)
				gl.TexCoord2f(0, 1); gl.Vertex2f(tx, y_bot)
				gl.End()

				gl.Disable(gl.TEXTURE_2D)
				return tender.NullValue, nil
			},
		},
		"draw_string_wrapped": &tender.BuiltinFunction{
			Name: "draw_string_wrapped",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) < 4 {
					return nil, tender.ErrInvalidArgCount
				}
				x := toFloat32(args[0])
				y := toFloat32(args[1])
				text := ""
				if strObj, ok := args[2].(*tender.String); ok {
					text = strObj.Value
				} else {
					text = fmt.Sprint(args[2])
				}
				width := toFloat32(args[3])

				lineSpacing := float32(1.5)
				if len(args) >= 5 {
					lineSpacing = toFloat32(args[4])
				}

				if text == "" {
					return tender.NullValue, nil
				}

				face := state.FontFace
				if face == nil {
					face = basicfont.Face7x13
				}

				metrics := face.Metrics()
				fontHeight := float32(metrics.Height) / 64.0
				if fontHeight == 0 {
					fontHeight = float32(metrics.Ascent+metrics.Descent) / 64.0
				}

				lines := wordWrap(face, text, float64(width))
				currentY := y
				for _, line := range lines {
					if line == "" {
						currentY += fontHeight * lineSpacing
						continue
					}

					info := getOrCreateTextTexture(face, line)

					gl.Enable(gl.TEXTURE_2D)
					gl.Enable(gl.BLEND)
					gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
					gl.BindTexture(gl.TEXTURE_2D, info.id)

					gl.Color4f(state.R, state.G, state.B, state.A)

					y_top := currentY - float32(info.ascent)
					y_bot := y_top + float32(info.height)
					x_right := x + float32(info.width)

					gl.Begin(gl.QUADS)
					gl.TexCoord2f(0, 0); gl.Vertex2f(x, y_top)
					gl.TexCoord2f(1, 0); gl.Vertex2f(x_right, y_top)
					gl.TexCoord2f(1, 1); gl.Vertex2f(x_right, y_bot)
					gl.TexCoord2f(0, 1); gl.Vertex2f(x, y_bot)
					gl.End()

					gl.Disable(gl.TEXTURE_2D)

					currentY += fontHeight * lineSpacing
				}
				return tender.NullValue, nil
			},
		},
		"encode": &tender.BuiltinFunction{
			Name: "encode",
			Value: func(args ...tender.Object) (tender.Object, error) {
				format := "png"
				if len(args) >= 1 {
					if strObj, ok := args[0].(*tender.String); ok {
						format = strObj.Value
					}
				}

				img := captureGLImage(state.Width, state.Height)
				buffer := new(bytes.Buffer)

				if format == "png" {
					err := png.Encode(buffer, img)
					if err != nil {
						return nil, err
					}
				} else if format == "jpeg" || format == "jpg" {
					err := jpeg.Encode(buffer, img, nil)
					if err != nil {
						return nil, err
					}
				} else {
					return nil, fmt.Errorf("unsupported encoding format: %s", format)
				}

				return &tender.Bytes{Value: buffer.Bytes()}, nil
			},
		},
		"save": &tender.BuiltinFunction{
			Name: "save",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) < 1 {
					return nil, tender.ErrInvalidArgCount
				}
				pathObj, ok := args[0].(*tender.String)
				if !ok {
					return nil, tender.ErrInvalidArgument
				}
				path := pathObj.Value

				format := "png"
				if len(args) >= 2 {
					if formatObj, ok := args[1].(*tender.String); ok {
						format = formatObj.Value
					}
				} else {
					if strings.HasSuffix(strings.ToLower(path), ".jpg") || strings.HasSuffix(strings.ToLower(path), ".jpeg") {
						format = "jpeg"
					}
				}

				img := captureGLImage(state.Width, state.Height)

				f, err := os.Create(path)
				if err != nil {
					return nil, err
				}
				defer f.Close()

				if format == "png" {
					err = png.Encode(f, img)
				} else if format == "jpeg" || format == "jpg" {
					err = jpeg.Encode(f, img, nil)
				} else {
					return nil, fmt.Errorf("unsupported save format: %s", format)
				}

				if err != nil {
					return nil, err
				}
				return tender.NullValue, nil
			},
		},
		"image": &tender.BuiltinFunction{
			Name: "image",
			Value: func(args ...tender.Object) (tender.Object, error) {
				img := captureGLImage(state.Width, state.Height)
				return makeImage(img), nil
			},
		},
	}
}