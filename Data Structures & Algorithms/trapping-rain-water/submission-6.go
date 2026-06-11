func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func trap(height []int) int {
	result := 0
	l, r := 0, len(height) - 1
	for (l < r){
		if (height[l] == 0){
			l++
			continue
		}
		if (height[r] == 0){
			r--
			continue
		}

		h := min(height[l], height[r])


		for(l < r && height[l] <= height[r]){
			if(height[l] > h){
				break
			}
			result += h - height[l]
			l++
		}

		for(l < r && height[l] > height[r]){
			if(height[r] > h){
				break
			}
			result += h - height[r]
			r --
		}
	}

	return result

}
