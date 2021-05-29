package main

import (
	"fmt"
	"strings"
)

// helps with performance
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + strings.Title(name)
}

func main() {
	fmt.Println(Hello("Fernando"))
}
