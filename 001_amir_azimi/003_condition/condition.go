package main

import "fmt"

func main() {
	x := 6.0
	y := 11.0

	if x > 10 {
		fmt.Println("X is great than 10")
	} else {
		fmt.Println("X is smaller than 10")
	}

	if av := x / y; av > 0.5 {
		fmt.Println("x is greater than half of y")
	} else {
		fmt.Println("x is smaller than half of y")
	}
}
