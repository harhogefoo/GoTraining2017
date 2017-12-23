#!/bin/sh
go build elemCounter.go
./fetch https://golang.org | ./elemCounter
