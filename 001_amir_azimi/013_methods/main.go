package main

import "fmt"

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	buy    bool
}

// Value name started with capital letter because we want that be a public method
//func resiver
//func () ==> point
func (trade *Trade) Value() float64 {
	value := float64(trade.Volume) * trade.Price

	if trade.buy {
		value = -value
	}
	return value
}

func main() {
	t1 := Trade{"APPLE", 10, 99.98, true}

	fmt.Println(t1.Value())
}
