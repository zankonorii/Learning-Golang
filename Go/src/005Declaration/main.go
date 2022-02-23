package main

import "fmt"

func main() {
	fmt.Println(y)

	x := 100+24*36
	fmt.Println(x)
	x = 2
	fmt.Println(x)

	//fmt.Println( i := 20) error
	//var y = 13	error
	//x := 15		error
	fmt.Println(y)
	//y = "this is a text"	error
	//x = "this is a text"	error

	test()
}

var y = 25


func test()  {
	//fmt.Println(x)	error
	fmt.Println(y)
}
