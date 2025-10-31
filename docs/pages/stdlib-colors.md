## Stdlib `colors`

The `colors` module provides functionality for printing colored and styled text to the terminal. It includes utilities for writing to standard output and error, as well as applying rich text styles.

## Functions

### `colors.stdout() → IOWriter`

Returns an `IOWriter` that writes to **standard output** (`stdout`) with color support.

```go
import "colors"

stdout := colors.stdout()
```

---

### `colors.stderr() → IOWriter`

Returns an `IOWriter` that writes to **standard error** (`stderr`) with color support.

```go
import "colors"

stderr := colors.stderr()
```

---

### `colors.style(text: string, ...props: map) → string`

Applies styles to a string and returns a formatted string.

#### Parameters:

* `text` — The string to style.
* `props` — Optional style properties as a map. Supported keys:

  * `"color"` — Text color (string, e.g., `"red"`, `"green"`, `"#ff00ff"`).
  * `"background"` — Background color (string).
  * `"bold"` — Boolean, makes text bold.
  * `"italic"` — Boolean, makes text italic.
  * `"underline"` — Boolean, underlines text.
  * `"strikethrough"` — Boolean, strikes through text.
  * `"width"` — Integer, sets text width.
  * `"height"` — Integer, sets text height.
  * `"align"` — Text alignment: `"left"`, `"center"`, `"right"`.
  * `"border"` — Border style: `"normal"`, `"rounded"`, `"thick"`, `"double"`.
  * `"border_top"` — Boolean, adds a top border.
  * `"margin"` — Integer, sets margin.
  * `"padding"` — Integer, sets padding.

#### Example Usage:

```go
import "fmt"
import "colors"

fmt.fprint(colors.stdout(), "Hello".red, "World".green, "\n")
fmt.fprintln(colors.stderr(), "Hello".red, "World".green)

fmt.fprint(colors.stdout(), colors.style("Hello", {
    "color": "#ff0000",
    "bold": true,
    "underline": true
}), "\n")

fmt.fprintln(colors.stderr(), colors.style("Error!", {
    "color": "#00ff00",
    "background": "#232323",
    "bold": true
}))
```