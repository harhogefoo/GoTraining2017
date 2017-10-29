package main

import (
	"net/http"
	"log"
	"fmt"
	"sync"
	"io"
	"image/gif"
	"image"
	"math"
	"math/rand"
	"image/color"
	"strconv"
)

var mu sync.Mutex
var count int

var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 255}}

const (
	whiteIndex = 0 // パレットの最初の色
	blackIndex = 1 // パレットの次の色
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	}
	http.HandleFunc("/", handler)
	// http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}


// counterは今までの呼び出し数を返します。
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	fCycles, isExist := r.Form["cycles"] // 発振器xが完了する周回の回数
	cycles := 5.0
	if isExist {
		cycles, _ = strconv.ParseFloat(fCycles[0], 64)
	}

	const (
		res     = 0.01 // 回転の分解能
		size    = 100  // 画像キャンバスは[-size..+size]の範囲で扱う
		nframes = 64   // アニメーションフレーム数
		delay   = 9    // 10ms単位でのフレーム間の遅延
	)
	freq := rand.Float64() * 3.0 // 発振器yの相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t+= res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意: エンコードエラーを無視
}


