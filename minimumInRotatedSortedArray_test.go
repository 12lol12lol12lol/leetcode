package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	testData := map[*struct {
		nums     []int
		expected int
	}]struct{}{
		{
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		}: {},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		}: {},
		{
			nums:     []int{11, 13, 15, 17},
			expected: 11,
		}: {},
		{
			nums:     []int{3, 1, 2},
			expected: 1,
		}: {},
	}
	for data := range testData {
		assert.Equal(t, data.expected, findMin(data.nums), fmt.Sprintf("error with args nums=%v", data.nums))
	}
}
