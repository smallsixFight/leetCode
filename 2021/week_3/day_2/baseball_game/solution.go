package baseball_game

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[棒球比赛](https://leetcode-cn.com/problems/baseball-game/)
标签：`栈`

简单思路：这道题有着很简单的思路，用栈来处理，或者用双向列表处理都行。

官网运行结果记录
执行用时：0ms
内存消耗：2.7MB

*/

func calPoints(ops []string) int {
	stack := make([]int, 0)
	for i := range ops {
		switch ops[i][0] {
		case 'C':
			stack = stack[:len(stack)-1]
		case 'D':
			stack = append(stack, stack[len(stack)-1]*2)
		case '+':
			stack = append(stack, stack[len(stack)-2]+stack[len(stack)-1])
		default:
			j, direct := 0, 1
			if ops[i][0] == '-' {
				j++
				direct = -1
			}
			num := int(ops[i][j] - '0')
			j++
			for ; j < len(ops[i]); j++ {
				num = num*10 + int(ops[i][j]-'0')
			}
			stack = append(stack, num*direct)
		}
	}
	res := 0
	for i := len(stack) - 1; i >= 0; i-- {
		res += stack[i]
	}
	return res
}
