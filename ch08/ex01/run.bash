#!/bin/sh
go build clock2.go
go build clockwall.go
TZ=US/Estern     ./clock2 -port 8010 &
TZ=Asia/Tokyo    ./clock2 -port 8020 &
TZ=Europe/London ./clock2 -port 8030 &

./clockwall

# killall clock2
