package main

import "testing"

func TestMax(t *testing.T) {
	for _, test := range []struct {
		vals []int
		valid bool
		expected int
	}{
		{nil, false, 0},
		{[]int{1}, true, 1},
		{[]int{1, 2, 4}, true, 4},
	} {
		m, err := max(test.vals...)
		if !test.valid {
			if err == nil {
				t.Errorf("err != nil expected for %v\n", test.vals)
			}
		}
		if m != test.expected {
			t.Errorf("m is %d, ut want %d\n", m, test.expected)
		}
	}
}

func TestMaxAtLeastOne(t *testing.T) {
	for _, test := range []struct {
		vals []int
		valid bool
		expected int
	}{
		{[]int{1}, true, 1},
		{[]int{1, 2, 4}, true, 4},
	} {
		m := maxAtLeastOne(test.vals[0], test.vals[1:]...)
		if m != test.expected {
			t.Errorf("m is %d, but want %d\n", m, test.expected)
		}
	}
}