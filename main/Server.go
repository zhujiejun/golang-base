package main

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
)
import "sync"

var lock sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("192.168.100.100:18080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	//lock.Lock()
	//count++
	//lock.Unlock()
	//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	lissajous(w)
}

func counter(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	lock.Unlock()
}

func lissajous(out io.Writer) {
	const (
		cycles   = 5
		delay    = 8
		nframe   = 64
		size     = 100
		resource = 0.01
	)
	phase := 0.0
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframe}
	for i := 0; i < nframe; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette.WebSafe)
		for t := 0.0; t < cycles*2*math.Pi; t += resource {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 15)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
