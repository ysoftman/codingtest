/*
https://leetcode.com/problems/coin-change-ii/
518. Coin Change II
Medium
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the number of combinations that make up that amount. If that amount of money cannot be made up by any combination of the coins, return 0.

You may assume that you have an infinite number of each kind of coin.
The answer is guaranteed to fit into a signed 32-bit integer.

Example 1:
Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

Example 2:
Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.

Example 3:
Input: amount = 10, coins = [10]
Output: 1

Constraints:
1 <= coins.length <= 300
1 <= coins[i] <= 5000
All the values of coins are unique.
0 <= amount <= 5000
*/
package main

import "fmt"

/*
amount = 5, coins = [1,2,5]
0 = 0, 1개
1 = 1, 1개
2 = 1+1, 2, 2개
3 = 1+1+1, 1+1+2, 2개
4 = 1+1+1+1, 1+1+2, 2+2, 3개
5 = 1+1+1+1+1, 1+1+1+2, 1+2+2, 5, 4개

dp[ 1 ] =  1
dp[ 2 ] =  1
dp[ 3 ] =  1
dp[ 4 ] =  1
dp[ 5 ] =  1

dp[ 2 ] = dp[2] + dp[2-2] = 2
dp[ 3 ] = dp[3] + dp[3-2] = 2
dp[ 4 ] = dp[4] + dp[4-2] = 3
dp[ 5 ] = dp[5] + dp[5-2] = 3

dp[ 5 ] = dp[5] + dp[5-5] = 4
*/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
			// fmt.Println("dp[",j,"] = ", dp[j])
		}
	}
	return dp[amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Println(change(3, []int{2}))
	fmt.Println(change(10, []int{10}))
}
