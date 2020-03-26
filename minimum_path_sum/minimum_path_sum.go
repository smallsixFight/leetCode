package minimum_path_sum

func MinPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for m := 0; m < len(grid); m++ {
		for n := 0; n < len(grid[0]); n++ {
			if m == 0 && n == 0 {
				continue
			}
			if m == 0 {
				grid[m][n] += grid[m][n-1]
			} else if n == 0 {
				grid[m][n] += grid[m-1][n]
			} else {
				if grid[m][n-1] < grid[m-1][n] {
					grid[m][n] += grid[m][n-1]
				} else {
					grid[m][n] += grid[m-1][n]
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
