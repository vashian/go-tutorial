package main

import "fmt"

func main() {
	fmt.Println("c:\\test\\one\\two\\three")
	fmt.Println(`c:\test\one\two\three`)

	html := `
<html>
<p>hi</p>
</html>
	`
	fmt.Println(html)

	name, family := "John", "Doe"

	fmt.Println(name + " " + family)
	fmt.Println(len(name))

	fmt.Println(len("very") + len(`\"cool\"`))
}