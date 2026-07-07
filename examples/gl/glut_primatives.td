import "gl"
import "glut"
import "glu"
import "os"

glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE | glut.DEPTH)
glut.init_window_size(450, 450)
glut.create_window("GLUT Primitives Demo")

gl.enable(gl.DEPTH_TEST)
gl.enable(gl.LIGHTING)
gl.enable(gl.LIGHT0)

gl.lightfv(gl.LIGHT0, gl.POSITION, 2.0, 3.0, 4.0, 1.0)
gl.lightfv(gl.LIGHT0, gl.DIFFUSE, 1.0, 1.0, 1.0, 1.0)

// Create a display list for the scene
var scene_list = gl.gen_lists(1)
gl.new_list(scene_list, gl.COMPILE)

// Draw various primitives
gl.color3f(1.0, 0.2, 0.2)
gl.push_matrix()
gl.translatef(-3.0, 2.0, 0.0)
glut.solid_sphere(0.8, 20, 20)
gl.pop_matrix()

gl.color3f(0.2, 1.0, 0.2)
gl.push_matrix()
gl.translatef(0.0, 2.0, 0.0)
glut.solid_cube(1.5)
gl.pop_matrix()

gl.color3f(0.2, 0.2, 1.0)
gl.push_matrix()
gl.translatef(3.0, 2.0, 0.0)
glut.solid_torus(0.3, 0.7, 20, 20)
gl.pop_matrix()

gl.color3f(1.0, 1.0, 0.2)
gl.push_matrix()
gl.translatef(-3.0, -1.0, 0.0)
glut.solid_cone(1.0, 1.5, 20, 20)
gl.pop_matrix()

gl.color3f(1.0, 0.5, 0.0)
gl.push_matrix()
gl.translatef(0.0, -1.0, 0.0)
glut.solid_dodecahedron()
gl.pop_matrix()

gl.color3f(0.8, 0.2, 0.8)
gl.push_matrix()
gl.translatef(3.0, -1.0, 0.0)
glut.solid_teapot(1.0)
gl.pop_matrix()

gl.end_list()

// Camera controls
var cam_rot_x = 0.0
var cam_rot_y = 0.0
var cam_dist = 10.0

glut.display_func(fn() {
    gl.clear_color(0.1, 0.1, 0.15, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
    gl.matrix_mode(gl.MODELVIEW)
    gl.load_identity()
    gl.translatef(0.0, 0.0, -cam_dist)
    gl.rotatef(cam_rot_x, 1.0, 0.0, 0.0)
    gl.rotatef(cam_rot_y, 0.0, 1.0, 0.0)
    
    gl.call_list(scene_list)
    
    glut.swap_buffers()
})

glut.reshape_func(fn(w, h) {
    gl.viewport(0, 0, w, h)
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    glu.perspective(45.0, w/h, 0.1, 50.0)
})

// Mouse interaction
var is_dragging = false
var last_x = 0
var last_y = 0

glut.mouse_func(fn(button, state, x, y) {
    if button == 0 {
        if state == 0 {
            is_dragging = true
            last_x = x
            last_y = y
        } else {
            is_dragging = false
        }
    }
    if button == 3 { cam_dist -= 0.5 }
    if button == 4 { cam_dist += 0.5 }
})

glut.motion_func(fn(x, y) {
    if is_dragging {
        cam_rot_y += (x - last_x) * 0.3
        cam_rot_x += (y - last_y) * 0.3
        last_x = x
        last_y = y
        glut.post_redisplay()
    }
})

glut.keyboard_func(fn(key, x, y) {
    if key == "q" || key == "Q" { os.exit(0) }
    if key == "r" || key == "R" {
        cam_rot_x = 0.0
        cam_rot_y = 0.0
        cam_dist = 10.0
    }
})

glut.main_loop()