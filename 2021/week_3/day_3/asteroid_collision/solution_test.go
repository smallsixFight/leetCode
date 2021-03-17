package asteroid_collision

import "testing"

func Test_asteroidCollision(t *testing.T) {
	t.Log(asteroidCollision([]int{5, 10, -5}))
	t.Log(asteroidCollision([]int{8, -8}))
	t.Log(asteroidCollision([]int{10, 2, -5}))
	t.Log(asteroidCollision([]int{-2, -1, 1, 2}))
	t.Log(asteroidCollision([]int{-2, -2, 1, -2}))
	t.Log(asteroidCollision([]int{-2, -2, 1, -1}))
}
