#!/bin/sh

go build main.go
main -s / a bc def
main -n a bc def
main -help