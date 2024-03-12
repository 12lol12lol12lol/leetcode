package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	peekMinNode := func() *ListNode {
		if list1 == nil && list2 == nil {
			return nil
		}
		if list1 == nil {
			tmp := list2
			list2 = list2.Next
			return tmp
		}
		if list2 == nil {
			tmp := list1
			list1 = list1.Next
			return tmp
		}
		if list1.Val < list2.Val {
			tmp := list1
			list1 = list1.Next
			return tmp
		}
		tmp := list2
		list2 = list2.Next
		return tmp
	}
	head := &ListNode{
		Val:  peekMinNode().Val,
		Next: nil,
	}
	current := head
	for list1 != nil || list2 != nil {
		current.Next = peekMinNode()
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}
	return head
}
