package main

import (
	"os"
	"golang.org/x/net/html"
	"fmt"
)

func main() {
	elem, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "text: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, elem) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode {
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		if n.Data == "img" || n.Data == "script" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					fmt.Println(a.Val)
					links = append(links, a.Val)
				}
			}
		}
		if n.Data == "link" {
			isStyleSheet := false
			var href string
			for _, a := range n.Attr {
				if a.Key == "rel" && a.Val == "stylesheet" {
					isStyleSheet = true
				}
				if a.Key == "href" {
					href = a.Val
				}
			}

			if isStyleSheet {
				links = append(links, href)
			}
		}
	}
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}
