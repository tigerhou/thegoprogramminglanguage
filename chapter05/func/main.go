package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	t := x - y
	first(t, 2)
	return
}

func first(x int, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}

//func Sin(value float64) float64

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)
	fmt.Printf("sub-> %d\n", sub(3, 5))
	fmt.Printf("first-> %d\n", first(3, 5))
	fmt.Printf("zero-> %d\n", zero(3, 7))
}
