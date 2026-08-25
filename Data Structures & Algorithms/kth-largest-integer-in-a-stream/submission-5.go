
// MinHeap là 1 slice int, implement interface heap.Interface
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] } // min-heap: phần tử nhỏ nhất ở gốc
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }

// Push/Pop phải dùng con trỏ vì cần thay đổi độ dài slice
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// KthLargest quản lý 1 min-heap kích thước k
// -> phần tử nhỏ nhất trong heap luôn là "kth largest" của toàn bộ stream
type KthLargest struct {
	k int
	h *MinHeap
}

func Constructor(k int, nums []int) *KthLargest {
	h := &MinHeap{}
	heap.Init(h)
	kl := &KthLargest{k: k, h: h}
	for _, n := range nums {
		kl.Add(n)
	}
	return kl
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl.h, val)
	if kl.h.Len() > kl.k {
		heap.Pop(kl.h) // loại phần tử nhỏ nhất, chỉ giữ k phần tử lớn nhất
	}
	return (*kl.h)[0] // gốc heap = kth largest hiện tại
}