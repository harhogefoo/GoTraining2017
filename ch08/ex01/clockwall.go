package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"strings"
	"os"
)

type clock struct {
	location string
	url string
}

var times = make([]string, 3)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("illegal argument.")
		return
	}

	clocks := make([]clock, 0)
	for _, arg := range os.Args[1:] {
		parsedArg := strings.Split(arg, "=")
		if len(parsedArg) != 2 {
			fmt.Println("illegal argument.")
			return
		} else {
			clocks = append(clocks, clock{parsedArg[0],parsedArg[1] })
		}
	}

	for i, clock := range clocks {
		go connect(clock, i)
	}

	select {}  // infinite loop
}

func connect(c clock, index int) {
	conn, err := net.Dial("tcp", c.url)
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
		times[index] = fmt.Sprintf("%s[%s]", c.location, bytes)
		fmt.Printf("\r%s", strings.Join(times, " "))
	}
}

