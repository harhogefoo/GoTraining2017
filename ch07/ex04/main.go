package main

import (
	"golang.org/x/net/html"
	"fmt"
	"os"
	"io"
)

func main() {
	if !Parse(os.Args[1]) {
		os.Exit(1)
	}
}

func Parse(contents string) bool {
	doc, err := html.Parse(newReader(contents))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		return false
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	return true
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}

type reader struct {
	bytes []byte
	next int
}

func (r *reader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	if r.next >= len(r.bytes) {
		return 0, io.EOF
	}

	bytes := len(r.bytes) - r.next
	if bytes > len(p) {
		bytes = len(p)
	}

	copy(p, r.bytes[r.next:r.next+ bytes])
	r.next += bytes
	return bytes, nil
}

func newReader(contents string) io.Reader {
	return &reader{[]byte(contents), 0}
}
