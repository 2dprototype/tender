# Builtins

## debug

Prints debugging information about the provided arguments.

```golang
debug("Hello, World!") // Output: Hello, World!
```

## sysout

Writes the provided arguments to the system standard output.

```golang
sysout "Hello, World!", "\n" // Output: Hello, World!
```

## print

Prints the provided arguments without adding a new line.

```golang
print("Hello, World!") // Output: Hello, World!
```

## println

Prints the provided arguments followed by a new line.

```golang
println("Hello, World!") // Output: Hello, World!
```

## reverse

Reverses the provided array or string.

```golang
arr := [1, 2, 3]
revArr := reverse(arr) // revArr == [3, 2, 1]

str := "Hello"
revStr := reverse(str) // revStr == "olleH"
```

## includes

Checks if the array or string includes the provided element.

```golang
arr := [1, 2, 3]
hasTwo := includes(arr, 2) // hasTwo == true

str := "Hello"
hasE := includes(str, 'e') // hasE == false
```

## indexof

Returns the index of the first occurrence of the provided element in the array or string.

```golang
arr := [1, 2, 3]
index := indexof(arr, 2) // index == 1

str := "Hello"
index := indexof(str, 'l') // index == 2
```

## lastindexof

Returns the index of the last occurrence of the provided element in the array or string.

```golang
arr := [1, 2, 3, 2]
index := lastindexof(arr, 2) // index == 3

str := "Hello"
index := lastindexof(str, 'l') // index == 3
```

## cap

Returns the capacity of the array.

```golang
arr := [1, 2, 3]
capacity := cap(arr) // capacity == 3
```

## len

Returns the number of elements if the given variable is array, string, map, or module map.

```golang
v := [1, 2, 3]
l := len(v) // l == 3
```

## copy

Creates a copy of the given variable. `copy` function calls `Object.Copy` interface method, which is expected to return a deep-copy of the value it holds.

```golang
v1 := [1, 2, 3]
v2 := v1
v3 := copy(v1)
v1[1] = 0
print(v2[1]) // "0"; 'v1' and 'v2' referencing the same array
print(v3[1]) // "2"; 'v3' not affected by 'v1'
```

## append

Appends object(s) to an array (first argument) and returns a new array object. (Like Go's `append` builtin.) Currently, this function takes array type only.

```golang
v := [1]
v = append(v, 2, 3) // v == [1, 2, 3]
```

## delete

Deletes the element with the specified key from the map type. First argument must be a map type and second argument must be a string type. (Like Go's `delete` builtin except keys are always string). `delete` returns `null` value if successful and it mutates given map.

```golang
v := {key: "value"}
delete(v, "key") // v == {}
```

## splice

Deletes and/or changes the contents of a given array and returns deleted items as a new array. `splice` is similar to JS `Array.prototype.splice()` except splice is a builtin function and first argument must an array. First argument must be an array, and if second and third arguments are provided those must be integers otherwise runtime error is returned.

Usage:

`deleted_items := splice(array[, start[, delete_count[, item1[, item2[, ...]]]])`

```golang
v := [1, 2, 3]
items := splice(v, 0) // items == [1, 2, 3], v == []
```

## sort

Sorts the provided array.

```golang
arr := [3, 1, 2]
sortedArr := sort(arr) // sortedArr == [1, 2, 3]
```


## range

Returns a new array of int with elements from `start` to `end`.

```golang
v := range(2, 5) // v == [2, 3, 4]
```

## freeze

Recursively converts a value into its fully immutable equivalent. Arrays
become immutable arrays, maps become immutable maps, and all nested containers
are converted in the same way. Primitives, strings, bytes, functions, and
errors are returned as-is.

If the input is already fully immutable (no mutable containers anywhere in the
tree), the original value is returned without allocating anything new.

```golang
config := freeze({
    limits: {max_retries: 3, timeout: 30},
    tags:   ["prod", "v2"]
})
// config, config.limits, and config.tags are all immutable arrays/maps.
is_immutable_map(config)          // true
is_immutable_map(config.limits)   // true
is_immutable_array(config.tags)   // true
```

`freeze` is particularly useful before sharing globals across concurrent
`Clone()` executions: a frozen value is shared by all clones at zero copy
cost because no clone can mutate it.

```golang
// Immutable input with no mutable children is returned as the same object.
a := immutable([1, 2, 3])
b := freeze(a)
// b == a 
```

## format

Returns a formatted string. The first argument must be a String object. See
[this](https://github.com/2dprototype/tender-free/blob/master/docs/formatting.md) for more
details on formatting.

```golang
a := [1, 2, 3]
s := format("Foo: %v", a) // s == "Foo: [1, 2, 3]"
```

## typeof

Returns the type name of an object.

```golang
typeof(1) // int
typeof("str") // string
typeof([1, 2, 3]) // array
```

## rune

Converts the provided rune to a integer.

```golang
r := rune('a') // r == 97
```

## string

Tries to convert an object to string object. See [Runtime Types](https://github.com/2dprototype/tender-free/blob/master/docs/runtime-types.md) for more details on type conversion.

```golang
x := string(123) //  x == "123"
```

## int

Tries to convert an object to int object. See [this](https://github.com/2dprototype/tender-free/blob/master/docs/runtime-types.md) for more details on type conversion.

```golang
v := int("123") //  v == 123
```

Optionally it can take the second argument, which will be returned if the first argument cannot be converted to int. Note that the second argument does not have to be int.

```golang
v = int(null, 10)    // v == 10
v = int(null, false) // v == false
```

## bool

Tries to convert an object to bool object. See [this](https://github.com/2dprototype/tender-free/blob/master/docs/runtime-types.md) for more details on type conversion.

```golang
v := bool(1) //  v == true
```

## float

Tries to convert an object to float object. See [this](https://github.com/2dprototype/tender-free/blob/master/docs/runtime-types.md) for more details on type conversion.

```golang
v := float("19.84") //  v == 19.84
```

Optionally it can take the second argument, which will be returned if the first argument cannot be converted to float. Note that the second argument does not have to be float.

```golang
v = float(null, 19.84)    // v == 19.84
v = float(null, false)    // v == false
```

## char

Tries to convert an object to char object. See [this](https://github.com/2dprototype/tender-free/blob/master/docs/runtime-types.md) for more details on type conversion.

```golang
v := char(89) //  v == 'Y'
```

Optionally it can take the second argument, which will be returned if the first argument cannot be converted to float. Note that the second argument does not have to be float.

```golang
v = char(null, 'X')    // v == 'X'
v = char(null, false)  // v == false
```

## bytes

Tries to convert an object to bytes object. See [this](https://github.com/2dprototype/tender-free/blob/master/docs/runtime-types.md) for more details on type conversion.

```golang
v := bytes("foo") //  v == [102 111 111]
```

If you pass an int to `bytes()` function, it will create a new byte object with the given size.

```golang
v := bytes(100)
```

## time

Tries to convert an object to time value.

```golang
v := time(1257894000) // 2009-11-10 23:00:00 +0000 UTC
```

## is_cycle

Returns `true` if the object's type is cycle. Or it returns `false`.

## is_int

Returns `true` if the object's type is int. Or it returns `false`.

## is_bigint

Returns `true` if the object's type is bigint. Or it returns `false`.

## is_float

Returns `true` if the object's type is float. Or it returns `false`.

## is_bigfloat

Returns `true` if the object's type is bigfloat. Or it returns `false`.

## is_string

Returns `true` if the object's type is string. Or it returns `false`.

## is_bool

Returns `true` if the object's type is bool. Or it returns `false`.

## is_char

Returns `true` if the object's type is char. Or it returns `false`.

## is_bytes

Returns `true` if the object's type is bytes. Or it returns `false`.

## is_array

Returns `true` if the object's type is array. Or it returns `false`.

## is_immutable_array

Returns `true` if the object's type is immutable array. Or it returns `false`.

## is_map

Returns `true` if the object's type is map. Or it returns `false`.

## is_immutable_map

immutable_map

Returns `true` if the object's type is immutable map. Or it returns `false`.

## is_time

Returns `true` if the object's type is time. Or it returns `false`.

## is_function

Returns `true` if the object's type is function. Or it returns `false`.

## is_null

Returns `true` if the object's type is null. Or it returns `false`.

## is_error

Returns `true` if the object's type is error. Or it returns `false`.

## is_iterable

Returns `true` if the object's type is iterable. Or it returns `false`.