package number_of_atoms

import (
	"fmt"
	"sort"
	"strings"
)

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[原子的数量](https://leetcode-cn.com/problems/number-of-atoms/)
标签：`栈`

简单思路：这道题的解题思路与前面做过的 `字符串解码` 几乎一模一样，只是把从左向右遍历改成从右向左遍历，需要处理原子符号的获取，并需要返回按字典序排序的字符串。这里大致写下处理逻辑：

- 使用两个栈，s1 用来存储需要重复的数字，s2 保存需要进行重复处理的化学式；
- 从右向左遍历，遇到原子或者数字加原子的数字，将连续的出现的原子及其数量组成一个 hashTable 放入 s2；
- 遇到数字加右括号，将数字放入 s1，开辟一个新的 hashTable 来放新的化学式原子计数；
- 遇到左括号时，取出 s1、s2 栈顶两个元素相乘，然后重新放入 s2；
- 最后取出 s2 栈顶元素，按字典序排序生成字符串并返回。

官网运行结果记录
执行用时：0ms
内存消耗：2.2MB

*/

func countOfAtoms(formula string) string {
	if len(formula) == 0 {
		return ""
	}
	s1, s2 := make([]int, 0), []map[string]int{{}}
	for i := len(formula) - 1; i >= 0; i-- {
		if formula[i] != '(' {
			num := 1
			if '0' <= formula[i] && formula[i] <= '9' {
				num = int(formula[i] - '0')
				j, nn := i-1, 10
				for j >= 0 && '0' <= formula[j] && formula[j] <= '9' {
					num = int(formula[j]-'0')*nn + num
					nn *= 10
					j--
				}
				i = j
			}
			if formula[i] == ')' {
				s1 = append(s1, num)
				s2 = append(s2, map[string]int{})
			} else {
				f := make([]byte, 0, 2)
				if 'a' <= formula[i] && formula[i] <= 'z' { // 原子符号长度最大两位，并且一定由大小写字母组成
					f = append(f, formula[i-1], formula[i])
					i--
				} else {
					f = append(f, formula[i])
				}
				s2[len(s2)-1][string(f)] += num
			}
		} else {
			num, m := s1[len(s1)-1], s2[len(s2)-1]
			s1, s2 = s1[:len(s1)-1], s2[:len(s2)-1]
			for k := range m {
				s2[len(s2)-1][k] += m[k] * num
			}
		}
	}
	keys := make([]string, 0, len(s2[0]))
	for k := range s2[0] {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	res := strings.Builder{}
	for _, k := range keys {
		if s2[0][k] > 1 {
			res.WriteString(fmt.Sprintf("%s%d", k, s2[0][k]))
		} else {
			res.WriteString(fmt.Sprintf("%s", k))
		}

	}
	return res.String()
}
