package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubtree(t *testing.T) {
	testData := map[string]struct {
		root     []int
		subRoot  []int
		expected bool
	}{
		"Case 1": {
			root:     []int{3, 4, 5, 1, 2},
			subRoot:  []int{4, 1, 2},
			expected: true,
		},
		"Case 2": {
			root:     []int{3, 4, 5, 1, 2, -1, -1, -1, -1, 0},
			subRoot:  []int{4, 1, 2},
			expected: false,
		},
	}
	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, isSubtree(slice2Tree(data.root), slice2Tree(data.subRoot)))
		})
	}
}
