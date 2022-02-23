package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	//we can define first, last and age again and can use both (in person and secretAgent)
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{first: "James", last: "Bond", age: 32},
		ltk:    true,
	}

	fmt.Println(sa1)
	fmt.Println(sa1.ltk)
	fmt.Println(sa1.first)
	fmt.Println(sa1.last)
	fmt.Println(sa1.age)
	//or
	fmt.Println(sa1.person.age)
}
