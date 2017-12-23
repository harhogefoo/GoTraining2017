#!/bin/sh
go build textNode.go
./fetch https://golang.org | ./textNode
