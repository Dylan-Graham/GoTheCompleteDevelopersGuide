package main

import "fmt"

type shape interface {
	getArea() int
}

type triangle struct {
	height float64
	base float64
}

type square struct {
	sideLength int
}

func main() {
	squareShape:= square{
		sideLength: 5,
	}

	printArea(squareShape)

	triangleShape := triangle{
		base: 2,
		height: 5,
	}

	printArea(triangleShape)
}

func (t triangle) getArea() int {
	return int(0.5 * t.height * t.base)
}

func (s square) getArea() int {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	area := s.getArea()
	fmt.Println(area)
}