package main

import "fmt"

// Slices act differently to: Struct and primitive data types: string, int, float, bool, array

// Value Types (primitives) use pointers: struts, string, int, float, bool, array.
// Reference types don't use pointers: slices, maps, channels, pointers, functions.

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	printSlice(mySlice)

}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func printSlice(s []string) {
	fmt.Println(s)
}