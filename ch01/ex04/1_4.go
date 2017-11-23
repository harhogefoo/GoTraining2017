// Dup1は標準入力から2回以上現れる行を出現回数と一緒に表示します。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"io"
)

var out io.Writer = os.Stdout

func main() {
	exec(os.Args)
}

func exec(args []string) {
	if !hasArgs(args) {
		fmt.Fprintln(out,"usage: go run 1_4.go filename1 filename2 ... ")
		return
	}
	counts := make(map[string]int)
	dupFiles := make(map[string][]string)

	files := args[1:]
	for _, fileName := range files {
		f := openFile(fileName)
		countLines(f, counts, dupFiles)
		f.Close()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(out,"%d\t%s\t%s\n", n, line, strings.Join(dupFiles[line], ","))
		}
	}
}

func hasArgs(args []string) bool {
	return len(args) > 1
}

func openFile(fileName string) *os.File {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "1_4: %v\n", err)
	}
	return f
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
	// 注意: input.Err()からのエラーの可能性を無視している
}

// 配列に特定の文字列が入っていればtrue, そうでなければfalse
func arrayContains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
