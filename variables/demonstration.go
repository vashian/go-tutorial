package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("2nd argument:", os.Args[2])
	fmt.Println("3rd argument:", os.Args[3])

	fmt.Println("Items inside os.Args:", len(os.Args))


	var (
		l = len(os.Args) - 1
		n1 = os.Args[1]
		n2 = os.Args[2]
		n3 = os.Args[3]
	)

	fmt.Println("There are", l, "people !")
	fmt.Println("Hello great", n1, "!")
	fmt.Println("Hello great", n2, "!")
	fmt.Println("Hello great", n3, "!")
	fmt.Println("Nice to meet you all.")
}