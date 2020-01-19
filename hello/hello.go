package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const surpriseSufix = "!"
const defaultUser = "World"
const langES = "Spanish"
const langFR = "French"

func Hello(name string, language string) string {

	if name == "" {
		name = defaultUser
	}

	return greetingPrefix(language) + name + surpriseSufix
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case langES:
		prefix = spanishHelloPrefix
	case langFR:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("", ""))
}
