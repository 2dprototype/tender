## Stdlib `path`

The `path` module provides functions for manipulating file paths and performing operations related to file paths.

### Functions

#### `join(path1, path2, ...)`

Joins any number of path elements into a single path, separating them with the operating system-specific separator.

- `path1, path2, ...`: Multiple `string` objects representing the path elements to be joined.

Returns a `string` containing the joined path.

#### `base(path)`

Returns the last element of the path, typically the file or directory name.

- `path`: A `string` representing the path.

Returns a `string` containing the base name.

#### `ext(path)`

Returns the file extension of the given path.

- `path`: A `string` representing the path.

Returns a `string` containing the file extension.

#### `clean(path)`

Returns the cleaned version of the path.

- `path`: A `string` representing the path to be cleaned.

Returns a `string` containing the cleaned path.

#### `dir(path)`

Returns the directory part of the given path.

- `path`: A `string` representing the path.

Returns a `string` containing the directory part.

#### `isabs(path)`

Checks whether the given path is absolute.

- `path`: A `string` representing the path.

Returns a `bool` indicating whether the path is absolute.

#### `abs(path)`

Returns the absolute path of the given path.

- `path`: A `string` representing the path.

Returns a `string` containing the absolute path.

#### `to_slash(path)`

Converts the path to use forward slashes ('/') as the separator.

- `path`: A `string` representing the path.

Returns a `string` with forward slashes.

#### `from_slash(path)`

Converts the path to use the native operating system separator.

- `path`: A `string` representing the path.

Returns a `string` with the native separator.

#### `vol(path)`

Returns the volume name of the given path.

- `path`: A `string` representing the path.

Returns a `string` containing the volume name.

#### `walklist(root)`

Walks the file tree rooted at the specified root path and returns a list of all visited files and directories.

- `root`: A `string` representing the root directory to start the walk.

Returns a `array` containing a list of paths visited during the walk.

#### `splitlist(paths)`

Splits the input string containing a list of paths into individual paths.

- `paths`: A `string` representing the list of paths.

Returns a `array` containing individual path elements.

### Example Usage

```go
import "path"

// Join path elements
joined_path := path.join("folder", "hello.txt")

println(joined_path)
```