func findKthLargest(nums []int, k int) int {
	target := len(nums) - k
	left, right := 0, len(nums)-1

	for left <= right {
		pivotIdx := left + rand.Intn(right-left+1)
		pivotFinalIdx := partition(nums, left, right, pivotIdx)

		if pivotFinalIdx == target {
			return nums[pivotFinalIdx]
		} else if pivotFinalIdx < target {
			left = pivotFinalIdx + 1
		} else {
			right = pivotFinalIdx - 1
		}
	}

	return -1
}

func partition(nums []int, left, right, pivotIdx int) int {
	pivotVal := nums[pivotIdx]
	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]
	storeIdx := left

	for i := left; i < right; i++ {
		if nums[i] < pivotVal {
			nums[storeIdx], nums[i] = nums[i], nums[storeIdx]
			storeIdx++
		}
	}
	nums[right], nums[storeIdx] = nums[storeIdx], nums[right]
	return storeIdx
}