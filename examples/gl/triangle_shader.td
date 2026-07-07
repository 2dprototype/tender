import "gl"
import "glut"
import "os"
import "math"

// Initialize Window
glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE)
glut.init_window_size(400, 400)
glut.create_window("Modern OpenGL - Animated Triangle")

// Vertex Shader
var vertex_src = `
#version 120
attribute vec2 position;
attribute vec3 color;
varying vec3 vColor;
void main() {
    vColor = color;
    gl_Position = vec4(position, 0.0, 1.0);
}
`

// Fragment Shader with time_-based color
var frag_src = `
#version 120
uniform float u_time_;
varying vec3 vColor;
void main() {
    float pulse = sin(u_time_) * 0.5 + 0.5;
    gl_FragColor = vec4(vColor * pulse, 1.0);
}
`

// Compile Shaders
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

// Get uniforms
var time__loc = gl.get_uniform_location(program, "u_time_")

// Triangle with colors: position (x,y), color (r,g,b)
var triangle_data = [
     0.0,  0.5,  1.0, 0.0, 0.0,  // Top - Red
    -0.5, -0.5,  0.0, 1.0, 0.0,  // Bottom Left - Green
     0.5, -0.5,  0.0, 0.0, 1.0   // Bottom Right - Blue
]

// Setup VAO/VBO
var vaos = gl.gen_vertex_arrays(1)
var vao = vaos[0]
gl.bind_vertex_array(vao)

var vbos = gl.gen_buffers(1)
var vbo = vbos[0]
gl.bind_buffer(gl.ARRAY_BUFFER, vbo)
gl.buffer_data(gl.ARRAY_BUFFER, triangle_data, gl.STATIC_DRAW)

// Position attribute (2 floats)
gl.enable_vertex_attrib_array(0)
gl.vertex_attrib_pointer(0, 2, gl.FLOAT, false, 20, 0)  // 5 floats * 4 bytes = 20

// Color attribute (3 floats)
gl.enable_vertex_attrib_array(1)
gl.vertex_attrib_pointer(1, 3, gl.FLOAT, false, 20, 8)   // 8 bytes offset for color

gl.bind_vertex_array(0)

// Animation loop
var time_ = 0.0

glut.display_func(fn() {
    gl.clear_color(0.1, 0.1, 0.2, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT)
    
    gl.use_program(program)
    
    time_ += 0.02
    gl.uniform1f(time__loc, time_)
    
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