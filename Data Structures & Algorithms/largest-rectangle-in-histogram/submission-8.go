func largestRectangleArea(heights []int) int {
    stack := []int{}
    result := 0

    for i := 0; i <= len(heights); i++ {
        curr := 0
        if i < len(heights) {
            curr = heights[i]
        }

        for len(stack) > 0 && curr < heights[stack[len(stack)-1]] {
            j := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            w := i
            if len(stack) > 0 {
                w = i - stack[len(stack)-1] - 1
            }

            if heights[j]*w > result {
                result = heights[j] * w
            }
        }
        stack = append(stack, i)
    }
    return result
}