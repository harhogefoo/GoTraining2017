package main

import (
	"testing"
)

func TestExpand(t *testing.T) {
	test := func(s string) string { return "hogehoge" }

	for _, test := range []struct {
		s        string
		f        func(string) string
		expected string
	}{
		{"$test", test, "hogehoge"},
		{"$test $piyo", test, "hogehoge hogehoge"},
		{"$test test", test, "hogehoge test"},
	} {
		result := expand(test.s, test.f)
		if result != test.expected {
			t.Errorf("%s, but want %s\n", result, test.expected)
		}
	}
}