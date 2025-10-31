## Stdlib `http`

The `http` module provides functionalities for making HTTP requests. This module supports various HTTP methods including GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD, and TRACE.

### Functions

#### `get(url, [body], [headers])`

Sends an HTTP GET request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

#### `post(url, [body], [headers])`

Sends an HTTP POST request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

#### `put(url, [body], [headers])`

Sends an HTTP PUT request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

#### `delete(url, [body], [headers])`

Sends an HTTP DELETE request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

#### `patch(url, [body], [headers])`

Sends an HTTP PATCH request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

#### `options(url, [body], [headers])`

Sends an HTTP OPTIONS request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

#### `head(url, [body], [headers])`

Sends an HTTP HEAD request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

#### `trace(url, [body], [headers])`

Sends an HTTP TRACE request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

### Example Usage

```go
import "http"

// GET request
response := http.get("https://example.com")

if !is_error(response) {
    println(string(response))
} 
else {
    println("Response: ", response)
}
```
