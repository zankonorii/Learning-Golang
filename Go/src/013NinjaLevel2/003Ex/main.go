package main

import "fmt"

/*
 * #3.	create Typed and UTyped constants. Print The values of Constants.
 */
const (
	a     = 42 //un type
	b int = 43 //type = int
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
}
