package online_stock_span

/*
来源：[leetCode](https://leetcode-cn.com/)
题目：[股票价格跨度](https://leetcode-cn.com/problems/online-stock-span/)
标签：`栈`

简单思路：这道题稍微反向想一下，就是向前遍历，找到第一个大于当前 price 的那一天，然后求时间差。接着，根据题目的难度和提示，知道不能使用暴力遍历的方法来处理，肯定超时。
接着，观察下题目给的例子——[100, 80, 60, 70, 60, 75, 85]，以 75 为例，它的前一个更大值是 80，可以发现，它的前面有个 70，而因为 `75>70` & `70 > 60`，那么 75 其实并不需要跟 60 比较大小，肯定包括 60 的股价那天，那么把 60 移除后，可以得到 [100, 80, 70, 60, 75, 85]，那么就可以发现这其实就是一个单调递减栈。
另外，由于需要计算间隔天数，可以使用一个二位数组作为栈元素的结构。

官网运行结果记录
执行用时：180ms(98%)/188ms(84%)
内存消耗：8.4MB(100%)/8.6MB(91%)

*/

type StockSpanner struct {
	stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{stack: make([][2]int, 0)}
}

func (this *StockSpanner) Next(price int) int {
	days := 1
	for len(this.stack) != 0 && this.stack[len(this.stack)-1][0] <= price {
		days += this.stack[len(this.stack)-1][1]
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, [2]int{price, days})
	return days
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
