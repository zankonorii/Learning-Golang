/*
 *	Defer is used to ensure that a function call is performed later in a program’s execution,
 *	usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.
 */
package main

import "fmt"

func safeValues(vals []int, index int) int {
	//Go anonymous function func() {}()
	// defer => clean resource
	defer func() { //try
		if err := recover(); err != nil { // recover => catch
			fmt.Printf("ERROR: %s\n", err)
		}
	}()

	return vals[index]
}

func main() {
	//Panic
	/*vals := []int{1, 2, 3}

	//panic = throw
	v := vals[10] // when key not exists return a panic error
	//output =>  panic: runtime error: index out of range [10] with length 3

	fmt.Println(v)*/

	/*file, err := os.Open("no-file")

	//make panic  , using panic is discouraged
	if err != nil {
		//return nil, err
		panic(err) // output => panic: open no-file: no such file or directory
	}
	defer file.Close()
	defer file.Close()
	fmt.Println("File opened")*/

	v := safeValues([]int{1, 2, 3}, 5) //	ERROR: runtime error: index out of range [5] with length 3
	fmt.Println(v)
}
