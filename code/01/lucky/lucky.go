package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Your lucky number is:")
	fmt.Println(rand.Intn(100) + 1)
}
