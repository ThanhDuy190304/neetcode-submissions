/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	
	var dfs func(n1 *TreeNode, n2 *TreeNode) bool

	dfs = func(n1 *TreeNode, n2 *TreeNode) bool{
		if (n1 != nil && n2 == nil){
			return false
		}else if (n1 == nil && n2 != nil){
			return false
		}else if (n1 == nil && n2 == nil){
			return true
		}
	
		if n1.Val != n2.Val{
			return false
		}

		return dfs(n1.Left, n2.Left) && dfs(n1.Right, n2.Right)
	}
	result := dfs(p, q)
	return result
}
