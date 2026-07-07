package stdlib

import (
	// "fmt"
	"image"
	jpeg "image/jpeg"
	bmp "golang.org/x/image/bmp"
	tiff "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	"image/png"
	"image/color"
	"image/draw"
	"os"
	"bytes"
	"github.com/2dprototype/tender"
)

var imageModule = map[string]tender.Object{
	"new": &tender.UserFunction{Value: imageNew},
	"load" : &tender.UserFunction{Value: imageLoad},
	"decode" : &tender.UserFunction{Value: imageDecode},
	"formats" : &tender.ImmutableArray{Value: []tender.Object{
			&tender.String{Value: "png"},
			&tender.String{Value: "jpeg"},
			&tender.String{Value: "bmp"},
			&tender.String{Value: "tiff"},
			&tender.String{Value: "webp"},
		},
	},
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

func makeImage(img image.Image) *tender.ImmutableMap {
	// Convert the image to *image.RGBA if it's not already
	if _, ok := img.(*image.RGBA); !ok {
		bounds := img.Bounds()
		newImg := image.NewRGBA(bounds)
		draw.Draw(newImg, bounds, img, bounds.Min, draw.Src)
		img = newImg
	}	

	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"filters": makeImageFilters(img),
			"encode" : &tender.UserFunction{
				Name: "encode",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
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
					buffer := new(bytes.Buffer)

					if format == "png" {
						err := png.Encode(buffer, img)
						if err != nil {
							return wrapError(err), nil
						}
					} else if format == "jpeg" {
						err := jpeg.Encode(buffer, img, nil)
						if err != nil {
							return wrapError(err), nil
						}
					} else if format == "tiff" {
						err := tiff.Encode(buffer, img, nil)
						if err != nil {
							return wrapError(err), nil
						}
					} else if format == "bmp" {
						err := bmp.Encode(buffer, img)
						if err != nil {
							return wrapError(err), nil
						}
					}

					return &tender.Bytes{Value: buffer.Bytes()}, nil
				},
			},
			"bounds": &tender.UserFunction{
				Name: "bounds",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					rect := img.Bounds()
					return makeRectangle(rect), nil
				},
			},
			"at": &tender.UserFunction{
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

					color := img.At(x, y)
					return makeColor(color), nil
				},
			},
			"pixels": &tender.UserFunction{
				Name: "pixels",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}
					bounds := img.Bounds()
					return &tender.Int{Value: int64((bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y))}, nil
				},
			},	
			"get_pixels": &tender.UserFunction{
				Name: "get_pixels",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 0 {
						return nil, tender.ErrWrongNumArguments
					}

					rgbaImage, ok := img.(*image.RGBA)
					if !ok {
						return nil, nil
					}

					pixels := make([]tender.Object, len(rgbaImage.Pix))
					for i, p := range rgbaImage.Pix {
						pixels[i] = &tender.Int{Value: int64(p)}
					}

					return &tender.Array{Value: pixels}, nil
				},
			},	
			"set_pixels": &tender.UserFunction{
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

					rgbaImage, ok := img.(*image.RGBA)
					if !ok {
						return nil, nil
					}

					if len(pixelArray.Value) > len(rgbaImage.Pix) {
						return &tender.Error{Value: &tender.String{Value: "Failed to set pixels: Length of pixel array is greater than image dimensions"}}, nil 
					}

					for i, pixel := range pixelArray.Value {
						val, _ := tender.ToInt(pixel)
						rgbaImage.Pix[i] = uint8(val)
					}

					return nil, nil
				},
			},
			"set": &tender.UserFunction{
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

					arr, ok := args[2].(*tender.Array)
					if !ok || len(arr.Value) != 4 {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "color",
							Expected: "[4]array",
							Found:    args[2].TypeName(),
						}
					}

					red, ok1 := tender.ToUint8(arr.Value[0])
					green, ok2 := tender.ToUint8(arr.Value[1])
					blue, ok3 := tender.ToUint8(arr.Value[2])
					alpha, ok4 := tender.ToUint8(arr.Value[3])

					if !ok1 || !ok2 || !ok3 || !ok4 {
						return nil, nil
					}

					img.(*image.RGBA).Set(x, y, color.RGBA{red, green, blue, alpha})
					return nil, nil
				},
			},
			"save": &tender.UserFunction{
				Name: "save",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 2 {
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
					format, ok := tender.ToString(args[1])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "path",
							Expected: "string",
							Found:    args[1].TypeName(),
						}
					}

					file, err := os.Create(path)
					if err != nil {
						return wrapError(err), nil
					}

					defer file.Close()

					if format == "png" {
						err = png.Encode(file, img)
						if err != nil {
							return wrapError(err), nil
						}
					} else if format == "jpeg" {
						err = jpeg.Encode(file, img, nil)
						if err != nil {
							return wrapError(err), nil
						}
					} else if format == "tiff" {
						err = tiff.Encode(file, img, nil)
						if err != nil {
							return wrapError(err), nil
						}
					} else if format == "bmp" {
						err = bmp.Encode(file, img)
						if err != nil {
							return wrapError(err), nil
						}
					}

					return nil, nil
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
	// Check if the color is RGBA or NRGBA
	if rgbaColor, ok := col.(color.RGBA); ok {
		return &tender.Array{Value: []tender.Object{
				&tender.Int{Value: int64(rgbaColor.R)},
				&tender.Int{Value: int64(rgbaColor.G)},
				&tender.Int{Value: int64(rgbaColor.B)},
				&tender.Int{Value: int64(rgbaColor.A)},
			},
		}
	} else {
		return &tender.Array{Value: []tender.Object{
				&tender.Int{Value: 0},
				&tender.Int{Value: 0},
				&tender.Int{Value: 0},
				&tender.Int{Value: 0},
			},
		}
	}
}

// func makeColor(col color.Color) *tender.ImmutableMap {
	// // Check if the color is RGBA or NRGBA
	// if rgbaColor, ok := col.(color.RGBA); ok {
		// return &tender.ImmutableMap{
			// Value: map[string]tender.Object{
				// "r":   &tender.Int{Value: int64(rgbaColor.R)},
				// "g": &tender.Int{Value: int64(rgbaColor.G)},
				// "b":  &tender.Int{Value: int64(rgbaColor.B)},
				// "a": &tender.Int{Value: int64(rgbaColor.A)},
			// },
		// }
	// } else {
		// return &tender.ImmutableMap{
			// Value: map[string]tender.Object{
				// "r":   &tender.Int{Value: 0},
				// "g": &tender.Int{Value: 0},
				// "b":  &tender.Int{Value: 0},
				// "a": &tender.Int{Value: 0},
			// },
		// }
	// }
// }


// func makeColor(col color.Color) *tender.ImmutableMap {
	// // Implement color object creation based on your needs
	// // You can access color components and create a custom structure
	// return &tender.ImmutableMap{
		// Value: map[string]tender.Object{
			// "red":   &tender.Int{Value: 0},
			// "green": &tender.Int{Value: 0},
			// "blue":  &tender.Int{Value: 0},
			// "alpha": &tender.Int{Value: 0},
		// },
	// }
// }

// func toColorArray(obj tender.Object) ([]uint8, bool) {
	// arr, ok := obj.(*tender.Array)
	// if !ok || len(arr.Value) != 3 {
		// return nil, false
	// }

	// red, ok1 := tender.ToInt(arr.Value[0])
	// green, ok2 := tender.ToInt(arr.Value[1])
	// blue, ok3 := tender.ToInt(arr.Value[2])

	// if !ok1 || !ok2 || !ok3 {
		// return nil, false
	// }

	// return []uint8{uint8(red), uint8(green), uint8(blue)}, true
// }
