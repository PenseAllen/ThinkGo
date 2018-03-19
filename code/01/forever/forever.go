package main

import (
	"fmt"
	"time"
)

const (
	timeFormat = "1/2 3:04:05 PM"
	remoteZone = "Pacific/Fiji"
	delay      = 1 * time.Second
)

func main() {
	remoteLoc, err := time.LoadLocation(remoteZone)
	if err != nil {
		panic(err)
	}
	for {
		now := time.Now()
		utc := now.In(time.UTC)
		local := now.In(time.Local)
		remote := now.In(remoteLoc)
		fmt.Print("UTC ", utc.Format(timeFormat), " | ")
		fmt.Print("Local ", local.Format(timeFormat), " | ")
		fmt.Println(remoteZone, remote.Format(timeFormat))
		time.Sleep(delay)
	}
}
