package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"strings"
)

var times = make([]string, 3)

func main() {
	addresses := []string{
		"localhost:8010",
		"localhost:8020",
		"localhost:8030",
	}

	for i, address := range  addresses {
		go connect(address, i)
	}

	select {}  // infinite loop
}

func connect(address string, index int) {
	fmt.Printf("connect: %s\n", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		times[index] = fmt.Sprintf("%s[%s]", address, bytes)
		fmt.Printf("\r%s", strings.Join(times, " "))
	}
}

