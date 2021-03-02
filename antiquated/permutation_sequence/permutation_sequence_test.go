package permutation_sequence

import "testing"

func TestGetPermutationTwo(t *testing.T) {
	t.Log(GetPermutation(3, 1) == GetPermutationTwo(3, 1))
	t.Log(GetPermutation(3, 2) == GetPermutationTwo(3, 2))
	t.Log(GetPermutation(3, 3) == GetPermutationTwo(3, 3))
	t.Log(GetPermutation(3, 4) == GetPermutationTwo(3, 4))
	t.Log(GetPermutation(3, 5) == GetPermutationTwo(3, 5))
	t.Log(GetPermutation(4, 5) == GetPermutationTwo(4, 5))
	t.Log(GetPermutation(4, 6) == GetPermutationTwo(4, 6))
	t.Log(GetPermutation(4, 9) == GetPermutationTwo(4, 9))
	t.Log(GetPermutation(5, 20) == GetPermutationTwo(5, 20))
	t.Log(GetPermutation(8, 67) == GetPermutationTwo(8, 67))
	t.Log(GetPermutation(5, 4) == GetPermutationTwo(5, 4))
	t.Log(GetPermutation(8, 4) == GetPermutationTwo(8, 4))
	t.Log(GetPermutation(9, 10) == GetPermutationTwo(9, 10))
	t.Log(GetPermutation(9, 54434) == GetPermutationTwo(9, 54434))
}

func TestGetPermutation(t *testing.T) {
	t.Log(GetPermutation(3, 1))
	t.Log(GetPermutation(3, 2))
	t.Log(GetPermutation(3, 3))
	t.Log(GetPermutation(3, 4))
	t.Log(GetPermutation(3, 5))
	t.Log(GetPermutation(4, 5))
	t.Log(GetPermutation(4, 6))
	t.Log(GetPermutation(4, 9))
	t.Log(GetPermutation(5, 4))
	t.Log(GetPermutation(8, 50))
	t.Log(GetPermutation(9, 80))
	t.Log(GetPermutation(9, 54434))
}
