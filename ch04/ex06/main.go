package main

import (
	"unicode/utf8"
	"unicode"
)

func squashSpaces(b []byte) []byte {
	if len(b) == 0 {
		return b
	}

	spaceBuf := make([]byte, 4)
	spaceSize := utf8.EncodeRune(spaceBuf, ' ')
	spaceBuf = spaceBuf[:spaceSize]
	inSpace := false

	current := 0
	var size int
	var r rune
	for next := 0; next < len(b); next += size {
		r, size = utf8.DecodeRune(b[next:])

		if r == utf8.RuneError {
			panic("Rune Error")
		}

		if unicode.IsSpace(r) {
			if !inSpace {
				copy(b[current:], spaceBuf)
				current += spaceSize
				inSpace = true
			}
			continue
		}

		copy(b[current:], b[next:next+size])
		current += size
		inSpace = false
	}

	return b[:current]
}
