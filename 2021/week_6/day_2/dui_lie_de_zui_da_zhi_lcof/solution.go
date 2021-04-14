package dui_lie_de_zui_da_zhi_lcof

/*
剑指 Offer 59 - II
来源：[leetCode](https://leetcode-cn.com/)
题目：[队列的最大值](https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/)
标签：`栈` `滑动窗口`

思路：可是使用一个数组和一个单调递减栈来实现。数组按照入队顺序保存的数据，单调递减栈则对入队数据进行排序，维持单调递减。
关于单调递减栈的使用，是由于根据队列的性质，若后续入队的元素为截至当前的最大元素，那么该元素之前的元素对于 `Max_value` 函数是没有意义的了，因为前面的元素都会先出队。

官网运行结果记录
执行用时：104ms(67%)/100ms(89%)
内存消耗：8.4MB(45%)/8.3MB(75%)

*/

type MaxQueue struct {
	data  []int
	stack []int
}

func Constructor() MaxQueue {
	return MaxQueue{data: make([]int, 0), stack: make([]int, 0)}
}

func (this *MaxQueue) Max_value() int {
	if len(this.stack) > 0 {
		return this.stack[0]
	}
	return -1
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.stack) > 0 && value > this.stack[len(this.stack)-1] {
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, value)
	this.data = append(this.data, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.data) > 0 {
		v := this.data[0]
		this.data = this.data[1:]
		if this.stack[0] == v {
			this.stack = this.stack[1:]
		}
		return v
	}
	return -1
}
