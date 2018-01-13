#!/bin/sh
go build main.go
./main -temp -18C
./main -temp 212F
./main -temp 273.15K
./main -help
