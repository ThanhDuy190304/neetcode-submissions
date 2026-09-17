func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}

		grid[i][j] = 0 //mark visited instead of a new visited array
		area := 1
		area += dfs(i-1, j)
		area += dfs(i+1, j)
		area += dfs(i, j-1)
		area += dfs(i, j+1)
		return area
	}

	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if area := dfs(i, j); area > result {
				result = area
			}
		}
	}

	return result
}