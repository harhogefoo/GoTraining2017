// Echo1は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"testing"
	"bytes"
)

func TestExec(t *testing.T) {
	var tests = []struct {
		args    []string
		want    string
	}{
		{[]string{"", "dup1.txt", "dup2.txt"}, "10\thogehoge\tdup1.txt,dup2.txt\n2\tfugafuga\tdup1.txt,dup2.txt\n2\tpiyopiyo\tdup1.txt,dup2.txt\n"},
		{[]string{""}, "usage: go run 1_4.go filename1 filename2 ... \n"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("1_4(%q", test.args)
		out = new(bytes.Buffer)
		exec(test.args)
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}