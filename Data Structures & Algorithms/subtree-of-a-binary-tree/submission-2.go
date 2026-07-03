/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// Bước 1: Serialize cả 2 cây thành chuỗi
	var sbRoot, sbSub strings.Builder
	serialize(root, &sbRoot)
	serialize(subRoot, &sbSub)

	strRoot := sbRoot.String()
	strSub := sbSub.String()

	// Bước 2: Dùng thuật toán KMP để tìm chuỗi con
	return kmpSearch(strRoot, strSub)
}

// Hàm duyệt Pre-order để chuyển cây thành chuỗi
func serialize(node *TreeNode, sb *strings.Builder) {
	if node == nil {
		sb.WriteString(",#") // Ký tự đại diện cho nil
		return
	}
	// Bắt đầu bằng dấu phẩy để phân tách rõ ràng các node mang số lớn/nhỏ
	sb.WriteString(",")
	sb.WriteString(strconv.Itoa(node.Val))
	
	serialize(node.Left, sb)
	serialize(node.Right, sb)
}

// Thuật toán KMP Search chuẩn chỉnh O(M + N)
func kmpSearch(text, pattern string) bool {
	n := len(text)
	m := len(pattern)
	if m == 0 {
		return true
	}

	// Tạo mảng lps (Longest Happy Prefix / Suffix)
	lps := make([]int, m)
	j := 0
	for i := 1; i < m; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = lps[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
			lps[i] = j
		}
	}

	// Bắt đầu khớp chuỗi
	j = 0
	for i := 0; i < n; i++ {
		for j > 0 && text[i] != pattern[j] {
			j = lps[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			return true // Tìm thấy chuỗi mẫu trùng khớp hoàn toàn
		}
	}

	return false
}