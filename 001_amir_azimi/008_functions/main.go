package main

import "fmt"

/*
	Base structure :
	func functionName(firstParameter parameter_type) return_type
*/

func add(a int, b int) int {
	return a + b
}

//you can don't use type hinting for firstly parameters
func mines(a, b int) int {
	return a - b
}

func multiple(a int, b int) int {
	return a * b
}

func division(a int, b int) float64 {
	if b != 0 {
		return float64(a / b)
	} else {
		return 0
	}
}

func swapCallByValue(x int, y int) (int, int) {
	x += y
	y = x - y
	x = x - y
	return x, y
}

//Call bt reference
func swap(x *int, y *int) {
	*x += *y
	*y = *x - *y
	*x = *x - *y
	fmt.Printf("values in function X : %v,  Y : %v\n", *x, *y)
}

func main() {
	a := 5
	b := 3

	fmt.Println("ADD is 	::\t", add(a, b))
	fmt.Println("Mines is	::\t", mines(a, b))
	fmt.Println("Multiple	::\t", multiple(a, b))
	fmt.Println("Division	::\t", division(a, b))

	//Call by value
	x, y := swapCallByValue(a, b)
	fmt.Printf("values in main X : %v,  Y : %v\t but function returns is %v, %v\n", a, b, x, y)

	//Call by reference
	swap(&a, &b)
	fmt.Printf("values in main X : %v,  Y : %v\n", a, b)
}
