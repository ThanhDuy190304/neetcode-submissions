func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}
		if grid[i][j] == 0 || visited[i][j] {
			return 0
		}

		visited[i][j] = true
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
			count := dfs(i, j)
			if count > result {
				result = count
			}
		}
	}

	return result
}