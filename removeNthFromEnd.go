package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	first, second := dummy, dummy
	for i := 0; i <= n; i++ {
		second = second.Next
	}
	for second != nil {
		second = second.Next
		first = first.Next
	}
	first.Next = first.Next.Next
	return dummy.Next
}
