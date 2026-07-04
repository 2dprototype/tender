/*
 * This is free and unencumbered software released into the public domain.
 * 
 * Anyone is free to copy, modify, publish, use, compile, sell, or
 * distribute this software, either in source code form or as a compiled
 * binary, for any purpose, commercial or non-commercial, and by any
 * means.
 * 
 * In jurisdictions that recognize copyright laws, the author or authors
 * of this software dedicate any and all copyright interest in the
 * software to the public domain. We make this dedication for the benefit
 * of the public at large and to the detriment of our heirs and
 * successors. We intend this dedication to be an overt act of
 * relinquishment in perpetuity of all present and future rights to this
 * software under copyright law.
 * 
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 * OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 * ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
 * OTHER DEALINGS IN THE SOFTWARE.
 * 
 * For more information, please refer to <http://unlicense.org>
 */


package glut

// #cgo CFLAGS: -ID:/lib/freeglut/include
// #cgo LDFLAGS: -LD:/lib/freeglut/bin/x64 -l:freeglut.dll
// #include <stdlib.h>
// #include <GL/glut.h>
// #include "gofunctions.h"
import "C"
import (
	"unsafe"
	"os"
	"strings"
)


// display modes for InitDisplayMode
const (
	RGB = C.GLUT_RGB
	RGBA = C.GLUT_RGBA
	INDEX = C.GLUT_INDEX
	SINGLE = C.GLUT_SINGLE
	DOUBLE = C.GLUT_DOUBLE
	ACCUM = C.GLUT_ACCUM
	ALPHA = C.GLUT_ALPHA
	DEPTH = C.GLUT_DEPTH
	STENCIL = C.GLUT_STENCIL
	MULTISAMPLE = C.GLUT_MULTISAMPLE
	STEREO = C.GLUT_STEREO
	LUMINANCE = C.GLUT_LUMINANCE
)

// mouse states
const (
	LEFT_BUTTON = C.GLUT_LEFT_BUTTON
	MIDDLE_BUTTON = C.GLUT_MIDDLE_BUTTON
	RIGHT_BUTTON = C.GLUT_RIGHT_BUTTON
	DOWN = C.GLUT_DOWN
	UP = C.GLUT_UP
)

// special key codes
const (
	KEY_F1 = C.GLUT_KEY_F1
	KEY_F2 = C.GLUT_KEY_F2
	KEY_F3 = C.GLUT_KEY_F3
	KEY_F4 = C.GLUT_KEY_F4
	KEY_F5 = C.GLUT_KEY_F5
	KEY_F6 = C.GLUT_KEY_F6
	KEY_F7 = C.GLUT_KEY_F7
	KEY_F8 = C.GLUT_KEY_F8
	KEY_F9 = C.GLUT_KEY_F9
	KEY_F10 = C.GLUT_KEY_F10
	KEY_F11 = C.GLUT_KEY_F11
	KEY_F12 = C.GLUT_KEY_F12
	KEY_LEFT = C.GLUT_KEY_LEFT
	KEY_UP = C.GLUT_KEY_UP
	KEY_RIGHT = C.GLUT_KEY_RIGHT
	KEY_DOWN = C.GLUT_KEY_DOWN
	KEY_PAGE_UP = C.GLUT_KEY_PAGE_UP
	KEY_PAGE_DOWN = C.GLUT_KEY_PAGE_DOWN
	KEY_HOME = C.GLUT_KEY_HOME
	KEY_END = C.GLUT_KEY_END
	KEY_INSERT = C.GLUT_KEY_INSERT
)

// entry/exit callback state
const (
	LEFT = C.GLUT_LEFT
	ENTERED = C.GLUT_ENTERED
)

// window and menu status
const (
	MENU_NOT_IN_USE = C.GLUT_MENU_NOT_IN_USE
	MENU_IN_USE = C.GLUT_MENU_IN_USE
	NOT_VISIBLE = C.GLUT_NOT_VISIBLE
	VISIBLE = C.GLUT_VISIBLE
	HIDDEN = C.GLUT_HIDDEN
	FULLY_RETAINED = C.GLUT_FULLY_RETAINED
	PARTIALLY_RETAINED = C.GLUT_PARTIALLY_RETAINED
	FULLY_COVERED = C.GLUT_FULLY_COVERED
)

// RGB color component specification for GetColor
const (
	RED = C.GLUT_RED
	GREEN = C.GLUT_GREEN
	BLUE = C.GLUT_BLUE
)

// the UseLayer parameters
const (
	NORMAL = C.GLUT_NORMAL
	OVERLAY = C.GLUT_OVERLAY
)

// fonts
const (
	STROKE_ROMAN = 0
	STROKE_MONO_ROMAN = 1
	BITMAP_9_BY_15 = 2
	BITMAP_8_BY_13 = 3
	BITMAP_TIMES_ROMAN_10 = 4
	BITMAP_TIMES_ROMAN_24 = 5
	BITMAP_HELVETICA_10 = 6
	BITMAP_HELVETICA_12 = 7
	BITMAP_HELVETICA_18 = 8
)

// the Get parameters
const (
	WINDOW_X = C.GLUT_WINDOW_X
	WINDOW_Y = C.GLUT_WINDOW_Y
	WINDOW_WIDTH = C.GLUT_WINDOW_WIDTH
	WINDOW_HEIGHT = C.GLUT_WINDOW_HEIGHT
	WINDOW_BUFFER_SIZE = C.GLUT_WINDOW_BUFFER_SIZE
	WINDOW_STENCIL_SIZE = C.GLUT_WINDOW_STENCIL_SIZE
	WINDOW_DEPTH_SIZE = C.GLUT_WINDOW_DEPTH_SIZE
	WINDOW_RED_SIZE = C.GLUT_WINDOW_RED_SIZE
	WINDOW_GREEN_SIZE = C.GLUT_WINDOW_GREEN_SIZE
	WINDOW_BLUE_SIZE = C.GLUT_WINDOW_BLUE_SIZE
	WINDOW_ALPHA_SIZE = C.GLUT_WINDOW_ALPHA_SIZE
	WINDOW_ACCUM_RED_SIZE = C.GLUT_WINDOW_ACCUM_RED_SIZE
	WINDOW_ACCUM_GREEN_SIZE = C.GLUT_WINDOW_ACCUM_GREEN_SIZE
	WINDOW_ACCUM_BLUE_SIZE = C.GLUT_WINDOW_ACCUM_BLUE_SIZE
	WINDOW_ACCUM_ALPHA_SIZE = C.GLUT_WINDOW_ACCUM_ALPHA_SIZE
	WINDOW_DOUBLEBUFFER = C.GLUT_WINDOW_DOUBLEBUFFER
	WINDOW_RGBA = C.GLUT_WINDOW_RGBA
	WINDOW_PARENT = C.GLUT_WINDOW_PARENT
	WINDOW_NUM_CHILDREN = C.GLUT_WINDOW_NUM_CHILDREN
	WINDOW_COLORMAP_SIZE = C.GLUT_WINDOW_COLORMAP_SIZE
	WINDOW_NUM_SAMPLES = C.GLUT_WINDOW_NUM_SAMPLES
	WINDOW_STEREO = C.GLUT_WINDOW_STEREO
	WINDOW_CURSOR = C.GLUT_WINDOW_CURSOR
	SCREEN_WIDTH = C.GLUT_SCREEN_WIDTH
	SCREEN_HEIGHT = C.GLUT_SCREEN_HEIGHT
	SCREEN_WIDTH_MM = C.GLUT_SCREEN_WIDTH_MM
	SCREEN_HEIGHT_MM = C.GLUT_SCREEN_HEIGHT_MM
	MENU_NUM_ITEMS = C.GLUT_MENU_NUM_ITEMS
	DISPLAY_MODE_POSSIBLE = C.GLUT_DISPLAY_MODE_POSSIBLE
	INIT_WINDOW_X = C.GLUT_INIT_WINDOW_X
	INIT_WINDOW_Y = C.GLUT_INIT_WINDOW_Y
	INIT_WINDOW_WIDTH = C.GLUT_INIT_WINDOW_WIDTH
	INIT_WINDOW_HEIGHT = C.GLUT_INIT_WINDOW_HEIGHT
	INIT_DISPLAY_MODE = C.GLUT_INIT_DISPLAY_MODE
	ELAPSED_TIME = C.GLUT_ELAPSED_TIME
	// glut api version >= 4 or xlib implementation >= 13
	WINDOW_FORMAT_ID = C.GLUT_WINDOW_FORMAT_ID
)

// the DeviceGet parameters
const (
	HAS_KEYBOARD = C.GLUT_HAS_KEYBOARD
	HAS_MOUSE = C.GLUT_HAS_MOUSE
	HAS_SPACEBALL = C.GLUT_HAS_SPACEBALL
	HAS_DIAL_AND_BUTTON_BOX = C.GLUT_HAS_DIAL_AND_BUTTON_BOX
	HAS_TABLET = C.GLUT_HAS_TABLET
	NUM_MOUSE_BUTTONS = C.GLUT_NUM_MOUSE_BUTTONS
	NUM_SPACEBALL_BUTTONS = C.GLUT_NUM_SPACEBALL_BUTTONS
	NUM_BUTTON_BOX_BUTTONS = C.GLUT_NUM_BUTTON_BOX_BUTTONS
	NUM_DIALS = C.GLUT_NUM_DIALS
	NUM_TABLET_BUTTONS = C.GLUT_NUM_TABLET_BUTTONS

	// glut api version >= 4 or xlib implementation >= 13

	DEVICE_IGNORE_KEY_REPEAT = C.GLUT_DEVICE_IGNORE_KEY_REPEAT
	DEVICE_KEY_REPEAT = C.GLUT_DEVICE_KEY_REPEAT
	HAS_JOYSTICK = C.GLUT_HAS_JOYSTICK
	OWNS_JOYSTICK = C.GLUT_OWNS_JOYSTICK
	JOYSTICK_BUTTONS = C.GLUT_JOYSTICK_BUTTONS
	JOYSTICK_AXES = C.GLUT_JOYSTICK_AXES
	JOYSTICK_POLL_RATE = C.GLUT_JOYSTICK_POLL_RATE
)

// the LayerGet parameters
const (
	OVERLAY_POSSIBLE = C.GLUT_OVERLAY_POSSIBLE
	LAYER_IN_USE = C.GLUT_LAYER_IN_USE
	HAS_OVERLAY = C.GLUT_HAS_OVERLAY
	TRANSPARENT_INDEX = C.GLUT_TRANSPARENT_INDEX
	NORMAL_DAMAGED = C.GLUT_NORMAL_DAMAGED
	OVERLAY_DAMAGED = C.GLUT_OVERLAY_DAMAGED
)

// glutVideoResizeGet parameters
// glut api version >= 4 or xlib implementation >= 9
const (
	VIDEO_RESIZE_POSSIBLE = C.GLUT_VIDEO_RESIZE_POSSIBLE
	VIDEO_RESIZE_IN_USE = C.GLUT_VIDEO_RESIZE_IN_USE
	VIDEO_RESIZE_X_DELTA = C.GLUT_VIDEO_RESIZE_X_DELTA
	VIDEO_RESIZE_Y_DELTA = C.GLUT_VIDEO_RESIZE_Y_DELTA
	VIDEO_RESIZE_WIDTH_DELTA = C.GLUT_VIDEO_RESIZE_WIDTH_DELTA
	VIDEO_RESIZE_HEIGHT_DELTA = C.GLUT_VIDEO_RESIZE_HEIGHT_DELTA
	VIDEO_RESIZE_X = C.GLUT_VIDEO_RESIZE_X
	VIDEO_RESIZE_Y = C.GLUT_VIDEO_RESIZE_Y
	VIDEO_RESIZE_WIDTH = C.GLUT_VIDEO_RESIZE_WIDTH
	VIDEO_RESIZE_HEIGHT = C.GLUT_VIDEO_RESIZE_HEIGHT
)

// the GetModifiers parameters
const (
	ACTIVE_SHIFT = C.GLUT_ACTIVE_SHIFT
	ACTIVE_CTRL = C.GLUT_ACTIVE_CTRL
	ACTIVE_ALT = C.GLUT_ACTIVE_ALT
)

// the SetCursor parameters
const (
	CURSOR_RIGHT_ARROW = C.GLUT_CURSOR_RIGHT_ARROW
	CURSOR_LEFT_ARROW = C.GLUT_CURSOR_LEFT_ARROW
	CURSOR_INFO = C.GLUT_CURSOR_INFO
	CURSOR_DESTROY = C.GLUT_CURSOR_DESTROY
	CURSOR_HELP = C.GLUT_CURSOR_HELP
	CURSOR_CYCLE = C.GLUT_CURSOR_CYCLE
	CURSOR_SPRAY = C.GLUT_CURSOR_SPRAY
	CURSOR_WAIT = C.GLUT_CURSOR_WAIT
	CURSOR_TEXT = C.GLUT_CURSOR_TEXT
	CURSOR_CROSSHAIR = C.GLUT_CURSOR_CROSSHAIR
	CURSOR_UP_DOWN = C.GLUT_CURSOR_UP_DOWN
	CURSOR_LEFT_RIGHT = C.GLUT_CURSOR_LEFT_RIGHT
	CURSOR_TOP_SIDE = C.GLUT_CURSOR_TOP_SIDE
	CURSOR_BOTTOM_SIDE = C.GLUT_CURSOR_BOTTOM_SIDE
	CURSOR_LEFT_SIDE = C.GLUT_CURSOR_LEFT_SIDE
	CURSOR_RIGHT_SIDE = C.GLUT_CURSOR_RIGHT_SIDE
	CURSOR_TOP_LEFT_CORNER = C.GLUT_CURSOR_TOP_LEFT_CORNER
	CURSOR_TOP_RIGHT_CORNER = C.GLUT_CURSOR_TOP_RIGHT_CORNER
	CURSOR_BOTTOM_RIGHT_CORNER = C.GLUT_CURSOR_BOTTOM_RIGHT_CORNER
	CURSOR_BOTTOM_LEFT_CORNER = C.GLUT_CURSOR_BOTTOM_LEFT_CORNER
	CURSOR_INHERIT = C.GLUT_CURSOR_INHERIT
	CURSOR_NONE = C.GLUT_CURSOR_NONE
	CURSOR_FULL_CROSSHAIR = C.GLUT_CURSOR_FULL_CROSSHAIR
)

// glut api version >= 4 or xlib implementation >= 13
const (
	KEY_REPEAT_OFF = C.GLUT_KEY_REPEAT_OFF
	KEY_REPEAT_ON = C.GLUT_KEY_REPEAT_ON
	KEY_REPEAT_DEFAULT = C.GLUT_KEY_REPEAT_DEFAULT
	JOYSTICK_BUTTON_A = C.GLUT_JOYSTICK_BUTTON_A
	JOYSTICK_BUTTON_B = C.GLUT_JOYSTICK_BUTTON_B
	JOYSTICK_BUTTON_C = C.GLUT_JOYSTICK_BUTTON_C
	JOYSTICK_BUTTON_D = C.GLUT_JOYSTICK_BUTTON_D
)

// glut api version >= 4 or xlib implementation >= 13
const (
	GAME_MODE_ACTIVE = C.GLUT_GAME_MODE_ACTIVE
	GAME_MODE_POSSIBLE = C.GLUT_GAME_MODE_POSSIBLE
	GAME_MODE_WIDTH = C.GLUT_GAME_MODE_WIDTH
	GAME_MODE_HEIGHT = C.GLUT_GAME_MODE_HEIGHT
	GAME_MODE_PIXEL_DEPTH = C.GLUT_GAME_MODE_PIXEL_DEPTH
	GAME_MODE_REFRESH_RATE = C.GLUT_GAME_MODE_REFRESH_RATE
	GAME_MODE_DISPLAY_CHANGED = C.GLUT_GAME_MODE_DISPLAY_CHANGED
)

type tWindowCallbacks struct {
	display func()
	overlayDisplay func()
	reshape func(width, height int)
	keyboard func(key uint8, x, y int)
	mouse func(button, state, x, y int)
	motion func(x, y int)
	passiveMotion func(x, y int)
	visibility func(state int)
	entry func(state int)
	special func(key, x, y int)
	spaceballMotion func(x, y, z int)
	spaceballRotate func(x, y, z int)
	spaceballButton func(button, state int)
	buttonBox func(button, state int)
	dials func(dial, value int)
	tabletMotion func(x, y int)
	tabletButton func(button, state, x, y int)
	menuStatus func(status, x, y int)
	menuState func(status int)
	windowStatus func(state int)
	keyboardUp func(key uint8, x, y int)
	specialUp func(key, x, y int)
	joystick func(buttonMask uint, x, y, z int)
	idle func()
}


var (
	windowCallbacks = make(map[int]*tWindowCallbacks)
	timers = make(map[int]func(timerId int))
	menus = make(map[int]func(value int))
)

/*
 * Initialization
 */

func Init() {
	argc := C.int(len(os.Args))
	argv := make([]*C.char, argc)
	argv_o := make([]*C.char, argc)

	for i, arg := range os.Args {
		argv[i] = C.CString(arg)
	}
	copy(argv_o, argv)

	defer func() {
		if int(argc) < len(os.Args) {
			for i, j := 0, 0; i < int(argc); i++ {
				argv_c := C.GoString(argv[i])

				for strings.Compare(os.Args[j], argv_c) != 0 {
					j = j + 1
				}
				os.Args[i] = os.Args[j]
				j = j + 1
			}
			os.Args = os.Args[:argc]
		}
		for _, arg := range argv_o {
			C.free(unsafe.Pointer(arg))
		}
	}()

	C.glutInit(&argc, &argv[0])
}

func InitWindowSize(width, height int) {
	C.glutInitWindowSize(C.int(width), C.int(height))
}

func InitWindowPosition(width, height int) {
	C.glutInitWindowPosition(C.int(width), C.int(height))
}

func InitDisplayMode(mode uint) {
	C.glutInitDisplayMode(C.uint(mode))
}

func MainLoop() {
	C.glutMainLoop()
}

/*
 * Window Management
 */

func CreateWindow(title string) (windowId int) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	windowId = int(C.glutCreateWindow(ctitle))
	windowCallbacks[windowId] = new(tWindowCallbacks)

	return windowId
}

func CreateSubWindow(windowId, x, y, width, height int) (subWindowId int) {
	subWindowId = int(C.glutCreateSubWindow(C.int(windowId), C.int(x), C.int(y), C.int(width), C.int(height)))
	windowCallbacks[subWindowId] = new(tWindowCallbacks)
	return subWindowId
}

func SetWindow(windowId int) {
	C.glutSetWindow(C.int(windowId))
}

func GetWindow() (windowId int) {
	windowId = int(C.glutGetWindow())
	return windowId
}

func DestroyWindow(windowId int) {
	C.glutDestroyWindow(C.int(windowId))
	delete(windowCallbacks, windowId)
}

func PostRedisplay() {
	C.glutPostRedisplay()
}

// glut api version >= 4 or xlib implementation >= 11
func PostWindowRedisplay(windowId int) {
	C.glutPostWindowRedisplay(C.int(windowId))
}

func SwapBuffers() {
	C.glutSwapBuffers()
}

func PositionWindow(x, y int) {
	C.glutPositionWindow(C.int(x), C.int(y))
}

func ReshapeWindow(width, height int) {
	C.glutReshapeWindow(C.int(width), C.int(height))
}

func FullScreen() {
	C.glutFullScreen()
}

func PopWindow() {
	C.glutPopWindow()
}

func PushWindow() {
	C.glutPushWindow()
}

func ShowWindow() {
	C.glutShowWindow()
}

func HideWindow() {
	C.glutHideWindow()
}

func IconifyWindow() {
	C.glutIconifyWindow()
}

func SetWindowTitle(title string) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	C.glutSetWindowTitle(ctitle)
}

func SetIconTitle(title string) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	C.glutSetIconTitle(ctitle)
}

func SetCursor(cursor int) {
	C.glutSetCursor(C.int(cursor))
}

// glut api version >= 4 or xlib implementation >= 11
func WarpPointer(x, y int) {
	C.glutWarpPointer(C.int(x), C.int(y))
}

/*
 * Overlay Management
 */

func EstablishOverlay() {
	C.glutEstablishOverlay()
}

// Use glut.NORMAL or glut.OVERLAY as layer.
func UseLayer(layer int) {
	C.glutUseLayer(C.GLenum(layer))
}

func RemoveOverlay() {
	C.glutRemoveOverlay()
}

func PostOverlayRedisplay() {
	C.glutPostOverlayRedisplay()
}

// glut api version >= 4 or xlib implementation >= 11
func PostWindowOverlayRedisplay(windowId int) {
	C.glutPostWindowOverlayRedisplay(C.int(windowId))
}

func ShowOverlay() {
	C.glutShowOverlay()
}

func HideOverlay() {
	C.glutHideOverlay()
}

/*
 * Menu Management
 */

func CreateMenu(menu func(value int)) (menuId int) {
	if menu != nil {
		menuId = int(C.create_menu())
	} else {
		menuId = int(C.create_menuNullCallback())
	}
	menus[menuId] = menu
	return menuId
}

func SetMenu(menuId int) {
	C.glutSetMenu(C.int(menuId))
}

func GetMenu() (menuId int) {
	menuId = int(C.glutGetMenu())
	return menuId
}

func DestroyMenu(menuId int) {
	C.glutDestroyMenu(C.int(menuId))
	delete(menus, menuId)
}

func AddMenuEntry(name string, value int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.glutAddMenuEntry(cname, C.int(value))
}

func AddSubMenu(name string, menuId int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.glutAddSubMenu(cname, C.int(menuId))
}

func ChangeToMenuEntry(entry int, name string, value int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.glutChangeToMenuEntry(C.int(entry), cname, C.int(value))
}

func ChangeToSubMenu(entry int, name string, menuId int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.glutChangeToSubMenu(C.int(entry), cname, C.int(menuId))
}

func RemoveMenuItem(entry int, name string, menuId int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.glutChangeToSubMenu(C.int(entry), cname, C.int(menuId))
}

func AttachMenu(button int) {
	C.glutAttachMenu(C.int(button))
}

func DetachMenu(button int) {
	C.glutDetachMenu(C.int(button))
}

/*
 * Callback Registration
 */

func DisplayFunc(display func()) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].display = display
	if display != nil {
		C.register_display()
	} else {
		C.unregister_display()
	}
}

func OverlayDisplayFunc(overlayDisplay func()) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].overlayDisplay = overlayDisplay
	if overlayDisplay != nil {
		C.register_overlayDisplay()
	} else {
		C.unregister_overlayDisplay()
	}
}

func ReshapeFunc(reshape func(width, height int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].reshape = reshape
	if reshape != nil {
		C.register_reshape()
	} else {
		C.unregister_reshape()
	}
}

func KeyboardFunc(keyboard func(key uint8, x, y int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].keyboard = keyboard
	if keyboard != nil {
		C.register_keyboard()
	} else {
		C.unregister_keyboard()
	}
}

func MouseFunc(mouse func(button, state, x, y int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].mouse = mouse
	if mouse != nil {
		C.register_mouse()
	} else {
		C.unregister_mouse()
	}
}

func MotionFunc(motion func(x, y int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].motion = motion
	if motion != nil {
		C.register_motion()
	} else {
		C.unregister_motion()
	}
}

func PassiveMotionFunc(passiveMotion func(x, y int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].passiveMotion = passiveMotion
	if passiveMotion != nil {
		C.register_passiveMotion()
	} else {
		C.unregister_passiveMotion()
	}
}

func VisibilityFunc(visibility func(state int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].visibility = visibility
	if visibility != nil {
		C.register_visibility()
	} else {
		C.unregister_visibility()
	}
}

func EntryFunc(entry func(state int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].entry = entry
	if entry != nil {
		C.register_entry()
	} else {
		C.unregister_entry()
	}
}

func SpecialFunc(special func(key, x, y int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].special = special
	if special != nil {
		C.register_special()
	} else {
		C.unregister_special()
	}
}

func SpaceballMotionFunc(spaceballMotion func(x, y, z int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].spaceballMotion = spaceballMotion
	if spaceballMotion != nil {
		C.register_spaceballMotion()
	} else {
		C.unregister_spaceballMotion()
	}
}

func SpaceballRotateFunc(spaceballRotate func(x, y, z int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].spaceballRotate = spaceballRotate
	if spaceballRotate != nil {
		C.register_spaceballRotate()
	} else {
		C.unregister_spaceballRotate()
	}
}

func SpaceballButtonFunc(spaceballButton func(button, state int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].spaceballButton = spaceballButton
	if spaceballButton != nil {
		C.register_spaceballButton()
	} else {
		C.unregister_spaceballButton()
	}
}

func ButtonBoxFunc(buttonBox func(button, state int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].buttonBox = buttonBox
	if buttonBox != nil {
		C.register_buttonBox()
	} else {
		C.unregister_buttonBox()
	}
}

func DialsFunc(dials func(dial, value int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].dials = dials
	if dials != nil {
		C.register_dials()
	} else {
		C.unregister_dials()
	}
}

func TabletMotionFunc(tabletMotion func(x, y int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].tabletMotion = tabletMotion
	if tabletMotion != nil {
		C.register_tabletMotion()
	} else {
		C.unregister_tabletMotion()
	}
}

func TabletButtonFunc(tabletButton func(button, state, x, y int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].tabletButton = tabletButton
	if tabletButton != nil {
		C.register_tabletButton()
	} else {
		C.unregister_tabletButton()
	}
}

func MenuStatusFunc(menuStatus func(status, x, y int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].menuStatus = menuStatus
	if menuStatus != nil {
		C.register_menuStatus()
	} else {
		C.unregister_menuStatus()
	}
}

func MenuStateFunc(menuState func(status int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].menuState = menuState
	if menuState != nil {
		C.register_menuState()
	} else {
		C.unregister_menuState()
	}
}

func IdleFunc(idle func()) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].idle = idle
	if idle != nil {
		C.register_idle()
	} else {
		C.unregister_idle()
	}
}

// Do not register a second timer with the same id, before the first run out.
// timer should not be nil.
func TimerFunc(msecs int, timer func(timerId int), timerId int) {
	timers[timerId] = timer
	C.register_timer(C.uint(msecs), C.int(timerId))
}

// glut api version >= 4 or xlib implementation >= 9
func WindowStatusFunc(windowStatus func(state int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].windowStatus = windowStatus
	if windowStatus != nil {
		C.register_windowStatus()
	} else {
		C.unregister_windowStatus()
	}
}

// glut api version >= 4 or xlib implementation >= 13
func KeyboardUpFunc(keyboardUp func(key uint8, x, y int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].keyboardUp = keyboardUp
	if keyboardUp != nil {
		C.register_keyboardUp()
	} else {
		C.unregister_keyboardUp()
	}
}

// glut api version >= 4 or xlib implementation >= 13
func SpecialUpFunc(specialUp func(key, x, y int)) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].specialUp = specialUp
	if specialUp != nil {
		C.register_specialUp()
	} else {
		C.unregister_specialUp()
	}
}

// glut api version >= 4 or xlib implementation >= 13
func JoystickFunc(joystick func(buttonMask uint, x, y, z int), pollInterval int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].joystick = joystick
	if joystick != nil {
		C.register_joystick(C.int(pollInterval))
	} else {
		C.unregister_joystick(C.int(pollInterval))
	}
}

/*
 * Color Index Colormap Management
 */

func SetColor(cell int, red, green, blue float32) {
	C.glutSetColor(C.int(cell), C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
}

func GetColor(cell, component int) (float32) {
	return float32(C.glutGetColor(C.int(cell), C.int(component)))
}

func CopyColormap(windowId int) {
	C.glutCopyColormap(C.int(windowId))
}

/*
 * State Retrieval
 */

func Get(state int) int {
	return int(C.glutGet(C.GLenum(state)))
}

func LayerGet(type_ int) int {
	return int(C.glutLayerGet(C.GLenum(type_)))
}

func DeviceGet(type_ int) int {
	return int(C.glutDeviceGet(C.GLenum(type_)))
}

func GetModifiers() int {
	return int(C.glutGetModifiers())
}

func ExtensionSupported(extension string) (supported bool) {
	cextension := C.CString(extension)
	defer C.free(unsafe.Pointer(cextension))
	supported = C.glutExtensionSupported(cextension) != 0
	return supported
}

/*
 * Font Rendering
 */

func fontPtr(font int) unsafe.Pointer {
	switch font {
		case STROKE_ROMAN: return C.const_GLUT_STROKE_ROMAN()
		case STROKE_MONO_ROMAN: return C.const_GLUT_STROKE_MONO_ROMAN()
		case BITMAP_9_BY_15: return C.const_GLUT_BITMAP_9_BY_15()
		case BITMAP_8_BY_13: return C.const_GLUT_BITMAP_8_BY_13()
		case BITMAP_TIMES_ROMAN_10: return C.const_GLUT_BITMAP_TIMES_ROMAN_10()
		case BITMAP_TIMES_ROMAN_24: return C.const_GLUT_BITMAP_TIMES_ROMAN_24()
		case BITMAP_HELVETICA_10: return C.const_GLUT_BITMAP_HELVETICA_10()
		case BITMAP_HELVETICA_12: return C.const_GLUT_BITMAP_HELVETICA_12()
		case BITMAP_HELVETICA_18: return C.const_GLUT_BITMAP_HELVETICA_18()
	}
	panic("unknown font")
}

func BitmapCharacter(font int, character rune) {
	C.glutBitmapCharacter(fontPtr(font), C.int(character))
}

func BitmapWidth(font int, character rune) int {
	return int(C.glutBitmapWidth(fontPtr(font), C.int(character)))
}

func StrokeCharacter(font int, character rune) {
	C.glutStrokeCharacter(fontPtr(font), C.int(character))
}

func StrokeWidth(font int, character rune) int {
	return int(C.glutStrokeWidth(fontPtr(font), C.int(character)))
}

// glut api version >= 4 or xlib implementation >= 11
func BitmapLength(font int, s string) int {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return int(C.glutBitmapLength(fontPtr(font), (*C.uchar)(unsafe.Pointer(cs))))
}

// glut api version >= 4 or xlib implementation >= 11
func StrokeLength(font int, s string) int {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return int(C.glutStrokeLength(fontPtr(font), (*C.uchar)(unsafe.Pointer(cs))))
}

/*
 * Geometric Object Rendering
 */

func SolidSphere(radius float64, slices, stacks int) {
	C.glutSolidSphereCompatibilityWrapper(C.double(radius), C.GLint(slices), C.GLint(stacks))
}
func WireSphere(radius float64, slices, stacks int) {
	C.glutWireSphereCompatibilityWrapper(C.double(radius), C.GLint(slices), C.GLint(stacks))
}
func SolidCube(size float64) {
	C.glutSolidCubeCompatibilityWrapper(C.double(size))
}
func WireCube(size float64) {
	C.glutWireCubeCompatibilityWrapper(C.double(size))
}
func SolidCone(base, height float64, slices, stacks int) {
	C.glutSolidConeCompatibilityWrapper(C.double(base), C.double(height), C.GLint(slices), C.GLint(stacks))
}
func WireCone(base, height float64, slices, stacks int) {
	C.glutWireConeCompatibilityWrapper(C.double(base), C.double(height), C.GLint(slices), C.GLint(stacks))
}
func SolidTorus(innerRadius, outerRadius float64, nsides, rings int) {
	C.glutSolidTorusCompatibilityWrapper(C.double(innerRadius), C.double(outerRadius), C.GLint(nsides), C.GLint(rings))
}
func WireTorus(innerRadius, outerRadius float64, nsides, rings int) {
	C.glutWireTorusCompatibilityWrapper(C.double(innerRadius), C.double(outerRadius), C.GLint(nsides), C.GLint(rings))
}
func SolidDodecahedron() {
	C.glutSolidDodecahedron()
}
func WireDodecahedron() {
	C.glutWireDodecahedron()
}
func SolidOctahedron() {
	C.glutSolidOctahedron()
}
func WireOctahedron() {
	C.glutWireOctahedron()
}
func SolidTetrahedron() {
	C.glutSolidTetrahedron()
}
func WireTetrahedron() {
	C.glutWireTetrahedron()
}
func SolidIcosahedron() {
	C.glutSolidIcosahedron()
}
func WireIcosahedron() {
	C.glutWireIcosahedron()
}
func SolidTeapot(size float64) {
	C.glutSolidTeapotCompatibilityWrapper(C.double(size))
}
func WireTeapot(size float64) {
	C.glutWireTeapotCompatibilityWrapper(C.double(size))
}

/*
 * video resize sub-API
 */

// glut api version >= 4 or xlib implementation >= 9
func VideoResizeGet(param int) int {
	return int(C.glutVideoResizeGet(C.GLenum(param)))
}

// glut api version >= 4 or xlib implementation >= 9
func SetupVideoResizing() {
	C.glutSetupVideoResizing()
}

// glut api version >= 4 or xlib implementation >= 9
func StopVideoResizing() {
	C.glutStopVideoResizing()
}

// glut api version >= 4 or xlib implementation >= 9
func VideoResize(x, y, width, height int) {
	C.glutVideoResize(C.int(x), C.int(y), C.int(width), C.int(height))
}

// glut api version >= 4 or xlib implementation >= 9
func VideoPan(x, y, width, height int) {
	C.glutVideoPan(C.int(x), C.int(y), C.int(width), C.int(height))
}

/*
 * debugging sub-API
 */

// glut api version >= 4 or xlib implementation >= 9
func ReportErrors() {
	C.glutReportErrors()
}

/*
 * device control sub-API
 */

// glut api version >= 4 or xlib implementation >= 13
func IgnoreKeyRepeat(ignore int) {
	C.glutIgnoreKeyRepeat(C.int(ignore))
}

// glut api version >= 4 or xlib implementation >= 13
func SetKeyRepeat(repeatMode int) {
	C.glutSetKeyRepeat(C.int(repeatMode))
}

// glut api version >= 4 or xlib implementation >= 13
func ForceJoystickFunc() {
	C.glutForceJoystickFunc()
}

/*
 * game mode sub-API
 */

// glut api version >= 4 or xlib implementation >= 13
func GameModeString(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.glutGameModeString(cs)
}

// glut api version >= 4 or xlib implementation >= 13
func EnterGameMode() int {
	return int(C.glutEnterGameMode())
}

// glut api version >= 4 or xlib implementation >= 13
func LeaveGameMode() {
	C.glutLeaveGameMode()
}

// glut api version >= 4 or xlib implementation >= 13
func GameModeGet(mode int) int {
	return int(C.glutGameModeGet(C.GLenum(mode)))
}

/*
 * Go Exported Functions
 */

//export goDisplay
func goDisplay() {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].display()
}

//export goOverlayDisplay
func goOverlayDisplay() {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].overlayDisplay()
}

//export goMenu
func goMenu(value C.int) {
	menuId := int(C.glutGetMenu())
	menus[menuId](int(value))
}

//export goReshape
func goReshape(width, height C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].reshape(int(width), int(height))
}

//export goKeyboard
func goKeyboard(key C.uchar, x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].keyboard(uint8(key), int(x), int(y))
}

//export goMouse
func goMouse(button, state, x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].mouse(int(button), int(state), int(x), int(y))
}

//export goMotion
func goMotion(x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].motion(int(x), int(y))
}

//export goPassiveMotion
func goPassiveMotion(x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].passiveMotion(int(x), int(y))
}

//export goVisibility
func goVisibility(state C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].visibility(int(state))
}

//export goEntry
func goEntry(state C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].entry(int(state))
}

//export goSpecial
func goSpecial(key, x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].special(int(key), int(x), int(y))
}

//export goSpaceballMotion
func goSpaceballMotion(x, y, z C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].spaceballMotion(int(x), int(y), int(z))
}

//export goSpaceballRotate
func goSpaceballRotate(x, y, z C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].spaceballRotate(int(x), int(y), int(z))
}

//export goSpaceballButton
func goSpaceballButton(button, state C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].spaceballButton(int(button), int(state))
}

//export goButtonBox
func goButtonBox(button, state C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].buttonBox(int(button), int(state))
}

//export goDials
func goDials(dial, value C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].dials(int(dial), int(value))
}

//export goTabletMotion
func goTabletMotion(x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].tabletMotion(int(x), int(y))
}

//export goTabletButton
func goTabletButton(button, state, x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].tabletButton(int(button), int(state), int(x), int(y))
}

//export goMenuStatus
func goMenuStatus(status, x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].menuStatus(int(status), int(x), int(y))
}

//export goMenuState
func goMenuState(status C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].menuState(int(status))
}

//export goIdle
func goIdle() {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].idle()
}

//export goTimer
func goTimer(timerId C.int) {
	goTimerId := int(timerId)
	timer := timers[goTimerId]
	delete(timers, goTimerId)
	timer(goTimerId)
}

//export goWindowStatus
func goWindowStatus(state C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].windowStatus(int(state))
}

//export goKeyboardUp
func goKeyboardUp(key C.uchar, x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].keyboardUp(uint8(key), int(x), int(y))
}

//export goSpecialUp
func goSpecialUp(key, x, y C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].specialUp(int(key), int(x), int(y))
}

//export goJoystick
func goJoystick(buttonMask C.uint, x, y, z C.int) {
	windowId := int(C.glutGetWindow())
	windowCallbacks[windowId].joystick(uint(buttonMask), int(x), int(y), int(z))
}







