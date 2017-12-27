package main

import "fmt"

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("引数がありません。")
	}
	m := vals[0]
	for _, v := range vals[1:] {
		if v < m {
			m = v
		}
	}
	return m, nil
}

func minAtLeastOne(one int, vals ...int) int {
	min := one

	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}
