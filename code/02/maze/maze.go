package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		if rand.Intn(2) == 0 {
			fmt.Print("╱") // BOX DRAWINGS LIGHT DIAGONAL UPPER RIGHT TO LOWER LEFT
		} else {
			fmt.Print("╲") // BOX DRAWINGS LIGHT DIAGONAL UPPER LEFT TO LOWER RIGHT
		}
		time.Sleep(5 * time.Millisecond)
	}
}
