package sort_list

import (
	"sort"
	"strconv"
)

/*
164
来源：[leetCode](https://leetcode-cn.com/)
题目：[最大数](https://leetcode-cn.com/problems/largest-number/)
标签：`排序`

简单思路：先把 `nums` 转换为对应的字符串数组，然后使用归并排序，最后返回归并排序后的字符串数组生成紫的字符串。

需要考虑处理前缀为 `0` 的字符。

largestNumber2 的比较方法值得学习。

官网运行结果记录
执行用时：4ms(largestNumber)/0ms(largestNumber2)
内存消耗：2.7MB/2.8MB(largestNumber)/2.3MB(largestNumber2)

*/

func largestNumber2(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := make([]byte, 0)
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}

func largestNumber(nums []int) string {
	arr := make([]string, len(nums))
	for i := range nums {
		temp := make([]byte, 0)
		for nums[i] > 0 {
			temp = append(temp, byte(nums[i]%10)+'0')
			nums[i] /= 10
		}
		if len(temp) == 0 {
			arr[i] = "0"
		} else {
			arr[i] = string(reverse(temp))
		}
	}
	bs := make([]byte, 0)
	arr = mergeSort(arr)
	for i := range arr {
		bs = append(bs, arr[i]...)
	}
	i := 0
	for ; i < len(bs)-1 && bs[i] == '0'; i++ {
	}
	return string(bs[i:])
}

func reverse(bs []byte) []byte {
	i, j := 0, len(bs)-1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return bs
}

func mergeSort(arr []string) []string {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	arr1 := mergeSort(arr[:mid])
	arr2 := mergeSort(arr[mid:])
	i, j := 0, 0
	res := make([]string, len(arr))
	for idx := 0; i < len(arr1) || j < len(arr2); idx++ {
		if i == len(arr1) {
			res[idx] = arr2[j]
			j++
		} else if j == len(arr2) {
			res[idx] = arr1[i]
			i++
		} else if compare(arr1[i], arr2[j]) {
			res[idx] = arr1[i]
			i++
		} else {
			res[idx] = arr2[j]
			j++
		}
	}
	return res
}

func compare(s1, s2 string) bool {
	for i, j := 0, 0; i < len(s1) || j < len(s2); i++ {
		if i == len(s1) {
			i = -1
			j--
		} else if j == len(s2) {
			j = -1
			i--
		} else if s1[i] < s2[j] {
			return false
		} else if s1[i] > s2[j] {
			return true
		}
		j++
	}
	return true
}
