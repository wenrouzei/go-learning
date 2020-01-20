package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			fmt.Println("放入", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("拿出i:", i, <-c)
		}
		quit <- 0
	}()
	fmt.Println("start")
	fibonacci(c, quit)
	fmt.Println("next")
}
