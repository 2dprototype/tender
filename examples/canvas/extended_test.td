// tender@v0.0.4-exp
// Extended Canvas 

import "canvas"
import "times"

var width = 800
var height = 600
var ctx = canvas.new_context(width, height)

// 1. Test Solid Pattern
println("Testing Solid Pattern...")
var solid = canvas.new_solid_pattern("#0000ff") // Blue
ctx.fillstyle(solid)
ctx.rect(10, 10, 100, 100)
ctx.fill()

// 2. Test Linear Gradient
println("Testing Linear Gradient...")
var lin = canvas.new_linear_gradient(150, 10, 250, 110)
lin.add_color_stop(0, "#ff0000") // Red
lin.add_color_stop(1, "#ffff00") // Yellow
ctx.fillstyle(lin)
ctx.rect(150, 10, 100, 100)
ctx.fill()

// 3. Test Radial Gradient
println("Testing Radial Gradient...")
var rad = canvas.new_radial_gradient(350, 60, 10, 350, 60, 50)
rad.add_color_stop(0, "#00ff00") // Green
rad.add_color_stop(1, "#000000") // Black
ctx.fillstyle(rad)
ctx.circle(350, 60, 50)
ctx.fill()

// 4. Test Conic Gradient
println("Testing Conic Gradient...")
var con = canvas.new_conic_gradient(500, 60, 0)
con.add_color_stop(0, "#00ffff") // Cyan
con.add_color_stop(1, "#ff00ff") // Magenta
ctx.fillstyle(con)
ctx.circle(500, 60, 50)
ctx.fill()


// 5. Test Line Caps and Joins
println("Testing Line Styles...")
ctx.strokestyle(solid)
ctx.linewidth(10)
ctx.linecap_round()
ctx.linejoin_round()
ctx.move_to(50, 150)
ctx.line_to(100, 200)
ctx.line_to(150, 150)
ctx.stroke()


// 7. Test Surface Pattern (requires loading the image we just saved or creating new one)
println("Testing Surface Pattern...")
var patCtx = canvas.new_context(20, 20)
patCtx.set_color("#666666")
patCtx.rect(0, 0, 20, 20)
patCtx.fill()
patCtx.set_color("#888888")
patCtx.circle(10, 10, 5)
patCtx.fill()


ctx.save_png("pattern.png")


println("Done.")
