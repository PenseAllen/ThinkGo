package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now())
		time.Sleep(1 * time.Second)
	}
}
