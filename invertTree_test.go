package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertTree(t *testing.T) {
	testData := map[string]struct {
		root     []int
		expected []int
	}{
		"Case 1": {
			root:     []int{4, 2, 7, 1, 3, 6, 9},
			expected: []int{4, 7, 2, 9, 6, 3, 1},
		},
		"Case 2": {
			root:     []int{2, 1, 3},
			expected: []int{2, 3, 1},
		},
		"Case 3": {
			root:     []int{},
			expected: []int{},
		},
	}
	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, tree2Slice(invertTree(slice2Tree(data.root))))
		})
	}
}
