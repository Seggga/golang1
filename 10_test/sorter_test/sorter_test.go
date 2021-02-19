package sorter_test

import (
	"reflect"
	"testing"

	"github.com/seggga/golang1/04_sort/sorter"
)

var tests = []struct {
	name  string
	input []int32
	want  []int32
}{
	{name: "negative", input: []int32{-1, -10, -100, -1000}, want: []int32{-1000, -100, -10, -1}},
	{name: "zeroes", input: []int32{0, 0, 0, 0, 0}, want: []int32{0, 0, 0, 0, 0}},
	{name: "different sign", input: []int32{15, 0, 5, -10, -15, -5, 10}, want: []int32{-15, -10, -5, 0, 5, 10, 15}},
}

func TestInsSort(t *testing.T) {
	for _, tc := range tests {
		got := sorter.InsSort(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("InsSort %s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func TestBubbleSort(t *testing.T) {
	for _, tc := range tests {
		got := sorter.BubbleSort(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("BubleSort %s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
