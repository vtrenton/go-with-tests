package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	switch language {
	case spanish:
		return spanishHelloPrefix + name
	case french:
		return frenchHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(Hello("world", "English"))
}
