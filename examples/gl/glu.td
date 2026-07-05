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