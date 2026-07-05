import "gl"
import "glut"
import "os"
import "math"

// Global state
var cam_rot_x = 0.0
var cam_rot_y = 0.0
var cam_distance = 8.0
var mouse_down = false
var last_mouse_x = 0
var last_mouse_y = 0
var selected_object = 0
var object_colors = []

fn set_perspective(fov, aspect, near, far) {
    f := 1.0 / math.tan(fov * math.pi / 360.0)
    
    gl.load_identity()
    gl.mult_matrixf(
        f / aspect, 0.0, 0.0, 0.0,
        0.0, f, 0.0, 0.0,
        0.0, 0.0, (far + near) / (near - far), -1.0,
        0.0, 0.0, (2.0 * far * near) / (near - far), 0.0
    )
}

glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE | glut.DEPTH)
glut.init_window_size(400, 400)
glut.init_window_position(100, 100)

win := glut.create_window("Interactive 3D Demo")

gl.enable(gl.DEPTH_TEST)
gl.enable(gl.LIGHTING)
gl.enable(gl.LIGHT0)
gl.enable(gl.NORMALIZE)

// Initialize object colors
object_colors = [
    [1.0, 0.0, 0.0],
    [0.0, 1.0, 0.0],
    [0.0, 0.0, 1.0],
    [1.0, 1.0, 0.0],
    [1.0, 0.0, 1.0],
    [0.0, 1.0, 1.0]
]

// Compile display lists for objects
// gen_lists returns the starting list ID
var list_start = gl.gen_lists(6)

// List 0: Sphere (list_start + 0)
gl.new_list(list_start, gl.COMPILE)
glut.solid_sphere(0.6, 24, 16)
gl.end_list()

// List 1: Cube (list_start + 1)
gl.new_list(list_start + 1, gl.COMPILE)
glut.solid_cube(1.0)
gl.end_list()

// List 2: Cone (list_start + 2)
gl.new_list(list_start + 2, gl.COMPILE)
glut.solid_cone(0.7, 1.2, 20, 10)
gl.end_list()

// List 3: Torus (list_start + 3)
gl.new_list(list_start + 3, gl.COMPILE)
glut.solid_torus(0.4, 0.7, 16, 16)
gl.end_list()

// List 4: Teapot (list_start + 4)
gl.new_list(list_start + 4, gl.COMPILE)
glut.solid_teapot(0.6)
gl.end_list()

// List 5: Dodecahedron (list_start + 5)
gl.new_list(list_start + 5, gl.COMPILE)
glut.solid_dodecahedron()
gl.end_list()

// Object positions in a circle
var object_angles = []
for i in range(0, 6) {
    object_angles.push(i * 60.0)
}

glut.display_func(fn() {
    gl.clear_color(0.1, 0.1, 0.2, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
    gl.matrix_mode(gl.MODELVIEW)
    gl.load_identity()
    
    // Camera orbit
    gl.translatef(0.0, 0.0, -cam_distance)
    gl.rotatef(cam_rot_x, 1.0, 0.0, 0.0)
    gl.rotatef(cam_rot_y, 0.0, 1.0, 0.0)
    
    // Draw grid (floor)
    gl.color3f(0.3, 0.3, 0.3)
    gl.begin(gl.LINES)
    for i in range(-5, 5) {
        gl.vertex3f(-5.0, -1.5, i)
        gl.vertex3f(5.0, -1.5, i)
        gl.vertex3f(i, -1.5, -5.0)
        gl.vertex3f(i, -1.5, 5.0)
    }
    gl.end()
    
    // Draw objects in a circle
    for i in range(0, 6) {
        angle := object_angles[i] * math.pi / 180.0
        x := 2.5 * math.cos(angle)
        z := 2.5 * math.sin(angle)
        
        gl.push_matrix()
        gl.translatef(x, 0.0, z)
        
        // Highlight selected object
        if i == selected_object {
            gl.color3f(1.0, 1.0, 1.0)
            glut.wire_sphere(0.8, 16, 12)
        }
        
        gl.color3f(object_colors[i][0], object_colors[i][1], object_colors[i][2])
        // Call the display list: list_start + i
        gl.call_list(list_start + i)
        gl.pop_matrix()
    }
    
    glut.swap_buffers()
})

glut.reshape_func(fn(w, h) {
    gl.viewport(0, 0, w, h)
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    set_perspective(45.0, w/h, 0.1, 100.0)
})

// Mouse interaction
glut.mouse_func(fn(button, state, x, y) {
    if button == glut.LEFT_BUTTON {
        if state == glut.DOWN {
            mouse_down = true
            last_mouse_x = x
            last_mouse_y = y
        } else {
            mouse_down = false
        }
    }
    if button == glut.RIGHT_BUTTON && state == glut.DOWN {
        // Cycle selection
        selected_object = (selected_object + 1) % 6
        println("Selected object:", selected_object)
    }
})

glut.motion_func(fn(x, y) {
    if mouse_down {
        dx := x - last_mouse_x
        dy := y - last_mouse_y
        cam_rot_y += dx * 0.3
        cam_rot_x += dy * 0.3
        
        // Clamp vertical rotation
        if cam_rot_x > 90.0 {
            cam_rot_x = 90.0
        } else if cam_rot_x < -90.0 {
            cam_rot_x = -90.0
        }
        
        last_mouse_x = x
        last_mouse_y = y
        glut.post_redisplay()
    }
})

glut.keyboard_func(fn(key, x, y) {
    if key == "q" || key == "Q" {
        os.exit(0)
    }
})

glut.special_func(fn(key, x, y) {
    if key == glut.KEY_UP {
        cam_distance -= 0.5
    } else if key == glut.KEY_DOWN {
        cam_distance += 0.5
    }
    // Clamp distance
    if cam_distance < 2.0 {
        cam_distance = 2.0
    } else if cam_distance > 20.0 {
        cam_distance = 20.0
    }
})

glut.main_loop()