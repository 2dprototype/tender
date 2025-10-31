## Stdlib `dll`

The `dll` module provides comprehensive functionalities for loading, managing, and calling functions from dynamic-link libraries (DLLs) in the Windows environment, including memory management, error handling, and system API interactions.

## Core Functions

### `new(dll_name)`
Loads a dynamic-link library using lazy loading (loads on first use).
- **Parameters**: `dll_name` - Name of the DLL file to load (string)
- **Returns**: DLL object with procedure access
- **Example**: `kernel32 := dll.new("kernel32.dll")`

### `load(dll_name)`
Immediately loads a dynamic-link library.
- **Parameters**: `dll_name` - Name of the DLL file to load (string)
- **Returns**: DLL object with procedure access
- **Example**: `user32 := dll.load("user32.dll")`

### `call_dll(dll_name, function_name, arguments...)`
Directly calls a DLL function without creating a DLL object.
- **Parameters**: `dll_name` - DLL name, `function_name` - Function name, `arguments` - Function arguments
- **Returns**: Function result as integer
- **Example**: `result := dll.call_dll("kernel32.dll", "Beep", 500, 300)`

## DLL Object Functions

### `proc(function_name)`
Retrieves a procedure address from the loaded DLL.
- **Parameters**: `function_name` - Name of the function to retrieve (string)
- **Returns**: Procedure object for calling the function
- **Example**: `message_box := user32.proc("MessageBoxW")`

### `handle`
Returns the handle to the loaded DLL.
- **Returns**: DLL handle as integer
- **Example**: `handle := kernel32.handle`

### `name`
Returns the name of the loaded DLL.
- **Returns**: DLL name as string
- **Example**: `name := kernel32.name`

### `unload()`
Unloads the DLL from memory (only for DLLs loaded with `load()`).
- **Returns**: Boolean indicating success
- **Example**: `success := dll_obj.unload()`

## Procedure Object Functions

### `call(arguments...)`
Calls the DLL function with the specified arguments.
- **Parameters**: `arguments` - Function arguments (integers, floats, strings, bytes, booleans, null, or pointers)
- **Returns**: Function result as integer
- **Example**: `result := add_func.call(3, 4)`

### `name`
Returns the name of the procedure.
- **Returns**: Function name as string
- **Example**: `func_name := proc.name`

### `addr`
Returns the memory address of the procedure.
- **Returns**: Function address as integer
- **Example**: `address := proc.addr`

## System Functions

### `last_error()`
Gets the last error code from Windows API.
- **Returns**: Last error code as integer
- **Example**: `error_code := dll.last_error()`

### `free_library(handle)`
Unloads a DLL from memory by handle.
- **Parameters**: `handle` - DLL handle to unload
- **Returns**: Boolean indicating success
- **Example**: `success := dll.free_library(handle)`

### `get_proc_address(handle, proc_name)`
Gets the address of a function in a loaded DLL.
- **Parameters**: `handle` - DLL handle, `proc_name` - Function name
- **Returns**: Function address as integer
- **Example**: `addr := dll.get_proc_address(handle, "FunctionName")`

## Memory Management

### `memory.alloc(size)`
Allocates memory from the process heap.
- **Parameters**: `size` - Number of bytes to allocate
- **Returns**: Pointer object to allocated memory
- **Example**: `mem := dll.memory.alloc(1024)`

### `memory.free(pointer)`
Frees previously allocated memory.
- **Parameters**: `pointer` - Pointer object to free
- **Returns**: Boolean indicating success
- **Example**: `success := dll.memory.free(mem)`

### `memory.copy(dest, src, size)`
Copies memory between pointers.
- **Parameters**: `dest` - Destination pointer, `src` - Source pointer, `size` - Number of bytes to copy
- **Returns**: Boolean indicating success
- **Example**: `success := dll.memory.copy(dest_ptr, src_ptr, 100)`

### `memory.read_string(pointer, max_len)`
Reads a null-terminated string from memory.
- **Parameters**: `pointer` - Pointer to string data, `max_len` - Maximum length to read
- **Returns**: String read from memory
- **Example**: `text := dll.memory.read_string(str_ptr, 256)`

## Pointer Operations

### `pointer.create(address)`
Creates a pointer object from an address.
- **Parameters**: `address` - Memory address
- **Returns**: Pointer object
- **Example**: `ptr := dll.pointer.create(0x12345678)`

### `pointer.offset(pointer, offset)`
Creates a new pointer with an offset from the original.
- **Parameters**: `pointer` - Base pointer, `offset` - Byte offset
- **Returns**: New pointer object
- **Example**: `new_ptr := dll.pointer.offset(ptr, 16)`

### `pointer.read_int(pointer)`
Reads a 4-byte integer from memory.
- **Parameters**: `pointer` - Pointer to integer data
- **Returns**: Integer value
- **Example**: `value := dll.pointer.read_int(int_ptr)`

### `pointer.write_int(pointer, value)`
Writes a 4-byte integer to memory.
- **Parameters**: `pointer` - Pointer to write to, `value` - Integer value to write
- **Returns**: Boolean indicating success
- **Example**: `success := dll.pointer.write_int(int_ptr, 42)`

### `pointer.read_bytes(pointer, size)`
Reads bytes from memory.
- **Parameters**: `pointer` - Pointer to data, `size` - Number of bytes to read
- **Returns**: Byte array
- **Example**: `data := dll.pointer.read_bytes(data_ptr, 100)`

### `pointer.write_bytes(pointer, data)`
Writes bytes to memory.
- **Parameters**: `pointer` - Pointer to write to, `data` - Byte array to write
- **Returns**: Boolean indicating success
- **Example**: `success := dll.pointer.write_bytes(data_ptr, byte_data)`

## Supported Argument Types

- **Integers**: Passed directly as 32/64-bit values
- **Floats**: Converted to integer representation (may lose precision)
- **Strings**: Converted to UTF-16 and passed as pointers
- **Bytes**: Passed as pointers to byte arrays
- **Booleans**: Converted to 0 (false) or 1 (true)
- **Null**: Passed as zero pointer
- **Pointers**: Passed directly as memory addresses



### Example Usage

#### Example DLL in C

```c
// example_dll.c

#include <stdio.h>
#include <stdlib.h>

#ifdef _WIN32
#define EXPORT __declspec(dllexport)
#else
#define EXPORT
#endif

EXPORT int add(int a, int b) {
    return a + b;
}
```

### Loading and Using the DLL in Tender

```go
// example.td
import "dll"

// Load the DLL
my_dll := dll.load("example_dll.dll")

// Get the 'add' function from the DLL
add_func := my_dll.proc("add")

// Call the 'add' function with arguments
result := add_func.call(3, 4)

// Print the result
println("Result of addition:", result)  // Output: Result of addition: 7
```


### How to compile dll

1. Compile the C code into a DLL. For example, using MinGW on Windows:

    ```
    gcc -shared -o example_dll.dll example_dll.c
    ```

2. Load and use the DLL in Tender as shown in the example above.

    ```
    tender test_dll.td
    ```
