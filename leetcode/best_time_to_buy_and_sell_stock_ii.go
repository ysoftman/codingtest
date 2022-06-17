/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
22. Best Time to Buy and Sell Stock II
Medium
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.
Find and return the maximum profit you can achieve.

Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.

Example 2:
Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.

Example 3:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.

Constraints:
1 <= prices.length <= 3 * 104
0 <= prices[i] <= 104
*/
package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
7,1,5,3,6,4

buy = 샀을때 가장 큰 이익
sell = 팔았을때 가장 큰 이익

사면 같은날 바로 팔아야 하기때문에, buy 계산후 바로 sell 을 계산하다.

buy = max(buy, sell-현재가), 사면 수익이 현재가 만큼 줄어드니까 -현재가
sell = max(sell, buy+현재가), 팔면 수익이 현재가 만큰 늘어나니까 +현재가

1day 가격 7
buy = max(-7, 0-7) => -7
sell = max(0, -7+0) => 0

2day 가격 1
buy = max(-7, 0-1) => -1
sell = max(0, -1+1) => 0

3day 가격 5
buy = max(-1, 0-5) => -1
sell = max(0, -1+5) => 4

4day 가격 3
buy = max(-1, 4-3) => 1
sell = max(4, 1+3) => 4

5day 가격 6
buy = max(1, 4-6) => 1
sell = max(4, 1+6) => 7

6day 가격 4
buy = max(1, 7-4) => 3
sell = max(7, 3+4) => 7
*/
func maxProfit(prices []int) int {
	buy := -prices[0]
	sell := 0
	for i := 0; i < len(prices); i++ {
		buy = max(buy, sell-prices[i])
		sell = max(sell, buy+prices[i])
		// fmt.Printf("buy:%v, sell:%v\n", buy, sell)
	}
	return sell
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
