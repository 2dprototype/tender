# Tender

**Tender** is a general-purpose programming language specially designed for image processing, 2D graphics, scripting, and more! Here is a quick [tutorial](docs/pages/tutorial.md). Also check the [docs](https://2dprototype.github.io/tender)!

## Overview

Tender compiles into bytecode and executes on a stack-based virtual machine (VM) written in native Golang.

## Features

- **Simple and highly readable syntax**  
- **Compiles to bytecode**  
- **Supports rich [built-in functions](docs/pages/builtins.md)**  
- **Includes an extensive [standard library](docs/pages/stdlib.md)**  
- **Optimized for 2D graphics**  
- **REPL (Read-Eval-Print Loop) for interactive development**
- **Rich type system** including int, float, string, bool, char, null, big integers, big floats, complex numbers, bytes, arrays (dynamic and immutable), maps (dynamic and immutable), tuples, time values, and error values
- **User-defined structs** with field types, nested structs, anonymous structs, and embedded fields
- **Pointers and references** for mutable data manipulation
- **Closures and first-class functions**
- **Advanced operators** including pipe operators (`<|`, `|>`), null coalescing (`??`), optional chaining (`?.`), ternary conditional (`? :`), compound assignment operators, and logical operators (`&&`, `||`)
- **Modular architecture** with import statements, module aliasing, selective imports, embedded file import (`embed()`), and file-based module loading
- **Runtime type introspection** with `typeof()` and type checking functions
- **Error handling** through the `error()` expression
- **Immutable data structures** via `immutable()` expression
- **Loop control** with `break` and `continue` statements
- **For loops** including traditional, for-in, conditional, and infinite loops
- **Variable declarations** with `var` and constants with `const`
- **Function definitions** with `fn` keyword
- **Export statements** for module exports
- **Built-in functions** for type conversion, type checking, collection manipulation, search operations, memory operations (pointer, deref, set), range generation, debugging, and pretty printing
- **Bytecode compilation** with compilation, execution, and parse-only modes
- **Comprehensive operator precedence** matching conventional expectations
- **Gob and CSV encoding/decoding** support
- **Cross-platform support** for Windows, macOS, and Linux

### Supported Standard Library

- [math](docs/pages/stdlib-math.md): Mathematical constants and functions  
- [cmplx](pages/stdlib-cmplx.md): Functions for complex numbers
- [os](docs/pages/stdlib-os.md): Platform-independent interface to OS functionality  
- [strings](docs/pages/stdlib-strings.md): String conversion, manipulation, and regular expressions  
- [times](docs/pages/stdlib-times.md): Time-related functions  
- [rand](docs/pages/stdlib-rand.md): Random number generation  
- [fmt](docs/pages/stdlib-fmt.md): Formatting functions  
- [json](docs/pages/stdlib-json.md): JSON handling functions  
- [xml](docs/pages/stdlib-xml.md): XML handling functions  
- [base64](docs/pages/stdlib-base64.md): Base64 encoding and decoding  
- [hex](docs/pages/stdlib-hex.md): Hexadecimal encoding and decoding  
- [colors](docs/pages/stdlib-colors.md): Functions to print colored text to the terminal  
- [gzip](docs/pages/stdlib-gzip.md): Gzip compression and decompression  
- [zip](docs/pages/stdlib-zip.md): ZIP archive manipulation  
- [tar](docs/pages/stdlib-tar.md): TAR archive creation and reading  
- [bufio](docs/pages/stdlib-bufio.md): Buffered I/O functions  
- [crypto](docs/pages/stdlib-crypto.md): Cryptographic functions  
- [path](docs/pages/stdlib-path.md): File path manipulation  
- [image](docs/pages/stdlib-image.md): Image manipulation  
- [canvas](docs/pages/stdlib-canvas.md): Drawing functions for canvases  
- [dll](docs/pages/stdlib-dll.md): Dynamic link library interactions  
- [io](docs/pages/stdlib-io.md): Input and output functions  
- [audio](docs/pages/stdlib-audio.md): Audio processing  
- [net](docs/pages/stdlib-net.md): Networking functions  
- [http](docs/pages/stdlib-http.md): HTTP client and server utilities  
- [websocket](docs/pages/stdlib-websocket.md): WebSocket communication utilities  
- **gob**: Gob Encoding/Decoding
- **csv**: CSV Encoding/Decoding
- **wui**: Windows GUI

## Quick Start

1. **Install Tender on your machine.**  
2. **Copy the sample code below:**

```go
// Basic example
str1 := "hello"
str2 := "world"

println(str1 + " " + str2)
```

```go
// Structs example
type user struct {
    name string
    age  int
}

u := user{name: "Alice", age: 30}
println("Name:", u.name, "Age:", u.age)

// Nested structs
type point struct {
    x, y int
}

type line struct {
    p1 point
    p2 point
}

l := line{
    p1: point{x: 0, y: 0},
    p2: point{x: 10, y: 10},
}
println("line from (", l.p1.x, ",", l.p1.y, ") to (", l.p2.x, ",", l.p2.y, ")")
```

```go
// Canvas drawing example (similar to JS Canvas)
import "canvas"
	
var ctx = canvas.new_context(100, 100)
ctx.hex("#0f0")          // Set color to green
ctx.dash(4, 2)           // Define dashed stroke
ctx.rect(25, 25, 50, 50) // Draw a rectangle
ctx.stroke()

ctx.save_png("out.png")  // Save output as PNG
```

3. **Save your code as `hello.td`** (use the `.td` extension).  
4. **Run your script using the following command:**

```bash
tender hello.td
```

---

## Installation

### Using Go

1. Install the latest version of Go.  
2. Run the following command to install:

```bash
go install github.com/2dprototype/tender/cli/tender@latest
```

### Manual Installation (Windows)

Precompiled binaries are available. Download them from the release tags.

---

## Documentation
Check the [docs](https://2dprototype.github.io/tender)!

- **[Runtime Types](docs/pages/runtime-types.md)**  
- **[Built-in Functions](docs/pages/builtins.md)**  
- **[Operators](docs/pages/operators.md)**  
- **[Standard Library](docs/pages/stdlib.md)**  

## Examples

### Basic Examples
```go
// Variable declarations
var name = "Tender"
const PI = 3.14159

// Functions
fn add(a, b) {
    return a + b
}

// Closures
fn make_counter() {
    var count = 0
    return fn() {
        count++
        return count
    }
}

// Arrays and maps
var arr = [1, 2, 3, 4, 5]
var map = { "key": "value" }

// Structs
type Person struct {
    name string
    age  int
}

var person = Person{name: "John", age: 25}
person.age = 26

// Pointers
var p = pointer(person)
var val = deref(p)
set(p, Person{name: "Jane", age: 30})

// Type conversion and checking
var num = int("123")
if is_string(num) {
    println("This is a string")
} else {
    println("This is not a string")
}

// Error handling
var result = error("something went wrong")
if is_error(result) {
    println(result.value)
}
```

### Advanced Examples
```go
// Pipe operators for functional composition
var result = [1, 2, 3, 4, 6] |> sort |> reverse |> println

// Null coalescing
var value = null ?? "default value"

// Optional chaining
var user = {	
    profile: {
        name: "jack"
    }
}
var name = user?.profile?.name
sysout name, "\n"

// Range generation
var numbers = range(0, 10, 2)  // [0, 2, 4, 6, 8]
sysout numbers, "\n"

// Module imports
import "math" as m
var sqrt2 = m.sqrt(2)
println(sqrt2)
```

Explore various examples demonstrating Tender's features in the [examples](examples) directory.

---

## Command Line Usage

Tender supports multiple operation modes:

```bash
# Start REPL (interactive mode)
tender

# Compile and run a source file
tender myapp.td

# Compile to bytecode
tender -o myapp myapp.td

# Run compiled bytecode
tender myapp

# Parse and output AST as JSON
tender -parse ast.json myapp.td

# Show version
tender -version
# or
tender -v

# Show help
tender -help
```

---

## Type System Overview

Tender provides a rich type system with support for:

| Type | Description | Example |
|------|-------------|---------|
| `int` | 64-bit integer | `42` |
| `float` | 64-bit floating point | `3.14159` |
| `bigint` | Arbitrary-precision integer | `12345678901234567890` |
| `bigfloat` | Arbitrary-precision float | `3.14159265358979323846` |
| `complex` | Complex number | `3+4i` |
| `string` | UTF-8 string | `"hello"` |
| `bool` | Boolean | `true` or `false` |
| `char` | Unicode character | `'a'` |
| `bytes` | Byte array | `[72, 101, 108, 108, 111]` |
| `array` | Dynamic array | `[1, 2, 3]` |
| `immutable-array` | Immutable array | `[1, 2, 3]` |
| `map` | Dynamic map | `{"key": value}` |
| `immutable-map` | Immutable map | `{"key": value}` |
| `tuple` | Fixed-size immutable sequence | `(1, "hello", true)` |
| `struct` | User-defined structure | `user{name: "Alice", age: 30}` |
| `time` | Time value | `time()` |
| `error` | Error value | `error("message")` |
| `pointer` | Reference to a value | `pointer(x)` |
| `null` | Null value | `null` |

---

## Dependencies

Tender uses the following dependencies:

- [go-mp3](https://github.com/hajimehoshi/go-mp3)  
- [gorilla/websocket](https://github.com/gorilla/websocket)  
- [ebitengine/oto/v3](https://github.com/ebitengine/oto/v3)  
- [exp/shiny](https://pkg.go.dev/golang.org/x/exp/shiny)  
- [fogleman/gg](https://github.com/fogleman/gg)  

---

## Syntax Highlighting

Syntax highlighting is currently available for:
- **Notepad++**: Download the configuration file [here](misc/syntax/npp_tender.xml)
- Support for additional editors coming soon

---

## License

Tender is distributed under the [MIT License](LICENSE), with additional licenses provided for third-party dependencies. See [LICENSE_GOLANG](LICENSE_GOLANG) and [LICENSE_TENGO](LICENSE_TENGO) for more information.

---

## Acknowledgments

Tender is written in Go, based on Tengo. We extend our gratitude to the contributors of Tengo for their valuable work.