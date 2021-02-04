//Print a gif lissajous file
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	blackIndex = 0
	greenIndex = 1
	oxRR       = 100
	oxGG       = 250
	oxBB       = 100
	oxff       = 100
)

var palettes = [][]color.Color{
	{color.Black, color.RGBA{100, 250, 100, 100}},
	{color.White, color.RGBA{100, 250, 100, 100}},
	{color.RGBA{10, 15, 20, 10}, color.RGBA{180, 150, 200, 150}},
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	for _, v := range os.Args[1:] {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Informe apenas números")
			os.Exit(1)
		}
		lissajous(os.Stdout, palettes[i])
	}

}

func lissajous(out io.Writer, palette []color.Color) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
