package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Print("╱╲╱") // even better: ╱╲╱
		time.Sleep(5 * time.Millisecond)
	}
}
