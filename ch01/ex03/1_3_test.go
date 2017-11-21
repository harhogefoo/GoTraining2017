// Echo1は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"testing"
	"bytes"
)

func TestPrint(t *testing.T) {
	for i := 0; i < 100; i++ {
		out = new(bytes.Buffer)
		array := createArray()
		plusEqualsSecond := plusEqual(array)
		joinSecond := join(array)

		if plusEqualsSecond < joinSecond {
			t.Error("no way!")
		}

		want := fmt.Sprintf("%.10fs elapsed\n%.10fs elapsed\n", plusEqualsSecond, joinSecond)
		got := out.(*bytes.Buffer).String()
		if got != want {
			t.Errorf("%q, want %q)", got, want)
		}
	}
}