package permutations_ii

import "testing"

func TestPermutation(t *testing.T) {
	t.Log(Permutation([]int{1, 1, 2}))
	t.Log(Permutation([]int{2, 1, 2}))
	t.Log(Permutation([]int{2, 2, 1}))
}
