package main

import "fmt"

/*
 *	6.	Create a slice to store the name of all the state in the United States of America.
 *		What is the length of your slice? What is the Capacity? Print out all the values,
 *		alone with their index position in the slice, without using  the range clause .
 *		Here is a list of the states:
 *		`Alabama`, `Alaska`, `Arizona`,	`Arkansas`,	`California`, `Colorado`, ``
 */

func main() {
	x := make([]string, 50, 50)
	x = []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Washington", `Colorado`}
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		fmt.Println(i)
		fmt.Println(x[i])
	}

	for i, v := range x {
		fmt.Println("Index is ", i)
		fmt.Println("Value is ", v)
	}

}
