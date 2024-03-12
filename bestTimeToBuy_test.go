package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	tableData := map[*struct {
		prices   []int
		expected int
	}]struct{}{
		{
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		}: {},
		{
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		}: {},
		{
			prices:   []int{1, 4, 2},
			expected: 3,
		}: {},
	}
	for data := range tableData {
		assert.Equal(t, data.expected, maxProfit(data.prices), fmt.Sprintf("error with prices=%v", data.prices))
	}
}
