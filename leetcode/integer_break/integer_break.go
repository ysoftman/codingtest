/*
https://leetcode.com/problems/integer-break/
343. Integer Break
Medium
Given an integer n, break it into the sum of k positive integers, where k >= 2, and maximize the product of those integers.
Return the maximum product you can get.

Example 1:
Input: n = 2
Output: 1
Explanation: 2 = 1 + 1, 1 × 1 = 1.

Example 2:
Input: n = 10
Output: 36
Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.

Constraints:
2 <= n <= 58
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
0 = 0
1 = 1
2 = 1*1=1
3 = 1*1*1=1, 1*2=2
4 = 1*1*1*1=1, 1*2*1=2, 2*2=4
5 = 1*1*1*1*1=1, 1*2*1*1=2, 1*2*2=4, 2*3=6
6 = 1*1*1*1*1*1=1, 1*1*2*1*1=2, 1*1*2*2=4, 2*3*1=6, 3*3=9
...

dynamic programming

10 = max(
dp[10-1] * dp[1]
...
dp[10-4] * dp[4] = 2*2*3*3 = 36
dp[10-5] * dp[5] = 6*6 = 36
...
dp[10-9] * dp[1]
)
*/
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			temp := max(i-j, dp[i-j]) * max(j, dp[j])
			dp[i] = max(dp[i], temp)
		}
	}
	// fmt.Println(dp)
	return dp[n]
}

func main() {
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(10))
	fmt.Println(integerBreak(25))
	fmt.Println(integerBreak(37))
	fmt.Println(integerBreak(58))
}
