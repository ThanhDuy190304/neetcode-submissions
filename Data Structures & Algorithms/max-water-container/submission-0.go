func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func maxArea(heights []int) int {
	
	left, right := 0, len(heights) - 1
	result := 0

	for (left < right){
		h := min(heights[left],  heights[right])
		area := h * (right - left)
		if (area > result){
			result = area
		}

		if (heights[left] == heights[right]){
			left ++
			right --
			continue
		}

		if (heights[left] < heights[right]){
			left++
			continue
		}

		if (heights[left] > heights[right]){
			right--
			continue
		}
	}
	return result

}
