package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	tableData := map[*struct {
		nums     []int
		expected [][]int
	}]struct{}{
		{
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		}: {},
		{
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		}: {},
		{
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		}: {},
		{
			nums:     []int{-2, 0, 0, 2, 2},
			expected: [][]int{{-2, 0, 2}},
		}: {},
	}
	for data := range tableData {
		assert.Equal(t, data.expected, threeSum(data.nums), fmt.Sprintf("threeSum error with args nums=%v", data.nums))
	}
	for data := range tableData {
		assert.Equal(t, data.expected, threeSumMap(data.nums), fmt.Sprintf("threeSumMap error with args nums=%v", data.nums))
	}
}
