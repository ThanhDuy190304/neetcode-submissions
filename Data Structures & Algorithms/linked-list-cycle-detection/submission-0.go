/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	mapLNs := make(map[*ListNode]struct{})
	for head != nil{
		if _, exists := mapLNs[head]; exists{
			return true
		}
		mapLNs[head] = struct{}{}
		head = head.Next
	}
	return false 
}
