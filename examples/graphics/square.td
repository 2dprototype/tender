import "graphics"

win := graphics.new_window(400, 400, "Rotating Square")

angle := 0.0

win.on_update(fn() {
    angle += 0.02
})

win.on_draw(fn() {
    win.clear(0.2, 0.2, 0.3)
    
    win.push()
    win.translate(200, 200)
    win.rotate(angle)
    win.translate(-50, -50)
    
    win.hex("#ff8800")
    win.rect(0, 0, 100, 100)
    win.fill()
    
    win.hex("#ffffff")
    win.set_line_width(2)
    win.stroke()
    
    win.pop()
})

win.run()