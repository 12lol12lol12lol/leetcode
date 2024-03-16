package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}

	queue := []*TreeNode{root}
	var currentValue *TreeNode
	for len(queue) > 0 {
		currentValue, queue = queue[0], queue[1:]
		if compareNodes(currentValue, subRoot) {
			if isSameTree(currentValue, subRoot) {
				return true
			}
		}
		if currentValue.Right != nil {
			queue = append(queue, currentValue.Right)
		}
		if currentValue.Left != nil {
			queue = append(queue, currentValue.Left)
		}
	}
	return false
}
