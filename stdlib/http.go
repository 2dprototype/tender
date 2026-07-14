package stdlib

import (
	"io"
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/2dprototype/tender"
)

var httpModule = map[string]tender.Object{
	"get":     &tender.NativeFunction{Name: "get", Value: httpGet},
	"post":    &tender.NativeFunction{Name: "post", Value: httpPost},
	"put":     &tender.NativeFunction{Name: "put", Value: httpPut},
	"delete":  &tender.NativeFunction{Name: "delete", Value: httpDelete},
	"patch":   &tender.NativeFunction{Name: "patch", Value: httpPatch},
	"options": &tender.NativeFunction{Name: "options", Value: httpOptions},
	"head":    &tender.NativeFunction{Name: "head", Value: httpHead},
	"trace":   &tender.NativeFunction{Name: "trace", Value: httpTrace},
	"listen_and_serve": &tender.NativeFunction{
		Name:      "listen_and_serve", 
		Value:     httpListenAndServe,
		NeedVMObj: true,
	},
}

// httpRequest creates an http.Request and wraps it in an object with helper methods.
func httpRequest(method string, args ...tender.Object) (ret tender.Object, err error) {
	if len(args) < 1 || len(args) > 3 {
		return nil, tender.ErrWrongNumArguments
	}

	url, _ := tender.ToString(args[0])
	var body []byte
	if len(args) > 1 {
		body, _ = tender.ToByteSlice(args[1])
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return wrapError(err), nil
	}

	if len(args) == 3 {
		headersObj, ok := args[2].(*tender.Map)
		if !ok {
			return nil, tender.ErrInvalidArgumentType{
				Name:     "headers",
				Expected: "map",
				Found:    args[2].TypeName(),
			}
		}
		for key, value := range headersObj.Value {
			valueStr, _ := value.(*tender.String)
			req.Header.Set(key, valueStr.Value)
		}
	}

	return makeHttpReq(req), nil
}

func makeHttpReq(req *http.Request) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"close":   tender.FromBool(req.Close),
			// "method":  &tender.String{Value: req.Method},
			"method":  &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					return &tender.String{Value: req.Method}, nil
				},
			},
			"url":     &tender.NativeFunction{Value: FuncARS(req.URL.String)},
			"headers": makeHeaderMap(req.Header),
			// Executes the request and returns only the body as bytes.
			"body": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					client := &http.Client{}
					resp, err := client.Do(req)
					if err != nil {
						return wrapError(err), nil
					}
					defer resp.Body.Close()
					respBody, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						return wrapError(err), nil
					}
					return &tender.Bytes{Value: respBody}, nil
				},
			},
			// Executes the request and returns a full response object.
			"execute": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					client := &http.Client{}
					resp, err := client.Do(req)
					if err != nil {
						return wrapError(err), nil
					}
					defer resp.Body.Close()
					respBody, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						return wrapError(err), nil
					}
					return makeHttpResponse(resp, respBody), nil
				},
			},
			// Get the value of a header by key.
			"get_header": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					key, _ := tender.ToString(args[0])
					return &tender.String{Value: req.Header.Get(key)}, nil
				},
			},
			// Set a header (overwrites if exists).
			"set_header": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 {
						return nil, tender.ErrWrongNumArguments
					}
					key, _ := tender.ToString(args[0])
					value, _ := tender.ToString(args[1])
					req.Header.Set(key, value)
					return nil, nil
				},
			},
			// Set the request body.
			"set_body": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					body, _ := tender.ToByteSlice(args[0])
					req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
					return nil, nil
				},
			},
			// Change the HTTP method.
			"set_method": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					m, _ := tender.ToString(args[0])
					req.Method = m
					return nil, nil
				},
			},
			// Change the URL.
			"set_url": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					urlStr, _ := tender.ToString(args[0])
					// Recreate the request URL; we reuse the current method and body.
					parsedReq, err := http.NewRequest(req.Method, urlStr, req.Body)
					if err != nil {
						return wrapError(err), nil
					}
					req.URL = parsedReq.URL
					return nil, nil
				},
			},
		},
	}
}

func makeHeaderMap(headers http.Header) *tender.Map {
	headerMap := make(map[string]tender.Object)
	for key, values := range headers {
		if len(values) > 0 {
			headerMap[key] = &tender.String{Value: values[0]}
		}
	}
	return &tender.Map{Value: headerMap}
}

func makeHttpResponse(resp *http.Response, body []byte) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"status":      &tender.Int{Value: int64(resp.StatusCode)},
			"status_text":  &tender.String{Value: resp.Status},
			"headers":     makeHeaderMap(resp.Header),
			"body":        &tender.Bytes{Value: body},
			"content_type": &tender.String{Value: resp.Header.Get("Content-Type")},
		},
	}
}

func httpGet(args ...tender.Object) (tender.Object, error) {
	return httpRequest("GET", args...)
}

func httpPost(args ...tender.Object) (tender.Object, error) {
	return httpRequest("POST", args...)
}

func httpPut(args ...tender.Object) (tender.Object, error) {
	return httpRequest("PUT", args...)
}

func httpDelete(args ...tender.Object) (tender.Object, error) {
	return httpRequest("DELETE", args...)
}

func httpPatch(args ...tender.Object) (tender.Object, error) {
	return httpRequest("PATCH", args...)
}

func httpOptions(args ...tender.Object) (tender.Object, error) {
	return httpRequest("OPTIONS", args...)
}

func httpHead(args ...tender.Object) (tender.Object, error) {
	return httpRequest("HEAD", args...)
}

func httpTrace(args ...tender.Object) (tender.Object, error) {
	return httpRequest("TRACE", args...)
}

// httpListenAndServe starts an HTTP server. 
func httpListenAndServe(args ...tender.Object) (tender.Object, error) {
	// 1. Extract the VM object injected by NeedVMObj: true
	if len(args) < 1 {
		return nil, tender.ErrWrongNumArguments
	}
	vmObj, ok := args[0].(*tender.VMObj)
	if !ok {
		// Fallback in case of internal engine mismatch
		return &tender.Error{Value: &tender.String{Value: "Internal Error: Missing VM context"}}, nil
	}
	vm := vmObj.Value

	// 2. Shift the args slice to process the actual user arguments
	args = args[1:]

	// 3. Validate user arguments (addr string, handler callable)
	if len(args) < 2 {
		return nil, tender.ErrWrongNumArguments
	}

	addr, _ := tender.ToString(args[0])
	handlerFunc := args[1]

	if !handlerFunc.CanCall() {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "handler",
			Expected: "callable function",
			Found:    handlerFunc.TypeName(),
		}
	}

	// 4. Setup the HTTP Server
	server := &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Read the incoming body safely
			bodyBytes, _ := io.ReadAll(r.Body)
			defer r.Body.Close()

			// Prepare wrapped request mapping for the tender script space
			wrappedReq := makeInboundReq(r, bodyBytes)
			
			// Setup an updated custom Response Writer handle
			wrappedRes, writerObj := makeInboundRes(w)

			// Execute handler code inside script space cleanly using WrapFuncCall
			// Pass the extracted 'vm' here instead of the placeholder function!
			_, err := tender.WrapFuncCall(vm, handlerFunc, wrappedReq, wrappedRes)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal script compilation/execution error"))
				return
			}

			// Commit delayed headers and buffered output back to the real pipeline
			writerObj.Flush(w)
		}),
	}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return &tender.Error{Value: &tender.String{Value: err.Error()}}, nil
	}

	return nil, nil
}

func makeInboundReq(req *http.Request, body []byte) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"method":      &tender.String{Value: req.Method},
			"url":         &tender.String{Value: req.URL.String()},
			"path":        &tender.String{Value: req.URL.Path},
			"headers":     makeHeaderMap(req.Header),
			"body":        &tender.Bytes{Value: body},
			"remote_addr": &tender.String{Value: req.RemoteAddr},
		},
	}
}

type inboundResponseWriter struct {
	statusCode int
	headers    http.Header
	bodyBuffer *bytes.Buffer
}

func (w *inboundResponseWriter) Flush(realWriter http.ResponseWriter) {
	for k, vv := range w.headers {
		for _, v := range vv {
			realWriter.Header().Add(k, v)
		}
	}
	if w.statusCode > 0 {
		realWriter.WriteHeader(w.statusCode)
	}
	_, _ = realWriter.Write(w.bodyBuffer.Bytes())
}

func makeInboundRes(realWriter http.ResponseWriter) (*tender.ImmutableMap, *inboundResponseWriter) {
	writerObj := &inboundResponseWriter{
		statusCode: http.StatusOK,
		headers:    make(http.Header),
		bodyBuffer: new(bytes.Buffer),
	}

	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"set_status": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 { return nil, tender.ErrWrongNumArguments }
					val, _ := tender.ToInt(args[0])
					writerObj.statusCode = val
					return nil, nil
				},
			},
			"set_header": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 { return nil, tender.ErrWrongNumArguments }
					k, _ := tender.ToString(args[0])
					v, _ := tender.ToString(args[1])
					writerObj.headers.Set(k, v)
					return nil, nil
				},
			},
			"write": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 { return nil, tender.ErrWrongNumArguments }
					switch b := args[0].(type) {
					case *tender.Bytes:
						writerObj.bodyBuffer.Write(b.Value)
					default:
						str, _ := tender.ToString(args[0])
						writerObj.bodyBuffer.WriteString(str)
					}
					return nil, nil
				},
			},
		},
	}, writerObj
}