package trapping_rain_water

import "testing"

func TestTrap(t *testing.T) {
	t.Log(Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	t.Log(Trap([]int{4, 1, 0, 0, 3, 2, 3}))
}
