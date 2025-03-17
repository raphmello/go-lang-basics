package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Printf("b.getGreeting(): %v\n", b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
