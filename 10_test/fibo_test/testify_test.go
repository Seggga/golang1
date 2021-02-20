package fibo_test

import (
	"testing"

	"github.com/seggga/golang1/05_recursion/fibo"
	"github.com/stretchr/testify/assert"
)

func TestRecursion(t *testing.T) {

	for _, tc := range tests {
		resp := fibo.FindNumRecursive(tc.input)
		assert.Equal(t, tc.want, resp)
	}
}
