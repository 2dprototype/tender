## Stdlib `image`

The `image` module provides comprehensive functionalities for working with images, including loading, decoding, creating new images, applying filters, and encoding images into various formats.

### Module Functions

#### `new(width, height)`
Creates a new image with the specified width and height.
- `width`: Width of the new image (integer)
- `height`: Height of the new image (integer)

#### `load(path)`
Loads an image from the specified file path.
- `path`: Path to the image file (string)

#### `decode(image_data)`
Decodes image data into an image object.
- `image_data`: Byte array containing image data (bytes)

#### `formats`
Array containing supported image formats: `["png", "jpeg", "bmp", "tiff", "webp"]`

### Image Object Functions

#### `encode(format)`
Encodes the image into the specified format and returns byte data.
- `format`: Format to encode the image ("png", "jpeg", "bmp", "tiff")

#### `bounds()`
Returns the bounding rectangle of the image as an object with:
- `min`: Object with `x`, `y` coordinates
- `max`: Object with `x`, `y` coordinates  
- `size`: Object with `width`, `height` dimensions

#### `at(x, y)`
Returns the color of the pixel at the specified coordinates as RGBA array.
- `x`: X-coordinate of the pixel (integer)
- `y`: Y-coordinate of the pixel (integer)

#### `pixels()`
Returns the total number of pixels in the image (integer).

#### `get_pixels()`
Returns a flat array containing all pixel values of the image.

#### `set_pixels(pixels)`
Sets the pixel values of the image from the given array.
- `pixels`: Array containing pixel values (array)

#### `set(x, y, color)`
Sets the color of the pixel at the specified coordinates.
- `x`: X-coordinate of the pixel (integer)
- `y`: Y-coordinate of the pixel (integer)
- `color`: Array containing RGBA values `[red, green, blue, alpha]` (array)

#### `save(path, format)`
Saves the image to the specified file path and format.
- `path`: Path to save the image (string)
- `format`: Format to save the image ("png", "jpeg", "bmp", "tiff")

### Image Filters

The `filters` object provides various image filter operations that return new filtered images:

#### `blur(radius)`
Applies Gaussian blur with specified radius.
- `radius`: Blur radius (integer)

#### `bnw(threshold)`
Converts image to black and white using luminance threshold.
- `threshold`: Luminance threshold (0-255)

#### `glitch(max_shift)`
Applies glitch effect by randomly shifting RGB channels.
- `max_shift`: Maximum pixel shift amount (integer)

#### `invert()`
Inverts all colors in the image.

#### `grayscale()`
Converts image to grayscale.

#### `sepia()`
Applies sepia tone filter.

#### `brightness(offset)`
Adjusts brightness by adding offset to each channel.
- `offset`: Brightness adjustment (-255 to 255)

#### `contrast(factor)`
Adjusts contrast by multiplying difference from middle gray.
- `factor`: Contrast multiplier (float)

#### `saturation(factor)`
Adjusts color saturation.
- `factor`: Saturation multiplier (0 = grayscale, 1 = original)

#### `sharpen()`
Applies sharpening filter.

#### `emboss()`
Applies emboss effect.

#### `edge()`
Applies edge detection filter.

#### `hue(angle)`
Adjusts hue by specified angle.
- `angle`: Hue angle in degrees (float)

#### `temperature(adjustment)`
Adjusts color temperature.
- `adjustment`: Temperature adjustment (positive = warmer, negative = cooler)

#### `vignette(intensity)`
Applies vignette (darkened edges) effect.
- `intensity`: Vignette strength (float)

#### `pixelate(size)`
Creates pixelated effect with specified block size.
- `size`: Pixel block size (integer)

#### `sobel()`
Applies Sobel edge detection operator.

### Example Usage

```go
import "image"

// Create a new image
img := image.new(255, 255)

for i := 0; i < 255; i++ {
	for j := 0; j < 255; j++ {
		color := [i, 0, j, 255]
		img.set(i, j, color)
	}
}

img.save("out.png", "png")
```


### Color Representation

Colors are represented as arrays of 4 integers: `[red, green, blue, alpha]` where each component ranges from 0 to 255.

### Performance Notes

- All filter operations are parallelized for optimal performance
- Filters return new image objects and do not modify the original
- Large images may require significant memory and processing time
- The module maintains image data in RGBA format internally

### Supported Formats

- **PNG**: Lossless compression with transparency support
- **JPEG**: Lossy compression suitable for photographs  
- **BMP**: Uncompressed bitmap format
- **TIFF**: High-quality format with various compression options
- **WebP**: Modern format with both lossy and lossless compression
