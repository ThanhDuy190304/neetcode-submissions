/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
		return nil
	}
	created := make(map[int]*Node)
	var dfs func(node *Node) *Node

	dfs = func(node *Node) *Node{
		if node == nil{
			return nil
		}
		if clone, ok := created[node.Val]; ok{
			return clone
		}
		clone := &Node{
			Val: node.Val,
		}
		created[node.Val] = clone
		for _, neighbor := range node.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(neighbor))
		}
		return clone
	}

	return dfs(node)
}
