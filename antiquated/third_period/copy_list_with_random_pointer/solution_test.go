package copy_list_with_random_pointer

import "testing"

func TestCopyRandomList(t *testing.T) {
	n5 := &Node{Val: 1}
	n4 := &Node{Val: 10, Next: n5}
	n3 := &Node{Val: 11, Next: n4}
	n2 := &Node{Val: 13, Next: n3}
	n1 := &Node{Val: 7, Next: n2}
	n2.Random = n1
	n3.Random = n5
	n4.Random = n3
	n5.Random = n1
	t.Log(CopyRandomList(n1))

}
