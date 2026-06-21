type Stack struct {
	idx []int
}

func (s *Stack) Push(i int) {
	s.idx = append(s.idx, i)
}

func (s *Stack) Pop() {
	s.idx = s.idx[:len(s.idx)-1]
}

func (s *Stack) Top() int {
	return s.idx[len(s.idx)-1]
}

func (s *Stack) Len() int {
	return len(s.idx)
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var st Stack

	for i := len(temperatures) - 1; i >= 0; i-- {
		for st.Len() > 0 && temperatures[st.Top()] <= temperatures[i] {
			st.Pop()
		}
		if st.Len() > 0 {
			result[i] = st.Top() - i
		}
		st.Push(i)
	}
	return result
}