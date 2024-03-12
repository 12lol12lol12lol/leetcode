package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	testData := map[*struct {
		head     []int
		expected []int
	}]struct{}{
		{
			head:     []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		}: {},
		{
			head:     []int{1, 2},
			expected: []int{2, 1},
		}: {},
		{
			head:     []int{},
			expected: []int{},
		}: {},
	}
	for data := range testData {
		assert.Equal(t, data.expected, linkedListToSlice(reverseList(sliceToLinkedList(data.head))), fmt.Sprintf("error with args head=%v", data.head))
	}

}
