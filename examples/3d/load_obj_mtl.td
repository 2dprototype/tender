import "gl"
import "glut"
import "glu"
import "os"
import "math"

glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE | glut.DEPTH)
glut.init_window_size(440, 440)
glut.init_window_position(100, 100)
glut.create_window("Interactive Anime Toon Engine - MTL Supported")

gl.enable(gl.DEPTH_TEST)
gl.enable(gl.LIGHTING)
gl.enable(gl.LIGHT0)
gl.enable(gl.TEXTURE_2D) 

gl.lightfv(gl.LIGHT0, gl.POSITION, 3.0, 4.0, 5.0, 1.0)
gl.lightfv(gl.LIGHT0, gl.DIFFUSE, 1.0, 1.0, 1.0, 1.0)


var vertex_src = `
#version 120
varying vec3 vNormal;
varying vec3 vLightDir;
varying vec2 vTexCoord;

void main() {
    // Transform normal into eye space
    vNormal = normalize(gl_NormalMatrix * gl_Normal);
    
    // Grab fixed light position from gl.lightfv automatically
    vLightDir = normalize(gl_LightSource[0].position.xyz);
    
    // Forward standard texture coordinates
    vTexCoord = gl_MultiTexCoord0.xy;
    
    gl_Position = gl_ModelViewProjectionMatrix * gl_Vertex;
}
`

var frag_src = `
#version 120
varying vec3 vNormal;
varying vec3 vLightDir;
varying vec2 vTexCoord;
uniform sampler2D u_texture;

void main() {
    // Calculate diffuse light intensity (dot product)
    float intensity = dot(normalize(vNormal), normalize(vLightDir));
    
    // --- Crisp Anime Cel-Shading Steps ---
    float shadowFactor;
    if (intensity > 0.5) {
        shadowFactor = 1.0;  // Full bright base color
    } else if (intensity > 0.15) {
        shadowFactor = 0.65; // Soft midtone shadow
    } else {
        shadowFactor = 0.35; // Deep anime shadow plate
    }
    
    // Sample texture and mix with cel steps
    vec4 texColor = texture2D(u_texture, vTexCoord);
    gl_FragColor = vec4(texColor.rgb * shadowFactor, texColor.a);
}
`

var vs = gl.create_shader(gl.VERTEX_SHADER)
gl.shader_source(vs, vertex_src)
gl.compile_shader(vs)
if gl.get_shader_iv(vs, gl.COMPILE_STATUS) == 0 {
    print("Vertex Shader Error: " + gl.get_shader_info_log(vs))
}

var fs = gl.create_shader(gl.FRAGMENT_SHADER)
gl.shader_source(fs, frag_src)
gl.compile_shader(fs)
if gl.get_shader_iv(fs, gl.COMPILE_STATUS) == 0 {
    print("Fragment Shader Error: " + gl.get_shader_info_log(fs))
}

var toon_program = gl.create_program()
gl.attach_shader(toon_program, vs)
gl.attach_shader(toon_program, fs)
gl.link_program(toon_program)

var tex_uniform_loc = gl.get_uniform_location(toon_program, "u_texture")
mesh_id := gl.load_obj_with_mtl("Lowpoly_tree_sample.obj", "Lowpoly_tree_sample.mtl")

var rot_x = 0.0
var rot_y = 0.0
var pos_x = -0.5 // Offset slightly to center the 0..1 cube
var pos_y = -3
var pos_z = -15.0 // Initial zoom

var is_dragging = false
var last_mouse_x = 0
var last_mouse_y = 0

glut.display_func(fn() {
    gl.clear_color(0.4, 0.55, 0.7, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    gl.matrix_mode(gl.MODELVIEW)
    gl.load_identity()
    gl.translatef(pos_x, pos_y, pos_z)
    gl.rotatef(rot_x, 1.0, 0.0, 0.0)
    gl.rotatef(rot_y, 0.0, 1.0, 0.0)
    gl.use_program(toon_program)
    gl.uniform1i(tex_uniform_loc, 0) 
    gl.draw_obj(mesh_id)
    gl.use_program(0) 
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
        } else {
            is_dragging = false
        }
    }
    
    if button == 3 && state == 0 { pos_z += 0.5; glut.post_redisplay() }
    if button == 4 && state == 0 { pos_z -= 0.5; glut.post_redisplay() }
})

glut.motion_func(fn(x, y) {
    if is_dragging {
        dx := x - last_mouse_x
        dy := y - last_mouse_y
        
        rot_y += dx * 0.5
        rot_x += dy * 0.5
        
        last_mouse_x = x
        last_mouse_y = y
        
        glut.post_redisplay()
    }
})

glut.keyboard_func(fn(key, x, y) {
    if key == "w" || key == "W" { pos_y += 0.2 }
    if key == "s" || key == "S" { pos_y -= 0.2 }
    if key == "a" || key == "A" { pos_x -= 0.2 }
    if key == "d" || key == "D" { pos_x += 0.2 }
    
    if key == "+" || key == "=" { pos_z += 0.5 }
    if key == "-" || key == "_" { pos_z -= 0.5 }
    
    if key == "q" || key == "Q" { os.exit(0) }
    
    glut.post_redisplay()
})

glut.main_loop()