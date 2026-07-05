import "gl"
import "glut"
import "glu"
import "os"
import "math"

glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE | glut.DEPTH)
glut.init_window_size(400, 400)
glut.init_window_position(100, 100)

win := glut.create_window("Vertex Arrays Demo")

gl.enable(gl.DEPTH_TEST)
gl.enable(gl.LIGHTING)
gl.enable(gl.LIGHT0)
gl.enable(gl.NORMALIZE)

// Define vertex data for a colored cube
var vertices = [
    // Front face
    -1.0, -1.0,  1.0,
     1.0, -1.0,  1.0,
     1.0,  1.0,  1.0,
    -1.0,  1.0,  1.0,
    // Back face
    -1.0, -1.0, -1.0,
    -1.0,  1.0, -1.0,
     1.0,  1.0, -1.0,
     1.0, -1.0, -1.0,
    // Top face
    -1.0,  1.0, -1.0,
    -1.0,  1.0,  1.0,
     1.0,  1.0,  1.0,
     1.0,  1.0, -1.0,
    // Bottom face
    -1.0, -1.0, -1.0,
     1.0, -1.0, -1.0,
     1.0, -1.0,  1.0,
    -1.0, -1.0,  1.0,
    // Right face
     1.0, -1.0, -1.0,
     1.0,  1.0, -1.0,
     1.0,  1.0,  1.0,
     1.0, -1.0,  1.0,
    // Left face
    -1.0, -1.0, -1.0,
    -1.0, -1.0,  1.0,
    -1.0,  1.0,  1.0,
    -1.0,  1.0, -1.0
]

var colors = [
    // Front face - Red
    1.0, 0.0, 0.0,
    1.0, 0.0, 0.0,
    1.0, 0.0, 0.0,
    1.0, 0.0, 0.0,
    // Back face - Blue
    0.0, 0.0, 1.0,
    0.0, 0.0, 1.0,
    0.0, 0.0, 1.0,
    0.0, 0.0, 1.0,
    // Top face - Green
    0.0, 1.0, 0.0,
    0.0, 1.0, 0.0,
    0.0, 1.0, 0.0,
    0.0, 1.0, 0.0,
    // Bottom face - Yellow
    1.0, 1.0, 0.0,
    1.0, 1.0, 0.0,
    1.0, 1.0, 0.0,
    1.0, 1.0, 0.0,
    // Right face - Magenta
    1.0, 0.0, 1.0,
    1.0, 0.0, 1.0,
    1.0, 0.0, 1.0,
    1.0, 0.0, 1.0,
    // Left face - Cyan
    0.0, 1.0, 1.0,
    0.0, 1.0, 1.0,
    0.0, 1.0, 1.0,
    0.0, 1.0, 1.0
]

// Normal data for each face
var normals = [
    0.0, 0.0, 1.0, // Front
    0.0, 0.0, -1.0, // Back
    0.0, 1.0, 0.0, // Top
    0.0, -1.0, 0.0, // Bottom
    1.0, 0.0, 0.0, // Right
    -1.0, 0.0, 0.0 // Left
]

var rot = 0.0

glut.display_func(fn() {
    gl.clear_color(0.1, 0.1, 0.2, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
    gl.matrix_mode(gl.MODELVIEW)
    gl.load_identity()
    gl.translatef(0.0, 0.0, -5.0)
    gl.rotatef(rot, 1.0, 0.0, 0.0)
    gl.rotatef(rot * 0.7, 0.0, 1.0, 0.0)
    
    // Enable vertex array states
    gl.enable_client_state(gl.VERTEX_ARRAY)
    gl.enable_client_state(gl.COLOR_ARRAY)
    gl.enable_client_state(gl.NORMAL_ARRAY)
    
    // Set up pointers
    gl.vertex_pointer(3, gl.FLOAT, 0, vertices)
    gl.color_pointer(3, gl.FLOAT, 0, colors)
    gl.normal_pointer(gl.FLOAT, 0, normals)
    
    // Draw each face separately with different normals
    for face in range(0, 6) {
        // Set normal for this face
        gl.normal3f(normals[face*3], normals[face*3+1], normals[face*3+2])
        gl.draw_arrays(gl.QUADS, face*4, 4)
    }
    
    // Disable vertex array states
    gl.disable_client_state(gl.VERTEX_ARRAY)
    gl.disable_client_state(gl.COLOR_ARRAY)
    gl.disable_client_state(gl.NORMAL_ARRAY)
    
    glut.swap_buffers()
})

glut.reshape_func(fn(w, h) {
    gl.viewport(0, 0, w, h)
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    glu.perspective(45.0, w/h, 0.1, 100.0)
})

glut.idle_func(fn() {
    rot += 0.5
    glut.post_redisplay()
})

glut.keyboard_func(fn(key, x, y) {
    if key == "q" || key == "Q" {
        os.exit(0)
    }
})

glut.main_loop()