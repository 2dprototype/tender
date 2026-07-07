import "gl"
import "glut"
import "glu"
import "os"

glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE | glut.DEPTH)
glut.init_window_size(440, 440)
glut.init_window_position(100, 50)
glut.create_window("Interactive Textured Engine")

gl.enable(gl.DEPTH_TEST)
gl.enable(gl.LIGHTING)
gl.enable(gl.LIGHT0)
gl.enable(gl.TEXTURE_2D) // IMPORTANT: Globally enable 2D textures!

// Set up lighting properties
gl.lightfv(gl.LIGHT0, gl.POSITION, 3.0, 4.0, 5.0, 1.0)
gl.lightfv(gl.LIGHT0, gl.DIFFUSE, 1.0, 1.0, 1.0, 1.0)

tex_id := gl.load_texture("textures.jpg")
mesh_id := gl.parse_obj(embed("model.obj"))

var rot_x = 0.0
var rot_y = 0.0
var pos_x = -0.5 // Offset slightly to center the 0..1 cube
var pos_y = -3
var pos_z = -15.0 // Initial zoom

var is_dragging = false
var last_mouse_x = 0.0
var last_mouse_y = 0.0

glut.display_func(fn() {
    gl.clear_color(0.05, 0.05, 0.08, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
    gl.matrix_mode(gl.MODELVIEW)
    gl.load_identity()
    
    gl.translatef(pos_x, pos_y, pos_z)

    gl.rotatef(rot_x, 1.0, 0.0, 0.0)
    gl.rotatef(rot_y, 0.0, 1.0, 0.0)
    
    // Bind texture to pipeline before processing the mesh display list
    gl.bind_texture(gl.TEXTURE_2D, tex_id)
    
    // Draw the hardware-cached geometry loop
    gl.draw_obj(mesh_id)
    
    glut.swap_buffers()
})

glut.reshape_func(fn(w, h) {
    gl.viewport(0, 0, w, h)
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    glu.perspective(45.0, w/h, 0.1, 100.0)
})


glut.mouse_func(fn(button, state, x, y) {
    if button == 0 {
        if state == 0 {
            is_dragging = true
            last_mouse_x = x
            last_mouse_y = y
        } 
        else {  
            is_dragging = false
        }
    }
    
    // Scroll wheel to zoom (Buttons 3 and 4 in standard GLUT)
    if button == 3 && state == 0 { pos_z += 0.5; glut.post_redisplay() }
    if button == 4 && state == 0 { pos_z -= 0.5; glut.post_redisplay() }
})

glut.motion_func(fn(x, y) {
    if is_dragging {
        var dx = x - last_mouse_x
        var dy = y - last_mouse_y
        
        // Update rotations based on mouse delta
        rot_y += dx * 0.5
        rot_x += dy * 0.5
        
        last_mouse_x = x
        last_mouse_y = y
        
        glut.post_redisplay()
    }
})

glut.keyboard_func(fn(key, x, y) {
    // Movement
    if key == "w" || key == "W" { pos_y += 0.2 }
    if key == "s" || key == "S" { pos_y -= 0.2 }
    if key == "a" || key == "A" { pos_x -= 0.2 }
    if key == "d" || key == "D" { pos_x += 0.2 }
    
    // Zoom keys (alternative to scroll wheel)
    if key == "+" || key == "=" { pos_z += 0.5 }
    if key == "-" || key == "_" { pos_z -= 0.5 }
    
    // Exit
    if key == "q" || key == "Q" { os.exit(0) }
    
    glut.post_redisplay()
})

glut.main_loop()