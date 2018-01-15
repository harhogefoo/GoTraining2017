#!/bin/sh
go build clock2.go
go build clockwall.go
TZ=US/Eastern     ./clock2 -port 8010 &
TZ=Asia/Tokyo    ./clock2 -port 8020 &
TZ=Europe/London ./clock2 -port 8030 &

./clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030

# killall clock2
