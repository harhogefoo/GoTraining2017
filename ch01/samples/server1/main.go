package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler) // 個々のリクエストに対してhandlerが呼ばれる
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

// handlerはリクエストされたURL r のPath要素を返します。
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
