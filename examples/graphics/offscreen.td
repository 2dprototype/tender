import "graphics"

// Offscreen rendering example
ctx := graphics.new_context(400, 400)
ctx.hex("#2a2a4a")
ctx.clear()
ctx.hex("#00ff88")
ctx.set_line_width(3)
ctx.roundrect(50, 50, 300, 200, 20, 20)
ctx.stroke()
ctx.hex("#ff6644")
ctx.circle(200, 150, 60)
ctx.fill()
ctx.hex("#0f0")
ctx.draw_string(10, 20, "hello")
ctx.save("output.png")