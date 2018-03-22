package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now().In(time.Local))
		time.Sleep(1 * time.Second)
	}
}
