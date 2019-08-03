package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{G: 0xFF, A: 0xFF},
	color.RGBA{R: 0xFF, A: 0xFF},
	color.RGBA{B: 0xFF, A: 0xFF},
}

func main() {
	http.HandleFunc("/", lissajousHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	v := queryValue(r.URL.Query(), "cycles", "5")
	c, err := strconv.Atoi(v)
	if err != nil {
		_, _ = fmt.Fprintf(w, "issajous: invalid cycles: %s	%v", v, err)
		return
	}
	lissajous(w, float64(c))
}

func queryValue(values url.Values, key string, defaultValue string) string {
	v := values.Get(key)
	if v == "" {
		v = defaultValue
	}
	return v
}

func lissajous(out io.Writer, cycles float64) {
	const (
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
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i%len(palette)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	_ = gif.EncodeAll(out, &anim)
}
