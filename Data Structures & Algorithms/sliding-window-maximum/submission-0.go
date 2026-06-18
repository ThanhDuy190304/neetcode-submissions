func maxSlidingWindow(nums []int, k int) []int {
    deque := []int{} 
    result := []int{}

    for r := 0; r < len(nums); r++ {

        if len(deque) > 0 && deque[0] < r-k+1 {
            deque = deque[1:]
        }

        for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[r] {
            deque = deque[:len(deque)-1]
        }

        deque = append(deque, r)

        if r >= k-1 {
            result = append(result, nums[deque[0]])
        }
    }

    return result
}