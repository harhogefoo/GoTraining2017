#!/bin/sh

go build main.go
./main sample.tar tar
./main sample.zip zip
