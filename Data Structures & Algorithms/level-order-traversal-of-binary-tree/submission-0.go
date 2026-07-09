/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var roots []*TreeNode
	var result [][]int
	if root == nil {
		return [][]int{}
	}

	roots = append(roots, root)
	result = append(result, []int{root.Val})
	
	var bfs func(roots []*TreeNode)

	bfs = func(roots []*TreeNode) {
		if len(roots) != 0 {

			var tmpVals []int
			var tmpNodes []*TreeNode

			for i := range roots {
				if roots[i].Left != nil {
					tmpVals = append(tmpVals, roots[i].Left.Val)
					tmpNodes = append(tmpNodes, roots[i].Left)
				}

				if roots[i].Right != nil {
					tmpVals = append(tmpVals, roots[i].Right.Val)
					tmpNodes = append(tmpNodes, roots[i].Right)
				}
			}
			if len(tmpVals) != 0{
				result = append(result, tmpVals)
			}

			bfs(tmpNodes)
		}
	}

	bfs(roots)

	return result
}
