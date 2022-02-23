package main

import "fmt"

//Slices and maps pass by reference
func doubleAt(values []int, index int) {
	values[index] *= 2
	fmt.Println("Double values is :\t", values[index])
}

func main() {
	values := []int{1, 2, 3, 4, 5}

	fmt.Println("values before function is :\t", values[2])
	doubleAt(values, 2)
	fmt.Println("values after function is :\t", values[2])
}
