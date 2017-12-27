package main

import(
	"fmt"
	"math"
)

func main() {
	fmt.Printf("square(2) = %d\n", square(2))
	fmt.Printf("log10(100) = %f\n", log10(100))
}

func log10(n float64) (r float64) {
	defer func() {
		if p := recover(); p != nil {
			r = p.(float64)
		}
	}()
	panic(math.Log10(n))
}

func square(n int) (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = p.(int)
		}
	}()
	panic(n * n)
}
