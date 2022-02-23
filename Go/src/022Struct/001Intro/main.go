package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	//First way
	//p1 := person{"James", "Bond"}
	//p2 := person{"Miss", "moneyPenny"}

	//Second way
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "moneyPenny",
		age:   27,
	}

	fmt.Println("First Person is ", p1)
	fmt.Println("Second Person is ", p2)

	//get type
	fmt.Printf("%T\n", p1)

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p1.age)

}
