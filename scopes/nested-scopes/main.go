package main

import "fmt"

var ok bool = true

func nested() {
	var ok bool = false

	fmt.Println("you nested var:", ok)
}

func main() {
	fmt.Println("you global var:", ok)

	nested()

	fmt.Println("you global var:", ok)

}