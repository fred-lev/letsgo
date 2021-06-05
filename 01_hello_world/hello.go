package main

import (
	"fmt"
)

const french = "French"
const spanish = "Spanish"
const russian = "Russian"

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const russianHelloPrefix = "Privet, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case russian:
		prefix = russianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Fred", ""))
}