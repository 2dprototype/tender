## Stdlib `bufio`

The `bufio` module provides functions for buffered I/O operations, particularly useful for reading input from standard input (stdin).

### Functions

#### `readline()`

Reads a line of text from standard input until a newline character (`"\n"`) is encountered.

#### `readstring(delimiter)`

Reads a string from standard input until the specified delimiter character is encountered.

- `delimiter`: Delimiter character to stop reading the string.

#### `readbytes(num_bytes)`

Reads a specified number of bytes from standard input.

- `num_bytes`: Number of bytes to read.

### Example Usage

```go
import "bufio"

line := bufio.readline()
println("Line read:", line)

// Read a string until a specific delimiter character
delimiter := "\n"
string_input := bufio.readstring(delimiter)
println("String read until delimiter:", string_input)

// Read a specified number of bytes
num_bytes := 10
bytes_read := bufio.readbytes(num_bytes)
println("Bytes read:", bytes_read)
```
