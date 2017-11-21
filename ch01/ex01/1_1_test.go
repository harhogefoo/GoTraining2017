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
		{[]string{}, "\n"},
		{[]string{}, "\n"},
		{[]string{"one", "two", "three"}, "one two three\n"},
		{[]string{"a", "b", "c"}, "a b c\n"},
		{[]string{"1", "2", "3"}, "1 2 3\n"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("printArgs(%q", test.args)
		out = new(bytes.Buffer)
		printArgs(test.args)
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}