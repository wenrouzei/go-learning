package main

import (
	"fmt"
	"sync"
)

type ss struct {
	a int
}

var a chan *ss
var bb chan []int
var wg sync.WaitGroup

func main() {
	a = make(chan *ss, 10)
	bb = make(chan []int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		var e, f bool
		for !e && !f {
			select {
			case i, ok := <-a:
				if ok {
					fmt.Println(fmt.Sprintf("%p %v", i, i))
				} else {
					e = true
				}
			case i, ok := <-bb:
				if ok {
					fmt.Println(i)
				} else {
					f = true
				}
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			b := ss{}
			var c []int
			b.a = i
			a <- &b
			c = append(c, i)
			bb <- c
		}
		close(a)
		close(bb)
	}()
	wg.Wait()
}
