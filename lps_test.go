package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLps(t *testing.T) {
	testData := map[string]struct {
		s        string
		expected []int
	}{
		"Case 1": {
			s:        "",
			expected: []int{},
		},
		"Case 2": {
			s:        "a",
			expected: []int{0},
		},
		"Case 3": {
			s:        "abacaba",
			expected: []int{0, 0, 1, 0, 1, 2, 3},
		},
		"Case 4": {
			s:        "abacabab",
			expected: []int{0, 0, 1, 0, 1, 2, 3, 2},
		},
	}
	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, lps(data.s))
		})
	}
}
