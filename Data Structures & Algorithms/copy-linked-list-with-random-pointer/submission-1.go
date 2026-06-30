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

	for cur := head; cur != nil; cur = cur.Next.Next {
		copyNode := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = copyNode
	}

	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}

	dummy := &Node{}
	copyCur := dummy
	for cur := head; cur != nil; cur = cur.Next {
		copyCur.Next = cur.Next
		copyCur = copyCur.Next
		cur.Next = cur.Next.Next
	}

	return dummy.Next
}
