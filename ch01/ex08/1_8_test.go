// Echo1は、そのコマンドライン引数を表示します。
package main

import (
	"fmt"
	"testing"
)

func TestAddPrefixIfNeeded(t *testing.T) {
	var tests = []struct {
		target string
		prefix string
		want    string
	}{
		{"google.com", "http://", "http://google.com"},
		{"http://google.com", "http://", "http://google.com"},
	}

	for _, test := range tests {
		description := fmt.Sprintf("AddPrefixIfNeeded(%q, %q", test.target, test.prefix)
		got := AddPrefixIfNeeded(test.target, test.prefix)
		if got != test.want {
			t.Errorf("%s = %q, want %q)", description, got, test.want)
		}
	}
}