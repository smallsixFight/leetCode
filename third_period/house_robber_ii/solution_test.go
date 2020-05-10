package house_robber_ii

import "testing"

func TestRob(t *testing.T) {
	t.Log(Rob([]int{2, 3, 2}))
	t.Log(Rob([]int{1, 2, 3, 1}))
	t.Log(Rob([]int{4, 1, 2, 7, 5, 3, 1}))
	t.Log(Rob([]int{1, 1, 3, 6, 7, 10, 7, 1, 8, 5, 9, 1, 4, 4, 3}))
}
