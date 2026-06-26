/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	
	var prev, result *ListNode
	p1, p2 := list1, list2

	for p1 != nil && p2 != nil{
		var cur *ListNode
		if p1.Val <= p2.Val{
			cur = &ListNode{
				Val: p1.Val,
				Next: nil,
			}
			p1 = p1.Next
		}else{
			cur = &ListNode{
				Val: p2.Val,
				Next: nil,
			}
			p2 = p2.Next
		}

		if prev != nil{
			prev.Next = cur
		}
		prev = cur

		if result == nil {
			result = cur
		}
	}

	for p1 != nil{
		cur := &ListNode{
				Val: p1.Val,
				Next: nil,
			}
		if prev != nil{
			prev.Next = cur
		}
		prev = cur
		p1 = p1.Next
		
		if result == nil {
			result = cur
		}
	}

	for p2 != nil{
		cur := &ListNode{
				Val: p2.Val,
				Next: nil,
			}
		if prev != nil{
			prev.Next = cur
		}
		prev = cur
		p2 = p2.Next

		if result == nil {
			result = cur
		}
	}

	return result
}
