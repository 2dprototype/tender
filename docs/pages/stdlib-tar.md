## Stdlib `tar`

The `tar` module provides functions for creating and reading TAR archives.

### Functions

- `writer()`: Returns a writer object for creating a new TAR archive. The writer object provides methods for adding files and closing the TAR archive.
  - `create(filename, content)`: Creates a new file entry in the TAR archive with the specified filename and content.
  - `bytes()`: Returns the byte representation of the created TAR archive.
  - `close()`: Closes the TAR archive writer.

- `reader(data)`: Returns an array containing information about each file in the TAR archive.
  - Each element in the array represents a file and contains the following attributes:
    - `name`: Name of the file.
    - `mode`: Permission mode of the file.
    - `size`: Size of the file.
    - `data`: Content of the file.

### Example Usage

```go
import "tar"
import "io"

// Create a new TAR archive
writer := tar.writer()
writer.create("hello.txt", "hello world")
writer.close()

io.writefile("hello.tar", writer.bytes())

// Reading TAR
reader := tar.reader(io.readfile("hello.tar"))
```