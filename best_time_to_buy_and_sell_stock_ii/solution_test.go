package best_time_to_buy_and_sell_stock_ii

import "testing"

func TestMaxProfit(t *testing.T) {
	t.Log(MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	t.Log(MaxProfit([]int{1, 2, 3, 4, 5}))
	t.Log(MaxProfit([]int{7, 6, 4, 3, 1}))
}
