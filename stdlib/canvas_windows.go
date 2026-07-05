//go:build windows

package stdlib

import (
	"fmt"
	"image"
	"image/draw"
	"sync/atomic"

	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/v/gg"
	"github.com/oakmound/shiny/driver"
	"github.com/oakmound/shiny/screen"
)

func init() {
	canvasModule["new_window"] = &tender.BuiltinFunction{
		Name:      "new_window",
		NeedVMObj: true,
		Value:     canvasNewWindow,
	}
}

var driverMainActive int32

func canvasNewWindow(args ...tender.Object) (ret tender.Object, err error) {
	vm := args[0].(*tender.VMObj).Value
	args = args[1:] // the first arg is VMObj inserted by VM
	if !(len(args) == 4 || len(args) == 2) {
		return nil, tender.ErrWrongNumArguments
	}

	// Prevent fatal crashes if user tries to open a second window
	if !atomic.CompareAndSwapInt32(&driverMainActive, 0, 1) {
		return nil, fmt.Errorf("a window is already open: driver.Main can only be called once per process")
	}
	defer atomic.StoreInt32(&driverMainActive, 0)

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
			Title:      title,
			Width:      width,
			Height:     height,
			Fullscreen: fullscreen,
			Borderless: borderless,
			TopMost:    topMost,
			NoScaling:  noScaling,
			X:          x,
			Y:          y,
		}
	}

	var callErr error

	// Use the Go exp/shiny package to create a window
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(wOpts)

		if err != nil {
			return
		}

		var ctx *gg.Context
		var screenImage screen.Image // Keep reference to reuse it
		var currentW, currentH int   // Track current dimensions

		wmap := &tender.ImmutableMap{
			Value: map[string]tender.Object{
				"release": &tender.UserFunction{
					Value: func(args ...tender.Object) (tender.Object, error) {
						// Clean up the persistent image buffer before releasing the window
						if screenImage != nil {
							screenImage.Release()
							screenImage = nil
						}
						w.Release()
						return nil, nil
					},
				},
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

						// SAFETY FIX: Prevent nil pointer panic if update is called before new_context
						if ctx == nil {
							return nil, fmt.Errorf("cannot update: context is nil. Call new_context first")
						}

						ww, _ := tender.ToInt(args[0])
						hh, _ := tender.ToInt(args[1])

						// SAFETY FIX: Only create a new OS image buffer if it doesn't exist,
						// or if the window dimensions have changed. Do not do this every frame.
						if screenImage == nil || currentW != ww || currentH != hh {
							if screenImage != nil {
								screenImage.Release() // Release old buffer
							}
							var imgErr error
							screenImage, imgErr = s.NewImage(image.Point{X: ww, Y: hh})
							if imgErr != nil {
								return wrapError(imgErr), nil
							}
							currentW = ww
							currentH = hh
						}

						// Draw to our persistent buffer
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

		defer func() {
			if screenImage != nil {
				screenImage.Release()
			}
			w.Release()
		}()

		if len(args) == 4 {
			_, callErr = tender.WrapFuncCall(vm, args[3], wmap)
		} else {
			_, callErr = tender.WrapFuncCall(vm, args[1], wmap)
		}
	})

	return nil, callErr
}
