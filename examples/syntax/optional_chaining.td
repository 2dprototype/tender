// tender@v0.0.4

println("\nOptional chaining (selector):")
var obj = {prop: "hello"}
println("obj?.prop:", obj?.prop)
obj = null
println("null?.prop:", obj?.prop)

println("\nOptional chaining (index):")
var arr = [1, 2, 3]
println("arr?.[0]:", arr?.[0])
arr = null
println("null?.[0]:", arr?.[0])

println("\nOptional chaining (call):")
greet := fn(name) { return "Hello, " + name }
println("greet?.(\"Alice\"): ", greet?.("Alice"))
greet = null
println("null?.(\"Alice\"): ", greet?.("Alice"))