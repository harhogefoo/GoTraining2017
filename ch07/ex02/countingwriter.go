package main

import "io"

type Wrapper struct {
	c int64
	w io.Writer
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var wp Wrapper
	wp.w = w
	return &wp, &(wp.c)
}

func (wp *Wrapper) Write(b []byte) (n int, err error) {
	n, err = wp.w.Write(b)
	wp.c += int64(n)
	return
}

