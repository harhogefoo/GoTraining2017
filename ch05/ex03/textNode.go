package main

import (
	"os"
	"golang.org/x/net/html"
	"fmt"
	"strings"
)

func main() {
	var textNode []string
	elem, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "text: %v\n", err)
		os.Exit(1)
	}

	printTextNode(elem, nil)
}

func printTextNode(n *html.Node, textNode []string) {
	if n == nil {
		return
	}
	switch n.Type {
	case html.ElementNode:
		textNode = append(textNode, n.Data)
	case html.TextNode:
		length := len(textNode)
		last := textNode[length-1]
		if last != "script" && last != "style" {
			trim := strings.TrimSpace(n.Data)
			if len(trim) > 0 {
				fmt.Printf("<%s>%s<%s>\n", last, n.Data, last)
			}
		}
	}

	printTextNode(n.FirstChild, textNode)
	printTextNode(n.NextSibling, textNode)
}
