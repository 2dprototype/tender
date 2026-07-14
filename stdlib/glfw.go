//go:build glfw

package stdlib

import (
	"fmt"
	"unsafe"

	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/v/glfw"
)

// glfwModule exposes essential GLFW functionality to Tender scripts
var glfwModule = map[string]tender.Object{
	// ================================================================
	// VERSION & INITIALIZATION
	// ================================================================

	"VERSION_MAJOR":    &tender.Int{Value: int64(glfw.VersionMajor)},
	"VERSION_MINOR":    &tender.Int{Value: int64(glfw.VersionMinor)},
	"VERSION_REVISION": &tender.Int{Value: int64(glfw.VersionRevision)},

	"init": &tender.NativeFunction{
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

	"terminate": &tender.NativeFunction{
		Name: "terminate",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glfw.Terminate()
			return tender.NullValue, nil
		},
	},

	"poll_events": &tender.NativeFunction{
		Name: "poll_events",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glfw.PollEvents()
			return tender.NullValue, nil
		},
	},

	"wait_events": &tender.NativeFunction{
		Name: "wait_events",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glfw.WaitEvents()
			return tender.NullValue, nil
		},
	},

	"get_time": &tender.NativeFunction{
		Name: "get_time",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			return &tender.Float{Value: glfw.GetTime()}, nil
		},
	},

	"set_time": &tender.NativeFunction{
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

	"window_hint": &tender.NativeFunction{
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

	"default_window_hints": &tender.NativeFunction{
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

	"create_window": &tender.NativeFunction{
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

	"destroy_window": &tender.NativeFunction{
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

	"window_should_close": &tender.NativeFunction{
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

	"set_window_should_close": &tender.NativeFunction{
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

	"set_window_title": &tender.NativeFunction{
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

	"get_window_size": &tender.NativeFunction{
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

	"set_window_size": &tender.NativeFunction{
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

	"get_framebuffer_size": &tender.NativeFunction{
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

	"get_window_pos": &tender.NativeFunction{
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

	"set_window_pos": &tender.NativeFunction{
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

	"iconify_window": &tender.NativeFunction{
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

	"restore_window": &tender.NativeFunction{
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

	"maximize_window": &tender.NativeFunction{
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

	"show_window": &tender.NativeFunction{
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

	"hide_window": &tender.NativeFunction{
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

	"focus_window": &tender.NativeFunction{
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

	"make_context_current": &tender.NativeFunction{
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

	"get_current_context": &tender.NativeFunction{
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

	"swap_buffers": &tender.NativeFunction{
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

	"swap_interval": &tender.NativeFunction{
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

	"get_key": &tender.NativeFunction{
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

	"get_mouse_button": &tender.NativeFunction{
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

	"get_cursor_pos": &tender.NativeFunction{
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

	"set_cursor_pos": &tender.NativeFunction{
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

	"set_input_mode": &tender.NativeFunction{
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

	"get_clipboard_string": &tender.NativeFunction{
		Name: "get_clipboard_string",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			return &tender.String{Value: glfw.GetClipboardString()}, nil
		},
	},

	"set_clipboard_string": &tender.NativeFunction{
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

	"get_primary_monitor": &tender.NativeFunction{
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

	"get_monitor_pos": &tender.NativeFunction{
		Name: "get_monitor_pos",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			m := (*glfw.Monitor)(unsafe.Pointer(uintptr(ptr.Value)))
			w, h := m.GetPos()
			return &tender.Array{Value: []tender.Object{
				&tender.Int{Value: int64(w)},
				&tender.Int{Value: int64(h)},
			}}, nil
		},
	},

	"get_monitor_name": &tender.NativeFunction{
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

	"get_video_mode": &tender.NativeFunction{
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

	"set_window_size_callback": &tender.NativeFunction{
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

	"set_framebuffer_size_callback": &tender.NativeFunction{
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

	"set_window_close_callback": &tender.NativeFunction{
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

	"set_key_callback": &tender.NativeFunction{
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

	"set_char_callback": &tender.NativeFunction{
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

	"set_mouse_button_callback": &tender.NativeFunction{
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

	"set_cursor_pos_callback": &tender.NativeFunction{
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

	"set_scroll_callback": &tender.NativeFunction{
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

	"FLOATING":              &tender.Int{Value: int64(glfw.Floating)},
	"FOCUS_ON_SHOW":         &tender.Int{Value: int64(glfw.FocusOnShow)},
	"CENTER_CURSOR":         &tender.Int{Value: int64(glfw.CenterCursor)},
	"SCALE_TO_MONITOR":      &tender.Int{Value: int64(glfw.ScaleToMonitor)},
	"SAMPLES":               &tender.Int{Value: int64(glfw.Samples)},
	"REFRESH_RATE":          &tender.Int{Value: int64(glfw.RefreshRate)},
	"STEREO":                &tender.Int{Value: int64(glfw.Stereo)},
	"SRGB_CAPABLE":          &tender.Int{Value: int64(glfw.SRGBCapable)},
	"RED_BITS":              &tender.Int{Value: int64(glfw.RedBits)},
	"GREEN_BITS":            &tender.Int{Value: int64(glfw.GreenBits)},
	"BLUE_BITS":             &tender.Int{Value: int64(glfw.BlueBits)},
	"ALPHA_BITS":            &tender.Int{Value: int64(glfw.AlphaBits)},
	"DEPTH_BITS":            &tender.Int{Value: int64(glfw.DepthBits)},
	"STENCIL_BITS":          &tender.Int{Value: int64(glfw.StencilBits)},
	"ACCUM_RED_BITS":        &tender.Int{Value: int64(glfw.AccumRedBits)},
	"ACCUM_GREEN_BITS":      &tender.Int{Value: int64(glfw.AccumGreenBits)},
	"ACCUM_BLUE_BITS":       &tender.Int{Value: int64(glfw.AccumBlueBits)},
	"ACCUM_ALPHA_BITS":      &tender.Int{Value: int64(glfw.AccumAlphaBits)},
	"AUX_BUFFERS":           &tender.Int{Value: int64(glfw.AuxBuffers)},

	// ================================================================
	// STANDARD CURSOR SHAPES
	// ================================================================

	"CURSOR_ARROW":     &tender.Int{Value: int64(glfw.ArrowCursor)},
	"CURSOR_IBEAM":     &tender.Int{Value: int64(glfw.IBeamCursor)},
	"CURSOR_CROSSHAIR": &tender.Int{Value: int64(glfw.CrosshairCursor)},
	"CURSOR_HAND":      &tender.Int{Value: int64(glfw.HandCursor)},
	"CURSOR_HRESIZE":   &tender.Int{Value: int64(glfw.HResizeCursor)},
	"CURSOR_VRESIZE":   &tender.Int{Value: int64(glfw.VResizeCursor)},

	// ================================================================
	// JOYSTICK CONSTANTS
	// ================================================================

	"JOYSTICK_1":  &tender.Int{Value: int64(glfw.Joystick1)},
	"JOYSTICK_2":  &tender.Int{Value: int64(glfw.Joystick2)},
	"JOYSTICK_3":  &tender.Int{Value: int64(glfw.Joystick3)},
	"JOYSTICK_4":  &tender.Int{Value: int64(glfw.Joystick4)},
	"JOYSTICK_5":  &tender.Int{Value: int64(glfw.Joystick5)},
	"JOYSTICK_6":  &tender.Int{Value: int64(glfw.Joystick6)},
	"JOYSTICK_7":  &tender.Int{Value: int64(glfw.Joystick7)},
	"JOYSTICK_8":  &tender.Int{Value: int64(glfw.Joystick8)},
	"JOYSTICK_9":  &tender.Int{Value: int64(glfw.Joystick9)},
	"JOYSTICK_10": &tender.Int{Value: int64(glfw.Joystick10)},
	"JOYSTICK_11": &tender.Int{Value: int64(glfw.Joystick11)},
	"JOYSTICK_12": &tender.Int{Value: int64(glfw.Joystick12)},
	"JOYSTICK_13": &tender.Int{Value: int64(glfw.Joystick13)},
	"JOYSTICK_14": &tender.Int{Value: int64(glfw.Joystick14)},
	"JOYSTICK_15": &tender.Int{Value: int64(glfw.Joystick15)},
	"JOYSTICK_16": &tender.Int{Value: int64(glfw.Joystick16)},

	"CONNECTED":    &tender.Int{Value: int64(glfw.Connected)},
	"DISCONNECTED": &tender.Int{Value: int64(glfw.Disconnected)},

	// ================================================================
	// PERIPHERAL EVENTS (for joystick callback)
	// ================================================================

	"PERIPHERAL_CONNECTED":    &tender.Int{Value: int64(glfw.Connected)},
	"PERIPHERAL_DISCONNECTED": &tender.Int{Value: int64(glfw.Disconnected)},

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
	// MORE KEYBOARD KEYS
	// ================================================================

	"KEY_GRAVE_ACCENT": &tender.Int{Value: int64(glfw.KeyGraveAccent)},
	"KEY_WORLD_1":      &tender.Int{Value: int64(glfw.KeyWorld1)},
	"KEY_WORLD_2":      &tender.Int{Value: int64(glfw.KeyWorld2)},
	"KEY_PRINT_SCREEN": &tender.Int{Value: int64(glfw.KeyPrintScreen)},
	"KEY_INSERT":       &tender.Int{Value: int64(glfw.KeyInsert)},
	"KEY_HOME":         &tender.Int{Value: int64(glfw.KeyHome)},
	"KEY_PAGE_UP":      &tender.Int{Value: int64(glfw.KeyPageUp)},
	"KEY_PAGE_DOWN":    &tender.Int{Value: int64(glfw.KeyPageDown)},
	"KEY_END":          &tender.Int{Value: int64(glfw.KeyEnd)},
	"KEY_CAPS_LOCK":    &tender.Int{Value: int64(glfw.KeyCapsLock)},
	"KEY_SCROLL_LOCK":  &tender.Int{Value: int64(glfw.KeyScrollLock)},
	"KEY_NUM_LOCK":     &tender.Int{Value: int64(glfw.KeyNumLock)},
	"KEY_PAUSE":        &tender.Int{Value: int64(glfw.KeyPause)},
	"KEY_MENU":         &tender.Int{Value: int64(glfw.KeyMenu)},

	// ================================================================
	// MORE MOUSE BUTTONS
	// ================================================================

	"MOUSE_BUTTON_4": &tender.Int{Value: int64(glfw.MouseButton4)},
	"MOUSE_BUTTON_5": &tender.Int{Value: int64(glfw.MouseButton5)},
	"MOUSE_BUTTON_6": &tender.Int{Value: int64(glfw.MouseButton6)},
	"MOUSE_BUTTON_7": &tender.Int{Value: int64(glfw.MouseButton7)},
	"MOUSE_BUTTON_8": &tender.Int{Value: int64(glfw.MouseButton8)},

	// ================================================================
	// MORE MODIFIER KEYS
	// ================================================================

	"MOD_SUPER": &tender.Int{Value: int64(glfw.ModSuper)},
	"MOD_CAPS_LOCK": &tender.Int{Value: int64(glfw.ModCapsLock)},
	"MOD_NUM_LOCK": &tender.Int{Value: int64(glfw.ModNumLock)},

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

	"LOCK_KEY_MODS": &tender.Int{Value: int64(glfw.LockKeyMods)},
	"RAW_MOUSE_MOTION": &tender.Int{Value: int64(glfw.RawMouseMotion)},

	// ================================================================
	// EXTRA HELPERS
	// ================================================================

	"get_version": &tender.NativeFunction{
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

	"get_version_string": &tender.NativeFunction{
		Name: "get_version_string",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			return &tender.String{Value: glfw.GetVersionString()}, nil
		},
	},

	"vulkan_supported": &tender.NativeFunction{
		Name: "vulkan_supported",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			return tender.FromBool(glfw.VulkanSupported()), nil
		},
	},
	
	
	// ================================================================
	// WINDOW ATTRIBUTES & QUERIES
	// ================================================================

	"get_input_mode": &tender.NativeFunction{
		Name: "get_input_mode",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			mode, okM := tender.ToInt(args[1])
			if !ok || !okM || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			return &tender.Int{Value: int64(win.GetInputMode(glfw.InputMode(mode)))}, nil
		},
	},

	"get_window_monitor": &tender.NativeFunction{
		Name: "get_window_monitor",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			m := win.GetMonitor()
			if m == nil {
				return tender.NullValue, nil
			}
			return &tender.Int{Value: int64(uintptr(unsafe.Pointer(m)))}, nil
		},
	},

	"set_window_monitor": &tender.NativeFunction{
		Name: "set_window_monitor",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 7 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			monitorPtr, okM := args[1].(*tender.Int)
			if !okM {
				return nil, tender.ErrInvalidArgument
			}
			xpos, okX := tender.ToInt(args[2])
			ypos, okY := tender.ToInt(args[3])
			width, okW := tender.ToInt(args[4])
			height, okH := tender.ToInt(args[5])
			refreshRate, okR := tender.ToInt(args[6])
			if !okX || !okY || !okW || !okH || !okR {
				return nil, tender.ErrInvalidArgument
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			var monitor *glfw.Monitor
			if monitorPtr.Value != 0 {
				monitor = (*glfw.Monitor)(unsafe.Pointer(uintptr(monitorPtr.Value)))
			}
			win.SetMonitor(monitor, xpos, ypos, width, height, refreshRate)
			return tender.NullValue, nil
		},
	},

	"get_window_attrib": &tender.NativeFunction{
		Name: "get_window_attrib",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			attrib, okA := tender.ToInt(args[1])
			if !ok || !okA || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			return &tender.Int{Value: int64(win.GetAttrib(glfw.Hint(attrib)))}, nil
		},
	},

	"set_window_attrib": &tender.NativeFunction{
		Name: "set_window_attrib",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			attrib, okA := tender.ToInt(args[1])
			value, okV := tender.ToInt(args[2])
			if !ok || !okA || !okV || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			win.SetAttrib(glfw.Hint(attrib), value)
			return tender.NullValue, nil
		},
	},

	"get_window_content_scale": &tender.NativeFunction{
		Name: "get_window_content_scale",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			x, y := win.GetContentScale()
			return &tender.Array{Value: []tender.Object{
				&tender.Float{Value: float64(x)},
				&tender.Float{Value: float64(y)},
			}}, nil
		},
	},

	"get_window_opacity": &tender.NativeFunction{
		Name: "get_window_opacity",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			return &tender.Float{Value: float64(win.GetOpacity())}, nil
		},
	},

	"set_window_opacity": &tender.NativeFunction{
		Name: "set_window_opacity",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			opacity, okO := tender.ToFloat64(args[1])
			if !ok || !okO || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			win.SetOpacity(float32(opacity))
			return tender.NullValue, nil
		},
	},

	"set_window_size_limits": &tender.NativeFunction{
		Name: "set_window_size_limits",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 5 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			minw, okMW := tender.ToInt(args[1])
			minh, okMH := tender.ToInt(args[2])
			maxw, okXW := tender.ToInt(args[3])
			maxh, okXH := tender.ToInt(args[4])
			if !ok || !okMW || !okMH || !okXW || !okXH || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			win.SetSizeLimits(minw, minh, maxw, maxh)
			return tender.NullValue, nil
		},
	},

	"set_window_aspect_ratio": &tender.NativeFunction{
		Name: "set_window_aspect_ratio",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			numer, okN := tender.ToInt(args[1])
			denom, okD := tender.ToInt(args[2])
			if !ok || !okN || !okD || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			win.SetAspectRatio(numer, denom)
			return tender.NullValue, nil
		},
	},

	"request_window_attention": &tender.NativeFunction{
		Name: "request_window_attention",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			win.RequestAttention()
			return tender.NullValue, nil
		},
	},

	"set_window_icon": &tender.NativeFunction{
		Name: "set_window_icon",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			// This is a simplified version - full image support would require more work
			// For now, just accept null to clear the icon
			if args[1] == tender.NullValue {
				win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
				win.SetIcon(nil)
				return tender.NullValue, nil
			}
			return nil, fmt.Errorf("set_window_icon requires image support (use null to clear)")
		},
	},

	// ================================================================
	// CURSORS
	// ================================================================

	"create_standard_cursor": &tender.NativeFunction{
		Name: "create_standard_cursor",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			shape, ok := tender.ToInt(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			cursor := glfw.CreateStandardCursor(glfw.StandardCursor(shape))
			if cursor == nil {
				return tender.NullValue, nil
			}
			return &tender.Int{Value: int64(uintptr(unsafe.Pointer(cursor)))}, nil
		},
	},

	"destroy_cursor": &tender.NativeFunction{
		Name: "destroy_cursor",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			cursor := (*glfw.Cursor)(unsafe.Pointer(uintptr(ptr.Value)))
			cursor.Destroy()
			return tender.NullValue, nil
		},
	},

	"set_window_cursor": &tender.NativeFunction{
		Name: "set_window_cursor",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			cursorPtr, okC := args[1].(*tender.Int)
			if !okC {
				return nil, tender.ErrInvalidArgument
			}
			win := (*glfw.Window)(unsafe.Pointer(uintptr(ptr.Value)))
			var cursor *glfw.Cursor
			if cursorPtr.Value != 0 {
				cursor = (*glfw.Cursor)(unsafe.Pointer(uintptr(cursorPtr.Value)))
			}
			win.SetCursor(cursor)
			return tender.NullValue, nil
		},
	},

	// ================================================================
	// TIMER
	// ================================================================

	"get_timer_frequency": &tender.NativeFunction{
		Name: "get_timer_frequency",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			return &tender.Int{Value: int64(glfw.GetTimerFrequency())}, nil
		},
	},

	"get_timer_value": &tender.NativeFunction{
		Name: "get_timer_value",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			return &tender.Int{Value: int64(glfw.GetTimerValue())}, nil
		},
	},

	// ================================================================
	// GAMEPAD & JOYSTICK
	// ================================================================

	"update_gamepad_mappings": &tender.NativeFunction{
		Name: "update_gamepad_mappings",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			mapping, ok := args[0].(*tender.String)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			return tender.FromBool(glfw.UpdateGamepadMappings(mapping.Value)), nil
		},
	},

	"joystick_present": &tender.NativeFunction{
		Name: "joystick_present",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			jid, ok := tender.ToInt(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			return tender.FromBool(glfw.Joystick(jid).Present()), nil
		},
	},

	"joystick_get_name": &tender.NativeFunction{
		Name: "joystick_get_name",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			jid, ok := tender.ToInt(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			name := glfw.Joystick(jid).GetName()
			if name == "" {
				return tender.NullValue, nil
			}
			return &tender.String{Value: name}, nil
		},
	},

	"joystick_is_gamepad": &tender.NativeFunction{
		Name: "joystick_is_gamepad",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			jid, ok := tender.ToInt(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			return tender.FromBool(glfw.Joystick(jid).IsGamepad()), nil
		},
	},

	"joystick_get_gamepad_name": &tender.NativeFunction{
		Name: "joystick_get_gamepad_name",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			jid, ok := tender.ToInt(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			name := glfw.Joystick(jid).GetGamepadName()
			if name == "" {
				return tender.NullValue, nil
			}
			return &tender.String{Value: name}, nil
		},
	},

	"joystick_get_gamepad_state": &tender.NativeFunction{
		Name: "joystick_get_gamepad_state",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			jid, ok := tender.ToInt(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			state := glfw.Joystick(jid).GetGamepadState()
			if state == nil {
				return tender.NullValue, nil
			}
			// Convert axes to array
			axes := make([]tender.Object, len(state.Axes))
			for i, a := range state.Axes {
				axes[i] = &tender.Float{Value: float64(a)}
			}
			// Convert buttons to array
			buttons := make([]tender.Object, len(state.Buttons))
			for i, b := range state.Buttons {
				buttons[i] = &tender.Int{Value: int64(b)}
			}
			return &tender.Map{
				Value: map[string]tender.Object{
					"axes":    &tender.Array{Value: axes},
					"buttons": &tender.Array{Value: buttons},
				},
			}, nil
		},
	},

	"joystick_get_axes": &tender.NativeFunction{
		Name: "joystick_get_axes",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			jid, ok := tender.ToInt(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			axes := glfw.Joystick(jid).GetAxes()
			if axes == nil {
				return tender.NullValue, nil
			}
			result := make([]tender.Object, len(axes))
			for i, a := range axes {
				result[i] = &tender.Float{Value: float64(a)}
			}
			return &tender.Array{Value: result}, nil
		},
	},

	"joystick_get_buttons": &tender.NativeFunction{
		Name: "joystick_get_buttons",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			jid, ok := tender.ToInt(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			buttons := glfw.Joystick(jid).GetButtons()
			if buttons == nil {
				return tender.NullValue, nil
			}
			result := make([]tender.Object, len(buttons))
			for i, b := range buttons {
				result[i] = &tender.Int{Value: int64(b)}
			}
			return &tender.Array{Value: result}, nil
		},
	},

	"joystick_get_hats": &tender.NativeFunction{
		Name: "joystick_get_hats",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			jid, ok := tender.ToInt(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			hats := glfw.Joystick(jid).GetHats()
			if hats == nil {
				return tender.NullValue, nil
			}
			result := make([]tender.Object, len(hats))
			for i, h := range hats {
				result[i] = &tender.Int{Value: int64(h)}
			}
			return &tender.Array{Value: result}, nil
		},
	},

	"joystick_get_guid": &tender.NativeFunction{
		Name: "joystick_get_guid",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			jid, ok := tender.ToInt(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			guid := glfw.Joystick(jid).GetGUID()
			if guid == "" {
				return tender.NullValue, nil
			}
			return &tender.String{Value: guid}, nil
		},
	},

	"set_joystick_callback": &tender.NativeFunction{
		Name:      "set_joystick_callback",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			cb := args[0]
			if cb != tender.NullValue && !cb.CanCall() {
				return nil, tender.ErrNotCallable
			}
			glfw.SetJoystickCallback(func(joy glfw.Joystick, event glfw.PeripheralEvent) {
				if _, err := tender.WrapFuncCall(vm, cb,
					&tender.Int{Value: int64(joy)},
					&tender.Int{Value: int64(event)},
				); err != nil {
					fmt.Println("GLFW joystick callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	// ================================================================
	// MONITOR EXTENDED INFO
	// ================================================================

	"get_monitor_video_modes": &tender.NativeFunction{
		Name: "get_monitor_video_modes",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			m := (*glfw.Monitor)(unsafe.Pointer(uintptr(ptr.Value)))
			modes := m.GetVideoModes()
			result := make([]tender.Object, len(modes))
			for i, mode := range modes {
				result[i] = &tender.Map{
					Value: map[string]tender.Object{
						"width":        &tender.Int{Value: int64(mode.Width)},
						"height":       &tender.Int{Value: int64(mode.Height)},
						"refresh_rate": &tender.Int{Value: int64(mode.RefreshRate)},
					},
				}
			}
			return &tender.Array{Value: result}, nil
		},
	},

	"get_monitor_physical_size": &tender.NativeFunction{
		Name: "get_monitor_physical_size",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			m := (*glfw.Monitor)(unsafe.Pointer(uintptr(ptr.Value)))
			width, height := m.GetPhysicalSize()
			return &tender.Array{Value: []tender.Object{
				&tender.Int{Value: int64(width)},
				&tender.Int{Value: int64(height)},
			}}, nil
		},
	},

	"get_monitor_workarea": &tender.NativeFunction{
		Name: "get_monitor_workarea",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			m := (*glfw.Monitor)(unsafe.Pointer(uintptr(ptr.Value)))
			x, y, width, height := m.GetWorkarea()
			return &tender.Array{Value: []tender.Object{
				&tender.Int{Value: int64(x)},
				&tender.Int{Value: int64(y)},
				&tender.Int{Value: int64(width)},
				&tender.Int{Value: int64(height)},
			}}, nil
		},
	},

	"get_monitor_gamma_ramp": &tender.NativeFunction{
		Name: "get_monitor_gamma_ramp",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			m := (*glfw.Monitor)(unsafe.Pointer(uintptr(ptr.Value)))
			ramp := m.GetGammaRamp()
			if ramp == nil {
				return tender.NullValue, nil
			}
			// Convert to Tender arrays
			red := make([]tender.Object, len(ramp.Red))
			green := make([]tender.Object, len(ramp.Green))
			blue := make([]tender.Object, len(ramp.Blue))
			for i := range ramp.Red {
				red[i] = &tender.Int{Value: int64(ramp.Red[i])}
				green[i] = &tender.Int{Value: int64(ramp.Green[i])}
				blue[i] = &tender.Int{Value: int64(ramp.Blue[i])}
			}
			return &tender.Map{
				Value: map[string]tender.Object{
					"red":   &tender.Array{Value: red},
					"green": &tender.Array{Value: green},
					"blue":  &tender.Array{Value: blue},
				},
			}, nil
		},
	},

	"set_monitor_gamma_ramp": &tender.NativeFunction{
		Name: "set_monitor_gamma_ramp",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			rampMap, okR := args[1].(*tender.Map)
			if !okR {
				return nil, tender.ErrInvalidArgument
			}
			redArr, okRArr := rampMap.Value["red"].(*tender.Array)
			greenArr, okGArr := rampMap.Value["green"].(*tender.Array)
			blueArr, okBArr := rampMap.Value["blue"].(*tender.Array)
			if !okRArr || !okGArr || !okBArr {
				return nil, tender.ErrInvalidArgument
			}
			// Ensure all arrays are same length
			if len(redArr.Value) != len(greenArr.Value) || len(redArr.Value) != len(blueArr.Value) {
				return nil, fmt.Errorf("color arrays must have same length")
			}
			ramp := &glfw.GammaRamp{
				Red:   make([]uint16, len(redArr.Value)),
				Green: make([]uint16, len(greenArr.Value)),
				Blue:  make([]uint16, len(blueArr.Value)),
			}
			for i := range redArr.Value {
				r, _ := tender.ToUint64(redArr.Value[i])
				g, _ := tender.ToUint64(greenArr.Value[i])
				b, _ := tender.ToUint64(blueArr.Value[i])
				ramp.Red[i] = uint16(r)
				ramp.Green[i] = uint16(g)
				ramp.Blue[i] = uint16(b)
			}
			m := (*glfw.Monitor)(unsafe.Pointer(uintptr(ptr.Value)))
			m.SetGammaRamp(ramp)
			return tender.NullValue, nil
		},
	},

	"set_monitor_gamma": &tender.NativeFunction{
		Name: "set_monitor_gamma",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			ptr, ok := args[0].(*tender.Int)
			if !ok || ptr.Value == 0 {
				return nil, tender.ErrInvalidArgument
			}
			gamma, okG := tender.ToFloat64(args[1])
			if !okG {
				return nil, tender.ErrInvalidArgument
			}
			m := (*glfw.Monitor)(unsafe.Pointer(uintptr(ptr.Value)))
			m.SetGamma(float32(gamma))
			return tender.NullValue, nil
		},
	},

	// ================================================================
	// ADDITIONAL EVENT CALLBACKS
	// ================================================================

	"set_cursor_enter_callback": &tender.NativeFunction{
		Name:      "set_cursor_enter_callback",
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
			win.SetCursorEnterCallback(func(_ *glfw.Window, entered bool) {
				if _, err := tender.WrapFuncCall(vm, cb,
					tender.FromBool(entered),
				); err != nil {
					fmt.Println("GLFW cursor enter callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	"set_drop_callback": &tender.NativeFunction{
		Name:      "set_drop_callback",
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
			win.SetDropCallback(func(_ *glfw.Window, paths []string) {
				pathObjs := make([]tender.Object, len(paths))
				for i, p := range paths {
					pathObjs[i] = &tender.String{Value: p}
				}
				if _, err := tender.WrapFuncCall(vm, cb,
					&tender.Array{Value: pathObjs},
				); err != nil {
					fmt.Println("GLFW drop callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	"set_focus_callback": &tender.NativeFunction{
		Name:      "set_focus_callback",
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
			win.SetFocusCallback(func(_ *glfw.Window, focused bool) {
				if _, err := tender.WrapFuncCall(vm, cb,
					tender.FromBool(focused),
				); err != nil {
					fmt.Println("GLFW focus callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	"set_iconify_callback": &tender.NativeFunction{
		Name:      "set_iconify_callback",
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
			win.SetIconifyCallback(func(_ *glfw.Window, iconified bool) {
				if _, err := tender.WrapFuncCall(vm, cb,
					tender.FromBool(iconified),
				); err != nil {
					fmt.Println("GLFW iconify callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	"set_maximize_callback": &tender.NativeFunction{
		Name:      "set_maximize_callback",
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
			win.SetMaximizeCallback(func(_ *glfw.Window, maximized bool) {
				if _, err := tender.WrapFuncCall(vm, cb,
					tender.FromBool(maximized),
				); err != nil {
					fmt.Println("GLFW maximize callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	"set_content_scale_callback": &tender.NativeFunction{
		Name:      "set_content_scale_callback",
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
			win.SetContentScaleCallback(func(_ *glfw.Window, xscale, yscale float32) {
				if _, err := tender.WrapFuncCall(vm, cb,
					&tender.Float{Value: float64(xscale)},
					&tender.Float{Value: float64(yscale)},
				); err != nil {
					fmt.Println("GLFW content scale callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	"set_pos_callback": &tender.NativeFunction{
		Name:      "set_pos_callback",
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
			win.SetPosCallback(func(_ *glfw.Window, xpos, ypos int) {
				if _, err := tender.WrapFuncCall(vm, cb,
					&tender.Int{Value: int64(xpos)},
					&tender.Int{Value: int64(ypos)},
				); err != nil {
					fmt.Println("GLFW pos callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	"set_refresh_callback": &tender.NativeFunction{
		Name:      "set_refresh_callback",
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
			win.SetRefreshCallback(func(_ *glfw.Window) {
				if _, err := tender.WrapFuncCall(vm, cb); err != nil {
					fmt.Println("GLFW refresh callback error:", err)
				}
			})
			return tender.NullValue, nil
		},
	},

	// ================================================================
	// UTILITY FUNCTIONS
	// ================================================================

	"post_empty_event": &tender.NativeFunction{
		Name: "post_empty_event",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glfw.PostEmptyEvent()
			return tender.NullValue, nil
		},
	},

	"raw_mouse_motion_supported": &tender.NativeFunction{
		Name: "raw_mouse_motion_supported",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			return tender.FromBool(glfw.RawMouseMotionSupported()), nil
		},
	},

	"get_key_name": &tender.NativeFunction{
		Name: "get_key_name",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			key, okK := tender.ToInt(args[0])
			scancode, okS := tender.ToInt(args[1])
			if !okK || !okS {
				return nil, tender.ErrInvalidArgument
			}
			name := glfw.GetKeyName(glfw.Key(key), scancode)
			if name == "" {
				return tender.NullValue, nil
			}
			return &tender.String{Value: name}, nil
		},
	},

	"get_key_scancode": &tender.NativeFunction{
		Name: "get_key_scancode",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			key, ok := tender.ToInt(args[0])
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			return &tender.Int{Value: int64(glfw.GetKeyScancode(glfw.Key(key)))}, nil
		},
	},	
	
}