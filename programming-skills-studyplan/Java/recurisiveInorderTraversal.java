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
  * Given the root of a binary tree, return the inorder traversal of its nodes' values.
  * https://leetcode.com/problems/binary-tree-inorder-traversal/description/?envType=study-plan-v2&envId=programming-skills
  */
class Solution {
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> ans = new ArrayList<>();
        helper(root,ans);
        return ans;
    }
    public void helper(TreeNode root, List<Integer> ans) {
        if(root == null) {
            return;
        }
        helper(root.left,ans);
        ans.add(root.val);
        helper(root.right,ans);
    }
}