package main

import "fmt"

const englishHelloPrefix = "Hello "

// Hello returns a greeting
func Hello() string {
	return "Hello World"
}

// HelloWithName returns a personalised greeting
func HelloWithName(name string) string {
	return "Hello " + name
}

// HelloWithPrefix returns greeting with fixed prefix
func HelloWithPrefix(name string) string {
	return englishHelloPrefix + name
}

// HelloWithCase returns a personalised greeting, defaulting to Hello,
// world if an empty name is passed
func HelloWithCase(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println("Getting started.")
}
