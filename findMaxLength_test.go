package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxLength(t *testing.T) {
	testData := map[string]struct {
		nums     []int
		expected int
	}{
		"Case 1": {
			nums:     []int{0, 1},
			expected: 2,
		},
		"Case 2": {
			nums:     []int{0, 1, 0},
			expected: 2,
		},
		"Case 3": {
			nums:     []int{0, 0, 1, 0, 0, 0, 1, 1},
			expected: 6,
		},
	}
	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, findMaxLength(data.nums))
		})
	}
}
