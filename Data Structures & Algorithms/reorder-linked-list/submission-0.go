/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    // 1. Find mid
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // 2. Reverse second half
    second := slow.Next
    slow.Next = nil
    var prev *ListNode
    for second != nil {
        next := second.Next
        second.Next = prev
        prev = second
        second = next
    }

    // 3. Merge
    first, second := head, prev
    for second != nil {
        tmp1, tmp2 := first.Next, second.Next
        first.Next = second
        second.Next = tmp1
        first = tmp1
        second = tmp2
    }
}
