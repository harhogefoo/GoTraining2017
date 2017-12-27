package main

import "testing"

func TestMin(t *testing.T) {
	for _, test := range []struct {
		vals []int
		valid bool
		expected int
	}{
		{nil, false, 0},
		{[]int{1}, true, 1},
		{[]int{1, 2, 4}, true, 1},
	} {
		m, err := min(test.vals...)
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

func TestMinAtLeastOne(t *testing.T) {
	for _, test := range []struct {
		vals []int
		valid bool
		expected int
	}{
		{[]int{1}, true, 1},
		{[]int{1, 2, 4}, true, 1},
	} {
		m := minAtLeastOne(test.vals[0], test.vals[1:]...)
		if m != test.expected {
			t.Errorf("m is %d, but want %d\n", m, test.expected)
		}
	}
}