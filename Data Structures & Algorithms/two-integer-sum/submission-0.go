func twoSum(nums []int, target int) []int {
    set := make(map[int]int)
    for i, num := range nums {
        second :=  target - num; 
       	if secondIndex, exists := set[second]; exists {
			return []int{secondIndex, i}
		}

        set[num] = i;
    }
    return nil
}
