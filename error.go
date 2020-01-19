package main

import (
	"errors"
	"fmt"
)

type Error struct {
	msg string
}

func (err *Error) Error() string {
	return err.msg
}

func main() {
	a := Error{msg: "fasdfsa"}
	err := errors.New("fsadfasdfas")
	fmt.Println(a.Error())
	fmt.Println(err)
}
