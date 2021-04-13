package bao_han_minhan_shu_de_zhan_lcof

/*
剑指 Offer 30
来源：[leetCode](https://leetcode-cn.com/)
题目：[包含min函数的栈](https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/)
标签：`栈` `设计`

思路：这道题之前也做过来着......在新元素入栈是同时记录截至当前最小的元素即可。

官网运行结果记录
执行用时：16ms(92%)/20ms(68%)
内存消耗：8.1MB(80%)

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

func (this *MinStack) Min() int {
	return this.data[len(this.data)-1][1]
}
