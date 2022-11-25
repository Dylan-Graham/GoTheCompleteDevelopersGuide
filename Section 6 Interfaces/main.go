package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func main () {
	ebot := englishBot{}
	printGreeting(ebot)

	sbot := spanishBot{}
	printGreeting(sbot);

}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// Custom logic
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	// Custom logic
	return "Hola!"
}