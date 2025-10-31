## Stdlib `io`

The `io` module provides functions for reading from and writing to files.

### Functions

- `readfile(path)`: Reads the contents of the file specified by the `path` parameter and returns it as a string.

- `writefile(path, content[, mode])`: Writes the `content` to the file specified by the `path` parameter. The optional `mode` parameter specifies the file mode (permission and mode bits). If not provided, the default mode is `0644`.

### Example Usage

```go
import "io"

// Read contents of a file
content := io.readfile("/path/to/file.txt")

// Write contents to a file
io.writefile("/path/to/file.txt", "Hello, World!")
```
