package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func (s *Saiyan) Super() {
	s.Power += 10000
}

func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
		Father: &Saiyan{
			Name:   "Goku",
			Power:  9001,
			Father: nil,
		},
	}
}

type contain struct {
	Saiyan
	test string
}

type aaa interface {
	echo()
}

type aa string

func (a aa) echo() {
	fmt.Println(a)
}

func main() {
	//goku := &Saiyan{"Goku", 9001, &Saiyan{"111", 232, nil}}
	//goku.Super()
	//fmt.Println(goku.Father) // 将会打印出 19001
	//a := NewSaiyan("abc", 1000)
	//b := NewSaiyan("abcd", 2000)
	//fmt.Println(a.Name)
	//fmt.Println(b.Name)
	c := 123
	fmt.Println(uint64(c))

	d := &c

	a := contain{Saiyan{"abc", 24323, nil}, "fasdfas"}
	a.Super()
	fmt.Println(a.Name, a.Saiyan)

	defer func() {
		fmt.Println("abc")
	}()

	fmt.Println("11111")

	fmt.Println(reflect.TypeOf(c), reflect.TypeOf(d))

	fmt.Println(unsafe.Sizeof(c), unsafe.Sizeof(d))

	var first aaa
	var second aa = "abc"
	first = second
	first.echo()
}
