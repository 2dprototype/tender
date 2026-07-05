# Stdlib `gl`

The `gl` module provides OpenGL bindings for 3D graphics rendering. It exposes a comprehensive set of OpenGL functions and constants for creating 3D applications with immediate mode rendering, display lists, vertex arrays, lighting, texturing, and more.

## Initialization

### `init()`

Initializes OpenGL context.

- **Returns**: `null`
- **Example**: `gl.init()`

---

## Matrix Operations

### `matrix_mode(mode)`

Sets the current matrix mode.

- **Parameters**: 
  - `mode` - Matrix mode constant (`gl.MODELVIEW`, `gl.PROJECTION`, or `gl.TEXTURE`)
- **Returns**: `null`
- **Example**: `gl.matrix_mode(gl.PROJECTION)`

### `load_identity()`

Replaces the current matrix with the identity matrix.

- **Returns**: `null`
- **Example**: `gl.load_identity()`

### `push_matrix()`

Pushes the current matrix stack.

- **Returns**: `null`
- **Example**: `gl.push_matrix()`

### `pop_matrix()`

Pops the current matrix stack.

- **Returns**: `null`
- **Example**: `gl.pop_matrix()`

### `mult_matrixf(m0, m1, m2, m3, m4, m5, m6, m7, m8, m9, m10, m11, m12, m13, m14, m15)`

Multiplies the current matrix by a 4x4 matrix.

- **Parameters**: 16 float values in column-major order
- **Returns**: `null`
- **Example**: 
  ```go
  gl.mult_matrixf(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1)
  ```

### `load_matrixf(m0, m1, ... m15)`

Replaces the current matrix with a 4x4 matrix.

- **Parameters**: 16 float values in column-major order
- **Returns**: `null`

### `translatef(x, y, z)`

Multiplies the current matrix by a translation matrix.

- **Parameters**: `x`, `y`, `z` - Translation values (float)
- **Returns**: `null`
- **Example**: `gl.translatef(1.0, 2.0, 3.0)`

### `rotatef(angle, x, y, z)`

Multiplies the current matrix by a rotation matrix.

- **Parameters**: `angle` - Rotation angle in degrees, `x`, `y`, `z` - Rotation axis
- **Returns**: `null`
- **Example**: `gl.rotatef(45.0, 0.0, 1.0, 0.0)`

### `scalef(x, y, z)`

Multiplies the current matrix by a scaling matrix.

- **Parameters**: `x`, `y`, `z` - Scale factors (float)
- **Returns**: `null`
- **Example**: `gl.scalef(2.0, 2.0, 2.0)`

### `ortho(left, right, bottom, top, zNear, zFar)`

Multiplies the current matrix by an orthographic projection matrix.

- **Parameters**: 
  - `left`, `right` - Left and right clipping planes (float)
  - `bottom`, `top` - Bottom and top clipping planes (float)
  - `zNear`, `zFar` - Near and far clipping planes (float)
- **Returns**: `null`
- **Example**: `gl.ortho(-1, 1, -1, 1, -1, 1)`

### `frustum(left, right, bottom, top, zNear, zFar)`

Multiplies the current matrix by a perspective projection matrix.

- **Parameters**:
  - `left`, `right` - Left and right clipping planes (float)
  - `bottom`, `top` - Bottom and top clipping planes (float)
  - `zNear`, `zFar` - Near and far clipping planes (float, must be positive)
- **Returns**: `null`
- **Example**: `gl.frustum(-1, 1, -1, 1, 1, 100)`

---

## Viewport and Scissor

### `viewport(x, y, width, height)`

Sets the viewport.

- **Parameters**: `x`, `y` - Lower-left corner, `width`, `height` - Viewport dimensions (int)
- **Returns**: `null`
- **Example**: `gl.viewport(0, 0, 800, 600)`

### `scissor(x, y, width, height)`

Sets the scissor rectangle.

- **Parameters**: `x`, `y` - Lower-left corner, `width`, `height` - Rectangle dimensions (int)
- **Returns**: `null`

---

## Clearing

### `clear(mask)`

Clears buffers.

- **Parameters**: `mask` - Bitwise OR of clear bits (`gl.COLOR_BUFFER_BIT`, `gl.DEPTH_BUFFER_BIT`, `gl.STENCIL_BUFFER_BIT`)
- **Returns**: `null`
- **Example**: `gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)`

### `clear_color(r, g, b, a)`

Sets the color buffer clear value.

- **Parameters**: `r`, `g`, `b`, `a` - Color components (float, 0.0-1.0)
- **Returns**: `null`
- **Example**: `gl.clear_color(0.2, 0.3, 0.4, 1.0)`

### `clear_depth(depth)`

Sets the depth buffer clear value.

- **Parameters**: `depth` - Depth value (float)
- **Returns**: `null`

### `clear_stencil(s)`

Sets the stencil buffer clear value.

- **Parameters**: `s` - Stencil value (int)
- **Returns**: `null`

### `clear_accum(r, g, b, a)`

Sets the accumulation buffer clear value.

- **Parameters**: `r`, `g`, `b`, `a` - Color components (float)
- **Returns**: `null`

---

## Primitives

### `begin(mode)`

Begins a primitive drawing mode.

- **Parameters**: `mode` - Primitive type (`gl.POINTS`, `gl.LINES`, `gl.TRIANGLES`, `gl.QUADS`, etc.)
- **Returns**: `null`
- **Example**:
  ```go
  gl.begin(gl.TRIANGLES)
  gl.vertex3f(0, 1, 0)
  gl.vertex3f(-1, -1, 0)
  gl.vertex3f(1, -1, 0)
  gl.end()
  ```

### `end()`

Ends a primitive drawing mode.

- **Returns**: `null`

### `vertex2f(x, y)`
### `vertex3f(x, y, z)`
### `vertex4f(x, y, z, w)`
### `vertex2d(x, y)`
### `vertex3d(x, y, z)`
### `vertex4d(x, y, z, w)`
### `vertex2i(x, y)`
### `vertex3i(x, y, z)`

Specifies a vertex (float, double, or integer versions).

- **Returns**: `null`

---

## Colors

### `color3f(r, g, b)`
### `color4f(r, g, b, a)`
### `color3ub(r, g, b)`
### `color4ub(r, g, b, a)`

Sets the current color (float or unsigned byte versions).

- **Parameters**: Color components (float: 0.0-1.0, byte: 0-255)
- **Returns**: `null`
- **Example**: `gl.color3f(1.0, 0.0, 0.0)`

---

## Normals

### `normal3f(nx, ny, nz)`

Sets the current normal vector.

- **Parameters**: `nx`, `ny`, `nz` - Normal components (float)
- **Returns**: `null`
- **Example**: `gl.normal3f(0.0, 0.0, 1.0)`

---

## Texture Coordinates

### `tex_coord2f(s, t)`

Sets the current texture coordinates.

- **Parameters**: `s`, `t` - Texture coordinates (float)
- **Returns**: `null`
- **Example**: `gl.tex_coord2f(0.5, 0.5)`

---

## Enable/Disable

### `enable(cap)`
### `disable(cap)`
### `is_enabled(cap)`

Enables, disables, or checks OpenGL capabilities.

- **Parameters**: `cap` - Capability constant (`gl.BLEND`, `gl.DEPTH_TEST`, `gl.LIGHTING`, etc.)
- **Returns**: `null` for `enable`/`disable`, `bool` for `is_enabled`
- **Example**: `gl.enable(gl.DEPTH_TEST)`

---

## Blending

### `blend_func(sfactor, dfactor)`

Sets the pixel blending function.

- **Parameters**: `sfactor` - Source factor, `dfactor` - Destination factor
- **Returns**: `null`
- **Example**: `gl.blend_func(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)`

### `blend_equation(mode)`

Sets the blend equation.

- **Parameters**: `mode` - Blend equation (`gl.FUNC_ADD`, `gl.FUNC_SUBTRACT`, etc.)
- **Returns**: `null`

### `blend_color(r, g, b, a)`

Sets the blend color.

- **Parameters**: `r`, `g`, `b`, `a` - Color components (float)
- **Returns**: `null`

---

## Depth Test

### `depth_func(func)`

Sets the depth test function.

- **Parameters**: `func` - Depth function (`gl.LESS`, `gl.LEQUAL`, `gl.GEQUAL`, etc.)
- **Returns**: `null`
- **Example**: `gl.depth_func(gl.LESS)`

### `depth_mask(flag)`

Enables or disables depth buffer writing.

- **Parameters**: `flag` - Boolean or integer
- **Returns**: `null`

### `depth_range(near, far)`

Sets the depth range.

- **Parameters**: `near`, `far` - Depth range values (float)
- **Returns**: `null`

---

## Alpha Test

### `alpha_func(func, ref)`

Sets the alpha test function.

- **Parameters**: `func` - Alpha function, `ref` - Reference value (float)
- **Returns**: `null`

---

## Culling

### `cull_face(mode)`

Specifies which faces to cull.

- **Parameters**: `mode` - `gl.FRONT`, `gl.BACK`, or `gl.FRONT_AND_BACK`
- **Returns**: `null`
- **Example**: `gl.cull_face(gl.BACK)`

### `front_face(mode)`

Sets which face is considered front-facing.

- **Parameters**: `mode` - `gl.CW` (clockwise) or `gl.CCW` (counter-clockwise)
- **Returns**: `null`

---

## Polygon Modes

### `polygon_mode(face, mode)`

Sets the polygon rendering mode.

- **Parameters**: `face` - `gl.FRONT_AND_BACK`, `mode` - `gl.POINT`, `gl.LINE`, or `gl.FILL`
- **Returns**: `null`

### `polygon_offset(factor, units)`

Sets the polygon offset.

- **Parameters**: `factor`, `units` - Offset values (float)
- **Returns**: `null`

---

## Line & Point Sizes

### `line_width(width)`

Sets the line width.

- **Parameters**: `width` - Line width (float)
- **Returns**: `null`

### `point_size(size)`

Sets the point size.

- **Parameters**: `size` - Point size (float)
- **Returns**: `null`

---

## Stencil

### `stencil_func(func, ref, mask)`

Sets the stencil test function.

- **Parameters**: `func` - Stencil function, `ref` - Reference value (int), `mask` - Mask (uint)
- **Returns**: `null`

### `stencil_mask(mask)`

Sets the stencil write mask.

- **Parameters**: `mask` - Mask value (uint)
- **Returns**: `null`

### `stencil_op(fail, zfail, zpass)`

Sets the stencil operation.

- **Parameters**: `fail` - Operation on stencil test failure, `zfail` - Operation on depth test failure, `zpass` - Operation on depth test pass
- **Returns**: `null`

---

## State Saving

### `push_attrib(mask)`
### `pop_attrib()`

Pushes and pops attribute stack.

- **Parameters**: `mask` - Attribute mask (uint)
- **Returns**: `null`

### `push_client_attrib(mask)`
### `pop_client_attrib()`

Pushes and pops client attribute stack.

- **Parameters**: `mask` - Client attribute mask (uint)
- **Returns**: `null`

---

## Shading

### `shade_model(mode)`

Sets the shading model.

- **Parameters**: `mode` - `gl.FLAT` or `gl.SMOOTH`
- **Returns**: `null`

---

## Hints

### `hint(target, mode)`

Sets hint parameters.

- **Parameters**: `target` - Hint target, `mode` - `gl.DONT_CARE`, `gl.FASTEST`, or `gl.NICEST`
- **Returns**: `null`

---

## Lighting

### `lightf(light, pname, param)`
### `lighti(light, pname, param)`
### `lightfv(light, pname, ...params)`

Sets light source parameters (float, int, or array versions).

- **Parameters**: `light` - Light number (`gl.LIGHT0`, `gl.LIGHT1`, etc.), `pname` - Parameter name, `param(s)` - Parameter value(s)
- **Returns**: `null`
- **Example**:
  ```go
  gl.lightfv(gl.LIGHT0, gl.POSITION, 1.0, 1.0, 1.0, 0.0)
  gl.lightf(gl.LIGHT0, gl.SPOT_CUTOFF, 45.0)
  ```

### `light_modelf(pname, param)`
### `light_modeli(pname, param)`

Sets lighting model parameters.

- **Returns**: `null`

### `materialf(face, pname, param)`
### `materialfv(face, pname, ...params)`

Sets material parameters.

- **Parameters**: `face` - `gl.FRONT`, `gl.BACK`, or `gl.FRONT_AND_BACK`
- **Returns**: `null`
- **Example**: `gl.materialfv(gl.FRONT_AND_BACK, gl.SPECULAR, 0.8, 0.8, 0.8, 1.0)`

### `color_material(face, mode)`

Makes material colors track the current color.

- **Parameters**: 
  - `face` - `gl.FRONT`, `gl.BACK`, or `gl.FRONT_AND_BACK`
  - `mode` - Which material parameters to track (`gl.AMBIENT`, `gl.DIFFUSE`, `gl.SPECULAR`, etc.)
- **Returns**: `null`
- **Example**: `gl.color_material(gl.FRONT_AND_BACK, gl.DIFFUSE)`

---

## Fog

### `fogf(pname, param)`
### `fogfv(pname, ...params)`

Sets fog parameters.

- **Parameters**: `pname` - Parameter name (`gl.FOG_COLOR`, `gl.FOG_DENSITY`, etc.)
- **Returns**: `null`
- **Example**:
  ```go
  gl.fogf(gl.FOG_DENSITY, 0.05)
  gl.fogfv(gl.FOG_COLOR, 0.5, 0.5, 0.5, 1.0)
  ```

---

## Textures

### `gen_textures(n)`

Generates texture names.

- **Parameters**: `n` - Number of textures to generate (int)
- **Returns**: Array of texture IDs
- **Example**: `textures := gl.gen_textures(1)`

### `bind_texture(target, texture)`

Binds a named texture.

- **Parameters**: `target` - `gl.TEXTURE_2D`, `gl.TEXTURE_CUBE_MAP`, etc., `texture` - Texture ID
- **Returns**: `null`
- **Example**: `gl.bind_texture(gl.TEXTURE_2D, textures[0])`

### `delete_textures(textures)`

Deletes named textures.

- **Parameters**: `textures` - Single texture ID or array of texture IDs
- **Returns**: `null`

### `tex_image1d(target, level, internal_format, width, border, format, type, pixels)`

Sets 1D texture image data.

- **Parameters**:
  - `target` - Texture target
  - `level` - Mipmap level (int)
  - `internal_format` - Internal format (int)
  - `width` - Image width (int)
  - `border` - Border width (int)
  - `format` - Pixel format
  - `type` - Pixel data type
  - `pixels` - Image data (bytes) or `null`
- **Returns**: `null`

### `tex_image2d(target, level, internal_format, width, height, border, format, type, pixels)`

Sets 2D texture image data.

- **Parameters**:
  - `target` - Texture target
  - `level` - Mipmap level (int)
  - `internal_format` - Internal format (int)
  - `width`, `height` - Image dimensions (int)
  - `border` - Border width (int)
  - `format` - Pixel format
  - `type` - Pixel data type
  - `pixels` - Image data (bytes) or `null`
- **Returns**: `null`
- **Example**:
  ```go
  gl.tex_image2d(gl.TEXTURE_2D, 0, gl.RGB, 256, 256, 0, gl.RGB, gl.UNSIGNED_BYTE, image_bytes)
  ```

### `tex_image3d(target, level, internal_format, width, height, depth, border, format, type, pixels)`

Sets 3D texture image data.

- **Parameters**:
  - `target` - Texture target
  - `level` - Mipmap level (int)
  - `internal_format` - Internal format (int)
  - `width`, `height`, `depth` - Image dimensions (int)
  - `border` - Border width (int)
  - `format` - Pixel format
  - `type` - Pixel data type
  - `pixels` - Image data (bytes) or `null`
- **Returns**: `null`

### `tex_sub_image1d(target, level, xoffset, width, format, type, pixels)`

Specifies a sub-rectangle of the current 1D texture image.

- **Returns**: `null`

### `tex_sub_image2d(target, level, xoffset, yoffset, width, height, format, type, pixels)`

Specifies a sub-rectangle of the current 2D texture image.

- **Returns**: `null`

### `tex_sub_image3d(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels)`

Specifies a sub-rectangle of the current 3D texture image.

- **Returns**: `null`

### `copy_tex_image2d(target, level, internal_format, x, y, width, height, border)`

Copies pixels from the framebuffer to a 2D texture image.

- **Parameters**:
  - `target` - Texture target
  - `level` - Mipmap level (int)
  - `internal_format` - Internal format (uint)
  - `x`, `y` - Framebuffer coordinates (int)
  - `width`, `height` - Image dimensions (int)
  - `border` - Border width (int)
- **Returns**: `null`

### `copy_tex_sub_image2d(target, level, xoffset, yoffset, x, y, width, height)`

Copies pixels from the framebuffer to a sub-rectangle of the current 2D texture.

- **Returns**: `null`

### `tex_parameterf(target, pname, param)`
### `tex_parameteri(target, pname, param)`
### `tex_parameterfv(target, pname, ...params)`
### `tex_parameteriv(target, pname, ...params)`

Sets texture parameters.

- **Parameters**: `target` - Texture target, `pname` - Parameter name, `param(s)` - Parameter value(s)
- **Returns**: `null`
- **Example**:
  ```go
  gl.tex_parameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
  gl.tex_parameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
  gl.tex_parameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
  ```

### `get_tex_parameteriv(target, pname)`
### `get_tex_parameterfv(target, pname)`

Queries texture parameter values.

- **Parameters**: `target` - Texture target, `pname` - Parameter name
- **Returns**: Parameter value (int or float)
- **Example**: `min_filter := gl.get_tex_parameteriv(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER)`

### `get_tex_image(target, level, format, type, pixels)`

Returns texture image data.

- **Parameters**:
  - `target` - Texture target
  - `level` - Mipmap level (int)
  - `format` - Pixel format
  - `type` - Pixel data type
  - `pixels` - Byte slice to receive the data
- **Returns**: `null`

### `tex_envf(target, pname, param)`
### `tex_envi(target, pname, param)`
### `tex_envfv(target, pname, ...params)`
### `tex_enviv(target, pname, ...params)`

Sets texture environment parameters.

- **Returns**: `null`

### `tex_genf(coord, pname, param)`
### `tex_gend(coord, pname, param)`
### `tex_geni(coord, pname, param)`

Sets texture coordinate generation parameters.

- **Returns**: `null`

---

## Pixel Operations

### `bitmap(width, height, xorig, yorig, xmove, ymove, bitmap)`

Draws a bitmap (monochrome image).

- **Parameters**:
  - `width`, `height` - Bitmap dimensions (int)
  - `xorig`, `yorig` - Origin of bitmap (float)
  - `xmove`, `ymove` - Cursor movement after drawing (float)
  - `bitmap` - Bitmap data (bytes) or `null`
- **Returns**: `null`

### `draw_pixels(width, height, format, type, pixels)`

Draws a rectangle of pixel data.

- **Parameters**:
  - `width`, `height` - Image dimensions (int)
  - `format` - Pixel format
  - `type` - Pixel data type
  - `pixels` - Pixel data (bytes)
- **Returns**: `null`

### `read_pixels(x, y, width, height, format, type, pixels)`

Reads pixels from the framebuffer.

- **Parameters**:
  - `x`, `y` - Lower-left corner (int)
  - `width`, `height` - Rectangle dimensions (int)
  - `format` - Pixel format
  - `type` - Pixel data type
  - `pixels` - Byte slice to receive the data
- **Returns**: `null`
- **Example**:
  ```go
  pixels := make([]byte, 800*600*4)
  gl.read_pixels(0, 0, 800, 600, gl.RGBA, gl.UNSIGNED_BYTE, pixels)
  ```

### `copy_pixels(x, y, width, height, type)`

Copies a block of pixels within the framebuffer.

- **Parameters**:
  - `x`, `y` - Source lower-left corner (int)
  - `width`, `height` - Rectangle dimensions (int)
  - `type` - `gl.COLOR`, `gl.DEPTH`, or `gl.STENCIL`
- **Returns**: `null`

### `pixel_zoom(xfactor, yfactor)`

Sets pixel zoom factors for draw/copy operations.

- **Parameters**: `xfactor`, `yfactor` - Zoom factors (float)
- **Returns**: `null`

### `raster_pos2i(x, y)`
### `raster_pos2f(x, y)`
### `raster_pos3f(x, y, z)`

Sets the current raster position.

- **Parameters**: Position coordinates (int or float)
- **Returns**: `null`
- **Example**: `gl.raster_pos2i(100, 200)`

### `color_mask(red, green, blue, alpha)`

Enables or disables writing of colour components.

- **Parameters**: Boolean values for each colour component
- **Returns**: `null`
- **Example**: `gl.color_mask(true, true, true, false)`

### `logic_op(opcode)`

Sets the logical operation for pixel writes.

- **Parameters**: `opcode` - Logical operation (`gl.CLEAR`, `gl.AND`, `gl.OR`, `gl.XOR`, etc.)
- **Returns**: `null`

### `pixel_storei(pname, param)`

Sets pixel storage modes.

- **Parameters**: `pname` - `gl.PACK_ALIGNMENT`, `gl.UNPACK_ALIGNMENT`, etc., `param` - Value (int)
- **Returns**: `null`

---

## Clipping Planes

### `clip_plane(plane, eq0, eq1, eq2, eq3)`

Defines a clipping plane.

- **Parameters**:
  - `plane` - Plane number (`gl.CLIP_PLANE0` through `gl.CLIP_PLANE5`)
  - `eq0`, `eq1`, `eq2`, `eq3` - Plane equation coefficients (float)
- **Returns**: `null`
- **Example**: `gl.clip_plane(gl.CLIP_PLANE0, 0, 1, 0, 0)`

---

## Vertex Arrays

### `enable_client_state(array)`
### `disable_client_state(array)`

Enables or disables client-side vertex arrays.

- **Parameters**: `array` - `gl.VERTEX_ARRAY`, `gl.NORMAL_ARRAY`, `gl.COLOR_ARRAY`, `gl.TEXTURE_COORD_ARRAY`
- **Returns**: `null`

### `vertex_pointer(size, type, stride, pointer)`
### `normal_pointer(type, stride, pointer)`
### `color_pointer(size, type, stride, pointer)`
### `tex_coord_pointer(size, type, stride, pointer)`

Specifies vertex array pointers.

- **Parameters**: `pointer` - Data (bytes) or `null`
- **Returns**: `null`
- **Example**:
  ```go
  gl.enable_client_state(gl.VERTEX_ARRAY)
  gl.vertex_pointer(3, gl.FLOAT, 0, vertex_data)
  ```

### `draw_arrays(mode, first, count)`

Renders primitives from array data.

- **Parameters**: `mode` - Primitive type, `first` - Starting index (int), `count` - Number of vertices (int)
- **Returns**: `null`

### `draw_elements(mode, count, type, indices)`

Renders primitives from indexed array data.

- **Parameters**: `indices` - Index data (bytes) or `null`
- **Returns**: `null`

---

## Display Lists

### `gen_lists(range)`

Generates display list names.

- **Parameters**: `range` - Number of display lists to generate (int)
- **Returns**: Starting display list ID (int)

### `new_list(list, mode)`

Begins a display list.

- **Parameters**: `list` - Display list ID (uint), `mode` - `gl.COMPILE` or `gl.COMPILE_AND_EXECUTE`
- **Returns**: `null`

### `end_list()`

Ends a display list.

- **Returns**: `null`

### `call_list(list)`

Executes a display list.

- **Parameters**: `list` - Display list ID (uint)
- **Returns**: `null`

### `delete_lists(list, range)`

Deletes display lists.

- **Parameters**: `list` - Starting display list ID (uint), `range` - Number of lists to delete (int)
- **Returns**: `null`

---

## Query & Sync

### `get_error()`

Returns the current OpenGL error code.

- **Returns**: Error code (int)
- **Example**:
  ```go
  err := gl.get_error()
  if err != gl.NO_ERROR {
      println("OpenGL error:", err)
  }
  ```

### `get_integerv(pname)`
### `get_floatv(pname)`
### `get_doublev(pname)`
### `get_booleanv(pname)`

Queries OpenGL state values.

- **Parameters**: `pname` - Parameter to query
- **Returns**: Query result (int, float, or bool)

### `get_string(name)`

Returns a string from OpenGL.

- **Parameters**: `name` - `gl.VENDOR`, `gl.RENDERER`, `gl.VERSION`, or `gl.EXTENSIONS`
- **Returns**: String
- **Example**: `println("OpenGL Version:", gl.get_string(gl.VERSION))`

---

## Flush & Finish

### `flush()`

Forces execution of OpenGL commands.

- **Returns**: `null`

### `finish()`

Waits for all OpenGL commands to complete.

- **Returns**: `null`

---

## Accumulation Buffer

### `accum(op, value)`

Performs accumulation buffer operations.

- **Parameters**: `op` - `gl.ACCUM`, `gl.LOAD`, `gl.RETURN`, `gl.MULT`, or `gl.ADD`, `value` - Operation value (float)
- **Returns**: `null`

---

## Render Mode

### `render_mode(mode)`

Sets the render mode.

- **Parameters**: `mode` - `gl.RENDER`, `gl.FEEDBACK`, or `gl.SELECT`
- **Returns**: Previous render mode (int)

---

## Constants Reference

### Matrix Modes
- `gl.MODELVIEW`
- `gl.PROJECTION`
- `gl.TEXTURE`

### Clear Bits
- `gl.COLOR_BUFFER_BIT`
- `gl.DEPTH_BUFFER_BIT`
- `gl.ACCUM_BUFFER_BIT`
- `gl.STENCIL_BUFFER_BIT`

### Primitive Types
- `gl.POINTS`
- `gl.LINES`
- `gl.LINE_LOOP`
- `gl.LINE_STRIP`
- `gl.TRIANGLES`
- `gl.TRIANGLE_STRIP`
- `gl.TRIANGLE_FAN`
- `gl.QUADS`
- `gl.QUAD_STRIP`
- `gl.POLYGON`

### Shading Models
- `gl.FLAT`
- `gl.SMOOTH`

### Capabilities
- `gl.BLEND`
- `gl.DEPTH_TEST`
- `gl.CULL_FACE`
- `gl.LIGHTING`
- `gl.LIGHT0` through `gl.LIGHT7`
- `gl.FOG`
- `gl.SCISSOR_TEST`
- `gl.STENCIL_TEST`
- `gl.ALPHA_TEST`
- `gl.NORMALIZE`
- `gl.COLOR_MATERIAL`

### Blend Factors
- `gl.ZERO`
- `gl.ONE`
- `gl.SRC_COLOR`
- `gl.ONE_MINUS_SRC_COLOR`
- `gl.DST_COLOR`
- `gl.ONE_MINUS_DST_COLOR`
- `gl.SRC_ALPHA`
- `gl.ONE_MINUS_SRC_ALPHA`
- `gl.DST_ALPHA`
- `gl.ONE_MINUS_DST_ALPHA`
- `gl.SRC_ALPHA_SATURATE`
- `gl.CONSTANT_COLOR`
- `gl.ONE_MINUS_CONSTANT_COLOR`
- `gl.CONSTANT_ALPHA`
- `gl.ONE_MINUS_CONSTANT_ALPHA`

### Blend Equations
- `gl.FUNC_ADD`
- `gl.FUNC_SUBTRACT`
- `gl.FUNC_REVERSE_SUBTRACT`
- `gl.MIN`
- `gl.MAX`

### Depth/Alpha Functions
- `gl.NEVER`
- `gl.LESS`
- `gl.EQUAL`
- `gl.LEQUAL`
- `gl.GREATER`
- `gl.NOTEQUAL`
- `gl.GEQUAL`
- `gl.ALWAYS`

### Cull Face Modes
- `gl.FRONT`
- `gl.BACK`
- `gl.FRONT_AND_BACK`

### Front Face
- `gl.CW` (clockwise)
- `gl.CCW` (counter-clockwise)

### Polygon Modes
- `gl.POINT`
- `gl.LINE`
- `gl.FILL`

### Hints
- `gl.DONT_CARE`
- `gl.FASTEST`
- `gl.NICEST`

### Fog Modes
- `gl.EXP`
- `gl.EXP2`

### Texture Targets
- `gl.TEXTURE_1D`
- `gl.TEXTURE_2D`
- `gl.TEXTURE_3D`
- `gl.TEXTURE_CUBE_MAP`
- `gl.TEXTURE_CUBE_MAP_POSITIVE_X` through `gl.TEXTURE_CUBE_MAP_NEGATIVE_Z`

### Texture Filters
- `gl.NEAREST`
- `gl.LINEAR`
- `gl.NEAREST_MIPMAP_NEAREST`
- `gl.LINEAR_MIPMAP_NEAREST`
- `gl.NEAREST_MIPMAP_LINEAR`
- `gl.LINEAR_MIPMAP_LINEAR`

### Texture Wrap Modes
- `gl.CLAMP`
- `gl.REPEAT`
- `gl.CLAMP_TO_EDGE`
- `gl.CLAMP_TO_BORDER`
- `gl.MIRRORED_REPEAT`

### Internal Formats
- `gl.ALPHA`, `gl.ALPHA4`, `gl.ALPHA8`, `gl.ALPHA12`, `gl.ALPHA16`
- `gl.LUMINANCE`, `gl.LUMINANCE4`, `gl.LUMINANCE8`, `gl.LUMINANCE12`, `gl.LUMINANCE16`
- `gl.LUMINANCE_ALPHA`, `gl.LUMINANCE4_ALPHA4`, `gl.LUMINANCE8_ALPHA8`
- `gl.RGB`, `gl.RGB4`, `gl.RGB8`, `gl.RGB10`, `gl.RGB12`, `gl.RGB16`
- `gl.RGBA`, `gl.RGBA2`, `gl.RGBA4`, `gl.RGB5_A1`, `gl.RGBA8`, `gl.RGB10_A2`, `gl.RGBA12`, `gl.RGBA16`
- `gl.DEPTH_COMPONENT`, `gl.DEPTH_COMPONENT16`, `gl.DEPTH_COMPONENT24`, `gl.DEPTH_COMPONENT32`

### Pixel Data Types
- `gl.UNSIGNED_BYTE`
- `gl.BYTE`
- `gl.UNSIGNED_SHORT`
- `gl.SHORT`
- `gl.UNSIGNED_INT`
- `gl.INT`
- `gl.FLOAT`
- `gl.DOUBLE`

### Vertex Arrays
- `gl.VERTEX_ARRAY`
- `gl.NORMAL_ARRAY`
- `gl.COLOR_ARRAY`
- `gl.TEXTURE_COORD_ARRAY`

### Error Codes
- `gl.NO_ERROR`
- `gl.INVALID_ENUM`
- `gl.INVALID_VALUE`
- `gl.INVALID_OPERATION`
- `gl.STACK_OVERFLOW`
- `gl.STACK_UNDERFLOW`
- `gl.OUT_OF_MEMORY`

### GetString Names
- `gl.VENDOR`
- `gl.RENDERER`
- `gl.VERSION`
- `gl.EXTENSIONS`

---

## Complete Example

```go
import "gl"

// Initialize OpenGL
gl.init()

// Set up the viewport
gl.viewport(0, 0, 800, 600)

// Set up projection matrix
gl.matrix_mode(gl.PROJECTION)
gl.load_identity()
gl.ortho(-1, 1, -1, 1, -1, 1)

// Set up modelview matrix
gl.matrix_mode(gl.MODELVIEW)
gl.load_identity()

// Clear the screen
gl.clear_color(0.2, 0.3, 0.4, 1.0)
gl.clear(gl.COLOR_BUFFER_BIT)

// Enable features
gl.enable(gl.BLEND)
gl.blend_func(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

// Draw a triangle
gl.begin(gl.TRIANGLES)
gl.color3f(1.0, 0.0, 0.0)
gl.vertex3f(400, 100, 0)
gl.color3f(0.0, 1.0, 0.0)
gl.vertex3f(200, 400, 0)
gl.color3f(0.0, 0.0, 1.0)
gl.vertex3f(600, 400, 0)
gl.end()

// Draw a textured quad
gl.enable(gl.TEXTURE_2D)
textures := gl.gen_textures(1)
gl.bind_texture(gl.TEXTURE_2D, textures[0])
gl.tex_parameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
gl.tex_parameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

// ... set texture data with gl.tex_image2d ...

gl.begin(gl.QUADS)
gl.tex_coord2f(0, 0)
gl.vertex3f(0, 0, 0)
gl.tex_coord2f(1, 0)
gl.vertex3f(100, 0, 0)
gl.tex_coord2f(1, 1)
gl.vertex3f(100, 100, 0)
gl.tex_coord2f(0, 1)
gl.vertex3f(0, 100, 0)
gl.end()

// Flush commands
gl.flush()
```