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

win := glut.create_window("3D Lighting Demo")

gl.enable(gl.DEPTH_TEST)
gl.enable(gl.LIGHTING)
gl.enable(gl.LIGHT0)
gl.enable(gl.LIGHT1)
gl.enable(gl.NORMALIZE)
gl.enable(gl.COLOR_MATERIAL)

// Setup light 0 - White light from above-right
gl.lightfv(gl.LIGHT0, gl.POSITION, 5.0, 5.0, 5.0, 1.0)
gl.lightfv(gl.LIGHT0, gl.DIFFUSE, 1.0, 1.0, 1.0, 1.0)
gl.lightfv(gl.LIGHT0, gl.SPECULAR, 1.0, 1.0, 1.0, 1.0)

// Setup light 1 - Colored light from below-left
gl.lightfv(gl.LIGHT1, gl.POSITION, -5.0, -5.0, 5.0, 1.0)
gl.lightfv(gl.LIGHT1, gl.DIFFUSE, 0.3, 0.3, 1.0, 1.0)
gl.lightfv(gl.LIGHT1, gl.SPECULAR, 0.3, 0.3, 1.0, 1.0)

// Material properties
gl.materialfv(gl.FRONT_AND_BACK, gl.SPECULAR, 0.8, 0.8, 0.8, 1.0)
gl.materialf(gl.FRONT_AND_BACK, gl.SHININESS, 32.0)

var rot = 0.0

glut.display_func(fn() {
    gl.clear_color(0.1, 0.1, 0.2, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
    gl.matrix_mode(gl.MODELVIEW)
    gl.load_identity()
    gl.translatef(0.0, 0.0, -8.0)
    gl.rotatef(rot, 0.0, 1.0, 0.0)
    
    // Draw a torus with lighting
    glut.solid_torus(0.8, 2.0, 32, 32)
    
    // Draw a sphere below
    gl.push_matrix()
    gl.translatef(0.0, -2.5, 0.0)
    gl.scalef(0.5, 0.5, 0.5)
    glut.solid_sphere(1.0, 32, 32)
    gl.pop_matrix()
    
    // Draw a cone above
    gl.push_matrix()
    gl.translatef(0.0, 2.5, 0.0)
    gl.scalef(0.6, 0.6, 0.6)
    glut.solid_cone(1.2, 2.0, 32, 16)
    gl.pop_matrix()
    
    // Draw a teapot on the side
    gl.push_matrix()
    gl.translatef(3.0, 0.0, 0.0)
    gl.scalef(0.5, 0.5, 0.5)
    glut.solid_teapot(1.0)
    gl.pop_matrix()
    
    // Draw a dodecahedron on the other side
    gl.push_matrix()
    gl.translatef(-3.0, 0.0, 0.0)
    gl.scalef(0.8, 0.8, 0.8)
    glut.solid_dodecahedron()
    gl.pop_matrix()
    
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