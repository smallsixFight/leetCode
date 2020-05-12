### 环形链表（Linked List Cycle）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/linked-list-cycle/)

#### 题目
设计一个支持`push`，`pop`，`top`操作，并能在常数时间内检索到最小元素的栈。

- `push(x)` —— 将元素 x 推入栈中。
- `pop()` —— 删除栈顶的元素。
- `top()` —— 获取栈顶元素。
- `getMin()` —— 检索栈中的最小元素。

示例：
```
输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
```

提示：

- `pop`、`top`和`getMin`操作总是在 **非空栈** 上调用。

#### 思路一
这道题的考点在于需要在常数时间内检索到最小的元素，也就是要求对入栈的元素在内部进行排序。

其实我想到两种方式，第一种是对入栈的元素进行排序，然后单独维护好入栈的顺序；第二种是保持入栈的顺序，然后单独维护排序。

这里我选择第二种。

#### 思路二
然而使用链表的导致 Pop() 和 Push() 的时间复杂度都是 O(n)，然后我去评论学习了，有了以下的思路。

由于 getMin() 不需要删除该元素，所以每个入栈的元素只需要保存两个数值，一个是入栈的值，一个是截至当前值入栈后最小的值即可。

#### 伪代码