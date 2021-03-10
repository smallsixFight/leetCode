package implement_queue_using_stacks

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[用栈实现队列](https://leetcode-cn.com/problems/implement-queue-using-stacks/)
标签：`栈` `设计`

简单思路：这道题与昨天做的 `用队列实现栈` 思路基本一致，一个栈（s1）用来作为输出栈，另一个栈（s2）作为输入栈。当 push 时，元素 push 到 s2，当 pop 时，从 s1 pop，若 s1 为空，则将 s2 的元素全部 pop 到 s2 中，再从 s2 进行 push。

由于不是每次 pop 和 peek 都会进行 updateOutQueue，每个元素最多只会 push 和 pop 两次，所以是均摊 O(1)，而 push 和 empty 的复杂度就是 O(1)。

官网运行结果记录
执行用时：0ms
内存消耗：1.9MB/2MB

*/

type MyQueue struct {
	s1 []int
	s2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s2 = append(this.s2, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.s1) == 0 {
		this.updateOutQueue()
	}
	v := this.s1[len(this.s1)-1]
	this.s1 = this.s1[:len(this.s1)-1]
	return v
}

func (this *MyQueue) updateOutQueue() {
	for i := len(this.s2) - 1; i >= 0; i-- {
		this.s1 = append(this.s1, this.s2[i])
	}
	this.s2 = make([]int, 0, len(this.s1))
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.s1) == 0 {
		this.updateOutQueue()
	}
	return this.s1[len(this.s1)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0 && len(this.s2) == 0
}
