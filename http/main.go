package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""), reflect.TypeOf(v).Kind())
	}
	//ctx, cancel := context.WithTimeout(r.Context(),time.Second*10)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	go func(ctx2 context.Context) {
		for {
			select {
			case <-ctx2.Done():
				fmt.Println(fmt.Sprintf("ctx2 done %v %p", ctx2.Err(), ctx2))
				return
			default:
				fmt.Println(fmt.Sprintf("ctx run %p", ctx2))
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	fmt.Fprintf(w, "Hello astaxie!") // 这个写入到 w 的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName)       // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
