package best_time_to_buy_and_sell_stock

// DP
func MaxProfitTwo(prices []int) int {
	maxPro, minD := 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[minD] > maxPro {
			maxPro = prices[i] - prices[minD]
		}
		if prices[i] < prices[minD] {
			minD = i
		}
	}
	return maxPro
}

func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxProfit := 0
	buyDay, sellDay := 0, 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] >= prices[buyDay] && i < sellDay {
			continue
		}
		if prices[i] < prices[buyDay] && i < sellDay {
			buyDay = i
			maxProfit = prices[sellDay] - prices[buyDay]
			continue
		}
		for k := i + 1; k < len(prices); k++ {
			if prices[k]-prices[i] > maxProfit {
				buyDay = i
				sellDay = k
				maxProfit = prices[k] - prices[i]
			}
		}
	}
	return maxProfit
}
