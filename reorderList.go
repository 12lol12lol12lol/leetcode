package main

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	ar := []*ListNode{head}
	next := head.Next
	for next != nil {
		ar = append(ar, next)
		next = next.Next
	}
	var left, right int
	for left, right = 0, len(ar)-1; (right - left) >= 0; left, right = left+1, right-1 {
		ar[left].Next = ar[right]
		ar[right].Next = ar[left+1]
	}
	ar[len(ar)-left].Next = nil
}
