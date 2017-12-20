package ex11

import "strings"

func Comma(s string) string {
	if len(s) == 0 {
		return s
	}

	if s[0:1] == "+" || s[0:1] == "-" {
		return s[0:1] + Comma(s[1:])
	}

	dotIndex := strings.IndexByte(s, '.')
	if dotIndex >= 0 {
		return Comma(s[:dotIndex]) + s[dotIndex:]
	}

	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}
