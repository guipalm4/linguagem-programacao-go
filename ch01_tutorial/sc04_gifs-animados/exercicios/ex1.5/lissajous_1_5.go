// ex1.5 generates GIF animations of random Lissajous figures, with a
// green-on-black palette.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Packages not needed by version in book.

var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 255}}

const (
	blackIndex = 0 // primeira cor da paleta
	greenIndex = 1 // próxima cor da paleta
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // número de revoluções completas do oscilador x
		res     = 0.001 // resolução angular
		size    = 100   // canvas da imagem cobre de [-size...+size]
		nFrames = 64    // número de quadros da animação
		delay   = 8     // tempo entre quadros em unidades de 10ms
	)

	freq := rand.Float64() * 3.0 //frequencia ralativa do oscilador y
	anim := gif.GIF{LoopCount: nFrames}
	phase := 0.0 // diferença de fase
	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTA: ignorando erros de codificação
}
