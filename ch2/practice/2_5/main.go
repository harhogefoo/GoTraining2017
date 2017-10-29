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

func PopCountV2(x uint64) int {
	var result byte
	for i := uint(0); i < 8; i++ {
		result += pc[byte(x>>(i*8))]
	}
	return int(result)
}

func PopCountV3(x uint64) int {
	var result int
	for i := uint(0); i < 64; i++ {
		result += int(byte(x&1))
		x = x >> 1
	}
	return result
}

func PopCountV4(x uint64) int {
	result := 0
	for x > 0 {
		if x != x&(x-1) {
			result++
			x = x & (x - 1)
		}
	}
	return result
}

func main() {
	x := uint64(1001)

	start := time.Now()
	fmt.Println(strconv.Itoa(PopCount(x)))
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(strconv.Itoa(PopCountV2(x)))
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(strconv.Itoa(PopCountV3(x)))
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(strconv.Itoa(PopCountV4(x)))
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())
}
