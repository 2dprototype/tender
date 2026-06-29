import "sync"
import { sleep } from "times"

// Protect access to a shared map
var mu = sync.mutex()
var cache = {}

fn set_(key, value) {
    mu.lock()
    cache[key] = value
    mu.unlock()
}

fn get_(key) {
    mu.lock()
    value := cache[key]
    mu.unlock()
    return value
}

// Use in concurrent environment
go(fn() {
    set_("user1", "Alice")
})

go(fn() {
    set_("user2", "Bob")
})

sleep(100)
println(get_("user1")) // Alice
println(get_("user2")) // Bob