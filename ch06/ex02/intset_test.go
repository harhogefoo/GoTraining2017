package main

import "testing"

func TestLen(t *testing.T) {
	for _, tc := range []struct {
		values   []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 144, 9, 42}, 4},
		{[]int{1, 32, 32 << 1, 32 << 2, 32 << 3, 32 << 8}, 6},
	} {
		var x IntSet
		for _, v := range tc.values {
			x.Add(v)
		}
		if x.Len() != tc.expected {
			t.Errorf("x.Len() is %d, but want %v", x.Len(), tc.expected)
		}
	}

	var x IntSet
	if x.Len() != 0 {
		t.Errorf("x.Len() is %d, but want 0", x.Len())
	}
}

func TestRemove(t *testing.T) {
	var x IntSet

	// Removes non-existing value
	x.Remove(1000)
	if x.Len() != 0 {
		t.Errorf("x.Len() is %d, but want 0", x.Len())
	}

	const max = 100000
	for i := 0; i < max; i++ {
		x.Add(i)
	}

	for i := 0; i < max; i++ {
		x.Remove(i)
		if x.Has(i) {
			t.Errorf("x.Has(%d) is true, but want false", i)
			continue
		}
		if x.Len() != (max - i - 1) {
			t.Errorf("x.Len() is %d, but want %d", x.Len(), max-i-1)
		}
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	const max = 100000
	for i := 0; i < max; i++ {
		x.Add(i)
	}
	x.Clear()

	if x.Len() != 0 {
		t.Errorf("x.Len() is %d, but want 0", x.Len())
	}
}

func TestCopy(t *testing.T) {
	var x IntSet
	const max = 100000
	for i := 0; i < max; i++ {
		x.Add(i)
	}

	c := x.Copy()
	x.Clear()

	if x.Len() != 0 {
		t.Errorf("x.Len() is %d, but want 0", x.Len())
	}

	for i := 0; i < max; i++ {
		if x.Has(i) {
			t.Errorf("c.Has(%d) is true, but want false", i)
		}
		if !c.Has(i) {
			t.Errorf("c.Has(%d) is false, but want true", i)
		}
	}
}


func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll()
	if x.Len() != 0 {
		t.Errorf("x.Len() is %d, but want 0", x.Len())
	}

	x.AddAll(1, 2, 3, 4, 5)
	if x.Len() != 5 {
		t.Errorf("x.Len() is %d, but want 5", x.Len())
	}

	for i := 1; i <= 5; i++ {
		if !x.Has(i) {
			t.Errorf("x.Has(%d) is false, but want true", i)
		}
	}
}

