/*
 * 高さに基づいて個々のポリゴンに色付け
 */
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z1 := corner(i+1, j)
			bx, by, z2 := corner(i, j)
			cx, cy, z3 := corner(i, j+1)
			dx, dy, z4 := corner(i+1, j+1)
			// fmt.Printf("%g %g %g %g\n", z1, z2, z3, z4)
			isRed := isSignedHeight(z1, z2, z3, z4)
			color := "#FF0000"
			if !isRed {
				color = "#0000FF"
			}
			if isFinite(ax) && isFinite(ay) &&
				isFinite(bx) && isFinite(by) &&
				isFinite(cx) && isFinite(cy) &&
				isFinite(dx) && isFinite(dy) {
				fmt.Printf(
					"<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='" + color + "' />\n",
						ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func isSignedHeight(z1, z2, z3, z4 float64) bool {
	z := (z1 + z2 + z3 + z3) / 4
	fmt.Printf("%g\n", z)
	if z > 0 {
		return true
	}
	return false
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func isFinite(f float64) bool {
	return !math.IsInf(f, 0)
}