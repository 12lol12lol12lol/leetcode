package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterReplacement(t *testing.T) {
	tableData := map[struct {
		s        string
		k        int
		expected int
	}]struct{}{
		{
			s:        "ABAB",
			k:        2,
			expected: 4,
		}: {},
		{
			s:        "AABABBA",
			k:        1,
			expected: 4,
		}: {},
		{
			s:        "ABAA",
			k:        0,
			expected: 2,
		}: {},
		{
			s:        "ABBB",
			k:        2,
			expected: 4,
		}: {},
	}
	for data := range tableData {
		assert.Equal(t, data.expected, characterReplacement(data.s, data.k), fmt.Sprintf("error with s=%v k=%v", data.s, data.k))
	}
}
