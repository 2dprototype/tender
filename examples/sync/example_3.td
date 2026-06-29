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