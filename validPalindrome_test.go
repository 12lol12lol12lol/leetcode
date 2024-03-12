package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	type testData struct {
		value    string
		expected bool
	}
	tableData := map[*testData]struct{}{
		{
			value:    "A man, a plan, a canal: Panama",
			expected: true,
		}: {},
		{
			value:    "race a car",
			expected: false,
		}: {},
		{
			value:    " ",
			expected: true,
		}: {},
	}
	for data := range tableData {
		assert.Equal(t, data.expected, isPalindrome(data.value), fmt.Sprintf("error with args value=%v", data.value))
	}

}
