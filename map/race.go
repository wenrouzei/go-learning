package main

import "sync"

// go run -race .\race.go TODO -race 帮助检测是否出现数据竞争
func main() {
	m := make(map[string]int, 1)
	m[`foo`] = 1

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			m[`foo`]++
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			m[`foo`]++
		}
	}()
	wg.Wait()
}
