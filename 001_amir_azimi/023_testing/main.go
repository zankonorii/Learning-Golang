package main

/*
 *
 *		go help test => for more information
 *
 *		Testes in go should follow this regex => name_test.go
 *		with above regex go can find and  run testes.
 *		usually test icon are difference from main source
 *
 *		in this session we want to test Abs and Sqrt function
 *
 *		Create a file with `sqrt_test.go` name.
 *		Test functions start with Test key word.
 *		In testes we just should find errors, it passes automatically.
 *		For stat testing , just enough type this command : go test
 *
 *		Test case:
 *
 *		go test -v : this command return more information about testing
 *		progressive like: time, status and etc.
 *
 *
 */

func Abs(val float64) float64 {
	if val < 0.0 {
		return -val
	}
	return val
}

func sqrt(val float64) (float64, error) {
	if val <= 0.0 {
		return 0.0, nil
	}

	guess, epsilon := 1.0, 0.00001

	for i := 0; i < 10000; i++ {
		if Abs(guess*guess-val) < epsilon {
			return guess, nil
		}

		guess = (val/guess + guess) / 2.0
	}

	return guess, nil
}
