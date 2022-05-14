package main

import (
	"fmt"
)

// there's no warning for package-level vars
var packageLevelVar string

func main() {
	var speed int
	var speedLimit float64
	var hybrideCar bool
	var model string

	fmt.Println(speed, speedLimit, hybrideCar)


	// print the value of empty string
	fmt.Println(model) // ""
	fmt.Printf("%T\n", model) // string
	fmt.Printf("%q\n", model) // """"

}