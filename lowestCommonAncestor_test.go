package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowestCommonAncestor(t *testing.T) {
	testData := map[string]struct {
		root     []int
		p        int
		q        int
		expected int
	}{
		"Case 1": {
			root:     []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5},
			p:        2,
			q:        8,
			expected: 6,
		},
		"Case 2": {
			root:     []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5},
			p:        2,
			q:        4,
			expected: 2,
		},
		"Case 3": {
			root:     []int{2, 1},
			p:        2,
			q:        1,
			expected: 2,
		},
	}
	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			p := lowestCommonAncestor(slice2Tree(data.root), &TreeNode{Val: data.p}, &TreeNode{Val: data.q})
			assert.Equal(t, data.expected, p.Val)
		})
	}
}
