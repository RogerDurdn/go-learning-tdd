package main

const (
	spanish          = "spanish"
	french           = "french"
	englishHelloPref = "Hello "
	spanishHelloPref = "Hola "
	frenchHelloPref  = "Bonjour "
)

func Greeting(to string, language string) string {
	if to == "" {
		to = "World"
	}
	prefix := englishHelloPref
	switch language {
	case spanish:
		prefix = spanishHelloPref
	case french:
		prefix = frenchHelloPref
	default:
	}
	return prefix + to
}
