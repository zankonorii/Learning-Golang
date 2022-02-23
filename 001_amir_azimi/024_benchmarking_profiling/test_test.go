package main

import (
	"fmt"
	"testing"
)

func almostEqual(a, b float64) bool {
	return Abs(a-b) < 0.001
}

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)

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
			out, err := Sqrt(tc.value)
			if err != nil {
				t.Fatalf("error")
			}
			if !almostEqual(out, tc.expected) {
				t.Fatalf("%f != %f", out, tc.expected)
			}
		})
	}
}

//Benchmarking
// *testing.B => B is Benchmark
func BenchmarkSqrt(b *testing.B) {

	//b.N : N = Number
	for i := 0; i < b.N; i++ {

		//parsing to float => float64(number)
		_, err := Sqrt(float64(i))

		if err != nil {
			b.Fatal(err)
		}
	}
}

/*
 *	after that for running benchmark test , run following command:
 *	go test -v -bench dir_path(.)
 *
 *	For running just benchmark tests:
 *	go test -v -bench dir_path -run A : A is name of a none test.
 *
 */

//Profiling
/*
 *	go test -v -bench dir_path -run A -cpuprofile=prof.out
 *	this is for profiling cpu .
 *	After running above command, a file created with name :
 *	prof.out that we cannot open that.
 *
 *	go tool : with this command we can use Go tools.
 *
 *	go tool pprof : this command give us ability to read
 *	created file.
 *
 *	go tool pprof file_name
 *
 *	After that an Interactive env run for us
 *
 *	with list function name we can see all details about functions
 *	like : list Sqrt
 *
 *	Profiling = how many speed of each line of code.
 *
 *
 */
