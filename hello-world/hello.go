package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello returns a greeting
func Hello() string {
	return "Hello, World"
}

// HelloWithName returns a personalised greeting
func HelloWithName(name string) string {
	return "Hello, " + name
}

// HelloWithPrefix returns greeting with fixed prefix
func HelloWithPrefix(name string) string {
	return englishHelloPrefix + name
}

// HelloWithCase returns a personalised greeting, defaulting to Hello,
// world if an empty name is passed
func HelloWithCase(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println("Getting started.")
}
