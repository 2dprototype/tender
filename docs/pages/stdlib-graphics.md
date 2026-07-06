## Stdlib `graphics`

The `graphics` module provides comprehensive 2D graphics functionalities including shape drawing, text rendering, image manipulation, and window management with OpenGL acceleration. It supports both offscreen rendering contexts and interactive windowed applications.

## Context Functions

### `new_context(width, height)`

Creates a new offscreen graphics context with the specified dimensions.

- **Parameters**: 
  - `width` - Width of the context in pixels (integer)
  - `height` - Height of the context in pixels (integer)
- **Returns**: Graphics context object with drawing methods
- **Example**: `ctx := graphics.new_context(800, 600)`

### `new_window(width, height, title)`

Creates a new window with the specified dimensions and title.

- **Parameters**:
  - `width` - Width of the window in pixels (integer)
  - `height` - Height of the window in pixels (integer)
  - `title` - Window title (string)
- **Returns**: Window context object with event handling methods
- **Example**: `win := graphics.new_window(800, 600, "My Game")`

## Color Functions

### `hex(color)`

Sets the current drawing color using hexadecimal color code.

- **Parameters**: `color` - Hex color string (e.g., `"#ff0000"`, `"#f00"`, `"#ff0000ff"`)
- **Returns**: `null`
- **Example**: `ctx.hex("#ff8800")`

### `rgb(r, g, b)`

Sets the current drawing color using RGB values.

- **Parameters**: 
  - `r` - Red component (0.0 to 1.0)
  - `g` - Green component (0.0 to 1.0)
  - `b` - Blue component (0.0 to 1.0)
- **Returns**: `null`
- **Example**: `ctx.rgb(1.0, 0.5, 0.0)`

### `rgba(r, g, b, a)`

Sets the current drawing color with alpha transparency.

- **Parameters**:
  - `r` - Red component (0.0 to 1.0)
  - `g` - Green component (0.0 to 1.0)
  - `b` - Blue component (0.0 to 1.0)
  - `a` - Alpha/opacity (0.0 to 1.0)
- **Returns**: `null`
- **Example**: `ctx.rgba(1.0, 0.0, 0.0, 0.5)`

## Path Construction Functions

### `move_to(x, y)`

Moves the current drawing position without drawing.

- **Parameters**: `x`, `y` - Target coordinates (float)
- **Returns**: `null`
- **Example**: `ctx.move_to(100, 100)`

### `line_to(x, y)`

Draws a line from current position to the specified coordinates.

- **Parameters**: `x`, `y` - Target coordinates (float)
- **Returns**: `null`
- **Example**: `ctx.line_to(200, 150)`

### `quadratic_to(x1, y1, x2, y2)`

Draws a quadratic Bezier curve from current position.

- **Parameters**:
  - `x1`, `y1` - Control point
  - `x2`, `y2` - End point
- **Returns**: `null`
- **Example**: `ctx.quadratic_to(150, 50, 200, 150)`

### `cubic_to(x1, y1, x2, y2, x3, y3)`

Draws a cubic Bezier curve from current position.

- **Parameters**:
  - `x1`, `y1` - First control point
  - `x2`, `y2` - Second control point
  - `x3`, `y3` - End point
- **Returns**: `null`
- **Example**: `ctx.cubic_to(100, 50, 200, 50, 250, 150)`

### `close_path()`

Closes the current path by connecting the last point to the first.

- **Returns**: `null`
- **Example**: `ctx.close_path()`

## Shape Functions

### `rect(x, y, width, height)`

Creates a rectangle path at the specified position.

- **Parameters**:
  - `x`, `y` - Top-left corner position
  - `width`, `height` - Rectangle dimensions
- **Returns**: `null`
- **Example**: `ctx.rect(50, 50, 200, 100)`

### `roundrect(x, y, width, height, rx, ry)`

Creates a rounded rectangle path.

- **Parameters**:
  - `x`, `y` - Top-left corner position
  - `width`, `height` - Rectangle dimensions
  - `rx`, `ry` - Corner radius in X and Y directions
- **Returns**: `null`
- **Example**: `ctx.roundrect(50, 50, 200, 100, 20, 20)`

### `circle(x, y, radius)`

Creates a circle path at the specified center.

- **Parameters**:
  - `x`, `y` - Center position
  - `radius` - Circle radius
- **Returns**: `null`
- **Example**: `ctx.circle(150, 100, 50)`

### `arc(x, y, radius, start_angle, end_angle)`

Creates a circular arc path.

- **Parameters**:
  - `x`, `y` - Center position
  - `radius` - Arc radius
  - `start_angle` - Starting angle in radians
  - `end_angle` - Ending angle in radians
- **Returns**: `null`
- **Example**: `ctx.arc(150, 100, 50, 0, 3.14)`

### `elliptical_arc(x, y, rx, ry, start_angle, end_angle)`

Creates an elliptical arc path.

- **Parameters**:
  - `x`, `y` - Center position
  - `rx`, `ry` - Radius in X and Y directions
  - `start_angle` - Starting angle in radians
  - `end_angle` - Ending angle in radians
- **Returns**: `null`
- **Example**: `ctx.elliptical_arc(150, 100, 60, 40, 0, 3.14)`

### `line(x1, y1, x2, y2)`

Creates a line segment path.

- **Parameters**: Start and end coordinates
- **Returns**: `null`
- **Example**: `ctx.line(50, 50, 200, 150)`

### `point(x, y)`

Draws a single point at the specified coordinates.

- **Parameters**: `x`, `y` - Point position
- **Returns**: `null`
- **Example**: `ctx.point(150, 100)`

## Drawing Functions

### `stroke()`

Draws the outline of the current path(s) with the current color and line width.

- **Returns**: `null`
- **Example**: `ctx.stroke()`

### `fill()`

Fills the current path(s) with the current color.

- **Returns**: `null`
- **Example**: `ctx.fill()`

### `clear([r, g, b, a])` or `clear(hex_color)`

Clears the context with the specified color.

- **Parameters**: RGB values, RGBA values, or hex color string
- **Returns**: `null`
- **Examples**:
  - `ctx.clear()` - Clears with current color
  - `ctx.clear(0.2, 0.2, 0.3)` - Clear with RGB
  - `ctx.clear(0.2, 0.2, 0.3, 1.0)` - Clear with RGBA
  - `ctx.clear("#1a1a2e")` - Clear with hex color

## Transform Functions

### `push()`

Saves the current transformation matrix onto the stack.

- **Returns**: `null`
- **Example**: `ctx.push()`

### `pop()`

Restores the most recently saved transformation matrix.

- **Returns**: `null`
- **Example**: `ctx.pop()`

### `translate(x, y)`

Translates the coordinate system.

- **Parameters**: `x`, `y` - Translation amounts
- **Returns**: `null`
- **Example**: `ctx.translate(100, 50)`

### `scale(sx, sy)`

Scales the coordinate system.

- **Parameters**: `sx`, `sy` - Scale factors
- **Returns**: `null`
- **Example**: `ctx.scale(2.0, 1.5)`

### `rotate(angle)`

Rotates the coordinate system.

- **Parameters**: `angle` - Rotation angle in radians
- **Returns**: `null`
- **Example**: `ctx.rotate(0.785)` // 45 degrees

## Style Functions

### `set_line_width(width)`

Sets the line width for subsequent stroke operations.

- **Parameters**: `width` - Line width in pixels (float)
- **Returns**: `null`
- **Example**: `ctx.set_line_width(2.5)`

## Image Functions

### `load_image(data)`

Loads an image from file path or byte data.

- **Parameters**: `data` - File path string or byte array
- **Returns**: Image object with `id`, `width`, and `height` properties
- **Example**:
  ```go
  img := ctx.load_image("sprite.png")
  // or
  img_data := embed("sprite.png")
  img := ctx.load_image(img_data)
  ```

### `draw_image(img, x, y)`

Draws an image at the specified position.

- **Parameters**:
  - `img` - Image object from `load_image()`, byte array, or file path string
  - `x`, `y` - Position to draw the image
- **Returns**: `null`
- **Example**:
  ```go
  img := ctx.load_image("sprite.png")
  ctx.draw_image(img, 100, 50)
  // or directly with bytes
  ctx.draw_image(embed("sprite.png"), 100, 50)
  ```

### `draw_image_rect(img, sx, sy, sw, sh, dx, dy, dw, dh)`

Draws a sub-rectangle of an image.

- **Parameters**:
  - `img` - Image object or byte data
  - `sx`, `sy` - Source rectangle position
  - `sw`, `sh` - Source rectangle dimensions
  - `dx`, `dy` - Destination position
  - `dw`, `dh` - Destination dimensions
- **Returns**: `null`
- **Example**:
  ```go
  ctx.draw_image_rect(img, 10, 10, 32, 32, 100, 50, 64, 64)
  ```

## Text Functions

### `load_font(path, size)`

Loads a font file for text rendering.

- **Parameters**:
  - `path` - Font file path (string)
  - `size` - Font size in points (float)
- **Returns**: `null`
- **Example**: `ctx.load_font("Arial.ttf", 24)`

### `set_font_size(size)`

Changes the current font size.

- **Parameters**: `size` - New font size in points (float)
- **Returns**: `null`
- **Example**: `ctx.set_font_size(32)`

### `measure_string(text)`

Measures the dimensions of the specified text.

- **Parameters**: `text` - Text to measure (string)
- **Returns**: Map with `width` and `height` properties (floats)
- **Example**: `measurements := ctx.measure_string("Hello World")`

### `draw_string(x, y, text)`

Draws text at the specified position (baseline alignment).

- **Parameters**:
  - `x`, `y` - Position for the text (baseline)
  - `text` - Text to draw (string)
- **Returns**: `null`
- **Example**: `ctx.draw_string(100, 100, "Hello World")`

### `draw_string_anchored(x, y, text, anchor_x, anchor_y)`

Draws text with anchor point alignment.

- **Parameters**:
  - `x`, `y` - Position for the text
  - `text` - Text to draw (string)
  - `anchor_x` - X anchor (0.0 = left, 0.5 = center, 1.0 = right)
  - `anchor_y` - Y anchor (0.0 = top, 0.5 = center, 1.0 = bottom)
- **Returns**: `null`
- **Example**: `ctx.draw_string_anchored(400, 300, "Centered", 0.5, 0.5)`

### `draw_string_wrapped(x, y, text, width[, line_spacing])`

Draws wrapped text within a specified width.

- **Parameters**:
  - `x`, `y` - Position for the text (top-left)
  - `text` - Text to draw (string)
  - `width` - Maximum line width in pixels (float)
  - `line_spacing` - Optional line spacing multiplier (default: 1.5)
- **Returns**: `null`
- **Example**: `ctx.draw_string_wrapped(50, 50, long_text, 200, 1.8)`

## Image Encoding Functions

### `encode([format])`

Encodes the current context as an image.

- **Parameters**: `format` - Optional output format: `"png"` or `"jpeg"` (default: `"png"`)
- **Returns**: Byte array containing encoded image data
- **Example**: 
  ```go
  png_data := ctx.encode()
  jpeg_data := ctx.encode("jpeg")
  ```

### `save(path[, format])`

Saves the current context to a file.

- **Parameters**:
  - `path` - Output file path (string)
  - `format` - Optional format: `"png"` or `"jpeg"` (auto-detected from extension)
- **Returns**: `null`
- **Example**:
  ```go
  ctx.save("output.png")
  ctx.save("output.jpg", "jpeg")
  ```

### `image()`

Returns the current context as an Image object.

- **Returns**: Image object for further processing
- **Example**: `img := ctx.image()`

## Window Event Functions

### `on_draw(callback)`

Registers a callback function for window rendering.

- **Parameters**: `callback` - Function to call when the window needs to redraw
- **Returns**: `null`
- **Example**:
  ```go
  win.on_draw(fn() {
      win.clear("#1a1a2e")
      win.rect(100, 100, 200, 150)
      win.fill()
  })
  ```

### `on_update(callback)`

Registers a callback function for continuous updates.

- **Parameters**: `callback` - Function to call on each frame update
- **Returns**: `null`
- **Example**:
  ```go
  win.on_update(fn() {
      // Update game logic here
  })
  ```

### `on_key(callback)`

Registers a callback for keyboard events.

- **Parameters**: `callback` - Function receiving (key, x, y)
- **Returns**: `null`
- **Example**:
  ```go
  win.on_key(fn(key, x, y) {
      if key == "escape" {
          // Handle escape key
      }
  })
  ```

Special keys are reported as strings: `"left"`, `"right"`, `"up"`, `"down"`, `"page_up"`, `"page_down"`, `"home"`, `"end"`, `"insert"`, `"f1"` through `"f12"`.

### `on_mouse(callback)`

Registers a callback for mouse button events.

- **Parameters**: `callback` - Function receiving (button, action, x, y)
- **Returns**: `null`
- **Example**:
  ```go
  win.on_mouse(fn(button, action, x, y) {
      if button == "left" && action == "down" {
          // Handle left click
      }
  })
  ```

Button names: `"left"`, `"middle"`, `"right"`
Actions: `"down"`, `"up"`

### `on_mouse_move(callback)`

Registers a callback for mouse movement events.

- **Parameters**: `callback` - Function receiving (x, y)
- **Returns**: `null`
- **Example**:
  ```go
  win.on_mouse_move(fn(x, y) {
      // Update cursor position
  })
  ```

### `run()`

Starts the window event loop and displays the window.

- **Returns**: `null` (blocks until window closes)
- **Example**: `win.run()`

### `width`

Returns the current width of the context/window.

- **Type**: Integer property
- **Example**: `w := ctx.width`

### `height`

Returns the current height of the context/window.

- **Type**: Integer property
- **Example**: `h := ctx.height`

## Complete Example

```go
import "graphics"

// Offscreen rendering example
ctx := graphics.new_context(400, 300)
ctx.hex("#2a2a4a")
ctx.clear()
ctx.hex("#00ff88")
ctx.set_line_width(3)
ctx.roundrect(50, 50, 300, 200, 20, 20)
ctx.stroke()
ctx.hex("#ff6644")
ctx.circle(200, 150, 60)
ctx.fill()
ctx.save("output.png")

// Window example with animation
win := graphics.new_window(800, 600, "Graphics Demo")
angle := 0.0

win.on_draw(fn() {
    win.clear("#1a1a2e")
    win.push()
    win.translate(400, 300)
    win.rotate(angle)
    win.hex("#ff6b35")
    win.rect(-50, -50, 100, 100)
    win.fill()
    win.pop()
})

win.on_update(fn() {
    angle += 0.01
})

win.on_key(fn(key, x, y) {
    if key == "escape" {
        // Exit (window run loop will continue)
    }
})

win.run()
```