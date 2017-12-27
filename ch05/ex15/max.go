package main

import "fmt"

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("引数がありません。")
	}
	m := vals[0]
	for _, v := range vals[1:] {
		if v > m {
			m = v
		}
	}
	return m, nil
}

func maxAtLeastOne(one int, vals ...int) int {
	max := one

	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}
