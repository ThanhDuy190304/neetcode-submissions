type Point struct {
	coords []int
	distSq int
}

type MaxHeap []Point

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].distSq > h[j].distSq } 
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Point))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, pt := range points {
		distSq := pt[0]*pt[0] + pt[1]*pt[1]
		p := Point{coords: pt, distSq: distSq}

		if h.Len() < k {
			heap.Push(h, p)
		} else if distSq < (*h)[0].distSq {
			heap.Pop(h)
			heap.Push(h, p)
		}
	}

	result := make([][]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(h).(Point).coords
	}

	return result
}
