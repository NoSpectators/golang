// Mandelbrot emits a PNG image of the Mandelbrot fractal.

package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"log"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	//THE BELOW LINE IS COMMENTED OUT FROM BOOK
	//png.Encode(os.Stdout, img) // NOTE: ignoring errors
	//instead of printing to stdout, I create png file. Looks better. 
	//https://golang.org/pkg/image/png/#pkg-examples  (see func Encode example)

	f, err := os.Create("mandelbrot.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
