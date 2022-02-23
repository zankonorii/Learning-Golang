package main

import "fmt"

//OOP

type Trade struct {
	/*
		if property name start with :
			Capital letter A,B,C,... 	=>	public
			small letter	a,b,c,...	=>	private
	*/
	Symbol string  //Stock symbol
	Volume int     // number of shares
	Price  float64 //Trade Price
	Buy    bool    //true if buy trade, false if sell trade
}

func main() {
	t1 := Trade{"Apple", 10, 99.98, true}

	fmt.Println(t1)
	fmt.Printf("%+v\n", t1)
	fmt.Printf("type of T1 object is %T\n", t1)

	fmt.Println("Access to each of properties : ")
	fmt.Println("Symbol ::\t", t1.Symbol)
	fmt.Println("Volume ::\t", t1.Volume)
	fmt.Println("Price ::\t", t1.Price)
	fmt.Println("Buy ::\t", t1.Buy)

	t2 := Trade{
		Price:  89.89,
		Symbol: "MicroSoft",
		Buy:    false,
		Volume: 5, //Trailing comma is mandatory
	}

	fmt.Println(t2)
	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Println(t3)
	fmt.Printf("%+v\n", t3)
}
