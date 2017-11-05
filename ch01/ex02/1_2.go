// Echo1は、そのコマンドライン引数を表示します。
package ex02

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%d: %s\n", i, os.Args[i])
	}
}