package main

import "fmt"

func main() {
	s := "Hello, playground"

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)

	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	/*
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		fmt.Println(i)	//tack an  error because i is not defined

	*/
	for i, v := range s {
		//fmt.Println(i, v)
		fmt.Printf("at index position %d we have hex %#x \n", i, v)
	}
}
