package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
type afrikaansBot struct {
	customGreeting string
}
type spanishBot struct {}

func main () {
	ebot := englishBot{}
	printGreeting(ebot)

	sbot := spanishBot{}
	printGreeting(sbot);

	abot := afrikaansBot{
		customGreeting: "Aangename kennis!",
	}
	printGreeting(abot);

}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// Custom logic
	return "Hello!"
}

func (sb spanishBot) getGreeting() string {
	// Custom logic
	return "Hola!"
}

func (ab afrikaansBot) getGreeting() string {
	// Custom logic
	return ab.customGreeting
}