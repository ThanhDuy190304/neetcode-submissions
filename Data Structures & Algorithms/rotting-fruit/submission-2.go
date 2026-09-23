func orangesRotting(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    
    // Cấp phát trước dung lượng tối đa để tránh re-allocation
    queue := make([][2]int, 0, m*n)
    countFresh := 0

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 2 {
                queue = append(queue, [2]int{i, j})
            } else if grid[i][j] == 1 {
                countFresh++
            }
        }
    }

    if countFresh == 0 {
        return 0
    }

    minutes := 0
    head := 0 // Dùng con trỏ thay vì queue = queue[1:]
    dirs := [][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

    for head < len(queue) && countFresh > 0 {
        levelSize := len(queue) - head
        for l := 0; l < levelSize; l++ {
            cur := queue[head]
            head++
            i, j := cur[0], cur[1]

            for _, d := range dirs {
                ni, nj := i+d[0], j+d[1]
                if ni < 0 || nj < 0 || ni >= m || nj >= n || grid[ni][nj] != 1 {
                    continue
                }
                grid[ni][nj] = 2
                countFresh--
                queue = append(queue, [2]int{ni, nj})
            }
        }
        minutes++
    }

    if countFresh > 0 {
        return -1
    }
    return minutes
}