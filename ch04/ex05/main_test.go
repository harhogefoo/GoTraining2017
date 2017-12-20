package ex05

import (
	"testing"
	"crypto/sha256"
	"fmt"
)

var data = []struct {
	s        []string
	expected []string
}{
	{
		[]string{},
		[]string{},
	},
	{
		[]string{"Hello"},
		[]string{"Hello"},
	},
	{
		[]string{"Hello", "World", "World"},
		[]string{"Hello", "World"},
	},
	{
		[]string{"Hello", "Hello", "World", "World"},
		[]string{"Hello", "World", "Hello", "World"},
	},
}

func TestRemoveNeighbor(t *testing.T) {
	for _, d := range data {
		result := removeNeighborDup(d.s)
		if len(result) != len(d.expected) {
			t.Errorf("Result length is %d, want %d",
				len(result), len(d.expected))
		}
		for i := 0; i < len(d.expected); i++ {
			if result[i] != d.expected[i] {
				t.Errorf("result[%d] is %s, want %s",
					i, result[i], d.expected[i])
			}
		}
	}
}
