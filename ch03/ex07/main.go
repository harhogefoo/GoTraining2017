package main

import (
	"image"
	"math/cmplx"
	"image/color"
	"image/png"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < width; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i:= uint8(0); i < iterations; i++ {
		z -= (z - 1 / (z * z * z)) / 4
		switch {
		case cmplx.Abs((1 + 0i) - z) < 1e-6:
			return color.RGBA{255 - contrast * i, 0, 0, 0xff}
		case cmplx.Abs((-1 + 0i) - z) < 1e-6:
			return color.RGBA{0, 255 - contrast * i, 0, 0xff}
		case cmplx.Abs((0 + 1i) - z) < 1e-6:
			return color.RGBA{0, 0, 255 - contrast * i, 0xff}
		case cmplx.Abs((0 - 1i) - z) < 1e-6:
			return color.RGBA{255 - contrast * i,255 - contrast * i, 0, 0xff}
		}
	}
	return color.Black
}