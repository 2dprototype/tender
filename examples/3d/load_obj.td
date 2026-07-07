import "gl"
import "glut"
import "glu"
import "os"

glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE | glut.DEPTH)
glut.init_window_size(400, 400)
glut.init_window_position(100, 100)

win := glut.create_window("OBJ Loader")

gl.enable(gl.DEPTH_TEST)
gl.enable(gl.LIGHTING)
gl.enable(gl.LIGHT0)
gl.enable(gl.NORMALIZE)
gl.enable(gl.COLOR_MATERIAL)

gl.lightfv(gl.LIGHT0, gl.POSITION, 5.0, 5.0, 5.0, 1.0)
gl.lightfv(gl.LIGHT0, gl.DIFFUSE, 0.9, 0.9, 0.9, 1.0)
gl.lightfv(gl.LIGHT0, gl.SPECULAR, 1.0, 1.0, 1.0, 1.0)

mesh_id := gl.load_obj("monkey.obj")

var rot = 0.0

glut.display_func(fn() {
    gl.clear_color(0.1, 0.1, 0.15, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    gl.matrix_mode(gl.MODELVIEW)
    gl.load_identity()
    gl.translatef(0.0, 0.0, -5.0)
    gl.rotatef(rot, 0.0, 1.0, 0.0)
    
    gl.draw_obj(mesh_id)
    
    glut.swap_buffers()
})

glut.reshape_func(fn(w, h) {
    gl.viewport(0, 0, w, h)
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    glu.perspective(45.0, w/h, 0.1, 100.0)
})

glut.idle_func(fn() {
    rot += 0.8
    glut.post_redisplay()
})

glut.keyboard_func(fn(key, x, y) {
    if key == "q" || key == "Q" {
        os.exit(0)
    }
})

glut.main_loop()