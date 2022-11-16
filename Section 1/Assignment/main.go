package main

import "fmt"

func main() {
	slice := []int{0,1,2,3,4,5,6,7,8,9,10}
	for i := range slice {
		if i % 2 == 0 {
			fmt.Println(i, " is even")
		}
		if i % 2 == 1 {
			fmt.Println(i, " is odd")
		}
	}
}