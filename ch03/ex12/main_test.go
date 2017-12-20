package ex12

import "testing"

var data = []struct {
	str1 string
	str2 string
	expected bool
}{
	{"abc", "abc", true},
	{"aca", "aac", true},
	{"hogehoge", "gehogeho", true},
	{"Hello World", "World Hello", true},
	{"anagram", "margana", true},
	{"hogehoge", "fugafuga", false},
}

func execute(t *testing.T, f func(string, string) bool) {
	for _, d := range data {
		result := f(d.str1, d.str2)
		if result != d.expected {
			t.Errorf("Result is %v, want %v", result, d.expected)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	execute(t, isAnagram)
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			isAnagram(d.str1, d.str2)
		}
	}
}
