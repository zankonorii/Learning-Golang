package main

import "fmt"

/*
 *	2.	Take the code from the previous exercise, then store the value of type person
 *	in map with the key of last name. Access each value in the map. Print out the values
 *	ringing over the slice.
 */

type person struct {
	first       string
	last        string
	favFavorite []string
}

func main() {
	p1 := person{
		first:       "James",
		last:        "Bond",
		favFavorite: []string{"one", "two", "three"},
	}

	p2 := person{
		first:       "Miss",
		last:        "moneyPenny",
		favFavorite: []string{"strawberry", "vanilla", "capuccino"},
	}
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)
	}
}
