package main

import "fmt"

/*
 * #6. Using iota, create 4 constants for the last 4 years. Print the constants values.
 */

const (
	a = 2018 + iota
	b = 2018 + iota
	c = 2018 + iota
	d = 2018 + iota
)

func main() {
	fmt.Println(a, b, c, d)
}
