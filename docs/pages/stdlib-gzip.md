## Stdlib `gzip`

The `gzip` module provides functions for compressing and decompressing data using the gzip compression format.

### Functions

- `compress(data)`: Compresses the input data using gzip compression and returns the compressed data as a byte slice.
- `decompress(data)`: Decompresses the input data using gzip decompression and returns the decompressed data as a byte slice.

### Example Usage

```go
import "gzip"

// Compress data
compressed_data := gzip.compress("Hello, World!")

// Decompress data
decompressed_data := gzip.decompress(compressed_data)
```