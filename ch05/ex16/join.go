package main

import "bytes"

func Join(sep string, strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	if len(strings) == 1 {
		return strings[0]
	}

	totalBytes := 0
	for _, s := range strings {
		totalBytes += len(s)
	}
	totalBytes += len(sep) * (len(strings) - 1)

	b := bytes.NewBuffer(make([]byte, 0, totalBytes))
	b.WriteString(strings[0])
	for _, s := range strings[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}
