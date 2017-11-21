// Echo1は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"testing"
	"bytes"
)

func TestPrint(t *testing.T) {
	var tests = []struct {
		args    []string
		want    string
	}{
		{[]string{}, ""},
		{[]string{}, ""},
		{[]string{"one", "two", "three"}, "0: one\n1: two\n2: three\n"},
		{[]string{"a", "b", "c"}, "0: a\n1: b\n2: c\n"},
		{[]string{"1", "2", "3"}, "0: 1\n1: 2\n2: 3\n"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("printArgsWithVerbose(%q", test.args)
		out = new(bytes.Buffer)
		printArgsWithVerbose(test.args)
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}