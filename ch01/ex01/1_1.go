// Echo1は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"os"
	"strings"
	"io"
)

var out io.Writer = os.Stdout // テスト中は変更される

func printArgs(args []string) {
	fmt.Fprintln(out, strings.Join(args, " "))
}
func main() {
	print(os.Args)
}