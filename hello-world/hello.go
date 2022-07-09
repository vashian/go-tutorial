package main

import "fmt"

const french = "French"
const spanish = "Spanish"

const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "

func main() {

	fmt.Println(Hello("Monica", "english"))

}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	
	switch language {

		case french:
			prefix = frenchHelloPrefix
		
		case spanish:
			prefix = spanishHelloPrefix
		
		default:
			prefix = englishHelloPrefix

	}
	return
}


// func Hello(name string, language string) string {

// 	if name == "" {
// 		name = "World"
// 	}

// 	if language == spanish {
// 		return fmt.Sprintf("%s%s!", spanishHelloPrefix, name)
// 	}

// 	if language == french {
// 		return fmt.Sprintf("%s%s!", frenchHelloPrefix, name)
// 	}

// 	return fmt.Sprintf("%s%s!", englishHelloPrefix, name)
// }