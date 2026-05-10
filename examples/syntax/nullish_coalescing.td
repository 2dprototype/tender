// tender@v0.0.4


println("\nNullish coalescing (??):")
var a = null
var b = "default"
println("null ?? 'default':", a ?? b)
a = "value"
println("'value' ?? 'default':", a ?? b)

// Special cases for ??
println("false ?? 'default':", false ?? "default")
println("0 ?? 'default':", 0 ?? "default")
println("'' ?? 'default':", "" ?? "default")
