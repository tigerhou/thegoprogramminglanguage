package main

import "time"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
		close(naturals)

	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		println(x)
	}
}
