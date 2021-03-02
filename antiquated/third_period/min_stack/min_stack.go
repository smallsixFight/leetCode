package min_stack

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinStack struct {
	data  []int
	order *ListNode
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:  make([]int, 0),
		order: &ListNode{},
	}
}

func (this *MinStack) Push(x int) {
	head := this.order
	for head.Next != nil && this.data[head.Next.Val] < x {
		head = head.Next
	}
	temp := head.Next
	head.Next = &ListNode{Val: len(this.data)}
	head.Next.Next = temp
	this.data = append(this.data, x)
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 || this.order.Next == nil {
		return
	}
	head := this.order
	for head.Next != nil && head.Next.Val != len(this.data)-1 {
		head = head.Next
	}
	head.Next = head.Next.Next
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return 0
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.data) == 0 {
		return 0
	}
	return this.data[this.order.Next.Val]
}
