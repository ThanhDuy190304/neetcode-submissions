func lengthOfLongestSubstring(s string) int {
	lastIndex := [256]int{}

	for i := range lastIndex {
		lastIndex[i] = -1
	}

	left := 0
	result := 0

	for right := 0; right < len(s); right++ {
		c := s[right]

		if lastIndex[c] >= left {
			left = lastIndex[c] + 1
		}

		lastIndex[c] = right

		length := right - left + 1
		if length > result {
			result = length
		}
	}

	return result
}