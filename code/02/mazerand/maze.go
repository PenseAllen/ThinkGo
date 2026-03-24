package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		var line string
		if rand.Intn(2) == 0 {
			line = "╱" // BOX DRAWINGS LIGHT DIAGONAL UPPER RIGHT TO LOWER LEFT
		} else {
			line = "╲" // BOX DRAWINGS LIGHT DIAGONAL UPPER LEFT TO LOWER RIGHT
		}
		fmt.Print(line)
		time.Sleep(5 * time.Millisecond)
	}
}
