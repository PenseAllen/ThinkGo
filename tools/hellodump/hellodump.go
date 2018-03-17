package main

import (
	"fmt"
	"io/ioutil"
)

const bytesPerLine = 20

func main() {
	data, err := ioutil.ReadFile("hello.elf")
	if err != nil {
		panic(err)
	}
	start := 0x083180  // determined with Hopper Disassembler
	stop := 0x0831e7
	for index, octet := range data[start:stop] {
		fmt.Printf("%02x ", octet)
		if index % bytesPerLine == bytesPerLine - 1 {
			fmt.Println()
		}
	}
	fmt.Println()
}

