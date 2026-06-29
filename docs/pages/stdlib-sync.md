## Stdlib `sync`

The `sync` module provides synchronization primitives for concurrent programming, allowing safe access to shared resources across multiple goroutines.

### Functions

#### `mutex()`

Creates a new mutex (mutual exclusion lock) for synchronizing access to shared resources.

Returns a `Mutex Object` with methods for locking and unlocking.

### Mutex Object Methods

#### `lock()`

Locks the mutex. If the mutex is already locked, the calling goroutine blocks until the mutex is available.

#### `unlock()`

Unlocks the mutex. It is a runtime error if the mutex is not locked on entry to unlock.

**Important**: A locked mutex is not associated with a particular goroutine. It is allowed for one goroutine to lock a mutex and for another goroutine to unlock it.

### Example Usage

```go
import "sync"

// Create a mutex
var mu = sync.mutex()

// Shared counter
var counter = 0

// Function to increment counter safely
fn increment() {
    mu.lock()
    counter = counter + 1
    mu.unlock()
}

// Start multiple goroutines
for i = 0; i < 10; i = i + 1 {
    go(fn() {
        increment()
    })
}

// Wait for all goroutines to finish (simplified)
sleep(1000)

println("Counter:", counter) // Should print 10
```

### Mutex with Shared Data Structure

```go
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
```

### Mutex with Error Handling

```go
import "sync"

var mu = sync.mutex()
var data = []

fn add_item(item) {
    mu.lock()
    
    // Critical section - ensure unlock even on error
    if len(data) > 100 {
        mu.unlock()
        return error("Data limit exceeded")
    }
    
    data = append(data, item)
    mu.unlock()
}

fn get_items() {
    mu.lock()
    result := data
    mu.unlock()
    return result
}

add_item(1)
add_item(2)
add_item(3)

get_items() |> println
```

### Important Notes

- Always pair `lock()` with `unlock()` to avoid deadlocks
- Mutexes are not reentrant - a goroutine cannot lock a mutex it already holds
- Use mutexes to protect shared state accessed by multiple goroutines
- For read-heavy workloads, consider using a read-write mutex (not currently implemented)