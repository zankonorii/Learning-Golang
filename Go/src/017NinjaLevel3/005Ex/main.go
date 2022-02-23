package main

import "fmt"

/*
 *	5. Print out the remainder (modulus) which is found for each
 * 		number between 10 and 100 when is is divided by 4.
 */

func main() {
	for x := 10; x <= 100; x++ {
		fmt.Printf("When %v divided by 4, the remainder (aka modulus) is %v\n", x, x%4)
	}
}
