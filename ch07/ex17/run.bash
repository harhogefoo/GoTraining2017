#!/bin/sh
go build fetch.go
go build xmlselect.go
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./xmlselect div div h2

echo ""
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./xmlselect div class="toc" h2

