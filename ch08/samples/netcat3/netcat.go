package main

import (
	"net"
	"log"
	"io"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // メインゴルーチンへ通知

	}()
}
