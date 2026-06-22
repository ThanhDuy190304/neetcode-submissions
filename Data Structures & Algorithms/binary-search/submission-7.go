func search(nums []int, target int) int {
	
	left, right := 0, len(nums) - 1
	idx := (left + right) / 2
	for (left <= right) {
		if(nums[idx] > target){
			right = idx - 1
		}else if(nums[idx] < target){
			left = idx + 1
		}else{
			return idx
		}
		idx = (right + left) / 2
	}
	return -1
}
