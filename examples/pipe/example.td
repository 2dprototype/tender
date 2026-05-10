// tender@v0.0.4

// Test Pipe Operator
add := fn(a, b) {
    return a + b 
}

5 |> add(10) |> println // Should print 15

// Test Pipe with complex expressions
[1, 2, 3] |> len |> println // Should print 3

// Test Backward Pipe
println <| add(20, 30) // Should print 50
