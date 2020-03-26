package unique_paths_ii

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	grid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	t.Log(UniquePathsWithObstacles(grid))
	grid = [][]int{{0, 1}}
	t.Log(UniquePathsWithObstacles(grid))
}
