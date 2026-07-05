import "gl"
import "glut"

glut.init()
gl.init()
glut.init_display_mode(glut.RGB | glut.SINGLE) // GLUT_RGB | GLUT_SINGLE
glut.init_window_size(400, 400)
glut.create_window("Tender Triangle Test")

glut.display_func(fn(){
    gl.clear_color(0.0, 0.0, 0.0, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT) // GL_COLOR_BUFFER_BIT = 16384
    gl.load_identity()
    gl.begin(gl.TRIANGLES) // GL_TRIANGLES = 4
    gl.color3f(1.0, 0.0, 0.0)
    gl.vertex3f(0.0, 0.75, 0.0)
    gl.color3f(0.0, 1.0, 0.0)
    gl.vertex3f(-0.75, -0.75, 0.0)
    gl.color3f(0.0, 0.0, 1.0)
    gl.vertex3f(0.75, -0.75, 0.0)
    gl.end()
    gl.flush()
})

glut.keyboard_func(fn(a, b, c){
    println(a, b, c)
})

glut.reshape_func(fn(w, h) {
    // (w, h) |> debug
    gl.viewport(0, 0, w, h)
})
    
glut.main_loop()