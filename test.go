package main

import "fmt"

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

	a := contain{Saiyan{"abc", 24323, nil}, "fasdfas"}
	a.Super()
	fmt.Println(a.Name, a.Saiyan)
}
