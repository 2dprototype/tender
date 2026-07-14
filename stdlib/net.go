package stdlib

import (
	"net"

	"github.com/2dprototype/tender"
)

var netModule = map[string]tender.Object{
	"dnslookup":        &tender.NativeFunction{Value: netDnsLookup},
	"resolve_tcp_addr": &tender.NativeFunction{Value: netResolveTCPAddr},
	"resolve_udp_addr": &tender.NativeFunction{Value: netResolveUDPAddr},
	"dial":             &tender.NativeFunction{Value: netDial},
	"dialtcp":          &tender.NativeFunction{Value: netDialTCP},
	"listen":           &tender.NativeFunction{Value: netListen},
}


func netDialTCP(args ...tender.Object) (tender.Object, error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}
	network, _ := args[0].(*tender.String)
	address, _ := args[1].(*tender.String)
	tcpAddr, err := net.ResolveTCPAddr(network.Value, address.Value)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP(network.Value, nil, tcpAddr)
	if err != nil {
		return wrapError(err), nil
	}
	return makeNetConn(conn), nil
}

func netDial(args ...tender.Object) (tender.Object, error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}
	network, _ := args[0].(*tender.String)
	address, _ := args[1].(*tender.String)
	conn, err := net.Dial(network.Value, address.Value)
	if err != nil {
		return wrapError(err), nil
	}
	return makeNetConn(conn), nil
}

func makeNetConn(conn net.Conn) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"close": &tender.NativeFunction{
				Value: FuncARE(conn.Close),
			},
			"read": &tender.NativeFunction{
				Value: FuncAYRIE(conn.Read),
			},
			"write": &tender.NativeFunction{
				Value: FuncAYRIE(conn.Write),
			},
			"remote_addr": &tender.String{Value: conn.RemoteAddr().String()},
			"local_addr":  &tender.String{Value: conn.LocalAddr().String()},
			"set_deadline": &tender.NativeFunction{
				Value: FuncATRE(conn.SetDeadline),
			},
			"set_readdeadline": &tender.NativeFunction{
				Value: FuncATRE(conn.SetReadDeadline),
			},
			"set_writedeadline": &tender.NativeFunction{
				Value: FuncATRE(conn.SetWriteDeadline),
			},
		},
	}
}

func netDnsLookup(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	host, _ := args[0].(*tender.String)
	addresses, err := net.LookupHost(host.Value)
	if err != nil {
		return wrapError(err), nil
	}
	results := make([]tender.Object, len(addresses))
	for i, addr := range addresses {
		results[i] = &tender.String{Value: addr}
	}
	return &tender.Array{Value: results}, nil
}


func netResolveTCPAddr(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}
	network, _ := args[0].(*tender.String)
	address, _ := args[1].(*tender.String)
	tcpAddr, err := net.ResolveTCPAddr(network.Value, address.Value)
	if err != nil {
		return wrapError(err), nil
	}
	return &tender.String{Value: tcpAddr.String()}, nil
}

func netResolveUDPAddr(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}
	network, _ := args[0].(*tender.String)
	address, _ := args[1].(*tender.String)
	udpAddr, err := net.ResolveUDPAddr(network.Value, address.Value)
	if err != nil {
		return wrapError(err), nil
	}
	return &tender.String{Value: udpAddr.String()}, nil
}

func netListen(args ...tender.Object) (tender.Object, error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}
	network, _ := tender.ToString(args[0])
	addr, _ := tender.ToString(args[1])

	ln, err := net.Listen(network, addr)
	if err != nil {
		return wrapError(err), nil
	}

	return makeNetListener(ln), nil
}

func makeNetListener(ln net.Listener) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"accept": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					conn, err := ln.Accept()
					if err != nil {
						return wrapError(err), nil
					}
					return makeNetConn(conn), nil
				},
			},
			"close": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					err := ln.Close()
					if err != nil {
						return wrapError(err), nil
					}
					return nil, nil
				},
			},
		},
	}
}
