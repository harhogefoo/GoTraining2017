package main

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html"
	"os"
	"bufio"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		fmt.Printf("[%s]\n=> Words : %d,  Images : %d\n", url, words, images)
	}
}

// CountWordsAndImagesはHTMLドキュメントに対するHTTP GETリクエストをurlへ
// 行い、そのドキュメント内に含まれる単語と画像の数を返します。
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	input := bufio.NewScanner(strings.NewReader(n.Data))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words++
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images = 1
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		wc, ic := countWordsAndImages(c)
		words += wc
		images += ic
	}

	return words, images
}
