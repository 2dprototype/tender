// tender@v0.0.5

sync := import("sync")
http := import("http")

// Global shared state
mu := sync.mutex()
var total_requests = 0

http.listen_and_serve(":8080", fn(req, res) {
    // Lock memory, safely mutate, unlock
    mu.lock()
    total_requests = total_requests + 1
    current_visitor := total_requests
    mu.unlock()

    // Router logic
    if req.path == "/" {
        res.set_status(200)
        res.set_header("Content-Type", "text/plain")
        res.write("Welcome! You are visitor #" + string(current_visitor))
    } else {
        res.set_status(404)
        res.write("Page not found on server.")
    }
})