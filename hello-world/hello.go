package main

import (
	"fmt"
)

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	switch language {
	case "Spanish":
		return spanishHelloPrefix + name
	case "French":
		return frenchHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(Hello("world", "English"))
}
