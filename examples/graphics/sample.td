import "graphics"
import { sin } from "math"

win := graphics.new_window(640, 480, "My Tender Game")

// Game state
pos_x := 320.0
pos_y := 240.0
angle := 0.0
scale := 1.0
color_phase := 0.0

// Keyboard state
keys := {}

win.on_key(fn(key, x, y) {
    keys[key] = true

    if key == "q" {
        println("Goodbye!")
    }
    if key == "r" {
        println("Reset position")
        pos_x = 320
        pos_y = 240
    }
    if key == "j" {
        println("Jump!")
    }
})

win.on_mouse(fn(button, action, x, y) {
    if button == "left" && action == "down" {
        println("Clicked at:", x, y)
        pos_x = x
        pos_y = y
    }
})

win.on_mouse_move(fn(x, y) {
    // Optional: show mouse position
    // println("Mouse at:", x, y)
})

win.on_update(fn() {
    // Movement with WASD
    if keys["w"] { pos_y -= 3 }
    if keys["s"] { pos_y += 3 }
    if keys["a"] { pos_x -= 3 }
    if keys["d"] { pos_x += 3 }
    
    // Clamp to screen
    if pos_x < 0 { pos_x = 0 }
    if pos_x > 640 { pos_x = 640 }
    if pos_y < 0 { pos_y = 0 }
    if pos_y > 480 { pos_y = 480 }
    
    // Update animations
    angle += 0.03
    scale = 1.0 + 0.3 * sin(angle * 2)
    color_phase += 0.02
})

win.on_draw(fn() {
    win.clear(0.1, 0.1, 0.2)
    // Calculate color based on position and time
    r := 0.5 + 0.5 * sin(color_phase)
    g := 0.5 + 0.5 * sin(color_phase + 2.0)
    b := 0.5 + 0.5 * sin(color_phase + 4.0)
    
    win.push()
    win.translate(pos_x, pos_y)
    win.scale(scale, scale)
    win.rotate(angle)
    win.translate(-30, -30)
    
    // Draw the player with rainbow colors
    win.rgb(r, g, b)
    win.rect(0, 0, 60, 60)
    win.fill()
    
    // Draw border
    win.rgb(1, 1, 1)
    win.set_line_width(3)
    win.stroke()
    
    win.pop()

    // Use hex color with alpha
    win.hex("#ff6b3544")
    win.circle(pos_x - 80, pos_y - 80, 30)
    win.fill()
    
    win.hex("#ff6b35")
    win.set_line_width(2)
    win.stroke()

    // A green rectangle with rounded corners (simulated)
    win.hex("#00ff0088")
    win.rect(500, 50, 100, 80)
    win.fill()
    win.hex("#00ff00")
    win.set_line_width(1)
    win.stroke()
    
    // A yellow line
    win.hex("#ffff00")
    win.set_line_width(3)
    win.line(50, 400, 200, 400)
    win.stroke()

    // Load a font (optional - comment if no font file)
    // win.load_font("assets/arial.ttf", 24)
    
    // Draw FPS counter
    win.hex("#ffffff")
    win.draw_string(10, 20, "Hello, Tender!")
    
    // Draw position info
    info := "Pos: (" + int(pos_x) + ", " + int(pos_y) + ")"
    win.draw_string(10, 40, info)
    
    // Draw controls help
    win.draw_string(10, 460, "WASD - Move | R - Reset | Q - Exit")
})

println("Game starting...")
println("Controls: WASD to move, R to reset, Click to teleport")

win.run()
println("Game ended.")