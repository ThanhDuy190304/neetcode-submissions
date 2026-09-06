func subsets(nums []int) [][]int {
	res := [][]int{{}}

	for _, n := range nums {
		size := len(res)
		for i := 0; i < size; i++ {
			newSet := make([]int, len(res[i]), len(res[i])+1)
			copy(newSet, res[i])
			newSet = append(newSet, n)
			res = append(res, newSet)
		}
	}

	return res
}