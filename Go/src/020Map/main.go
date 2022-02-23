package main

import "fmt"

func main() {
	/*
	 *	 Map
	 *	map[Key_Type]value_type{values}
	 */
	m := map[string]int{
		"James":            32,
		"Miss Money Penny": 27,
	}
	fmt.Println(m) //	map[James:32 Miss Money Penny:27]

	fmt.Println(m["James"])            //	32
	fmt.Println(m["Miss Money Penny"]) //	27
	fmt.Println(m["key_not_exist"])    //	0
	fmt.Println(m)                     //	map[James:32 Miss Money Penny:27]

	//Check value in map
	value, ok := m["value_looking"]
	fmt.Println(value) //	0
	fmt.Println(ok)    //	false

	value2, ok2 := m["James"]
	fmt.Println(value2) //	32
	fmt.Println(ok2)    //	true

	//Ex:
	if v, ok := m["key"]; ok {
		fmt.Println("This is The If Print", v)
	}

	if v, ok := m["James"]; ok {
		fmt.Println("This is The If Print", v)
	}

	//Add to Map
	m["todd"] = 33

	for key, value := range m {
		fmt.Println("Key is  : ", key)
		fmt.Println("value is : ", value)
	}

	xi := []int{4, 5, 8, 9, 1, 2, 35}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	/*
	 *	How to Delete & Entry map
	 *		delete(<map name> , "key")
	 */

	//Delete an exists key
	delete(m, "James")
	fmt.Println(m)

	//Delete a none exists key
	delete(m, "Ian Fleming")
	fmt.Println(m)

}
