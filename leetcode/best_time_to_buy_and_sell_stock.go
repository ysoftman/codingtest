/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

121. Best Time to Buy and Sell Stock

You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
*/
package main

import "fmt"

/*
y축은 산날, x축은 판날일때 dynamic programming
   7  1  5  3  6  4
7  0 -6 -2 -4 -1 -3
1     0  4  2  5  3
5        0 -2  1 -1
3           0  3  1
6              0 -2
4                 0

하지만 time limit exceeded
*/
func maxProfit2(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if max < prices[j]-prices[i] {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}

/*
 dynamic programming
 이익 = max(이익, 현재가격-이전 낮은 가격)
 이전낮은 가격 = min(현재가격, 이전 낮은 가격)

 max(profit,7-7)  max(profit,1-7)  max(profit, 5-1)  max(profit, 3-1)  max(profit, 6-1)  max(profit, 4-1)
 min(7, 7)        min(1,7)         min(5, 1)         min(3,1)          min(6,1)          min(4,1)
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfit(prices []int) int {
	profit := 0
	previousValue := prices[0]
	for i := 0; i < len(prices); i++ {
		profit = max(profit, prices[i]-previousValue)
		previousValue = min(prices[i], previousValue)
	}
	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}
