package main

import (
	"net/http"
	"log"
	"fmt"
	"math"
	"strconv"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		setup(r)
		polygon(w, r)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

var width int64 = 600
var height int64 = 320
var hexStr string = "FFFFFF"

var (
	cells         = 100
	xyrange       = 30.0
	xyscale       = float64(width) / 2 / xyrange
	zscale        = float64(height) * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func setup(r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	fHeight, isExist := r.Form["height"]
	if isExist {
		height, _ = strconv.ParseInt(fHeight[0], 10, 64)
	}
	fWidth, isExist := r.Form["width"]
	if isExist {
		width, _ = strconv.ParseInt(fWidth[0], 10, 64)
	}
	fColor, isExist := r.Form["color"]
	if isExist {
		hexStr = fColor[0]
	}

	cells = 100
	xyrange = 30.0
	xyscale = float64(width) / 2 / xyrange
	zscale = float64(height) * 0.4
	angle = math.Pi / 6
}

func polygon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>`)
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			// fmt.Printf("%g %g %g %g\n", z1, z2, z3, z4)
			if isFinite(ax) && isFinite(ay) &&
				isFinite(bx) && isFinite(by) &&
				isFinite(cx) && isFinite(cy) &&
				isFinite(dx) && isFinite(dy) {
				fmt.Fprintf(w,
					"<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#" + hexStr + "' />\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintln(w,"</svg>")
	fmt.Fprintln(w, `</body></html>`)
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)
	z := f(x, y)
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*float64(zscale)
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func isFinite(f float64) bool {
	return !math.IsInf(f, 0)
}


