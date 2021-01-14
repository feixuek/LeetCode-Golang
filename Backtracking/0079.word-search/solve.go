func exist(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])
	target := []byte(word)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == target[0] {
				if backtrace(board, target, i, j, row, col, 0) {
					return true
				}
			}
		}
	}
	return false
}

func backtrace(board [][]byte, word []byte, startRow, startCol, row, col, target int) bool {
	if target == len(word) {
		return true
	}
	if startRow < 0 || startCol < 0 || startRow >= row || startCol >= col || board[startRow][startCol] != word[target] {
		return false
	}
	tmp := board[startRow][startCol]
	board[startRow][startCol] = '0'
	res := backtrace(board, word, startRow, startCol+1, row, col, target+1) || backtrace(board, word, startRow, startCol-1, row, col, target+1) || backtrace(board, word, startRow+1, startCol, row, col, target+1) || backtrace(board, word, startRow-1, startCol, row, col, target+1)
	board[startRow][startCol] = tmp
	return res
}