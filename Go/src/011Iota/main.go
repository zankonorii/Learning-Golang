package main

import "fmt"

const (
	x = iota
	//y = iota
	y
	//z = iota
	z
)

const (
	a = iota
	b
	c
	d
)

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	fmt.Println(c)
	fmt.Printf("%T\n", c)

	fmt.Println(d)
	fmt.Printf("%T\n", d)
}
