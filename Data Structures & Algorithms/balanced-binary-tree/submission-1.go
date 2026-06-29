/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
func isBalanced(root *TreeNode) bool {
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        l, r := dfs(node.Left), dfs(node.Right)

        if l == -1 || r == -1 || abs(l-r) > 1 {
            return -1 
        }

        return 1 + max(l, r)
    }

    return dfs(root) != -1
}