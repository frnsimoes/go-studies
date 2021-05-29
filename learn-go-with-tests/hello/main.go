package main

// helps with performance
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "

func greetingPrefix(language string) (prefix string) {

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "Portuguese":
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + CapitalizedName(name)
}

func main() {

}
