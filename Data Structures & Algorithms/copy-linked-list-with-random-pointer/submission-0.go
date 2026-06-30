/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	nodeMap := make(map[*Node]*Node) 
	cur := head
	dummy := &Node{}
	copyCur := dummy
	nodeMap[head] = dummy
	for cur != nil {
		if cur.Random != nil {
			if copyRandom, exists := nodeMap[cur.Random]; !exists {
				copyCur.Random = &Node{}
				nodeMap[cur.Random] = copyCur.Random
			} else {
				copyCur.Random = copyRandom
			}
		}

		if cur.Next != nil {
			if copyNext, exists := nodeMap[cur.Next]; exists {
				copyCur.Next = copyNext
			} else {
				copyCur.Next = &Node{}
				nodeMap[cur.Next] = copyCur.Next
			}
		}

		copyCur.Val = cur.Val
		copyCur = copyCur.Next
		cur = cur.Next
	}

	return dummy
}
