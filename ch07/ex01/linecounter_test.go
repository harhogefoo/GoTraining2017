package main

import (
	"testing"
)

func TestLineCounter(t *testing.T) {
	data := []struct {
		s        string
		expected int
	}{
		{"今日はいい天気だ", 1},
		{"今日はいい天気だ\nな", 2},
		{"今日はいい天気だ\nな\nも", 3},
	}

	var c LineCounter
	for _, d := range data {
		c = 0

		bytes := []byte(d.s)
		n, err := c.Write(bytes)

		if err != nil {
			t.Errorf("Unpexected Error : %v", err)
			continue
		}

		if n != len(bytes) {
			t.Errorf("Written bytes is %d, want %d", n, len(bytes))
			continue
		}

		if c != LineCounter(d.expected) {
			t.Errorf("Result is %d, want %d", c, d.expected)
		}
	}
}