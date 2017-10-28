package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		dstName := "out.txt"
		dst, err := os.Create(dstName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(dst, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		resp.Body.Close()
		dst.Close()
	}
}
