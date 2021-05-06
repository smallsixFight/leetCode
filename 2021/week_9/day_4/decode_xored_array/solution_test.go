package same_tree

import "testing"

func Test_decode(t *testing.T) {
	t.Log(decode([]int{1, 2, 3}, 1))    // [1,0,2,1]
	t.Log(decode([]int{6, 2, 7, 3}, 4)) // [4,2,0,7,4]
}
