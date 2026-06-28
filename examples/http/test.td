// tender@v0.0.5

import "http" 

var hit_count = 0

http.listen_and_serve(":8080", fn(req, res) {
    hit_count = hit_count + 1 // DATA RACE!
    res.write("Hits: " + string(hit_count))
})