package design_a_stack_with_increment_operation

/*
1381
来源：[leetCode](https://leetcode-cn.com/)
题目：[设计一个支持增量操作的栈](https://leetcode-cn.com/problems/design-a-stack-with-increment-operation/)
标签：`栈` `设计`

思路：使用一个数组加一个指针来实现。
内存消耗怎么都提不上去……

官网运行结果记录
执行用时：28ms(88%)
内存消耗：7.2MB(9.3%)

*/

type CustomStack struct {
	data []int
	p    int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{data: make([]int, maxSize), p: -1}
}

func (this *CustomStack) Push(x int) {
	if len(this.data) == this.p+1 {
		return
	}
	this.p++
	this.data[this.p] = x
}

func (this *CustomStack) Pop() int {
	res := -1
	if this.p >= 0 {
		res = this.data[this.p]
		this.p--
	}
	return res
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < k && i <= this.p; i++ {
		this.data[i] += val
	}
}
