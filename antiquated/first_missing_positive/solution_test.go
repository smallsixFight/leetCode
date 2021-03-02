package first_missing_positive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	t.Log(FirstMissingPositive([]int{1, 2, 0}))
	t.Log(FirstMissingPositive([]int{3, 4, -1, 1}))
	t.Log(FirstMissingPositive([]int{7, 8, 9, 11, 12}))
	t.Log(FirstMissingPositive([]int{1, 2, 3}))
}
