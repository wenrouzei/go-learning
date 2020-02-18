package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	seconds := map[string]int{"a": 1, "b": 2}
	seconds["c"] = 111
	for i, s := range seconds {
		// 计数加 1
		wg.Add(2)
		go func(i string, s int) {
			// 计数减 1
			defer wg.Done()
			fmt.Printf("goroutine %v 结束\n", i)
		}(i, s)
		go func(i string, s int) {
			// 计数减 1
			defer wg.Done()
			fmt.Printf("goroutine %v 结束\n", i)
		}(i, s)
	}

	// 等待执行结束
	wg.Wait()
	fmt.Println("所有 goroutine 执行结束")
}
