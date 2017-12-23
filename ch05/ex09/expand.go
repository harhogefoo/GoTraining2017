package main

import "regexp"

var pattern = regexp.MustCompile(`(\$\w*)`)

func expand(s string, f func(string) string) string {
	return pattern.ReplaceAllStringFunc(s,
		func(sub string) string {
			return f(sub[1:])
		})
}