package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotInteger(t *testing.T) {
	testData := map[string]struct {
		n        int
		expected int
	}{
		"Case 1": {
			n:        8,
			expected: 6,
		},
		"Case 2": {
			n:        1,
			expected: 1,
		},
		"Case 3": {
			n:        4,
			expected: -1,
		},
	}
	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, pivotInteger(data.n))
		})
	}
}
