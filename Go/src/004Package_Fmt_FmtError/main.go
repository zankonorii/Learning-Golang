package main

import "fmt"

func main() {
	/*n , err := fmt.Println("hello every one, this is a text" , 4 , 5.2 , true)
	fmt.Println(n)
	fmt.Println(err)*/

	//_ allow to dont see it and compiler dont secret on it when dont use it.
	n, _ := fmt.Println("hello every one, this is a text", 4, 5.2, true)
	fmt.Println(n)
}
