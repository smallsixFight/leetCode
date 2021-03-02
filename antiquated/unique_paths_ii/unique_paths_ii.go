package unique_paths_ii

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	pre, cur := make([]int, len(obstacleGrid[0])), make([]int, len(obstacleGrid[0]))
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if i == 0 && j == 0 {
				if obstacleGrid[0][0] != 1 {
					cur[j] = 1
				}
			} else if i == 0 && obstacleGrid[0][j] != 1 {
				cur[j] = cur[j-1]
			} else if j == 0 && obstacleGrid[i][0] != 1 {
				cur[j] = pre[j]
			} else {
				if obstacleGrid[i][j] != 1 {
					cur[j] = cur[j-1] + pre[j]
				} else {
					cur[j] = 0
				}
			}
		}
		pre = cur
	}
	return cur[len(obstacleGrid[0])-1]
}
