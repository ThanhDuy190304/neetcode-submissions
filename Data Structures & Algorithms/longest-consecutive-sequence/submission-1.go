func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})

	for _, num := range nums{
		m[num] = struct{}{}
	}
	result := 0
	for num := range m {
		if _, exists := m[num -1]; exists{
			continue
		}
		tmp := 1; i := 1
		for{
			if _, exists := m[num + i]; exists{
				tmp++
				i++ 
				continue
			}
			break
		}
		
		if tmp > result {
			result = tmp
		}
	}

	return result
}
