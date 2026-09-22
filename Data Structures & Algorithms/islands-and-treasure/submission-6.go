func islandsAndTreasure(grid [][]int) {
	m := len(grid)
	n := len(grid[0])
	queue := [][2]int{}
	for i := 0; i < m; i++{
		for j := 0; j < n; j ++ {
			if grid[i][j] == 0{
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0{
		cur := queue[0]
		queue = queue[1:]
		i, j := cur[0], cur[1]

		for _, d := range dirs {
			ni, nj := d[0] + i, d[1] + j
			if ni < 0 || nj < 0 || ni >= m || nj >= n{
				continue
			} 
			if grid[ni][nj] != 2147483647{
				continue
			}
			grid[ni][nj] = 1 + grid[i][j]
			queue = append(queue, [2]int{ni, nj})
		}
	}

}
