package main

import "fmt"

/*
 *	 10. Write down what these print:
 *	 	ftm.Println(true && true)	True
 *	 	fmt.Println(true && false)	False
 *	 	fmt.Println(true || true)	True
 *	 	fmt.Println(true || false)	True
 *	 	fmt.Println(!true)			False
 */

func main() {
	fmt.Println("True && True  : ", true && true)
	fmt.Println("True && false : ", true && false)
	fmt.Println("True || True  : ", true || true)
	fmt.Println("True || false : ", true || false)
	fmt.Println("!True", !true)
}
