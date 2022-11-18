package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contactInfo
}

type contactInfo struct {
	email string
	zip int
}

func main () {
	var dylan person
	dylan.firstName = "Dylan"
	dylan.lastName = "Graham"
	dylan.updateFirstName("Dyl")
	dylan.print()
	// Notice name isn't updated due to reference

	maguina := person{
		firstName: "Maguina",
		lastName: "Ramilo Henry",
		contactInfo: contactInfo{
			email: "maguina@cool.com",
			zip: 7441,
		},
	}
	maguina.print()

}

func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)

}