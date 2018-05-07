package main

import (
	"fmt"
	"time"
)

const (
	numericFields = "1 2 3 4 5 06 -07"
  timeFormat = "Jan 02, 15:04:05"
)

func main() {
	remoteLoc, _ := time.LoadLocation("Pacific/Fiji")
	now := time.Now()
	utc := now.In(time.UTC)
	local := now.In(time.Local)
	remote := now.In(remoteLoc)
	fmt.Println("UTC   =", utc.Format(timeFormat))
	fmt.Println("Local =", local.Format(timeFormat))
	fmt.Println("Fiji  =", remote.Format(timeFormat))
}
