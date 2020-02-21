### 有效的括号（Valid Parentheses）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/valid-parentheses/)

#### 题目
给定一个只包括`'('，')'，'{'，'}'，'['，']'`的字符串，判断字符串是否有效。

有效字符串需满足：
1. 左括号必须用相同类型的右括号闭合。
2. 左括号必须以正确的顺序闭合。

注意空字符串可被认为是有效字符串。

示例1：
```
输入: "()"
输出: true
```
示例2：
```
输入: "()[]{}"
输出: true
```
示例3：
```
输入: "([)]"
输出: false
```

#### 思路
使用一个栈，将左括号 push，遇到右括号则 pop，查看是否匹配，不匹配或栈为空则返回 FALSE。  
另外，如果栈的大小超过未比较的字符串的一半长度，也直接返回 FALSE。

IsValid() 使用了我自己写的 stack 结构， IsValid2() 使用 slice 来实现类似 stack。

#### 伪代码
```
IS-VALID(s)
    stack = new Stack()
    for i = 0 to s.length -1
        if s[i] == '(' || s[i] == '[' || s[i] == '{'
            stack.push(s[i])
            if stack.length > s[i-(s.length-1)]
                return false
        else
            if stack.isEmpty
                return false
            val = stack.pop
            if s[i] == ')' && val != '('
                return false
            elseif s[i] == ']' && val != '['
                return false
            elseif s[i] == '}' && val != '{'
                return false            
    return stack.isEmpty
```