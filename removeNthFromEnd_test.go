package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	testData := map[string]struct {
		head     []int
		n        int
		expected []int
	}{
		"Case 1": {
			head:     []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{1, 2, 3, 5},
		},
		"Case 2": {
			head:     []int{1},
			n:        1,
			expected: []int{},
		},
		"Case 3": {
			head:     []int{1, 2},
			n:        1,
			expected: []int{1},
		},
		"Case 4": {
			head:     []int{1, 2},
			n:        2,
			expected: []int{2},
		},
	}
	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, linkedListToSlice(removeNthFromEnd(sliceToLinkedList(data.head), data.n)))
		})
	}
}
