package main

import (
	"testing"
	"os"
)


func TestCountingWriter(t *testing.T) {
	data := []string{
		"hogehoge fugafuga\n",
		"How is the weather today??\n",
	}

	w, c := CountingWriter(os.Stdout)

	var total int64 = 0

	for _, d := range data {
		bytes := []byte(d)
		w.Write(bytes)
		total += int64(len(bytes))

		if *c != total {
			t.Errorf("count is %d, want %d", *c, total)
		}
	}
}

