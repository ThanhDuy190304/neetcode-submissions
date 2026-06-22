func largestRectangleArea(heights []int) int {
	var stackIdx []int
	var result int

	for i := 0; i < len(heights); i++{
		if(len(stackIdx) == 0){
			if (heights[i] > result){
				result = heights[i]
			}
			stackIdx = append(stackIdx, i)
			continue
		}
		if(heights[i] == heights[stackIdx[len(stackIdx) - 1]]){
			tmp := heights[i] * (stackIdx[len(stackIdx) - 1] - i + 1) 
			if (tmp > result){
				result = tmp
			}
		}else if(heights[i] > heights[stackIdx[len(stackIdx) - 1]]){
			if (heights[i] > result){
				result = heights[i]
			}
			stackIdx = append(stackIdx, i)
			continue
		}

		for len(stackIdx) != 0 && heights[i] < heights[stackIdx[len(stackIdx)-1]] {
			j := stackIdx[len(stackIdx)-1]
			stackIdx = stackIdx[:len(stackIdx)-1]  

			var w int
			if len(stackIdx) == 0 {
				w = i  
			} else {
				w = i - stackIdx[len(stackIdx)-1] - 1
			}

			if heights[j]*w > result {
				result = heights[j] * w
			}
		}
		stackIdx = append(stackIdx, i)

	}

	for i := 0; i < len(stackIdx); i++{
		tmp := 0
		if(i == 0){
			tmp = heights[stackIdx[i]] * len(heights)
		}else{
			tmp = heights[stackIdx[i]] * (stackIdx[len(stackIdx) - 1] - stackIdx[i-1])
		}

		if(tmp > result){
			result = tmp
		}
	}
	return result
}