package contains_duplicate_iii

import (
	"sort"
)

/*
220
来源：[leetCode](https://leetcode-cn.com/)
题目：[存在重复元素 III](https://leetcode-cn.com/problems/contains-duplicate-iii/)
标签：`排序` `ordered map`

简单思路：遍历 nums，使用一个 hashMap 记录出现的数字和索引位置，使用 arr 数组记录不重复的数字。
对 arr 进行排序，先比较 `arr[i]` 和 `arr[j]`，符合 `abs[arr[j]-arr[i] <= t`时，再通过 hashMap 比较是否存在 `abs[i-j] <= k`，存在直接返回 true，否则持续比较到结束。


官网运行结果记录
执行用时：44ms/48ms(containsNearbyAlmostDuplicate)/460ms(containsNearbyAlmostDuplicate2)
内存消耗：8.7MB/8.9MB(containsNearbyAlmostDuplicate)/5.9MB(containsNearbyAlmostDuplicate2)

*/

func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	for i := range nums {
		for j := i + 1; j < len(nums) && j <= i+k; j++ {
			if abs(nums[j]-nums[i]) <= t {
				return true
			}
		}
	}
	return false
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	m, arr := make(map[int][]int), make([]int, 0)
	for i := range nums {
		if _, ok := m[nums[i]]; !ok {
			arr = append(arr, nums[i])
		}
		m[nums[i]] = append(m[nums[i]], i)
	}
	sort.Ints(arr)
	//fmt.Println(arr)
	//fmt.Println(m)
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if abs(arr[j]-arr[i]) > t { // 后续都不需要比较
				break // break 改成 continue 直接从 328ms 变成 44ms，但内存占用从 8MB 变为 10.5MB
			}
			l1, l2 := m[arr[i]], m[arr[j]]
			for x, y := 0, 0; x < len(l1) && y < len(l2); {
				if (i != j || x != y) && abs(l1[x]-l2[y]) <= k {
					return true
				}
				if l1[x] > l2[y] {
					y++
				} else {
					x++
				}
				if i == j && x == y { // 加了后，44ms 变成 36ms，内存占用从 10.5MB 变为 8.7MB
					y++
				}
			}
		}

	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
