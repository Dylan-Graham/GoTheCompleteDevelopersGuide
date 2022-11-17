package main

import "fmt"

type person struct {
	firstName string
	lastName string
	id identification
}

type identification struct {
	idNumber int
	passportNumber string
}

func main () {
	var dylan person
	dylan.firstName = "Dylan"
	dylan.lastName = "Graham"
	dylan.id.idNumber = 9607881
	dylan.id.passportNumber = "A330120"
	fmt.Printf("%+v\n", dylan)

	maguina := person{
		firstName: "Maguina",
		lastName: "Ramilo Henry",
		id: identification{
			idNumber: 123123,
			passportNumber: "A33",
		},
	}
	fmt.Printf("%+v",maguina)

}