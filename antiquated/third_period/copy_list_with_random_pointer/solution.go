package copy_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	top := head
	m := make(map[*Node]*Node)
	for head != nil {
		if _, ok := m[head]; !ok {
			n1 := &Node{Val: head.Val}
			m[head] = n1
		}
		if head.Next != nil {
			if _, ok := m[head.Next]; !ok {
				n1 := &Node{Val: head.Next.Val}
				m[head].Next = n1
				m[head.Next] = n1
			} else {
				m[head].Next = m[head.Next]
			}
		}
		if head.Random != nil {
			if _, ok := m[head.Random]; !ok {
				n1 := &Node{Val: head.Random.Val}
				m[head].Random = n1
				m[head.Random] = n1
			} else {
				m[head].Random = m[head.Random]
			}
		}
		head = head.Next
	}
	return m[top]
}
