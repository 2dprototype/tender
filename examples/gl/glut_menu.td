import "gl"
import "glut"
import "os"

glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE)
glut.init_window_size(400, 400)
glut.create_window("GLUT Menu Demo")

// Create submenu
var submenu = glut.create_menu(fn(value) {
    print("Submenu selected: " + string(value))
    if value == 10 { gl.clear_color(1.0, 0.0, 0.0, 1.0) }
    if value == 11 { gl.clear_color(0.0, 1.0, 0.0, 1.0) }
    if value == 12 { gl.clear_color(0.0, 0.0, 1.0, 1.0) }
})

glut.add_menu_entry("Red", 10)
glut.add_menu_entry("Green", 11)
glut.add_menu_entry("Blue", 12)

// Create main menu
var main_menu = glut.create_menu(fn(value) {
    print("Menu selected: " + string(value))
    if value == 1 { os.exit(0) }
    if value == 2 { gl.clear_color(0.2, 0.2, 0.2, 1.0) }
})

glut.add_menu_entry("Exit", 1)
glut.add_menu_entry("Reset Background", 2)
glut.add_sub_menu("Change Color", submenu)

// Attach menu to right mouse button
glut.attach_menu(main_menu)

glut.display_func(fn() {
    gl.clear(gl.COLOR_BUFFER_BIT)
    glut.swap_buffers()
})

glut.idle_func(fn() {
    gl.clear(gl.COLOR_BUFFER_BIT)
    glut.swap_buffers()
})

glut.keyboard_func(fn(key, x, y) {
    if key == "q" || key == "Q" { os.exit(0) }
})

glut.main_loop()