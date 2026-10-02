package stdlib

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	jpeg "image/jpeg"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/2dprototype/tender"
	bmp "golang.org/x/image/bmp"
	xdraw "golang.org/x/image/draw"
	tiff "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

var imageModule = map[string]tender.Object{
	"new":        &tender.NativeFunction{Value: imageNew},
	"load":       &tender.NativeFunction{Value: imageLoad},
	"decode":     &tender.NativeFunction{Value: imageDecode},
	"from_bytes": &tender.NativeFunction{Value: imageFromBytes},
	"dimensions": &tender.NativeFunction{Value: imageDimensions},
	"save":       &tender.NativeFunction{Value: imageModuleSave},
	"formats": &tender.ImmutableArray{Value: []tender.Object{
		&tender.String{Value: "png"},
		&tender.String{Value: "jpeg"},
		&tender.String{Value: "jpg"},
		&tender.String{Value: "bmp"},
		&tender.String{Value: "tiff"},
		&tender.String{Value: "webp"},
	}},
}

func normalizeFormat(format string) string {
	f := strings.ToLower(strings.TrimSpace(format))
	f = strings.TrimPrefix(f, ".")
	switch f {
	case "jpg", "jpeg":
		return "jpeg"
	case "png":
		return "png"
	case "bmp":
		return "bmp"
	case "tiff", "tif":
		return "tiff"
	}
	return f
}

func detectFormatFromPath(path string) string {
	ext := filepath.Ext(path)
	return normalizeFormat(ext)
}

func encodeImage(img image.Image, format string, quality int) ([]byte, error) {
	fmtNorm := normalizeFormat(format)
	var buf bytes.Buffer
	switch fmtNorm {
	case "png":
		if err := png.Encode(&buf, img); err != nil {
			return nil, err
		}
	case "jpeg":
		opt := &jpeg.Options{Quality: 85}
		if quality > 0 && quality <= 100 {
			opt.Quality = quality
		}
		if err := jpeg.Encode(&buf, img, opt); err != nil {
			return nil, err
		}
	case "tiff":
		if err := tiff.Encode(&buf, img, nil); err != nil {
			return nil, err
		}
	case "bmp":
		if err := bmp.Encode(&buf, img); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}
	return buf.Bytes(), nil
}

func imageDimensions(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	b, err := ToFileData(args[0])
	if err != nil {
		return nil, tender.ErrInvalidArgumentType{Name: "src", Expected: "string or bytes"}
	}

	cfg, format, err := image.DecodeConfig(bytes.NewReader(b))
	if err != nil {
		return wrapError(err), nil
	}

	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"width":  &tender.Int{Value: int64(cfg.Width)},
			"height": &tender.Int{Value: int64(cfg.Height)},
			"format": &tender.String{Value: format},
		},
	}, nil
}

func imageFromBytes(args ...tender.Object) (tender.Object, error) {
	if len(args) != 3 {
		return nil, tender.ErrWrongNumArguments
	}
	data, ok := tender.ToByteSlice(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{Name: "bytes", Expected: "bytes"}
	}
	w, ok1 := tender.ToInt(args[1])
	h, ok2 := tender.ToInt(args[2])
	if !ok1 || !ok2 {
		return nil, tender.ErrInvalidArgumentType{Name: "width/height", Expected: "int"}
	}
	expectedLen := w * h * 4
	if len(data) < expectedLen {
		return nil, fmt.Errorf("insufficient byte length: got %d, expected %d for %dx%d RGBA", len(data), expectedLen, w, h)
	}
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	copy(rgba.Pix, data[:expectedLen])
	return makeImage(rgba), nil
}

func imageModuleSave(args ...tender.Object) (tender.Object, error) {
	if len(args) < 2 {
		return nil, tender.ErrWrongNumArguments
	}
	imgObj, ok := args[0].(*tender.ImmutableMap)
	if !ok {
		return nil, tender.ErrInvalidArgumentType{Name: "image", Expected: "image"}
	}
	saveFn, ok := imgObj.Value["save"].(*tender.NativeFunction)
	if !ok {
		return nil, fmt.Errorf("invalid image object: missing save method")
	}
	return saveFn.Value(args[1:]...)
}

func imageDecode(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	imageBytes, ok := tender.ToByteSlice(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	buffer := bytes.NewBuffer(imageBytes)

	img, _, err := image.Decode(buffer)
	if err != nil {
		return wrapError(err), nil
	}

	return makeImage(img), nil
}

func imageLoad(args ...tender.Object) (tender.Object, error) {
	if len(args) != 1 {
		return nil, tender.ErrWrongNumArguments
	}

	path, ok := tender.ToString(args[0])
	if !ok {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "path",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}

	file, err := os.Open(tender.ResolvePath(path))
	if err != nil {
		return wrapError(err), nil
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return wrapError(err), nil
	}

	return makeImage(img), nil
}

func imageNew(args ...tender.Object) (ret tender.Object, err error) {
	if len(args) != 2 {
		return nil, tender.ErrWrongNumArguments
	}

	width, ok1 := tender.ToInt(args[0])
	height, ok2 := tender.ToInt(args[1])

	if !ok1 || !ok2 {
		return nil, tender.ErrInvalidArgumentType{
			Name:     "width/height",
			Expected: "int",
			Found:    args[0].TypeName(),
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	return makeImage(img), nil
}

func parseColorArg(obj tender.Object) (color.RGBA, bool) {
	if arr, ok := obj.(*tender.Array); ok {
		if len(arr.Value) >= 3 {
			r, _ := tender.ToInt(arr.Value[0])
			g, _ := tender.ToInt(arr.Value[1])
			b, _ := tender.ToInt(arr.Value[2])
			a := 255
			if len(arr.Value) >= 4 {
				a, _ = tender.ToInt(arr.Value[3])
			}
			return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}, true
		}
	}
	if str, ok := tender.ToString(obj); ok {
		hex := strings.TrimPrefix(str, "#")
		var r, g, b, a uint32 = 0, 0, 0, 255
		if len(hex) == 3 {
			fmt.Sscanf(hex, "%1x%1x%1x", &r, &g, &b)
			r |= r << 4
			g |= g << 4
			b |= b << 4
		} else if len(hex) == 4 {
			fmt.Sscanf(hex, "%1x%1x%1x%1x", &r, &g, &b, &a)
			r |= r << 4
			g |= g << 4
			b |= b << 4
			a |= a << 4
		} else if len(hex) == 6 {
			fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
		} else if len(hex) == 8 {
			fmt.Sscanf(hex, "%02x%02x%02x%02x", &r, &g, &b, &a)
		} else {
			return color.RGBA{}, false
		}
		return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}, true
	}
	return color.RGBA{}, false
}

func cropImage(rgba *image.RGBA, x, y, w, h int) *image.RGBA {
	bounds := rgba.Bounds()
	reqRect := image.Rect(x, y, x+w, y+h)
	intersect := reqRect.Intersect(bounds)
	if intersect.Empty() {
		return image.NewRGBA(image.Rect(0, 0, w, h))
	}
	dst := image.NewRGBA(image.Rect(0, 0, intersect.Dx(), intersect.Dy()))
	draw.Draw(dst, dst.Bounds(), rgba, intersect.Min, draw.Src)
	return dst
}

func resizeImage(src image.Image, w, h int, filter string) *image.RGBA {
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	var scaler xdraw.Scaler = xdraw.BiLinear
	switch strings.ToLower(strings.TrimSpace(filter)) {
	case "nearest", "point":
		scaler = xdraw.NearestNeighbor
	case "approx_bilinear":
		scaler = xdraw.ApproxBiLinear
	case "catmull_rom", "bicubic":
		scaler = xdraw.CatmullRom
	}
	scaler.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)
	return dst
}

func flipHorizontal(src *image.RGBA) *image.RGBA {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := src.RGBAAt(bounds.Min.X+x, bounds.Min.Y+y)
			dst.SetRGBA(w-1-x, y, c)
		}
	}
	return dst
}

func flipVertical(src *image.RGBA) *image.RGBA {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := src.RGBAAt(bounds.Min.X+x, bounds.Min.Y+y)
			dst.SetRGBA(x, h-1-y, c)
		}
	}
	return dst
}

func rotate90(src *image.RGBA) *image.RGBA {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, h, w))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := src.RGBAAt(bounds.Min.X+x, bounds.Min.Y+y)
			dst.SetRGBA(h-1-y, x, c)
		}
	}
	return dst
}

func rotate180(src *image.RGBA) *image.RGBA {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := src.RGBAAt(bounds.Min.X+x, bounds.Min.Y+y)
			dst.SetRGBA(w-1-x, h-1-y, c)
		}
	}
	return dst
}

func rotate270(src *image.RGBA) *image.RGBA {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, h, w))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := src.RGBAAt(bounds.Min.X+x, bounds.Min.Y+y)
			dst.SetRGBA(y, w-1-x, c)
		}
	}
	return dst
}

func rotateAngle(src *image.RGBA, angleRad float64) *image.RGBA {
	bounds := src.Bounds()
	w, h := float64(bounds.Dx()), float64(bounds.Dy())
	sin, cos := math.Abs(math.Sin(angleRad)), math.Abs(math.Cos(angleRad))
	newW := int(math.Ceil(w*cos + h*sin))
	newH := int(math.Ceil(w*sin + h*cos))
	dst := image.NewRGBA(image.Rect(0, 0, newW, newH))

	cx, cy := w/2.0, h/2.0
	ncx, ncy := float64(newW)/2.0, float64(newH)/2.0
	c, s := math.Cos(-angleRad), math.Sin(-angleRad)

	for dy := 0; dy < newH; dy++ {
		for dx := 0; dx < newW; dx++ {
			ox := float64(dx) - ncx
			oy := float64(dy) - ncy
			sx := ox*c - oy*s + cx
			sy := ox*s + oy*c + cy
			isx := int(math.Round(sx))
			isy := int(math.Round(sy))
			if isx >= 0 && isx < int(w) && isy >= 0 && isy < int(h) {
				dst.SetRGBA(dx, dy, src.RGBAAt(bounds.Min.X+isx, bounds.Min.Y+isy))
			}
		}
	}
	return dst
}

func cloneRGBA(src *image.RGBA) *image.RGBA {
	dst := image.NewRGBA(src.Bounds())
	copy(dst.Pix, src.Pix)
	return dst
}

func compositeImage(base *image.RGBA, overlay image.Image, x, y int, op draw.Op) *image.RGBA {
	dst := cloneRGBA(base)
	oBounds := overlay.Bounds()
	rect := image.Rect(x, y, x+oBounds.Dx(), y+oBounds.Dy())
	draw.Draw(dst, rect, overlay, oBounds.Min, op)
	return dst
}

func makeImage(img image.Image) *tender.ImmutableMap {
	// Convert the image to *image.RGBA if it's not already
	var rgbaImage *image.RGBA
	if r, ok := img.(*image.RGBA); ok {
		rgbaImage = r
	} else {
		bounds := img.Bounds()
		rgbaImage = image.NewRGBA(bounds)
		draw.Draw(rgbaImage, bounds, img, bounds.Min, draw.Src)
		img = rgbaImage
	}

	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"filters": makeImageFilters(rgbaImage),
			"width": &tender.NativeFunction{
				Name: "width",
				Value: func(args ...tender.Object) (tender.Object, error) {
					return &tender.Int{Value: int64(rgbaImage.Bounds().Dx())}, nil
				},
			},
			"height": &tender.NativeFunction{
				Name: "height",
				Value: func(args ...tender.Object) (tender.Object, error) {
					return &tender.Int{Value: int64(rgbaImage.Bounds().Dy())}, nil
				},
			},
			"size": &tender.NativeFunction{
				Name: "size",
				Value: func(args ...tender.Object) (tender.Object, error) {
					return &tender.ImmutableMap{
						Value: map[string]tender.Object{
							"width":  &tender.Int{Value: int64(rgbaImage.Bounds().Dx())},
							"height": &tender.Int{Value: int64(rgbaImage.Bounds().Dy())},
						},
					}, nil
				},
			},
			"crop": &tender.NativeFunction{
				Name: "crop",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 4 {
						return nil, tender.ErrWrongNumArguments
					}
					x, _ := tender.ToInt(args[0])
					y, _ := tender.ToInt(args[1])
					w, _ := tender.ToInt(args[2])
					h, _ := tender.ToInt(args[3])
					return makeImage(cropImage(rgbaImage, x, y, w, h)), nil
				},
			},
			"sub_image": &tender.NativeFunction{
				Name: "sub_image",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 4 {
						return nil, tender.ErrWrongNumArguments
					}
					x, _ := tender.ToInt(args[0])
					y, _ := tender.ToInt(args[1])
					w, _ := tender.ToInt(args[2])
					h, _ := tender.ToInt(args[3])
					return makeImage(cropImage(rgbaImage, x, y, w, h)), nil
				},
			},
			"resize": &tender.NativeFunction{
				Name: "resize",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) < 2 {
						return nil, tender.ErrWrongNumArguments
					}
					w, _ := tender.ToInt(args[0])
					h, _ := tender.ToInt(args[1])
					filter := "bilinear"
					if len(args) >= 3 {
						filter, _ = tender.ToString(args[2])
					}
					return makeImage(resizeImage(rgbaImage, w, h, filter)), nil
				},
			},
			"scale": &tender.NativeFunction{
				Name: "scale",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) < 1 {
						return nil, tender.ErrWrongNumArguments
					}
					factor, _ := tender.ToFloat64(args[0])
					if factor <= 0 {
						return nil, fmt.Errorf("scale factor must be positive, got %f", factor)
					}
					w := int(math.Round(float64(rgbaImage.Bounds().Dx()) * factor))
					h := int(math.Round(float64(rgbaImage.Bounds().Dy()) * factor))
					if w < 1 {
						w = 1
					}
					if h < 1 {
						h = 1
					}
					filter := "bilinear"
					if len(args) >= 2 {
						filter, _ = tender.ToString(args[1])
					}
					return makeImage(resizeImage(rgbaImage, w, h, filter)), nil
				},
			},
			"flip_h": &tender.NativeFunction{
				Name: "flip_h",
				Value: func(args ...tender.Object) (tender.Object, error) {
					return makeImage(flipHorizontal(rgbaImage)), nil
				},
			},
			"flip_v": &tender.NativeFunction{
				Name: "flip_v",
				Value: func(args ...tender.Object) (tender.Object, error) {
					return makeImage(flipVertical(rgbaImage)), nil
				},
			},
			"rotate90": &tender.NativeFunction{
				Name: "rotate90",
				Value: func(args ...tender.Object) (tender.Object, error) {
					return makeImage(rotate90(rgbaImage)), nil
				},
			},
			"rotate180": &tender.NativeFunction{
				Name: "rotate180",
				Value: func(args ...tender.Object) (tender.Object, error) {
					return makeImage(rotate180(rgbaImage)), nil
				},
			},
			"rotate270": &tender.NativeFunction{
				Name: "rotate270",
				Value: func(args ...tender.Object) (tender.Object, error) {
					return makeImage(rotate270(rgbaImage)), nil
				},
			},
			"rotate": &tender.NativeFunction{
				Name: "rotate",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					deg, _ := tender.ToFloat64(args[0])
					return makeImage(rotateAngle(rgbaImage, deg*math.Pi/180.0)), nil
				},
			},
			"clone": &tender.NativeFunction{
				Name: "clone",
				Value: func(args ...tender.Object) (tender.Object, error) {
					return makeImage(cloneRGBA(rgbaImage)), nil
				},
			},
			"copy": &tender.NativeFunction{
				Name: "copy",
				Value: func(args ...tender.Object) (tender.Object, error) {
					return makeImage(cloneRGBA(rgbaImage)), nil
				},
			},
			"composite": &tender.NativeFunction{
				Name: "composite",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) < 3 {
						return nil, tender.ErrWrongNumArguments
					}
					overlay, err := decodeImageArg(args[0])
					if err != nil {
						return wrapError(err), nil
					}
					x, _ := tender.ToInt(args[1])
					y, _ := tender.ToInt(args[2])
					op := draw.Over
					if len(args) >= 4 {
						if opStr, ok := tender.ToString(args[3]); ok && strings.ToLower(opStr) == "src" {
							op = draw.Src
						}
					}
					res := compositeImage(rgbaImage, overlay, x, y, op)
					return makeImage(res), nil
				},
			},
			"fill": &tender.NativeFunction{
				Name: "fill",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) < 1 {
						return nil, tender.ErrWrongNumArguments
					}
					c, ok := parseColorArg(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{Name: "color", Expected: "array or hex string"}
					}
					bounds := rgbaImage.Bounds()
					fillBounds := bounds
					if len(args) >= 5 {
						x, _ := tender.ToInt(args[1])
						y, _ := tender.ToInt(args[2])
						w, _ := tender.ToInt(args[3])
						h, _ := tender.ToInt(args[4])
						fillBounds = image.Rect(x, y, x+w, y+h).Intersect(bounds)
					}
					draw.Draw(rgbaImage, fillBounds, &image.Uniform{C: c}, image.Point{}, draw.Src)
					return tender.NullValue, nil
				},
			},
			"clear": &tender.NativeFunction{
				Name: "clear",
				Value: func(args ...tender.Object) (tender.Object, error) {
					c := color.RGBA{0, 0, 0, 0}
					if len(args) >= 1 {
						if parsed, ok := parseColorArg(args[0]); ok {
							c = parsed
						}
					}
					draw.Draw(rgbaImage, rgbaImage.Bounds(), &image.Uniform{C: c}, image.Point{}, draw.Src)
					return tender.NullValue, nil
				},
			},
			"raw_bytes": &tender.NativeFunction{
				Name: "raw_bytes",
				Value: func(args ...tender.Object) (tender.Object, error) {
					return &tender.Bytes{Value: append([]byte(nil), rgbaImage.Pix...)}, nil
				},
			},
			"bytes": &tender.NativeFunction{
				Name: "bytes",
				Value: func(args ...tender.Object) (tender.Object, error) {
					return &tender.Bytes{Value: append([]byte(nil), rgbaImage.Pix...)}, nil
				},
			},
			"encode": &tender.NativeFunction{
				Name: "encode",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) < 1 {
						return nil, tender.ErrWrongNumArguments
					}
					format, ok := tender.ToString(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string",
							Found:    args[0].TypeName(),
						}
					}
					quality := 85
					if len(args) >= 2 {
						if q, ok := tender.ToInt(args[1]); ok {
							quality = q
						}
					}
					data, err := encodeImage(rgbaImage, format, quality)
					if err != nil {
						return wrapError(err), nil
					}
					return &tender.Bytes{Value: data}, nil
				},
			},
			"bounds": &tender.NativeFunction{
				Name: "bounds",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					rect := rgbaImage.Bounds()
					return makeRectangle(rect), nil
				},
			},
			"at": &tender.NativeFunction{
				Name: "at",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 {
						return nil, tender.ErrWrongNumArguments
					}

					x, ok1 := tender.ToInt(args[0])
					y, ok2 := tender.ToInt(args[1])

					if !ok1 || !ok2 {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "x/y",
							Expected: "int",
							Found:    args[0].TypeName(),
						}
					}

					color := rgbaImage.At(x, y)
					return makeColor(color), nil
				},
			},
			"pixels": &tender.NativeFunction{
				Name: "pixels",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					bounds := rgbaImage.Bounds()
					return &tender.Int{Value: int64((bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y))}, nil
				},
			},
			"get_pixels": &tender.NativeFunction{
				Name: "get_pixels",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}

					pixels := make([]tender.Object, len(rgbaImage.Pix))
					for i, p := range rgbaImage.Pix {
						pixels[i] = &tender.Int{Value: int64(p)}
					}

					return &tender.Array{Value: pixels}, nil
				},
			},
			"set_pixels": &tender.NativeFunction{
				Name: "set_pixels",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}

					pixelArray, ok := args[0].(*tender.Array)
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "pixels",
							Expected: "array",
							Found:    args[0].TypeName(),
						}
					}

					if len(pixelArray.Value) > len(rgbaImage.Pix) {
						return &tender.Error{Value: &tender.String{Value: "Failed to set pixels: Length of pixel array is greater than image dimensions"}}, nil
					}

					for i, pixel := range pixelArray.Value {
						val, _ := tender.ToInt(pixel)
						rgbaImage.Pix[i] = uint8(val)
					}

					return tender.NullValue, nil
				},
			},
			"set": &tender.NativeFunction{
				Name: "set",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 3 {
						return nil, tender.ErrWrongNumArguments
					}

					x, ok1 := tender.ToInt(args[0])
					y, ok2 := tender.ToInt(args[1])

					if !ok1 || !ok2 {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "x/y",
							Expected: "int",
							Found:    args[0].TypeName(),
						}
					}

					col, ok := parseColorArg(args[2])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "color",
							Expected: "array or hex string",
							Found:    args[2].TypeName(),
						}
					}

					rgbaImage.Set(x, y, col)
					return tender.NullValue, nil
				},
			},
			"save": &tender.NativeFunction{
				Name: "save",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) < 1 {
						return nil, tender.ErrWrongNumArguments
					}

					path, ok := tender.ToString(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "path",
							Expected: "string",
							Found:    args[0].TypeName(),
						}
					}
					format := ""
					if len(args) >= 2 {
						format, _ = tender.ToString(args[1])
					}
					if format == "" {
						format = detectFormatFromPath(path)
						if format == "" {
							format = "png"
						}
					}
					quality := 85
					if len(args) >= 3 {
						if q, ok := tender.ToInt(args[2]); ok {
							quality = q
						}
					}

					data, err := encodeImage(rgbaImage, format, quality)
					if err != nil {
						return wrapError(err), nil
					}

					if err := os.WriteFile(tender.ResolvePath(path), data, 0666); err != nil {
						return wrapError(err), nil
					}

					return tender.NullValue, nil
				},
			},

			// Channel getters
			"get_red": &tender.NativeFunction{
				Name: "get_red",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}

					bounds := rgbaImage.Bounds()
					width := bounds.Dx()
					height := bounds.Dy()

					data := make([]int64, width*height)
					for y := 0; y < height; y++ {
						for x := 0; x < width; x++ {
							idx := y*width + x
							data[idx] = int64(rgbaImage.RGBAAt(x, y).R)
						}
					}

					return &tender.Matrix[int64]{
						Rows: height,
						Cols: width,
						Data: data,
					}, nil
				},
			},
			"get_green": &tender.NativeFunction{
				Name: "get_green",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}

					bounds := rgbaImage.Bounds()
					width := bounds.Dx()
					height := bounds.Dy()

					data := make([]int64, width*height)
					for y := 0; y < height; y++ {
						for x := 0; x < width; x++ {
							idx := y*width + x
							data[idx] = int64(rgbaImage.RGBAAt(x, y).G)
						}
					}

					return &tender.Matrix[int64]{
						Rows: height,
						Cols: width,
						Data: data,
					}, nil
				},
			},
			"get_blue": &tender.NativeFunction{
				Name: "get_blue",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}

					bounds := rgbaImage.Bounds()
					width := bounds.Dx()
					height := bounds.Dy()

					data := make([]int64, width*height)
					for y := 0; y < height; y++ {
						for x := 0; x < width; x++ {
							idx := y*width + x
							data[idx] = int64(rgbaImage.RGBAAt(x, y).B)
						}
					}

					return &tender.Matrix[int64]{
						Rows: height,
						Cols: width,
						Data: data,
					}, nil
				},
			},
			"get_alpha": &tender.NativeFunction{
				Name: "get_alpha",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}

					bounds := rgbaImage.Bounds()
					width := bounds.Dx()
					height := bounds.Dy()

					data := make([]int64, width*height)
					for y := 0; y < height; y++ {
						for x := 0; x < width; x++ {
							idx := y*width + x
							data[idx] = int64(rgbaImage.RGBAAt(x, y).A)
						}
					}

					return &tender.Matrix[int64]{
						Rows: height,
						Cols: width,
						Data: data,
					}, nil
				},
			},

			// Channel setters
			"set_red": &tender.NativeFunction{
				Name: "set_red",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}

					mat, ok := args[0].(*tender.Matrix[int64])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "matrix",
							Expected: "matrix:int",
							Found:    args[0].TypeName(),
						}
					}

					bounds := rgbaImage.Bounds()
					width := bounds.Dx()
					height := bounds.Dy()

					// Check dimensions
					if mat.Rows != height || mat.Cols != width {
						return &tender.Error{
							Value: &tender.String{
								Value: fmt.Sprintf("matrix dimensions mismatch: expected %dx%d, got %dx%d",
									height, width, mat.Rows, mat.Cols),
							},
						}, nil
					}

					// Set red channel
					for y := 0; y < height; y++ {
						for x := 0; x < width; x++ {
							idx := y*width + x
							val := mat.Data[idx]
							if val < 0 || val > 255 {
								return &tender.Error{
									Value: &tender.String{
										Value: fmt.Sprintf("value out of range (0-255): %d at (%d,%d)", val, x, y),
									},
								}, nil
							}
							c := rgbaImage.RGBAAt(x, y)
							c.R = uint8(val)
							rgbaImage.SetRGBA(x, y, c)
						}
					}

					return tender.NullValue, nil
				},
			},
			"set_green": &tender.NativeFunction{
				Name: "set_green",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}

					mat, ok := args[0].(*tender.Matrix[int64])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "matrix",
							Expected: "matrix:int",
							Found:    args[0].TypeName(),
						}
					}

					bounds := rgbaImage.Bounds()
					width := bounds.Dx()
					height := bounds.Dy()

					if mat.Rows != height || mat.Cols != width {
						return &tender.Error{
							Value: &tender.String{
								Value: fmt.Sprintf("matrix dimensions mismatch: expected %dx%d, got %dx%d",
									height, width, mat.Rows, mat.Cols),
							},
						}, nil
					}

					for y := 0; y < height; y++ {
						for x := 0; x < width; x++ {
							idx := y*width + x
							val := mat.Data[idx]
							if val < 0 || val > 255 {
								return &tender.Error{
									Value: &tender.String{
										Value: fmt.Sprintf("value out of range (0-255): %d at (%d,%d)", val, x, y),
									},
								}, nil
							}
							c := rgbaImage.RGBAAt(x, y)
							c.G = uint8(val)
							rgbaImage.SetRGBA(x, y, c)
						}
					}

					return tender.NullValue, nil
				},
			},
			"set_blue": &tender.NativeFunction{
				Name: "set_blue",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}

					mat, ok := args[0].(*tender.Matrix[int64])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "matrix",
							Expected: "matrix:int",
							Found:    args[0].TypeName(),
						}
					}

					bounds := rgbaImage.Bounds()
					width := bounds.Dx()
					height := bounds.Dy()

					if mat.Rows != height || mat.Cols != width {
						return &tender.Error{
							Value: &tender.String{
								Value: fmt.Sprintf("matrix dimensions mismatch: expected %dx%d, got %dx%d",
									height, width, mat.Rows, mat.Cols),
							},
						}, nil
					}

					for y := 0; y < height; y++ {
						for x := 0; x < width; x++ {
							idx := y*width + x
							val := mat.Data[idx]
							if val < 0 || val > 255 {
								return &tender.Error{
									Value: &tender.String{
										Value: fmt.Sprintf("value out of range (0-255): %d at (%d,%d)", val, x, y),
									},
								}, nil
							}
							c := rgbaImage.RGBAAt(x, y)
							c.B = uint8(val)
							rgbaImage.SetRGBA(x, y, c)
						}
					}

					return tender.NullValue, nil
				},
			},
			"set_alpha": &tender.NativeFunction{
				Name: "set_alpha",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}

					mat, ok := args[0].(*tender.Matrix[int64])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "matrix",
							Expected: "matrix:int",
							Found:    args[0].TypeName(),
						}
					}

					bounds := rgbaImage.Bounds()
					width := bounds.Dx()
					height := bounds.Dy()

					if mat.Rows != height || mat.Cols != width {
						return &tender.Error{
							Value: &tender.String{
								Value: fmt.Sprintf("matrix dimensions mismatch: expected %dx%d, got %dx%d",
									height, width, mat.Rows, mat.Cols),
							},
						}, nil
					}

					for y := 0; y < height; y++ {
						for x := 0; x < width; x++ {
							idx := y*width + x
							val := mat.Data[idx]
							if val < 0 || val > 255 {
								return &tender.Error{
									Value: &tender.String{
										Value: fmt.Sprintf("value out of range (0-255): %d at (%d,%d)", val, x, y),
									},
								}, nil
							}
							c := rgbaImage.RGBAAt(x, y)
							c.A = uint8(val)
							rgbaImage.SetRGBA(x, y, c)
						}
					}

					return tender.NullValue, nil
				},
			},
		},
	}
}

func makeRectangle(rect image.Rectangle) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"min": &tender.ImmutableMap{
				Value: map[string]tender.Object{
					"x": &tender.Int{Value: int64(rect.Min.X)},
					"y": &tender.Int{Value: int64(rect.Min.Y)},
				},
			},
			"max": &tender.ImmutableMap{
				Value: map[string]tender.Object{
					"x": &tender.Int{Value: int64(rect.Max.X)},
					"y": &tender.Int{Value: int64(rect.Max.Y)},
				},
			},
			"size": &tender.ImmutableMap{
				Value: map[string]tender.Object{
					"width":  &tender.Int{Value: int64(rect.Dx())},
					"height": &tender.Int{Value: int64(rect.Dy())},
				},
			},
		},
	}
}

func makeColor(col color.Color) *tender.Array {
	r, g, b, a := col.RGBA()
	return &tender.Array{Value: []tender.Object{
		&tender.Int{Value: int64(uint8(r >> 8))},
		&tender.Int{Value: int64(uint8(g >> 8))},
		&tender.Int{Value: int64(uint8(b >> 8))},
		&tender.Int{Value: int64(uint8(a >> 8))},
	}}
}