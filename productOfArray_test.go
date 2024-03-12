package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	type testData struct {
		nums     []int
		expected []int
	}

	tableData := map[*testData]struct{}{
		{
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		}: {},
		{
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		}: {},
	}
	for v := range tableData {
		assert.Equal(t, v.expected, productExceptSelf(v.nums), fmt.Sprintf("error with args nums=%v", v.nums))
	}

}
