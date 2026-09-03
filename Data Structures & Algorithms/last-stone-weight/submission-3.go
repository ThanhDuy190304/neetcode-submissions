type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } 
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func lastStoneWeight(stones []int) int {
	// Ép kiểu trực tiếp từ slice ban đầu
	h := (*MaxHeap)(&stones)
	heap.Init(h) // Thực hiện heapify tại chỗ với độ phức tạp O(N)

	// Mô phỏng quá trình đập đá
	for h.Len() > 1 {
		first := heap.Pop(h).(int)
		second := heap.Pop(h).(int)

		if first != second {
			heap.Push(h, first-second)
		}
	}

	if h.Len() == 0 {
		return 0
	}
	return (*h)[0]
}