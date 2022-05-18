package main

import "fmt"

func main() {
	intNumber := 5
	floadNumber := 4.5

	intNumber = intNumber * int(floadNumber)

	fmt.Println(intNumber)

	intNumber = 5
	floadNumber = floadNumber * float64(intNumber)

	fmt.Println(floadNumber)
}