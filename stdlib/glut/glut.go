package glut

import (
	"errors"
	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/v/glut"
)

var (
	ErrInvalidArgCount     = errors.New("invalid number of arguments")
	ErrInvalidArgumentType = errors.New("invalid argument type")
	ErrNotCallable         = errors.New("runtime error: object is not callable")
)

// Package-level registry to hold references to state for native CGo callbacks
var (
	runningVM       *tender.VM
	displayCallback tender.Object
	reshapeCallback tender.Object
	keyboardCallback tender.Object
	mouseCallback tender.Object
	motionCallback tender.Object
)

// Safe wrapper that uses WrapFuncCall properly
func invokeTenderFunc(vm *tender.VM, fn tender.Object, args ...tender.Object) {
	if fn == nil || vm == nil {
		return
	}
	
	// Use WrapFuncCall with the function and arguments
	// WrapFuncCall takes (vm, args...) where args[0] is the function
	allArgs := append([]tender.Object{fn}, args...)
	_, _ = tender.WrapFuncCall(vm, allArgs...)
}

// Native callbacks routed out to Tender runtime configurations
func nativeDisplay() {
	if displayCallback != nil && runningVM != nil {
		invokeTenderFunc(runningVM, displayCallback)
		glut.SwapBuffers()
	}
}

func nativeReshape(w, h int) {
	if reshapeCallback != nil && runningVM != nil {
		wObj := &tender.Int{Value: int64(w)}
		hObj := &tender.Int{Value: int64(h)}
		invokeTenderFunc(runningVM, reshapeCallback, wObj, hObj)
	}
}

func nativeKeyboard(key uint8, x, y int) {
	if keyboardCallback != nil && runningVM != nil {
		keyObj := &tender.Int{Value: int64(key)}
		xObj := &tender.Int{Value: int64(x)}
		yObj := &tender.Int{Value: int64(y)}
		invokeTenderFunc(runningVM, keyboardCallback, keyObj, xObj, yObj)
	}
}

func nativeMouse(button, state, x, y int) {
	if mouseCallback != nil && runningVM != nil {
		buttonObj := &tender.Int{Value: int64(button)}
		stateObj := &tender.Int{Value: int64(state)}
		xObj := &tender.Int{Value: int64(x)}
		yObj := &tender.Int{Value: int64(y)}
		invokeTenderFunc(runningVM, mouseCallback, buttonObj, stateObj, xObj, yObj)
	}
}

func nativeMotion(x, y int) {
	if motionCallback != nil && runningVM != nil {
		xObj := &tender.Int{Value: int64(x)}
		yObj := &tender.Int{Value: int64(y)}
		invokeTenderFunc(runningVM, motionCallback, xObj, yObj)
	}
}

// Module exposes the GLUT lifecycle to Tender scripts
var Module = &tender.BuiltinModule{
	Attrs: map[string]tender.Object{
		// Constants
		"RGB":    &tender.Int{Value: int64(glut.RGB)},
		"RGBA":   &tender.Int{Value: int64(glut.RGBA)},
		"DOUBLE": &tender.Int{Value: int64(glut.DOUBLE)},
		"DEPTH":  &tender.Int{Value: int64(glut.DEPTH)},

		// System Functions
		"init": &tender.BuiltinFunction{
			Name: "init",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				glut.Init()
				return tender.NullValue, nil
			},
		},

		"initDisplayMode": &tender.BuiltinFunction{
			Name: "initDisplayMode",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				mode, ok := args[0].(*tender.Int)
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				glut.InitDisplayMode(uint(mode.Value))
				return tender.NullValue, nil
			},
		},

		"initWindowSize": &tender.BuiltinFunction{
			Name: "initWindowSize",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				w, okW := args[0].(*tender.Int)
				h, okH := args[1].(*tender.Int)
				if !okW || !okH {
					return nil, ErrInvalidArgumentType
				}
				glut.InitWindowSize(int(w.Value), int(h.Value))
				return tender.NullValue, nil
			},
		},

		"initWindowPosition": &tender.BuiltinFunction{
			Name: "initWindowPosition",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				x, okX := args[0].(*tender.Int)
				y, okY := args[1].(*tender.Int)
				if !okX || !okY {
					return nil, ErrInvalidArgumentType
				}
				glut.InitWindowPosition(int(x.Value), int(y.Value))
				return tender.NullValue, nil
			},
		},

		"createWindow": &tender.BuiltinFunction{
			Name: "createWindow",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				title, ok := args[0].(*tender.String)
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				id := glut.CreateWindow(title.Value)
				return &tender.Int{Value: int64(id)}, nil
			},
		},

		"destroyWindow": &tender.BuiltinFunction{
			Name: "destroyWindow",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				win, ok := args[0].(*tender.Int)
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				glut.DestroyWindow(int(win.Value))
				return tender.NullValue, nil
			},
		},

		"getWindow": &tender.BuiltinFunction{
			Name: "getWindow",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				id := glut.GetWindow()
				return &tender.Int{Value: int64(id)}, nil
			},
		},

		"setWindow": &tender.BuiltinFunction{
			Name: "setWindow",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				win, ok := args[0].(*tender.Int)
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				glut.SetWindow(int(win.Value))
				return tender.NullValue, nil
			},
		},

		// Callback Configuration Hooking into VM Runtime State
		"displayFunc": &tender.BuiltinFunction{
			Name: "displayFunc",
			NeedVMObj: true,
			Value: func(args ...tender.Object) (tender.Object, error) {
				vm := args[0].(*tender.VMObj).Value
				args = args[1:]
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				
				// Check if the object is callable using CanCall()
				if args[0] != tender.NullValue && !args[0].CanCall() {
					return nil, ErrNotCallable
				}

				glut.DisplayFunc(func(){
					tender.WrapFuncCall(vm, args[0])
				})
				return tender.NullValue, nil
			},
		},

		"reshapeFunc": &tender.BuiltinFunction{
			Name: "reshapeFunc",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}

				// Check if the object is callable using CanCall()
				if args[0] != tender.NullValue && !args[0].CanCall() {
					return nil, ErrNotCallable
				}

				reshapeCallback = args[0]
				glut.ReshapeFunc(nativeReshape)
				return tender.NullValue, nil
			},
		},

		"keyboardFunc": &tender.BuiltinFunction{
			Name: "keyboardFunc",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}

				if args[0] != tender.NullValue && !args[0].CanCall() {
					return nil, ErrNotCallable
				}

				keyboardCallback = args[0]
				glut.KeyboardFunc(nativeKeyboard)
				return tender.NullValue, nil
			},
		},

		"mouseFunc": &tender.BuiltinFunction{
			Name: "mouseFunc",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}

				if args[0] != tender.NullValue && !args[0].CanCall() {
					return nil, ErrNotCallable
				}

				mouseCallback = args[0]
				glut.MouseFunc(nativeMouse)
				return tender.NullValue, nil
			},
		},

		"motionFunc": &tender.BuiltinFunction{
			Name: "motionFunc",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}

				if args[0] != tender.NullValue && !args[0].CanCall() {
					return nil, ErrNotCallable
				}

				motionCallback = args[0]
				glut.MotionFunc(nativeMotion)
				return tender.NullValue, nil
			},
		},

		"postRedisplay": &tender.BuiltinFunction{
			Name: "postRedisplay",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				glut.PostRedisplay()
				return tender.NullValue, nil
			},
		},

		"swapBuffers": &tender.BuiltinFunction{
			Name: "swapBuffers",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				glut.SwapBuffers()
				return tender.NullValue, nil
			},
		},

		"mainLoop": &tender.BuiltinFunction{
			Name: "mainLoop",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				glut.MainLoop()
				return tender.NullValue, nil
			},
		},

		// Keyboard modifier constants
		"ACTIVE_SHIFT": &tender.Int{Value: int64(glut.ACTIVE_SHIFT)},
		"ACTIVE_CTRL":  &tender.Int{Value: int64(glut.ACTIVE_CTRL)},
		"ACTIVE_ALT":   &tender.Int{Value: int64(glut.ACTIVE_ALT)},

		"getModifiers": &tender.BuiltinFunction{
			Name: "getModifiers",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				mod := glut.GetModifiers()
				return &tender.Int{Value: int64(mod)}, nil
			},
		},

		// Menu functions
		"createMenu": &tender.BuiltinFunction{
			Name: "createMenu",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}

				if args[0] != tender.NullValue && !args[0].CanCall() {
					return nil, ErrNotCallable
				}

				// Store callback in a closure
				cb := args[0]
				id := glut.CreateMenu(func(value int) {
					if runningVM != nil && cb != nil {
						valObj := &tender.Int{Value: int64(value)}
						invokeTenderFunc(runningVM, cb, valObj)
					}
				})
				return &tender.Int{Value: int64(id)}, nil
			},
		},

		"addMenuEntry": &tender.BuiltinFunction{
			Name: "addMenuEntry",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				name, ok := args[0].(*tender.String)
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				value, ok := args[1].(*tender.Int)
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				glut.AddMenuEntry(name.Value, int(value.Value))
				return tender.NullValue, nil
			},
		},

		"addSubMenu": &tender.BuiltinFunction{
			Name: "addSubMenu",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				name, ok := args[0].(*tender.String)
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				menuId, ok := args[1].(*tender.Int)
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				glut.AddSubMenu(name.Value, int(menuId.Value))
				return tender.NullValue, nil
			},
		},

		"attachMenu": &tender.BuiltinFunction{
			Name: "attachMenu",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				button, ok := args[0].(*tender.Int)
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				glut.AttachMenu(int(button.Value))
				return tender.NullValue, nil
			},
		},

		// Mouse button constants
		"LEFT_BUTTON":   &tender.Int{Value: int64(glut.LEFT_BUTTON)},
		"MIDDLE_BUTTON": &tender.Int{Value: int64(glut.MIDDLE_BUTTON)},
		"RIGHT_BUTTON":  &tender.Int{Value: int64(glut.RIGHT_BUTTON)},

		// Display mode constants
		"SINGLE": &tender.Int{Value: int64(glut.SINGLE)},
		"INDEX":  &tender.Int{Value: int64(glut.INDEX)},
		"STENCIL": &tender.Int{Value: int64(glut.STENCIL)},
		"ACCUM":   &tender.Int{Value: int64(glut.ACCUM)},

		// Window operations
		"positionWindow": &tender.BuiltinFunction{
			Name: "positionWindow",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				x, okX := args[0].(*tender.Int)
				y, okY := args[1].(*tender.Int)
				if !okX || !okY {
					return nil, ErrInvalidArgumentType
				}
				glut.PositionWindow(int(x.Value), int(y.Value))
				return tender.NullValue, nil
			},
		},

		"reshapeWindow": &tender.BuiltinFunction{
			Name: "reshapeWindow",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 2 {
					return nil, ErrInvalidArgCount
				}
				w, okW := args[0].(*tender.Int)
				h, okH := args[1].(*tender.Int)
				if !okW || !okH {
					return nil, ErrInvalidArgumentType
				}
				glut.ReshapeWindow(int(w.Value), int(h.Value))
				return tender.NullValue, nil
			},
		},

		"fullScreen": &tender.BuiltinFunction{
			Name: "fullScreen",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				glut.FullScreen()
				return tender.NullValue, nil
			},
		},

		"hideWindow": &tender.BuiltinFunction{
			Name: "hideWindow",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				glut.HideWindow()
				return tender.NullValue, nil
			},
		},

		"showWindow": &tender.BuiltinFunction{
			Name: "showWindow",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 0 {
					return nil, ErrInvalidArgCount
				}
				glut.ShowWindow()
				return tender.NullValue, nil
			},
		},

		"setWindowTitle": &tender.BuiltinFunction{
			Name: "setWindowTitle",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				title, ok := args[0].(*tender.String)
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				glut.SetWindowTitle(title.Value)
				return tender.NullValue, nil
			},
		},

		// Solid and Wire shapes
		"solidSphere": &tender.BuiltinFunction{
			Name: "solidSphere",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 3 {
					return nil, ErrInvalidArgCount
				}
				radius, okR := args[0].(*tender.Float)
				slices, okS := args[1].(*tender.Int)
				stacks, okT := args[2].(*tender.Int)
				if !okR || !okS || !okT {
					return nil, ErrInvalidArgumentType
				}
				glut.SolidSphere(radius.Value, int(slices.Value), int(stacks.Value))
				return tender.NullValue, nil
			},
		},

		"wireSphere": &tender.BuiltinFunction{
			Name: "wireSphere",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 3 {
					return nil, ErrInvalidArgCount
				}
				radius, okR := args[0].(*tender.Float)
				slices, okS := args[1].(*tender.Int)
				stacks, okT := args[2].(*tender.Int)
				if !okR || !okS || !okT {
					return nil, ErrInvalidArgumentType
				}
				glut.WireSphere(radius.Value, int(slices.Value), int(stacks.Value))
				return tender.NullValue, nil
			},
		},

		"solidCube": &tender.BuiltinFunction{
			Name: "solidCube",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				size, ok := args[0].(*tender.Float)
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				glut.SolidCube(size.Value)
				return tender.NullValue, nil
			},
		},

		"wireCube": &tender.BuiltinFunction{
			Name: "wireCube",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				size, ok := args[0].(*tender.Float)
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				glut.WireCube(size.Value)
				return tender.NullValue, nil
			},
		},

		"solidTeapot": &tender.BuiltinFunction{
			Name: "solidTeapot",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				size, ok := args[0].(*tender.Float)
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				glut.SolidTeapot(size.Value)
				return tender.NullValue, nil
			},
		},

		"wireTeapot": &tender.BuiltinFunction{
			Name: "wireTeapot",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 1 {
					return nil, ErrInvalidArgCount
				}
				size, ok := args[0].(*tender.Float)
				if !ok {
					return nil, ErrInvalidArgumentType
				}
				glut.WireTeapot(size.Value)
				return tender.NullValue, nil
			},
		},

		"solidCone": &tender.BuiltinFunction{
			Name: "solidCone",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, ErrInvalidArgCount
				}
				base, okB := args[0].(*tender.Float)
				height, okH := args[1].(*tender.Float)
				slices, okS := args[2].(*tender.Int)
				stacks, okT := args[3].(*tender.Int)
				if !okB || !okH || !okS || !okT {
					return nil, ErrInvalidArgumentType
				}
				glut.SolidCone(base.Value, height.Value, int(slices.Value), int(stacks.Value))
				return tender.NullValue, nil
			},
		},

		"wireCone": &tender.BuiltinFunction{
			Name: "wireCone",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, ErrInvalidArgCount
				}
				base, okB := args[0].(*tender.Float)
				height, okH := args[1].(*tender.Float)
				slices, okS := args[2].(*tender.Int)
				stacks, okT := args[3].(*tender.Int)
				if !okB || !okH || !okS || !okT {
					return nil, ErrInvalidArgumentType
				}
				glut.WireCone(base.Value, height.Value, int(slices.Value), int(stacks.Value))
				return tender.NullValue, nil
			},
		},

		"solidTorus": &tender.BuiltinFunction{
			Name: "solidTorus",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, ErrInvalidArgCount
				}
				inner, okI := args[0].(*tender.Float)
				outer, okO := args[1].(*tender.Float)
				nsides, okN := args[2].(*tender.Int)
				rings, okR := args[3].(*tender.Int)
				if !okI || !okO || !okN || !okR {
					return nil, ErrInvalidArgumentType
				}
				glut.SolidTorus(inner.Value, outer.Value, int(nsides.Value), int(rings.Value))
				return tender.NullValue, nil
			},
		},

		"wireTorus": &tender.BuiltinFunction{
			Name: "wireTorus",
			Value: func(args ...tender.Object) (tender.Object, error) {
				if len(args) != 4 {
					return nil, ErrInvalidArgCount
				}
				inner, okI := args[0].(*tender.Float)
				outer, okO := args[1].(*tender.Float)
				nsides, okN := args[2].(*tender.Int)
				rings, okR := args[3].(*tender.Int)
				if !okI || !okO || !okN || !okR {
					return nil, ErrInvalidArgumentType
				}
				glut.WireTorus(inner.Value, outer.Value, int(nsides.Value), int(rings.Value))
				return tender.NullValue, nil
			},
		},
	},
}

// SetVMContext allows your compiler/runner to hook the active execution state 
// down to the stdlib registry before launching scripts or entering the main loop.
func SetVMContext(vm *tender.VM) {
	runningVM = vm
}