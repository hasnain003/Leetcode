package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTreeRecursive(root *TreeNode) *TreeNode {
	if root == nil || (root.Right == nil && root.Left == nil) {
		return root
	}
	right := root.Right
	root.Right = invertTreeRecursive(root.Left)
	root.Left = invertTreeRecursive(right)
	return root
}
