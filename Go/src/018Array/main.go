package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	fmt.Println(len(x))

	x[3] = 42
	fmt.Println(x)

	fmt.Println(x[4])
}
