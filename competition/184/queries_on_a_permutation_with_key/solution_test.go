package queries_on_a_permutation_with_key

import "testing"

func TestProcessQueries(t *testing.T) {
	t.Log(ProcessQueries([]int{3, 1, 2, 1}, 5))
	t.Log(ProcessQueries([]int{4, 1, 2, 2}, 4))
	t.Log(ProcessQueries([]int{7, 5, 5, 8, 3}, 8))
}
