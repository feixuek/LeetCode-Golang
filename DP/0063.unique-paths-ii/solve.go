func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return -1
	}
	var res [][]int
	for i := 0; i < len(obstacleGrid); i++ {
		res = append(res, make([]int, len(obstacleGrid[i])))
		if i == 0 {
			if obstacleGrid[i][0] == 0 {
				res[i][0] = 1
			}
			for j := 1; j < len(obstacleGrid[i]); j++ {
				if res[i][j-1] == 1 && obstacleGrid[i][j] == 0 {
					res[i][j] = 1
				}
			}
		} else {
			if res[i-1][0] == 1 && obstacleGrid[i][0] == 0 {
				res[i][0] = 1
			}
		}
	}
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 0 {
				switch {
				case obstacleGrid[i-1][j] == 1 && obstacleGrid[i][j-1] == 1:
				case obstacleGrid[i-1][j] == 1:
					res[i][j] = res[i][j-1]
				case obstacleGrid[i][j-1] == 1:
					res[i][j] = res[i-1][j]
				default:
					res[i][j] = res[i-1][j] + res[i][j-1]
				}
			}
		}
	}
	return res[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}