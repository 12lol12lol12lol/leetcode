package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	testData := map[*struct {
		nums     []int
		target   int
		expected int
	}]struct{}{
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		}: {},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		}: {},
		{
			nums:     []int{1},
			target:   0,
			expected: -1,
		}: {},
		{
			nums:     []int{1, 3},
			target:   3,
			expected: 1,
		}: {},
	}
	for data := range testData {
		assert.Equal(t, data.expected, search(data.nums, data.target), fmt.Sprintf("error with args nums=%v target=%v", data.nums, data.target))
	}
}
