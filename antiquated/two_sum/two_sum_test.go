package two_sum

import "testing"

func TestFirstMethod(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := FirstMethod(nums, target)
	if result != nil {
		t.Logf("%d + %d = %d", nums[result[0]], nums[result[1]], target)
	} else {
		t.Log("the result is nil")
	}
}

func TestSecondMethod(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := SecondMethod(nums, target)
	if result != nil {
		t.Logf("%d + %d = %d", nums[result[0]], nums[result[1]], target)
	} else {
		t.Log("the result is nil")
	}
}
