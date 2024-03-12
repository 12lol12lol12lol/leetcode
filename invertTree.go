package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func slice2Tree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	head := &TreeNode{
		Val:   arr[0],
		Left:  nil,
		Right: nil,
	}
	if len(arr) <= 1 {
		return head
	}
	deque := []*TreeNode{head}
	i := 1
	var current *TreeNode
	for i < len(arr) {
		current, deque = deque[0], deque[1:]
		if i < len(arr) {
			current.Left = &TreeNode{Val: arr[i]}
			deque = append(deque, current.Left)
			i++
		}
		if i < len(arr) {
			current.Right = &TreeNode{Val: arr[i]}
			deque = append(deque, current.Right)
			i++
		}

	}
	return head
}

func tree2Slice(head *TreeNode) []int {
	res := []int{}
	if head == nil {
		return res
	}
	deque := []*TreeNode{head}
	var current *TreeNode
	for len(deque) > 0 {
		current, deque = deque[0], deque[1:]
		res = append(res, current.Val)
		if current.Left != nil {
			deque = append(deque, current.Left)
		}
		if current.Right != nil {
			deque = append(deque, current.Right)
		}

	}
	return res
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := []*TreeNode{root}
	var current *TreeNode
	for len(stack) > 0 {
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		current.Left, current.Right = current.Right, current.Left
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
	}
	return root
}
