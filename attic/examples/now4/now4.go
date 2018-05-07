package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		now := time.Now().In(time.Local)
		fmt.Println(now.Format("3:04:05 PM"))
		time.Sleep(1 * time.Second)
	}
}
