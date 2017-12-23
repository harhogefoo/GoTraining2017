package main

import (
	"golang.org/x/net/html"
	"fmt"
	"os"
	"net/http"
	"bytes"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startNode, endNode)
	return nil
}

// forEachNodeはnから始まるツリー内の個々のノードxに対して
// 関数pre(x)とpost(x)を呼び出します。その2つの関数はオプションです。
// pre は子ノードを訪れる前に呼び出され(前順: preorder)、
// post は子ノードを訪れた後に呼び出されます(後順: postorder)。
func forEachNode(n *html.Node,
	pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startNode(n *html.Node) {
	switch n.Type {
	case html.ErrorNode:
	case html.TextNode:
		startTextNode(n)
		return
	case html.DocumentNode:
	case html.ElementNode:
		startElementNode(n)
		return
	case html.CommentNode:
	case html.DoctypeNode:
	}
}

func startTextNode(n *html.Node) {
	text := strings.TrimSpace(n.Data)
	if text != "" {
		fmt.Printf("%s", text)
	}
}

func startElementNode(n *html.Node) {
	depth++
	if n.FirstChild == nil {
		return
	}

	attrs := attributes(n.Attr)
	if attrs == "" {
		fmt.Printf("\n%*s<%s>", depth, "", n.Data)
	} else {
		fmt.Printf("\n%*s<%s %s>", depth, "", n.Data, attrs)
	}
}

func endElementNode(n *html.Node) {
	if n.FirstChild == nil {
		attrs := attributes(n.Attr)
		if attrs == "" {
			switch n.Data {
			case "br":
				fmt.Printf("<%s/>\n", n.Data)
			default:
				fmt.Printf("\n%*s<%s />", depth, "", n.Data)
			}
		} else {
			fmt.Printf("\n%*s<%s %s />", depth, "", n.Data, attrs)
		}
	} else {
		switch n.Data {
		case "a", "code", "title", "tt", "h1":
			fmt.Printf("</%s>", n.Data)
		default:
			fmt.Printf("\n%*s</%s>", depth, "", n.Data)
		}
	}
	depth--
}

func endNode(n *html.Node) {
	switch n.Type {
	case html.ErrorNode:
	case html.TextNode:
		return
	case html.DocumentNode:
	case html.ElementNode:
		endElementNode(n)
		return
	case html.CommentNode:
	case html.DoctypeNode:
	}
}

func attributes(attr []html.Attribute) string {
	var buf bytes.Buffer

	for i, a := range attr {
		if i != 0 {
			buf.WriteString(" ")
		}
		if a.Namespace == "" {
			buf.WriteString(a.Key)
			buf.WriteString(`="`)
			buf.WriteString(a.Val)
			buf.WriteString(`"`)
		} else {
			buf.WriteString(a.Namespace)
			buf.WriteString(":")
			buf.WriteString(a.Key)
			buf.WriteString(`="`)
			buf.WriteString(a.Val)
			buf.WriteString(`"`)
		}
	}
	return buf.String()
}
