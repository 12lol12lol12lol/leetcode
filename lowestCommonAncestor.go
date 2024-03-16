package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	currentNode := root

	for {
		if currentNode.Val > p.Val && currentNode.Val > q.Val {
			currentNode = currentNode.Left
		} else if currentNode.Val < p.Val && currentNode.Val < q.Val {
			currentNode = currentNode.Right
		} else {
			break
		}
	}
	return currentNode
}
