package main

import (
	"time"
	"fmt"
	"strconv"
)

// pc[i] は i のポピュレーションカウントです。
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountはxのポピュレーションカウント(1が設定されているビット数)を返します
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountv2(x uint64) int {
	var result byte
	for i := uint(0); i < 8; i++ {
		result += pc[byte(x>>(i*8))]
	}
	return int(result)
}

func main() {
	x := uint64(1000)

	start := time.Now()
	fmt.Println(strconv.Itoa(PopCount(x)))
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(strconv.Itoa(PopCountv2(x)))
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())
}
