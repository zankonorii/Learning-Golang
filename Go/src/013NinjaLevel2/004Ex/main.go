package main

import "fmt"

/*
 *	#4.	Write a Program that
 *		1. assigns an into a variable
 *		2. prints that int in decimal, binary and hex
 *		3. shifts the bits of that into over 1 position to the left, and assigns that to variable
 *		4. prints that int in decimal, binary and hex
 */

func main() {
	x := 42
	fmt.Printf("Decimal %d\tBinary %b\tHex %#x\n", x, x, x)
	x = x << 1
	fmt.Printf("Decimal %d\tBinary %b\tHex %#x\n", x, x, x)
}
