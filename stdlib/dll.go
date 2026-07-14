//go:build windows

package stdlib


import (
	"syscall"
	"unsafe"
	"encoding/binary"
	"runtime"
	"unicode/utf16"
	"fmt"

	"github.com/2dprototype/tender"
)

var dllModule = map[string]tender.Object{
	"call_dll":      &tender.NativeFunction{Name: "call_dll", Value: callDLLWrapper},
	"new":           &tender.NativeFunction{Name: "new", Value: dllNew},
	"load":          &tender.NativeFunction{Name: "load", Value: dllLoad},
	"last_error":    &tender.NativeFunction{Name: "last_error", Value: getLastError},
	"free_library":  &tender.NativeFunction{Name: "free_library", Value: freeLibrary},
	"get_proc_address": &tender.NativeFunction{Name: "get_proc_address", Value: getProcAddress},
	"memory":        &tender.NativeFunction{Name: "memory", Value: memoryOperations},
	"callback":      &tender.NativeFunction{Name: "callback", Value: createCallback},
	"struct":        &tender.NativeFunction{Name: "struct", Value: createStruct},
	"pointer":       &tender.NativeFunction{Name: "pointer", Value: pointerOperations},
}

// Extended DLL structure to track loaded libraries
type DLLContext struct {
	dll      *syscall.DLL
	lazyDLL  *syscall.LazyDLL
	isLazy   bool
	handle   uintptr
}

var loadedDLLs = make(map[uintptr]*DLLContext)

func dllLoad(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	
	dllName, ok := args[0].(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "dll_name",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	
	dll, err := syscall.LoadDLL(dllName.Value)
	if err != nil {
		return wrapError(err), nil
	}
	
	// Store the DLL context
	ctx := &DLLContext{
		dll:    dll,
		isLazy: false,
		handle: uintptr(unsafe.Pointer(dll)),
	}
	loadedDLLs[ctx.handle] = ctx
	
	return makeDLLObject(ctx), nil
}

func dllNew(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	
	dllName, ok := args[0].(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "dll_name",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	
	dll := syscall.NewLazyDLL(dllName.Value)
	
	ctx := &DLLContext{
		lazyDLL: dll,
		isLazy:  true,
		handle:  uintptr(unsafe.Pointer(dll)),
	}
	loadedDLLs[ctx.handle] = ctx
	
	return makeDLLObject(ctx), nil
}

func makeDLLObject(ctx *DLLContext) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"proc": &tender.NativeFunction{
				Name:  "proc",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					functionName, ok := args[0].(*tender.String)
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "function_name",
							Expected: "string",
							Found:    args[0].TypeName(),
						}
					}
					
					if ctx.isLazy {
						proc := ctx.lazyDLL.NewProc(functionName.Value)
						return makeProcObject(proc, true, ctx), nil
					} else {
						proc, err := ctx.dll.FindProc(functionName.Value)
						if err != nil {
							return wrapError(err), nil
						}
						return makeProcObject(proc, false, ctx), nil
					}
				},
			},
			"handle": &tender.Int{Value: int64(ctx.handle)},
			"name": &tender.String{Value: func() string {
				if ctx.isLazy {
					return ctx.lazyDLL.Name
				}
				return ctx.dll.Name
			}()},
			"unload": &tender.NativeFunction{
				Name: "unload",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if ctx.isLazy {
						// LazyDLL doesn't have explicit unload
						return tender.TrueValue, nil
					} else {
						err := ctx.dll.Release()
						if err != nil {
							return wrapError(err), nil
						}
						delete(loadedDLLs, ctx.handle)
						return tender.TrueValue, nil
					}
				},
			},
		},
	}
}

func makeProcObject(proc interface{}, isLazy bool, ctx *DLLContext) *tender.ImmutableMap {
	procMap := map[string]tender.Object{
		"call": &tender.NativeFunction{
			Name:  "call",
			Value: makeProcCaller(proc, isLazy),
		},
		"name": &tender.String{Value: func() string {
			if isLazy {
				return proc.(*syscall.LazyProc).Name
			}
			return proc.(*syscall.Proc).Name
		}()},
		"addr": &tender.Int{Value: func() int64 {
			if isLazy {
				return int64(proc.(*syscall.LazyProc).Addr())
			}
			return int64(proc.(*syscall.Proc).Addr())
		}()},
	}
	
	return &tender.ImmutableMap{Value: procMap}
}

func makeProcCaller(proc interface{}, isLazy bool) tender.CallableFunc {
	return func(args ...tender.Object) (tender.Object, error) {
		var uintptrArgs []uintptr
		
		// Parse arguments with type support
		for i, arg := range args {
			switch val := arg.(type) {
			case *tender.Int:
				uintptrArgs = append(uintptrArgs, uintptr(val.Value))
			case *tender.Float:
				// Convert float to uintptr (note: this may lose precision on 32-bit)
				uintptrArgs = append(uintptrArgs, uintptr(val.Value))
			case *tender.String:
				// For strings, we need to ensure they're properly converted to UTF16
				utf16Str := utf16.Encode([]rune(val.Value + "\x00"))
				ptr := unsafe.Pointer(&utf16Str[0])
				uintptrArgs = append(uintptrArgs, uintptr(ptr))
				// Keep reference to prevent GC
				runtime.KeepAlive(utf16Str)
			case *tender.Bytes:
				if len(val.Value) > 0 {
					uintptrArgs = append(uintptrArgs, uintptr(unsafe.Pointer(&val.Value[0])))
				} else {
					uintptrArgs = append(uintptrArgs, 0)
				}
			case *tender.Null:
				uintptrArgs = append(uintptrArgs, 0)
			case *tender.Bool:
				if val.IsFalsy() {
					uintptrArgs = append(uintptrArgs, 0)
				} else {
					uintptrArgs = append(uintptrArgs, 1)
				}
			default:
				// Check if it's a pointer object
				if ptr, ok := isPointerObject(arg); ok {
					uintptrArgs = append(uintptrArgs, uintptr(ptr))
				} else {
					return nil, tender.ErrInvalidArgumentType{
						Name:     fmt.Sprintf("argument_%d", i),
						Expected: "integer, float, string, bytes, bool, null, or pointer",
						Found:    arg.TypeName(),
					}
				}
			}
		}
		
		var ret uintptr
		var err error
		
		if isLazy {
			lazyProc := proc.(*syscall.LazyProc)
			ret, _, err = lazyProc.Call(uintptrArgs...)
		} else {
			regularProc := proc.(*syscall.Proc)
			ret, _, err = regularProc.Call(uintptrArgs...)
		}
		
		if err != nil {
			// Check if it's a success with error code (common in Windows API)
			if ret != 0 {
				// Many Windows APIs return non-zero for success, but set last error
				return &tender.Int{Value: int64(ret)}, nil
			}
			return wrapError(err), nil
		}
		
		return &tender.Int{Value: int64(ret)}, nil
	}
}

// Enhanced callDLLWrapper with better error handling
func callDLLWrapper(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) < 2 {
		return nil, tender.ErrWrongNumArguments
	}
	
	dllName, ok := args[0].(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "dll_name",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	
	functionName, ok := args[1].(*tender.String)
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "function_name",
			Expected: "string",
			Found:    args[1].TypeName(),
		}
	}
	
	var uintptrArgs []uintptr
	for i, arg := range args[2:] {
		switch val := arg.(type) {
		case *tender.Int:
			uintptrArgs = append(uintptrArgs, uintptr(val.Value))
		case *tender.Float:
			uintptrArgs = append(uintptrArgs, uintptr(val.Value))
		case *tender.String:
			utf16Str := utf16.Encode([]rune(val.Value + "\x00"))
			ptr := unsafe.Pointer(&utf16Str[0])
			uintptrArgs = append(uintptrArgs, uintptr(ptr))
			runtime.KeepAlive(utf16Str)
		case *tender.Bytes:
			if len(val.Value) > 0 {
				uintptrArgs = append(uintptrArgs, uintptr(unsafe.Pointer(&val.Value[0])))
			} else {
				uintptrArgs = append(uintptrArgs, 0)
			}
		case *tender.Null:
			uintptrArgs = append(uintptrArgs, 0)
		case *tender.Bool:
			if val.IsFalsy() {
				uintptrArgs = append(uintptrArgs, 0)
			} else {
				uintptrArgs = append(uintptrArgs, 1)
			}
		default:
			if ptr, ok := isPointerObject(arg); ok {
				uintptrArgs = append(uintptrArgs, uintptr(ptr))
			} else {
				return nil, tender.ErrInvalidArgumentType{
					Name:     fmt.Sprintf("argument_%d", i+2),
					Expected: "integer, float, string, bytes, bool, null, or pointer",
					Found:    arg.TypeName(),
				}
			}
		}
	}
	
	result, err := callDLL(dllName.Value, functionName.Value, uintptrArgs...)
	if err != nil {
		return wrapError(err), nil
	}
	
	return &tender.Int{Value: int64(result)}, nil
}

func callDLL(dllName, functionName string, args ...uintptr) (ret uintptr, err error) {
	dll := syscall.NewLazyDLL(dllName)
	proc := dll.NewProc(functionName)
	
	ret, _, err = proc.Call(args...)
	return ret, err
}

// Error handling functions
func getLastError(args ...tender.Object) (ret tender.Object, err error) {
	errCode := syscall.GetLastError()
	return wrapError(errCode), nil
}

func freeLibrary(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}
	
	handle, ok := tender.ToInt(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "handle",
			Expected: "int",
			Found:    args[0].TypeName(),
		}
	}
	
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	freeLibraryProc := kernel32.NewProc("FreeLibrary")
	
	result, _, err := freeLibraryProc.Call(uintptr(handle))
	if result == 0 {
		return wrapError(err), nil
	}
	
	return tender.TrueValue, nil
}

func getProcAddress(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}
	
	handle, ok := tender.ToInt(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "handle",
			Expected: "int",
			Found:    args[0].TypeName(),
		}
	}
	
	procName, ok := tender.ToString(args[1])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "proc_name",
			Expected: "string",
			Found:    args[1].TypeName(),
		}
	}
	
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getProcAddressProc := kernel32.NewProc("GetProcAddress")
	
	// Convert proc name to byte pointer
	procNameBytes := append([]byte(procName), 0)
	addr, _, err := getProcAddressProc.Call(uintptr(handle), uintptr(unsafe.Pointer(&procNameBytes[0])))
	
	if addr == 0 {
		return wrapError(err), nil
	}
	
	return &tender.Int{Value: int64(addr)}, nil
}

// Memory operations
func memoryOperations(args ...tender.Object) (ret tender.Object, err error) {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"alloc": &tender.NativeFunction{
				Name: "alloc",
				Value: func(args ...tender.Object) (tender.Object, error) {
					size := 1024
					if len(args) > 0 {
						sizeArg, ok := tender.ToInt(args[0])
						if !ok {
							return nil, tender.ErrInvalidArgumentType{
								Name:     "size",
								Expected: "int",
								Found:    args[0].TypeName(),
							}
						}
						size = sizeArg
					}
					
					kernel32 := syscall.NewLazyDLL("kernel32.dll")
					heapAlloc := kernel32.NewProc("HeapAlloc")
					getProcessHeap := kernel32.NewProc("GetProcessHeap")
					
					heap, _, _ := getProcessHeap.Call()
					ptr, _, err := heapAlloc.Call(heap, 0, uintptr(size))
					
					if ptr == 0 {
						return wrapError(err), nil
					}
					
					return createPointerObject(ptr, size), nil
				},
			},
			"free": &tender.NativeFunction{
				Name: "free",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					
					ptr, ok := isPointerObject(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "pointer",
							Expected: "pointer",
							Found:    args[0].TypeName(),
						}
					}
					
					kernel32 := syscall.NewLazyDLL("kernel32.dll")
					heapFree := kernel32.NewProc("HeapFree")
					getProcessHeap := kernel32.NewProc("GetProcessHeap")
					
					heap, _, _ := getProcessHeap.Call()
					result, _, err := heapFree.Call(heap, 0, uintptr(ptr))
					
					if result == 0 {
						return wrapError(err), nil
					}
					
					return tender.TrueValue, nil
				},
			},
			"copy": &tender.NativeFunction{
				Name: "copy",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 3 {
						return nil, tender.ErrWrongNumArguments
					}
					
					destPtr, ok1 := isPointerObject(args[0])
					srcPtr, ok2 := isPointerObject(args[1])
					size, ok3 := tender.ToInt(args[2])
					
					if !ok1 || !ok2 || !ok3 {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "arguments",
							Expected: "pointer, pointer, int",
							Found:    fmt.Sprintf("%s, %s, %s", args[0].TypeName(), args[1].TypeName(), args[2].TypeName()),
						}
					}
					
					kernel32 := syscall.NewLazyDLL("kernel32.dll")
					rtlMoveMemory := kernel32.NewProc("RtlMoveMemory")
					
					rtlMoveMemory.Call(uintptr(destPtr), uintptr(srcPtr), uintptr(size))
					return tender.TrueValue, nil
				},
			},
			"read_string": &tender.NativeFunction{
				Name: "read_string",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) < 1 {
						return nil, tender.ErrWrongNumArguments
					}
					
					ptrObj, ok := isPointerObject(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "pointer",
							Expected: "pointer",
							Found:    args[0].TypeName(),
						}
					}
					
					maxLen := 256
					if len(args) > 1 {
						maxLenArg, ok := tender.ToInt(args[1])
						if ok {
							maxLen = maxLenArg
						}
					}
					
					// Read null-terminated UTF-16 string from pointer
					var result []uint16
					for i := 0; i < maxLen; i++ {
						// Read 2 bytes at a time (UTF-16 character)
						charPtr := unsafe.Pointer(ptrObj + uintptr(i*2))
						char := *(*uint16)(charPtr)
						if char == 0 {
							break
						}
						result = append(result, char)
					}
					
					if len(result) > 0 {
						return &tender.String{Value: string(utf16.Decode(result))}, nil
					}
					
					return &tender.String{Value: ""}, nil
				},
			},
		},
	}, nil
}

// Callback support (simplified)
func createCallback(args ...tender.Object) (ret tender.Object, err error) {
	// This is a complex feature that would require proper implementation
	// For now, return a placeholder
	return &tender.String{Value: "Callback support requires additional implementation"}, nil
}

// Struct creation for passing structured data
func createStruct(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) == 0 {
		return nil, tender.ErrWrongNumArguments
	}
	
	// Simple struct creation - just allocate memory and return pointer
	size := 1024
	if len(args) > 0 {
		sizeArg, ok := tender.ToInt(args[0])
		if ok {
			size = sizeArg
		}
	}
	
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	heapAlloc := kernel32.NewProc("HeapAlloc")
	getProcessHeap := kernel32.NewProc("GetProcessHeap")
	
	heap, _, _ := getProcessHeap.Call()
	ptr, _, err := heapAlloc.Call(heap, 0, uintptr(size))
	
	if ptr == 0 {
		return wrapError(err), nil
	}
	
	// Zero the memory
	zeroMemory := kernel32.NewProc("RtlZeroMemory")
	zeroMemory.Call(ptr, uintptr(size))
	
	return createPointerObject(ptr, size), nil
}

// Pointer operations
func pointerOperations(args ...tender.Object) (ret tender.Object, err error) {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"create": &tender.NativeFunction{
				Name: "create",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					
					addr, ok := tender.ToInt(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "address",
							Expected: "int",
							Found:    args[0].TypeName(),
						}
					}
					
					return createPointerObject(uintptr(addr), 0), nil
				},
			},
			"offset": &tender.NativeFunction{
				Name: "offset",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 {
						return nil, tender.ErrWrongNumArguments
					}
					
					ptr, ok := isPointerObject(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "pointer",
							Expected: "pointer",
							Found:    args[0].TypeName(),
						}
					}
					
					offset, ok := tender.ToInt(args[1])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "offset",
							Expected: "int",
							Found:    args[1].TypeName(),
						}
					}
					
					newPtr := ptr + uintptr(offset)
					return createPointerObject(newPtr, 0), nil
				},
			},
			"read_int": &tender.NativeFunction{
				Name: "read_int",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					
					ptr, ok := isPointerObject(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "pointer",
							Expected: "pointer",
							Found:    args[0].TypeName(),
						}
					}
					
					// Read 4-byte integer directly from memory
					data := (*[4]byte)(unsafe.Pointer(ptr))[:]
					value := int32(binary.LittleEndian.Uint32(data))
					return &tender.Int{Value: int64(value)}, nil
				},
			},
			"write_int": &tender.NativeFunction{
				Name: "write_int",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 {
						return nil, tender.ErrWrongNumArguments
					}
					
					ptr, ok := isPointerObject(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "pointer",
							Expected: "pointer",
							Found:    args[0].TypeName(),
						}
					}
					
					value, ok := tender.ToInt(args[1])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "value",
							Expected: "int",
							Found:    args[1].TypeName(),
						}
					}
					
					// Write 4-byte integer directly to memory
					data := (*[4]byte)(unsafe.Pointer(ptr))[:]
					binary.LittleEndian.PutUint32(data, uint32(value))
					return tender.TrueValue, nil
				},
			},
			"read_bytes": &tender.NativeFunction{
				Name: "read_bytes",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 {
						return nil, tender.ErrWrongNumArguments
					}
					
					ptr, ok := isPointerObject(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "pointer",
							Expected: "pointer",
							Found:    args[0].TypeName(),
						}
					}
					
					size, ok := tender.ToInt(args[1])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "size",
							Expected: "int",
							Found:    args[1].TypeName(),
						}
					}
					
					if size <= 0 {
						return &tender.Bytes{Value: []byte{}}, nil
					}
					
					// Read bytes directly from memory
					data := make([]byte, size)
					for i := 0; i < size; i++ {
						data[i] = *(*byte)(unsafe.Pointer(ptr + uintptr(i)))
					}
					
					return &tender.Bytes{Value: data}, nil
				},
			},
			"write_bytes": &tender.NativeFunction{
				Name: "write_bytes",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 {
						return nil, tender.ErrWrongNumArguments
					}
					
					ptr, ok := isPointerObject(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "pointer",
							Expected: "pointer",
							Found:    args[0].TypeName(),
						}
					}
					
					data, ok := args[1].(*tender.Bytes)
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "data",
							Expected: "bytes",
							Found:    args[1].TypeName(),
						}
					}
					
					// Write bytes directly to memory
					for i, b := range data.Value {
						*(*byte)(unsafe.Pointer(ptr + uintptr(i))) = b
					}
					
					return tender.TrueValue, nil
				},
			},
		},
	}, nil
}

// Helper functions
func createPointerObject(ptr uintptr, size int) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"address": &tender.Int{Value: int64(ptr)},
			"size":    &tender.Int{Value: int64(size)},
			"type":    &tender.String{Value: "pointer"},
		},
	}
}

func isPointerObject(obj tender.Object) (uintptr, bool) {
	if imm, ok := obj.(*tender.ImmutableMap); ok {
		if addr, exists := imm.Value["address"]; exists {
			if addrInt, ok := addr.(*tender.Int); ok {
				return uintptr(addrInt.Value), true
			}
		}
	}
	return 0, false
}
