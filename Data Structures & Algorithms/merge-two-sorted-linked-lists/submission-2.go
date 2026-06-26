/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    cur := dummy
    p1, p2 := list1, list2

    for p1 != nil && p2 != nil {
        if p1.Val <= p2.Val {
            cur.Next = p1
            p1 = p1.Next
        } else {
            cur.Next = p2
            p2 = p2.Next
        }
        cur = cur.Next
    }

    if p1 != nil {
        cur.Next = p1
    } else {
        cur.Next = p2
    }

    return dummy.Next
}
