// tender@v0.0.7

import "graphics"
import { sin, cos, pi } from "math"
import { printf } from "fmt"

// Create window and query dimensions
win := graphics.new_window(800, 600, "Tender 2D Graphics Engine - Comprehensive Test")
w := win.width
h := win.height
printf("Created window of size: %v x %v\n", w, h)

// Load Resources (Font and Image)
win.load_font("IBMPlexSans-Bold.ttf", 20)
gopher := win.load_image("gopher.png")
printf("Loaded image: width=%v, height=%v\n", gopher.width, gopher.height)

// State variables for animation and user input
angle := 0.0
mouseX := 400
mouseY := 300
lastKey := "None"
mouseBtnState := "None"

// Register Event Callbacks
win.on_key(fn(key, x, y) {
    lastKey = key
    printf("Key pressed: '%v' at mouse position (%v, %v)\n", key, x, y)
})

win.on_mouse(fn(button, action, x, y) {
    mouseBtnState = button + " " + action
    printf("Mouse click: button=%v, action=%v at (%v, %v)\n", button, action, x, y)
})

win.on_mouse_move(fn(x, y) {
    mouseX = x
    mouseY = y
})

win.on_draw(fn() {
    // Test Clear (with Hex Color)
    win.clear("#1a1a24")

    // Test Basic Shapes & Colors (rgb/rgba/hex, set_line_width, line, rect, circle, stroke, fill)
    win.rgba(1.0, 0.0, 0.0, 0.5) // Semi-transparent Red
    win.rect(20, 20, 150, 100)
    win.fill()

    win.hex("#00ffff")
    win.set_line_width(4)
    win.rect(20, 20, 150, 100)
    win.stroke()

    win.rgb(0.0, 1.0, 0.0) // Green Line
    win.line(200, 20, 350, 120)
    win.stroke()

    win.hex("#ffff00") // Yellow Circle outline
    win.set_line_width(2)
    win.circle(450, 70, 40)
    win.stroke()

    // Test Curves (move_to, line_to, quadratic_to, cubic_to, arc, elliptical_arc)
    win.push()
    win.hex("#ff00ff")
    win.set_line_width(3)
    win.move_to(20, 200)
    win.quadratic_to(100, 120, 180, 200)
    win.stroke()

    win.hex("#00ff88")
    win.move_to(200, 200)
    win.cubic_to(250, 120, 300, 280, 350, 200)
    win.stroke()

    // Arc (angle1 to angle2 in radians)
    win.hex("#ff8800")
    win.arc(480, 200, 40, 0.0, pi)
    win.stroke()

    // Elliptical Arc
    win.hex("#ff5555")
    win.elliptical_arc(620, 200, 50, 30, 0.0, 1.5 * pi)
    win.stroke()
    win.pop()

    // Test Transforms & State (push, pop, translate, scale, rotate)
    win.push()
    win.translate(100, 350)
    win.rotate(angle)
    win.scale(1.2, 1.2)
    win.hex("#ffffff")
    win.rect(-30, -30, 60, 60)
    win.set_line_width(2)
    win.stroke()
    win.pop()

    // Test Image Drawing (draw_image, draw_image_rect)
    // Draw full image
    win.draw_image(gopher, 250, 300)

    // Draw portion of image cropped/scaled (spritesheet style)
    // Src: (10, 10, 80, 80) Dest: (450, 300, 120, 120)
    win.draw_image_rect(gopher, 10, 10, 80, 80, 450, 300, 120, 120)

    // Test Text rendering & Measuring (measure_string, set_font_size, draw_string, draw_string_anchored, draw_string_wrapped)
    win.hex("#ffffff")
    win.set_font_size(16)
    win.draw_string(20, 500, "Last Key: " + lastKey)
    win.draw_string(20, 530, "Mouse Button: " + mouseBtnState)
    win.draw_string(20, 560, "Mouse Position: (" + string(mouseX) + ", " + string(mouseY) + ")")

    // Measure String
    m := win.measure_string("Anchored Text")
    win.hex("#ff00ff")
    win.draw_string_anchored(400, 500, "Anchored Text (W:" + string(m.width) + " H:" + string(m.height) + ")", 0.5, 0.5)

    // Wrapped Text
    longText := "This is a very long text string designed to test the word wrap functionality of the Tender 2D graphics engine."
    win.hex("#00ffcc")
    win.draw_string_wrapped(600, 450, longText, 180, 1.3)
})

win.on_update(fn() {
    angle = angle + 0.03
})

//  Run event loop
win.run()
