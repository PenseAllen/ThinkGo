package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Print(time.Now(), " | ")
		time.Sleep(20 * time.Millisecond)
	}
}
