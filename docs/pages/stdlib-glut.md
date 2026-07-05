# Stdlib `glut`

The `glut` module provides OpenGL Utility Toolkit (GLUT) bindings for creating windowed OpenGL applications. It handles window creation, event processing, input handling, and rendering callbacks.

---

## Initialization

### `init()`

Initializes GLUT.

- **Returns**: `null`
- **Example**: `glut.init()`

### `init_display_mode(mode)`

Sets the initial display mode for GLUT windows.

- **Parameters**: `mode` - Display mode flags (bitwise OR of constants)
- **Returns**: `null`
- **Example**: `glut.init_display_mode(glut.RGBA | glut.DOUBLE | glut.DEPTH)`

### `init_window_size(width, height)`

Sets the initial window size.

- **Parameters**: `width`, `height` - Window dimensions (int)
- **Returns**: `null`
- **Example**: `glut.init_window_size(800, 600)`

### `init_window_position(x, y)`

Sets the initial window position.

- **Parameters**: `x`, `y` - Window position (int)
- **Returns**: `null`
- **Example**: `glut.init_window_position(100, 100)`

---

## Window Management

### `create_window(title)`

Creates a new GLUT window.

- **Parameters**: `title` - Window title (string)
- **Returns**: Window ID (int)
- **Example**:
  ```go
  win := glut.create_window("My OpenGL Window")
  ```

### `create_sub_window(parent_window, x, y, width, height)`

Creates a sub-window.

- **Parameters**: `parent_window` - Parent window ID, `x`, `y` - Position, `width`, `height` - Dimensions (int)
- **Returns**: Sub-window ID (int)

### `destroy_window(window_id)`

Destroys a GLUT window.

- **Parameters**: `window_id` - Window ID (int)
- **Returns**: `null`

### `get_window()`

Returns the ID of the current window.

- **Returns**: Window ID (int)
- **Example**: `current := glut.get_window()`

### `set_window(window_id)`

Sets the current window.

- **Parameters**: `window_id` - Window ID (int)
- **Returns**: `null`

### `set_window_title(title)`

Sets the title of the current window.

- **Parameters**: `title` - Window title (string)
- **Returns**: `null`
- **Example**: `glut.set_window_title("New Title")`

### `set_icon_title(title)`

Sets the icon title of the current window.

- **Parameters**: `title` - Icon title (string)
- **Returns**: `null`

### `position_window(x, y)`

Repositions the current window.

- **Parameters**: `x`, `y` - New window position (int)
- **Returns**: `null`

### `reshape_window(width, height)`

Reshapes the current window.

- **Parameters**: `width`, `height` - New dimensions (int)
- **Returns**: `null`

### `full_screen()`

Switches the current window to full-screen mode.

- **Returns**: `null`
- **Example**: `glut.full_screen()`

### `hide_window()`

Hides the current window.

- **Returns**: `null`

### `show_window()`

Shows the current window.

- **Returns**: `null`

### `iconify_window()`

Iconifies (minimizes) the current window.

- **Returns**: `null`

### `push_window()`

Pushes the current window onto the window stack.

- **Returns**: `null`

### `pop_window()`

Pops the current window from the window stack.

- **Returns**: `null`

### `warp_pointer(x, y)`

Warps the mouse pointer to the specified position.

- **Parameters**: `x`, `y` - Position (int)
- **Returns**: `null`

### `set_cursor(cursor)`

Sets the cursor for the current window.

- **Parameters**: `cursor` - Cursor constant
- **Returns**: `null`
- **Example**: `glut.set_cursor(glut.CURSOR_CROSSHAIR)`

---

## Callback Functions

### `display_func(callback)`

Registers the display callback function.

- **Parameters**: `callback` - Function with no parameters
- **Returns**: `null`
- **Example**:
  ```go
  glut.display_func(fn() {
      gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
      // Rendering code
      glut.swap_buffers()
  })
  ```

### `reshape_func(callback)`

Registers the reshape callback function.

- **Parameters**: `callback` - Function with `(width, height)` parameters (int)
- **Returns**: `null`
- **Example**:
  ```go
  glut.reshape_func(fn(w, h) {
      gl.viewport(0, 0, w, h)
  })
  ```

### `keyboard_func(callback)`

Registers the keyboard callback function.

- **Parameters**: `callback` - Function with `(key, x, y)` parameters (string, int, int)
- **Returns**: `null`
- **Example**:
  ```go
  glut.keyboard_func(fn(key, x, y) {
      if key == "q" {
          os.exit(0)
      }
  })
  ```

### `keyboard_up_func(callback)`

Registers the keyboard release callback function.

- **Parameters**: `callback` - Function with `(key, x, y)` parameters (string, int, int)
- **Returns**: `null`

### `special_func(callback)`

Registers the special key callback function (arrow keys, F-keys, etc.).

- **Parameters**: `callback` - Function with `(key, x, y)` parameters (int, int, int)
- **Returns**: `null`
- **Example**:
  ```go
  glut.special_func(fn(key, x, y) {
      if key == glut.KEY_LEFT {
          // Handle left arrow key
      }
  })
  ```

### `special_up_func(callback)`

Registers the special key release callback function.

- **Parameters**: `callback` - Function with `(key, x, y)` parameters (int, int, int)
- **Returns**: `null`

### `mouse_func(callback)`

Registers the mouse button callback function.

- **Parameters**: `callback` - Function with `(button, state, x, y)` parameters (int, int, int, int)
- **Returns**: `null`
- **Example**:
  ```go
  glut.mouse_func(fn(button, state, x, y) {
      if button == glut.LEFT_BUTTON && state == glut.DOWN {
          // Handle left click
      }
  })
  ```

### `motion_func(callback)`

Registers the mouse motion callback function (when a button is pressed).

- **Parameters**: `callback` - Function with `(x, y)` parameters (int, int)
- **Returns**: `null`

### `passive_motion_func(callback)`

Registers the passive mouse motion callback function (no buttons pressed).

- **Parameters**: `callback` - Function with `(x, y)` parameters (int, int)
- **Returns**: `null`

### `entry_func(callback)`

Registers the window entry callback function.

- **Parameters**: `callback` - Function with `(state)` parameter (int)
- **Returns**: `null`
- **Example**:
  ```go
  glut.entry_func(fn(state) {
      if state == glut.ENTERED {
          // Mouse entered the window
      }
  })
  ```

### `visibility_func(callback)`

Registers the window visibility callback function.

- **Parameters**: `callback` - Function with `(state)` parameter (int)
- **Returns**: `null`

### `idle_func(callback)`

Registers the idle callback function (called when no events are pending).

- **Parameters**: `callback` - Function with no parameters
- **Returns**: `null`
- **Example**:
  ```go
  glut.idle_func(fn() {
      // Background processing
      glut.post_redisplay()
  })
  ```

### `timer_func(msecs, callback, timer_id)`

Registers a timer callback.

- **Parameters**: `msecs` - Time in milliseconds (int), `callback` - Function with `(id)` parameter (int), `timer_id` - Timer identifier (int)
- **Returns**: `null`
- **Example**:
  ```go
  glut.timer_func(100, fn(id) {
      // Timer triggered
  }, 0)
  ```

### `overlay_display_func(callback)`

Registers the overlay display callback function.

- **Parameters**: `callback` - Function with no parameters
- **Returns**: `null`

---

## Menus

### `create_menu(callback)`

Creates a new menu.

- **Parameters**: `callback` - Function with `(value)` parameter (int)
- **Returns**: Menu ID (int)
- **Example**:
  ```go
  menu := glut.create_menu(fn(value) {
      println("Menu item selected:", value)
  })
  ```

### `destroy_menu(menu_id)`

Destroys a menu.

- **Parameters**: `menu_id` - Menu ID (int)
- **Returns**: `null`

### `get_menu()`

Returns the ID of the current menu.

- **Returns**: Menu ID (int)

### `set_menu(menu_id)`

Sets the current menu.

- **Parameters**: `menu_id` - Menu ID (int)
- **Returns**: `null`

### `add_menu_entry(name, value)`

Adds an entry to the current menu.

- **Parameters**: `name` - Menu entry name (string), `value` - Value for callback (int)
- **Returns**: `null`
- **Example**:
  ```go
  glut.add_menu_entry("Option 1", 1)
  glut.add_menu_entry("Option 2", 2)
  ```

### `add_sub_menu(name, sub_menu_id)`

Adds a sub-menu to the current menu.

- **Parameters**: `name` - Menu entry name (string), `sub_menu_id` - Sub-menu ID (int)
- **Returns**: `null`

### `change_to_menu_entry(entry, name, value)`

Changes an existing menu entry.

- **Parameters**: `entry` - Entry index (int), `name` - New name (string), `value` - New value (int)
- **Returns**: `null`

### `change_to_sub_menu(entry, name, sub_menu_id)`

Changes an existing menu entry to a sub-menu.

- **Parameters**: `entry` - Entry index (int), `name` - New name (string), `sub_menu_id` - Sub-menu ID (int)
- **Returns**: `null`

### `remove_menu_item(entry, name, menu_id)`

Removes a menu item.

- **Parameters**: `entry` - Entry index (int), `name` - Name (string), `menu_id` - Menu ID (int)
- **Returns**: `null`

### `attach_menu(button)`

Attaches the current menu to a mouse button.

- **Parameters**: `button` - Mouse button (`glut.LEFT_BUTTON`, `glut.RIGHT_BUTTON`, or `glut.MIDDLE_BUTTON`)
- **Returns**: `null`
- **Example**: `glut.attach_menu(glut.RIGHT_BUTTON)`

### `detach_menu(button)`

Detaches the current menu from a mouse button.

- **Parameters**: `button` - Mouse button (int)
- **Returns**: `null`

### `menu_state_func(callback)`

Registers the menu state callback function.

- **Parameters**: `callback` - Function with `(status)` parameter (int)
- **Returns**: `null`

### `menu_status_func(callback)`

Registers the menu status callback function.

- **Parameters**: `callback` - Function with `(status, x, y)` parameters (int, int, int)
- **Returns**: `null`

---

## Colors

### `set_color(cell, red, green, blue)`

Sets a color in the colormap.

- **Parameters**: `cell` - Color cell index (int), `red`, `green`, `blue` - Color components (float, 0.0-1.0)
- **Returns**: `null`

### `get_color(cell, component)`

Gets a color component from the colormap.

- **Parameters**: `cell` - Color cell index (int), `component` - `glut.RED`, `glut.GREEN`, or `glut.BLUE` (int)
- **Returns**: Color component value (float)

### `copy_colormap(window_id)`

Copies the colormap from another window.

- **Parameters**: `window_id` - Source window ID (int)
- **Returns**: `null`

---

## State Queries

### `get(state)`

Gets GLUT state values.

- **Parameters**: `state` - State constant (`glut.WINDOW_WIDTH`, `glut.WINDOW_HEIGHT`, etc.)
- **Returns**: State value (int)
- **Example**:
  ```go
  width := glut.get(glut.WINDOW_WIDTH)
  height := glut.get(glut.WINDOW_HEIGHT)
  ```

### `device_get(type)`

Gets device state values.

- **Parameters**: `type` - Device constant (`glut.HAS_KEYBOARD`, `glut.NUM_MOUSE_BUTTONS`, etc.)
- **Returns**: Device state value (int)

### `layer_get(type)`

Gets overlay layer state values.

- **Parameters**: `type` - Layer constant (`glut.OVERLAY_POSSIBLE`, etc.)
- **Returns**: Layer state value (int)

### `video_resize_get(param)`

Gets video resize state values.

- **Parameters**: `param` - Video resize parameter constant
- **Returns**: Value (int)

### `game_mode_get(mode)`

Gets game mode state values.

- **Parameters**: `mode` - Game mode constant
- **Returns**: Value (int)

---

## Modifiers

### `get_modifiers()`

Returns the current modifier key state.

- **Returns**: Bitmask of active modifiers (`glut.ACTIVE_SHIFT`, `glut.ACTIVE_CTRL`, `glut.ACTIVE_ALT`)
- **Example**:
  ```go
  mods := glut.get_modifiers()
  if mods & glut.ACTIVE_CTRL != 0 {
      // Ctrl key is pressed
  }
  ```

### `ignore_key_repeat(ignore)`

Sets whether to ignore key repeat events.

- **Parameters**: `ignore` - 0 for repeat enabled, 1 for disabled (int)
- **Returns**: `null`

### `set_key_repeat(repeat_mode)`

Sets the key repeat mode.

- **Parameters**: `repeat_mode` - `glut.KEY_REPEAT_OFF`, `glut.KEY_REPEAT_ON`, or `glut.KEY_REPEAT_DEFAULT`
- **Returns**: `null`

---

## Overlay Functions

### `establish_overlay()`

Establishes an overlay for the current window.

- **Returns**: `null`

### `remove_overlay()`

Removes the overlay from the current window.

- **Returns**: `null`

### `use_layer(layer)`

Sets the current layer to use.

- **Parameters**: `layer` - `glut.NORMAL` or `glut.OVERLAY`
- **Returns**: `null`

### `show_overlay()`

Shows the overlay.

- **Returns**: `null`

### `hide_overlay()`

Hides the overlay.

- **Returns**: `null`

### `post_overlay_redisplay()`

Posts a redisplay event for the overlay.

- **Returns**: `null`

### `post_window_overlay_redisplay(window_id)`

Posts a redisplay event for the overlay of a specific window.

- **Parameters**: `window_id` - Window ID (int)
- **Returns**: `null`

---

## Redisplay

### `post_redisplay()`

Posts a redisplay event for the current window.

- **Returns**: `null`

### `post_window_redisplay(window_id)`

Posts a redisplay event for a specific window.

- **Parameters**: `window_id` - Window ID (int)
- **Returns**: `null`

---

## Double Buffering

### `swap_buffers()`

Swaps the front and back buffers for double-buffered windows.

- **Returns**: `null`
- **Example**: `glut.swap_buffers()`

---

## Main Loop

### `main_loop()`

Enters the GLUT main event processing loop.

- **Returns**: `null` (this function never returns)
- **Example**: `glut.main_loop()`

---

## Font Functions

### `bitmap_character(font, char)`

Renders a bitmap character.

- **Parameters**: `font` - Font constant, `char` - Character to render (rune)
- **Returns**: `null`

### `bitmap_length(font, string)`

Returns the length of a string in pixels for a bitmap font.

- **Parameters**: `font` - Font constant, `string` - String to measure
- **Returns**: Length in pixels (int)

### `bitmap_width(font, char)`

Returns the width of a character in a bitmap font.

- **Parameters**: `font` - Font constant, `char` - Character (rune)
- **Returns**: Width in pixels (int)

### `stroke_character(font, char)`

Renders a stroke character.

- **Parameters**: `font` - Font constant, `char` - Character to render (rune)
- **Returns**: `null`

### `stroke_length(font, string)`

Returns the length of a string for a stroke font.

- **Parameters**: `font` - Font constant, `string` - String to measure
- **Returns**: Length (int)

### `stroke_width(font, char)`

Returns the width of a character in a stroke font.

- **Parameters**: `font` - Font constant, `char` - Character (rune)
- **Returns**: Width (int)

---

## Solid and Wire Shapes

### `solid_sphere(radius, slices, stacks)`
### `wire_sphere(radius, slices, stacks)`

Draws a sphere.

- **Parameters**: `radius` - Sphere radius (float), `slices` - Number of slices (int), `stacks` - Number of stacks (int)
- **Returns**: `null`
- **Example**: `glut.solid_sphere(1.0, 32, 32)`

### `solid_cube(size)`
### `wire_cube(size)`

Draws a cube.

- **Parameters**: `size` - Cube side length (float)
- **Returns**: `null`
- **Example**: `glut.solid_cube(2.0)`

### `solid_teapot(size)`
### `wire_teapot(size)`

Draws a teapot (classic GLUT object).

- **Parameters**: `size` - Teapot size (float)
- **Returns**: `null`

### `solid_cone(base, height, slices, stacks)`
### `wire_cone(base, height, slices, stacks)`

Draws a cone.

- **Parameters**: `base` - Base radius (float), `height` - Height (float), `slices` - Number of slices (int), `stacks` - Number of stacks (int)
- **Returns**: `null`

### `solid_torus(inner, outer, nsides, rings)`
### `wire_torus(inner, outer, nsides, rings)`

Draws a torus.

- **Parameters**: `inner` - Inner radius (float), `outer` - Outer radius (float), `nsides` - Number of sides (int), `rings` - Number of rings (int)
- **Returns**: `null`

### `solid_dodecahedron()`
### `wire_dodecahedron()`

Draws a dodecahedron.

- **Returns**: `null`

### `solid_icosahedron()`
### `wire_icosahedron()`

Draws an icosahedron.

- **Returns**: `null`

### `solid_octahedron()`
### `wire_octahedron()`

Draws an octahedron.

- **Returns**: `null`

### `solid_tetrahedron()`
### `wire_tetrahedron()`

Draws a tetrahedron.

- **Returns**: `null`

---

## Game Mode

### `enter_game_mode()`

Enters game mode (full-screen with custom resolution).

- **Returns**: Window ID (int)
- **Example**:
  ```go
  glut.game_mode_string("800x600:32@60")
  win := glut.enter_game_mode()
  ```

### `leave_game_mode()`

Leaves game mode.

- **Returns**: `null`

### `game_mode_string(string)`

Sets the game mode configuration string.

- **Parameters**: `string` - Mode string (e.g., "800x600:32@60")
- **Returns**: `null`

---

## Video Resizing

### `setup_video_resizing()`

Sets up video resizing.

- **Returns**: `null`

### `stop_video_resizing()`

Stops video resizing.

- **Returns**: `null`

### `video_resize(x, y, width, height)`

Resizes the video display.

- **Parameters**: `x`, `y` - Position, `width`, `height` - Dimensions (int)
- **Returns**: `null`

### `video_pan(x, y, width, height)`

Pans the video display.

- **Parameters**: `x`, `y` - Position, `width`, `height` - Dimensions (int)
- **Returns**: `null`

---

## Utility Functions

### `extension_supported(extension)`

Checks if an OpenGL extension is supported.

- **Parameters**: `extension` - Extension name (string)
- **Returns**: Boolean indicating support
- **Example**:
  ```go
  if glut.extension_supported("GL_ARB_multitexture") {
      // Use multitexture
  }
  ```

### `force_joystick_func()`

Forces the joystick callback to be called.

- **Returns**: `null`

### `report_errors()`

Reports GLUT errors to stderr.

- **Returns**: `null`

---

## Constants Reference

### Display Modes
- `glut.RGB`
- `glut.RGBA`
- `glut.INDEX`
- `glut.SINGLE`
- `glut.DOUBLE`
- `glut.ACCUM`
- `glut.ALPHA`
- `glut.DEPTH`
- `glut.STENCIL`
- `glut.MULTISAMPLE`
- `glut.STEREO`
- `glut.LUMINANCE`

### Mouse Buttons
- `glut.LEFT_BUTTON`
- `glut.MIDDLE_BUTTON`
- `glut.RIGHT_BUTTON`
- `glut.DOWN`
- `glut.UP`

### Special Keys
- `glut.KEY_F1` through `glut.KEY_F12`
- `glut.KEY_LEFT`
- `glut.KEY_UP`
- `glut.KEY_RIGHT`
- `glut.KEY_DOWN`
- `glut.KEY_PAGE_UP`
- `glut.KEY_PAGE_DOWN`
- `glut.KEY_HOME`
- `glut.KEY_END`
- `glut.KEY_INSERT`

### Modifiers
- `glut.ACTIVE_SHIFT`
- `glut.ACTIVE_CTRL`
- `glut.ACTIVE_ALT`

### Entry/Exit States
- `glut.LEFT`
- `glut.ENTERED`

### Font Constants
- `glut.STROKE_ROMAN`
- `glut.STROKE_MONO_ROMAN`
- `glut.BITMAP_9_BY_15`
- `glut.BITMAP_8_BY_13`
- `glut.BITMAP_TIMES_ROMAN_10`
- `glut.BITMAP_TIMES_ROMAN_24`
- `glut.BITMAP_HELVETICA_10`
- `glut.BITMAP_HELVETICA_12`
- `glut.BITMAP_HELVETICA_18`

### Cursors
- `glut.CURSOR_RIGHT_ARROW`
- `glut.CURSOR_LEFT_ARROW`
- `glut.CURSOR_INFO`
- `glut.CURSOR_DESTROY`
- `glut.CURSOR_HELP`
- `glut.CURSOR_CYCLE`
- `glut.CURSOR_SPRAY`
- `glut.CURSOR_WAIT`
- `glut.CURSOR_TEXT`
- `glut.CURSOR_CROSSHAIR`
- `glut.CURSOR_UP_DOWN`
- `glut.CURSOR_LEFT_RIGHT`
- `glut.CURSOR_NONE`
- `glut.CURSOR_FULL_CROSSHAIR`

---

## Complete Example

```go
import "gl"
import "glut"
import "os"

// Initialize GLUT
glut.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE | glut.DEPTH)
glut.init_window_size(800, 600)
glut.init_window_position(100, 100)

// Create window
win := glut.create_window("GLUT OpenGL Demo")
println("Window created with ID:", win)

// Setup callbacks
glut.display_func(fn() {
    gl.clear_color(0.2, 0.3, 0.4, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
    gl.matrix_mode(gl.MODELVIEW)
    gl.load_identity()
    
    // Draw a rotating triangle
    gl.begin(gl.TRIANGLES)
    gl.color3f(1.0, 0.0, 0.0)
    gl.vertex3f(0.0, 0.5, 0.0)
    gl.color3f(0.0, 1.0, 0.0)
    gl.vertex3f(-0.5, -0.5, 0.0)
    gl.color3f(0.0, 0.0, 1.0)
    gl.vertex3f(0.5, -0.5, 0.0)
    gl.end()
    
    glut.swap_buffers()
})

glut.reshape_func(fn(w, h) {
    gl.viewport(0, 0, w, h)
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    gl.ortho(-1.0, 1.0, -1.0, 1.0, -1.0, 1.0)
})

glut.keyboard_func(fn(key, x, y) {
    if key == "q" || key == "Q" {
        os.exit(0)
    }
})

glut.idle_func(fn() {
    // Rotate the scene
    gl.push_matrix()
    gl.rotatef(0.5, 0.0, 1.0, 0.0)
    glut.post_redisplay()
})

// Create menu
menu := glut.create_menu(fn(value) {
    println("Menu selection:", value)
})
glut.add_menu_entry("Option 1", 1)
glut.add_menu_entry("Option 2", 2)
glut.attach_menu(glut.RIGHT_BUTTON)

// Start main loop
glut.main_loop()
```