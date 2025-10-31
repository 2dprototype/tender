## Stdlib `net`

The `net` module provides functionalities for network communication, including DNS lookup, TCP and UDP address resolution, and various methods for establishing network connections.

### Functions

#### `dnslookup(host)`

Performs a DNS lookup for the given host.

- `host`: The hostname to lookup.

Returns an array of IP addresses associated with the given host.

#### `resolve_tcp_addr(network, address)`

Resolves a TCP address.

- `network`: Network type (e.g., "tcp", "tcp4", "tcp6").
- `address`: The address to resolve.

Returns a string representation of the resolved TCP address.

#### `resolve_udp_addr(network, address)`

Resolves a UDP address.

- `network`: Network type (e.g., "udp", "udp4", "udp6").
- `address`: The address to resolve.

Returns a string representation of the resolved UDP address.

#### `dial(network, address)`

Establishes a generic network connection.

- `network`: Network type (e.g., "tcp", "udp").
- `address`: The address to connect to.

Returns a connection object with methods for interaction.

#### `dialtcp(network, address)`

Establishes a TCP network connection.

- `network`: Network type (e.g., "tcp", "tcp4", "tcp6").
- `address`: The address to connect to.

Returns a connection object with methods for interaction.

### Connection Object Methods

#### `close()`

Closes the network connection.

#### `read(buffer)`

Reads data from the connection into the buffer.

- `buffer`: The buffer to read data into.

Returns the number of bytes read.

#### `write(buffer)`

Writes data from the buffer to the connection.

- `buffer`: The buffer containing data to write.

Returns the number of bytes written.

#### `local_addr()`

Returns the local address of the connection.

#### `remote_addr()`

Returns the remote address of the connection.

#### `set_deadline(t)`

Sets the read and write deadlines for the connection.

- `t`: The deadline time.

#### `set_readdeadline(t)`

Sets the read deadline for the connection.

- `t`: The read deadline time.

#### `set_writedeadline(t)`

Sets the write deadline for the connection.

- `t`: The write deadline time.

### Example Usage

```go
import "net"

// DNS Lookup
addresses := net.dnslookup("example.com")
println(addresses)

// Resolve TCP Address
tcp_addr := net.resolve_tcp_addr("tcp", "example.com:80")
println(tcp_addr)

// Resolve UDP Address
udp_addr := net.resolve_udp_addr("udp", "example.com:80")
println(udp_addr)

// Dial a TCP Connection
conn := net.dialtcp("tcp", "example.com:80")
if !is_error(conn) {
    println("Connected to: ", conn.remote_addr())

    // Write to the connection
    msg := bytes("GET / HTTP/1.1\r\nHost: example.com\r\n\r\n")
    n := conn.write(msg)
    println("Bytes written: ", n)

    // Read from the connection
    buf := bytes(1024)
    n = conn.read(buf)
    println("Response: ", string(buf[:n]))

    // Close the connection
    conn.close()
}
```
