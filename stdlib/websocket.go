package stdlib

import (
	"net/http"
	"fmt"

	"github.com/2dprototype/tender"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow all origins for the script
}

var websocketModule = map[string]tender.Object{
	"dial": &tender.NativeFunction{Value: wsDial},
	"listen_and_serve": &tender.NativeFunction{
		Name:      "listen_and_serve",
		Value:     wsListenAndServe,
		NeedVMObj: true,
	},
}

func wsListenAndServe(args ...tender.Object) (tender.Object, error) {
	if len(args) < 3 {
		return nil, tender.ErrWrongNumArguments
	}
	vmObj, ok := args[0].(*tender.VMObj)
	if !ok {
		return &tender.Error{Value: &tender.String{Value: "Internal Error: Missing VM context"}}, nil
	}
	vm := vmObj.Value

	addr, _ := tender.ToString(args[1])
	handlerFunc := args[2]

	if !handlerFunc.CanCall() {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "handler",
			Expected: "callable function",
			Found:    handlerFunc.TypeName(),
		}
	}

	server := &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				return
			}
			// Don't defer close here; let the script manage the connection lifecycle
			
			wrappedConn := makeWsConn(conn)
			
			// Execute script callback
			_, err = tender.WrapFuncCall(vm, handlerFunc, wrappedConn)
			if err != nil {
				fmt.Println(err)
			}
		}),
	}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return wrapError(err), nil
	}

	return nil, nil
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

	return makeWsConn(conn), nil
}

func makeWsConn(conn *websocket.Conn) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"read_message": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					t, b, err := conn.ReadMessage()
					if err != nil {
						return wrapError(err), nil
					}
					return &tender.Array{
						Value: []tender.Object{
							&tender.Int{Value: int64(t)},
							&tender.Bytes{Value: b},
						},
					}, nil
				},
			},
			"write_message": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 { return nil, tender.ErrWrongNumArguments }
					t, _ := tender.ToInt(args[0])
					b, _ := tender.ToByteSlice(args[1])
					
					err := conn.WriteMessage(int(t), b)
					if err != nil {
						return wrapError(err), nil
					}
					return nil, nil
				},
			},
			"close": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					err := conn.Close()
					if err != nil {
						return wrapError(err), nil
					}
					return nil, nil
				},
			},
			"remote_addr": &tender.String{Value: conn.RemoteAddr().String()},
			"local_addr":  &tender.String{Value: conn.LocalAddr().String()},
		},
	}
}