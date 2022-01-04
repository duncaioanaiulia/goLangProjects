package hello

import "fmt"

const enHelloPrefix = "Hello, "
const spHelloPrefix = "Hola, "
const frHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frHelloPrefix
	case spanish:
		prefix = spHelloPrefix
	default:
		prefix = enHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
