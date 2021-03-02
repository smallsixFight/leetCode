package unique_paths

import "testing"

func TestUniquePathsTwo(t *testing.T) {
	t.Log(UniquePathsTwo(3, 2))
	t.Log(UniquePathsTwo(7, 3))
	t.Log(UniquePaths(23, 12) == UniquePathsTwo(23, 12))
	t.Log(UniquePaths(32, 12) == UniquePathsTwo(32, 12))
	t.Log(UniquePaths(44, 13) == UniquePathsTwo(44, 13))
	t.Log(UniquePaths(13, 4) == UniquePathsTwo(13, 4))
	t.Log(UniquePaths(33, 12) == UniquePathsTwo(33, 12))
	t.Log(UniquePaths(59, 5))
	t.Log(UniquePathsTwo(59, 5))
}

func TestUniquePaths(t *testing.T) {
	t.Log(UniquePaths(3, 2))
	t.Log(UniquePaths(7, 3))
}
