package single_number_ii

import "testing"

func TestSingleNumbers(t *testing.T) {
	t.Log(SingleNumbers([]int{2, 2, 3, 2}))
	t.Log(SingleNumbers([]int{0, 1, 0, 1, 0, 1, 99}))
}
