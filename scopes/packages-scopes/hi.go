package main

import "fmt"

func hey() {
	fmt.Println("Hey!")
}

// func bye() {
// 	fmt.Println("Bye!")
// }

// ***** EXPLANATION *****
//
// ERROR: "bye" function "redeclared"
//        in "this block"
//
// "this block" means = "main package"
//
// "redeclared" means = you're using the same name
//   in the same scope again
//
// main package's scope is:
// all source-code files which are in the same main package