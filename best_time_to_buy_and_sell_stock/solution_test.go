package best_time_to_buy_and_sell_stock

import "testing"

func TestMaxProfitTwo(t *testing.T) {
	t.Log(5 == MaxProfitTwo([]int{7, 1, 5, 3, 6, 4}))
	t.Log(0 == MaxProfitTwo([]int{7, 6, 4, 3, 1}))
	t.Log(2 == MaxProfitTwo([]int{2, 1, 2, 1, 0, 1, 2}))
	t.Log(4 == MaxProfitTwo([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	t.Log(0 == MaxProfitTwo([]int{}))
}

func TestMaxProfit(t *testing.T) {
	t.Log(5 == MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	t.Log(0 == MaxProfit([]int{7, 6, 4, 3, 1}))
	t.Log(2 == MaxProfit([]int{2, 1, 2, 1, 0, 1, 2}))
	t.Log(4 == MaxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}
