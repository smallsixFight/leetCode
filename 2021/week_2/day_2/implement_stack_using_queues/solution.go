package implement_stack_using_queues

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[用队列实现栈](https://leetcode-cn.com/problems/implement-stack-using-queues/)
标签：`栈` `设计`

简单思路：一个队列（q1）作为输出队列，一个队列（q2）作为输入队列，在 push 时先将值放入 q2，然后将 q1 的所有值依次入队到 q2，这样就维持了栈的结构了，然后将 q1 和 q2 作交换。pop 和 top 均取 a1 的第一个值就可以。

单个队列实现：每次 push 的时候，只要记录队列的原长度，在添加新元素后，将新元素前面的元素依次出队后重新入队就可以了。

官网运行结果记录
执行用时：0ms
内存消耗：1.9MB

*/

type MyStack struct {
	queue1 []int
	queue2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue2 = append(this.queue2, x)
	this.queue2 = append(this.queue2, this.queue1...)
	this.queue1 = make([]int, 0, len(this.queue2))
	this.queue1, this.queue2 = this.queue2, this.queue1
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	v := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return v
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue1[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}
