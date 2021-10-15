package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var color_palette = []color.Color{color.White, color.Black}

const(
	white_Index = 0 // first color in the palette
	black_Index = 1 // following color in the palette
)

func main() {
	random_gifs(os.Stdout)
}

func random_gifs(out io.Writer) {
	const(
		cycle = 5 // number of complete x oscillator revolutions
		resol = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i :=0; i<nframes; i++{
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, color_palette)
		for t:= 0.0; t < cycle*2*math.Pi; t+=resol{
			x:= math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), black_Index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay,delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}