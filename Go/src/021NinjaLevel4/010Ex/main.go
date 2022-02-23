package main

import "fmt"

/*
 *	10.	 Using the code from the previous example, delete a record to your map.
 * 	Now Print the map out using the "range" loop
 */

func main() {
	//	I get this section from previous Example.
	m := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, "Martinis", "Women"},
		"moneyPenny_miss": []string{"James Bond", "Literature", `Computer Since`},
		"no_dr":           []string{"Being evil", `Ice Cream`, "Sunsets"},
	}

	//Add to map
	m["fleming_ian"] = []string{"streaks", `cigars`, `espionage`}
	//all up codes given from Ex number nine.

	delete(m, "no_dr")
	for key, value := range m {
		fmt.Println("This is the record for ", key)
		for i, v := range value {
			fmt.Println("\t", i, v)
		}
	}

}
