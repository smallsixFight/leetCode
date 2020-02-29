package three_sum_closest

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var res int
	if target > 0 {
		res = math.MinInt32 / 10
	} else {
		res = math.MaxInt32 / 10
	}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return target
			}
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum < target {
				for l < r && nums[l+1] == nums[l] {
					l++
				}
				l++
			} else {
				for l < r && nums[r-1] == nums[r] {
					r--
				}
				r--
			}
		}
	}
	return res
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
