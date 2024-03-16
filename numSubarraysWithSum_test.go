package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSubarraysWithSum(t *testing.T) {
	testData := map[string]struct {
		nums     []int
		goal     int
		expected int
	}{
		"Case 1": {
			nums:     []int{1, 0, 1, 0, 1},
			goal:     2,
			expected: 4,
		},
		"Case 2": {
			nums:     []int{0, 0, 0, 0, 0},
			goal:     0,
			expected: 15,
		},
	}
	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, numSubarraysWithSum(data.nums, data.goal))
		})
	}
}
