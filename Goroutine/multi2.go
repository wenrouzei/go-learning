package main

import (
	"fmt"
)

func main() {
	c, quit := make(chan int), make(chan int)
	go func() {
		c <- 2    // 添加数据
		quit <- 1 // 发送完成信号
	}()
	for isQuit := false; !isQuit; {
		// 监视信道c的数据流出
		select {
		case v := <-c:
			fmt.Printf("received %d from c \r\n", v)
		case q := <-quit:
			isQuit = true
			fmt.Println("退出", q)
			// quit信道有输出，关闭for循环
		}
	}
}
