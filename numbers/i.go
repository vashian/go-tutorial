package main

import "fmt"

func main () {
	fmt.Println(-3.0 * 2) // float value?

	fmt.Println(4 / 2) // two int result is int

	// reminder operator
	// only use with integer
	fmt.Println(5 % 2)

	fmt.Println(-(-3))
	fmt.Println(- -3)

	var (
		scoreOne = 10
		scoreTwo = 20
		average float64
	)

	average = float64(scoreOne + scoreTwo) / 2
	fmt.Println("Average:", average)


	ratio := 1.0 / 10.0

	fmt.Println("Ratio:", ratio) // 0.1

	for range [...]int{10:0} {
		ratio += 1.0 / 10.0
	}

	fmt.Println(ratio)


	x := 3 / 2
	
	fmt.Println(x)
}