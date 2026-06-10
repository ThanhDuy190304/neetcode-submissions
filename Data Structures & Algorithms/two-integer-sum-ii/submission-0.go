func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers) - 1

	for(left < right){
		if (target == numbers[right] + numbers[left]){
			return []int{left + 1, right + 1}
		}
		if(target < numbers[right] + numbers[left]){
			right --
			continue
		}
		if(target > numbers[right] + numbers[left]){
			left ++
			continue
		}
	}

	return nil
}
