func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	mResult := make(map[[3]int]struct{})

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			total := nums[i] + nums[j]
			right := len(nums) - 1

			for j < right {
				if nums[right]+total == 0 {
					mResult[[3]int{
						nums[i],
						nums[j],
						nums[right],
					}] = struct{}{}
					break
				}

				right--
			}
		}
	}

	var result [][]int

	for k := range mResult {
		result = append(result, []int{
			k[0],
			k[1],
			k[2],
		})
	}

	return result
}