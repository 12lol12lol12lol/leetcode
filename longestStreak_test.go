package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	type testData struct {
		nums     []int
		expected int
	}
	tableData := map[*testData]struct{}{
		{
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		}: {},
		{
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		}: {},
	}
	for data := range tableData {
		assert.Equal(t, data.expected, longestConsecutive(data.nums), fmt.Sprintf("error with args nums=%v", data.nums))
	}

}
