# Stdlib `glu`

The `glu` module provides GLU (OpenGL Utility Library) bindings for Tender. It offers higher-level utility functions for OpenGL programming, including NURBS, tessellation, quadric surfaces, and projection utilities.

---

## Core Functions

### `build_2d_mipmaps(target, internal_format, width, height, format, type, pixels)`

Builds a 2D mipmap pyramid from a single image.

- **Parameters**:
  - `target` - Texture target (`gl.TEXTURE_2D`, etc.)
  - `internal_format` - Internal texture format (int)
  - `width` - Image width (int)
  - `height` - Image height (int)
  - `format` - Pixel format (`gl.RGB`, `gl.RGBA`, etc.)
  - `type` - Pixel data type (`gl.UNSIGNED_BYTE`, etc.)
  - `pixels` - Image data (bytes) or `null`
- **Returns**: Error code (int), `0` on success
- **Example**:
  ```go
  err := glu.build_2d_mipmaps(gl.TEXTURE_2D, gl.RGB, 512, 512, gl.RGB, gl.UNSIGNED_BYTE, image_bytes)
  if err != 0 {
      println("Mipmap generation failed:", glu.error_string(err))
  }
  ```

### `look_at(eye_x, eye_y, eye_z, center_x, center_y, center_z, up_x, up_y, up_z)`

Defines a viewing transformation using the camera position, look-at target, and up vector.

- **Parameters**: All coordinates as float64
- **Returns**: `null`
- **Example**:
  ```go
  glu.look_at(0.0, 0.0, 5.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)
  ```

### `perspective(fovy, aspect, z_near, z_far)`

Sets up a perspective projection matrix.

- **Parameters**:
  - `fovy` - Field of view angle in degrees (float64)
  - `aspect` - Aspect ratio (width/height) (float64)
  - `z_near` - Near clipping plane distance (float64)
  - `z_far` - Far clipping plane distance (float64)
- **Returns**: `null`
- **Example**:
  ```go
  glu.perspective(45.0, 800.0/600.0, 0.1, 100.0)
  ```

### `project(obj_x, obj_y, obj_z, model_matrix, proj_matrix, viewport)`

Projects an object coordinate to window coordinates.

- **Parameters**:
  - `obj_x`, `obj_y`, `obj_z` - Object coordinates (float64)
  - `model_matrix` - 16 float64 values (column-major)
  - `proj_matrix` - 16 float64 values (column-major)
  - `viewport` - 4 int32 values `[x, y, width, height]`
- **Returns**: Array of 3 floats `[win_x, win_y, win_z]`
- **Example**:
  ```go
  model := [0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0]
  proj := [0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0]
  view := [0, 0, 800, 600]
  result := glu.project(0.0, 0.0, 0.0, model..., proj..., view...)
  println("Window position:", result[0], result[1], result[2])
  ```

### `unproject(win_x, win_y, win_z, model_matrix, proj_matrix, viewport)`

Projects window coordinates to object coordinates.

- **Parameters**:
  - `win_x`, `win_y`, `win_z` - Window coordinates (float64)
  - `model_matrix` - 16 float64 values (column-major)
  - `proj_matrix` - 16 float64 values (column-major)
  - `viewport` - 4 int32 values `[x, y, width, height]`
- **Returns**: Array of 3 floats `[obj_x, obj_y, obj_z]`
- **Example**:
  ```go
  result := glu.unproject(mouse_x, screen_height - mouse_y, 0.0, model..., proj..., view...)
  println("Object position:", result[0], result[1], result[2])
  ```

### `error_string(error_code)`

Returns a human-readable description of a GLU error code.

- **Parameters**: `error_code` - GLU error code (uint32)
- **Returns**: Error description string
- **Example**:
  ```go
  err := glu.build_2d_mipmaps(...)
  if err != 0 {
      println("Error:", glu.error_string(err))
  }
  ```

---

## Quadric Functions

Quadric objects are used to render spheres, cylinders, and disks.

### `new_quadric()`

Creates a new quadric object.

- **Returns**: Quadric handle (int)
- **Example**:
  ```go
  q := glu.new_quadric()
  ```

### `sphere(quadric, radius, slices, stacks)`

Draws a sphere.

- **Parameters**:
  - `quadric` - Quadric handle (int)
  - `radius` - Sphere radius (float32)
  - `slices` - Number of longitudinal divisions (int)
  - `stacks` - Number of latitudinal divisions (int)
- **Returns**: `null`
- **Example**:
  ```go
  q := glu.new_quadric()
  glu.sphere(q, 1.0, 32, 32)
  ```

### `cylinder(quadric, base_radius, top_radius, height, slices, stacks)`

Draws a cylinder.

- **Parameters**:
  - `quadric` - Quadric handle (int)
  - `base_radius` - Bottom radius (float32)
  - `top_radius` - Top radius (float32)
  - `height` - Cylinder height (float32)
  - `slices` - Number of longitudinal divisions (int)
  - `stacks` - Number of latitudinal divisions (int)
- **Returns**: `null`
- **Example**:
  ```go
  glu.cylinder(q, 1.0, 0.5, 2.0, 32, 16)
  ```

### `disk(quadric, inner_radius, outer_radius, slices, loops)`

Draws a disk.

- **Parameters**:
  - `quadric` - Quadric handle (int)
  - `inner_radius` - Inner radius (float32)
  - `outer_radius` - Outer radius (float32)
  - `slices` - Number of radial divisions (int)
  - `loops` - Number of concentric rings (int)
- **Returns**: `null`
- **Example**:
  ```go
  glu.disk(q, 0.5, 2.0, 32, 16)
  ```

### `partial_disk(quadric, inner_radius, outer_radius, slices, loops, start_angle, sweep_angle)`

Draws a partial disk.

- **Parameters**:
  - `quadric` - Quadric handle (int)
  - `inner_radius` - Inner radius (float32)
  - `outer_radius` - Outer radius (float32)
  - `slices` - Number of radial divisions (int)
  - `loops` - Number of concentric rings (int)
  - `start_angle` - Starting angle in degrees (float32)
  - `sweep_angle` - Sweep angle in degrees (float32)
- **Returns**: `null`
- **Example**:
  ```go
  glu.partial_disk(q, 0.5, 2.0, 32, 16, 0.0, 180.0) // Half disk
  ```

---

## Tessellation Constants

### Tessellation Callback Types
- `glu.TESS_BEGIN_DATA`
- `glu.TESS_VERTEX_DATA`
- `glu.TESS_END_DATA`
- `glu.TESS_ERROR_DATA`
- `glu.TESS_EDGE_FLAG_DATA`
- `glu.TESS_COMBINE_DATA`

### Tessellation Properties
- `glu.TESS_WINDING_RULE`
- `glu.TESS_BOUNDARY_ONLY`
- `glu.TESS_TOLERANCE`

### Winding Rules
- `glu.TESS_WINDING_ODD`
- `glu.TESS_WINDING_NONZERO`
- `glu.TESS_WINDING_POSITIVE`
- `glu.TESS_WINDING_NEGATIVE`
- `glu.TESS_WINDING_ABS_GEQ_TWO`

---

## Complete Examples

### Example 1: Perspective Camera Setup

```go
import "gl"
import "glu"
import "glut"

glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE | glut.DEPTH)
glut.init_window_size(400, 400)
glut.create_window("GLU Demo")

glut.reshape_func(fn(w, h) {
    gl.viewport(0, 0, w, h)
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    glu.perspective(45.0, float(w)/float(h), 0.1, 100.0)
    gl.matrix_mode(gl.MODELVIEW)
})

glut.display_func(fn() {
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    gl.load_identity()
    glu.look_at(0.0, 2.0, 5.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)
    
    // Draw a sphere
    q := glu.new_quadric()
    gl.color3f(1.0, 0.5, 0.0)
    glu.sphere(q, 1.0, 32, 32)
    
    glut.swap_buffers()
})

glut.main_loop()
```

### Example 2: Rendering Various Quadrics

```go
import "gl"
import "glut"
import "glu"

glut.init()
gl.init()
glut.init_display_mode(glut.RGB | glut.DOUBLE | glut.DEPTH) // Use DOUBLE and DEPTH
glut.init_window_size(400, 400)
glut.create_window("GLU Test")

// Set up projection matrix in reshape callback
glut.reshape_func(fn(w, h) {
    gl.viewport(0, 0, w, h)
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    glu.perspective(45.0, float(w)/float(h), 0.1, 100.0)
    gl.matrix_mode(gl.MODELVIEW)
})

fn draw_scene() {
    q := glu.new_quadric()
    
    // Draw a sphere at (-3, 0, 0)
    gl.push_matrix()
    gl.translatef(-3.0, 0.0, 0.0)
    gl.color3f(1.0, 0.0, 0.0)
    glu.sphere(q, 0.8, 32, 32)
    gl.pop_matrix()
    
    // Draw a cylinder at (0, 0, 0)
    gl.push_matrix()
    gl.translatef(0.0, 0.0, 0.0)
    gl.color3f(0.0, 1.0, 0.0)
    glu.cylinder(q, 0.8, 0.3, 1.5, 32, 16)
    gl.pop_matrix()
    
    // Draw a disk at (3, 0, 0)
    gl.push_matrix()
    gl.translatef(3.0, 0.0, 0.0)
    gl.color3f(0.0, 0.0, 1.0)
    glu.disk(q, 0.2, 0.8, 32, 16)
    gl.pop_matrix()
    
    // Draw a partial disk at (0, 0, 3)
    gl.push_matrix()
    gl.translatef(0.0, 0.0, 3.0)
    gl.color3f(1.0, 1.0, 0.0)
    glu.partial_disk(q, 0.0, 0.8, 32, 16, 0.0, 270.0)
    gl.pop_matrix()
}

// Register with GLUT callbacks
glut.display_func(fn() {
    gl.clear_color(0.2, 0.3, 0.4, 1.0) // Set clear color
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
    gl.matrix_mode(gl.MODELVIEW)
    gl.load_identity()
    glu.look_at(0.0, 2.0, 6.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)
    
    // Enable depth testing
    gl.enable(gl.DEPTH_TEST)
    gl.depth_func(gl.LESS)
    
    draw_scene()
    glut.swap_buffers()
})
    
glut.main_loop()
```
