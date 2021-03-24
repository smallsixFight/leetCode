package odd_even_jump

import (
	"sort"
)

/*
975
来源：[leetCode](https://leetcode-cn.com/)
题目：[奇偶跳](https://leetcode-cn.com/problems/odd-even-jump/)
标签：`栈` `动态规划` `Ordered Map`

简单思路：这道题的很明确，就是先找出数组中每个元素分别进行奇偶跳的下一个位置，然后再从左向右遍历进行跳跃尝试，能跳到最后则记录加一。
然而，这道题的主要问题就是如何友好地解决找到每个元素奇偶跳的下一个位置。一开始，想到了用单调栈，不过在联想到要排序的时候就临时放弃了。
接着想到一个新的思路，在尝试后发现不符合，于是写了一次暴力法，提交后果不其然超时。
然后，去看了下题解，发现居然真的可以用单调栈来处理，于是重新理了一遍，做了出来。

教训：困难类型的题如果有思路就先尽量尝试，不要想到过程复杂或者耗时就放弃，先做出来再说。

官网运行结果记录
执行用时：56ms/68ms
内存消耗：6.7MB

*/

func oddEvenJumps2(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	s := make([]int, len(arr))
	for i := range arr {
		s[i] = i
	}
	sort.Slice(s, func(i, j int) bool {
		if arr[s[i]] == arr[s[j]] {
			return s[i] < s[j]
		}
		return arr[s[i]] < arr[s[j]]
	})
	//fmt.Println(s)

	record := make([][2]int, len(arr))
	getNext := func(idx int) {
		stack := make([]int, 0)
		for i := 0; i < len(arr); i++ {
			for len(stack) > 0 && s[i] > stack[len(stack)-1] {
				record[stack[len(stack)-1]][idx] = s[i]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, s[i])
		}
	}
	getNext(0)

	sort.Slice(s, func(i, j int) bool {
		if arr[s[i]] == arr[s[j]] {
			return s[i] < s[j]
		}
		return arr[s[i]] > arr[s[j]]
	})
	getNext(1)
	//fmt.Println(record)
	res := 0
	for i := range record {
		if i == len(record)-1 {
			res++
			break
		}
		cur := record[i][0]
		idx := 1
		for cur > 0 {
			if cur == len(arr)-1 {
				break
			}
			cur = record[cur][idx]
			if idx == 0 {
				idx = 1
			} else {
				idx = 0
			}
		}
		if cur == len(arr)-1 {
			res++
		}
	}
	return res
}

func oddEvenJumps(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	record := make([][2]int, len(arr))
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] >= arr[i] {
				if record[i][0] == 0 || arr[j] < arr[record[i][0]] {
					record[i][0] = j
				}
			}
			if arr[j] <= arr[i] {
				if record[i][1] == 0 || arr[j] > arr[record[i][1]] {
					record[i][1] = j
				}
			}
		}
	}
	//fmt.Println(record)
	res := 0
	for i := range record {
		if i == len(record)-1 {
			res++
			break
		}
		cur := record[i][0]
		idx := 1
		for cur > 0 {
			if cur == len(arr)-1 {
				break
			}
			cur = record[cur][idx]
			if idx == 0 {
				idx = 1
			} else {
				idx = 0
			}
		}
		if cur == len(arr)-1 {
			res++
		}
	}
	return res
}

/*
for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] == arr[i+1] {
			record[i][0], record[i][1] = i+1, i+1
		} else if arr[i] < arr[i+1] {
			record[i][0] = i+1
			for j := record[i+1][1]; j > 0 && arr[i] <= arr[j]; j = record[j][1] {
				if arr[j] < arr[record[i][0]] {
					record[i][0] = j
				}
				if arr[i] == arr[j] {
					break
				}
			}
			for j := record[i+1][1]; j > 0; j = record[j][1] {
				if arr[i] >= arr[j] {
					record[i][1] = j
					break
				}
			}
		} else {
			record[i][1] = i+1
			for j :=record[i+1][0]; j > 0 && arr[j] <= arr[i]; j = record[j][0] {
				if arr[j] > arr[record[i][1]] {
					record[i][1] = j
				}
				if arr[i] == arr[j] {
					break
				}
			}

			for j := i+1; j > 0; j = record[j][0] {
				if arr[i] <= arr[j] {
					record[i][0] = j
					break
				}
			}
		}
	}
*/
