package main

import "fmt"

func main() {
	type a struct {
		A int
		B []int
	}
	b := a{A: 1, B: []int{1, 2, 3}}
	c := b.A
	d := b.B
	fmt.Printf("%p %v %d %d", b.B, b.B, cap(b.B), len(b.B))
	fmt.Println("\n##################")
	fmt.Printf("%p %d %p %d", &c, c, &b.A, b.A)
	fmt.Println("\n##################")
	fmt.Printf("%p %v %d %d", d, d, cap(d), len(d))
	b.A = 66
	b.B[0] = 4
	fmt.Println("\n##################")
	fmt.Println(b)
	fmt.Println("\n##################")
	fmt.Printf("%p %v %d %d", b.B, b.B, cap(b.B), len(b.B))
	fmt.Println("\n##################")
	fmt.Printf("%p %d %p %d", &c, c, &b.A, b.A)
	fmt.Println("\n##################")
	fmt.Printf("%p %v %d %d", d, d, cap(d), len(d))
}
