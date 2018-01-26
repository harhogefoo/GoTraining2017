package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png" // register PNG decoder
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	img, err := read(os.Stdin)
	if err != nil {
		fmt.Println("Error")
		return
	}

	switch os.Args[1] {
	case "jpg", "jpeg":
		toJPEG(img, os.Stdout)
	case "png":
		toPNG(img, os.Stdout)
	default:
		usage()
	}
}

func read (in io.Reader) (image.Image, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return nil, err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return img, nil

}

func usage () {
	fmt.Println("Usage: [png|jpg|jpeg]")
	os.Exit(1)
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}

//!-main

/*
//!+with
$ go build gopl.io/ch3/mandelbrot
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
Input format = png
//!-with

//!+without
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
jpeg: image: unknown format
//!-without
*/