package main

import (
	"fmt"
)

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
	dylan.updatePassByValueFirstName("Dyl")
	dylan.print()
	// Notice name isn't updated due to pass by value

	maguina := person{
		firstName: "Maguina",
		lastName: "Ramilo Henry",
		contactInfo: contactInfo{
			email: "maguina@cool.com",
			zip: 7441,
		},
	}
	maguina.print()

	maguina.updateFirstName("Maz")
	maguina.print()

}

func (p person) updatePassByValueFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	pointerToPerson.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}