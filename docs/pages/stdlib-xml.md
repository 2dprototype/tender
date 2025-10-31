## Stdlib `xml`

The `xml` module provides comprehensive functionality for parsing, generating, and manipulating XML data. It supports both document-style and data-oriented XML processing with flexible object mapping.

## Core Functions

### `decode(xml_string)`
Parses XML string into Tender objects using the `@` prefix for attributes and `#` key for text content.
- **Parameters**: `xml_string` - XML data as string or bytes
- **Returns**: Tender object representing the XML structure
- **Example**: `data := xml.decode("<root><item id='1'>Text</item></root>")`

### `encode(object)`
Converts Tender objects into XML string using the `@` prefix for attributes and `#` key for text content.
- **Parameters**: `object` - Tender object to encode as XML
- **Returns**: XML string
- **Example**: `xml_string := xml.encode(data)`

### `escape(text)`
Escapes XML special characters in a string.
- **Parameters**: `text` - String to escape
- **Returns**: Escaped string
- **Example**: `escaped := xml.escape("AT&T <special>")`

### `unescape(text)`
Unescapes XML entities back to their original characters.
- **Parameters**: `text` - String with XML entities
- **Returns**: Unescaped string
- **Example**: `original := xml.unescape("AT&amp;T &lt;special&gt;")`

## XML Structure Convention

The module uses a specific convention for representing XML in Tender objects:

- **Element names**: Used as object keys
- **Attributes**: Prefixed with `@` (e.g., `{"@id": "1"}`)
- **Text content**: Stored under `#` key
- **Multiple elements**: Become arrays
- **Mixed content**: Text and elements stored together

## Examples

### Basic XML Parsing

```javascript
import "xml"

// Simple XML parsing
xml_data := `
<book id="123">
    <title>XML Guide</title>
    <author>John Doe</author>
    <price currency="USD">29.99</price>
</book>
`

// Parse XML
book := xml.decode(xml_data)
println("Book title:", book.book.title)
println("Book ID:", book.book["@id"])
println("Price:", book.book.price["#"])
println("Currency:", book.book.price["@currency"])

// Output structure:
// {
//   "book": {
//     "@id": "123",
//     "title": "XML Guide",
//     "author": "John Doe", 
//     "price": {
//       "@currency": "USD",
//       "#": "29.99"
//     }
//   }
// }
```

### XML Generation

```javascript
import "xml"

// Create data structure for XML
data := {
    "catalog": {
        "@version": "1.0",
        "book": [
            {
                "@id": "1",
                "title": "First Book",
                "author": "Author One",
                "price": {
                    "@currency": "USD",
                    "#": "19.99"
                }
            },
            {
                "@id": "2", 
                "title": "Second Book",
                "author": "Author Two",
                "price": {
                    "@currency": "EUR", 
                    "#": "17.50"
                }
            }
        ]
    }
}

// Generate XML
xml_output := xml.encode(data)
println(string(xml_output))

// Output:
// <catalog version="1.0">
//   <book id="1">
//     <title>First Book</title>
//     <author>Author One</author>
//     <price currency="USD">19.99</price>
//   </book>
//   <book id="2">
//     <title>Second Book</title>
//     <author>Author Two</author>
//     <price currency="EUR">17.50</price>
//   </book>
// </catalog>
```