#!/bin/sh
go build mandelbrot.go
go build converter.go
./mandelbrot | ./converter jpg > mandelbrot.jpg
./mandelbrot | ./converter jpeg > mandelbrot.jpeg
./mandelbrot | ./converter png > mandelbrot.png
./mandelbrot | ./converter mov

