package largest_number

import "strings"

func LargestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	arr := make([][]byte, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			arr = append(arr, []byte{'0'})
			continue
		}
		stack := make([]byte, 0)
		for nums[i] > 0 {
			stack = append(stack, byte(nums[i]%10)+'0')
			nums[i] /= 10
		}
		temp := make([]byte, len(stack))
		a := 0
		for i := len(stack) - 1; i >= 0; i-- {
			temp[a] = stack[i]
			a++
		}
		arr = append(arr, temp)
	}
	arr = customSort(arr)
	res := strings.Builder{}

	for i := 0; i < len(arr); i++ {
		res.Write(arr[i])
	}
	str := res.String()
	st := 0
	for ; st < len(str)-1; st++ {
		if str[st] != '0' {
			break
		}
	}
	return str[st:]
}

func customSort(arr [][]byte) [][]byte {
	if len(arr) < 2 {
		return arr
	}
	arr1, arr2 := arr[:len(arr)/2], arr[len(arr)/2:]
	arr1 = customSort(arr1)
	arr2 = customSort(arr2)
	res := make([][]byte, len(arr))
	for i, j, idx := 0, 0, 0; i < len(arr1) || j < len(arr2); {
		if i == len(arr1) {
			res[idx] = arr2[j]
			j++
		} else if j == len(arr2) {
			res[idx] = arr1[i]
			i++
		} else if string(arr1[i]) == string(arr2[j]) {
			res[idx] = arr1[i]
			idx++
			res[idx] = arr2[j]
			i++
			j++
		} else if arr1[i][0] > arr2[j][0] {
			res[idx] = arr1[i]
			i++
		} else if arr1[i][0] < arr2[j][0] {
			res[idx] = arr2[j]
			j++
		} else {
			if isBigger(arr1[i], arr2[j]) {
				res[idx] = arr1[i]
				i++
			} else {
				res[idx] = arr2[j]
				j++
			}
		}
		idx++
	}
	return res
}

func isBigger(arr1, arr2 []byte) bool {
	if len(arr1) == 0 || len(arr2) == 0 {
		return true
	}
	for i := 0; ; i++ {
		if i == len(arr1) {
			return isBigger(arr1, arr2[i:])
		} else if i == len(arr2) {
			return isBigger(arr1[i:], arr2)
		}
		if arr1[i] > arr2[i] {
			return true
		} else if arr1[i] < arr2[i] {
			return false
		}
	}
}
