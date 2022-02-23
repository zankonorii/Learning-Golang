package main

import "fmt"

//way 1
/*const x = 15
const y = 15.2
const z = "James Bond"*/

//way 2
/*const (
	x = 15
	y = 15.2
	z = "James Bond"
)*/

//way 3
const (
	x int     = 15
	y float64 = 15.2
	z string  = "James Bond"
)

func main() {
	fmt.Println("Value Of X is ", x)
	fmt.Printf("%T\n", x)

	fmt.Println("Value Of Y is ", y)
	fmt.Printf("%T\n", y)

	fmt.Println("Value Of Z is ", z)
	fmt.Printf("%T\n", z)
	/*
		x = 42	// Tack Error
		fmt.Println("Value Of X is ", x)
		fmt.Printf("%T\n", x)*/
}
