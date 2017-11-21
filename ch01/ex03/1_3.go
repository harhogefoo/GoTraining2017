// Echo1は、そのコマンドライン引数を表示します。
package main

import (
	"strings"
	"time"
	"io"
	"fmt"
	"os"
)

var out io.Writer = os.Stdout

func plusEqual(array []string) float64 {
	start := time.Now()
	var output string
	for i := 0; i < len(array); i++ {
		output += array[i]
	}
	seconds := time.Since(start).Seconds()
	fmt.Fprintf(out,"%.10fs elapsed\n", seconds)
	return seconds
}

func join(array []string) float64 {
	start := time.Now()
	strings.Join(array, "")
	seconds := time.Since(start).Seconds()
	fmt.Fprintf(out,"%.10fs elapsed\n", seconds)
	return seconds
}

func createArray() []string {
	array := make([]string, 10000)
	for i := 0; i < len(array); i++ {
		array[i] = "hogehoge"
	}
	return array
}

func main() {
	array := createArray()
	// +=
	plusEqual(array)
	// Join
	join(array)
}
