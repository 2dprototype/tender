import "gl"
import "glut"
import "glu"
import "os"

glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE | glut.DEPTH)
glut.init_window_size(400, 400)
glut.create_window("FBO Render to Texture")

// Create texture for FBO
var tex_id
var tex_ids = gl.gen_textures(1)
tex_id = tex_ids[0]
gl.bind_texture(gl.TEXTURE_2D, tex_id)
gl.tex_image2d(gl.TEXTURE_2D, 0, gl.RGBA, 512, 512, 0, gl.RGBA, gl.UNSIGNED_BYTE, null)
gl.tex_parameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
gl.tex_parameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

// Create FBO
var fbo_ids = gl.gen_framebuffers(1)
var fbo = fbo_ids[0]
gl.bind_framebuffer(gl.FRAMEBUFFER, fbo)
gl.framebuffer_texture2d(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, tex_id, 0)
gl.bind_framebuffer(gl.FRAMEBUFFER, 0)

// Rotation for scene
var rot_angle = 0.0

glut.display_func(fn() {
    // First pass: Render to FBO
    gl.bind_framebuffer(gl.FRAMEBUFFER, fbo)
    gl.viewport(0, 0, 512, 512)
    gl.clear_color(0.2, 0.3, 0.4, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    glu.perspective(60.0, 1.0, 0.1, 100.0)
    
    gl.matrix_mode(gl.MODELVIEW)
    gl.load_identity()
    gl.translatef(0.0, 0.0, -5.0)
    gl.rotatef(rot_angle, 1.0, 1.0, 0.0)
    
    // Draw something to the FBO
    gl.color3f(1.0, 0.0, 0.0)
    glut.solid_teapot(1.0)
    
    // Second pass: Render to screen using FBO texture
    gl.bind_framebuffer(gl.FRAMEBUFFER, 0)
    var w = glut.get(glut.WINDOW_WIDTH)
    var h = glut.get(glut.WINDOW_HEIGHT)
    gl.viewport(0, 0, w, h)
    gl.clear_color(0.1, 0.1, 0.1, 1.0)
    gl.clear(gl.COLOR_BUFFER_BIT)
    
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    gl.ortho(0.0, 1.0, 1.0, 0.0, -1.0, 1.0)
    
    gl.matrix_mode(gl.MODELVIEW)
    gl.load_identity()
    
    // Display the FBO texture
    gl.enable(gl.TEXTURE_2D)
    gl.bind_texture(gl.TEXTURE_2D, tex_id)
    gl.begin(gl.QUADS)
    gl.tex_coord2f(0.0, 0.0)
    gl.vertex2f(0.0, 0.0)
    gl.tex_coord2f(1.0, 0.0)
    gl.vertex2f(1.0, 0.0)
    gl.tex_coord2f(1.0, 1.0)
    gl.vertex2f(1.0, 1.0)
    gl.tex_coord2f(0.0, 1.0)
    gl.vertex2f(0.0, 1.0)
    gl.end()
    gl.disable(gl.TEXTURE_2D)
    
    glut.swap_buffers()
})

glut.reshape_func(fn(w, h) {
    gl.viewport(0, 0, w, h)
})

glut.idle_func(fn() {
    rot_angle += 0.5
    glut.post_redisplay()
})

glut.keyboard_func(fn(key, x, y) {
    if key == "q" || key == "Q" { os.exit(0) }
})

glut.main_loop()