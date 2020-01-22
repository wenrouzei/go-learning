package main

import (
	"fmt"
	"strconv"
)

func queryUserById(id int) chan string {
	c := make(chan string)
	go func() {
		c <- "姓名" + strconv.Itoa(id)
	}()
	return c
}

func main() {
	//三个协程同时并发查询，缩小执行时间，
	//本来一次查询需要1秒，顺序执行就得3秒，
	//现在并发执行总共1秒就执行完成
	name1 := queryUserById(1)
	name2 := queryUserById(2)
	name3 := queryUserById(3)
	//从通道中获取执行结果
	fmt.Println(<-name2)
	fmt.Println(<-name1)
	fmt.Println(<-name3)
}
