/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
/*
* Given the root of a binary tree, invert the tree, and return its root.
*https://leetcode.com/problems/invert-binary-tree/description/?envType=study-plan-v2&envId=programming-skills
 */
class Solution {
    public TreeNode invertTree(TreeNode root) {
        if (root == null || (root.left == null && root.right == null)) {
            return root;
        }
        TreeNode left = invertTree(root.left);
        TreeNode right =  invertTree(root.right);
        root.left = right;
        root.right = left;
        return root;
    }
}