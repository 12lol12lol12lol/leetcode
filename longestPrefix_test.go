package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPrefix(t *testing.T) {
	testData := map[string]struct {
		s        string
		expected string
	}{
		"Case 1": {
			s:        "level",
			expected: "l",
		},
		"Case 2": {
			s:        "ababab",
			expected: "abab",
		},
	}
	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, longestPrefix(data.s))
		})
	}
}
