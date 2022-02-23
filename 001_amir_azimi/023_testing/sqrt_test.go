package main

import (
	"fmt"
	"testing"
)

func almostEqual(a, b float64) bool {
	return Abs(a-b) < 0.001
}

func TestSimple(t *testing.T) {
	val, err := sqrt(2)

	if err != nil {
		t.Fatalf("error in calculation: %s\n", err) //fatal => error , f => format
	}

	//	if this is not True , test passed automatically
	if !almostEqual(val, 1.414214) {
		t.Fatalf("bad value - %f\n", val)
	}

}

type testCase struct {
	value, expected float64
}

func TestMany(t *testing.T) {
	testcase := []testCase{
		{0.0, 0.0},
		{2, 1.414214},
		{9.0, 3.0},
	}

	for _, tc := range testcase {
		//t.Run(first argument, callback function => second argument)
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
			out, err := sqrt(tc.value)
			if err != nil {
				t.Fatalf("error")
			}
			if !almostEqual(out, tc.expected) {
				t.Fatalf("%f != %f", out, tc.expected)
			}
		})
	}
}
