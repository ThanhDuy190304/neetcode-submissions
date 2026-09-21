/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil{
		return nil
	}
	visited := make(map[int]struct{})
	created := make(map[int]*Node)

	array := make(map[int][]int)

	var dfs func(n *Node)
	dfs = func(n *Node) {
		if n == nil {
			return
		}
		if _, ok := visited[n.Val]; ok {
			return
		}

		visited[n.Val] = struct{}{}
		var tmp []int

		for _, node := range n.Neighbors {
			tmp = append(tmp, node.Val)
			dfs(node)
		}

		array[n.Val] = tmp
	}

	dfs(node)

	for i, row := range array{
		var node *Node
		if created[i] == nil{
			node = &Node{
				Val: i,
			}
			created[i] = node
		}else{
			node = created[i]
		}

		for _, val := range row{
			var neighbor *Node 
			if created[val] == nil{
				neighbor = &Node{
					Val: val,
				}
				created[val] = neighbor
			}else{
				neighbor = created[val]
			}
			node.Neighbors = append(node.Neighbors, neighbor)
		}
	} 

	return created[node.Val]

}
