package stdlib

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/2dprototype/tender"
)

var httpModule = map[string]tender.Object{
	"get":     &tender.UserFunction{Name: "get", Value: httpGet},
	"post":    &tender.UserFunction{Name: "post", Value: httpPost},
	"put":     &tender.UserFunction{Name: "put", Value: httpPut},
	"delete":  &tender.UserFunction{Name: "delete", Value: httpDelete},
	"patch":   &tender.UserFunction{Name: "patch", Value: httpPatch},
	"options": &tender.UserFunction{Name: "options", Value: httpOptions},
	"head":    &tender.UserFunction{Name: "head", Value: httpHead},
	"trace":            &tender.UserFunction{Name: "trace", Value: httpTrace},
	"listen_and_serve": &tender.UserFunction{Name: "listen_and_serve", Value: httpListenAndServe},
}

type ServerResponse struct {
	*tender.ImmutableMap
	W http.ResponseWriter
}

type ServerRequest struct {
	*tender.ImmutableMap
	R *http.Request
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
			"method":  &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					return &tender.String{Value: req.Method}, nil
				},
			},
			"url":     &tender.UserFunction{Value: FuncARS(req.URL.String)},
			"headers": makeHeaderMap(req.Header),
			// Executes the request and returns only the body as bytes.
			"body": &tender.UserFunction{
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
			"execute": &tender.UserFunction{
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
			"get_header": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					key, _ := tender.ToString(args[0])
					return &tender.String{Value: req.Header.Get(key)}, nil
				},
			},
			// Set a header (overwrites if exists).
			"set_header": &tender.UserFunction{
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
			"set_body": &tender.UserFunction{
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
			"set_method": &tender.UserFunction{
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
			"set_url": &tender.UserFunction{
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

func httpListenAndServe(args ...tender.Object) (tender.Object, error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}
	addr, _ := tender.ToString(args[0])
	handler := args[1]
	if !handler.CanCall() {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "handler",
			Expected: "callable",
			Found:    handler.TypeName(),
		}
	}

	err := http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := makeServerRequest(r)
		resp := makeServerResponse(w)
		handler.Call(resp, req)
	}))
	return wrapError(err), nil
}

func makeServerRequest(r *http.Request) *ServerRequest {
	m := &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"method":      &tender.String{Value: r.Method},
			"url":         &tender.UserFunction{Value: FuncARS(r.URL.String)},
			"header":      makeHeaderMap(r.Header),
			"remote_addr": &tender.String{Value: r.RemoteAddr},
			"body": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					b, err := ioutil.ReadAll(r.Body)
					if err != nil {
						return wrapError(err), nil
					}
					return &tender.Bytes{Value: b}, nil
				},
			},
		},
	}
	return &ServerRequest{ImmutableMap: m, R: r}
}

func makeServerResponse(w http.ResponseWriter) *ServerResponse {
	m := &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"write": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					b, _ := tender.ToByteSlice(args[0])
					n, err := w.Write(b)
					if err != nil {
						return wrapError(err), nil
					}
					return &tender.Int{Value: int64(n)}, nil
				},
			},
			"write_header": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					code, _ := tender.ToInt(args[0])
					w.WriteHeader(int(code))
					return nil, nil
				},
			},
			"set_header": &tender.UserFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 {
						return nil, tender.ErrWrongNumArguments
					}
					k, _ := tender.ToString(args[0])
					v, _ := tender.ToString(args[1])
					w.Header().Set(k, v)
					return nil, nil
				},
			},
		},
	}
	return &ServerResponse{ImmutableMap: m, W: w}
}
