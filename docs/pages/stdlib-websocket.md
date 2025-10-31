## Stdlib `websocket`

The `websocket` module provides functionalities for establishing and managing WebSocket connections using the `gorilla/websocket` library.

### Functions

#### `dial(url)`

Establishes a WebSocket connection to the specified URL.

- `url`: The URL to connect to.

Returns a `WebSocket Connection Object` with methods for interaction.

### WebSocket Connection Object Methods

#### `read_message()`

Reads a message from the WebSocket connection.

Returns an array containing the message type and message data.

#### `write_message(type, message)`

Writes a message to the WebSocket connection.

- `type`: The type of the message (e.g., text, binary).
- `message`: The message data to be sent.

#### `close()`

Closes the WebSocket connection.

#### `local_addr()`

Returns the local address of the WebSocket connection.

#### `remote_addr()`

Returns the remote address of the WebSocket connection.

#### `set_readdeadline(t)`

Sets the read deadline for the WebSocket connection.

- `t`: The read deadline time.

#### `set_writedeadline(t)`

Sets the write deadline for the WebSocket connection.

- `t`: The write deadline time.

### Example Usage

```go
import "websocket"

// Dial a WebSocket connection
var conn = websocket.dial("ws://example.com/socket")

//Read message from server
go(fn(){
	for {
		read := conn.read_message()
		println(read)
	}
})

//Write message
conn.write_message(1, "hello")

```
