package trapping_rain_water

func Trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	p1, p2 := 0, 0
	res := 0
	for p1+1 < len(height) {
		if height[p1+1] < height[p1] {
			p2 = p1 + 1
			break
		}
		p1++
	}
	for p2 < len(height) {
		for p2 < len(height)-1 && height[p2] < height[p1] {
			p2++
		}
		if height[p2] >= height[p1] {
			for i := p1 + 1; i < p2; i++ {
				res += height[p1] - height[i]
			}
			p1 = p2
		}
		p2++
	}
	p2--
	if p1 != p2 {
		for p2 > p1 && height[p2] <= height[p2-1] {
			p2--
		}
		for i := p2 - 1; i > p1; i-- {
			if height[i] < height[p2] {
				res += height[p2] - height[i]
			} else {
				p2 = i
			}
		}
	}
	return res
}
