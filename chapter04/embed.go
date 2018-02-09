package main

import "fmt"

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var wl Wheel
	wl.X = 8
	wl.Y = 8
	wl.Radius = 3
	wl.Spokes = 50
	fmt.Println(wl)
	fmt.Printf("%#v\n", wl)
	w := Wheel{Circle{Point{8, 8}, 3}, 50}
	fmt.Printf("%#v\n", w)
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 3,
		},
		Spokes: 50,
	}
}
