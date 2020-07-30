package guan_mai

/*
题目：一个芬兰人进了一个房间，房间有一排椅子，椅子上有一些人坐着，还剩一些空位，他要选择一个位子坐下，
这个位子要尽可能远离已经坐着的人，请给出算法。请自行定义数据结构和输入输出。

题目分析：一排椅子为一个数组，已经坐着人的位置上值为 1，没坐人的位置值为 0。
找一个极可能远离已经坐着的人的位置：找的位置与相邻的人的距离越大越好。

思路：使用快慢双指针法，从起点开始，分别找出起点到第一个有人的位置的距离、两个已经有人的位置最大距离、最后一个有人的位置到最后一把椅子的距离。
取其中的最大值，两个已有人的位置的最大距离要取中间的位置。
*/

func findSeat(seats []int) int {
	res := -1
	sp, fp := -1, 0
	hl, ml, tl := -1, -1, -1
	prep := sp
	for fp < len(seats) {
		if seats[fp] == 1 {
			if sp == -1 {
				hl = fp - sp
			} else if fp-sp > ml {
				ml = fp - sp
			}
			prep = sp
			if fp != len(seats)-1 {
				sp = fp
			}
		} else {
			if fp == len(seats)-1 {
				tl = fp - sp
			}
		}
		fp++
	}
	ml = ml / 2
	if hl > ml && hl >= tl {
		res = 0
	} else if tl >= ml {
		res = len(seats) - 1
	} else {
		res = prep + ml
	}
	return res
}
