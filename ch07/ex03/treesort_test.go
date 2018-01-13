package treesort

import (
	"testing"
	"sort"
	"math/rand"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}


func TestString(t *testing.T) {
	data := []struct {
		values []int
	}{
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
		{[]int{8, 9, 6, 7, 4, 5, 2, 3, 0, 1}},
		{[]int{5, 6, 7, 8, 9, 0, 1, 2, 3, 4}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	expected := "{0 1 2 3 4 5 6 7 8 9}"

	for _, d := range data {
		var root *tree
		for _, v := range d.values {
			root = add(root, v)
		}

		if root.String() != expected {
			t.Errorf("Result is %s, but want %s", root.String(), expected)
		}
	}
}