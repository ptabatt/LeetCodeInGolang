package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TwoSum(t *testing.T) {
	tests := []struct {
		inputNums   []int
		inputTarget int
		want        []int
	}{
		{
			inputNums:   []int{2, 7, 11, 15},
			inputTarget: 9,
			want:        []int{0, 1},
		},
		{
			inputNums:   []int{3, 2, 4},
			inputTarget: 6,
			want:        []int{1, 2},
		},
		{
			inputNums:   []int{3, 3},
			inputTarget: 6,
			want:        []int{0, 1},
		},
		{
			inputNums:   []int{2, 5, 5, 11},
			inputTarget: 10,
			want:        []int{1, 2},
		},
	}

	for _, tc := range tests {
		got := twoSum(tc.inputNums, tc.inputTarget)
		assert.Equal(t, tc.want, got)
	}
}
