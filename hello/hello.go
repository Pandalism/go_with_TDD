package main

import "fmt"

const englishHelloPrefix = "Hello, "
const surpriseSufix = "!"

func Hello(name string) string {

	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name + surpriseSufix
}

func main() {
	fmt.Println(Hello("world"))
}
