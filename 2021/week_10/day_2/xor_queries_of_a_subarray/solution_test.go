package xor_queries_of_a_subarray

import (
	"testing"
)

func Test_xorQueries(t *testing.T) {
	t.Log(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))  // [2 7 14 8]
	t.Log(xorQueries([]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}})) // [8 0 4 4]
	//t.Log(31^0)
}
