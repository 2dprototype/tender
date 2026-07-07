//go:build glfw

package stdlib

import (
	"fmt"
	"unsafe"

	"github.com/2dprototype/tender"
	"github.com/go-gl/glfw/v3.4/glfw"
)

// glfwModule exposes essential GLFW functionality to Tender scripts
var glfwModule = map[string]tender.Object{
	// ================================================================
	// VERSION & INITIALIZATION
	// ================================================================

	"VERSION_MAJOR":    &tender.Int{Value: int64(glfw.VersionMajor)},
	"VERSION_MINOR":    &tender.Int{Value: int64(glfw.VersionMinor)},
	"VERSION_REVISION": &tender.Int{Value: int64(glfw.VersionRevision)},

	"init": &tender.BuiltinFunction{
		Name: "init",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			if err := glfw.Init(); err != nil {
				return nil, fmt.Errorf("glfw init failed: %w", err)
			}
			return tender.NullValue, nil
		},
	},

	"terminate": &tender.BuiltinFunction{
		Name: "terminate",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glfw.Terminate()
			return tender.NullValue, nil
		},
	},

	"poll_events": &tender.BuiltinFunction{
		Name: "poll_events",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glfw.PollEvents()
			return tender.NullValue, nil
		},
	},

	"wait_events": &tender.BuiltinFunction{
		Name: "wait_events",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glfw.WaitEvents()
			return tender.NullValue, nil
		},
	},

	"get_time": &tender.BuiltinFunction{
		Name: "get_time",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			return &tender.Float{Value: glfw.GetTime()}, nil
		},
	},

	"set_time": &tender.BuiltinFunction{
		Name: "set_time",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			t, ok := tender.ToFloat64(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glfw.SetTime(t)
			return tender.NullValue, nil
		},
	},

	// ================================================================
	// WINDOW HINTS
	// ================================================================

	"window_hint": &tender.BuiltinFunction{
		Name: "window_hint",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			hint, okH := tender.ToInt(args[0])
			val, okV := tender.ToInt(args[1])
			if !okH || !okV {
				return nil, tender.ErrInvalidArgument
			}
			glfw.WindowHint(glfw.Hint(hint), val)
			return tender.NullValue, nil
		},
	},

	"default_window_hints": &tender.BuiltinFunction{
		Name: "default_window_hints",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glfw.DefaultWindowHints()
			return tender.NullValue, nil
		},
	},

	// ================================================================
	// WINDOW MANAGEMENT
	// ================================================================

	"create_window": &tender.BuiltinFunction{
		Name: "create_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			w, okW := tender.ToInt(args[0])
			h, okH := tender.ToInt(args[1])
			title, okT := args[2].(*tender.String)
			if !okW || !okH || !okT {
				return nil, tender.ErrInvalidArgument
			}
			win, err := glfw.CreateWindow(w, h, title.Value, nil, nil)
			if err != nil {
				return nil, fmt.Errorf("create_window failed: %w", err)
			}
			return &tender.Int{Value: int64(uintptr(unsafe.Pointer(win)))}, nil
		},
	},

	"destroy_window": &tender.BuiltinFunction{
		Name: "destroy_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).Destroy()
			return tender.NullValue, nil
		},
	},

	"window_should_close": &tender.BuiltinFunction{
		Name: "window_should_close",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return tender.FalseValue, nil
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			return tender.FromBool(win.ShouldClose()), nil
		},
	},

	"set_window_should_close": &tender.BuiltinFunction{
		Name: "set_window_should_close",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			val, _ := tender.ToBool(args[1])
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).SetShouldClose(val)
			return tender.NullValue, nil
		},
	},

	"set_window_title": &tender.BuiltinFunction{
		Name: "set_window_title",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			title, okT := args[1].(*tender.String)
			if !ok || !okT || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).SetTitle(title.Value)
			return tender.NullValue, nil
		},
	},

	"get_window_size": &tender.BuiltinFunction{
		Name: "get_window_size",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			w, h := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).GetSize()
			return &tender.Array{Value: []tender.Object{
				&tender.Int{Value: int64(w)},
				&tender.Int{Value: int64(h)},
			}}, nil
		},
	},

	"set_window_size": &tender.BuiltinFunction{
		Name: "set_window_size",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			w, okW := tender.ToInt(args[1])
			h, okH := tender.ToInt(args[2])
			if !ok || !okW || !okH || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).SetSize(w, h)
			return tender.NullValue, nil
		},
	},

	"get_framebuffer_size": &tender.BuiltinFunction{
		Name: "get_framebuffer_size",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			w, h := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).GetFramebufferSize()
			return &tender.Array{Value: []tender.Object{
				&tender.Int{Value: int64(w)},
				&tender.Int{Value: int64(h)},
			}}, nil
		},
	},

	"get_window_pos": &tender.BuiltinFunction{
		Name: "get_window_pos",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			x, y := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).GetPos()
			return &tender.Array{Value: []tender.Object{
				&tender.Int{Value: int64(x)},
				&tender.Int{Value: int64(y)},
			}}, nil
		},
	},

	"set_window_pos": &tender.BuiltinFunction{
		Name: "set_window_pos",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			x, okX := tender.ToInt(args[1])
			y, okY := tender.ToInt(args[2])
			if !ok || !okX || !okY || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).SetPos(x, y)
			return tender.NullValue, nil
		},
	},

	"iconify_window": &tender.BuiltinFunction{
		Name: "iconify_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).Iconify()
			return tender.NullValue, nil
		},
	},

	"restore_window": &tender.BuiltinFunction{
		Name: "restore_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).Restore()
			return tender.NullValue, nil
		},
	},

	"maximize_window": &tender.BuiltinFunction{
		Name: "maximize_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).Maximize()
			return tender.NullValue, nil
		},
	},

	"show_window": &tender.BuiltinFunction{
		Name: "show_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).Show()
			return tender.NullValue, nil
		},
	},

	"hide_window": &tender.BuiltinFunction{
		Name: "hide_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).Hide()
			return tender.NullValue, nil
		},
	},

	"focus_window": &tender.BuiltinFunction{
		Name: "focus_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).Focus()
			return tender.NullValue, nil
		},
	},

	// ================================================================
	// CONTEXT MANAGEMENT
	// ================================================================

	"make_context_current": &tender.BuiltinFunction{
		Name: "make_context_current",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).MakeContextCurrent()
			return tender.NullValue, nil
		},
	},

	"get_current_context": &tender.BuiltinFunction{
		Name: "get_current_context",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			win := glfw.GetCurrentContext()
			if win == nil {
				return tender.NullValue, nil
			}
			return &tender.Int{Value: int64(uintptr(unsafe.Pointer(win)))}, nil
		},
	},

	"swap_buffers": &tender.BuiltinFunction{
		Name: "swap_buffers",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).SwapBuffers()
			return tender.NullValue, nil
		},
	},

	"swap_interval": &tender.BuiltinFunction{
		Name: "swap_interval",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			interval, ok := tender.ToInt(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glfw.SwapInterval(interval)
			return tender.NullValue, nil
		},
	},

	// ================================================================
	// INPUT HANDLING (Simple)
	// ================================================================

	"get_key": &tender.BuiltinFunction{
		Name: "get_key",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			key, okK := tender.ToInt(args[1])
			if !ok || !okK || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			return &tender.Int{Value: int64((*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).GetKey(glfw.Key(key)))}, nil
		},
	},

	"get_mouse_button": &tender.BuiltinFunction{
		Name: "get_mouse_button",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			btn, okB := tender.ToInt(args[1])
			if !ok || !okB || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			return &tender.Int{Value: int64((*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).GetMouseButton(glfw.MouseButton(btn)))}, nil
		},
	},

	"get_cursor_pos": &tender.BuiltinFunction{
		Name: "get_cursor_pos",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			x, y := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).GetCursorPos()
			return &tender.Array{Value: []tender.Object{
				&tender.Float{Value: x},
				&tender.Float{Value: y},
			}}, nil
		},
	},

	"set_cursor_pos": &tender.BuiltinFunction{
		Name: "set_cursor_pos",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			x, okX := tender.ToFloat64(args[1])
			y, okY := tender.ToFloat64(args[2])
			if !ok || !okX || !okY || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).SetCursorPos(x, y)
			return tender.NullValue, nil
		},
	},

	"set_input_mode": &tender.BuiltinFunction{
		Name: "set_input_mode",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			mode, okM := tender.ToInt(args[1])
			val, okV := tender.ToInt(args[2])
			if !ok || !okM || !okV || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			(*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value))).SetInputMode(glfw.InputMode(mode), val)
			return tender.NullValue, nil
		},
	},

	// ================================================================
	// CLIPBOARD
	// ================================================================

	"get_clipboard_string": &tender.BuiltinFunction{
		Name: "get_clipboard_string",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			return &tender.String{Value: glfw.GetClipboardString()}, nil
		},
	},

	"set_clipboard_string": &tender.BuiltinFunction{
		Name: "set_clipboard_string",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			str, ok := args[0].(*tender.String)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glfw.SetClipboardString(str.Value)
			return tender.NullValue, nil
		},
	},

	// ================================================================
	// MONITORS
	// ================================================================

	"get_primary_monitor": &tender.BuiltinFunction{
		Name: "get_primary_monitor",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			m := glfw.GetPrimaryMonitor()
			if m == nil {
				return tender.NullValue, nil
			}
			return &tender.Int{Value: int64(uintptr(unsafe.Pointer(m)))}, nil
		},
	},

	// "get_monitor_size": &tender.BuiltinFunction{
		// Name: "get_monitor_size",
		// Value: func(args ...tender.Object) (tender.Object, error) {
			// if len(args) != 1 {
				// return nil, tender.ErrInvalidArgCount
			// }
			// ptr, ok := args[0].(*tender.Int)
			// if !ok || ptr.Value == 0 {
				// return nil, tender.ErrInvalidArgument
			// }
			// m := (*glfw.Monitor)(unsafe.Pointer(uintptr(ptr.Value)))
			// w, h := m.GetSize()
			// return &tender.Array{Value: []tender.Object{
				// &tender.Int{Value: int64(w)},
				// &tender.Int{Value: int64(h)},
			// }}, nil
		// },
	// },

	"get_monitor_name": &tender.BuiltinFunction{
		Name: "get_monitor_name",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			m := (*glfw.Monitor)(unsafe.Pointer(uintptr(ptr.Value)))
			return &tender.String{Value: m.GetName()}, nil
		},
	},

	"get_video_mode": &tender.BuiltinFunction{
		Name: "get_video_mode",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			m := (*glfw.Monitor)(unsafe.Pointer(uintptr(ptr.Value)))
			mode := m.GetVideoMode()
			return &tender.Map{
				Value: map[string]tender.Object{
					"width":        &tender.Int{Value: int64(mode.Width)},
					"height":       &tender.Int{Value: int64(mode.Height)},
					"refresh_rate": &tender.Int{Value: int64(mode.RefreshRate)},
				},
			}, nil
		},
	},

	// ================================================================
	// CALLBACKS (VM-Aware)
	// ================================================================

	"set_window_size_callback": &tender.BuiltinFunction{
		Name:      "set_window_size_callback",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			cb := args[1]
			if cb != tender.NullValue && !cb.CanCall() {
				return nil, tender.ErrNotCallable
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			win.SetSizeCallback(func(_ *glfw.Window, w, h int) {
				if _, err := tender.WrapFuncCall(vm, cb,
					&tender.Int{Value: int64(w)},
					&tender.Int{Value: int64(h)},
				); err != nil {
					fmt.Println("GLFW size callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	"set_framebuffer_size_callback": &tender.BuiltinFunction{
		Name:      "set_framebuffer_size_callback",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			cb := args[1]
			if cb != tender.NullValue && !cb.CanCall() {
				return nil, tender.ErrNotCallable
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			win.SetFramebufferSizeCallback(func(_ *glfw.Window, w, h int) {
				if _, err := tender.WrapFuncCall(vm, cb,
					&tender.Int{Value: int64(w)},
					&tender.Int{Value: int64(h)},
				); err != nil {
					fmt.Println("GLFW framebuffer callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	"set_window_close_callback": &tender.BuiltinFunction{
		Name:      "set_window_close_callback",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			cb := args[1]
			if cb != tender.NullValue && !cb.CanCall() {
				return nil, tender.ErrNotCallable
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			win.SetCloseCallback(func(_ *glfw.Window) {
				if _, err := tender.WrapFuncCall(vm, cb); err != nil {
					fmt.Println("GLFW close callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	"set_key_callback": &tender.BuiltinFunction{
		Name:      "set_key_callback",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			cb := args[1]
			if cb != tender.NullValue && !cb.CanCall() {
				return nil, tender.ErrNotCallable
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			win.SetKeyCallback(func(_ *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
				if _, err := tender.WrapFuncCall(vm, cb,
					&tender.Int{Value: int64(key)},
					&tender.Int{Value: int64(scancode)},
					&tender.Int{Value: int64(action)},
					&tender.Int{Value: int64(mods)},
				); err != nil {
					fmt.Println("GLFW key callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	"set_char_callback": &tender.BuiltinFunction{
		Name:      "set_char_callback",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			cb := args[1]
			if cb != tender.NullValue && !cb.CanCall() {
				return nil, tender.ErrNotCallable
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			win.SetCharCallback(func(_ *glfw.Window, char rune) {
				if _, err := tender.WrapFuncCall(vm, cb,
					&tender.Char{Value: char},
				); err != nil {
					fmt.Println("GLFW char callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	"set_mouse_button_callback": &tender.BuiltinFunction{
		Name:      "set_mouse_button_callback",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			cb := args[1]
			if cb != tender.NullValue && !cb.CanCall() {
				return nil, tender.ErrNotCallable
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			win.SetMouseButtonCallback(func(_ *glfw.Window, btn glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
				if _, err := tender.WrapFuncCall(vm, cb,
					&tender.Int{Value: int64(btn)},
					&tender.Int{Value: int64(action)},
					&tender.Int{Value: int64(mods)},
				); err != nil {
					fmt.Println("GLFW mouse button callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	"set_cursor_pos_callback": &tender.BuiltinFunction{
		Name:      "set_cursor_pos_callback",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			cb := args[1]
			if cb != tender.NullValue && !cb.CanCall() {
				return nil, tender.ErrNotCallable
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			win.SetCursorPosCallback(func(_ *glfw.Window, x, y float64) {
				if _, err := tender.WrapFuncCall(vm, cb,
					&tender.Float{Value: x},
					&tender.Float{Value: y},
				); err != nil {
					fmt.Println("GLFW cursor pos callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	"set_scroll_callback": &tender.BuiltinFunction{
		Name:      "set_scroll_callback",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			cb := args[1]
			if cb != tender.NullValue && !cb.CanCall() {
				return nil, tender.ErrNotCallable
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			win.SetScrollCallback(func(_ *glfw.Window, xoff, yoff float64) {
				if _, err := tender.WrapFuncCall(vm, cb,
					&tender.Float{Value: xoff},
					&tender.Float{Value: yoff},
				); err != nil {
					fmt.Println("GLFW scroll callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	// ================================================================
	// WINDOW HINT CONSTANTS
	// ================================================================

	"RESIZABLE":               &tender.Int{Value: int64(glfw.Resizable)},
	"VISIBLE":                 &tender.Int{Value: int64(glfw.Visible)},
	"DECORATED":               &tender.Int{Value: int64(glfw.Decorated)},
	"FOCUSED":                 &tender.Int{Value: int64(glfw.Focused)},
	"MAXIMIZED":               &tender.Int{Value: int64(glfw.Maximized)},
	"TRANSPARENT_FRAMEBUFFER": &tender.Int{Value: int64(glfw.TransparentFramebuffer)},

	"CONTEXT_VERSION_MAJOR": &tender.Int{Value: int64(glfw.ContextVersionMajor)},
	"CONTEXT_VERSION_MINOR": &tender.Int{Value: int64(glfw.ContextVersionMinor)},
	"OPENGL_PROFILE":        &tender.Int{Value: int64(glfw.OpenGLProfile)},
	"OPENGL_CORE_PROFILE":   &tender.Int{Value: int64(glfw.OpenGLCoreProfile)},
	"OPENGL_COMPAT_PROFILE": &tender.Int{Value: int64(glfw.OpenGLCompatProfile)},

	"DOUBLEBUFFER": &tender.Int{Value: int64(glfw.DoubleBuffer)},

	// ================================================================
	// KEYBOARD KEYS
	// ================================================================

	"KEY_ESCAPE":  &tender.Int{Value: int64(glfw.KeyEscape)},
	"KEY_ENTER":   &tender.Int{Value: int64(glfw.KeyEnter)},
	"KEY_SPACE":   &tender.Int{Value: int64(glfw.KeySpace)},
	"KEY_TAB":     &tender.Int{Value: int64(glfw.KeyTab)},
	"KEY_BACKSPACE": &tender.Int{Value: int64(glfw.KeyBackspace)},
	"KEY_DELETE":  &tender.Int{Value: int64(glfw.KeyDelete)},
	"KEY_RIGHT":   &tender.Int{Value: int64(glfw.KeyRight)},
	"KEY_LEFT":    &tender.Int{Value: int64(glfw.KeyLeft)},
	"KEY_DOWN":    &tender.Int{Value: int64(glfw.KeyDown)},
	"KEY_UP":      &tender.Int{Value: int64(glfw.KeyUp)},

	"KEY_0": &tender.Int{Value: int64(glfw.Key0)},
	"KEY_1": &tender.Int{Value: int64(glfw.Key1)},
	"KEY_2": &tender.Int{Value: int64(glfw.Key2)},
	"KEY_3": &tender.Int{Value: int64(glfw.Key3)},
	"KEY_4": &tender.Int{Value: int64(glfw.Key4)},
	"KEY_5": &tender.Int{Value: int64(glfw.Key5)},
	"KEY_6": &tender.Int{Value: int64(glfw.Key6)},
	"KEY_7": &tender.Int{Value: int64(glfw.Key7)},
	"KEY_8": &tender.Int{Value: int64(glfw.Key8)},
	"KEY_9": &tender.Int{Value: int64(glfw.Key9)},

	"KEY_A": &tender.Int{Value: int64(glfw.KeyA)},
	"KEY_B": &tender.Int{Value: int64(glfw.KeyB)},
	"KEY_C": &tender.Int{Value: int64(glfw.KeyC)},
	"KEY_D": &tender.Int{Value: int64(glfw.KeyD)},
	"KEY_E": &tender.Int{Value: int64(glfw.KeyE)},
	"KEY_F": &tender.Int{Value: int64(glfw.KeyF)},
	"KEY_G": &tender.Int{Value: int64(glfw.KeyG)},
	"KEY_H": &tender.Int{Value: int64(glfw.KeyH)},
	"KEY_I": &tender.Int{Value: int64(glfw.KeyI)},
	"KEY_J": &tender.Int{Value: int64(glfw.KeyJ)},
	"KEY_K": &tender.Int{Value: int64(glfw.KeyK)},
	"KEY_L": &tender.Int{Value: int64(glfw.KeyL)},
	"KEY_M": &tender.Int{Value: int64(glfw.KeyM)},
	"KEY_N": &tender.Int{Value: int64(glfw.KeyN)},
	"KEY_O": &tender.Int{Value: int64(glfw.KeyO)},
	"KEY_P": &tender.Int{Value: int64(glfw.KeyP)},
	"KEY_Q": &tender.Int{Value: int64(glfw.KeyQ)},
	"KEY_R": &tender.Int{Value: int64(glfw.KeyR)},
	"KEY_S": &tender.Int{Value: int64(glfw.KeyS)},
	"KEY_T": &tender.Int{Value: int64(glfw.KeyT)},
	"KEY_U": &tender.Int{Value: int64(glfw.KeyU)},
	"KEY_V": &tender.Int{Value: int64(glfw.KeyV)},
	"KEY_W": &tender.Int{Value: int64(glfw.KeyW)},
	"KEY_X": &tender.Int{Value: int64(glfw.KeyX)},
	"KEY_Y": &tender.Int{Value: int64(glfw.KeyY)},
	"KEY_Z": &tender.Int{Value: int64(glfw.KeyZ)},

	"KEY_F1":  &tender.Int{Value: int64(glfw.KeyF1)},
	"KEY_F2":  &tender.Int{Value: int64(glfw.KeyF2)},
	"KEY_F3":  &tender.Int{Value: int64(glfw.KeyF3)},
	"KEY_F4":  &tender.Int{Value: int64(glfw.KeyF4)},
	"KEY_F5":  &tender.Int{Value: int64(glfw.KeyF5)},
	"KEY_F6":  &tender.Int{Value: int64(glfw.KeyF6)},
	"KEY_F7":  &tender.Int{Value: int64(glfw.KeyF7)},
	"KEY_F8":  &tender.Int{Value: int64(glfw.KeyF8)},
	"KEY_F9":  &tender.Int{Value: int64(glfw.KeyF9)},
	"KEY_F10": &tender.Int{Value: int64(glfw.KeyF10)},
	"KEY_F11": &tender.Int{Value: int64(glfw.KeyF11)},
	"KEY_F12": &tender.Int{Value: int64(glfw.KeyF12)},

	"KEY_LEFT_SHIFT":   &tender.Int{Value: int64(glfw.KeyLeftShift)},
	"KEY_LEFT_CONTROL": &tender.Int{Value: int64(glfw.KeyLeftControl)},
	"KEY_LEFT_ALT":     &tender.Int{Value: int64(glfw.KeyLeftAlt)},
	"KEY_RIGHT_SHIFT":  &tender.Int{Value: int64(glfw.KeyRightShift)},
	"KEY_RIGHT_CONTROL": &tender.Int{Value: int64(glfw.KeyRightControl)},
	"KEY_RIGHT_ALT":    &tender.Int{Value: int64(glfw.KeyRightAlt)},

	// ================================================================
	// KEY ACTIONS
	// ================================================================

	"PRESS":   &tender.Int{Value: int64(glfw.Press)},
	"RELEASE": &tender.Int{Value: int64(glfw.Release)},
	"REPEAT":  &tender.Int{Value: int64(glfw.Repeat)},

	// ================================================================
	// MOUSE BUTTONS
	// ================================================================

	"MOUSE_BUTTON_LEFT":   &tender.Int{Value: int64(glfw.MouseButtonLeft)},
	"MOUSE_BUTTON_RIGHT":  &tender.Int{Value: int64(glfw.MouseButtonRight)},
	"MOUSE_BUTTON_MIDDLE": &tender.Int{Value: int64(glfw.MouseButtonMiddle)},

	// ================================================================
	// CURSOR MODES
	// ================================================================

	"CURSOR_NORMAL":   &tender.Int{Value: int64(glfw.CursorNormal)},
	"CURSOR_HIDDEN":   &tender.Int{Value: int64(glfw.CursorHidden)},
	"CURSOR_DISABLED": &tender.Int{Value: int64(glfw.CursorDisabled)},

	// ================================================================
	// MODIFIER KEYS
	// ================================================================

	"MOD_SHIFT":   &tender.Int{Value: int64(glfw.ModShift)},
	"MOD_CONTROL": &tender.Int{Value: int64(glfw.ModControl)},
	"MOD_ALT":     &tender.Int{Value: int64(glfw.ModAlt)},

	// ================================================================
	// EXTRA HELPERS
	// ================================================================

	"get_version": &tender.BuiltinFunction{
		Name: "get_version",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			major, minor, rev := glfw.GetVersion()
			return &tender.Array{
				Value: []tender.Object{
					&tender.Int{Value: int64(major)},
					&tender.Int{Value: int64(minor)},
					&tender.Int{Value: int64(rev)},
				},
			}, nil
		},
	},

	"get_version_string": &tender.BuiltinFunction{
		Name: "get_version_string",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			return &tender.String{Value: glfw.GetVersionString()}, nil
		},
	},

	"vulkan_supported": &tender.BuiltinFunction{
		Name: "vulkan_supported",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			return tender.FromBool(glfw.VulkanSupported()), nil
		},
	},
}