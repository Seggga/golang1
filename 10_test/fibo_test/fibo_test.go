package fibo_test

import (
	"testing"

	"github.com/seggga/golang1/05_recursion/fibo"
)

var tests = []struct {
	name  string
	input uint
	want  uint
}{
	{name: "table-test one", input: 1, want: 1},
	{name: "table-test two", input: 2, want: 1},
	{name: "table-test 30", input: 31, want: 832040},
}

func TestFindNumRecursive(t *testing.T) {

	for _, tc := range tests {
		if got := fibo.FindNumRecursive(tc.input); got != tc.want {
			t.Errorf("%s: recursive expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func TestFindNumLinear(t *testing.T) {

	for _, tc := range tests {
		if got := fibo.FindNumLinear(tc.input); got != tc.want {
			t.Errorf("%s: linear expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func TestFindNumRecursiveMap(t *testing.T) {

	for _, tc := range tests {

		optiMap := make(map[uint]uint)
		optiMap[0] = 0
		optiMap[1] = 1

		if got := fibo.FindNumRecursiveMap(tc.input, optiMap); got != tc.want {
			t.Errorf("%s: with map expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
