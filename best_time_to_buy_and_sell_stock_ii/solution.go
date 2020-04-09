package best_time_to_buy_and_sell_stock_ii

func MaxProfit(prices []int) int {
	maxPro, buyD := 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[buyD] {
			buyD = i
		} else if i == len(prices)-1 || prices[i+1] < prices[i] {
			if prices[i]-prices[buyD] > 0 {
				maxPro += prices[i] - prices[buyD]
			}
			buyD = i + 1
		}
	}
	return maxPro
}
