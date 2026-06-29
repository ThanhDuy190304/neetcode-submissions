/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    l, r := dummy, dummy

    // Đẩy r đi n+1 bước
    for i := 0; i <= n; i++ {
        r = r.Next
    }

    // Di chuyển đồng thời đến khi r == nil
    for r != nil {
        l = l.Next
        r = r.Next
    }

    // l đang ở node trước node cần xóa
    l.Next = l.Next.Next
    return dummy.Next
}