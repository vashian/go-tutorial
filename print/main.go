package main

import (
	"fmt"
	"os"
)

func main() {
	name, family := os.Args[1], os.Args[2]

	fullname := "Your name is %s and your family is %s"

	fmt.Printf(fullname, name, family)
}