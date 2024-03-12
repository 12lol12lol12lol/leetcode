package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorderList(t *testing.T) {
	testData := map[string]struct {
		head     []int
		expected []int
	}{
		"Case 1": {
			head:     []int{1, 2, 3, 4},
			expected: []int{1, 4, 2, 3},
		},
		"Case 2": {
			head:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 5, 2, 4, 3},
		},
	}
	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			head := sliceToLinkedList(data.head)
			reorderList(head)
			assert.Equal(t, data.expected, linkedListToSlice(head))
		})
	}
}
