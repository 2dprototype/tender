import "gl"
import "glut"
import "os"
import "math"
import "rand"
import "times"

// Constants
const GRID_WIDTH = 10
const GRID_HEIGHT = 20
const BLOCK_SIZE = 20
const WINDOW_WIDTH = GRID_WIDTH * BLOCK_SIZE + 200
const WINDOW_HEIGHT = GRID_HEIGHT * BLOCK_SIZE + 40

// Tetromino shapes
const SHAPES = [
    // I
    [[1, 1, 1, 1]],
    // O
    [[1, 1],
     [1, 1]],
    // T
    [[0, 1, 0],
     [1, 1, 1]],
    // S
    [[0, 1, 1],
     [1, 1, 0]],
    // Z
    [[1, 1, 0],
     [0, 1, 1]],
    // L
    [[1, 0, 0],
     [1, 1, 1]],
    // J
    [[0, 0, 1],
     [1, 1, 1]]
]

const COLORS = [
    [0.0, 0.0, 0.0],      // Empty
    [0.0, 1.0, 1.0],      // I - Cyan
    [1.0, 1.0, 0.0],      // O - Yellow
    [1.0, 0.0, 1.0],      // T - Purple
    [0.0, 1.0, 0.0],      // S - Green
    [1.0, 0.0, 0.0],      // Z - Red
    [1.0, 0.5, 0.0],      // L - Orange
    [0.0, 0.0, 1.0]       // J - Blue
]

// Game state
grid := []
for i := 0; i < GRID_HEIGHT; i++ {
    row := []
    for j := 0; j < GRID_WIDTH; j++ {
        row = append(row, 0)
    }
    grid = append(grid, row)
}

current_piece := 0
current_rotation := 0
current_x := GRID_WIDTH / 2 - 1
current_y := 0
next_piece := 0
score := 0
level := 1
lines_cleared := 0
game_over := false
drop_counter := 0
drop_interval := 30
move_cooldown := 0
das_active := false
const DAS_DELAY = 12 // Frames to wait before auto-repeat kicks in
const ARR = 3        // Auto-Repeat Rate (frames per repeated move)

// Keyboard state
key_left := false
key_right := false
key_down := false
key_rotate := false

fn rand_int() {
	min := 1
	max := 100
	return int(math.floor(rand.float() * (max - min + 1)) + min)
}

// Helper: Get piece shape with rotation
fn get_piece_shape(piece, rotation) {
    shape := SHAPES[piece]
    rows := len(shape)
    cols := len(shape[0])
    
    if rotation == 0 {
        return shape
    }
    
    rotated := []
    for i := 0; i < cols; i++ {
        row := []
        for j := 0; j < rows; j++ {
            if rotation == 1 {
                // 90 degrees clockwise
                row = append(row, shape[rows - 1 - j][i])
            } else if rotation == 2 {
                // 180 degrees
                row = append(row, shape[rows - 1 - i][cols - 1 - j])
            } else if rotation == 3 {
                // 270 degrees clockwise
                row = append(row, shape[j][cols - 1 - i])
            }
        }
        rotated = append(rotated, row)
    }
    return rotated
}

// Helper: Check if piece can be placed
fn can_place(shape, x, y) {
    rows := len(shape)
    cols := len(shape[0])
    
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if shape[i][j] != 0 {
                grid_x := x + j
                grid_y := y + i
                
                if grid_x < 0 || grid_x >= GRID_WIDTH || grid_y >= GRID_HEIGHT {
                    return false
                }
                if grid_y >= 0 && grid[grid_y][grid_x] != 0 {
                    return false
                }
            }
        }
    }
    return true
}

// Helper: Lock current piece in grid
fn lock_piece() {
    shape := get_piece_shape(current_piece, current_rotation)
    rows := len(shape)
    cols := len(shape[0])
    
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if shape[i][j] != 0 {
                grid_x := current_x + j
                grid_y := current_y + i
                
                if grid_y < 0 {
                    game_over = true
                    return
                }
                grid[grid_y][grid_x] = current_piece + 1
            }
        }
    }
    
    // Check for completed lines
    lines := 0
    i := GRID_HEIGHT - 1
    for i >= 0 {
        full := true
        for j := 0; j < GRID_WIDTH; j++ {
            if grid[i][j] == 0 {
                full = false
                break
            }
        }
        if full {
            // Remove line
            k := i
            for k > 0 {
                grid[k] = grid[k-1]
                k--
            }
            // Create new empty top row
            new_row := []
            for j := 0; j < GRID_WIDTH; j++ {
                new_row = append(new_row, 0)
            }
            grid[0] = new_row
            lines++
            // Recheck same row (don't decrement i)
        } else {
            i--
        }
    }
    
    // Update score and level
    if lines > 0 {
        lines_cleared += lines
        if lines == 1 {
            score += 100 * level
        } else if lines == 2 {
            score += 300 * level
        } else if lines == 3 {
            score += 500 * level
        } else if lines == 4 {
            score += 800 * level
        }
        level = lines_cleared / 10 + 1
        drop_interval = 30 - (level - 1) * 2
        if drop_interval < 5 {
            drop_interval = 5
        }
    }
    
    // Spawn new piece
    current_piece = next_piece
    next_piece = rand_int() % len(SHAPES)
    current_rotation = 0
    current_x = GRID_WIDTH / 2 - 1
    current_y = 0
    
    // Check if new piece can be placed
    shape2 := get_piece_shape(current_piece, current_rotation)
    if !can_place(shape2, current_x, current_y) {
        game_over = true
    }
}

// Helper: Rotate piece
fn rotate_piece() {
    shape := get_piece_shape(current_piece, (current_rotation + 1) % 4)
    if can_place(shape, current_x, current_y) {
        current_rotation = (current_rotation + 1) % 4
    }
}

// Helper: Move piece
fn move_piece(dx, dy) {
    shape := get_piece_shape(current_piece, current_rotation)
    if can_place(shape, current_x + dx, current_y + dy) {
        current_x += dx
        current_y += dy
        return true
    }
    return false
}

// Helper: Hard drop
fn hard_drop() {
    shape := get_piece_shape(current_piece, current_rotation)
    for can_place(shape, current_x, current_y + 1) {
        current_y++
    }
    lock_piece()
}

// Update game
fn update() {
    if game_over {
        return
    }
    
    // Handle left/right with DAS (Delayed Auto Shift)
    if key_left || key_right {
        if move_cooldown == 0 {
            // Move immediately
            if key_left {
                move_piece(-1, 0)
            } else if key_right {
                move_piece(1, 0)
            }
            
            // Set cooldown
            if !das_active {
                move_cooldown = DAS_DELAY // Initial pause
                das_active = true
            } else {
                move_cooldown = ARR       // Fast repeat
            }
        } else {
            move_cooldown--
        }
    } else {
        move_cooldown = 0
        das_active = false
	}
    
    // Handle Rotation
    if key_rotate {
        rotate_piece()
        key_rotate = false
    }
    
    // Auto-drop & Soft drop
    drop_speed := 1
    if key_down {
        drop_speed = 4 // Multiplier for soft dropping
    }
    
    drop_counter += drop_speed
    if drop_counter >= drop_interval {
        drop_counter = 0
        if !move_piece(0, 1) {
            lock_piece()
        }
    }
}

// Draw functions
fn draw_block(x, y, color_index) {
    px := x * BLOCK_SIZE + 10
    py := y * BLOCK_SIZE + 30
    
    // Main block
    color := COLORS[color_index]
    gl.color3f(color[0], color[1], color[2])
    gl.begin(gl.QUADS)
    gl.vertex2f(px, py)
    gl.vertex2f(px + BLOCK_SIZE, py)
    gl.vertex2f(px + BLOCK_SIZE, py + BLOCK_SIZE)
    gl.vertex2f(px, py + BLOCK_SIZE)
    gl.end()
    
    // Border
    gl.color3f(0.3, 0.3, 0.3)
    gl.line_width(1)
    gl.begin(gl.LINE_LOOP)
    gl.vertex2f(px, py)
    gl.vertex2f(px + BLOCK_SIZE, py)
    gl.vertex2f(px + BLOCK_SIZE, py + BLOCK_SIZE)
    gl.vertex2f(px, py + BLOCK_SIZE)
    gl.end()
    
    // Highlight
    if color_index > 0 {
        gl.color3f(0.8, 0.8, 0.8)
        gl.begin(gl.LINES)
        gl.vertex2f(px + 2, py + 2)
        gl.vertex2f(px + BLOCK_SIZE - 2, py + 2)
        gl.vertex2f(px + 2, py + 2)
        gl.vertex2f(px + 2, py + BLOCK_SIZE - 2)
        gl.end()
    }
}

fn draw_grid() {
    // Draw placed blocks
    y := 0
    for y < GRID_HEIGHT {
        x := 0
        for x < GRID_WIDTH {
            if grid[y][x] != 0 {
                draw_block(x, y, grid[y][x])
            }
            x++
        }
        y++
    }
    
    // Draw current piece
    shape := get_piece_shape(current_piece, current_rotation)
    rows := len(shape)
    cols := len(shape[0])
    
    i := 0
    for i < rows {
        j := 0
        for j < cols {
            if shape[i][j] != 0 {
                grid_x := current_x + j
                grid_y := current_y + i
                if grid_y >= 0 {
                    draw_block(grid_x, grid_y, current_piece + 1)
                }
            }
            j++
        }
        i++
    }
    
    // Draw ghost piece (where piece will land)
    ghost_y := current_y
    shape2 := get_piece_shape(current_piece, current_rotation)
    for can_place(shape2, current_x, ghost_y + 1) {
        ghost_y++
    }
    
    if ghost_y != current_y {
        gl.color4f(0.5, 0.5, 0.5, 0.3)
        i := 0
        for i < rows {
            j := 0
            for j < cols {
                if shape2[i][j] != 0 {
                    grid_x := current_x + j
                    grid_y := ghost_y + i
                    if grid_y >= 0 {
                        px := grid_x * BLOCK_SIZE + 10
                        py := grid_y * BLOCK_SIZE + 30
                        gl.begin(gl.QUADS)
                        gl.vertex2f(px, py)
                        gl.vertex2f(px + BLOCK_SIZE, py)
                        gl.vertex2f(px + BLOCK_SIZE, py + BLOCK_SIZE)
                        gl.vertex2f(px, py + BLOCK_SIZE)
                        gl.end()
                    }
                }
                j++
            }
            i++
        }
    }
}


fn draw_info() {
    // Draw next piece preview
    gl.color3f(1.0, 1.0, 1.0)
    gl.raster_pos2i(GRID_WIDTH * BLOCK_SIZE + 30, 60)
    next_text := "Next:"
    idx := 0
    for idx < len(next_text) {
        glut.bitmap_character(glut.BITMAP_HELVETICA_18, next_text[idx])
        idx++
    }
    
    next_shape := get_piece_shape(next_piece, 0)
    rows := len(next_shape)
    cols := len(next_shape[0])
    
    i := 0
    for i < rows {
        j := 0
        for j < cols {
            if next_shape[i][j] != 0 {
                px := GRID_WIDTH * BLOCK_SIZE + 30 + j * 25
                py := 80 + i * 25
                color := COLORS[next_piece + 1]
                gl.color3f(color[0], color[1], color[2])
                gl.begin(gl.QUADS)
                gl.vertex2f(px, py)
                gl.vertex2f(px + 20, py)
                gl.vertex2f(px + 20, py + 20)
                gl.vertex2f(px, py + 20)
                gl.end()
            }
            j++
        }
        i++
    }
    
    // Draw score
    gl.color3f(1.0, 1.0, 1.0)
    gl.raster_pos2i(GRID_WIDTH * BLOCK_SIZE + 30, 180)
    score_text := "Score: " + string(score)
    idx = 0
    for idx < len(score_text) {
        glut.bitmap_character(glut.BITMAP_HELVETICA_18, score_text[idx])
        idx++
    }
    
    // Draw level
    gl.raster_pos2i(GRID_WIDTH * BLOCK_SIZE + 30, 210)
    level_text := "Level: " + string(level)
    idx = 0
    for idx < len(level_text) {
        glut.bitmap_character(glut.BITMAP_HELVETICA_18, level_text[idx])
        idx++
    }
    
    // Draw lines
    gl.raster_pos2i(GRID_WIDTH * BLOCK_SIZE + 30, 240)
    lines_text := "Lines: " + string(lines_cleared)
    idx = 0
    for idx < len(lines_text) {
        glut.bitmap_character(glut.BITMAP_HELVETICA_18, lines_text[idx])
        idx++
    }
    
    // Game over message
    if game_over {
        gl.color3f(1.0, 0.0, 0.0)
        gl.raster_pos2i(GRID_WIDTH * BLOCK_SIZE / 2 - 60, WINDOW_HEIGHT / 2)
        game_over_text := "GAME OVER!"
        idx = 0
        for idx < len(game_over_text) {
            glut.bitmap_character(glut.BITMAP_HELVETICA_18, game_over_text[idx])
            idx++
        }
        gl.raster_pos2i(GRID_WIDTH * BLOCK_SIZE / 2 - 70, WINDOW_HEIGHT / 2 + 30)
        restart_text := "Press R to restart"
        idx = 0
        for idx < len(restart_text) {
            glut.bitmap_character(glut.BITMAP_HELVETICA_18, restart_text[idx])
            idx++
        }
    }
}

// Display callback
fn display() {
    gl.clear(gl.COLOR_BUFFER_BIT)
    gl.load_identity()
    
    // Draw background
    gl.color3f(0.1, 0.1, 0.15)
    gl.begin(gl.QUADS)
    gl.vertex2f(0, 0)
    gl.vertex2f(WINDOW_WIDTH, 0)
    gl.vertex2f(WINDOW_WIDTH, WINDOW_HEIGHT)
    gl.vertex2f(0, WINDOW_HEIGHT)
    gl.end()
    
    // Draw game area background
    gl.color3f(0.05, 0.05, 0.08)
    gl.begin(gl.QUADS)
    gl.vertex2f(10, 30)
    gl.vertex2f(GRID_WIDTH * BLOCK_SIZE + 10, 30)
    gl.vertex2f(GRID_WIDTH * BLOCK_SIZE + 10, GRID_HEIGHT * BLOCK_SIZE + 30)
    gl.vertex2f(10, GRID_HEIGHT * BLOCK_SIZE + 30)
    gl.end()
    
    draw_grid()
    draw_info()
    
    glut.swap_buffers()
}

// Reshape callback
fn reshape(w, h) {
    gl.viewport(0, 0, w, h)
    gl.matrix_mode(gl.PROJECTION)
    gl.load_identity()
    gl.ortho(0, w, h, 0, -1, 1)
    gl.matrix_mode(gl.MODELVIEW)
}

// Keyboard callback
fn keyboard(key, x, y) {
    if key == "q" || key == "Q" {
        os.exit(0)
    }
    
    if key == "r" || key == "R" {
        // Restart game
        i := 0
        for i < GRID_HEIGHT {
            j := 0
            for j < GRID_WIDTH {
                grid[i][j] = 0
                j++
            }
            i++
        }
        score = 0
        level = 1
        lines_cleared = 0
        game_over = false
        current_piece = rand_int() % len(SHAPES)
        next_piece = (rand_int() + 1) % len(SHAPES)
        current_rotation = 0
        current_x = GRID_WIDTH / 2 - 1
        current_y = 0
        drop_counter = 0
        drop_interval = 30
        return
    }
    
    if key == " " {
        if !game_over {
            hard_drop()
        }
        return
    }
    
    if key == "w" || key == "W" {
        if !game_over {
            key_rotate = true
        }
    }
    
    if key == "a" || key == "A" {
        key_left = true
    }
    if key == "d" || key == "D" {
        key_right = true
    }
    if key == "s" || key == "S" {
        key_down = true
    }
}

// Keyboard up callback
fn keyboard_up(key, x, y) {
    if key == "a" || key == "A" {
        key_left = false
    }
    if key == "d" || key == "D" {
        key_right = false
    }
    if key == "s" || key == "S" {
        key_down = false
    }
}

// Special key callback
fn special(key, x, y) {
    if key == glut.KEY_UP {
        if !game_over {
            key_rotate = true
        }
    }
    if key == glut.KEY_LEFT {
        key_left = true
    }
    if key == glut.KEY_RIGHT {
        key_right = true
    }
    if key == glut.KEY_DOWN {
        key_down = true
    }
}

// Special key up callback
fn special_up(key, x, y) {
    if key == glut.KEY_LEFT {
        key_left = false
    }
    if key == glut.KEY_RIGHT {
        key_right = false
    }
    if key == glut.KEY_DOWN {
        key_down = false
    }
}

// Timer callback - runs at ~60 FPS
fn timer(val) {
    update()
    glut.post_redisplay()
    glut.timer_func(16, timer, 0)
}

// Main
glut.init()
gl.init()
glut.init_display_mode(glut.RGBA | glut.DOUBLE)
glut.init_window_size(WINDOW_WIDTH, WINDOW_HEIGHT)
glut.init_window_position(100, 100)
glut.create_window("Tetris")

// Set initial pieces
current_piece = rand_int() % len(SHAPES)
next_piece = (rand_int() + 1) % len(SHAPES)

gl.clear_color(0.1, 0.1, 0.15, 1.0)

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