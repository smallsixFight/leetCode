package jump_game

import "testing"

func TestCanJump(t *testing.T) {
	t.Log(CanJump([]int{2, 3, 1, 1, 4}))
	t.Log(CanJump([]int{3, 2, 1, 0, 4}))
}
