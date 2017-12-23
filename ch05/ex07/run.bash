#!/bin/sh
go build htmlprint.go
./htmlprint http://gopl.io > out.html
