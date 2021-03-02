package min_stack_2

type node struct {
	val    int
	curMin int
}

type MinStack struct {
	data []node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{data: make([]node, 0)}
}

func (this *MinStack) Push(x int) {
	n := node{val: x, curMin: x}
	if len(this.data) > 0 && this.data[len(this.data)-1].curMin < x {
		n.curMin = this.data[len(this.data)-1].curMin
	}
	this.data = append(this.data, n)
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1].val
}

func (this *MinStack) GetMin() int {
	return this.data[len(this.data)-1].curMin
}
