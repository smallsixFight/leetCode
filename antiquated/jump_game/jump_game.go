package jump_game

func CanJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	curPoint := 0
	for nums[curPoint] > 0 {
		if curPoint+nums[curPoint] >= len(nums)-1 {
			return true
		}
		p1 := curPoint
		temp := 0
		for i := nums[p1]; i > 0; i-- {
			if i+p1 >= len(nums)-1 {
				return true
			}
			if p1+i+nums[p1+i] > temp {
				temp = p1 + i + nums[p1+i]
				curPoint = p1 + i
			}
		}
	}
	return false
}
