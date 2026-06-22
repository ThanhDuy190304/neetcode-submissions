func carFleet(target int, position []int, speed []int) int {
	idx := make([]int, len(position))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return position[idx[i]] > position[idx[j]]
	})

	var stack []float64

	for i := 0; i < len(position); i++{
		timeToTarget := float64(target - position[idx[i]]) / float64(speed[idx[i]])
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
