package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	testData := map[string]struct {
		n        int
		expected []string
	}{
		"Case 1": {
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		// "Case 2": {
		// 	n:        1,
		// 	expected: []string{"()"},
		// },
	}
	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, generateParenthesis(data.n))
		})
	}
}
