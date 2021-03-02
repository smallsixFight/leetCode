package climbing_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	t.Log(1 == ClimbStairs(1))
	t.Log(2 == ClimbStairs(2))
	t.Log(3 == ClimbStairs(3))
	t.Log(5 == ClimbStairs(4))
}

func TestClimbStairs2(t *testing.T) {
	t.Log(1 == ClimbStairs2(1))
	t.Log(2 == ClimbStairs2(2))
	t.Log(3 == ClimbStairs2(3))
	t.Log(5 == ClimbStairs2(4))
}
