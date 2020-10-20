package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	fmt.Println(fmt.Sprintf("%p %p %v %v", a, b, a, b))
	a = append(a, 4) //a指针已改变
	a = append(a, 5)
	a = append(a, 6)
	fmt.Println(fmt.Sprintf("%p %p %v %v", a, b, a, b))
}
