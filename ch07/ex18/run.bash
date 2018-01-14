#!/bin/sh
go build fetch.go
go build xml.go
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 > sample.xml

./xml < sample.xml > out.txt
