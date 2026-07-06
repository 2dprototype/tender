import "graphics"
import { sin } from "math"

// Create a window - 800x600 size
win := graphics.new_window(800, 600, "Tender 2D Graphics Demo")

angle := 0.0

// Register keyboard events
win.on_key(fn(key, x, y) {
    if key == "q" or key == "special_27" { // ESC or q to exit
        // Window close handler or logic
        win.hex("#ff0000")
    }
})

// Register mouse events
win.on_mouse(fn(button, action, x, y) {
    if action == "down" {
        // Draw an indicator or react to click
    }
})

// Main draw callback
win.on_draw(fn() {
    // Clear screen with a nice modern dark gray-blue background
    win.clear("#181824")

    // Draw a pulsating circle in the middle
    scale_val := 1.0 + 0.25 * sin(angle * 2.5)
    win.push()
    win.translate(400, 200)
    win.scale(scale_val, scale_val)
    win.hex("#44aaff33") // Semi-transparent blue
    win.circle(0, 0, 80)
    win.fill()
    win.pop()

    // Draw a rotating orange square
    win.push()
    win.translate(400, 450)
    win.rotate(angle)
    win.hex("#ff8800")
    win.rect(-50, -50, 100, 100)
    win.fill()
    win.hex("#ffffff")
    win.set_line_width(2)
    win.stroke()
    win.pop()

    // Draw Bezier Curves
    win.push()
    win.hex("#00ff88")
    win.set_line_width(3)
    win.move_to(100, 300)
    win.quadratic_to(200, 100, 300, 300)
    win.stroke()
    win.pop()

    // Render Text
    win.hex("#ffffff")
    // Simple drawing (falls back to default font if none loaded)
    win.draw_string(20, 40, "Press arrow keys or 'q' to interact")
    win.draw_string_anchored(400, 30, "TENDER HIGH PERFORMANCE 2D", 0.5, 0.5)
})

// Main update callback for physics/animation
win.on_update(fn() {
    angle += 0.02
})

// Run the window event loop
win.run()
