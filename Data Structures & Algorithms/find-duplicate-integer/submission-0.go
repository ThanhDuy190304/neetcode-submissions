func findDuplicate(nums []int) int {
	
	slow, fast := nums[0], nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	slow2 := 0
    for slow != slow2 {
        slow = nums[slow]
        slow2 = nums[slow2]
    }

    return slow

}
