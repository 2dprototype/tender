package stdlib

import (
	"github.com/2dprototype/tender"
	"github.com/2dprototype/tender/v/glut"
)

// Module exposes the GLUT lifecycle to Tender scripts
var glutModule = map[string]tender.Object{
	// ============================================================
	// DISPLAY MODE CONSTANTS
	// ============================================================
	"RGB":         &tender.Int{Value: int64(glut.RGB)},
	"RGBA":        &tender.Int{Value: int64(glut.RGBA)},
	"INDEX":       &tender.Int{Value: int64(glut.INDEX)},
	"SINGLE":      &tender.Int{Value: int64(glut.SINGLE)},
	"DOUBLE":      &tender.Int{Value: int64(glut.DOUBLE)},
	"ACCUM":       &tender.Int{Value: int64(glut.ACCUM)},
	"ALPHA":       &tender.Int{Value: int64(glut.ALPHA)},
	"DEPTH":       &tender.Int{Value: int64(glut.DEPTH)},
	"STENCIL":     &tender.Int{Value: int64(glut.STENCIL)},
	"MULTISAMPLE": &tender.Int{Value: int64(glut.MULTISAMPLE)},
	"STEREO":      &tender.Int{Value: int64(glut.STEREO)},
	"LUMINANCE":   &tender.Int{Value: int64(glut.LUMINANCE)},

	// ============================================================
	// MOUSE BUTTON CONSTANTS
	// ============================================================
	"LEFT_BUTTON":   &tender.Int{Value: int64(glut.LEFT_BUTTON)},
	"MIDDLE_BUTTON": &tender.Int{Value: int64(glut.MIDDLE_BUTTON)},
	"RIGHT_BUTTON":  &tender.Int{Value: int64(glut.RIGHT_BUTTON)},
	"DOWN":          &tender.Int{Value: int64(glut.DOWN)},
	"UP":            &tender.Int{Value: int64(glut.UP)},

	// ============================================================
	// SPECIAL KEY CONSTANTS
	// ============================================================
	"KEY_F1":        &tender.Int{Value: int64(glut.KEY_F1)},
	"KEY_F2":        &tender.Int{Value: int64(glut.KEY_F2)},
	"KEY_F3":        &tender.Int{Value: int64(glut.KEY_F3)},
	"KEY_F4":        &tender.Int{Value: int64(glut.KEY_F4)},
	"KEY_F5":        &tender.Int{Value: int64(glut.KEY_F5)},
	"KEY_F6":        &tender.Int{Value: int64(glut.KEY_F6)},
	"KEY_F7":        &tender.Int{Value: int64(glut.KEY_F7)},
	"KEY_F8":        &tender.Int{Value: int64(glut.KEY_F8)},
	"KEY_F9":        &tender.Int{Value: int64(glut.KEY_F9)},
	"KEY_F10":       &tender.Int{Value: int64(glut.KEY_F10)},
	"KEY_F11":       &tender.Int{Value: int64(glut.KEY_F11)},
	"KEY_F12":       &tender.Int{Value: int64(glut.KEY_F12)},
	"KEY_LEFT":      &tender.Int{Value: int64(glut.KEY_LEFT)},
	"KEY_UP":        &tender.Int{Value: int64(glut.KEY_UP)},
	"KEY_RIGHT":     &tender.Int{Value: int64(glut.KEY_RIGHT)},
	"KEY_DOWN":      &tender.Int{Value: int64(glut.KEY_DOWN)},
	"KEY_PAGE_UP":   &tender.Int{Value: int64(glut.KEY_PAGE_UP)},
	"KEY_PAGE_DOWN": &tender.Int{Value: int64(glut.KEY_PAGE_DOWN)},
	"KEY_HOME":      &tender.Int{Value: int64(glut.KEY_HOME)},
	"KEY_END":       &tender.Int{Value: int64(glut.KEY_END)},
	"KEY_INSERT":    &tender.Int{Value: int64(glut.KEY_INSERT)},

	// ============================================================
	// ENTRY/EXIT CONSTANTS
	// ============================================================
	"LEFT":    &tender.Int{Value: int64(glut.LEFT)},
	"ENTERED": &tender.Int{Value: int64(glut.ENTERED)},

	// ============================================================
	// WINDOW/MENU STATUS CONSTANTS
	// ============================================================
	"MENU_NOT_IN_USE":    &tender.Int{Value: int64(glut.MENU_NOT_IN_USE)},
	"MENU_IN_USE":        &tender.Int{Value: int64(glut.MENU_IN_USE)},
	"NOT_VISIBLE":        &tender.Int{Value: int64(glut.NOT_VISIBLE)},
	"VISIBLE":            &tender.Int{Value: int64(glut.VISIBLE)},
	"HIDDEN":             &tender.Int{Value: int64(glut.HIDDEN)},
	"FULLY_RETAINED":     &tender.Int{Value: int64(glut.FULLY_RETAINED)},
	"PARTIALLY_RETAINED": &tender.Int{Value: int64(glut.PARTIALLY_RETAINED)},
	"FULLY_COVERED":      &tender.Int{Value: int64(glut.FULLY_COVERED)},

	// ============================================================
	// RGB COLOR COMPONENT CONSTANTS
	// ============================================================
	"RED":   &tender.Int{Value: int64(glut.RED)},
	"GREEN": &tender.Int{Value: int64(glut.GREEN)},
	"BLUE":  &tender.Int{Value: int64(glut.BLUE)},

	// ============================================================
	// LAYER CONSTANTS
	// ============================================================
	"NORMAL":  &tender.Int{Value: int64(glut.NORMAL)},
	"OVERLAY": &tender.Int{Value: int64(glut.OVERLAY)},

	// ============================================================
	// FONT CONSTANTS
	// ============================================================
	"STROKE_ROMAN":          &tender.Int{Value: int64(glut.STROKE_ROMAN)},
	"STROKE_MONO_ROMAN":     &tender.Int{Value: int64(glut.STROKE_MONO_ROMAN)},
	"BITMAP_9_BY_15":        &tender.Int{Value: int64(glut.BITMAP_9_BY_15)},
	"BITMAP_8_BY_13":        &tender.Int{Value: int64(glut.BITMAP_8_BY_13)},
	"BITMAP_TIMES_ROMAN_10": &tender.Int{Value: int64(glut.BITMAP_TIMES_ROMAN_10)},
	"BITMAP_TIMES_ROMAN_24": &tender.Int{Value: int64(glut.BITMAP_TIMES_ROMAN_24)},
	"BITMAP_HELVETICA_10":   &tender.Int{Value: int64(glut.BITMAP_HELVETICA_10)},
	"BITMAP_HELVETICA_12":   &tender.Int{Value: int64(glut.BITMAP_HELVETICA_12)},
	"BITMAP_HELVETICA_18":   &tender.Int{Value: int64(glut.BITMAP_HELVETICA_18)},

	// ============================================================
	// GET PARAMETER CONSTANTS
	// ============================================================
	"WINDOW_X":                 &tender.Int{Value: int64(glut.WINDOW_X)},
	"WINDOW_Y":                 &tender.Int{Value: int64(glut.WINDOW_Y)},
	"WINDOW_WIDTH":             &tender.Int{Value: int64(glut.WINDOW_WIDTH)},
	"WINDOW_HEIGHT":            &tender.Int{Value: int64(glut.WINDOW_HEIGHT)},
	"WINDOW_BUFFER_SIZE":       &tender.Int{Value: int64(glut.WINDOW_BUFFER_SIZE)},
	"WINDOW_STENCIL_SIZE":      &tender.Int{Value: int64(glut.WINDOW_STENCIL_SIZE)},
	"WINDOW_DEPTH_SIZE":        &tender.Int{Value: int64(glut.WINDOW_DEPTH_SIZE)},
	"WINDOW_RED_SIZE":          &tender.Int{Value: int64(glut.WINDOW_RED_SIZE)},
	"WINDOW_GREEN_SIZE":        &tender.Int{Value: int64(glut.WINDOW_GREEN_SIZE)},
	"WINDOW_BLUE_SIZE":         &tender.Int{Value: int64(glut.WINDOW_BLUE_SIZE)},
	"WINDOW_ALPHA_SIZE":        &tender.Int{Value: int64(glut.WINDOW_ALPHA_SIZE)},
	"WINDOW_ACCUM_RED_SIZE":    &tender.Int{Value: int64(glut.WINDOW_ACCUM_RED_SIZE)},
	"WINDOW_ACCUM_GREEN_SIZE":  &tender.Int{Value: int64(glut.WINDOW_ACCUM_GREEN_SIZE)},
	"WINDOW_ACCUM_BLUE_SIZE":   &tender.Int{Value: int64(glut.WINDOW_ACCUM_BLUE_SIZE)},
	"WINDOW_ACCUM_ALPHA_SIZE":  &tender.Int{Value: int64(glut.WINDOW_ACCUM_ALPHA_SIZE)},
	"WINDOW_DOUBLEBUFFER":      &tender.Int{Value: int64(glut.WINDOW_DOUBLEBUFFER)},
	"WINDOW_RGBA":              &tender.Int{Value: int64(glut.WINDOW_RGBA)},
	"WINDOW_PARENT":            &tender.Int{Value: int64(glut.WINDOW_PARENT)},
	"WINDOW_NUM_CHILDREN":      &tender.Int{Value: int64(glut.WINDOW_NUM_CHILDREN)},
	"WINDOW_COLORMAP_SIZE":     &tender.Int{Value: int64(glut.WINDOW_COLORMAP_SIZE)},
	"WINDOW_NUM_SAMPLES":       &tender.Int{Value: int64(glut.WINDOW_NUM_SAMPLES)},
	"WINDOW_STEREO":            &tender.Int{Value: int64(glut.WINDOW_STEREO)},
	"WINDOW_CURSOR":            &tender.Int{Value: int64(glut.WINDOW_CURSOR)},
	"WINDOW_FORMAT_ID":         &tender.Int{Value: int64(glut.WINDOW_FORMAT_ID)},
	"SCREEN_WIDTH":             &tender.Int{Value: int64(glut.SCREEN_WIDTH)},
	"SCREEN_HEIGHT":            &tender.Int{Value: int64(glut.SCREEN_HEIGHT)},
	"SCREEN_WIDTH_MM":          &tender.Int{Value: int64(glut.SCREEN_WIDTH_MM)},
	"SCREEN_HEIGHT_MM":         &tender.Int{Value: int64(glut.SCREEN_HEIGHT_MM)},
	"MENU_NUM_ITEMS":           &tender.Int{Value: int64(glut.MENU_NUM_ITEMS)},
	"DISPLAY_MODE_POSSIBLE":    &tender.Int{Value: int64(glut.DISPLAY_MODE_POSSIBLE)},
	"INIT_WINDOW_X":            &tender.Int{Value: int64(glut.INIT_WINDOW_X)},
	"INIT_WINDOW_Y":            &tender.Int{Value: int64(glut.INIT_WINDOW_Y)},
	"INIT_WINDOW_WIDTH":        &tender.Int{Value: int64(glut.INIT_WINDOW_WIDTH)},
	"INIT_WINDOW_HEIGHT":       &tender.Int{Value: int64(glut.INIT_WINDOW_HEIGHT)},
	"INIT_DISPLAY_MODE":        &tender.Int{Value: int64(glut.INIT_DISPLAY_MODE)},
	"ELAPSED_TIME":             &tender.Int{Value: int64(glut.ELAPSED_TIME)},

	// ============================================================
	// DEVICE GET PARAMETERS
	// ============================================================
	"HAS_KEYBOARD":            &tender.Int{Value: int64(glut.HAS_KEYBOARD)},
	"HAS_MOUSE":               &tender.Int{Value: int64(glut.HAS_MOUSE)},
	"HAS_SPACEBALL":           &tender.Int{Value: int64(glut.HAS_SPACEBALL)},
	"HAS_DIAL_AND_BUTTON_BOX": &tender.Int{Value: int64(glut.HAS_DIAL_AND_BUTTON_BOX)},
	"HAS_TABLET":              &tender.Int{Value: int64(glut.HAS_TABLET)},
	"HAS_JOYSTICK":            &tender.Int{Value: int64(glut.HAS_JOYSTICK)},
	"OWNS_JOYSTICK":           &tender.Int{Value: int64(glut.OWNS_JOYSTICK)},
	"NUM_MOUSE_BUTTONS":       &tender.Int{Value: int64(glut.NUM_MOUSE_BUTTONS)},
	"NUM_SPACEBALL_BUTTONS":   &tender.Int{Value: int64(glut.NUM_SPACEBALL_BUTTONS)},
	"NUM_BUTTON_BOX_BUTTONS":  &tender.Int{Value: int64(glut.NUM_BUTTON_BOX_BUTTONS)},
	"NUM_DIALS":               &tender.Int{Value: int64(glut.NUM_DIALS)},
	"NUM_TABLET_BUTTONS":      &tender.Int{Value: int64(glut.NUM_TABLET_BUTTONS)},
	"JOYSTICK_BUTTONS":        &tender.Int{Value: int64(glut.JOYSTICK_BUTTONS)},
	"JOYSTICK_AXES":           &tender.Int{Value: int64(glut.JOYSTICK_AXES)},
	"JOYSTICK_POLL_RATE":      &tender.Int{Value: int64(glut.JOYSTICK_POLL_RATE)},
	"DEVICE_IGNORE_KEY_REPEAT": &tender.Int{Value: int64(glut.DEVICE_IGNORE_KEY_REPEAT)},
	"DEVICE_KEY_REPEAT":        &tender.Int{Value: int64(glut.DEVICE_KEY_REPEAT)},

	// ============================================================
	// LAYER GET PARAMETERS
	// ============================================================
	"OVERLAY_POSSIBLE":  &tender.Int{Value: int64(glut.OVERLAY_POSSIBLE)},
	"LAYER_IN_USE":      &tender.Int{Value: int64(glut.LAYER_IN_USE)},
	"HAS_OVERLAY":       &tender.Int{Value: int64(glut.HAS_OVERLAY)},
	"TRANSPARENT_INDEX": &tender.Int{Value: int64(glut.TRANSPARENT_INDEX)},
	"NORMAL_DAMAGED":    &tender.Int{Value: int64(glut.NORMAL_DAMAGED)},
	"OVERLAY_DAMAGED":   &tender.Int{Value: int64(glut.OVERLAY_DAMAGED)},

	// ============================================================
	// VIDEO RESIZE PARAMETERS
	// ============================================================
	"VIDEO_RESIZE_POSSIBLE":     &tender.Int{Value: int64(glut.VIDEO_RESIZE_POSSIBLE)},
	"VIDEO_RESIZE_IN_USE":       &tender.Int{Value: int64(glut.VIDEO_RESIZE_IN_USE)},
	"VIDEO_RESIZE_X_DELTA":      &tender.Int{Value: int64(glut.VIDEO_RESIZE_X_DELTA)},
	"VIDEO_RESIZE_Y_DELTA":      &tender.Int{Value: int64(glut.VIDEO_RESIZE_Y_DELTA)},
	"VIDEO_RESIZE_WIDTH_DELTA":  &tender.Int{Value: int64(glut.VIDEO_RESIZE_WIDTH_DELTA)},
	"VIDEO_RESIZE_HEIGHT_DELTA": &tender.Int{Value: int64(glut.VIDEO_RESIZE_HEIGHT_DELTA)},
	"VIDEO_RESIZE_X":            &tender.Int{Value: int64(glut.VIDEO_RESIZE_X)},
	"VIDEO_RESIZE_Y":            &tender.Int{Value: int64(glut.VIDEO_RESIZE_Y)},
	"VIDEO_RESIZE_WIDTH":        &tender.Int{Value: int64(glut.VIDEO_RESIZE_WIDTH)},
	"VIDEO_RESIZE_HEIGHT":       &tender.Int{Value: int64(glut.VIDEO_RESIZE_HEIGHT)},

	// ============================================================
	// MODIFIER CONSTANTS
	// ============================================================
	"ACTIVE_SHIFT": &tender.Int{Value: int64(glut.ACTIVE_SHIFT)},
	"ACTIVE_CTRL":  &tender.Int{Value: int64(glut.ACTIVE_CTRL)},
	"ACTIVE_ALT":   &tender.Int{Value: int64(glut.ACTIVE_ALT)},

	// ============================================================
	// CURSOR CONSTANTS
	// ============================================================
	"CURSOR_RIGHT_ARROW":         &tender.Int{Value: int64(glut.CURSOR_RIGHT_ARROW)},
	"CURSOR_LEFT_ARROW":          &tender.Int{Value: int64(glut.CURSOR_LEFT_ARROW)},
	"CURSOR_INFO":                &tender.Int{Value: int64(glut.CURSOR_INFO)},
	"CURSOR_DESTROY":             &tender.Int{Value: int64(glut.CURSOR_DESTROY)},
	"CURSOR_HELP":                &tender.Int{Value: int64(glut.CURSOR_HELP)},
	"CURSOR_CYCLE":               &tender.Int{Value: int64(glut.CURSOR_CYCLE)},
	"CURSOR_SPRAY":               &tender.Int{Value: int64(glut.CURSOR_SPRAY)},
	"CURSOR_WAIT":                &tender.Int{Value: int64(glut.CURSOR_WAIT)},
	"CURSOR_TEXT":                &tender.Int{Value: int64(glut.CURSOR_TEXT)},
	"CURSOR_CROSSHAIR":           &tender.Int{Value: int64(glut.CURSOR_CROSSHAIR)},
	"CURSOR_UP_DOWN":             &tender.Int{Value: int64(glut.CURSOR_UP_DOWN)},
	"CURSOR_LEFT_RIGHT":          &tender.Int{Value: int64(glut.CURSOR_LEFT_RIGHT)},
	"CURSOR_TOP_SIDE":            &tender.Int{Value: int64(glut.CURSOR_TOP_SIDE)},
	"CURSOR_BOTTOM_SIDE":         &tender.Int{Value: int64(glut.CURSOR_BOTTOM_SIDE)},
	"CURSOR_LEFT_SIDE":           &tender.Int{Value: int64(glut.CURSOR_LEFT_SIDE)},
	"CURSOR_RIGHT_SIDE":          &tender.Int{Value: int64(glut.CURSOR_RIGHT_SIDE)},
	"CURSOR_TOP_LEFT_CORNER":     &tender.Int{Value: int64(glut.CURSOR_TOP_LEFT_CORNER)},
	"CURSOR_TOP_RIGHT_CORNER":    &tender.Int{Value: int64(glut.CURSOR_TOP_RIGHT_CORNER)},
	"CURSOR_BOTTOM_RIGHT_CORNER": &tender.Int{Value: int64(glut.CURSOR_BOTTOM_RIGHT_CORNER)},
	"CURSOR_BOTTOM_LEFT_CORNER":  &tender.Int{Value: int64(glut.CURSOR_BOTTOM_LEFT_CORNER)},
	"CURSOR_INHERIT":             &tender.Int{Value: int64(glut.CURSOR_INHERIT)},
	"CURSOR_NONE":                &tender.Int{Value: int64(glut.CURSOR_NONE)},
	"CURSOR_FULL_CROSSHAIR":      &tender.Int{Value: int64(glut.CURSOR_FULL_CROSSHAIR)},

	// ============================================================
	// KEY REPEAT CONSTANTS
	// ============================================================
	"KEY_REPEAT_OFF":     &tender.Int{Value: int64(glut.KEY_REPEAT_OFF)},
	"KEY_REPEAT_ON":      &tender.Int{Value: int64(glut.KEY_REPEAT_ON)},
	"KEY_REPEAT_DEFAULT": &tender.Int{Value: int64(glut.KEY_REPEAT_DEFAULT)},

	// ============================================================
	// JOYSTICK BUTTON CONSTANTS
	// ============================================================
	"JOYSTICK_BUTTON_A": &tender.Int{Value: int64(glut.JOYSTICK_BUTTON_A)},
	"JOYSTICK_BUTTON_B": &tender.Int{Value: int64(glut.JOYSTICK_BUTTON_B)},
	"JOYSTICK_BUTTON_C": &tender.Int{Value: int64(glut.JOYSTICK_BUTTON_C)},
	"JOYSTICK_BUTTON_D": &tender.Int{Value: int64(glut.JOYSTICK_BUTTON_D)},

	// ============================================================
	// GAME MODE CONSTANTS
	// ============================================================
	"GAME_MODE_ACTIVE":          &tender.Int{Value: int64(glut.GAME_MODE_ACTIVE)},
	"GAME_MODE_POSSIBLE":        &tender.Int{Value: int64(glut.GAME_MODE_POSSIBLE)},
	"GAME_MODE_WIDTH":           &tender.Int{Value: int64(glut.GAME_MODE_WIDTH)},
	"GAME_MODE_HEIGHT":          &tender.Int{Value: int64(glut.GAME_MODE_HEIGHT)},
	"GAME_MODE_PIXEL_DEPTH":     &tender.Int{Value: int64(glut.GAME_MODE_PIXEL_DEPTH)},
	"GAME_MODE_REFRESH_RATE":    &tender.Int{Value: int64(glut.GAME_MODE_REFRESH_RATE)},
	"GAME_MODE_DISPLAY_CHANGED": &tender.Int{Value: int64(glut.GAME_MODE_DISPLAY_CHANGED)},

	// ============================================================
	// SYSTEM FUNCTIONS
	// ============================================================
	"init": &tender.BuiltinFunction{
		Name: "init",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.Init()
			return tender.NullValue, nil
		},
	},

	"init_display_mode": &tender.BuiltinFunction{
		Name: "init_display_mode",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			mode, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.InitDisplayMode(uint(mode.Value))
			return tender.NullValue, nil
		},
	},

	"init_window_size": &tender.BuiltinFunction{
		Name: "init_window_size",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			w, okW := args[0].(*tender.Int)
			h, okH := args[1].(*tender.Int)
			if !okW || !okH {
				return nil, tender.ErrInvalidArgument
			}
			glut.InitWindowSize(int(w.Value), int(h.Value))
			return tender.NullValue, nil
		},
	},

	"init_window_position": &tender.BuiltinFunction{
		Name: "init_window_position",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := args[0].(*tender.Int)
			y, okY := args[1].(*tender.Int)
			if !okX || !okY {
				return nil, tender.ErrInvalidArgument
			}
			glut.InitWindowPosition(int(x.Value), int(y.Value))
			return tender.NullValue, nil
		},
	},

	"create_window": &tender.BuiltinFunction{
		Name: "create_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			title, ok := args[0].(*tender.String)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			id := glut.CreateWindow(title.Value)
			return &tender.Int{Value: int64(id)}, nil
		},
	},

	"create_sub_window": &tender.BuiltinFunction{
		Name: "create_sub_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 5 {
				return nil, tender.ErrInvalidArgCount
			}
			winId, ok1 := args[0].(*tender.Int)
			x, ok2 := args[1].(*tender.Int)
			y, ok3 := args[2].(*tender.Int)
			w, ok4 := args[3].(*tender.Int)
			h, ok5 := args[4].(*tender.Int)
			if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 {
				return nil, tender.ErrInvalidArgument
			}
			id := glut.CreateSubWindow(int(winId.Value), int(x.Value), int(y.Value), int(w.Value), int(h.Value))
			return &tender.Int{Value: int64(id)}, nil
		},
	},

	"destroy_window": &tender.BuiltinFunction{
		Name: "destroy_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			win, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.DestroyWindow(int(win.Value))
			return tender.NullValue, nil
		},
	},

	"get_window": &tender.BuiltinFunction{
		Name: "get_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			id := glut.GetWindow()
			return &tender.Int{Value: int64(id)}, nil
		},
	},

	"set_window": &tender.BuiltinFunction{
		Name: "set_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			win, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.SetWindow(int(win.Value))
			return tender.NullValue, nil
		},
	},

	"set_window_title": &tender.BuiltinFunction{
		Name: "set_window_title",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			title, ok := args[0].(*tender.String)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.SetWindowTitle(title.Value)
			return tender.NullValue, nil
		},
	},

	"set_icon_title": &tender.BuiltinFunction{
		Name: "set_icon_title",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			title, ok := args[0].(*tender.String)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.SetIconTitle(title.Value)
			return tender.NullValue, nil
		},
	},

	"position_window": &tender.BuiltinFunction{
		Name: "position_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := args[0].(*tender.Int)
			y, okY := args[1].(*tender.Int)
			if !okX || !okY {
				return nil, tender.ErrInvalidArgument
			}
			glut.PositionWindow(int(x.Value), int(y.Value))
			return tender.NullValue, nil
		},
	},

	"reshape_window": &tender.BuiltinFunction{
		Name: "reshape_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			w, okW := args[0].(*tender.Int)
			h, okH := args[1].(*tender.Int)
			if !okW || !okH {
				return nil, tender.ErrInvalidArgument
			}
			glut.ReshapeWindow(int(w.Value), int(h.Value))
			return tender.NullValue, nil
		},
	},

	"full_screen": &tender.BuiltinFunction{
		Name: "full_screen",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.FullScreen()
			return tender.NullValue, nil
		},
	},

	"hide_window": &tender.BuiltinFunction{
		Name: "hide_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.HideWindow()
			return tender.NullValue, nil
		},
	},

	"show_window": &tender.BuiltinFunction{
		Name: "show_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.ShowWindow()
			return tender.NullValue, nil
		},
	},

	"iconify_window": &tender.BuiltinFunction{
		Name: "iconify_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.IconifyWindow()
			return tender.NullValue, nil
		},
	},

	"push_window": &tender.BuiltinFunction{
		Name: "push_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.PushWindow()
			return tender.NullValue, nil
		},
	},

	"pop_window": &tender.BuiltinFunction{
		Name: "pop_window",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.PopWindow()
			return tender.NullValue, nil
		},
	},

	"warp_pointer": &tender.BuiltinFunction{
		Name: "warp_pointer",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			x, okX := args[0].(*tender.Int)
			y, okY := args[1].(*tender.Int)
			if !okX || !okY {
				return nil, tender.ErrInvalidArgument
			}
			glut.WarpPointer(int(x.Value), int(y.Value))
			return tender.NullValue, nil
		},
	},

	"set_cursor": &tender.BuiltinFunction{
		Name: "set_cursor",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			cursor, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.SetCursor(int(cursor.Value))
			return tender.NullValue, nil
		},
	},

	// ============================================================
	// CALLBACK FUNCTIONS (VM-Aware)
	// ============================================================
	"display_func": &tender.BuiltinFunction{
		Name:      "display_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.DisplayFunc(func() {
				tender.WrapFuncCall(vm, args[0])
			})
			return tender.NullValue, nil
		},
	},

	"reshape_func": &tender.BuiltinFunction{
		Name:      "reshape_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.ReshapeFunc(func(w, h int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(w)},
					&tender.Int{Value: int64(h)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"keyboard_func": &tender.BuiltinFunction{
		Name:      "keyboard_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.KeyboardFunc(func(key byte, x, y int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.String{Value: string(key)},
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"keyboard_up_func": &tender.BuiltinFunction{
		Name:      "keyboard_up_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.KeyboardUpFunc(func(key byte, x, y int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.String{Value: string(key)},
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"special_func": &tender.BuiltinFunction{
		Name:      "special_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.SpecialFunc(func(key, x, y int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(key)},
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"special_up_func": &tender.BuiltinFunction{
		Name:      "special_up_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.SpecialUpFunc(func(key, x, y int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(key)},
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"mouse_func": &tender.BuiltinFunction{
		Name:      "mouse_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.MouseFunc(func(button, state, x, y int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(button)},
					&tender.Int{Value: int64(state)},
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"motion_func": &tender.BuiltinFunction{
		Name:      "motion_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.MotionFunc(func(x, y int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"passive_motion_func": &tender.BuiltinFunction{
		Name:      "passive_motion_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.PassiveMotionFunc(func(x, y int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"entry_func": &tender.BuiltinFunction{
		Name:      "entry_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.EntryFunc(func(state int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(state)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"visibility_func": &tender.BuiltinFunction{
		Name:      "visibility_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.VisibilityFunc(func(state int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(state)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"idle_func": &tender.BuiltinFunction{
		Name:      "idle_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.IdleFunc(func() {
				tender.WrapFuncCall(vm, args[0])
			})
			return tender.NullValue, nil
		},
	},

	"timer_func": &tender.BuiltinFunction{
		Name:      "timer_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			msecs, ok1 := args[0].(*tender.Int)
			if !ok1 {
				return nil, tender.ErrInvalidArgument
			}
			if args[1] != tender.NullValue && !args[1].CanCall() {
				return nil, tender.ErrNotCallable
			}
			timerId, ok3 := args[2].(*tender.Int)
			if !ok3 {
				return nil, tender.ErrInvalidArgument
			}
			cb := args[1]
			glut.TimerFunc(int(msecs.Value), func(id int) {
				tender.WrapFuncCall(vm, cb, &tender.Int{Value: int64(id)})
			}, int(timerId.Value))
			return tender.NullValue, nil
		},
	},

	"overlay_display_func": &tender.BuiltinFunction{
		Name:      "overlay_display_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.OverlayDisplayFunc(func() {
				tender.WrapFuncCall(vm, args[0])
			})
			return tender.NullValue, nil
		},
	},

	"spaceball_motion_func": &tender.BuiltinFunction{
		Name:      "spaceball_motion_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.SpaceballMotionFunc(func(x, y, z int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
					&tender.Int{Value: int64(z)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"spaceball_rotate_func": &tender.BuiltinFunction{
		Name:      "spaceball_rotate_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.SpaceballRotateFunc(func(x, y, z int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
					&tender.Int{Value: int64(z)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"spaceball_button_func": &tender.BuiltinFunction{
		Name:      "spaceball_button_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.SpaceballButtonFunc(func(button, state int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(button)},
					&tender.Int{Value: int64(state)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"button_box_func": &tender.BuiltinFunction{
		Name:      "button_box_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.ButtonBoxFunc(func(button, state int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(button)},
					&tender.Int{Value: int64(state)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"dials_func": &tender.BuiltinFunction{
		Name:      "dials_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.DialsFunc(func(dial, value int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(dial)},
					&tender.Int{Value: int64(value)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"tablet_motion_func": &tender.BuiltinFunction{
		Name:      "tablet_motion_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.TabletMotionFunc(func(x, y int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"tablet_button_func": &tender.BuiltinFunction{
		Name:      "tablet_button_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.TabletButtonFunc(func(button, state, x, y int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(button)},
					&tender.Int{Value: int64(state)},
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"joystick_func": &tender.BuiltinFunction{
		Name:      "joystick_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			pollInterval, ok := args[1].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.JoystickFunc(func(buttonMask uint, x, y, z int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(buttonMask)},
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
					&tender.Int{Value: int64(z)},
				)
			}, int(pollInterval.Value))
			return tender.NullValue, nil
		},
	},

	"window_status_func": &tender.BuiltinFunction{
		Name:      "window_status_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.WindowStatusFunc(func(state int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(state)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"menu_state_func": &tender.BuiltinFunction{
		Name:      "menu_state_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.MenuStateFunc(func(status int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(status)},
				)
			})
			return tender.NullValue, nil
		},
	},

	"menu_status_func": &tender.BuiltinFunction{
		Name:      "menu_status_func",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			glut.MenuStatusFunc(func(status, x, y int) {
				tender.WrapFuncCall(vm, args[0],
					&tender.Int{Value: int64(status)},
					&tender.Int{Value: int64(x)},
					&tender.Int{Value: int64(y)},
				)
			})
			return tender.NullValue, nil
		},
	},

	// ============================================================
	// MENU FUNCTIONS
	// ============================================================
	"create_menu": &tender.BuiltinFunction{
		Name:      "create_menu",
		NeedVMObj: true,
		Value: func(args ...tender.Object) (tender.Object, error) {
			vm := args[0].(*tender.VMObj).Value
			args = args[1:]
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			if args[0] != tender.NullValue && !args[0].CanCall() {
				return nil, tender.ErrNotCallable
			}
			cb := args[0]
			id := glut.CreateMenu(func(value int) {
				tender.WrapFuncCall(vm, cb, &tender.Int{Value: int64(value)})
			})
			return &tender.Int{Value: int64(id)}, nil
		},
	},

	"destroy_menu": &tender.BuiltinFunction{
		Name: "destroy_menu",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			menuId, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.DestroyMenu(int(menuId.Value))
			return tender.NullValue, nil
		},
	},

	"get_menu": &tender.BuiltinFunction{
		Name: "get_menu",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			id := glut.GetMenu()
			return &tender.Int{Value: int64(id)}, nil
		},
	},

	"set_menu": &tender.BuiltinFunction{
		Name: "set_menu",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			menuId, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.SetMenu(int(menuId.Value))
			return tender.NullValue, nil
		},
	},

	"add_menu_entry": &tender.BuiltinFunction{
		Name: "add_menu_entry",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			name, ok := args[0].(*tender.String)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			value, ok := args[1].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.AddMenuEntry(name.Value, int(value.Value))
			return tender.NullValue, nil
		},
	},

	"add_sub_menu": &tender.BuiltinFunction{
		Name: "add_sub_menu",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			name, ok := args[0].(*tender.String)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			menuId, ok := args[1].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.AddSubMenu(name.Value, int(menuId.Value))
			return tender.NullValue, nil
		},
	},

	"change_to_menu_entry": &tender.BuiltinFunction{
		Name: "change_to_menu_entry",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			entry, ok1 := args[0].(*tender.Int)
			name, ok2 := args[1].(*tender.String)
			value, ok3 := args[2].(*tender.Int)
			if !ok1 || !ok2 || !ok3 {
				return nil, tender.ErrInvalidArgument
			}
			glut.ChangeToMenuEntry(int(entry.Value), name.Value, int(value.Value))
			return tender.NullValue, nil
		},
	},

	"change_to_sub_menu": &tender.BuiltinFunction{
		Name: "change_to_sub_menu",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			entry, ok1 := args[0].(*tender.Int)
			name, ok2 := args[1].(*tender.String)
			menuId, ok3 := args[2].(*tender.Int)
			if !ok1 || !ok2 || !ok3 {
				return nil, tender.ErrInvalidArgument
			}
			glut.ChangeToSubMenu(int(entry.Value), name.Value, int(menuId.Value))
			return tender.NullValue, nil
		},
	},

	"remove_menu_item": &tender.BuiltinFunction{
		Name: "remove_menu_item",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			entry, ok1 := args[0].(*tender.Int)
			name, ok2 := args[1].(*tender.String)
			menuId, ok3 := args[2].(*tender.Int)
			if !ok1 || !ok2 || !ok3 {
				return nil, tender.ErrInvalidArgument
			}
			glut.RemoveMenuItem(int(entry.Value), name.Value, int(menuId.Value))
			return tender.NullValue, nil
		},
	},

	"attach_menu": &tender.BuiltinFunction{
		Name: "attach_menu",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			button, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.AttachMenu(int(button.Value))
			return tender.NullValue, nil
		},
	},

	"detach_menu": &tender.BuiltinFunction{
		Name: "detach_menu",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			button, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.DetachMenu(int(button.Value))
			return tender.NullValue, nil
		},
	},

	// ============================================================
	// GET / SET COLOR FUNCTIONS
	// ============================================================
	"set_color": &tender.BuiltinFunction{
		Name: "set_color",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			cell, ok1 := args[0].(*tender.Int)
			red, ok2 := args[1].(*tender.Float)
			green, ok3 := args[2].(*tender.Float)
			blue, ok4 := args[3].(*tender.Float)
			if !ok1 || !ok2 || !ok3 || !ok4 {
				return nil, tender.ErrInvalidArgument
			}
			glut.SetColor(int(cell.Value), float32(red.Value), float32(green.Value), float32(blue.Value))
			return tender.NullValue, nil
		},
	},

	"get_color": &tender.BuiltinFunction{
		Name: "get_color",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			cell, ok1 := args[0].(*tender.Int)
			component, ok2 := args[1].(*tender.Int)
			if !ok1 || !ok2 {
				return nil, tender.ErrInvalidArgument
			}
			val := glut.GetColor(int(cell.Value), int(component.Value))
			return &tender.Float{Value: float64(val)}, nil
		},
	},

	"copy_colormap": &tender.BuiltinFunction{
		Name: "copy_colormap",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			windowId, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.CopyColormap(int(windowId.Value))
			return tender.NullValue, nil
		},
	},

	// ============================================================
	// GET FUNCTION
	// ============================================================
	"get": &tender.BuiltinFunction{
		Name: "get",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			state, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			val := glut.Get(int(state.Value))
			return &tender.Int{Value: int64(val)}, nil
		},
	},

	// ============================================================
	// MODIFIER FUNCTIONS
	// ============================================================
	"get_modifiers": &tender.BuiltinFunction{
		Name: "get_modifiers",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			mod := glut.GetModifiers()
			return &tender.Int{Value: int64(mod)}, nil
		},
	},

	"ignore_key_repeat": &tender.BuiltinFunction{
		Name: "ignore_key_repeat",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			ignore, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.IgnoreKeyRepeat(int(ignore.Value))
			return tender.NullValue, nil
		},
	},

	"set_key_repeat": &tender.BuiltinFunction{
		Name: "set_key_repeat",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			repeatMode, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.SetKeyRepeat(int(repeatMode.Value))
			return tender.NullValue, nil
		},
	},

	// ============================================================
	// DEVICE FUNCTIONS
	// ============================================================
	"device_get": &tender.BuiltinFunction{
		Name: "device_get",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			type_, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			val := glut.DeviceGet(int(type_.Value))
			return &tender.Int{Value: int64(val)}, nil
		},
	},

	// ============================================================
	// OVERLAY FUNCTIONS
	// ============================================================
	"establish_overlay": &tender.BuiltinFunction{
		Name: "establish_overlay",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.EstablishOverlay()
			return tender.NullValue, nil
		},
	},

	"remove_overlay": &tender.BuiltinFunction{
		Name: "remove_overlay",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.RemoveOverlay()
			return tender.NullValue, nil
		},
	},

	"use_layer": &tender.BuiltinFunction{
		Name: "use_layer",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			layer, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.UseLayer(int(layer.Value))
			return tender.NullValue, nil
		},
	},

	"show_overlay": &tender.BuiltinFunction{
		Name: "show_overlay",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.ShowOverlay()
			return tender.NullValue, nil
		},
	},

	"hide_overlay": &tender.BuiltinFunction{
		Name: "hide_overlay",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.HideOverlay()
			return tender.NullValue, nil
		},
	},

	"post_overlay_redisplay": &tender.BuiltinFunction{
		Name: "post_overlay_redisplay",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.PostOverlayRedisplay()
			return tender.NullValue, nil
		},
	},

	"post_window_overlay_redisplay": &tender.BuiltinFunction{
		Name: "post_window_overlay_redisplay",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			windowId, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.PostWindowOverlayRedisplay(int(windowId.Value))
			return tender.NullValue, nil
		},
	},

	"layer_get": &tender.BuiltinFunction{
		Name: "layer_get",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			type_, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			val := glut.LayerGet(int(type_.Value))
			return &tender.Int{Value: int64(val)}, nil
		},
	},

	// ============================================================
	// POST REDISPLAY FUNCTIONS
	// ============================================================
	"post_redisplay": &tender.BuiltinFunction{
		Name: "post_redisplay",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.PostRedisplay()
			return tender.NullValue, nil
		},
	},

	"post_window_redisplay": &tender.BuiltinFunction{
		Name: "post_window_redisplay",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			windowId, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.PostWindowRedisplay(int(windowId.Value))
			return tender.NullValue, nil
		},
	},

	// ============================================================
	// SWAP BUFFERS
	// ============================================================
	"swap_buffers": &tender.BuiltinFunction{
		Name: "swap_buffers",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.SwapBuffers()
			return tender.NullValue, nil
		},
	},

	// ============================================================
	// MAIN LOOP
	// ============================================================
	"main_loop": &tender.BuiltinFunction{
		Name: "main_loop",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.MainLoop()
			return tender.NullValue, nil
		},
	},

	// ============================================================
	// GAME MODE FUNCTIONS
	// ============================================================
	"enter_game_mode": &tender.BuiltinFunction{
		Name: "enter_game_mode",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			val := glut.EnterGameMode()
			return &tender.Int{Value: int64(val)}, nil
		},
	},

	"leave_game_mode": &tender.BuiltinFunction{
		Name: "leave_game_mode",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.LeaveGameMode()
			return tender.NullValue, nil
		},
	},

	"game_mode_string": &tender.BuiltinFunction{
		Name: "game_mode_string",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			s, ok := args[0].(*tender.String)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.GameModeString(s.Value)
			return tender.NullValue, nil
		},
	},

	"game_mode_get": &tender.BuiltinFunction{
		Name: "game_mode_get",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			mode, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			val := glut.GameModeGet(int(mode.Value))
			return &tender.Int{Value: int64(val)}, nil
		},
	},

	"force_joystick_func": &tender.BuiltinFunction{
		Name: "force_joystick_func",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.ForceJoystickFunc()
			return tender.NullValue, nil
		},
	},

	// ============================================================
	// VIDEO RESIZE FUNCTIONS
	// ============================================================
	"setup_video_resizing": &tender.BuiltinFunction{
		Name: "setup_video_resizing",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.SetupVideoResizing()
			return tender.NullValue, nil
		},
	},

	"stop_video_resizing": &tender.BuiltinFunction{
		Name: "stop_video_resizing",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.StopVideoResizing()
			return tender.NullValue, nil
		},
	},

	"video_resize": &tender.BuiltinFunction{
		Name: "video_resize",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			x, ok1 := args[0].(*tender.Int)
			y, ok2 := args[1].(*tender.Int)
			w, ok3 := args[2].(*tender.Int)
			h, ok4 := args[3].(*tender.Int)
			if !ok1 || !ok2 || !ok3 || !ok4 {
				return nil, tender.ErrInvalidArgument
			}
			glut.VideoResize(int(x.Value), int(y.Value), int(w.Value), int(h.Value))
			return tender.NullValue, nil
		},
	},

	"video_pan": &tender.BuiltinFunction{
		Name: "video_pan",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			x, ok1 := args[0].(*tender.Int)
			y, ok2 := args[1].(*tender.Int)
			w, ok3 := args[2].(*tender.Int)
			h, ok4 := args[3].(*tender.Int)
			if !ok1 || !ok2 || !ok3 || !ok4 {
				return nil, tender.ErrInvalidArgument
			}
			glut.VideoPan(int(x.Value), int(y.Value), int(w.Value), int(h.Value))
			return tender.NullValue, nil
		},
	},

	"video_resize_get": &tender.BuiltinFunction{
		Name: "video_resize_get",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			param, ok := args[0].(*tender.Int)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			val := glut.VideoResizeGet(int(param.Value))
			return &tender.Int{Value: int64(val)}, nil
		},
	},

	// ============================================================
	// FONT FUNCTIONS
	// ============================================================
	"bitmap_character": &tender.BuiltinFunction{
		Name: "bitmap_character",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			font, ok1 := args[0].(*tender.Int)
			char, ok2 := args[1].(*tender.Char)
			if !ok1 || !ok2 {
				return nil, tender.ErrInvalidArgument
			}
			glut.BitmapCharacter(int(font.Value), char.Value)
			return tender.NullValue, nil
		},
	},

	"bitmap_length": &tender.BuiltinFunction{
		Name: "bitmap_length",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			font, ok1 := args[0].(*tender.Int)
			s, ok2 := args[1].(*tender.String)
			if !ok1 || !ok2 {
				return nil, tender.ErrInvalidArgument
			}
			val := glut.BitmapLength(int(font.Value), s.Value)
			return &tender.Int{Value: int64(val)}, nil
		},
	},

	"bitmap_width": &tender.BuiltinFunction{
		Name: "bitmap_width",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			font, ok1 := args[0].(*tender.Int)
			char, ok2 := args[1].(*tender.Char)
			if !ok1 || !ok2 {
				return nil, tender.ErrInvalidArgument
			}
			val := glut.BitmapWidth(int(font.Value), char.Value)
			return &tender.Int{Value: int64(val)}, nil
		},
	},

	"stroke_character": &tender.BuiltinFunction{
		Name: "stroke_character",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			font, ok1 := args[0].(*tender.Int)
			char, ok2 := args[1].(*tender.Char)
			if !ok1 || !ok2 {
				return nil, tender.ErrInvalidArgument
			}
			glut.StrokeCharacter(int(font.Value), char.Value)
			return tender.NullValue, nil
		},
	},

	"stroke_length": &tender.BuiltinFunction{
		Name: "stroke_length",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			font, ok1 := args[0].(*tender.Int)
			s, ok2 := args[1].(*tender.String)
			if !ok1 || !ok2 {
				return nil, tender.ErrInvalidArgument
			}
			val := glut.StrokeLength(int(font.Value), s.Value)
			return &tender.Int{Value: int64(val)}, nil
		},
	},

	"stroke_width": &tender.BuiltinFunction{
		Name: "stroke_width",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 2 {
				return nil, tender.ErrInvalidArgCount
			}
			font, ok1 := args[0].(*tender.Int)
			char, ok2 := args[1].(*tender.Char)
			if !ok1 || !ok2 {
				return nil, tender.ErrInvalidArgument
			}
			val := glut.StrokeWidth(int(font.Value), char.Value)
			return &tender.Int{Value: int64(val)}, nil
		},
	},

	// ============================================================
	// SOLID AND WIRE SHAPES
	// ============================================================
	"solid_sphere": &tender.BuiltinFunction{
		Name: "solid_sphere",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			radius, okR := args[0].(*tender.Float)
			slices, okS := args[1].(*tender.Int)
			stacks, okT := args[2].(*tender.Int)
			if !okR || !okS || !okT {
				return nil, tender.ErrInvalidArgument
			}
			glut.SolidSphere(radius.Value, int(slices.Value), int(stacks.Value))
			return tender.NullValue, nil
		},
	},

	"wire_sphere": &tender.BuiltinFunction{
		Name: "wire_sphere",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 3 {
				return nil, tender.ErrInvalidArgCount
			}
			radius, okR := args[0].(*tender.Float)
			slices, okS := args[1].(*tender.Int)
			stacks, okT := args[2].(*tender.Int)
			if !okR || !okS || !okT {
				return nil, tender.ErrInvalidArgument
			}
			glut.WireSphere(radius.Value, int(slices.Value), int(stacks.Value))
			return tender.NullValue, nil
		},
	},

	"solid_cube": &tender.BuiltinFunction{
		Name: "solid_cube",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			size, ok := args[0].(*tender.Float)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.SolidCube(size.Value)
			return tender.NullValue, nil
		},
	},

	"wire_cube": &tender.BuiltinFunction{
		Name: "wire_cube",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			size, ok := args[0].(*tender.Float)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.WireCube(size.Value)
			return tender.NullValue, nil
		},
	},

	"solid_teapot": &tender.BuiltinFunction{
		Name: "solid_teapot",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			size, ok := args[0].(*tender.Float)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.SolidTeapot(size.Value)
			return tender.NullValue, nil
		},
	},

	"wire_teapot": &tender.BuiltinFunction{
		Name: "wire_teapot",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			size, ok := args[0].(*tender.Float)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			glut.WireTeapot(size.Value)
			return tender.NullValue, nil
		},
	},

	"solid_cone": &tender.BuiltinFunction{
		Name: "solid_cone",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			base, okB := args[0].(*tender.Float)
			height, okH := args[1].(*tender.Float)
			slices, okS := args[2].(*tender.Int)
			stacks, okT := args[3].(*tender.Int)
			if !okB || !okH || !okS || !okT {
				return nil, tender.ErrInvalidArgument
			}
			glut.SolidCone(base.Value, height.Value, int(slices.Value), int(stacks.Value))
			return tender.NullValue, nil
		},
	},

	"wire_cone": &tender.BuiltinFunction{
		Name: "wire_cone",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			base, okB := args[0].(*tender.Float)
			height, okH := args[1].(*tender.Float)
			slices, okS := args[2].(*tender.Int)
			stacks, okT := args[3].(*tender.Int)
			if !okB || !okH || !okS || !okT {
				return nil, tender.ErrInvalidArgument
			}
			glut.WireCone(base.Value, height.Value, int(slices.Value), int(stacks.Value))
			return tender.NullValue, nil
		},
	},

	"solid_torus": &tender.BuiltinFunction{
		Name: "solid_torus",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			inner, okI := args[0].(*tender.Float)
			outer, okO := args[1].(*tender.Float)
			nsides, okN := args[2].(*tender.Int)
			rings, okR := args[3].(*tender.Int)
			if !okI || !okO || !okN || !okR {
				return nil, tender.ErrInvalidArgument
			}
			glut.SolidTorus(inner.Value, outer.Value, int(nsides.Value), int(rings.Value))
			return tender.NullValue, nil
		},
	},

	"wire_torus": &tender.BuiltinFunction{
		Name: "wire_torus",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 4 {
				return nil, tender.ErrInvalidArgCount
			}
			inner, okI := args[0].(*tender.Float)
			outer, okO := args[1].(*tender.Float)
			nsides, okN := args[2].(*tender.Int)
			rings, okR := args[3].(*tender.Int)
			if !okI || !okO || !okN || !okR {
				return nil, tender.ErrInvalidArgument
			}
			glut.WireTorus(inner.Value, outer.Value, int(nsides.Value), int(rings.Value))
			return tender.NullValue, nil
		},
	},

	"solid_dodecahedron": &tender.BuiltinFunction{
		Name: "solid_dodecahedron",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.SolidDodecahedron()
			return tender.NullValue, nil
		},
	},

	"wire_dodecahedron": &tender.BuiltinFunction{
		Name: "wire_dodecahedron",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.WireDodecahedron()
			return tender.NullValue, nil
		},
	},

	"solid_icosahedron": &tender.BuiltinFunction{
		Name: "solid_icosahedron",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.SolidIcosahedron()
			return tender.NullValue, nil
		},
	},

	"wire_icosahedron": &tender.BuiltinFunction{
		Name: "wire_icosahedron",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.WireIcosahedron()
			return tender.NullValue, nil
		},
	},

	"solid_octahedron": &tender.BuiltinFunction{
		Name: "solid_octahedron",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.SolidOctahedron()
			return tender.NullValue, nil
		},
	},

	"wire_octahedron": &tender.BuiltinFunction{
		Name: "wire_octahedron",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.WireOctahedron()
			return tender.NullValue, nil
		},
	},

	"solid_tetrahedron": &tender.BuiltinFunction{
		Name: "solid_tetrahedron",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.SolidTetrahedron()
			return tender.NullValue, nil
		},
	},

	"wire_tetrahedron": &tender.BuiltinFunction{
		Name: "wire_tetrahedron",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.WireTetrahedron()
			return tender.NullValue, nil
		},
	},

	"extension_supported": &tender.BuiltinFunction{
		Name: "extension_supported",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 1 {
				return nil, tender.ErrInvalidArgCount
			}
			extension, ok := args[0].(*tender.String)
			if !ok {
				return nil, tender.ErrInvalidArgument
			}
			supported := glut.ExtensionSupported(extension.Value)
			if supported {
				return tender.TrueValue, nil
			}
			return tender.FalseValue, nil
		},
	},

	"report_errors": &tender.BuiltinFunction{
		Name: "report_errors",
		Value: func(args ...tender.Object) (tender.Object, error) {
			if len(args) != 0 {
				return nil, tender.ErrInvalidArgCount
			}
			glut.ReportErrors()
			return tender.NullValue, nil
		},
	},
}