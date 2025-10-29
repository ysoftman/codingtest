/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
309. Best Time to Buy and Sell Stock with Cooldown
Medium
You are given an array prices where prices[i] is the price of a given stock on the ith day.

Find the maximum profit you can achieve. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times) with the following restrictions:

After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).
Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

Example 1:
Input: prices = [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]

Example 2:
Input: prices = [1]
Output: 0

Constraints:
1 <= prices.length <= 5000
0 <= prices[i] <= 1000
*/
package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dynamic programming
/*
s0,1,2 상태
s0 ---buy---> s1 ---sell---> s2 ---rest(cooldown 쉬기)---> s3

자세한 설명 참고
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/discuss/75928/Share-my-DP-solution-(By-State-Machine-Thinking)
*/
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	s0 := make([]int, len(prices))
	s1 := make([]int, len(prices))
	s2 := make([]int, len(prices))
	// 처음 시작시 쉬기(reset) 0
	s0[0] = 0
	// 첫날 가격에 buy 산가격 만큼 손해
	s1[0] = -prices[0]
	// sell 한 s2 가장 작은 값으로 시작
	s2[0] = -1 << 31
	for i := 1; i < len(prices); i++ {
		s0[i] = max(s0[i-1], s2[i-1])
		s1[i] = max(s1[i-1], s0[i-1]-prices[i])
		s2[i] = s1[i-1] + prices[i]
	}
	return max(s0[len(prices)-1], s2[len(prices)-1])
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1}))
}
