import "graphics"
import { exit } from "os"

win := graphics.new_window(400, 400, "Graphics Demo")
angle := 0.0

win.on_draw(fn() {
    win.clear("#1a1a2e")
    win.push()
    win.translate(200, 200)
    win.rotate(angle)
    win.hex("#ff6b35")
    win.rect(-50, -50, 100, 100)
    win.fill()
    win.pop()
})

win.on_update(fn() {
    angle += 0.01
})

win.on_key(fn(key, x, y) {
    if key == "q" {
        exit(0)
    }
})

win.run()