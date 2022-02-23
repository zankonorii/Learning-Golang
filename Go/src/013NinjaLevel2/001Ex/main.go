package main

import "fmt"

/*
 *	#1.	Write a program that prints a number in decimal, binary and hex.
 */
func main() {
	x := 25

	fmt.Printf("Decimal is %d, Binary is %b, Hex is %x\n", x, x, x)
	//OR
	fmt.Printf("Decimal is %d, Binary is %b, Hex is %#x\n", x, x, x)
}
