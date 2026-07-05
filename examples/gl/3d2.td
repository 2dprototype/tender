import "gl"
import "glut"
import "glu"
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
var quadric = glu.new_quadric()

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
glut.init_window_size(450, 450)
glut.init_window_position(100, 100)

win := glut.create_window("Interactive 3D Demo with GLU Shapes")

gl.enable(gl.DEPTH_TEST)
gl.enable(gl.LIGHTING)
gl.enable(gl.LIGHT0)
gl.enable(gl.LIGHT1)
gl.enable(gl.NORMALIZE)

// Setup lights
gl.lightfv(gl.LIGHT0, gl.POSITION, 5.0, 5.0, 5.0, 1.0)
gl.lightfv(gl.LIGHT0, gl.DIFFUSE, 1.0, 1.0, 1.0, 1.0)

gl.lightfv(gl.LIGHT1, gl.POSITION, -5.0, -5.0, 5.0, 1.0)
gl.lightfv(gl.LIGHT1, gl.DIFFUSE, 0.3, 0.3, 1.0, 1.0)

// Initialize object colors
object_colors = [
    [1.0, 0.0, 0.0],
    [0.0, 1.0, 0.0],
    [0.0, 0.0, 1.0],
    [1.0, 1.0, 0.0],
    [1.0, 0.0, 1.0],
    [0.0, 1.0, 1.0],
    [1.0, 0.5, 0.0],
    [0.5, 0.0, 1.0],
    [0.0, 1.0, 0.5],
    [1.0, 0.5, 0.5],
    [0.5, 1.0, 0.5],
    [0.5, 0.5, 1.0]
]

// Compile display lists for objects
// gen_lists returns the starting list ID
var list_start = gl.gen_lists(12)

// List 0: GLUT Solid Sphere
gl.new_list(list_start, gl.COMPILE)
glut.solid_sphere(0.6, 24, 16)
gl.end_list()

// List 1: GLUT Solid Cube
gl.new_list(list_start + 1, gl.COMPILE)
glut.solid_cube(1.0)
gl.end_list()

// List 2: GLUT Solid Cone
gl.new_list(list_start + 2, gl.COMPILE)
glut.solid_cone(0.7, 1.2, 20, 10)
gl.end_list()

// List 3: GLUT Solid Torus
gl.new_list(list_start + 3, gl.COMPILE)
glut.solid_torus(0.4, 0.7, 16, 16)
gl.end_list()

// List 4: GLUT Solid Teapot
gl.new_list(list_start + 4, gl.COMPILE)
glut.solid_teapot(0.6)
gl.end_list()

// List 5: GLUT Solid Dodecahedron
gl.new_list(list_start + 5, gl.COMPILE)
glut.solid_dodecahedron()
gl.end_list()

// List 6: GLU Sphere (Quadric)
gl.new_list(list_start + 6, gl.COMPILE)
glu.sphere(quadric, 0.6, 32, 32)
gl.end_list()

// List 7: GLU Cylinder
gl.new_list(list_start + 7, gl.COMPILE)
glu.cylinder(quadric, 0.6, 0.6, 1.2, 32, 16)
gl.end_list()

// List 8: GLU Cone (Cylinder with top = 0)
gl.new_list(list_start + 8, gl.COMPILE)
glu.cylinder(quadric, 0.7, 0.0, 1.2, 32, 16)
gl.end_list()

// List 9: GLU Disk
gl.new_list(list_start + 9, gl.COMPILE)
glu.disk(quadric, 0.0, 0.7, 32, 16)
gl.end_list()

// List 10: GLU Partial Disk
gl.new_list(list_start + 10, gl.COMPILE)
glu.partial_disk(quadric, 0.2, 0.7, 32, 16, 45.0, 270.0)
gl.end_list()

// List 11: GLUT Solid Icosahedron
gl.new_list(list_start + 11, gl.COMPILE)
glut.solid_icosahedron()
gl.end_list()

// Object positions in a circle
var object_angles = []
var num_objects = 12
for i in range(0, num_objects) {
    object_angles.push(i * 360.0 / num_objects)
}

var object_names = [
    "Sphere (GLUT)",
    "Cube (GLUT)",
    "Cone (GLUT)",
    "Torus (GLUT)",
    "Teapot (GLUT)",
    "Dodecahedron (GLUT)",
    "Sphere (GLU)",
    "Cylinder (GLU)",
    "Cone (GLU)",
    "Disk (GLU)",
    "Partial Disk (GLU)",
    "Icosahedron (GLUT)"
]

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
    for i in range(-5, 6) {
        gl.vertex3f(-5.0, -1.5, i)
        gl.vertex3f(5.0, -1.5, i)
        gl.vertex3f(i, -1.5, -5.0)
        gl.vertex3f(i, -1.5, 5.0)
    }
    gl.end()
    
    // Draw objects in a circle
    for i in range(0, num_objects) {
        angle := object_angles[i] * math.pi / 180.0
        x := 3.5 * math.cos(angle)
        z := 3.5 * math.sin(angle)
        
        gl.push_matrix()
        gl.translatef(x, 0.0, z)
        
        // Rotate each object to face center
        gl.rotatef(-object_angles[i], 0.0, 1.0, 0.0)
        
        // Highlight selected object
        if i == selected_object {
            gl.color3f(1.0, 1.0, 1.0)
            glut.wire_sphere(0.9, 16, 12)
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
        selected_object = (selected_object + 1) % num_objects
        println("Selected:", selected_object, "-", object_names[selected_object])
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
    if key == " " {
        // Reset camera
        cam_rot_x = 0.0
        cam_rot_y = 0.0
        cam_distance = 8.0
        println("Camera reset")
    }
    if key >= "1" && key <= "9" {
        idx := int(key) - int('1')
        if idx < num_objects {
            selected_object = idx
            println("Selected:", selected_object, "-", object_names[selected_object])
        }
    }
    if key == "0" {
        selected_object = 9
        println("Selected:", selected_object, "-", object_names[selected_object])
    }
})

glut.special_func(fn(key, x, y) {
    if key == glut.KEY_UP {
        cam_distance -= 0.5
    } else if key == glut.KEY_DOWN {
        cam_distance += 0.5
    } else if key == glut.KEY_LEFT {
        // Rotate objects
        for i in range(0, num_objects) {
            object_angles[i] += 5.0
        }
    } else if key == glut.KEY_RIGHT {
        for i in range(0, num_objects) {
            object_angles[i] -= 5.0
        }
    }
    // Clamp distance
    if cam_distance < 2.0 {
        cam_distance = 2.0
    } else if cam_distance > 20.0 {
        cam_distance = 20.0
    }
})

// Display help in console
println("=== Interactive 3D Demo with GLU Shapes ===")
println("Controls:")
println("  Left Mouse Drag: Rotate camera")
println("  Right Click: Cycle through objects")
println("  Keys 1-9, 0: Select specific object")
println("  Space: Reset camera")
println("  Arrow Up/Down: Zoom in/out")
println("  Arrow Left/Right: Rotate objects")
println("  Q: Quit")
println("")
println("Objects:")
for i in range(0, num_objects) {
    println("  ", i, ":", object_names[i])
}

glut.main_loop()