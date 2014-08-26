package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"code.google.com/p/draw2d/draw2d"
)

const (
	w, h   = 200, 200
	border = 20
)

var (
	black = color.RGBA{0x33, 0x33, 0x33, 0xFF}
	white = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	gray  = color.RGBA{0xDC, 0xDC, 0xDC, 0xFF}
)

func main() {
	for i, f := range map[string]func(float64) float64{
		"linear": func(n float64) float64 {
			return n
		},
		"sineIn": func(n float64) float64 {
			return -math.Cos(n*(math.Pi/2.0)) + 1
		},
		"sineOut": func(n float64) float64 {
			return math.Sin(n * (math.Pi / 2.0))
		},
		"sineInOut": func(n float64) float64 {
			return -0.5 * (math.Cos(math.Pi*n) - 1)
		},
		"backIn": func(n float64) float64 {
			s := 1.70158
			return n * n * ((s+1)*n - s)
		},
		"backOut": func(n float64) float64 {
			n = n - 1
			s := 1.70158
			return n*n*((s+1)*n+s) + 1
		},
		"elasticOut": func(n float64) float64 {
			return math.Pow(2, -10*n)*math.Sin((n-0.075)*(2*math.Pi)/0.3) + 1
		},
		"bounceOut": func(n float64) float64 {
			s, p := 7.5625, 2.75
			if n < (1.0 / p) {
				return s * n * n
			}
			if n < (2.0 / p) {
				n -= (1.5 / p)
				return s*n*n + 0.75
			}
			if n < (2.5 / p) {
				n -= (2.25 / p)
				return s*n*n + 0.9375
			}
			n -= (2.625 / p)
			return s*n*n + 0.984375
		},
	} {
		saveToPng("images/"+i+".png", createImage(f))
	}
}

func createImage(easing func(float64) float64) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	gc := draw2d.NewGraphicContext(img)
	gc.SetLineWidth(1)

	// border
	gc.SetStrokeColor(gray)
	gc.MoveTo(border, 0)
	gc.LineTo(border, h)
	gc.MoveTo(w-border, 0)
	gc.LineTo(w-border, h)
	gc.Stroke()

	// graph
	gc.SetStrokeColor(black)
	gc.MoveTo(border, border)

	var x, y float64
	for i := 0.0; i < 1; i += 1.0 / w {
		y = i * (h - border*2)
		x = easing(i) * (w - border*2)
		gc.LineTo(x+border, y+border)
	}
	gc.Stroke()

	return img
}

func saveToPng(file string, img image.Image) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}
