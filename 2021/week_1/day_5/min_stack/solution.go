package min_stack

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[最小栈](https://leetcode-cn.com/problems/min-stack/)
标签：`栈`

简单思路：重点在于如何在常数范围内获取栈内的最小。一种方法是维护两个 slice，一个是按照 push 顺序的数据，
一个是排序后的数据，但这样每次 push 都要进行排序，pop 也需要遍历查找去删除排序后的数据，而且每次 getMin
都只是获取排序 slice 的第一个值而已。
那么在 push 一个值的时候同时记录截止当前的最小值，每次 push 时的最小值只需要将当前的值跟栈顶的元素的最小值比较
取其最小值存起来就好。

官网运行结果记录
执行用时：20ms/24ms
内存消耗：8.1MB/9.1MB

*/

type node struct {
	val    int
	curMin int
}

type MinStack struct {
	data []node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
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

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
