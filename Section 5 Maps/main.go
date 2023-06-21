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

	colors2 := map[int]string{
		1: "Yeet",
	}

	printColours(colors2)
}

func printMap(c map[string]string) {
	fmt.Println(c)
	for key, value := range c {
		fmt.Printf("Brand:%s Model:%s\n", key, value)
	}
}

func printColours(c map[int]string) {
	fmt.Println(c)
	for key, value := range c {
		fmt.Printf("Key:%d Color:%s\n", key, value)
	}
}