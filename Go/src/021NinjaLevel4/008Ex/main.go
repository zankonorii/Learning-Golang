package main

import "fmt"

/*
*	8.	Create a map with key of Type string which is person's last_first name, and a value of type
*	 []string which stores their favorite things. Store three records in your map. Print out all of
*	 the values. along with their index position in the slice, without using the range clause.
 */
func main() {
	m := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, "Martinis", "Women"},
		"moneyPenny_miss": []string{"James Bond", "Literature", `Computer Since`},
		"no_dr":           []string{"Being evil", `Ice Cream`, "Sunsets"},
	}

	//fmt.Println(m)
	for key, value := range m {
		fmt.Println("This is the record for ", key)
		for i, v := range value {
			fmt.Println("\t ", i, v)
		}
	}
}
