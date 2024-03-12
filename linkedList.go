package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToLinkedList(nums []int) *ListNode {
	if len(nums) < 1 {
		return nil
	}
	firstNote := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	if len(nums) < 2 {
		return firstNote
	}
	currentNote := firstNote
	for i := 1; i < len(nums); i++ {
		currentNote.Next = &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		currentNote = currentNote.Next
	}
	return firstNote
}

func linkedListToSlice(head *ListNode) []int {
	res := []int{}
	current := head
	for current != nil {
		res = append(res, current.Val)
		current = current.Next
	}
	return res
}

func printLinkesList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%v ", current.Val)
		current = current.Next
	}
	fmt.Println()
}
