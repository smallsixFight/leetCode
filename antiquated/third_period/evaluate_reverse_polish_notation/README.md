###  逆波兰表达式求值（Evaluate Reverse Polish Notation）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)

#### 题目
根据逆波兰表示法，求表达式的值。

有效的运算符包括`+`，`-`，`*`，`/`。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

说明：

- 整数除法只保留整数部分。
- 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

示例 1：
```
输入: ["2", "1", "+", "3", "*"]
输出: 9
解释: ((2 + 1) * 3) = 9
```
示例 2：
```
输入: ["4", "13", "5", "/", "+"]
输出: 6
解释: (4 + (13 / 5)) = 6
```
示例 3：
```
输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
输出: 22
解释: 
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
```

#### 思路
逆波兰其实就是对栈的应用，如果是数字则入栈，如果是运算符号，则取出栈顶的两个元素，把运算后的结果入栈，最后返回栈底的结果。

#### 伪代码
```
ENV-RPN(token)
    stack = new Stack()
    for i = 0 to token.length -1
        switch token[i]:
        case: "*":
            v1 = stack.pop()
            v2 = stack.pop()
            stack.push(v1 * v2)
        case: "/":
            v1 = stack.pop()
            v2 = stack.pop()
            stack.push(v2 / v1)
        case: "+":
            v1 = stack.pop()
            v2 = stack.pop()
            stack.push(v2 + v1)
        case: "-":
            v1 = stack.pop()
            v2 = stack.pop()
            stack.push(v2 - v1)
        default:
            isNe = false
            num = 0
            for j = 0 to token[i].length -1
                if token[i][j] == '-'
                    isNe = true
                else if token[i][j] >= '0' && token[i][j] <= '9'
                    num = num * 10 + token[i][j] - '0'
            if isNe
                stack.push(-num)
            else
                stack.push(num)
    return stack.pop()
```