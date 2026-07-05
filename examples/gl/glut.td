import "gl"
import "glut"
import "os"

// Initialize GLUT
glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE | glut.DEPTH)
glut.init_window_size(400, 400)
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