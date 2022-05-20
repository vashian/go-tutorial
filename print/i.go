package main

import "fmt"


func main() {
	n1, n2, n3 := 1, 2, 3

	fmt.Println("Number one is:", n1,"- Number two is:", n2,"- Number three is:", n3)

	fmt.Printf("Number one is: %d - Number two is: %d - Number three is: %d\n", n1,n2,n3)

	var fullName string

  // prints the string in quoted-form like this ""
	fmt.Printf("%q\n", fullName)

	fullName = "John Doe"

	fmt.Printf("%q\n", fullName)


	fmt.Println("helloworld")
	fmt.Println("hello\nworld")
	fmt.Println("hello\\n\"world\"")

	var (
		age int
		weight float64
		stduent bool
		name string
	)

	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", weight)
	fmt.Printf("%T\n", stduent)
	fmt.Printf("%T\n", name)


	var (
		planet = "jupiter"
		distance = 778.5
		orbitalPeriod = 6.67
		hasLife = false
	)

	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v\n", distance)
	fmt.Printf("Orbital Period: %v\n", orbitalPeriod)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)

	// argument indexing - indexes start from 1
	fmt.Printf(
		"%v is %v away. Think! %[2]v kms! %[1]v OMG.\n",
		planet, distance,
	)
}