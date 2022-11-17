package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main () {
	dylan := person{firstName: "Dylan", lastName: "Graham"}

	fmt.Println(dylan.lastName)
	dylan.lastName = "New Last Name"
	fmt.Println(dylan.lastName)

	var zeroValuedPerson person
	fmt.Println(zeroValuedPerson)

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%+v", alex)

}