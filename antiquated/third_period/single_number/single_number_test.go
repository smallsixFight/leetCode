package single_number

import "testing"

func TestSingleNumber(t *testing.T) {
	t.Log(SingleNumber([]int{1}))
	t.Log(SingleNumber([]int{2, 2, 1}))
	t.Log(SingleNumber([]int{4, 1, 2, 1, 2}))
}
