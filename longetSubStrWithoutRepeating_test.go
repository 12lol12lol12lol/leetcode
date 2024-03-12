package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tableData := map[struct {
		s        string
		expected int
	}]struct{}{
		{
			s:        "abcabcbb",
			expected: 3,
		}: {},
		{
			s:        "bbbbb",
			expected: 1,
		}: {},
		{
			s:        "pwwkew",
			expected: 3,
		}: {},
		{
			s:        " ",
			expected: 1,
		}: {},
	}
	for data := range tableData {
		assert.Equal(t, data.expected, lengthOfLongestSubstring(data.s), fmt.Sprintf("error with s=%v", data.s))
	}
}
