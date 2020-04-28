package find_peak_element

import "testing"

func TestFindPeakElement(t *testing.T) {
	t.Log(FindPeakElement([]int{1, 2, 3, 1}))
	t.Log(FindPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	t.Log(FindPeakElement([]int{1, 2}))
	t.Log(FindPeakElement([]int{2, 1}))
}
