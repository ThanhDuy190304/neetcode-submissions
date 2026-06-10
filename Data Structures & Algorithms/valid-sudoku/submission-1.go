func isValidSudoku(board [][]byte) bool {
	checkAndAdd := func(seen map[byte]struct{}, value byte) bool {
		if value == '.' {
			return true
		}

		if _, exists := seen[value]; exists {
			return false
		}

		seen[value] = struct{}{}
		return true
	}

	for i := 0; i < 9; i++ {
		// Kiểm tra hàng i
		rowMap := make(map[byte]struct{}, 9)

		for j := 0; j < 9; j++ {
			if !checkAndAdd(rowMap, board[i][j]) {
				return false
			}
		}

		// Kiểm tra cột i
		colMap := make(map[byte]struct{}, 9)

		for j := 0; j < 9; j++ {
			if !checkAndAdd(colMap, board[j][i]) {
				return false
			}
		}

		// Kiểm tra box i
		xBoxOrigin := (i/3)*3 + 1
		yBoxOrigin := (i%3)*3 + 1

		boxMap := make(map[byte]struct{}, 9)

		if !checkAndAdd(boxMap, board[xBoxOrigin][yBoxOrigin]) {
			return false
		}

		for j := -1; j <= 1; j++ {
			if j == 0 {
				continue
			}

			if !checkAndAdd(
				boxMap,
				board[xBoxOrigin+j][yBoxOrigin],
			) {
				return false
			}

			if !checkAndAdd(
				boxMap,
				board[xBoxOrigin][yBoxOrigin+j],
			) {
				return false
			}

			if !checkAndAdd(
				boxMap,
				board[xBoxOrigin+j][yBoxOrigin+j],
			) {
				return false
			}

			if !checkAndAdd(
				boxMap,
				board[xBoxOrigin-j][yBoxOrigin+j],
			) {
				return false
			}
		}
	}

	return true
}