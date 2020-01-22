package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	//带有缓冲区的通道
	cInt := make(chan int, 10)
	go func() {
		//product  ，循环往通道中写入一个元素
		for i := 0; i < 100; i++ {
			cInt <- i
		}
		//关闭通道
		close(cInt)
	}()
	go func() {
		defer wg.Done()
		//consumer   遍历通道消费元素并打印
		for temp := range cInt {
			fmt.Println(temp)
			//len函数可以查看当前通道元素个数
			fmt.Println("当前通道元素个数", len(cInt))
		}
	}()
	wg.Wait()
}
