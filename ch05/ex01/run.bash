#!/bin/bash

go build fetch.go
go build findlinks1.go
./fetch https://golang.org | ./findlinks1
