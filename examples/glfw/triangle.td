import "gl"
import "glfw"

glfw.init()
gl.init()

// Set OpenGL version hints (optional, for modern OpenGL)
// For compatibility with fixed-function pipeline, use 2.1 or lower
glfw.window_hint(glfw.CONTEXT_VERSION_MAJOR, 2)
glfw.window_hint(glfw.CONTEXT_VERSION_MINOR, 1)

window := glfw.create_window(400, 400, "Tender Triangle Test (GLFW)")

glfw.make_context_current(window)

wh := glfw.get_window_size(window)
width := wh[0]
height := wh[1]
gl.viewport(0, 0, width, height)

gl.matrix_mode(gl.PROJECTION)
gl.load_identity()
gl.ortho(-1.0, 1.0, -1.0, 1.0, -1.0, 1.0)
gl.matrix_mode(gl.MODELVIEW)
gl.load_identity()

glfw.set_key_callback(window, fn(key, scancode, action, mods){
    println("Key: " + key + ", scancode: " + scancode + ", action: " + action + ", mods: " + mods)
    if key == glfw.KEY_ESCAPE && action == glfw.PRESS {
        glfw.set_window_should_close(window, 1)
    }
})

glfw.set_window_size_callback(window, fn(w, h){
    gl.viewport(0, 0, w, h)
    println("Window resized to: " + w + "x" + h)
})

glfw.set_mouse_button_callback(window, fn(button, action, mods){
    if action == glfw.PRESS {
        println("Mouse button " + button + " pressed")
    }
})

for !glfw.window_should_close(window) {
    gl.clear_color(0.0, 0.0, 0.0, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT)
    gl.load_identity()
    gl.begin(gl.TRIANGLES)
    gl.color3f(1.0, 0.0, 0.0) 
    gl.vertex3f(0.0, 0.75, 0.0)
    gl.color3f(0.0, 1.0, 0.0)
    gl.vertex3f(-0.75, -0.75, 0.0)
    gl.color3f(0.0, 0.0, 1.0) 
    gl.vertex3f(0.75, -0.75, 0.0)
    gl.end()
    gl.flush()
    glfw.swap_buffers(window)
    glfw.poll_events()
}

glfw.destroy_window(window)
glfw.terminate()