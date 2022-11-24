package main

import "fmt"

func main()  {
	// var colors map[string]string


	// map[key]value
	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#4bf745",
	// 	"blue": "#8as134",
	// }

	colors := make(map[string]string)
	colors["white"] = "#ffffff"

	delete(colors,"white")

	fmt.Println(colors)
}