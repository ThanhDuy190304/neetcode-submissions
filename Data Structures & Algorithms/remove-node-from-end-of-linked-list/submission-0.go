/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 0
	tmp := head
	for tmp != nil{
		l++
		tmp = tmp.Next
	}
	if n == l{
		return head.Next
	}

	var prev *ListNode
	next := head
	for i := 0; i < l; i++{
		if i == l - n{
			prev.Next = next.Next 
			break
		}
		prev = next
		next = next.Next
	}
	return head
}
