type Item struct {
	value int
	freq  int
}

type MinHeap []Item
var _ heap.Interface = (*MinHeap)(nil)

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	item := x.(Item)
	*h = append(*h, item)
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)

	item := old[n-1]
	*h = old[:n-1]

	return item
}

func topKFrequent(nums []int, k int) []int {
    frequency := make(map[int]int)
    for _, num := range nums {
        frequency[num] ++ 
    }
    
    minHeap := &MinHeap{}
	heap.Init(minHeap)

    for value, freq := range frequency {
		heap.Push(minHeap, Item{
			value: value,
			freq:  freq,
		})

		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

    result := make([]int, minHeap.Len())

	for i := len(result) - 1; i >= 0; i-- {
		item := heap.Pop(minHeap).(Item)
		result[i] = item.value
	}

	return result

}
