net := import("net")

server := net.listen("tcp", "127.0.0.1:9000")

for {
    // Blocks until a TCP client connects (e.g. via Telnet or Netcat)
    client_conn := server.accept()

    // Immediately spin off a background worker so the server can accept the next guy!
    go(fn() {
        inbound_data := client_conn.read(512)
        
        client_conn.write(bytes("ACK: "))
        client_conn.write(inbound_data)
        client_conn.close()
    })
}