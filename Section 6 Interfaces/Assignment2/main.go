package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]

	byteArr, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	text := string(byteArr)

	fmt.Println(text)
}