package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io"
)

func main() {

	for i := 0; i < 2; i++ {
		fetchWithPrint()
	}
}

func fetchWithPrint() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // ゴルーチンを開始
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // chチャネルから受信
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // chチャネルへ送信
		return
	}

	filename := start.String() + ".txt"
	dst, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	nbytes, err := io.Copy(dst, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	resp.Body.Close() // 資源をリークさせない
	dst.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
