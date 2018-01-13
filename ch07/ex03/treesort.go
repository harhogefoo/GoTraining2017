package treesort

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	return "{" + t.toString() + "}"
}

func (t *tree) toString() string {
	if t == nil {
		return ""
	}

	var result string

	if t.left != nil {
		result = t.left.toString() + " "
	}

	result = fmt.Sprintf("%s%d", result, t.value)

	if t.right != nil {
		result = result + " " + t.right.toString()
	}
	return result
}