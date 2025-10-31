## Stdlib `canvas`

The `canvas` module provides functionalities for creating and manipulating graphical elements.

### Functions

#### `new_context(width, height)`

Creates a new canvas context with the specified width and height.

- `width`: Width of the canvas context.
- `height`: Height of the canvas context.

#### `load_image(image_bytes)`

Loads an image from the provided byte array.

- `image_bytes`: Byte array representing the image data.

#### `radians(degrees)`

Converts degrees to radians.

- `degrees`: Angle in degrees.

#### `degrees(radians)`

Converts radians to degrees.

- `radians`: Angle in radians.

#### `new_window(width, height, title, update_function)`

Creates a new window with the specified width, height, and title. The `update_function` is called to update the content of the window.

- `width`: Width of the window.
- `height`: Height of the window.
- `title`: Title of the window.
- `update_function`: Function responsible for updating the content of the window.

```go
import "canvas"

canvas.new_window(600, 400, "Hello World", fn(w){
	for {}
})
```

#### `new_window(options, update_function)`

*Note: Only supports on `windows` platform.*

Creates a new window with the specified options. The `update_function` is called to update the content of the window.

```go
import "canvas"

var options = {
	width: 600,
	height: 400,
	title: "Hello World",
	fullscreen: false,
	borderless: false,
	top_most: false,
	no_scaling: false,
	x: 200,
	y: 10
}

canvas.new_window(options, fn(w){
	for {}
})
```


### canvas Context Functions

#### Drawing Functions

- `drawimage(image_bytes, x, y)`
- `drawimage_anchored(image_bytes, x, y, anchor_x, anchor_y)`
- `savepng(filename)`
- `point(x, y)`
- `line(x1, y1, x2, y2)`
- `rect(x, y, width, height)`
- `polygon(sides, x_center, y_center, radius, rotation)`
- `rounderect(x, y, width, height, radius)`
- `circle(x, y, radius)`
- `arc(x, y, radius, start_angle, end_angle)`
- `ellipse(x, y, x_radius, y_radius)`
- `ellipsearc(x, y, x_radius, y_radius, start_angle, end_angle)`

#### Style Functions

- `set_pixel(x, y)`
- `rgb(r, g, b)`
- `rgba(r, g, b, a)`
- `rgb255(r, g, b)`
- `rgba255(r, g, b, a)`
- `hex(color)`
- `linewidth(width)`
- `dashoffset(offset)`
- `dash(...)`

#### Path Functions

- `move_to(x, y)`
- `line_to(x, y)`
- `quadratic_to(x1, y1, x2, y2)`
- `cubic_to(x1, y1, x2, y2, x3, y3)`
- `closepath()`
- `clearpath()`
- `newsubpath()`
- `clear()`

#### Drawing State Functions

- `stroke()`
- `fill()`
- `stroke_preserve()`
- `fill_preserve()`
- `push()`
- `pop()`
- `clip()`
- `clip_preserve()`
- `resetclip()`

#### Text Functions

- `text(text, x, y)`
- `text_anchored(text, x, y, anchor_x, anchor_y)`
- `measure_text(text)`
- `measure_multiline_text(text)`
- `load_fontface(font_path, size)`
- `fontface(font_data, size)`
- `fontheight()`

#### Transformation Functions

- `identity()`
- `translate(x, y)`
- `scale(scale_x, scale_y)`
- `rotate(angle)`
- `shear(shear_x, shear_y)`
- `scaleabout(scale_x, scale_y, x, y)`
- `rotateabout(angle, x, y)`
- `shearabout(shear_x, shear_y, x, y)`
- `transform_point(x, y)`

#### Other Functions

- `height()`
- `width()`
- `wordwrap(text, max_width)`

#### Image Functions

- `image()`

### Example Usage

```go
import "canvas"

ctx := canvas.new_context(200, 200)

ctx.hex(`#121212`)
ctx.clear()
ctx.hex("#0f0")
ctx.dash(4)
ctx.rect(50, 50, 100, 100)
ctx.stroke()	

ctx.save_png("output.png")
```
