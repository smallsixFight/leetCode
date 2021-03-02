package number_of_islands

func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	res := 0
	bfs := func(x, y int) {
		queue := [][2]int{{x, y}}
		for len(queue) > 0 {
			l := len(queue)
			for i := 0; i < l; i++ {
				a, b := queue[i][0], queue[i][1]
				if a > 0 && grid[a-1][b] == '1' {
					grid[a-1][b] = 'a'
					queue = append(queue, [2]int{a - 1, b})
				}
				if a < len(grid)-1 && grid[a+1][b] == '1' {
					grid[a+1][b] = 'a'
					queue = append(queue, [2]int{a + 1, b})
				}
				if b > 0 && grid[a][b-1] == '1' {
					grid[a][b-1] = 'a'
					queue = append(queue, [2]int{a, b - 1})
				}
				if b < len(grid[a])-1 && grid[a][b+1] == '1' {
					grid[a][b+1] = 'a'
					queue = append(queue, [2]int{a, b + 1})
				}
			}
			queue = queue[l:]
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				grid[i][j] = 'a'
				bfs(i, j)
			}
		}
	}
	return res
}
