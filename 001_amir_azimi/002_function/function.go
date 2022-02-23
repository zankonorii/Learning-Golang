package main

import "fmt"

func main() {
	//var x int
	//var y int

	x := 1.0
	y := 2.0

	fmt.Printf("X : %v type of %T\n", x, x)
	fmt.Printf("Y : %v type of %T\n", y, y)

	var mean float64 // float32 give error because x and y are float64
	mean = (x + y) / 2

	fmt.Printf("Mean : %v type of %T\n", mean, mean)
}
