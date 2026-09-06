func subsets(nums []int) [][]int {
	n := len(nums)
	total := 1 << n
	res := make([][]int, 0, total)

	for mask := 0; mask < total; mask++ {
		var subset []int
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				subset = append(subset, nums[i])
			}
		}
		res = append(res, subset)
	}

	return res
}