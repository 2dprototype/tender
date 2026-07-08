# Stdlib `glfw`

The `glfw` module provides bindings for the GLFW library, enabling window creation, OpenGL context management, and input handling for 3D applications. It is designed to work alongside the `gl` module for complete OpenGL applications.

## Initialization

### `init()`

Initializes the GLFW library.

- **Returns**: `null`
- **Example**: `glfw.init()`

### `terminate()`

Terminates the GLFW library and destroys all remaining windows.

- **Returns**: `null`
- **Example**: `glfw.terminate()`

### `poll_events()`

Processes all pending events (keyboard, mouse, window, joystick).

- **Returns**: `null`
- **Example**: `glfw.poll_events()`

### `wait_events()`

Waits for events to arrive and processes them. Blocks the current thread.

- **Returns**: `null`
- **Example**: `glfw.wait_events()`

### `post_empty_event()`

Posts an empty event to wake up `wait_events()`.

- **Returns**: `null`

### `get_time()`

Returns the elapsed time since GLFW was initialized, in seconds.

- **Returns**: Float (seconds)
- **Example**: `time := glfw.get_time()`

### `set_time(time)`

Sets the GLFW timer to the specified value.

- **Parameters**: `time` - Time in seconds (float)
- **Returns**: `null`

### `get_timer_frequency()`

Returns the frequency of the GLFW high-resolution timer.

- **Returns**: Integer (frequency in Hz)

### `get_timer_value()`

Returns the current value of the GLFW high-resolution timer.

- **Returns**: Integer (timer value)

---

## Version Information

### `get_version()`

Returns the GLFW version.

- **Returns**: Array `[major, minor, revision]`
- **Example**: `major, minor, rev := glfw.get_version()`

### `get_version_string()`

Returns the GLFW version string.

- **Returns**: String
- **Example**: `glfw.get_version_string()`

### Constants

- `glfw.VERSION_MAJOR`
- `glfw.VERSION_MINOR`
- `glfw.VERSION_REVISION`

---

## Window Management

### `create_window(width, height, title)`

Creates a new window.

- **Parameters**:
  - `width` - Window width (int)
  - `height` - Window height (int)
  - `title` - Window title (string)
- **Returns**: Window ID (int)
- **Example**: `win := glfw.create_window(800, 600, "My Window")`

### `destroy_window(window)`

Destroys the specified window.

- **Parameters**: `window` - Window ID (int)
- **Returns**: `null`

### `window_should_close(window)`

Checks if the window should close.

- **Parameters**: `window` - Window ID (int)
- **Returns**: Boolean

### `set_window_should_close(window, value)`

Sets the close flag for the window.

- **Parameters**:
  - `window` - Window ID (int)
  - `value` - Boolean
- **Returns**: `null`

### `get_window_size(window)`

Gets the window size.

- **Parameters**: `window` - Window ID (int)
- **Returns**: Array `[width, height]`
- **Example**: `w, h := glfw.get_window_size(win)`

### `set_window_size(window, width, height)`

Sets the window size.

- **Parameters**:
  - `window` - Window ID (int)
  - `width` - New width (int)
  - `height` - New height (int)
- **Returns**: `null`

### `get_framebuffer_size(window)`

Gets the framebuffer size (useful for OpenGL viewport).

- **Parameters**: `window` - Window ID (int)
- **Returns**: Array `[width, height]`
- **Example**: `w, h := glfw.get_framebuffer_size(win)`

### `get_window_pos(window)`

Gets the window position.

- **Parameters**: `window` - Window ID (int)
- **Returns**: Array `[x, y]`

### `set_window_pos(window, x, y)`

Sets the window position.

- **Parameters**:
  - `window` - Window ID (int)
  - `x`, `y` - Position (int)
- **Returns**: `null`

### `set_window_title(window, title)`

Sets the window title.

- **Parameters**:
  - `window` - Window ID (int)
  - `title` - New title (string)
- **Returns**: `null`

### `iconify_window(window)`

Iconifies (minimizes) the window.

- **Parameters**: `window` - Window ID (int)
- **Returns**: `null`

### `restore_window(window)`

Restores the window from iconified or maximized state.

- **Parameters**: `window` - Window ID (int)
- **Returns**: `null`

### `maximize_window(window)`

Maximizes the window.

- **Parameters**: `window` - Window ID (int)
- **Returns**: `null`

### `show_window(window)`

Shows the window.

- **Parameters**: `window` - Window ID (int)
- **Returns**: `null`

### `hide_window(window)`

Hides the window.

- **Parameters**: `window` - Window ID (int)
- **Returns**: `null`

### `focus_window(window)`

Brings the window to focus.

- **Parameters**: `window` - Window ID (int)
- **Returns**: `null`

### `request_window_attention(window)`

Requests user attention for the window (flashes taskbar entry).

- **Parameters**: `window` - Window ID (int)
- **Returns**: `null`

### `get_window_monitor(window)`

Gets the monitor the window is on.

- **Parameters**: `window` - Window ID (int)
- **Returns**: Monitor ID (int) or `null` if window is windowed

### `set_window_monitor(window, monitor, xpos, ypos, width, height, refreshRate)`

Sets the monitor for the window. If monitor is 0, the window becomes windowed.

- **Parameters**:
  - `window` - Window ID (int)
  - `monitor` - Monitor ID (int) or 0 for windowed
  - `xpos`, `ypos` - Position (int)
  - `width`, `height` - Size (int)
  - `refreshRate` - Refresh rate (int)
- **Returns**: `null`

### `get_window_attrib(window, attrib)`

Gets a window attribute.

- **Parameters**:
  - `window` - Window ID (int)
  - `attrib` - Attribute hint constant
- **Returns**: Integer value

### `set_window_attrib(window, attrib, value)`

Sets a window attribute.

- **Parameters**:
  - `window` - Window ID (int)
  - `attrib` - Attribute hint constant
  - `value` - Value (int)
- **Returns**: `null`

### `get_window_content_scale(window)`

Gets the content scale of the window.

- **Parameters**: `window` - Window ID (int)
- **Returns**: Array `[xscale, yscale]`

### `get_window_opacity(window)`

Gets the window opacity.

- **Parameters**: `window` - Window ID (int)
- **Returns**: Float (0.0 - 1.0)

### `set_window_opacity(window, opacity)`

Sets the window opacity.

- **Parameters**:
  - `window` - Window ID (int)
  - `opacity` - Opacity value (float, 0.0 - 1.0)
- **Returns**: `null`

### `set_window_size_limits(window, minw, minh, maxw, maxh)`

Sets the window size limits. Use 0 for no limit.

- **Parameters**:
  - `window` - Window ID (int)
  - `minw`, `minh` - Minimum size (int)
  - `maxw`, `maxh` - Maximum size (int)
- **Returns**: `null`

### `set_window_aspect_ratio(window, numer, denom)`

Sets the window aspect ratio constraint.

- **Parameters**:
  - `window` - Window ID (int)
  - `numer`, `denom` - Aspect ratio numerator and denominator (int)
- **Returns**: `null`

### `set_window_icon(window, icon)`

Sets the window icon. Currently only supports `null` to clear the icon.

- **Parameters**:
  - `window` - Window ID (int)
  - `icon` - `null` to clear, or image data (not yet implemented)
- **Returns**: `null`

---

## Context Management

### `make_context_current(window)`

Makes the window's OpenGL context current.

- **Parameters**: `window` - Window ID (int)
- **Returns**: `null`
- **Example**: `glfw.make_context_current(win)`

### `get_current_context()`

Gets the window with the current OpenGL context.

- **Returns**: Window ID (int) or `null`

### `swap_buffers(window)`

Swaps the front and back buffers.

- **Parameters**: `window` - Window ID (int)
- **Returns**: `null`

### `swap_interval(interval)`

Sets the swap interval (V-Sync).

- **Parameters**: `interval` - 0 for V-Sync off, 1 for V-Sync on
- **Returns**: `null`

---

## Window Hints

### `window_hint(hint, value)`

Sets a window hint before creating the window.

- **Parameters**:
  - `hint` - Hint constant
  - `value` - Value (int)
- **Returns**: `null`
- **Example**:
  ```go
  glfw.window_hint(glfw.RESIZABLE, 0)
  glfw.window_hint(glfw.CONTEXT_VERSION_MAJOR, 3)
  glfw.window_hint(glfw.CONTEXT_VERSION_MINOR, 3)
  glfw.window_hint(glfw.OPENGL_PROFILE, glfw.OPENGL_CORE_PROFILE)
  ```

### `default_window_hints()`

Resets all window hints to their default values.

- **Returns**: `null`

### Window Hint Constants

#### Window Behavior
- `glfw.RESIZABLE`
- `glfw.VISIBLE`
- `glfw.DECORATED`
- `glfw.FOCUSED`
- `glfw.MAXIMIZED`
- `glfw.FLOATING`
- `glfw.FOCUS_ON_SHOW`
- `glfw.TRANSPARENT_FRAMEBUFFER`
- `glfw.CENTER_CURSOR`
- `glfw.SCALE_TO_MONITOR`

#### OpenGL Context
- `glfw.CONTEXT_VERSION_MAJOR`
- `glfw.CONTEXT_VERSION_MINOR`
- `glfw.OPENGL_PROFILE`
- `glfw.OPENGL_CORE_PROFILE`
- `glfw.OPENGL_COMPAT_PROFILE`
- `glfw.DOUBLEBUFFER`

#### Framebuffer
- `glfw.SAMPLES` (MSAA)
- `glfw.REFRESH_RATE`
- `glfw.STEREO`
- `glfw.SRGB_CAPABLE`
- `glfw.RED_BITS`
- `glfw.GREEN_BITS`
- `glfw.BLUE_BITS`
- `glfw.ALPHA_BITS`
- `glfw.DEPTH_BITS`
- `glfw.STENCIL_BITS`
- `glfw.ACCUM_RED_BITS`
- `glfw.ACCUM_GREEN_BITS`
- `glfw.ACCUM_BLUE_BITS`
- `glfw.ACCUM_ALPHA_BITS`
- `glfw.AUX_BUFFERS`

---

## Input Handling

### Keyboard

#### `get_key(window, key)`

Gets the state of a key.

- **Parameters**:
  - `window` - Window ID (int)
  - `key` - Key constant
- **Returns**: Action constant (`glfw.PRESS`, `glfw.RELEASE`, or `glfw.REPEAT`)
- **Example**: `state := glfw.get_key(win, glfw.KEY_ESCAPE)`

#### `get_key_name(key, scancode)`

Gets the name of a key.

- **Parameters**:
  - `key` - Key constant
  - `scancode` - Scan code (int)
- **Returns**: Key name (string) or `null`

#### `get_key_scancode(key)`

Gets the platform-specific scan code for a key.

- **Parameters**: `key` - Key constant
- **Returns**: Scan code (int)

#### Key Constants

**Alphanumeric:**
- `glfw.KEY_A` through `glfw.KEY_Z`
- `glfw.KEY_0` through `glfw.KEY_9`

**Function Keys:**
- `glfw.KEY_F1` through `glfw.KEY_F12`

**Navigation:**
- `glfw.KEY_UP`, `glfw.KEY_DOWN`, `glfw.KEY_LEFT`, `glfw.KEY_RIGHT`
- `glfw.KEY_HOME`, `glfw.KEY_END`, `glfw.KEY_PAGE_UP`, `glfw.KEY_PAGE_DOWN`

**Modifiers:**
- `glfw.KEY_LEFT_SHIFT`, `glfw.KEY_RIGHT_SHIFT`
- `glfw.KEY_LEFT_CONTROL`, `glfw.KEY_RIGHT_CONTROL`
- `glfw.KEY_LEFT_ALT`, `glfw.KEY_RIGHT_ALT`

**Other:**
- `glfw.KEY_ESCAPE`, `glfw.KEY_ENTER`, `glfw.KEY_SPACE`, `glfw.KEY_TAB`
- `glfw.KEY_BACKSPACE`, `glfw.KEY_DELETE`, `glfw.KEY_INSERT`
- `glfw.KEY_PRINT_SCREEN`, `glfw.KEY_PAUSE`
- `glfw.KEY_CAPS_LOCK`, `glfw.KEY_SCROLL_LOCK`, `glfw.KEY_NUM_LOCK`
- `glfw.KEY_GRAVE_ACCENT`, `glfw.KEY_WORLD_1`, `glfw.KEY_WORLD_2`
- `glfw.KEY_MENU`

#### Key Actions
- `glfw.PRESS`
- `glfw.RELEASE`
- `glfw.REPEAT`

### Mouse

#### `get_mouse_button(window, button)`

Gets the state of a mouse button.

- **Parameters**:
  - `window` - Window ID (int)
  - `button` - Mouse button constant
- **Returns**: Action constant (`glfw.PRESS` or `glfw.RELEASE`)

#### `get_cursor_pos(window)`

Gets the cursor position.

- **Parameters**: `window` - Window ID (int)
- **Returns**: Array `[x, y]` (float)

#### `set_cursor_pos(window, x, y)`

Sets the cursor position.

- **Parameters**:
  - `window` - Window ID (int)
  - `x`, `y` - Position (float)
- **Returns**: `null`

#### `get_input_mode(window, mode)`

Gets the input mode.

- **Parameters**:
  - `window` - Window ID (int)
  - `mode` - Input mode constant
- **Returns**: Value (int)

#### `set_input_mode(window, mode, value)`

Sets the input mode.

- **Parameters**:
  - `window` - Window ID (int)
  - `mode` - Input mode constant
  - `value` - Value (int)
- **Returns**: `null`
- **Example**: `glfw.set_input_mode(win, glfw.CURSOR, glfw.CURSOR_DISABLED)`

#### Mouse Button Constants
- `glfw.MOUSE_BUTTON_LEFT`
- `glfw.MOUSE_BUTTON_RIGHT`
- `glfw.MOUSE_BUTTON_MIDDLE`
- `glfw.MOUSE_BUTTON_4` through `glfw.MOUSE_BUTTON_8`

#### Cursor Mode Constants
- `glfw.CURSOR_NORMAL` - Normal cursor operation
- `glfw.CURSOR_HIDDEN` - Cursor is hidden
- `glfw.CURSOR_DISABLED` - Cursor is hidden and locked

#### Input Mode Constants
- `glfw.CURSOR` - Cursor mode
- `glfw.LOCK_KEY_MODS` - Lock key modifier detection
- `glfw.RAW_MOUSE_MOTION` - Raw mouse motion

---

## Cursors

### `create_standard_cursor(shape)`

Creates a standard system cursor.

- **Parameters**: `shape` - Cursor shape constant
- **Returns**: Cursor ID (int) or `null`
- **Example**: `cursor := glfw.create_standard_cursor(glfw.CURSOR_HAND)`

### `destroy_cursor(cursor)`

Destroys a cursor.

- **Parameters**: `cursor` - Cursor ID (int)
- **Returns**: `null`

### `set_window_cursor(window, cursor)`

Sets the cursor for a window.

- **Parameters**:
  - `window` - Window ID (int)
  - `cursor` - Cursor ID (int) or 0 for default
- **Returns**: `null`

### Cursor Shape Constants
- `glfw.CURSOR_ARROW`
- `glfw.CURSOR_IBEAM`
- `glfw.CURSOR_CROSSHAIR`
- `glfw.CURSOR_HAND`
- `glfw.CURSOR_HRESIZE`
- `glfw.CURSOR_VRESIZE`

---

## Modifier Keys

### Modifier Constants
- `glfw.MOD_SHIFT`
- `glfw.MOD_CONTROL`
- `glfw.MOD_ALT`
- `glfw.MOD_SUPER`
- `glfw.MOD_CAPS_LOCK`
- `glfw.MOD_NUM_LOCK`

---

## Monitors

### `get_primary_monitor()`

Gets the primary monitor.

- **Returns**: Monitor ID (int) or `null`

### `get_monitor_pos(monitor)`

Gets the monitor position.

- **Parameters**: `monitor` - Monitor ID (int)
- **Returns**: Array `[x, y]`

### `get_monitor_name(monitor)`

Gets the monitor name.

- **Parameters**: `monitor` - Monitor ID (int)
- **Returns**: Monitor name (string)

### `get_video_mode(monitor)`

Gets the current video mode of the monitor.

- **Parameters**: `monitor` - Monitor ID (int)
- **Returns**: Map with `width`, `height`, `refresh_rate`

### `get_monitor_video_modes(monitor)`

Gets all available video modes for the monitor.

- **Parameters**: `monitor` - Monitor ID (int)
- **Returns**: Array of video mode maps

### `get_monitor_physical_size(monitor)`

Gets the physical size of the monitor in millimeters.

- **Parameters**: `monitor` - Monitor ID (int)
- **Returns**: Array `[width, height]`

### `get_monitor_workarea(monitor)`

Gets the work area of the monitor (excluding taskbars).

- **Parameters**: `monitor` - Monitor ID (int)
- **Returns**: Array `[x, y, width, height]`

### `get_monitor_gamma_ramp(monitor)`

Gets the current gamma ramp of the monitor.

- **Parameters**: `monitor` - Monitor ID (int)
- **Returns**: Map with `red`, `green`, `blue` arrays

### `set_monitor_gamma_ramp(monitor, ramp)`

Sets the gamma ramp of the monitor.

- **Parameters**:
  - `monitor` - Monitor ID (int)
  - `ramp` - Map with `red`, `green`, `blue` arrays (uint16 values)
- **Returns**: `null`

### `set_monitor_gamma(monitor, gamma)`

Sets the gamma of the monitor using a generated ramp.

- **Parameters**:
  - `monitor` - Monitor ID (int)
  - `gamma` - Gamma value (float)
- **Returns**: `null`

---

## Joysticks & Gamepads

### Joystick Constants
- `glfw.JOYSTICK_1` through `glfw.JOYSTICK_16`

### Joystick States
- `glfw.CONNECTED`
- `glfw.DISCONNECTED`

### Joystick Functions

#### `joystick_present(jid)`

Checks if a joystick is present.

- **Parameters**: `jid` - Joystick ID (int)
- **Returns**: Boolean

#### `joystick_get_name(jid)`

Gets the joystick name.

- **Parameters**: `jid` - Joystick ID (int)
- **Returns**: Name (string) or `null`

#### `joystick_is_gamepad(jid)`

Checks if the joystick is a gamepad.

- **Parameters**: `jid` - Joystick ID (int)
- **Returns**: Boolean

#### `joystick_get_gamepad_name(jid)`

Gets the gamepad name.

- **Parameters**: `jid` - Joystick ID (int)
- **Returns**: Name (string) or `null`

#### `joystick_get_gamepad_state(jid)`

Gets the gamepad state.

- **Parameters**: `jid` - Joystick ID (int)
- **Returns**: Map with `axes` and `buttons` arrays

#### `joystick_get_axes(jid)`

Gets the joystick axes.

- **Parameters**: `jid` - Joystick ID (int)
- **Returns**: Array of float values

#### `joystick_get_buttons(jid)`

Gets the joystick button states.

- **Parameters**: `jid` - Joystick ID (int)
- **Returns**: Array of button states (`glfw.PRESS` or `glfw.RELEASE`)

#### `joystick_get_hats(jid)`

Gets the joystick hat states.

- **Parameters**: `jid` - Joystick ID (int)
- **Returns**: Array of hat states

#### `joystick_get_guid(jid)`

Gets the joystick GUID.

- **Parameters**: `jid` - Joystick ID (int)
- **Returns**: GUID (string) or `null`

### `update_gamepad_mappings(mapping)`

Updates the gamepad mappings from a string.

- **Parameters**: `mapping` - Gamepad mapping string
- **Returns**: Boolean (success)

---

## Clipboard

### `get_clipboard_string()`

Gets the clipboard contents.

- **Returns**: String

### `set_clipboard_string(str)`

Sets the clipboard contents.

- **Parameters**: `str` - String to set
- **Returns**: `null`

---

## Vulkan

### `vulkan_supported()`

Checks if Vulkan is supported.

- **Returns**: Boolean

### `raw_mouse_motion_supported()`

Checks if raw mouse motion is supported.

- **Returns**: Boolean

---

## Callbacks

All callback functions accept a callable as their last argument. The callable will be invoked when the event occurs. Callbacks are optional - pass `null` to remove an existing callback.

### `set_window_size_callback(window, callback)`

Called when the window size changes.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func(width, height)` or `null`
- **Returns**: `null`

### `set_framebuffer_size_callback(window, callback)`

Called when the framebuffer size changes.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func(width, height)` or `null`
- **Returns**: `null`

### `set_window_close_callback(window, callback)`

Called when the user attempts to close the window.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func()` or `null`
- **Returns**: `null`

### `set_key_callback(window, callback)`

Called when a key is pressed, released, or repeated.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func(key, scancode, action, mods)` or `null`
- **Returns**: `null`

### `set_char_callback(window, callback)`

Called when a character is typed.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func(char)` or `null`
- **Returns**: `null`

### `set_mouse_button_callback(window, callback)`

Called when a mouse button is pressed or released.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func(button, action, mods)` or `null`
- **Returns**: `null`

### `set_cursor_pos_callback(window, callback)`

Called when the cursor moves.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func(x, y)` or `null`
- **Returns**: `null`

### `set_cursor_enter_callback(window, callback)`

Called when the cursor enters or leaves the window.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func(entered)` or `null`
- **Returns**: `null`

### `set_scroll_callback(window, callback)`

Called when the mouse scroll wheel is used.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func(xoffset, yoffset)` or `null`
- **Returns**: `null`

### `set_joystick_callback(callback)`

Called when a joystick connects or disconnects.

- **Parameters**:
  - `callback` - Function `func(jid, event)` or `null`
- **Returns**: `null`

### `set_drop_callback(window, callback)`

Called when files are dropped onto the window.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func(paths)` or `null` (paths is an array of strings)
- **Returns**: `null`

### `set_focus_callback(window, callback)`

Called when the window gains or loses focus.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func(focused)` or `null`
- **Returns**: `null`

### `set_iconify_callback(window, callback)`

Called when the window is iconified or restored.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func(iconified)` or `null`
- **Returns**: `null`

### `set_maximize_callback(window, callback)`

Called when the window is maximized or restored.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func(maximized)` or `null`
- **Returns**: `null`

### `set_content_scale_callback(window, callback)`

Called when the window content scale changes.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func(xscale, yscale)` or `null`
- **Returns**: `null`

### `set_pos_callback(window, callback)`

Called when the window position changes.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func(xpos, ypos)` or `null`
- **Returns**: `null`

### `set_refresh_callback(window, callback)`

Called when the window needs refreshing.

- **Parameters**:
  - `window` - Window ID (int)
  - `callback` - Function `func()` or `null`
- **Returns**: `null`

---

## Complete Example

```go
import "gl"
import "glfw"

glfw.init()
gl.init()

// Set OpenGL version hints (optional, for modern OpenGL)
// For compatibility with fixed-function pipeline, use 2.1 or lower
glfw.window_hint(glfw.CONTEXT_VERSION_MAJOR, 2)
glfw.window_hint(glfw.CONTEXT_VERSION_MINOR, 1)

window := glfw.create_window(400, 400, "Tender Triangle Test (GLFW)")

glfw.make_context_current(window)

wh := glfw.get_window_size(window)
width := wh[0]
height := wh[1]
gl.viewport(0, 0, width, height)

gl.matrix_mode(gl.PROJECTION)
gl.load_identity()
gl.ortho(-1.0, 1.0, -1.0, 1.0, -1.0, 1.0)
gl.matrix_mode(gl.MODELVIEW)
gl.load_identity()

glfw.set_key_callback(window, fn(key, scancode, action, mods){
    println("Key: " + key + ", scancode: " + scancode + ", action: " + action + ", mods: " + mods)
    if key == glfw.KEY_ESCAPE && action == glfw.PRESS {
        glfw.set_window_should_close(window, 1)
    }
})

glfw.set_window_size_callback(window, fn(w, h){
    gl.viewport(0, 0, w, h)
    println("Window resized to: " + w + "x" + h)
})

glfw.set_mouse_button_callback(window, fn(button, action, mods){
    if action == glfw.PRESS {
        println("Mouse button " + button + " pressed")
    }
})

for !glfw.window_should_close(window) {
    gl.clear_color(0.0, 0.0, 0.0, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT)
    gl.load_identity()
    gl.begin(gl.TRIANGLES)
    gl.color3f(1.0, 0.0, 0.0) 
    gl.vertex3f(0.0, 0.75, 0.0)
    gl.color3f(0.0, 1.0, 0.0)
    gl.vertex3f(-0.75, -0.75, 0.0)
    gl.color3f(0.0, 0.0, 1.0) 
    gl.vertex3f(0.75, -0.75, 0.0)
    gl.end()
    gl.flush()
    glfw.swap_buffers(window)
    glfw.poll_events()
}

glfw.destroy_window(window)
glfw.terminate()
```