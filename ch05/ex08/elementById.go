package main

import (
	"net/http"
	"fmt"
	"os"
	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: findElement url id")
		os.Exit(1)
	}
	findElement(os.Args[1], os.Args[2])
}

func findElement(url, id string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	node := ElementById(doc, id)
	if node == nil {
		fmt.Printf("No Element with \"%s\" attribute Found\n", id)
	} else {
		printNode(node)
	}
}

func printNode(n *html.Node) {
	fmt.Printf("<%s", n.Data)
	for _, a := range n.Attr {
		fmt.Printf(" %s='%s'", a.Key, a.Val)
	}
	fmt.Println(">")
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}

	if post != nil {
		if !post(n) {
			return false
		}
	}
	return true
}

func ElementById(doc *html.Node, id string) *html.Node {
	var node *html.Node

	forEachNode(doc, func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}

		for _, a := range n.Attr {
			if a.Key == id {
				node = n
				return false
			}
		}
		return true
	}, nil)

	return node
}
