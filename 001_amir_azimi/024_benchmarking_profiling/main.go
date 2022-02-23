package main

/*
 *	This is something that showing over app speed
 *	=> almost it is for monitoring and analyzing our Go Code
 */

func Abs(val float64) float64 {
	if val < 0.0 {
		return -val
	}
	return val
}

func Sqrt(val float64) (float64, error) {
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
