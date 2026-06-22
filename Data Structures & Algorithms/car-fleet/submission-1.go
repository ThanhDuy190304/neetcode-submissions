func carFleet(target int, position []int, speed []int) int {
	pairs := make([][2]int, len(position))
	for i := range position {
    	pairs[i] = [2]int{position[i], speed[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
    	return pairs[i][0] > pairs[j][0]
	})
	var stack []float64

	for i := 0; i < len(position); i++{
		timeToTarget := float64(target - pairs[i][0]) / float64(pairs[i][1])
		if (len(stack) == 0){
			stack = append(stack, timeToTarget)
			continue
		}	
		if (timeToTarget > stack[len(stack) - 1]){
			stack = append(stack, timeToTarget)
			continue
		}
	}

	return len(stack)
}
