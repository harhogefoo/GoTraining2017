/*
 * mathパッケージの別の関数で可視化を試す
 */
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 420
	cells         = 500
	radius        = 200
	strokeWidth   = 2.0
)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		x1 := 200 + radius * math.Cos(float64(i))
		y1 := 200 + radius * math.Sin(float64(i))
		x2 := 200 + radius * math.Cos(float64(i) + strokeWidth)
		y2 := 200 + radius * math.Sin(float64(i) + strokeWidth)
		fmt.Printf("<polygon points='%g,%g %g,%g'/>\n", x1, y1, x2, y2)
	}
	fmt.Println("</svg>")
}
