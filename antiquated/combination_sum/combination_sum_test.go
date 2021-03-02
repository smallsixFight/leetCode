package combination_sum

import "testing"

func TestCombinationSum(t *testing.T) {
	t.Log(CombinationSum([]int{2, 3, 5, 7}, 2))
	t.Log(CombinationSum([]int{2, 3, 5, 7}, 3))
	t.Log(CombinationSum([]int{2, 3, 5, 7}, 5))
	t.Log(CombinationSum([]int{2, 3, 5, 7}, 7))
	t.Log(CombinationSum([]int{2, 3, 6, 7}, 7))
	t.Log(CombinationSum([]int{2, 3, 5}, 8))
	t.Log(CombinationSum([]int{8, 7, 4, 3}, 11))
	t.Log(CombinationSum([]int{7, 3, 2}, 18))
}
