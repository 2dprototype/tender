import "http"
import "json"

async_fetch := go(fn() {
    req := http.get("https://api.github.com")
    resp := req.execute()
    
    return resp.body
})

// Do other work here while the HTTP request flies across the internet...

// Pull the returned value out of the goroutine VM
final_status := async_fetch.result()

debug(final_status |> json.decode)