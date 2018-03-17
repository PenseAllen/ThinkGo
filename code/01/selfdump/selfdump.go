package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

const showLines = 5
const bytesPerLine = 20

func main() {
	fmt.Print("Contents of: ", os.Args[0])
	data, err := ioutil.ReadFile(os.Args[0])
	if err != nil {
		panic(err)
	}
	start := len(data) / 2
	stop := start + showLines * bytesPerLine
	if stop > len(data) {
		stop = len(data)
	}
	for index, octet := range data[start:stop] {
		if index % bytesPerLine == 0 {
			fmt.Println()
		}
		fmt.Printf("%02x ", octet)
	}
	fmt.Println()
}

