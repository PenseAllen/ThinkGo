package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Print("╱╲╱") // better: ╱╲
		time.Sleep(5 * time.Millisecond)
	}
}
