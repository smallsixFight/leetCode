package same_tree

import "testing"

func Test_isSameTree(t *testing.T) {
	t1 := TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	t2 := TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	t.Log(isSameTree(&t1, &t2))

	t3 := TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	t4 := TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	t.Log(isSameTree(&t3, &t4))
}
