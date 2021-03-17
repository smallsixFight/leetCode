package asteroid_collision

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[行星碰撞](https://leetcode-cn.com/problems/asteroid-collision/)
标签：`栈`

简单思路：根据示例，先给题目里的行星加个二向箔，使用一个栈来保存行星的信息，只有 `asteroids[i-1]` 行星向右移动，`asteroids[i]` 向左移动，才会发生碰撞。大致处理逻辑如下：

- 使用一个栈作记录，从左向右遍历；
- 当 `asteroids[i] `为负数时，判断栈顶行星是否为正数（栈不为空的情况），正数则会发生碰撞，进行碰撞计算，否则加入栈中；
- 遍历结束后返回结果集。

官网运行结果记录
执行用时：4ms/8ms
内存消耗：4.5MB

*/

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))
	for i := range asteroids {
		if asteroids[i] < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			isNeedAdd := true
			for len(stack) > 0 && stack[len(stack)-1] > 0 {
				v := stack[len(stack)-1] + asteroids[i]
				if v == 0 {
					stack = stack[:len(stack)-1]
					isNeedAdd = false
					break
				} else if v < 0 {
					stack = stack[:len(stack)-1]
				} else {
					isNeedAdd = false
					break
				}
			}
			if isNeedAdd {
				stack = append(stack, asteroids[i])
			}
			continue
		}
		stack = append(stack, asteroids[i])
	}
	return stack
}
