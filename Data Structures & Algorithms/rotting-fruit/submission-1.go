func orangesRotting(grid [][]int) int {
    m := len(grid)
	n := len(grid[0])

	var queue [][2]int
	countFresh := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}else if grid[i][j] == 1{
				countFresh++
			}	
		}
	}

	if countFresh == 0 {
		return 0
	}
	
	minutes := 0
	dirs := [][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

	for len(queue) > 0 && countFresh != 0{
		levelSize := len(queue)
		for l := 0; l < levelSize; l++ {
			cur := queue[0]
			queue = queue[1:]
			i, j := cur[0], cur[1]
			for _, d := range dirs{
				ni, nj := i + d[0], j + d[1]
				if ni < 0 || nj < 0 || ni >= m || nj >= n{
					continue
				}
				if grid[ni][nj] != 1 {
					continue
				}
				grid[ni][nj] = 2
				countFresh--
				queue = append(queue, [2]int{ni, nj})
			}
		}
		
		minutes++
	}

	if countFresh != 0 { return -1 } else { return minutes }

}
