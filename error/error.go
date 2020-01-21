package main

import (
	"errors"
	"log"
	"net/http"
)

type appError struct {
	Error   error
	Message string
	Code    int
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		http.Error(w, e.Message, e.Code)
	}
}

func viewRecord(w http.ResponseWriter, r *http.Request) *appError {
	return &appError{errors.New("test"), "Can't display record", 500}
}

func main() {
	http.Handle("/", appHandler(viewRecord)) // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
