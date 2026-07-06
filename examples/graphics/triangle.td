import "graphics"

win := graphics.new_window(400, 400, "Hello Triangle")

win.on_draw(fn() {
    win.clear(0, 0, 0)
    win.rgb(1, 0, 0)
    win.move_to(200, 50)
    win.line_to(50, 350)
    win.line_to(350, 350)
    win.close_path()
    win.fill()
})

win.run()