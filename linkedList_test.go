package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceToLinkedList(t *testing.T) {
	testData := map[*struct {
		nums     []int
		expected *ListNode
	}]struct{}{
		{
			nums: []int{1, 2, 3, 4},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		}: {},
	}
	for data := range testData {
		assert.Equal(t, data.expected, sliceToLinkedList(data.nums))
	}
}

func TestLinkedListToSlice(t *testing.T) {
	testData := map[*struct {
		firstNode *ListNode
		expected  []int
	}]struct{}{
		{
			expected: []int{1, 2, 3, 4},
			firstNode: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		}: {},
	}
	for data := range testData {
		assert.Equal(t, data.expected, linkedListToSlice(data.firstNode))
	}
}
