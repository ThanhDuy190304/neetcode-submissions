func threeSum(nums []int) [][]int {
	pairBySum := make(map[int][][2]int)
	indicesByValue := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]

			pairBySum[sum] = append(
				pairBySum[sum],
				[2]int{i, j},
			)
		}

		indicesByValue[nums[i]] = append(indicesByValue[nums[i]], i)
	}

	result := make([][]int, 0)
	seen := make(map[[3]int]struct{})

	for sum, pairs := range pairBySum {
		indices, exists := indicesByValue[-sum]
		if !exists {
			continue
		}

		for _, pair := range pairs {
			for _, thirdIndex := range indices {
				if thirdIndex == pair[0] || thirdIndex == pair[1] {
					continue
				}

				triple := []int{
					nums[pair[0]],
					nums[pair[1]],
					nums[thirdIndex],
				}

				sort.Ints(triple)

				key := [3]int{
					triple[0],
					triple[1],
					triple[2],
				}

				if _, exists := seen[key]; exists {
					continue
				}

				seen[key] = struct{}{}
				result = append(result, triple)
			}
		}
	}

	return result
}