import "gl"
import "glut"
import "glu"
import "os"
import "canvas"
import "math"

glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE | glut.DEPTH)
glut.init_window_size(400, 400)
glut.init_window_position(100, 100)

win := glut.create_window("Texture Mapping Demo")

gl.enable(gl.DEPTH_TEST)
gl.enable(gl.TEXTURE_2D)

// Create a procedural texture using canvas
var texture_data = []
var tex_size = 64

ctx := canvas.new_context(tex_size, tex_size)

// Draw a checkerboard pattern
for i in range(0, tex_size) {
    for j in range(0, tex_size) {
        if ((i/8) % 2) == ((j/8) % 2) {
            ctx.rgb(1.0, 1.0, 1.0)
        } else {
            ctx.rgb(0.2, 0.2, 0.2)
        }
        ctx.set_pixel(i, j)
    }
}

// Generate texture
textures := gl.gen_textures(1)
gl.bind_texture(gl.TEXTURE_2D, textures[0])

gl.tex_parameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
gl.tex_parameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
gl.tex_parameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
gl.tex_parameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

// Convert canvas to bytes and upload texture
gl.tex_image2d(gl.TEXTURE_2D, 0, gl.RGB, tex_size, tex_size, 0, gl.RGB, gl.UNSIGNED_BYTE, ctx.image().encode("png"))

var rot = 0.0
var zoom = -5.0

glut.display_func(fn() {
    gl.clear_color(0.1, 0.1, 0.2, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
    gl.matrix_mode(gl.MODELVIEW)
    gl.load_identity()
    gl.translatef(0.0, 0.0, zoom)
    gl.rotatef(rot, 0.0, 1.0, 0.0)
    
    // Draw a textured cube
    gl.bind_texture(gl.TEXTURE_2D, textures[0])
    gl.begin(gl.QUADS)
    
    // Front face
    gl.tex_coord2f(0.0, 0.0)
    gl.vertex3f(-1.0, -1.0,  1.0)
    gl.tex_coord2f(1.0, 0.0)
    gl.vertex3f( 1.0, -1.0,  1.0)
    gl.tex_coord2f(1.0, 1.0)
    gl.vertex3f( 1.0,  1.0,  1.0)
    gl.tex_coord2f(0.0, 1.0)
    gl.vertex3f(-1.0,  1.0,  1.0)
    
    // Back face
    gl.tex_coord2f(0.0, 0.0)
    gl.vertex3f(-1.0, -1.0, -1.0)
    gl.tex_coord2f(0.0, 1.0)
    gl.vertex3f(-1.0,  1.0, -1.0)
    gl.tex_coord2f(1.0, 1.0)
    gl.vertex3f( 1.0,  1.0, -1.0)
    gl.tex_coord2f(1.0, 0.0)
    gl.vertex3f( 1.0, -1.0, -1.0)
    
    // Top face
    gl.tex_coord2f(0.0, 0.0)
    gl.vertex3f(-1.0,  1.0, -1.0)
    gl.tex_coord2f(0.0, 1.0)
    gl.vertex3f(-1.0,  1.0,  1.0)
    gl.tex_coord2f(1.0, 1.0)
    gl.vertex3f( 1.0,  1.0,  1.0)
    gl.tex_coord2f(1.0, 0.0)
    gl.vertex3f( 1.0,  1.0, -1.0)
    
    // Bottom face
    gl.tex_coord2f(0.0, 0.0)
    gl.vertex3f(-1.0, -1.0, -1.0)
    gl.tex_coord2f(1.0, 0.0)
    gl.vertex3f( 1.0, -1.0, -1.0)
    gl.tex_coord2f(1.0, 1.0)
    gl.vertex3f( 1.0, -1.0,  1.0)
    gl.tex_coord2f(0.0, 1.0)
    gl.vertex3f(-1.0, -1.0,  1.0)
    
    // Right face
    gl.tex_coord2f(0.0, 0.0)
    gl.vertex3f( 1.0, -1.0, -1.0)
    gl.tex_coord2f(1.0, 0.0)
    gl.vertex3f( 1.0,  1.0, -1.0)
    gl.tex_coord2f(1.0, 1.0)
    gl.vertex3f( 1.0,  1.0,  1.0)
    gl.tex_coord2f(0.0, 1.0)
    gl.vertex3f( 1.0, -1.0,  1.0)
    
    // Left face
    gl.tex_coord2f(0.0, 0.0)
    gl.vertex3f(-1.0, -1.0, -1.0)
    gl.tex_coord2f(0.0, 1.0)
    gl.vertex3f(-1.0, -1.0,  1.0)
    gl.tex_coord2f(1.0, 1.0)
    gl.vertex3f(-1.0,  1.0,  1.0)
    gl.tex_coord2f(1.0, 0.0)
    gl.vertex3f(-1.0,  1.0, -1.0)
    
    gl.end()
    
    glut.swap_buffers()
})

glut.reshape_func(fn(w, h) {
    gl.viewport(0, 0, w, h)
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    glu.perspective(45.0, w/h, 0.1, 100.0)
})

glut.keyboard_func(fn(key, x, y) {
    if key == "q" || key == "Q" {
        os.exit(0)
    }
})

glut.special_func(fn(key, x, y) {
    if key == glut.KEY_UP {
        zoom += 0.5
    } else if key == glut.KEY_DOWN {
        zoom -= 0.5
    }
})

glut.idle_func(fn() {
    rot += 0.5
    glut.post_redisplay()
})

glut.main_loop()