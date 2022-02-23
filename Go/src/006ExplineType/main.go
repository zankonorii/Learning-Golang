package main

import "fmt"

var x = 123

func main() {
	//x := 1
	var x = 5
	fmt.Println(x)

	foo()

	t := `this is a text`

	fmt.Println(t)

	//zero value
	var z int
	var q bool
	var c string
	var a float32

	fmt.Println("z is : ", z)
	fmt.Printf("%T\n", z)

	fmt.Println("q is : ", q)
	fmt.Printf("%T\n", q)

	fmt.Println("c is : ",c)
	fmt.Printf("%T\n", c)
	fmt.Println("a is : ",a)
	fmt.Printf("%T\n", a)
}

func foo() {
	fmt.Println("Foo : ", x)
	fmt.Printf("%T\n", x)
}
