package leaf_similar_trees

import (
	"testing"
)

func Test_decode(t *testing.T) {
	t.Log(decode([]int{3, 1}))       // [1 2 3]
	t.Log(decode([]int{6, 5, 4, 6})) // [2 4 1 5 3]
}
