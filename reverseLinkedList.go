package main

/*
- Definition for singly-linked list.
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var previuous *ListNode
	current := head
	for {
		next := current.Next
		current.Next = previuous
		previuous = current
		if next == nil {
			break
		}
		current = next
	}
	return current
}
