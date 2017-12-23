#!/bin/sh
go build variableNode.go
./fetch https://golang.org | ./variableNode
