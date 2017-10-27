// Dup1は標準入力から2回以上現れる行を出現回数と一緒に表示します。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if !hasArgs() {
		fmt.Println("usage: go run 1_4.go filename1 filename2 ... ")
		return
	}
	counts := make(map[string]int)
	dupFiles := make(map[string][]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, dupFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, dupFiles)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(dupFiles[line], ","))
		}
	}
}

func hasArgs() bool {
	return len(os.Args) < 1
}

func countLines(f *os.File, counts map[string]int, dupFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text]++

		if !arrayContains(dupFiles[text], f.Name()) {
			dupFiles[text] = append(dupFiles[text], f.Name())
		}
	}
	// 注意: input.Err()からのエラーの可能性を蒸ししている
}

func arrayContains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
