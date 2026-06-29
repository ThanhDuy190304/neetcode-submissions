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
	if root == nil{
		return true
	}

	result := true
	var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int{
		if node == nil{
			return 0
		}
		l, r := dfs(node.Left), dfs(node.Right)
		if abs(l - r) > 1{
			result = false
			return 0
		}
		return 1 + max(l, r)
	}

	dfs(root)
	return result
}
