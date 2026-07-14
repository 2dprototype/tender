package stdlib

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"runtime"
	"sync"

	"github.com/2dprototype/tender"
)

// makeImageFilters wraps an image.Image and returns an ImmutableMap of filter functions.
func makeImageFilters(img image.Image) *tender.ImmutableMap {
	return &tender.ImmutableMap{
		Value: map[string]tender.Object{
			"blur": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					r, _ := tender.ToInt(args[0])
					newImg := applyBlurParallelOptimized(img, r)
					return makeImage(newImg), nil
				},
			},
			"bnw": &tender.NativeFunction{
				Name: "bnw",
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					t, _ := tender.ToUint8(args[0])
					return makeImage(applyBnWParallel(img, t)), nil
				},
			},
			// "glitch" filter: optionally accepts a maxShift (default 20)
			"glitch": &tender.NativeFunction{
				Value: func(args ...tender.Object) (tender.Object, error) {
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					m, _ := tender.ToInt(args[0])
					return makeImage(applyGlitchParallel(img, m)), nil
				},
			},
			"invert": &tender.NativeFunction{
				Name: "invert",
				Value: func(args ...tender.Object) (tender.Object, error) {
					newImg := applyInvertParallel(img)
					return makeImage(newImg), nil
				},
			},
			"grayscale": &tender.NativeFunction{
				Name: "grayscale",
				Value: func(args ...tender.Object) (tender.Object, error) {
					newImg := applyGrayscaleParallel(img)
					return makeImage(newImg), nil
				},
			},
			"sepia": &tender.NativeFunction{
				Name: "sepia",
				Value: func(args ...tender.Object) (tender.Object, error) {
					newImg := applySepiaParallel(img)
					return makeImage(newImg), nil
				},
			},
			"brightness": &tender.NativeFunction{
				Name: "brightness",
				Value: func(args ...tender.Object) (tender.Object, error) {
					// Required: offset (can be negative)
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					offset, ok := tender.ToInt(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "offset",
							Expected: "int",
							Found:    args[0].TypeName(),
						}
					}
					newImg := applyBrightnessParallel(img, offset)
					return makeImage(newImg), nil
				},
			},
			"contrast": &tender.NativeFunction{
				Name: "contrast",
				Value: func(args ...tender.Object) (tender.Object, error) {
					// Required: factor (float, e.g., 1.0 = no change)
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					factor, ok := tender.ToFloat64(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "factor",
							Expected: "float",
							Found:    args[0].TypeName(),
						}
					}
					newImg := applyContrastParallel(img, factor)
					return makeImage(newImg), nil
				},
			},
			"saturation": &tender.NativeFunction{
				Name: "saturation",
				Value: func(args ...tender.Object) (tender.Object, error) {
					// Required: factor (float; 0 = grayscale, 1 = original)
					if len(args) != 1 {
						return nil, tender.ErrWrongNumArguments
					}
					factor, ok := tender.ToFloat64(args[0])
					if !ok {
						return nil, tender.ErrInvalidArgumentType{
							Name:     "factor",
							Expected: "float",
							Found:    args[0].TypeName(),
						}
					}
					newImg := applySaturationParallel(img, factor)
					return makeImage(newImg), nil
				},
			},
			"sharpen": &tender.NativeFunction{
				Name: "sharpen",
				Value: func(args ...tender.Object) (tender.Object, error) {
					newImg := applyConvolutionParallel(img, [][]float64{
						{0, -1, 0},
						{-1, 5, -1},
						{0, -1, 0},
					}, 1, 0)
					return makeImage(newImg), nil
				},
			},
			"emboss": &tender.NativeFunction{
				Name: "emboss",
				Value: func(args ...tender.Object) (tender.Object, error) {
					// Using an emboss kernel with an offset to recenter colors.
					newImg := applyConvolutionParallel(img, [][]float64{
						{-2, -1, 0},
						{-1, 1, 1},
						{0, 1, 2},
					}, 1, 128)
					return makeImage(newImg), nil
				},
			},
			"edge": &tender.NativeFunction{
				Name: "edge",
				Value: func(args ...tender.Object) (tender.Object, error) {
					// Edge detection kernel; offset added for visibility.
					newImg := applyConvolutionParallel(img, [][]float64{
						{1, 1, 1},
						{1, -8, 1},
						{1, 1, 1},
					}, 1, 128)
					return makeImage(newImg), nil
				},
			},
            // New filters
            "hue": &tender.NativeFunction{
                Name: "hue",
                Value: func(args ...tender.Object) (tender.Object, error) {
                    if len(args) != 1 {
                        return nil, tender.ErrWrongNumArguments
                    }
                    hue, ok := tender.ToFloat64(args[0])
                    if !ok {
                        return nil, tender.ErrInvalidArgumentType{
                            Name:     "hue",
                            Expected: "float",
                            Found:    args[0].TypeName(),
                        }
                    }
                    return makeImage(applyHueParallel(img, hue)), nil
                },
            },
            
            "temperature": &tender.NativeFunction{
                Name: "temperature",
                Value: func(args ...tender.Object) (tender.Object, error) {
                    if len(args) != 1 {
                        return nil, tender.ErrWrongNumArguments
                    }
                    temp, ok := tender.ToFloat64(args[0])
                    if !ok {
                        return nil, tender.ErrInvalidArgumentType{
                            Name:     "temperature",
                            Expected: "float",
                            Found:    args[0].TypeName(),
                        }
                    }
                    return makeImage(applyTemperatureParallel(img, temp)), nil
                },
            },
            
            "vignette": &tender.NativeFunction{
                Name: "vignette",
                Value: func(args ...tender.Object) (tender.Object, error) {
                    if len(args) != 1 {
                        return nil, tender.ErrWrongNumArguments
                    }
                    intensity, ok := tender.ToFloat64(args[0])
                    if !ok {
                        return nil, tender.ErrInvalidArgumentType{
                            Name:     "intensity",
                            Expected: "float",
                            Found:    args[0].TypeName(),
                        }
                    }
                    return makeImage(applyVignetteParallel(img, intensity)), nil
                },
            },
            
            "pixelate": &tender.NativeFunction{
                Name: "pixelate",
                Value: func(args ...tender.Object) (tender.Object, error) {
                    if len(args) != 1 {
                        return nil, tender.ErrWrongNumArguments
                    }
                    size, ok := tender.ToInt(args[0])
                    if !ok {
                        return nil, tender.ErrInvalidArgumentType{
                            Name:     "size",
                            Expected: "int",
                            Found:    args[0].TypeName(),
                        }
                    }
                    return makeImage(applyPixelateParallel(img, size)), nil
                },
            },
            
            "sobel": &tender.NativeFunction{
                Name: "sobel",
                Value: func(args ...tender.Object) (tender.Object, error) {
                    return makeImage(applySobelEdgeDetection(img)), nil
                },
            },
        },
    }
}
// -------------------
// Filter Implementations
// -------------------

// applyBlurParallelOptimized applies a blur filter using a sliding window approach in parallel.
func applyBlurParallelOptimized(img image.Image, radius int) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	blurImg := image.NewRGBA(bounds)
	pixels := blurImg.Pix

	numWorkers := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	rowHeight := height / numWorkers

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		startY := i * rowHeight
		endY := (i + 1) * rowHeight
		if i == numWorkers-1 {
			endY = height
		}
		go func(startY, endY int) {
			defer wg.Done()
			for y := startY; y < endY; y++ {
				for x := 0; x < width; x++ {
					var rSum, gSum, bSum, count float64
					for ky := -radius; ky <= radius; ky++ {
						for kx := -radius; kx <= radius; kx++ {
							nx, ny := x+kx, y+ky
							if nx >= 0 && nx < width && ny >= 0 && ny < height {
								r, g, b, _ := img.At(nx+bounds.Min.X, ny+bounds.Min.Y).RGBA()
								rSum += float64(r >> 8)
								gSum += float64(g >> 8)
								bSum += float64(b >> 8)
								count++
							}
						}
					}
					offset := (y*width + x) * 4
					pixels[offset] = uint8(clamp(rSum/count, 0, 255))
					pixels[offset+1] = uint8(clamp(gSum/count, 0, 255))
					pixels[offset+2] = uint8(clamp(bSum/count, 0, 255))
					pixels[offset+3] = 255
				}
			}
		}(startY, endY)
	}
	wg.Wait()
	return blurImg
}

// applyBnWParallel converts the image to black and white using a luminance threshold.
func applyBnWParallel(img image.Image, threshold uint8) image.Image {
	bounds := img.Bounds()
	bnwImg := image.NewRGBA(bounds)
	numWorkers := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	rowHeight := bounds.Dy() / numWorkers

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		startY := bounds.Min.Y + i*rowHeight
		endY := bounds.Min.Y + (i+1)*rowHeight
		if i == numWorkers-1 {
			endY = bounds.Max.Y
		}
		go func(startY, endY int) {
			defer wg.Done()
			for y := startY; y < endY; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					r, g, b, a := img.At(x, y).RGBA()
					lum := 0.3*float64(r>>8) + 0.59*float64(g>>8) + 0.11*float64(b>>8)
					var val uint8
					if lum > float64(threshold) {
						val = 255
					}
					bnwImg.Set(x, y, color.RGBA{val, val, val, uint8(a >> 8)})
				}
			}
		}(startY, endY)
	}
	wg.Wait()
	return bnwImg
}

// applyGlitchParallel applies a glitch effect by shifting RGB channels randomly.
func applyGlitchParallel(img image.Image, maxShift int) image.Image {
	bounds := img.Bounds()
	glitchImg := image.NewRGBA(bounds)
	numWorkers := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	rowHeight := bounds.Dy() / numWorkers

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		startY := bounds.Min.Y + i*rowHeight
		endY := bounds.Min.Y + (i+1)*rowHeight
		if i == numWorkers-1 {
			endY = bounds.Max.Y
		}
		go func(startY, endY int) {
			defer wg.Done()
			for y := startY; y < endY; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					rx := x + rand.Intn(maxShift) - maxShift/2
					gx := x
					bx := x - rand.Intn(maxShift) + maxShift/2
					if rx < bounds.Min.X || rx >= bounds.Max.X {
						rx = x
					}
					if gx < bounds.Min.X || gx >= bounds.Max.X {
						gx = x
					}
					if bx < bounds.Min.X || bx >= bounds.Max.X {
						bx = x
					}
					r, _, _, _ := img.At(rx, y).RGBA()
					_, g, _, _ := img.At(gx, y).RGBA()
					_, _, b, a := img.At(bx, y).RGBA()
					glitchImg.Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
				}
			}
		}(startY, endY)
	}
	wg.Wait()
	return glitchImg
}

// applyInvertParallel inverts the image colors.
func applyInvertParallel(img image.Image) image.Image {
	bounds := img.Bounds()
	out := image.NewRGBA(bounds)
	numWorkers := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	rowHeight := bounds.Dy() / numWorkers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		startY := bounds.Min.Y + i*rowHeight
		endY := bounds.Min.Y + (i+1)*rowHeight
		if i == numWorkers-1 {
			endY = bounds.Max.Y
		}
		go func(startY, endY int) {
			defer wg.Done()
			for y := startY; y < endY; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					orig := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
					inv := color.RGBA{
						R: 255 - orig.R,
						G: 255 - orig.G,
						B: 255 - orig.B,
						A: orig.A,
					}
					out.Set(x, y, inv)
				}
			}
		}(startY, endY)
	}
	wg.Wait()
	return out
}

// applyGrayscaleParallel converts the image to grayscale.
func applyGrayscaleParallel(img image.Image) image.Image {
	bounds := img.Bounds()
	out := image.NewRGBA(bounds)
	numWorkers := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	rowHeight := bounds.Dy() / numWorkers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		startY := bounds.Min.Y + i*rowHeight
		endY := bounds.Min.Y + (i+1)*rowHeight
		if i == numWorkers-1 {
			endY = bounds.Max.Y
		}
		go func(startY, endY int) {
			defer wg.Done()
			for y := startY; y < endY; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					orig := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
					// Using standard luminance formula.
					lum := uint8(0.299*float64(orig.R) + 0.587*float64(orig.G) + 0.114*float64(orig.B))
					out.Set(x, y, color.RGBA{R: lum, G: lum, B: lum, A: orig.A})
				}
			}
		}(startY, endY)
	}
	wg.Wait()
	return out
}

// applySepiaParallel applies a sepia tone to the image.
func applySepiaParallel(img image.Image) image.Image {
	bounds := img.Bounds()
	out := image.NewRGBA(bounds)
	numWorkers := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	rowHeight := bounds.Dy() / numWorkers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		startY := bounds.Min.Y + i*rowHeight
		endY := bounds.Min.Y + (i+1)*rowHeight
		if i == numWorkers-1 {
			endY = bounds.Max.Y
		}
		go func(startY, endY int) {
			defer wg.Done()
			for y := startY; y < endY; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					orig := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
					r := clamp(0.393*float64(orig.R)+0.769*float64(orig.G)+0.189*float64(orig.B), 0, 255)
					g := clamp(0.349*float64(orig.R)+0.686*float64(orig.G)+0.168*float64(orig.B), 0, 255)
					b := clamp(0.272*float64(orig.R)+0.534*float64(orig.G)+0.131*float64(orig.B), 0, 255)
					out.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), orig.A})
				}
			}
		}(startY, endY)
	}
	wg.Wait()
	return out
}

// applyBrightnessParallel adjusts brightness by adding an offset to each color channel.
func applyBrightnessParallel(img image.Image, offset int) image.Image {
	bounds := img.Bounds()
	out := image.NewRGBA(bounds)
	numWorkers := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	rowHeight := bounds.Dy() / numWorkers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		startY := bounds.Min.Y + i*rowHeight
		endY := bounds.Min.Y + (i+1)*rowHeight
		if i == numWorkers-1 {
			endY = bounds.Max.Y
		}
		go func(startY, endY int) {
			defer wg.Done()
			for y := startY; y < endY; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					orig := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
					r := clamp(float64(orig.R+uint8(offset)), 0, 255)
					g := clamp(float64(orig.G+uint8(offset)), 0, 255)
					b := clamp(float64(orig.B+uint8(offset)), 0, 255)
					out.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), orig.A})
				}
			}
		}(startY, endY)
	}
	wg.Wait()
	return out
}

// applyContrastParallel adjusts contrast by a factor.
func applyContrastParallel(img image.Image, factor float64) image.Image {
	bounds := img.Bounds()
	out := image.NewRGBA(bounds)
	numWorkers := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	rowHeight := bounds.Dy() / numWorkers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		startY := bounds.Min.Y + i*rowHeight
		endY := bounds.Min.Y + (i+1)*rowHeight
		if i == numWorkers-1 {
			endY = bounds.Max.Y
		}
		go func(startY, endY int) {
			defer wg.Done()
			for y := startY; y < endY; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					orig := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
					// Adjust contrast: new = ((old - 128) * factor) + 128.
					r := clamp(((float64(orig.R)-128)*factor)+128, 0, 255)
					g := clamp(((float64(orig.G)-128)*factor)+128, 0, 255)
					b := clamp(((float64(orig.B)-128)*factor)+128, 0, 255)
					out.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), orig.A})
				}
			}
		}(startY, endY)
	}
	wg.Wait()
	return out
}

// applySaturationParallel adjusts saturation by interpolating between grayscale and original color.
func applySaturationParallel(img image.Image, factor float64) image.Image {
	bounds := img.Bounds()
	out := image.NewRGBA(bounds)
	numWorkers := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	rowHeight := bounds.Dy() / numWorkers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		startY := bounds.Min.Y + i*rowHeight
		endY := bounds.Min.Y + (i+1)*rowHeight
		if i == numWorkers-1 {
			endY = bounds.Max.Y
		}
		go func(startY, endY int) {
			defer wg.Done()
			for y := startY; y < endY; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					orig := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
					// Calculate luminance.
					lum := 0.299*float64(orig.R) + 0.587*float64(orig.G) + 0.114*float64(orig.B)
					// Interpolate between gray and original.
					r := clamp(lum+factor*(float64(orig.R)-lum), 0, 255)
					g := clamp(lum+factor*(float64(orig.G)-lum), 0, 255)
					b := clamp(lum+factor*(float64(orig.B)-lum), 0, 255)
					out.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), orig.A})
				}
			}
		}(startY, endY)
	}
	wg.Wait()
	return out
}

// applyConvolutionParallel applies a convolution filter using the provided kernel, divisor, and offset.
func applyConvolutionParallel(img image.Image, kernel [][]float64, divisor, offset float64) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	out := image.NewRGBA(bounds)
	numWorkers := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	rowHeight := height / numWorkers
	kSize := len(kernel)
	kRadius := kSize / 2

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		startY := i * rowHeight
		endY := (i + 1) * rowHeight
		if i == numWorkers-1 {
			endY = height
		}
		go func(startY, endY int) {
			defer wg.Done()
			for y := startY; y < endY; y++ {
				for x := 0; x < width; x++ {
					var rAcc, gAcc, bAcc float64
					for ky := 0; ky < kSize; ky++ {
						for kx := 0; kx < kSize; kx++ {
							ix := x + kx - kRadius
							iy := y + ky - kRadius
							// Clamp to bounds.
							if ix < 0 {
								ix = 0
							}
							if iy < 0 {
								iy = 0
							}
							if ix >= width {
								ix = width - 1
							}
							if iy >= height {
								iy = height - 1
							}
							r, g, b, _ := img.At(ix+bounds.Min.X, iy+bounds.Min.Y).RGBA()
							weight := kernel[ky][kx]
							rAcc += float64(r>>8) * weight
							gAcc += float64(g>>8) * weight
							bAcc += float64(b>>8) * weight
						}
					}
					rVal := clamp((rAcc/divisor)+offset, 0, 255)
					gVal := clamp((gAcc/divisor)+offset, 0, 255)
					bVal := clamp((bAcc/divisor)+offset, 0, 255)
					out.Set(x+bounds.Min.X, y+bounds.Min.Y, color.RGBA{uint8(rVal), uint8(gVal), uint8(bVal), 255})
				}
			}
		}(startY, endY)
	}
	wg.Wait()
	return out
}

// clamp ensures value is within min and max.
func clamp(value, min, max float64) float64 {
	return math.Max(min, math.Min(max, value))
}

// New filter implementations

// applyHueParallel adjusts the hue of the image
func applyHueParallel(img image.Image, hue float64) image.Image {
    bounds := img.Bounds()
    out := image.NewRGBA(bounds)
    numWorkers := runtime.GOMAXPROCS(0)
    var wg sync.WaitGroup
    rowHeight := bounds.Dy() / numWorkers

    hueRad := hue * math.Pi / 180

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        startY := bounds.Min.Y + i*rowHeight
        endY := bounds.Min.Y + (i+1)*rowHeight
        if i == numWorkers-1 {
            endY = bounds.Max.Y
        }
        go func(startY, endY int) {
            defer wg.Done()
            for y := startY; y < endY; y++ {
                for x := bounds.Min.X; x < bounds.Max.X; x++ {
                    r, g, b, a := img.At(x, y).RGBA()
                    rf, gf, bf := float64(r>>8), float64(g>>8), float64(b>>8)
                    
                    // Convert RGB to HSL
                    h, s, l := rgbToHsl(rf, gf, bf)
                    
                    // Adjust hue
                    h = math.Mod(h+hueRad, 2*math.Pi)
                    if h < 0 {
                        h += 2 * math.Pi
                    }
                    
                    // Convert back to RGB
                    rf, gf, bf = hslToRgb(h, s, l)
                    
                    out.Set(x, y, color.RGBA{
                        uint8(clamp(rf, 0, 255)),
                        uint8(clamp(gf, 0, 255)),
                        uint8(clamp(bf, 0, 255)),
                        uint8(a >> 8),
                    })
                }
            }
        }(startY, endY)
    }
    wg.Wait()
    return out
}

// applyTemperatureParallel adjusts color temperature
func applyTemperatureParallel(img image.Image, temp float64) image.Image {
    bounds := img.Bounds()
    out := image.NewRGBA(bounds)
    numWorkers := runtime.GOMAXPROCS(0)
    var wg sync.WaitGroup
    rowHeight := bounds.Dy() / numWorkers

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        startY := bounds.Min.Y + i*rowHeight
        endY := bounds.Min.Y + (i+1)*rowHeight
        if i == numWorkers-1 {
            endY = bounds.Max.Y
        }
        go func(startY, endY int) {
            defer wg.Done()
            for y := startY; y < endY; y++ {
                for x := bounds.Min.X; x < bounds.Max.X; x++ {
                    r, g, b, a := img.At(x, y).RGBA()
                    rf, gf, bf := float64(r>>8), float64(g>>8), float64(b>>8)
                    
                    // Adjust temperature (warm/cool)
                    if temp > 0 {
                        // Warm - increase red, decrease blue
                        rf += temp * 0.5
                        bf -= temp * 0.3
                    } else {
                        // Cool - decrease red, increase blue
                        rf += temp * 0.5
                        bf -= temp * 0.3
                    }
                    
                    out.Set(x, y, color.RGBA{
                        uint8(clamp(rf, 0, 255)),
                        uint8(clamp(gf, 0, 255)),
                        uint8(clamp(bf, 0, 255)),
                        uint8(a >> 8),
                    })
                }
            }
        }(startY, endY)
    }
    wg.Wait()
    return out
}

// applyVignetteParallel applies vignette effect
func applyVignetteParallel(img image.Image, intensity float64) image.Image {
    bounds := img.Bounds()
    out := image.NewRGBA(bounds)
    width, height := bounds.Dx(), bounds.Dy()
    centerX, centerY := float64(width)/2, float64(height)/2
    maxDist := math.Sqrt(centerX*centerX + centerY*centerY)

    numWorkers := runtime.GOMAXPROCS(0)
    var wg sync.WaitGroup
    rowHeight := bounds.Dy() / numWorkers

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        startY := bounds.Min.Y + i*rowHeight
        endY := bounds.Min.Y + (i+1)*rowHeight
        if i == numWorkers-1 {
            endY = bounds.Max.Y
        }
        go func(startY, endY int) {
            defer wg.Done()
            for y := startY; y < endY; y++ {
                for x := bounds.Min.X; x < bounds.Max.X; x++ {
                    r, g, b, a := img.At(x, y).RGBA()
                    
                    // Calculate distance from center
                    dist := math.Sqrt(math.Pow(float64(x-bounds.Min.X)-centerX, 2) + 
                                    math.Pow(float64(y-bounds.Min.Y)-centerY, 2))
                    
                    // Calculate vignette factor
                    factor := 1.0 - (dist/maxDist)*intensity
                    factor = clamp(factor, 0, 1)
                    
                    out.Set(x, y, color.RGBA{
                        uint8(clamp(float64(r>>8)*factor, 0, 255)),
                        uint8(clamp(float64(g>>8)*factor, 0, 255)),
                        uint8(clamp(float64(b>>8)*factor, 0, 255)),
                        uint8(a >> 8),
                    })
                }
            }
        }(startY, endY)
    }
    wg.Wait()
    return out
}

// applyPixelateParallel creates pixelated effect
func applyPixelateParallel(img image.Image, size int) image.Image {
    bounds := img.Bounds()
    out := image.NewRGBA(bounds)
    width, height := bounds.Dx(), bounds.Dy()

    numWorkers := runtime.GOMAXPROCS(0)
    var wg sync.WaitGroup
    blockHeight := height / numWorkers

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        startY := i * blockHeight
        endY := (i + 1) * blockHeight
        if i == numWorkers-1 {
            endY = height
        }
        go func(startY, endY int) {
            defer wg.Done()
            for y := startY; y < endY; y += size {
                for x := 0; x < width; x += size {
                    // Calculate block bounds
                    blockWidth := size
                    blockHeight := size
                    if x+blockWidth > width {
                        blockWidth = width - x
                    }
                    if y+blockHeight > endY {
                        blockHeight = endY - y
                    }
                    
                    // Calculate average color in block
                    var rSum, gSum, bSum, count float64
                    for by := y; by < y+blockHeight; by++ {
                        for bx := x; bx < x+blockWidth; bx++ {
                            r, g, b, _ := img.At(bx+bounds.Min.X, by+bounds.Min.Y).RGBA()
                            rSum += float64(r >> 8)
                            gSum += float64(g >> 8)
                            bSum += float64(b >> 8)
                            count++
                        }
                    }
                    
                    avgR := rSum / count
                    avgG := gSum / count
                    avgB := bSum / count
                    
                    // Fill block with average color
                    for by := y; by < y+blockHeight; by++ {
                        for bx := x; bx < x+blockWidth; bx++ {
                            out.Set(bx+bounds.Min.X, by+bounds.Min.Y, color.RGBA{
                                uint8(avgR),
                                uint8(avgG),
                                uint8(avgB),
                                255,
                            })
                        }
                    }
                }
            }
        }(startY, endY)
    }
    wg.Wait()
    return out
}

// applySobelEdgeDetection applies Sobel edge detection
func applySobelEdgeDetection(img image.Image) image.Image {
    bounds := img.Bounds()
    out := image.NewRGBA(bounds)
    width, height := bounds.Dx(), bounds.Dy()

    // Sobel kernels
    sobelX := [3][3]float64{
        {-1, 0, 1},
        {-2, 0, 2},
        {-1, 0, 1},
    }
    sobelY := [3][3]float64{
        {-1, -2, -1},
        {0, 0, 0},
        {1, 2, 1},
    }

    numWorkers := runtime.GOMAXPROCS(0)
    var wg sync.WaitGroup
    rowHeight := height / numWorkers

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        startY := i * rowHeight
        endY := (i + 1) * rowHeight
        if i == numWorkers-1 {
            endY = height
        }
        go func(startY, endY int) {
            defer wg.Done()
            for y := startY; y < endY; y++ {
                for x := 0; x < width; x++ {
                    var gx, gy float64
                    
                    // Apply Sobel kernels
                    for ky := -1; ky <= 1; ky++ {
                        for kx := -1; kx <= 1; kx++ {
                            nx, ny := x+kx, y+ky
                            if nx >= 0 && nx < width && ny >= 0 && ny < height {
                                r, _, _, _ := img.At(nx+bounds.Min.X, ny+bounds.Min.Y).RGBA()
                                gray := float64(r >> 8)
                                gx += gray * sobelX[ky+1][kx+1]
                                gy += gray * sobelY[ky+1][kx+1]
                            }
                        }
                    }
                    
                    // Calculate gradient magnitude
                    magnitude := math.Sqrt(gx*gx + gy*gy)
                    magnitude = clamp(magnitude, 0, 255)
                    
                    out.Set(x+bounds.Min.X, y+bounds.Min.Y, color.RGBA{
                        uint8(255 - magnitude),
                        uint8(255 - magnitude),
                        uint8(255 - magnitude),
                        255,
                    })
                }
            }
        }(startY, endY)
    }
    wg.Wait()
    return out
}

// RGB to HSL conversion
func rgbToHsl(r, g, b float64) (h, s, l float64) {
    r /= 255
    g /= 255
    b /= 255
    
    max := math.Max(r, math.Max(g, b))
    min := math.Min(r, math.Min(g, b))
    
    l = (max + min) / 2
    
    if max == min {
        h = 0
        s = 0
    } else {
        d := max - min
        if l > 0.5 {
            s = d / (2 - max - min)
        } else {
            s = d / (max + min)
        }
        
        switch max {
        case r:
            h = (g - b) / d
            if g < b {
                h += 6
            }
        case g:
            h = (b-r)/d + 2
        case b:
            h = (r-g)/d + 4
        }
        h *= math.Pi / 3
    }
    return
}

// HSL to RGB conversion
func hslToRgb(h, s, l float64) (r, g, b float64) {
    if s == 0 {
        r, g, b = l, l, l
    } else {
        var q float64
        if l < 0.5 {
            q = l * (1 + s)
        } else {
            q = l + s - l*s
        }
        p := 2*l - q
        
        r = hueToRgb(p, q, h+2*math.Pi/3)
        g = hueToRgb(p, q, h)
        b = hueToRgb(p, q, h-2*math.Pi/3)
    }
    
    r *= 255
    g *= 255
    b *= 255
    return
}

func hueToRgb(p, q, t float64) float64 {
    if t < 0 {
        t += 2 * math.Pi
    }
    if t > 2*math.Pi {
        t -= 2 * math.Pi
    }
    
    t /= 2 * math.Pi
    
    if t < 1.0/6.0 {
        return p + (q-p)*6*t
    }
    if t < 1.0/2.0 {
        return q
    }
    if t < 2.0/3.0 {
        return p + (q-p)*(2.0/3.0-t)*6
    }
    return p
}