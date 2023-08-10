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
 * Given the root of a binary tree, return the inorder traversal of its nodes' values.
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/?envType=study-plan-v2&envId=programming-skills
 */
func iterativeInorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	curr := root
	for curr != nil || len(stack) != 0 {
		for curr != nil { // Left
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, curr.Val) // Root
		curr = curr.Right           // Right
	}
	return ans
}
