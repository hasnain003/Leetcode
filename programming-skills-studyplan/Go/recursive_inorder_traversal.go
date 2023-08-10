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
func recursiveInorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	helper(root, &ans)
	return ans
}

func helper(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, ans)
	*ans = append(*ans, root.Val)
	helper(root.Right, ans)
}
