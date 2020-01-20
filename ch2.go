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
				println(v)
			case t := <-time.After(5 * time.Second):
				println("timeout")
				fmt.Println(t)
				o <- true
				break
			}
		}
	}()
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	fmt.Println(<-o)
	fmt.Println("end")
}
