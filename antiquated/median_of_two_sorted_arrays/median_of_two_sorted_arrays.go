package median_of_two_sorted_arrays

func FindMedianSortedArrays(nums1, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	iMin, iMax, halfLen := 0, m, (m+n+1)/2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && nums1[i] < nums2[j-1] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			max1 := 0
			if i == 0 {
				max1 = nums2[j-1]
			} else if j == 0 { // m == n
				max1 = nums1[i-1]
			} else {
				max1 = max(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(max1)
			}
			max2 := 0
			if i == m {
				max2 = nums2[j]
			} else if j == n { // i == 0 && m == n
				max2 = nums1[i]
			} else {
				max2 = min(nums1[i], nums2[j])
			}
			return float64(max1+max2) / 2
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
