import "gl"
import "glut"
import "glu"
import "os"
import "math"

glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE | glut.DEPTH)
glut.init_window_size(400, 400)
glut.create_window("Display List Performance Demo")

gl.enable(gl.DEPTH_TEST)
gl.enable(gl.LIGHTING)
gl.enable(gl.LIGHT0)

gl.lightfv(gl.LIGHT0, gl.POSITION, 2.0, 5.0, 3.0, 1.0)

// Create display list for complex object
var object_list = gl.gen_lists(1)
gl.new_list(object_list, gl.COMPILE)

// Build a complex scene
for i := 0; i < 50; i++ {
    gl.push_matrix()
    gl.translatef(
        math.sin(i * 0.5) * 4,
        math.cos(i * 0.3) * 3,
        math.sin(i * 0.7) * 2
    )
    gl.rotatef(i * 15, 1.0, 0.5, 0.3)
    gl.color3f(
        math.sin(i * 0.1) * 0.5 + 0.5,
        math.cos(i * 0.2) * 0.5 + 0.5,
        math.sin(i * 0.15) * 0.5 + 0.5
    )
    glut.solid_torus(0.2, 0.5, 10, 10)
    gl.pop_matrix()
}
gl.end_list()

// Animation
var rot_angle = 0.0

glut.display_func(fn() {
    gl.clear_color(0.1, 0.1, 0.2, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
    gl.matrix_mode(gl.MODELVIEW)
    gl.load_identity()
    gl.translatef(0.0, 0.0, -12.0)
    gl.rotatef(rot_angle, 0.0, 1.0, 0.0)
    
    // Call the display list (fast!)
    gl.call_list(object_list)
    
    glut.swap_buffers()
})

glut.reshape_func(fn(w, h) {
    gl.viewport(0, 0, w, h)
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    glu.perspective(45.0, w/h, 0.1, 50.0)
})

glut.idle_func(fn() {
    rot_angle += 0.5
    glut.post_redisplay()
})

glut.keyboard_func(fn(key, x, y) {
    if key == "q" || key == "Q" { os.exit(0) }
})

glut.main_loop()