package ex07

import "testing"

func TestReverseUtf8(t *testing.T) {
	data := []struct {
		s        string
		expected string
	}{
		{
			"abcdefgh",
			"hgfedcba"},
		{
			"ほげほげふがふが",
			"がふがふげほげほ"},
	}

	for _, d := range data {
		b := []byte(d.s)
		ReverseUtf8(b)
		result := string(b)
		if result != d.expected {
			t.Errorf("Result is %s, want %s", result, d.expected)
		}
	}
}
