import "gl"
import "glut"
import "glu"
import "os"
import "math"

// Global rotation variables
var rot_x = 0.0
var rot_y = 0.0
var rot_z = 0.0

glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE | glut.DEPTH)
glut.init_window_size(400, 400)
glut.init_window_position(100, 100)

win := glut.create_window("3D Rotating Cube")
println("Window created:", win)

// Enable depth testing for 3D rendering
gl.enable(gl.DEPTH_TEST)

glut.display_func(fn() {
    gl.clear_color(0.1, 0.1, 0.2, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
    gl.matrix_mode(gl.MODELVIEW)
    gl.load_identity()
    gl.translatef(0.0, 0.0, -5.0)
    
    // Apply rotations
    gl.rotatef(rot_x, 1.0, 0.0, 0.0)
    gl.rotatef(rot_y, 0.0, 1.0, 0.0)
    gl.rotatef(rot_z, 0.0, 0.0, 1.0)
    
    // Draw colorful cube
    gl.begin(gl.QUADS)
    
    // Front face (red)
    gl.color3f(1.0, 0.0, 0.0)
    gl.vertex3f(-1.0, -1.0,  1.0)
    gl.vertex3f( 1.0, -1.0,  1.0)
    gl.vertex3f( 1.0,  1.0,  1.0)
    gl.vertex3f(-1.0,  1.0,  1.0)
    
    // Back face (blue)
    gl.color3f(0.0, 0.0, 1.0)
    gl.vertex3f(-1.0, -1.0, -1.0)
    gl.vertex3f(-1.0,  1.0, -1.0)
    gl.vertex3f( 1.0,  1.0, -1.0)
    gl.vertex3f( 1.0, -1.0, -1.0)
    
    // Top face (green)
    gl.color3f(0.0, 1.0, 0.0)
    gl.vertex3f(-1.0,  1.0, -1.0)
    gl.vertex3f(-1.0,  1.0,  1.0)
    gl.vertex3f( 1.0,  1.0,  1.0)
    gl.vertex3f( 1.0,  1.0, -1.0)
    
    // Bottom face (yellow)
    gl.color3f(1.0, 1.0, 0.0)
    gl.vertex3f(-1.0, -1.0, -1.0)
    gl.vertex3f( 1.0, -1.0, -1.0)
    gl.vertex3f( 1.0, -1.0,  1.0)
    gl.vertex3f(-1.0, -1.0,  1.0)
    
    // Right face (magenta)
    gl.color3f(1.0, 0.0, 1.0)
    gl.vertex3f( 1.0, -1.0, -1.0)
    gl.vertex3f( 1.0,  1.0, -1.0)
    gl.vertex3f( 1.0,  1.0,  1.0)
    gl.vertex3f( 1.0, -1.0,  1.0)
    
    // Left face (cyan)
    gl.color3f(0.0, 1.0, 1.0)
    gl.vertex3f(-1.0, -1.0, -1.0)
    gl.vertex3f(-1.0, -1.0,  1.0)
    gl.vertex3f(-1.0,  1.0,  1.0)
    gl.vertex3f(-1.0,  1.0, -1.0)
    
    gl.end()
    
    glut.swap_buffers()
})

glut.reshape_func(fn(w, h) {
    gl.viewport(0, 0, w, h)
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    glu.perspective(45.0, w/h, 0.1, 100.0)
})

glut.idle_func(fn() {
    rot_x += 0.5
    rot_y += 1.0
    rot_z += 0.3
    glut.post_redisplay()
})

glut.keyboard_func(fn(key, x, y) {
    if key == "q" || key == "Q" {
        os.exit(0)
    }
    if key == "r" {
        rot_x = 0
        rot_y = 0
        rot_z = 0
    }
})

glut.main_loop()