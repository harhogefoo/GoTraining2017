package main

import (
	"golang.org/x/net/html"
	"testing"
	"os"
)

func TestElementsByTagName(t *testing.T) {
	f, err := os.Open("index.html")
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	defer f.Close()

	doc, err := html.Parse(f)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	if len(images) != 29 {
		t.Errorf("len(%v) is not 29", images)
	}

	if len(headings) != 28 {
		t.Errorf("len(%v) is not 28", headings)
	}
}




