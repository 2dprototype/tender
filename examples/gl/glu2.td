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