package basic_calculator_iv

import (
	"fmt"
	"strconv"
)

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[基本计算器 IV](https://leetcode-cn.com/problems/basic-calculator-iv/)
标签：`栈` `哈希表` `字符串`

思路：这道题...拆，都可以拆，分成下面的处理步骤：

- 以 `evalvars` 和 `evalints` 生成一个 hashTable；
- 把原来的表达式的空格和数字进行处理，生成一个新的表达式；
- 接着按照 `基本计数器 II` 的思路进行乘法分配律和加法处理生成需要整合的子表达式列表；
- 排序生成结果返回。

官网运行结果记录
执行用时：ms
内存消耗：MB

*/

// deprecated
func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {
	ht := make(map[string]int, len(evalvars))
	for i := range evalvars {
		ht[evalvars[i]] = evalints[i]
	}
	exp := generateExp([]byte(expression), ht)
	if len(exp) == 0 {
		return nil
	}
	s1, s2 := [][]byte{{}}, make([]byte, 0)
	for i := 0; i < len(exp); i++ {
		if exp[i] == '(' {
			if exp[i+1] == '(' {
				s2 = append(s2, exp[i])
			}
			continue
		} else if exp[i] == ')' {
			//for exp[i-1] == ')' && len(s2) > 0 && s2[len(s2)-1] != '(' && exp[i] != '*' {
			//	s1[len(s1)-2] = cal(s1[len(s1)-2], s1[len(s1)-1], s2[len(s2)-1])
			//	s1 = s1[:len(s1)-1]
			//	s2 = s2[:len(s2)-1]
			//}
			s1 = append(s1, []byte{})
		} else if (exp[i] == '+' || exp[i] == '-' || exp[i] == '*') && exp[i+1] == '(' {
			//if len(s2) > 0 && exp[i] != '*' {
			//	s1[len(s1)-2] = cal(s1[len(s1)-2], s1[len(s1)-1], s2[len(s2)-1])
			//	s1 = s1[:len(s1)-1]
			//	s2 = s2[:len(s2)-1]
			//}
			s2 = append(s2, exp[i])
		} else {
			s1[len(s1)-1] = append(s1[len(s1)-1], exp[i])
		}
	}
	for i := range s1 {
		fmt.Print(string(s1[i]) + " ")
	}
	fmt.Println(string(s2))
	return nil
}

// deprecated
//func cal(s1, s2 []byte, r byte) []byte {
//	s1, num1 := getExp(s1)
//	s2, num2 := getExp(s2)
//	res := make([]byte, 0)
//	switch r {
//	case '*':
//
//	case '+':
//		res = append(res, s1...)
//		res = append(res, '+')
//		res = append(res, s2...)
//		if v := num1 + num2; v != 0 {
//			if v > 0 {
//				res = append(res, '+')
//			}
//			res = append(res, strconv.FormatInt(int64(v), 10)...)
//		}
//	case '-':
//		res = append(res, s1...)
//		res = append(res, '-')
//		for i := range s2 {
//			if s2[i] == '-' {
//				if res[len(res)-1] == '-' {
//					res[len(res)-1] = '+'
//				} else {
//					res = append(res, '+')
//				}
//			} else if s2[i] == '+' {
//				res = append(res, '-')
//			} else {
//				res = append(res, s2[i])
//			}
//		}
//		if v := num1 - num2; v != 0 {
//			if v > 0 {
//				res = append(res, '+')
//			}
//			res = append(res, strconv.FormatInt(int64(v), 10)...)
//		}
//	}
//	return res
//}

func getExp(s []byte) (exp []byte, num int) {
	exp, direct := make([]byte, 0, len(s)), 1
	for i := 0; i < len(s); i++ {
		if isNum(s[i]) && !((len(exp) > 0 && exp[len(exp)-1] == '*') || (i+1 < len(s) && s[i+1] == '*')) {
			n := int(s[i] - '0')
			i++
			if exp[len(exp)-1] == '-' {
				direct = -1
			}
			for ; i < len(s) && isNum(s[i]); i++ {
				n = n*10 + int(s[i]-'0')
			}
			i--
			num += n * direct
			exp = exp[:len(exp)-1]
			direct = 1
		} else {
			exp = append(exp, s[i])
		}
	}
	return
}

func isNum(b byte) bool {
	return '0' <= b && b <= '9'
}

func generateExp(pre []byte, m map[string]int) []byte {
	res := make([]byte, 0, len(pre))
	num := 0
	for i := 0; i < len(pre); i++ {
		if pre[i] == ' ' {
			continue
		}
		if 'a' <= pre[i] && pre[i] <= 'z' {
			temp := []byte{pre[i]}
			j := i + 1
			for ; j < len(pre) && 'a' <= pre[j] && pre[j] <= 'z'; j++ {
				temp = append(temp, pre[j])
			}
			i = j - 1
			if v, ok := m[string(temp)]; ok {
				// 是否为系数
				isCoef := (len(res) > 0 && res[len(res)-1] == '*') || (j+1 < len(pre) && pre[j+1] == '*')
				if isCoef {
					res = append(res, strconv.FormatInt(int64(v), 10)...)
				} else {
					if len(res) > 0 {
						res = res[:len(res)-1]
					}
					num += v
				}
			} else {
				res = append(res, temp...)
			}
		} else if isNum(pre[i]) {
			v, j := int(pre[i]-'0'), i+1
			for ; j < len(pre) && isNum(pre[j]); j++ {
				v = v*10 + int(pre[j]-'0')
			}
			isCoef := (len(res) > 0 && res[len(res)-1] == '*') || (j+1 < len(pre) && pre[j+1] == '*')
			if isCoef {
				res = append(res, strconv.FormatInt(int64(v), 10)...)
			} else {
				if len(res) > 0 {
					if res[len(res)-1] == '-' {
						v = -v
					}
					res = res[:len(res)-1]
				}
				num += v
			}
			i = j - 1
		} else {
			if (pre[i] == '(' || pre[i] == ')') && num != 0 {
				if num > 0 {
					res = append(res, '+')
				}
				res = append(res, strconv.FormatInt(int64(num), 10)...)
				num = 0
			}
			res = append(res, pre[i])
		}
	}
	if num != 0 {
		if num > 0 {
			res = append(res, '+')
		}
		res = append(res, strconv.FormatInt(int64(num), 10)...)
	}
	if len(res) > 0 && res[0] == '+' {
		res = res[1:]
	}
	fmt.Println(string(res))
	return res
}
