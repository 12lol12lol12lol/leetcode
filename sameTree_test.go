package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	testData := map[string]struct {
		p        []int
		q        []int
		expected bool
	}{
		"Case 1": {
			p:        []int{1, 2, 3},
			q:        []int{1, 2, 3},
			expected: true,
		},
		"Case 2": {
			p:        []int{1, 2},
			q:        []int{1, -1, 2},
			expected: false,
		},
		"Case 3": {
			p:        []int{1, 2, 1},
			q:        []int{1, 1, 2},
			expected: false,
		},
	}
	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, isSameTree(slice2Tree(data.p), slice2Tree(data.q)))
		})
	}
}
