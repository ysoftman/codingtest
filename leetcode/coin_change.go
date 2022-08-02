/*
https://leetcode.com/problems/coin-change/
322. Coin Change
Medium
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.
Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.
You may assume that you have an infinite number of each kind of coin.

Example 1:
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:
Input: coins = [2], amount = 3
Output: -1

Example 3:
Input: coins = [1], amount = 0
Output: 0

// 점화식
1,2,5

0 => 0

1 => 1(1)
1

2 => 1(2)
1,1
2

3 => 2(1,2)
1,1,1
1,2

4 => 2(2,2)
1,1,1,1
1,1,2
2,2

5 => 1(5)
1,1,1,1,1
1,1,1,2
1,2,2
5

6 => 2(1,5)
1,1,1,1,1,1
1,1,1,1,2
1,1,2,2
2,2,2
1,5

7 => 2(1,5)
1,1,1,1,1,1,1
1,1,1,1,1,2
1,1,1,2,2
1,2,2,2
1,1,5
2,5


amount=-1 이하
==> invaild

amount=0 인경우
0

amount=1 인경우
1-1원 => 0원 만드는 최소개수(0) + 1(1원)
중 최소 => 1

amount=4 인경우
4-1원 => 3원 만드는 최소개수(2) + 1(1원)
4-2원 => 2원 만드는 최소개수(1) + 1(2원)
중 최소 => 2

amount=5 인경우
5-1원 => 4원 만드는 최소개수(2) + 1(1원)
5-2원 => 3원 만드는 최소개수(2) + 1(2원)
5-5원 => 0원 만드는 최소개수(0) + 1(5원)
중 최소 => 1

amount=6 인경우
6-1원 => 5원 만드는 최소개수(1) + 1(1원)
6-2원 => 4원 만드는 최소개수(2) + 1(2원)
6-5원 => 1원 만드는 최소개수(1) + 1(5원)
중 최소 => 2

amount=7 인경우
7-1원 => 6원 만드는 최소개수(2) + 1(1원)
7-2원 => 5원 만드는 최소개수(1) + 1(2원)
7-5원 => 2원 만드는 최소개수(1) + 1(5원)
중 최소 => 2


dp[7] = min(dp[7-1], dp[7-2], dp[7-5]) + 1

bottom-up 방식으로 채워나가자
0 1 2 3 4 5 6 7 => amount
0 1 1 2 2 1 2 2 => min

*/
package main

import "fmt"

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		min := 1<<31 - 1
		for _, coin := range coins {
			if i-coin >= 0 && min >= dp[i-coin] {
				min = dp[i-coin]
			}
		}
		dp[i] = min + 1
	}
	if dp[amount] == (1<<31-1)+1 {
		return -1
	}
	return dp[amount]
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 5))
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{2}, 4))
}
