package contains_duplicate_iii

import "testing"

func Test_containsNearbyAlmostDuplicate2(t *testing.T) {
	t.Log(containsNearbyAlmostDuplicate2([]int{1, 2, 3, 1}, 3, 0))       // true
	t.Log(containsNearbyAlmostDuplicate2([]int{1, 0, 1, 1}, 1, 2))       // true
	t.Log(containsNearbyAlmostDuplicate2([]int{1, 5, 9, 1, 5, 9}, 2, 3)) // false
	t.Log(containsNearbyAlmostDuplicate2([]int{-1, -1}, 1, 0))           // true
}

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	t.Log(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))       // true
	t.Log(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))       // true
	t.Log(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3)) // false
	t.Log(containsNearbyAlmostDuplicate([]int{-1, -1}, 1, 0))           // true
}
