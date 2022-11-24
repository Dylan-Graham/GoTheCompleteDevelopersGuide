package main

import "fmt"

func main()  {
	// var colors map[string]string

	colors := make(map[string]string)
	colors["white"] = "#ffffff"

	delete(colors,"white")

	fmt.Println(colors)

	// Interating
	// map[key]value
	kite := map[string]string{
		"naish": "pivot",
		"duotone": "rebel",
		"airush": "lift",
	}

	printMap(kite)
}

func printMap(c map[string]string) {
	fmt.Println(c)
	for key, value := range c {
		fmt.Printf("Brand:%s Model:%s\n", key, value)
	}
}