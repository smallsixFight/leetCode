package container_with_most_water

func MaxArea2(height []int) int {
	p1 := 0
	max := 0
	for p1 < len(height)-1 {
		for p2 := 0; p2 < len(height); p2++ {
			h := 0
			if height[p1] > height[p2] {
				h = height[p2]
			} else {
				h = height[p1]
			}
			area := (p2 - p1) * h
			if area > max {
				max = area
			}
		}
		p1++
	}
	return max
}

func MaxArea(height []int) int {
	max, p1, p2 := 0, 0, len(height)-1
	for p1 < p2 {
		h := 0
		if height[p1] > height[p2] {
			h = height[p2]
			p2--
		} else {
			h = height[p1]
			p1++
		}
		area := h * (p2 - p1 + 1)
		if area > max {
			max = area
		}
	}
	return max
}
