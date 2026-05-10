// tender@v0.0.4

canvas := import("canvas")

ctx := canvas.new_context(100, 100)


res := canvas.load_font("testfont", 12, "IBMPlexSans-Bold.ttf")
ctx.hex("#00ff00")
ctx.text("Wooh", 10, 10)

ctx.save_png("out.png")