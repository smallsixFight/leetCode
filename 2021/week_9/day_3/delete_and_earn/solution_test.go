package same_tree

import "testing"

func Test_deleteAndEarn(t *testing.T) {
	t.Log(deleteAndEarn([]int{3, 4, 2}))                   // 6
	t.Log(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))          // 9
	t.Log(deleteAndEarn([]int{1, 1, 1, 2, 4, 5, 5, 5, 6})) // 18
}
