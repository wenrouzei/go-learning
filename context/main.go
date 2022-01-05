package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println(fmt.Sprintf("context.TODO %p %p", context.TODO(), context.TODO()))
	fmt.Println(fmt.Sprintf("context.Background %p %p", context.Background(), context.Background()))
	ctx, cancel := context.WithCancel(context.TODO())
	ctx3, _ := context.WithCancel(context.TODO())
	fmt.Println(fmt.Sprintf("context.WithCancel(context.TODO()) %p %p", ctx, ctx3))
	ctx2 := context.WithValue(ctx, "abc", "cccc")
	fmt.Println(fmt.Sprintf("%p %p %v", ctx, ctx2, ctx2.Value("abc")))
	out := make(chan struct{})
	go func(ctx context.Context) {
		ctx1, cancel1 := context.WithTimeout(ctx, time.Second*6)
		fmt.Println(fmt.Sprintf("%p %p", ctx1.Done(), ctx.Done()), 11111111111)
		for {
			select {
			case v := <-ctx1.Done():
				fmt.Println("ctx1 done", v, ctx1.Err())
				out <- struct{}{}
				return
			default:
				fmt.Println(fmt.Sprintf("ctx1 echo %p %p", ctx, ctx1))
				time.Sleep(2 * time.Second)
			}
		}
		cancel1()
	}(ctx)
	time.Sleep(6 * time.Second)
	cancel()
	<-out
}
