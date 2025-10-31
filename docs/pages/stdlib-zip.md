## Stdlib `zip`

The `zip` module provides functions for creating and reading ZIP archives.

### Functions

- `writer()`: Returns a writer object for creating a new ZIP archive. The writer object provides methods for adding files, setting comment, closing, and flushing the ZIP archive.
  - `create(filename, content)`: Creates a new file entry in the ZIP archive with the specified filename and content.
  - `bytes()`: Returns the byte representation of the created ZIP archive.
  - `close()`: Closes the ZIP archive writer.
  - `flush()`: Flushes any buffered data to the underlying writer.
  - `set_comment(comment)`: Sets the comment for the ZIP archive.
  - `set_offset(offset)`: Sets the offset for the next file entry in the ZIP archive.

- `reader(data)`: Returns a reader object for reading a ZIP archive from the provided byte slice.
  - `files`: Returns an array containing information about each file in the ZIP archive.
    - `name`: Name of the file.
    - `comment`: Comment associated with the file.
    - `non_utf8`: Indicates whether the file name is not encoded in UTF-8.
    - `creator_version`: Version of the software that created the ZIP archive.
    - `reader_version`: Version needed to read the ZIP archive.
    - `method`: Compression method used for the file.
    - `modified`: Last modification time of the file.
    - `modified_time`: Modified time in seconds since January 1, 1970 UTC.
    - `modified_date`: Modified date in MS-DOS date format.
    - `crc32`: CRC-32 checksum of the file contents.
    - `compressed_size`: Size of the compressed file data.
    - `uncompressed_size`: Size of the uncompressed file data.
    - `extra`: Extra data associated with the file.
    - `read()`: Function to read the content of the file.

  - `comment`: Comment associated with the ZIP archive.

### Example Usage

```go
import "zip"
import "io"

// Create a new ZIP archive
writer := zip.writer()
writer.create("hello.txt", "hello world")
writer.set_comment("this is a comment!")
writer.close()

io.writefile("hello.zip", writer.bytes())

// Reading ZIP
reader := zip.reader(io.readfile("hello.zip"))

for file in reader.files {
	println(file.read())
}
```