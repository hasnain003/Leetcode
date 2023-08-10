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
 Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
https://leetcode.com/problems/maximum-depth-of-binary-tree/description/?envType=study-plan-v2&envId=programming-skills
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type NodeDepth struct {
	node  *TreeNode
	depth int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []NodeDepth{{root, 1}}
	maxDepth := 0
	for len(stack) != 0 {
		depthNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node := depthNode.node
		depth := depthNode.depth
		if node.Left != nil {
			stack = append(stack, NodeDepth{node.Left, depth + 1})
		}
		if node.Right != nil {
			stack = append(stack, NodeDepth{node.Right, depth + 1})
		}
		maxDepth = max(maxDepth, depth)
	}
	return maxDepth
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
