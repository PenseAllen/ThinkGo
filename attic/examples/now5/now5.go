package main

import (
	"fmt"
	"time"
)

const timeFormat = "Jan 02, 3:04:05 PM"

func main() {
	remoteLoc, _ := time.LoadLocation("Pacific/Fiji")
	for {
		now := time.Now()
		utc := now.In(time.UTC)
		local := now.In(time.Local)
		remote := now.In(remoteLoc)
		fmt.Print("UTC: ", utc.Format(timeFormat))
		fmt.Print(" | Local: ", local.Format(timeFormat))
		fmt.Print(" | Fiji: ", remote.Format(timeFormat))
		fmt.Println()
		time.Sleep(1 * time.Second)
	}
}
