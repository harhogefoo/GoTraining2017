package ex03

const Size = 10

func reverse(a *[Size]int) {
	for i, j := 0, Size-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
