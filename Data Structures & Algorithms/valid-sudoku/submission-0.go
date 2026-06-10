func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		//check i col
		rMap := make(map[byte]struct{}, 9)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.'{
				continue
			}
			if _, exists := rMap[board[i][j]]; exists{
				return false
			}
			rMap[board[i][j]] = struct{}{}	
		}
		//check i row
		cMap := make(map[byte]struct{}, 9)
		for j := 0; j < 9; j++ {
			if board[j][i] == '.'{
				continue
			}
			if _, exists := cMap[board[j][i]]; exists{
				return false
			}
			cMap[board[j][i]] = struct{}{}	
		}

		//check box 
		xBoxOrigin := (i/3)*3 + 1
		yBoxOrigin := (i%3)*3 + 1
		boxMap := make(map[byte]struct{}, 9)
		boxMap[board[xBoxOrigin][yBoxOrigin]] = struct{}{}
		for j := -1; j <= 1; j++ {
			if(j == 0){
				continue
			}
			if (board[xBoxOrigin + j][yBoxOrigin]) != '.' {
				if _, exists := boxMap[board[xBoxOrigin + j][yBoxOrigin]]; exists{
					return false
				}
				boxMap[board[xBoxOrigin + j][yBoxOrigin]] = struct{}{}
			}

			if (board[xBoxOrigin][yBoxOrigin + j]) != '.' {
				if _, exists := boxMap[board[xBoxOrigin][yBoxOrigin+j]]; exists{
					return false
				}
				boxMap[board[xBoxOrigin][yBoxOrigin+j]] = struct{}{}
			}

			if (board[xBoxOrigin+j][yBoxOrigin + j]) != '.' {
				if _, exists := boxMap[board[xBoxOrigin+j][yBoxOrigin+j]]; exists{
					return false
				}
				boxMap[board[xBoxOrigin+j][yBoxOrigin+j]] = struct{}{}
			}

			if (board[xBoxOrigin-j][yBoxOrigin + j]) != '.' {
				if _, exists := boxMap[board[xBoxOrigin-j][yBoxOrigin+j]]; exists{
					return false
				}
				boxMap[board[xBoxOrigin-j][yBoxOrigin+j]] = struct{}{}
			}			
		}
	}
	
	return true
}