package number_of_ways_to_paint_N_3_grid

func NumOfWaysTwo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 12
	}
	var temp = 1000000007
	repeat := 6
	unrepeat := 6
	for i := 2; i <= n; i++ {
		newrep := (repeat*3)%temp + unrepeat*2%temp
		newunrep := repeat*2%temp + unrepeat*2%temp
		repeat = newrep
		unrepeat = newunrep
	}
	return (repeat + unrepeat) % temp
}

func NumOfWays(n int) int {
	if n == 1 {
		return 12
	} else if n == 2 {
		return 54
	}
	pre := []int{30, 24}
	moVal := 1000000007
	for i := 3; i <= n; i++ {
		p1 := pre[0]*3%moVal + pre[1]*2%moVal
		p2 := pre[0]*2%moVal + pre[1]*2%moVal
		pre[0] = p1
		pre[1] = p2
	}
	return (pre[0] + pre[1]) % moVal
}
