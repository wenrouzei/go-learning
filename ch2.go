package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case t := <-time.After(5 * time.Second):
				fmt.Println("timeout")
				fmt.Println(t)
				o <- true
				break
			}
		}
	}()
	go func() {
		fmt.Println("3333")
	}()
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	c <- 10
	fmt.Println(<-o)
	fmt.Println("end")
}
