# Getting Started  

Welcome to **Tender**, a flexible programming language designed for simplicity and power. Let’s dive in!

---

## **1. Variables in Tender**  

In Tender, variables are defined using the `:=` operator along with `var`. You don’t need to declare the type explicitly—Tender infers the type based on the assigned value.

### **Example: Variable Declaration**  
```go
name := "Alice"    // String variable  
age := 30          // Integer variable  
pi := 3.14         // Floating-point variable  
is_admin := false  // Boolean variable
var year = bigint(2024) //bigint datatype
var c = complex(pi, age) //complex datatype
```



### **Reassigning Variables**  
Variables in Tender are **dynamically typed**, meaning you can assign a value of a different type to an existing variable.

```go
age = "Thirty"  // Now 'age' holds a String instead of an Integer
println(age)    // Output: Thirty
```

### **Constants**
Constants are variables that cannot be reassigned after their initial declaration. They are defined using the `const` keyword.

```go
const pi = 3.14159
pi = 3.14 // Compile Error: cannot assign to constant 'pi'
```

---

## **2. Data Types**  

Tender supports several primitive data types:

| **Type**     | **Example**              | **Description**                     |
|--------------|--------------------------|-------------------------------------|
| string       | `"Hello"`                | Sequence of characters.             |
| int          | `42`                     | Whole numbers.                      |
| float        | `-19.84`                 | Numbers with decimal points.        |
| bool         | `true` / `false`         | Logical values.                     |
| char         | `'A'`                    | Single character.                   |
| time         | `time()`                 | time (`time.Time` in Go)            |
| bigint       | `bigint(42)`             | Arbitrary-precision integer.        |
| bigfloat     | `bigfloat(1000.11)`      | Arbitrary-precision float.          |
| complex      | `complex(1, 2)`          | Complex number (`complex128` in go).|

---

## **3. Arrays and Maps**  

### **Arrays**  
Arrays hold an ordered list of values. They can contain elements of any type.  
```go
fruits := ["Apple", "Banana", "Cherry"]  
println(fruits[0])  // Output: Apple
```

### **Maps**  
Maps are key-value pairs, similar to dictionaries in other languages.  
```go
person := {name: "Alice", age: 30, is_admin: true}  
println(person["name"])  // Output: Alice
```

---

## **4. Control Flow Statements**  

### **If/Else Statements**  
Tender supports `if/else` for decision-making. You can include an **init statement** before the condition.  
```go
if age := 30; age > 18 {
    println("Adult")
}
else {
    println("Minor")
}
```

---

## **5. Loops**  

### **For Loop**  
The `for` loop is used to iterate over arrays or perform repeated actions.  
```go
numbers := [1, 2, 3, 4]
for i := 0; i < len(numbers); i++ {
    println(numbers[i])
}
```

### **For-Each Loop**  
```go
var fruits = ["mango", "apple", "banana"]

for i, fruit in fruits {
	if fruit == apple {
		continue
	}
	else if fruit i == fruits.length {
		break
	}
    println(fruit)
}
```

### **For-condition**  

```go
for true {
    println(fruit)
}
```

### **For-infinity**  

```go
for {}
```

---

## **6. Functions in Tender**  

Functions are defined using the `fn` keyword. Functions can take arguments and return values.

### **Defining a Function**  
```go
greet := fn(name) {
    return "Hello, " + name
}
println(greet("Alice"))  // Output: Hello, Alice
```

*or*

```go
fn greet(name) {
    return "Hello, " + name
}
println(greet("Alice"))  // Output: Hello, Alice
```

### **Passing Functions as Arguments**  
Tender supports higher-order functions—functions that take other functions as parameters.  
```go
each := fn(arr, action) {
    for x in arr {
        action(x)
    }
}
each([1, 2, 3], fn(n) {
    println(n)
})
```

---

## **7. Closures**  

Closures are functions that capture variables from their outer scope.  
```go
counter := fn() {
    count := 0
    return fn() {
        count += 1
        return count
    }
}
next := counter()
println(next())  // Output: 1
println(next())  // Output: 2
```

---

## **8. Recursion and Tail-Call Optimization**  

Tender supports recursion and optimizes tail-recursive calls for better performance.  
```go
factorial := fn(n, acc) {
    if n == 0 {
        return acc
    }
    return factorial(n - 1, acc * n)
}
println(factorial(5, 1))  // Output: 120
```

---

## **9. Slicing Strings and Arrays**  

You can extract parts of strings or arrays using **slices**.  
```go
str := "hello world"
println(str[0:5])  // Output: hello

arr := [10, 20, 30, 40, 50]
println(arr[1:4])  // Output: [20, 30, 40]
```

---

## **10. Type Conversion**  

Tender provides built-in functions to convert between data types.  
```go
str_num := string(42)  // "42"
int_num := int("99")   // 99
float_num := float(3)  // 3.0
bool_val := bool(1)    // true
```

## **12. Built-in Functions**  

| **Function**   | **Description**                           |
|----------------|-------------------------------------------|
| `println()`    | Prints output to the console.             |
| `len()`        | Returns the length of an array or string. |
| `append()`     | Adds elements to an array.                |
|more...||

See all [builtin functions](builtins.md)!

---

## **13. Tuples**

Tuples are immutable ordered collections of objects, similar to arrays but enclosed in parentheses `()`.

```go
t := (1, 2, "three")
println(t[0])      // Output: 1
println(t.length)  // Output: 3

// Single-element tuple requires a trailing comma
t1 := (10,) 
```

---

## **14. Pipe Operator**

The pipe operators `|>` and `<|` allow for functional-style chaining of expressions.

```go
add := fn(a, b) { return a + b }
double := fn(x) { return x * 2 }

// Forward pipe: x |> f  => f(x)
res := 5 |> double |> double
println(res) // Output: 20

// Pipe into function call: x |> f(y) => f(x, y)
res2 := 10 |> add(20)
println(res2) // Output: 30

// Backward pipe: f <| x => f(x)
println <| add(1, 2) // Output: 3
```

---

## **15. Embedding Files**

The `embed` expression allows you to include the contents of a file as a `Bytes` object directly in your script at compile-time.

```go
// Embed a text file or binary data
logo := embed("logo.png")
config := embed("config.json")

println(is_bytes(logo)) // Output: true
```

---
