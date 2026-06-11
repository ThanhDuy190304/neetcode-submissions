func lengthOfLongestSubstring(s string) int {
	left := 0
	result := 0
	lastIndex := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		if index, exists := lastIndex[s[right]]; exists && index >= left {
			left = index + 1
		}

		lastIndex[s[right]] = right

		length := right - left + 1
		if length > result {
			result = length
		}
	}

	return result
}