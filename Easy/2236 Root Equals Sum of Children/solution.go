// https://leetcode.com/problems/root-equals-sum-of-children/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkTree(root *TreeNode) bool {
    if root.Val == root.Left.Val + root.Right.Val{
        return true
    }else{
        return false
    }
}