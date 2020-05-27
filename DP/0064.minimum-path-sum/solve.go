func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var res [][]int
	for i := 0; i < len(grid); i++ {
		res = append(res, make([]int, len(grid[i])))
	}
	res[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		res[i][0] = grid[i][0] + res[i-1][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		res[0][i] = grid[0][i] + res[0][i-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			res[i][j] = min(res[i-1][j], res[i][j-1]) + grid[i][j]
		}
	}
	return res[len(grid)-1][len(grid[0])-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}