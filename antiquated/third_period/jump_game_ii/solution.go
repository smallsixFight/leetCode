package jump_game_ii

func Jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	res := 1
	curIdx := 0
	for nums[curIdx] > 0 {
		if curIdx+nums[curIdx] >= len(nums)-1 {
			return res
		}
		res++
		p1 := curIdx
		furthest := 0
		for i := nums[p1]; i > 0; i-- {
			if p1+i+nums[p1+i] > furthest {
				furthest = p1 + i + nums[p1+i]
				curIdx = p1 + i
			}
		}
	}
	return -1
}
