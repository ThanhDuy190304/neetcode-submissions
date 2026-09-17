func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	
	visited := make([][]bool, m)

	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		//out of bounds
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		// water or visited 
		if grid[i][j] == '0' || visited[i][j] {
			return
		}

		visited[i][j] = true 

		dfs(i-1, j) // top
		dfs(i+1, j) // bottom
		dfs(i, j-1) // lefft
		dfs(i, j+1) // right
	}
		
	
	result := 0 
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				dfs(i, j)
				result++
			}
		}
	}

	return result
}
