package main

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}

	}()

	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	for {
		println(<-squares)
	}
}
