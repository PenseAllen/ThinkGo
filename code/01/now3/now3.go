package main

import (
	"fmt"
	"time"
)

const (
	timeFormat = "1 2 3 4 5 06 -07"
)

func main() {
	now := time.Now()
	utc := now.In(time.UTC)
	local := now.In(time.Local)
	remoteLoc, _ := time.LoadLocation("Pacific/Fiji")
	remote := now.In(remoteLoc)
	fmt.Println("UTC   =", utc.Format(timeFormat))
	fmt.Println("Local =", local.Format(timeFormat))
	fmt.Println("Fiji  =", remote.Format(timeFormat))
}
