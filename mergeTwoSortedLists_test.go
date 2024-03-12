package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	testData := map[*struct {
		list1    []int
		list2    []int
		expected []int
	}]struct{}{
		{
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		}: {},
		{
			list1:    []int{},
			list2:    []int{},
			expected: []int{},
		}: {},
		{
			list1:    []int{},
			list2:    []int{0},
			expected: []int{0},
		}: {},
	}
	for data := range testData {
		assert.Equal(
			t, data.expected,
			linkedListToSlice(mergeTwoLists(sliceToLinkedList(data.list1), sliceToLinkedList(data.list2))),
			fmt.Sprintf("error with args list1=%v list2=%v", data.list1, data.list2),
		)
	}
}
