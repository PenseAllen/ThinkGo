package main

import (
	"fmt"
	"time"
)

func main() {
start:
	fmt.Print("J. Random Hacker ")
	time.Sleep(10 * time.Millisecond)
	goto start
}
