package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
* Given the root of a binary tree, invert the tree, and return its root.
*https://leetcode.com/problems/invert-binary-tree/description/?envType=study-plan-v2&envId=programming-skills
 */
func invertTreeIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			top := queue[0]
			queue = queue[1:]
			swapChild(top)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
	}
	return root
}

func swapChild(root *TreeNode) {
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
}
