// Create a 1-slot buffered channel
job_queue := makechan(1)

// Spawn background worker
worker := go(fn() {
    // Simulate some heavy math or database lookup
    secret_answer := 42
    
    job_queue.send(secret_answer)
})

// Main script thread waits for the channel payload
payload := job_queue.recv()

// Verify the background thread closed cleanly
worker.wait()