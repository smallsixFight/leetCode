package same_tree

import "testing"

func Test_xorOperation(t *testing.T) {
	t.Log(xorOperation(5, 0))  // 8
	t.Log(xorOperation(4, 3))  // 8
	t.Log(xorOperation(1, 7))  // 7
	t.Log(xorOperation(10, 5)) // 2
}
