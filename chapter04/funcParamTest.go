package main

//import "fmt"

type Point struct{ X, Y int }

/*
func main() {
	p := Point{1, 2}
	//fmt.Println(Scale(p, 5))
	//fmt.Println(Scale_p(p, 5))
	//Scale_p(p, 5)

	Scale_p1(p, 5)
	fmt.Println(p)
	//fmt.Println(Scale_p(p, 3))
	p1 := Scale_p(p, 3)
	fmt.Printf("p1==p --> %v\n", &p1 == &p)
	fmt.Println(p)
	Scale_p2(&p, 5)
	fmt.Println(p)
}
*/
func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

//结构体类型的值可以作为参数传递给函数或者作为函数的返回值。
func Scale_p(p Point, factor int) Point {
	p.X *= factor
	p.Y *= factor
	return p
}

func Scale_p1(p Point, factor int) {
	p.X *= factor
	p.Y *= factor
}

//出于效率的考虑，大型的结构体通常都使用结构体指针的方式直接传递给函数或者从函数中返回
//这种方式在函数需要修改结构体内容的时候也是必需的，在 GO 这种按值调用的语言中，调用的函数接收到的是实参的一个副本，并不是实参的引用。
func Scale_p2(p *Point, factor int) {
	p.X *= factor
	p.Y *= factor
}
