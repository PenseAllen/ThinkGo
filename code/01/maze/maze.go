package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		if rand.Intn(2) == 0 {
			fmt.Print("/")
		} else {
			fmt.Print("\\")
		}
		time.Sleep(5 * time.Millisecond)
	}
}
