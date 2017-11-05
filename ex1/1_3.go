// Echo1は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	array := make([]string, 10000)
	for i := 0; i < len(array); i++ {
		array[i] = "hogehoge"
	}

	// +=
	start := time.Now()
	var output string
	for i := 0; i < len(array); i++ {
		output += array[i]
	}
	// fmt.Println(output)
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())

	// Join
	start = time.Now()
	output = strings.Join(array, "")
	// fmt.Println(output)
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())
}
