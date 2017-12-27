package main

import "testing"

var tests = []struct {
	strings []string
	sep string
	expected string
}{
	{nil, "", ""},
	{[]string{"hoge"}, " ", "hoge"},
	{[]string{"hoge", "hoge"}, " ", "hoge hoge"},
	{[]string{"hoge", "hoge", "hoge"}, ",", "hoge,hoge,hoge"},
}

func TestJoin(t *testing.T) {
	for _, test := range tests {
		result := Join(test.sep, test.strings...)
		if result != test.expected {
			t.Errorf("Join(%s, %v) = %s, want %s",
				test.sep, test.strings, result, test.expected)
		}
	}
}
