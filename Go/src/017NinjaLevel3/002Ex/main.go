package main

import "fmt"

/*
 *	2. Print every rune code point of the uppercase alphabet three times.
 *		Your output should look like this:
 *			1
 *				U+0041'A'
 *				U+0041'A'
 *				U+0041'A'
*			2
 *				U+0041'B'
 *				U+0041'B'
 *				U+0041'B'
 *			...through the rest of the alphabet characters
*/

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i - 67)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
