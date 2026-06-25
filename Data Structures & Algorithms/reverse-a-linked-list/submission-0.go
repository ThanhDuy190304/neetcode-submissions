/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode
	for {
		if head == nil {
			break
		}
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
