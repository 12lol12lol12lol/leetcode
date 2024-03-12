package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	type testData struct {
		nums     []int
		target   int
		expected []int
	}
	tableData := map[*testData]struct{}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		}: {},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		}: {},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		}: {},
	}
	for v := range tableData {
		assert.Equal(t, v.expected, twoSum(v.nums, v.target), fmt.Sprintf("error with args nums=%v target=%v", v.nums, v.target))
	}
}
