package main

import (
	"fmt"
	"time"
)

const (
	timeFormat = "1/2 3:08:05 PM"
	remoteZone = "Pacific/Fiji"
	delay      = 1 * time.Second
)

func main() {
	remoteLoc, _ := time.LoadLocation(remoteZone)
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
