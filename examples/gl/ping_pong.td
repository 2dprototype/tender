import "gl"
import "glut"
import "os"
import "rand"
import "math"

// Window dimensions
const WIDTH = 400
const HEIGHT = 400

// Paddle properties
const PADDLE_WIDTH = 10
const PADDLE_HEIGHT = 60
const PADDLE_SPEED = 5

// Ball properties
const BALL_SIZE = 8
const BALL_SPEED_X = 4
const BALL_SPEED_Y = 3

// Game state
leftY := HEIGHT/2 - PADDLE_HEIGHT/2
rightY := HEIGHT/2 - PADDLE_HEIGHT/2
ballX := WIDTH/2
ballY := HEIGHT/2
ballVx := BALL_SPEED_X
ballVy := BALL_SPEED_Y
scoreLeft := 0
scoreRight := 0

// Keyboard state
keyW := false
keyS := false
keyUp := false
keyDown := false

// Helper: draw a filled rectangle
fn draw_rect(x1, y1, x2, y2) {
    gl.begin(gl.QUADS)
    gl.vertex2f(x1, y1)
    gl.vertex2f(x2, y1)
    gl.vertex2f(x2, y2)
    gl.vertex2f(x1, y2)
    gl.end()
}

// Helper: draw text using GLUT bitmap font
fn draw_text(x, y, text) {
    gl.raster_pos2i(x, y)
    for i := 0; i < len(text); i++ {
        glut.bitmap_character(glut.BITMAP_HELVETICA_18, text[i])
    }
}

fn rand_int(min, max) {
	return int(math.floor(rand.float() * (max - min + 1)) + min)
}

// Reset ball after a score, direction depends on who scored
fn reset_ball(direction) {
    ballX = WIDTH/2
    ballY = HEIGHT/2
    ballVx = BALL_SPEED_X * direction   // +1 for right, -1 for left
    // Add some randomness to vertical speed
    if rand_int(0, 100) % 2 == 0 {
        ballVy = BALL_SPEED_Y
    } else {
        ballVy = -BALL_SPEED_Y
    }
}

// Update game state each frame
fn update() {
    // Move paddles
    if keyW && leftY > 0 {
        leftY -= PADDLE_SPEED
    }
    if keyS && leftY < HEIGHT - PADDLE_HEIGHT {
        leftY += PADDLE_SPEED
    }
    if keyUp && rightY > 0 {
        rightY -= PADDLE_SPEED
    }
    if keyDown && rightY < HEIGHT - PADDLE_HEIGHT {
        rightY += PADDLE_SPEED
    }

    // Move ball
    ballX += ballVx
    ballY += ballVy

    // Wall bounce (top/bottom)
    if ballY - BALL_SIZE/2 < 0 || ballY + BALL_SIZE/2 > HEIGHT {
        ballVy = -ballVy
    }

    // Left paddle collision
    if ballVx < 0 &&
       ballX - BALL_SIZE/2 < PADDLE_WIDTH &&
       ballX - BALL_SIZE/2 > 0 {
        if ballY > leftY && ballY < leftY + PADDLE_HEIGHT {
            ballVx = -ballVx
            // Adjust vertical speed based on where on paddle the ball hit
            rel := (ballY - (leftY + PADDLE_HEIGHT/2)) / (PADDLE_HEIGHT/2)
            ballVy = rel * BALL_SPEED_Y * 1.5
        }
    }

    // Right paddle collision
    if ballVx > 0 &&
       ballX + BALL_SIZE/2 > WIDTH - PADDLE_WIDTH &&
       ballX + BALL_SIZE/2 < WIDTH {
        if ballY > rightY && ballY < rightY + PADDLE_HEIGHT {
            ballVx = -ballVx
            rel := (ballY - (rightY + PADDLE_HEIGHT/2)) / (PADDLE_HEIGHT/2)
            ballVy = rel * BALL_SPEED_Y * 1.5
        }
    }

    // Scoring
    if ballX - BALL_SIZE/2 < 0 {
        scoreRight++
        reset_ball(1)   // ball goes right
    } else if ballX + BALL_SIZE/2 > WIDTH {
        scoreLeft++
        reset_ball(-1)  // ball goes left
    }
}

// Display callback
fn display() {
    gl.clear(gl.COLOR_BUFFER_BIT)
    gl.load_identity()

    // Draw center line (dashed effect with a solid line)
    gl.color3f(0.4, 0.4, 0.4)
    gl.line_width(1)
    gl.begin(gl.LINES)
    gl.vertex2f(WIDTH/2, 0)
    gl.vertex2f(WIDTH/2, HEIGHT)
    gl.end()

    // Draw paddles (white)
    gl.color3f(1.0, 1.0, 1.0)
    draw_rect(0, leftY, PADDLE_WIDTH, leftY + PADDLE_HEIGHT)
    draw_rect(WIDTH - PADDLE_WIDTH, rightY, WIDTH, rightY + PADDLE_HEIGHT)

    // Draw ball (white)
    draw_rect(ballX - BALL_SIZE/2, ballY - BALL_SIZE/2,
              ballX + BALL_SIZE/2, ballY + BALL_SIZE/2)

    // Draw scores
    scoreText := "Score: " + string(scoreLeft) + " - " + string(scoreRight)
    gl.color3f(1.0, 1.0, 1.0)
    draw_text(WIDTH/2 - 80, HEIGHT - 30, scoreText)

    glut.swap_buffers()
}

// Reshape callback
fn reshape(w, h) {
    gl.viewport(0, 0, w, h)
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    gl.ortho(0, WIDTH, HEIGHT, 0, -1, 1)  // y-axis inverted: origin top‑left
    gl.matrix_mode(gl.MODELVIEW)
}

// Keyboard press
fn keyboard(key, x, y) {
    if key == "q" || key == "Q" {
        os.exit(0)
    }
    if key == "w" || key == "W" { keyW = true }
    if key == "s" || key == "S" { keyS = true }
}

// Keyboard release
fn keyboard_up(key, x, y) {
    if key == "w" || key == "W" { keyW = false }
    if key == "s" || key == "S" { keyS = false }
}

// Special key press (arrows)
fn special(key, x, y) {
    if key == glut.KEY_UP   { keyUp = true }
    if key == glut.KEY_DOWN { keyDown = true }
}

// Special key release
fn special_up(key, x, y) {
    if key == glut.KEY_UP   { keyUp = false }
    if key == glut.KEY_DOWN { keyDown = false }
}

// Timer callback – runs at ~60 FPS
fn timer(val) {
    update()
    glut.post_redisplay()
    glut.timer_func(16, timer, 0)
}

// ----- Main program -----
glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE)
glut.init_window_size(WIDTH, HEIGHT)
glut.init_window_position(100, 100)
glut.create_window("Ping Pong")

gl.clear_color(0.0, 0.0, 0.0, 1.0)

// Register callbacks
glut.display_func(display)
glut.reshape_func(reshape)
glut.keyboard_func(keyboard)
glut.keyboard_up_func(keyboard_up)
glut.special_func(special)
glut.special_up_func(special_up)

// Start game loop
glut.timer_func(16, timer, 0)
glut.main_loop()