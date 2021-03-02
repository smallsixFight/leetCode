### 买卖股票的最佳时机（Best Time to Buy and Sell Stock）
[原题来源：力扣（LeetCode）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)

#### 题目
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

注意：你不能在买入股票前卖出股票。

示例 1：
```
输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
```
示例 2：
```
输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
```

#### 思路一（暴力法）
这道题可以将前`n-1`个作为买股票的那天，然后遍历后面是否存在利润大于当前`maxProfit`的利润，如果有，则替换。

当然，不必每个点都进行遍历，可以把之前可以买入股票最低价格的那天记录起来，如果当前的价格大于之前的购买价格并且在卖出的天数之前，则可以直接跳过。

#### 思路二（DP）
这道题可以用 DP 来做。大致思路如下：

第`n`天卖出能够获取的最大利润是前`n-1`天能够能够买入价格最低的那一天的价格差，所以可以以此构建一个子结构。

#### 吐槽
额...我认为这道题跟股票没任何关系。

#### 伪代码
```
// DP
MAX-PROFIT(prices)
    maxProfit = 0
    minDay = 0
    for i = 1 to prices.length -1
        if prices[i] - prices[minDay] > maxProfit
            maxProfit = prices[i] - prices[minDay]
        if prices[i] < prices[minDay]
            minDay = i
    return maxProfit
    
// 暴力破解
MAX-PROFIT(prices)
    maxProfit = 0
    buyDay, sellDay = 0, 0
    for i = 0 to prices.length -2
        if prices[i] >= prices[buyDay] && i < sellDay
            continue
        if prices[i] < prices[buyDay] && i < sellDay
            buyDay = i
            maxProfit = prices[sellDay] - prices[buyDay]
            continue
        for k = i +1 to prices.length -1
            if prices[k] - prices[i] > maxProfit
                buyDay = i
                sellDay = k
                maxProfit = prices[k] - prices[i]
    return maxProfit
```