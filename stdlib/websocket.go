package stdlib

import (
	"github.com/gorilla/websocket"
	"github.com/2dprototype/tender"
)


var websocketModule = map[string]tender.Object{
	"dial":    &tender.UserFunction{Value: wsDial},
	"upgrade": &tender.UserFunction{Value: wsUpgrade},
}

func wsUpgrade(args ...tender.Object) (tender.Object, error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}

	res, ok1 := args[0].(*ServerResponse)
	req, ok2 := args[1].(*ServerRequest)

	if !ok1 {
		return nil, tender.ErrInvalidArgumentType{Name: "response", Expected: "http.ServerResponse", Found: args[0].TypeName()}
	}
	if !ok2 {
		return nil, tender.ErrInvalidArgumentType{Name: "request", Expected: "http.ServerRequest", Found: args[1].TypeName()}
	}

	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(res.W, req.R, nil)
	if err != nil {
		return wrapError(err), nil
	}

	return makeWsConn(*conn), nil
}

func wsDial(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	url, _ := tender.ToString(args[0])
	
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return wrapError(err), nil
	}

	return makeWsConn(*conn), nil
}

func makeWsConn(conn websocket.Conn) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"read_message": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					t, b, err := conn.ReadMessage()
					if err != nil {
						return wrapError(err), nil
					}
					return &tender.Array{Value: 
						[]tender.Object{
							&tender.Int{Value: int64(t)},
							&tender.Bytes{Value: b},
						},
					}, nil
				},
			},	
			"write_message": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 {
						return nil, tender.ErrWrongNumArguments
					}
					t, _ := tender.ToInt(args[0])
					b, _ := tender.ToByteSlice(args[1])
					err := conn.WriteMessage(t, b)
					if err != nil {
						return wrapError(err), nil
					}
					return nil, nil
				},
			},
			"close": &tender.UserFunction{
				Value: FuncARE(conn.Close),
			},
			"local_addr": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					addr := conn.LocalAddr()
					return &tender.ImmutableMap{
						Value: map[string]tender.Object{
							"string" :  &tender.UserFunction{Value: FuncARS(addr.String)},
							"network" :  &tender.UserFunction{Value: FuncARS(addr.Network)},
						},
					}, nil
				},
			},
			"remote_addr": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					addr := conn.RemoteAddr()
					return &tender.ImmutableMap{
						Value: map[string]tender.Object{
							"string" :  &tender.UserFunction{Value: FuncARS(addr.String)},
							"network" :  &tender.UserFunction{Value: FuncARS(addr.Network)},
						},
					}, nil
				},
			},
			"set_readdeadline": &tender.UserFunction{
				Value: FuncATRE(conn.SetReadDeadline),
			},
			"set_writedeadline": &tender.UserFunction{
				Value: FuncATRE(conn.SetWriteDeadline),
			},
		},
	}
}