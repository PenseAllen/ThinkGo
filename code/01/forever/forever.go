package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		fmt.Print(os.Args[1])
		time.Sleep(10 * time.Millisecond)
	}
}
