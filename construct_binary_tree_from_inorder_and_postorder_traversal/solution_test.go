package construct_binary_tree_from_inorder_and_postorder_traversal

import "testing"

func TestBuildTree(t *testing.T) {
	tree := BuildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	t.Log(printTree(tree))
}
