package main

import (
	"fmt"
	"os"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

//In Go, we don't have __constructor__ method, but we have something like that
//Note :: Error is optional
func newTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol cannot be empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("volume must be grater than zero (was %d)", volume)
	}

	if price <= 0.0 {
		return nil, fmt.Errorf("price must be grater than 0.0 was(%d)", price)
	}

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}
	return trade, nil
}

func newTradeWithoutErrorAndWithoutValidation(symbol string, volume int, price float64, buy bool) *Trade {

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}
	return trade
}

func newTradeWithoutError(symbol string, volume int, price float64, buy bool) *Trade {
	if symbol == "" {
		return nil
	}

	if volume <= 0 {
		return nil
	}

	if price <= 0.0 {
		return nil
	}

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}
	return trade
}

func main() {
	t1, err1 := newTrade("Microsoft", 10, 100.00, true)

	if err1 != nil {
		fmt.Printf("ERROR : can't create trade - %s\n", err1)
		os.Exit(1) // 1 => programmer error
	} else {
		fmt.Println(t1)
	}
}
