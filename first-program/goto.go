package main

import "fmt"

func main() {
	myFunc()
}

func myFunc() {
	i := 0
Here:   // label ends with ":"
	fmt.Println(i)
	i++

	if i == 1 {
		fmt.Println("I'm in the if statement")
		return
	}
	goto Here   // jump to label "Here"
}