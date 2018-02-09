package main

import(
	"fmt"
)

func main(){
	p := new(int)
	q := new(int)

	fmt.Println(p==q)
	fmt.Println(*p==*q)
}