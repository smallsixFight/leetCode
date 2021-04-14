package min_stack_lcci

/*
面试题 03.02
来源：[leetCode](https://leetcode-cn.com/)
题目：[栈的最小值](https://leetcode-cn.com/problems/min-stack-lcci/)
标签：`栈`

思路：昨天的重复题目，累计第三次重复。

官网运行结果记录
执行用时：20ms(91%)
内存消耗：8.2MB(90%)

*/

type MinStack struct {
	data [][2]int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{data: make([][2]int, 0)}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Push(x int) {
	if len(this.data) > 0 {
		this.data = append(this.data, [2]int{x, min(x, this.data[len(this.data)-1][1])})
	} else {
		this.data = append(this.data, [2]int{x, x})
	}
}

func (this *MinStack) Pop() {
	if len(this.data) > 0 {
		this.data = this.data[:len(this.data)-1]
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.data[len(this.data)-1][1]
}
