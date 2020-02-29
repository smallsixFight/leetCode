package four_sum

import "sort"

func FourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			l, r := j+1, len(nums)-1
			if nums[i]+nums[j] > target && nums[j] > 0 {
				break
			}
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l+1] == nums[l] {
						l++
					}
					for l < r && nums[r-1] == nums[r] {
						r--
					}
					l++
					r--
				} else if sum > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return res
}
