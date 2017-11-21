// Echo1は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"os"
	"io"
)


var out io.Writer = os.Stdout // テスト中は変更される

func printArgsWithVerbose(args []string) {
	for i := 0; i < len(args); i++ {
		fmt.Fprintf(out, "%d: %s\n", i, args[i])
	}
}

func main() {
	printArgsWithVerbose(os.Args)
}