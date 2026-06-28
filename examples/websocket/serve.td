// tender@v0.0.5

websocket := import("websocket")

websocket.listen_and_serve(":8080", fn(conn) {
    // Keep this specific client's thread alive in a listening loop
    for {
        msg := conn.read_message() 
        
        if is_error(msg) {
            break 
        }

        msg_type := msg[0]
        raw_bytes := msg[1]

        // Echo the exact message back to the browser
        conn.write_message(msg_type, raw_bytes)
    }
    
    conn.close()
})