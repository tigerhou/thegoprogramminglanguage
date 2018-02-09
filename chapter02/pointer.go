package main

import (
	"fmt"
)

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)
	var x1, y int
	fmt.Println(&x1 == &x1, &x1 == &y, &x1 == nil)
	var _ = f()
	fmt.Println(f() == f())
	v := 1
	inc(&v)
	fmt.Println(inc(&v))
}
func f() *int {
	v := 1
	return &v
}
func inc(p *int) int {
	*p++
	return *p
}
