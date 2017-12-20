package ex10

import "bytes"

func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

func Comma2(s string) string {
	n := len(s)
	start, end := 0, n % 3
	if end == 0 {
		start, end = 0, 3
	}

	buffer := bytes.NewBuffer(make([]byte, 0, n + (n - 1) / 3))

	for end <= n {
		buffer.WriteString(s[start:end])
		if end < n {
			buffer.WriteString(",")
		}
		start, end = end, end + 3
	}
	return buffer.String()
}