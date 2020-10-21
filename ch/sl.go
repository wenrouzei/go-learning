package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	fmt.Println(fmt.Sprintf("%p %p %v %v", a, b, a, b))
	fmt.Println(cap(a), len(a))
	a = append(a, 4) //第一次扩容,底层数组不可改变，指向新的（1024以下2倍，以上1.25倍）容量数组，a指针已改变
	fmt.Println(cap(a), len(a))
	fmt.Println(fmt.Sprintf("%p %p %v %v", a, b, a, b))
	a = append(a, 5)
	fmt.Println(fmt.Sprintf("%p %p %v %v", a, b, a, b))
	a = append(a, 6)
	fmt.Println(cap(a), len(a))
	fmt.Println(fmt.Sprintf("%p %p %v %v", a, b, a, b))
	a = append(a, 7) //第二次扩容,底层数组不可改变，指向新的（1024以下2倍，以上1.25倍）容量数组，a指针再次改变
	fmt.Println(cap(a), len(a))
	fmt.Println(fmt.Sprintf("%p %p %v %v", a, b, a, b))
	a = append(a, 8)
	fmt.Println(fmt.Sprintf("%p %p %v %v", a, b, a, b))
	a = append(a, 9)
	fmt.Println(fmt.Sprintf("%p %p %v %v", a, b, a, b))
	a = append(a, 10)
	fmt.Println(fmt.Sprintf("%p %p %v %v", a, b, a, b))
}
