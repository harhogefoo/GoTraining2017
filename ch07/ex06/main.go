package main

import (
	"github.com/harhogefoo/go_training2017/ch07/ex06/tempflag"
	"flag"
	"fmt"
)

var temp = tempflag.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

