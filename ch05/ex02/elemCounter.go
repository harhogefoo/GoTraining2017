package main

import (
	"os"
	"golang.org/x/net/html"
	"fmt"
)

func main() {
	elemCount := make(map[string]int)

	elem, err := html.Parse(os.Stdin)
	fmt.Println(elem)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elemCounter: %v\n", err)
		os.Exit(1)
	}

	visitElem(elem, elemCount)

	for e, c := range elemCount {
		fmt.Printf("%10s: %d\n", e, c)
	}
}

func visitElem(n *html.Node, elemCount map[string]int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		elemCount[n.Data]++
	}

	visitElem(n.FirstChild, elemCount)
	visitElem(n.NextSibling, elemCount)
}
