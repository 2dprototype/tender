// Int comparisons
println(5 <=> 5)     // Expected: 0
println(3 <=> 5)     // Expected: -1
println(7 <=> 5)     // Expected: 1

// Mixed Int and Float comparisons
println(5 <=> 5.0)   // Expected: 0
println(3.5 <=> 5)   // Expected: -1
println(7 <=> 5.5)   // Expected: 1

// Float comparisons
println(5.5 <=> 5.5) // Expected: 0
println(3.5 <=> 5.5) // Expected: -1
println(7.5 <=> 5.5) // Expected: 1

// Char comparisons
println('a' <=> 'a') // Expected: 0
println('a' <=> 'b') // Expected: -1
println('b' <=> 'a') // Expected: 1

// String comparisons
println("abc" <=> "abc") // Expected: 0
println("abc" <=> "def") // Expected: -1
println("def" <=> "abc") // Expected: 1
