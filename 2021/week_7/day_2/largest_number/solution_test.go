package sort_list

import "testing"

func Test_largestNumber(t *testing.T) {
	//t.Log(largestNumber([]int{10, 2}), "210")                 // 210
	//t.Log(largestNumber([]int{3, 30, 34, 5, 9}), "9534330")   // 9534330
	//t.Log(largestNumber([]int{1}), "1")                       // 1
	//t.Log(largestNumber([]int{10}), "10")                     // 10
	//t.Log(largestNumber([]int{34323, 3432}), "343234323")     // 343234323
	//t.Log(largestNumber([]int{10, 2, 9, 39, 17}), "93921710") // 93921710
	t.Log(largestNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
	t.Log(largestNumber([]int{0, 0}))
}

func Test_compare(t *testing.T) {
	t.Log(compare("30", "3"))
	t.Log(compare("3", "30"))
	t.Log(compare("33", "3"))
	t.Log(compare("33334", "3"))
	t.Log(compare("3", "33334"))
	t.Log(compare("34", "3"))
	t.Log(compare("33", "34"))
	t.Log(compare("3", "34"))
	t.Log(compare("34323", "3432"))
	t.Log(compare("2", "17"))
	t.Log(compare("17", "2"))
	t.Log(compare("1", "2"))
}
