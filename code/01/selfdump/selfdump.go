package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

const limit = 60
const bytesPerLn = 20

func main() {
	fmt.Print("Contents of: ", os.Args[0])
	data, err := ioutil.ReadFile(os.Args[0])
	if err != nil {
		panic(err)
	}
	stop := len(data)
	if stop > limit {
		stop = limit
	}
	for index, octet := range data[:stop] {
		if index % bytesPerLn == 0 {
			fmt.Println()
		}
		fmt.Printf("%02x ", octet)
	}
	fmt.Println()
}

