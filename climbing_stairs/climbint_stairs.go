package climbing_stairs

func ClimbStairs2(n int) int {
	if n < 3 {
		return n
	}
	a, b := 1, 2
	for i := 3; i < n; i++ {
		a, b = b, a+b
	}
	return a + b
}

func ClimbStairs(n int) int {
	if n < 3 {
		return n
	}
	arr := make([]int, n)
	arr[0] = 1
	arr[1] = 2
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n-1]
}
