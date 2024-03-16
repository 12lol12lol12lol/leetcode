package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	queue := [][2]*TreeNode{{p, q}}
	for len(queue) > 0 {
		first, second := queue[0][0], queue[0][1]
		queue = queue[1:]
		if !compareNodes(first, second) {
			return false
		}
		if first.Left != nil || second.Left != nil {
			queue = append(queue, [2]*TreeNode{first.Left, second.Left})
		}
		if first.Right != nil || second.Right != nil {
			queue = append(queue, [2]*TreeNode{first.Right, second.Right})
		}

	}
	return true
}

func compareNodes(first, second *TreeNode) bool {
	if first == nil || second == nil {
		return first == second
	}
	return first.Val == second.Val
}
