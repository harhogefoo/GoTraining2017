#!/bin/sh

go build fetchall.go
./fetchall https://golang.org http://gopl.io https://godoc.org
