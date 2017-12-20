package ex10

import "testing"

var data = []struct {
	input string
	expected string
}{
	{"", ""},
	{"1", "1"},
	{"111", "111"},
	{"1111", "1,111"},
	{"111111", "111,111"},
	{"1111111", "1,111,111"},
	{"1111111111", "1,111,111,111"},
}

func execute(t *testing.T, f func(string) string) {
	for _, d := range data {
		result := f(d.input)
		if result != d.expected {
			t.Errorf("Result is %s, want %s", result, d.expected)
		}
	}
}

func TestComma(t *testing.T) {
	execute(t, Comma)
}

func TestComma2(t *testing.T) {
	execute(t, Comma2)
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			Comma(d.input)
		}
	}
}

func BenchmarkComma2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			Comma2(d.input)
		}
	}
}