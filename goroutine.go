package main

import (
	"fmt"
	"runtime"
)

func say(s string, n int) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s, n)
	}
}

func main() {
	go say("world", 1) // 开一个新的 Goroutines 执行
	say("hello", 2)    // 当前 Goroutines 执行
}
