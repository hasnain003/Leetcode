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
 * Given a binary tree, find its minimum depth.
 * The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
 * https://leetcode.com/problems/minimum-depth-of-binary-tree/description/?envType=study-plan-v2&envId=programming-skills
*/

class Solution {
    public int minDepth(TreeNode root) {
        if(root == null) {
            return 0;
        }
        Queue<TreeNode> q = new LinkedList<>();
        q.add(root);
        int depth = 0;
        while (!q.isEmpty()) {
            int n = q.size(); // Size of current level
            depth++;
            for (int i = 0; i < n;i++) {
                TreeNode node = q.remove();
                if(node.left == null && node.right == null) {
                    return depth;
                }
                if(node.left != null) {
                    q.add(node.left);
                }
                if(node.right != null ) {
                    q.add(node.right);
                }
            }
        }
        return depth;
    }
}