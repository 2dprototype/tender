// tender@v0.0.5

import "http"

http.listen_and_serve(":8080", fn(req, res) {
    if req.path == "/health" {
        res.set_status(200)
        res.write("OK")
    } 
	else {
        res.set_status(404)
        res.set_header("Content-Type", "text/plain")
        res.write("Not Found")
    }
})