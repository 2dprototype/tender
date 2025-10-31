## Stdlib `fmt`

The `fmt` module provides functions for formatted I/O operations similar to those in the standard Go `fmt` package. 

### Functions

#### `print(args...)`

Prints the arguments to standard output without a newline.

- `args`: The objects to print. Each object will be converted to a string.

Returns `null`.

#### `println(args...)`

Prints the arguments to standard output with a newline.

- `args`: The objects to print. Each object will be converted to a string.

Returns `null`.

#### `printf(format, args...)`

Prints the formatted string to standard output.

- `format`: A format string.
- `args`: The objects to format and print.

Returns `null`.

#### `sprintf(format, args...)`

Returns a formatted string.

- `format`: A format string.
- `args`: The objects to format.

Returns the formatted string as a `string`.

#### `scanln()`

Reads a line from standard input.

Returns the input as a `string`.

#### `fprint(IOWriter, args...)`

Prints the arguments to standard output without a newline using IOWriter.

- `IOWriter`: stdlib.IOWriter object.
- `args`: The objects to format.

#### `fprintln(IOWriter, args...)`

Prints the arguments to standard output with a newline using IOWriter.

- `IOWriter`: stdlib.IOWriter object.
- `args`: The objects to format.

Returns the formatted string as a `string`.

### Example Usage

```go
import "fmt"

// Print example
fmt.print("Hello, ", "world")

// Println example
fmt.println("Hello, ", "world")

// Printf example
fmt.printf("Hello, %s", "world")

// Sprintf example
formatted_string := fmt.sprintf("Hello, %s", "world")
println(formatted_string)

// Scanln example
input := fmt.scanln()
println("You entered: ", input)
```

#### `fprint, fprintln` Example

```go
import "fmt"
import "colors"
import "os"

fmt.fprint(colors.stdout, "Hello".red, "World".green, "\n")
fmt.fprintln(colors.stderr, "Hello".red, "World".green)
fmt.fprintln(os.stdout, "Hello".red, "World".green, "\n")
```