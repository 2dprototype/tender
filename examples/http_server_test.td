import "http"
import "fmt"
import "times"
import "websocket"

fn handler(res, req) {
	// Debug print
	fmt.println("Server received request:", req.method, req.url())

	if req.url() == "/ws" {
		fmt.println("Upgrading to WebSocket...")
		conn := websocket.upgrade(res, req)
		fmt.println("WebSocket upgraded.")
		
		// Send a welcome message
		conn.write_message(1, bytes("hello via ws"))

		// Echo back
		for {
			msg := conn.read_message()
			if is_error(msg) { break } // Error 
			conn.write_message(1, msg[1])
		}
		conn.close()
		return
	}

	res.write(bytes("Hello from Tender server!"))
}

// Start server in a separate thread
go(fn() {
	fmt.println("Starting server on :8081...")
	err := http.listen_and_serve(":8081", handler)
	if err {
		fmt.println("Server error:", err)
	}
})

// Give server time to start
times.sleep(times.second)

// Client test: HTTP GET
fmt.println("Testing HTTP GET...")
resp := http.get("http://localhost:8081/")
if is_error(resp) {
    fmt.println("HTTP GET Error:", resp)
} else {
    fmt.println("Response status:", resp.status)
    fmt.println("Response body:", string(resp.body()))

    if resp.status != 200 {
        fmt.println("HTTP Test Failed!")
    } else {
        fmt.println("HTTP Test Passed!")
    }
}

// Client test: WebSocket
fmt.println("Testing WebSocket...")
ws := websocket.dial("ws://localhost:8081/ws")
if is_error(ws) {
    fmt.println("WebSocket Dial Error:", ws)
} else {
    msg := ws.read_message()
    fmt.println("WS Message 1:", string(msg[1]))

    // Send echo
    ws.write_message(1, bytes("echo me"))
    msg2 := ws.read_message()
    fmt.println("WS Message 2:", string(msg2[1]))

    ws.close()
    fmt.println("WebSocket Test Finished.")
}
