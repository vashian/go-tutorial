package main

import (
	"fmt"
	"strconv"
)

func main() {
	// %b verb prints binary
	fmt.Printf("%b\n", 0)
	fmt.Printf("%b\n", 1)

	// %02b prints 2 bits with leading zeros if any
	fmt.Printf("%08b = %d\n", 0, 0)
	fmt.Printf("%08b = %d\n", 128, 128)
	fmt.Printf("%08b = %d\n", 255, 255)

	i, _ := strconv.ParseInt("00000010", 2, 64)
	fmt.Println(i)
}