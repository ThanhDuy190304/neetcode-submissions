/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }

    var isSameTree func(n1 *TreeNode, n2 *TreeNode) bool
    isSameTree = func(n1 *TreeNode, n2 *TreeNode) bool {
        if n1 == nil && n2 == nil {
            return true
        }
        if n1 == nil || n2 == nil {
            return false
        }
        if n1.Val != n2.Val {
            return false
        }
        return isSameTree(n1.Left, n2.Left) && isSameTree(n1.Right, n2.Right)
    }

    if isSameTree(root, subRoot) {
        return true
    }
    
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}