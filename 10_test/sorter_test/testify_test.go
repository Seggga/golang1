package sorter_test

import (
	"testing"

	"github.com/seggga/golang1/04_sort/sorter"
	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {

	for _, tc := range tests {
		resp := sorter.InsSort(tc.input)
		assert.Equal(t, tc.want, resp)
	}
}
