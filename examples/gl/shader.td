// tender@v0.0.7

import "gl"
import "glut"
import "os"
import "math"

// Initialize Window Context
glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE)
glut.init_window_size(600, 600)
glut.create_window("Tender - Basic Shader Example")

// Define Minimal Shaders
var vertex_src = `
#version 120
attribute vec2 position;
void main() {
    gl_Position = vec4(position, 0.0, 1.0);
}
`

var frag_src = `
#version 120
uniform float u_blue;
void main() {
    // Red = 1.0, Green = 0.5, Blue animates over time
    gl_FragColor = vec4(1.0, 0.5, u_blue, 1.0);
}
`

// Compile and Link Shader Program
var vs = gl.create_shader(gl.VERTEX_SHADER)
gl.shader_source(vs, vertex_src)
gl.compile_shader(vs)
if gl.get_shader_iv(vs, gl.COMPILE_STATUS) == 0 {
    print("Vertex Shader Error: " + gl.get_shader_info_log(vs))
}

var fs = gl.create_shader(gl.FRAGMENT_SHADER)
gl.shader_source(fs, frag_src)
gl.compile_shader(fs)
if gl.get_shader_iv(fs, gl.COMPILE_STATUS) == 0 {
    print("Fragment Shader Error: " + gl.get_shader_info_log(fs))
}

var program = gl.create_program()
gl.attach_shader(program, vs)
gl.attach_shader(program, fs)
gl.link_program(program)

// Get uniform location for our color animator
var blue_loc = gl.get_uniform_location(program, "u_blue")

// Setup Simple Triangle Geometry (X, Y coordinates)
var triangle_vertices = [
     0.0,  0.5,  // Top
    -0.5, -0.5,  // Bottom Left
     0.5, -0.5   // Bottom Right
]

var vaos = gl.gen_vertex_arrays(1)
var vao = vaos[0]

var vbos = gl.gen_buffers(1)
var vbo = vbos[0]

gl.bind_vertex_array(vao)
gl.bind_buffer(gl.ARRAY_BUFFER, vbo)

// Upload flat float values using your updated Go marshaling function
gl.buffer_data(gl.ARRAY_BUFFER, triangle_vertices, gl.STATIC_DRAW)

// Map attribute index 0 to 'position' in the vertex shader
gl.enable_vertex_attrib_array(0)
gl.vertex_attrib_pointer(0, 2, gl.FLOAT, false, 8, 0) // 2 floats * 4 bytes = 8 stride

gl.bind_vertex_array(0)

// Render Loop Animation
var time_ = 0.0

glut.display_func(fn() {
    gl.clear_color(0.1, 0.1, 0.1, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT)
    
    // Bind shader program
    gl.use_program(program)
    
    // Pass time-based values to the uniform (oscillate blue component between 0.0 and 1.0)
    time_ += 0.03
    var blue_value = (math.sin(time_) + 1.0) * 0.5
    gl.uniform1f(blue_loc, blue_value)
    
    // Bind geometry and draw
    gl.bind_vertex_array(vao)
    gl.draw_arrays(gl.TRIANGLES, 0, 3)
    gl.bind_vertex_array(0)
    
    glut.swap_buffers()
})

glut.idle_func(fn() {
    glut.post_redisplay()
})

glut.keyboard_func(fn(key, x, y) {
    if key == "q" || key == "Q" { os.exit(0) }
})

glut.main_loop()