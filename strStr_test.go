package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	testData := map[string]struct {
		haystack string
		needle   string
		expected int
	}{
		"Case 1": {
			haystack: "sadbutsad",
			needle:   "sad",
			expected: 0,
		},
		"Case 2": {
			haystack: "leetcode",
			needle:   "leto",
			expected: -1,
		},
	}
	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, strStr(data.haystack, data.needle))
		})
	}
}
