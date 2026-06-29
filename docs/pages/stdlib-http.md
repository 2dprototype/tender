## Stdlib `http`

The `http` module provides functionalities for making HTTP requests and creating HTTP servers. This module supports various HTTP methods including GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD, and TRACE.

### Functions

#### `get(url, [body], [headers])`

Sends an HTTP GET request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

Returns an HTTP request object that can be used to execute the request.

#### `post(url, [body], [headers])`

Sends an HTTP POST request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

Returns an HTTP request object that can be used to execute the request.

#### `put(url, [body], [headers])`

Sends an HTTP PUT request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

Returns an HTTP request object that can be used to execute the request.

#### `delete(url, [body], [headers])`

Sends an HTTP DELETE request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

Returns an HTTP request object that can be used to execute the request.

#### `patch(url, [body], [headers])`

Sends an HTTP PATCH request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

Returns an HTTP request object that can be used to execute the request.

#### `options(url, [body], [headers])`

Sends an HTTP OPTIONS request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

Returns an HTTP request object that can be used to execute the request.

#### `head(url, [body], [headers])`

Sends an HTTP HEAD request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

Returns an HTTP request object that can be used to execute the request.

#### `trace(url, [body], [headers])`

Sends an HTTP TRACE request to the specified URL.

- `url`: The URL to which the request is sent.
- `body` (optional): Request body as a byte array.
- `headers` (optional): Map of request headers.

Returns an HTTP request object that can be used to execute the request.

#### `listen_and_serve(addr, handler)`

Starts an HTTP server on the specified address.

- `addr`: The address to listen on (e.g., ":8080").
- `handler`: A callback function that receives the request and response objects.

Returns an error object if the server fails to start.

### HTTP Request Object Methods

#### `body()`

Executes the HTTP request and returns the response body as bytes.

#### `execute()`

Executes the HTTP request and returns a full response object.

#### `close()`

Returns whether the connection should be closed after the request.

#### `method()`

Gets the HTTP method of the request.

#### `url()`

Gets the URL of the request.

#### `headers`

Returns a map of request headers.

#### `get_header(key)`

Gets the value of a specific header.

- `key`: The header key.

Returns the header value as a string.

#### `set_header(key, value)`

Sets a header value.

- `key`: The header key.
- `value`: The header value.

#### `set_body(data)`

Sets the request body.

- `data`: The body data as bytes.

#### `set_method(method)`

Changes the HTTP method of the request.

- `method`: The new HTTP method.

#### `set_url(url)`

Changes the URL of the request.

- `url`: The new URL.

### HTTP Response Object Properties

- `status`: The response status code.
- `status_text`: The response status text.
- `headers`: A map of response headers.
- `body`: The response body as bytes.
- `content_type`: The response content type.

### HTTP Request Object (Client) Example

```go
import "http"

// Create and execute a GET request
req := http.get("https://api.example.com/data")
resp := req.execute()
if !is_error(resp) {
    println("Status:", resp.status)
    println("Body:", string(resp.body))
}

// Create a request with custom headers
headers := {"Authorization": "Bearer token123"}
req = http.get("https://api.example.com/protected", null, headers)
body := req.body()
println(string(body))

// POST request with JSON body
json_data := bytes(`{"name": "John", "age": 30}`)
req = http.post("https://api.example.com/users", json_data)
req.set_header("Content-Type", "application/json")
resp = req.execute()
println("Created user:", string(resp))
```

### HTTP Response Object (Server) Example

```go
import "http"

// Create an HTTP server
http.listen_and_serve(":8080", fn(req, res) {
    println("Request from:", req.remote_addr)
    println("Method:", req.method)
    println("Path:", req.path)
    
    // Set response headers
    res.set_header("Content-Type", "text/plain")
    
    // Write response
    res.set_status(200)
    res.write("Hello, World!")
    
    // Access request body
    if len(req.body) > 0 {
        println("Body:", string(req.body))
    }
})

// Server with routing logic
http.listen_and_serve(":8080", fn(req, res) {
    if req.path == "/" {
        res.set_header("Content-Type", "text/html")
        res.write("<h1>Home Page</h1>")
    } 
    else if req.path == "/api/data" && req.method == "GET" {
        res.set_header("Content-Type", "application/json")
        res.write(`{"status": "ok", "data": [1, 2, 3]}`)
    } 
    else {
        res.set_status(404)
        res.write("Not Found")
    }
})
```

### Combined Example

```go
import "http"

// Server that handles different methods
http.listen_and_serve(":8080", fn(req, res) {
    if req.path == "/" {
        res.set_header("Content-Type", "text/html")
        res.write("<h1>Welcome!</h1><p>Try /api/hello</p>")
    } 
    else if req.path == "/api/hello" {
        if req.method == "GET" {
            res.set_header("Content-Type", "application/json")
            res.write(`{"message": "Hello, GET!"}`)
        }
        else if req.method == "POST" {
            res.set_header("Content-Type", "application/json")
            res.write(`{"message": "Hello, POST!", "received": ${string(req.body)}}`)
        }
        else {
            res.set_status(405)
            res.write("Method Not Allowed")
        }
    } 
    else {
        res.set_status(404)
        res.write("Page not found")
    }
})
```