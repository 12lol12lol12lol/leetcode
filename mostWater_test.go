package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	tableData := map[*struct {
		height   []int
		expected int
	}]struct{}{
		{
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		}: {},
		{
			height:   []int{1, 1},
			expected: 1,
		}: {},
	}
	for data := range tableData {
		assert.Equal(t, data.expected, maxArea(data.height), fmt.Sprintf("error with args height=%v", data.height))
	}
}
