package main

import "fmt"

func main() {
	var (
		name string 
		age int
	)

	name = "Monica"
	age = 28


	var firstName, lastName string = "John", "Doe"

	fmt.Println(name, age)
	fmt.Println(firstName, lastName)

	// short declaration
	safe := true
	a, b := 1, 2

	fmt.Println(safe,a,b)
}