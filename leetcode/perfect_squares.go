/*
https://leetcode.com/problems/perfect-squares/
279. Perfect Squares
Medium
Given an integer n, return the least number of perfect square numbers that sum to n.

A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself. For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.

Example 1:
Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.

Example 2:
Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.

Constraints:
1 <= n <= 104
*/
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
1 = 1 => 1
2 = 1+1 => 2
3 = 1+1+1 => 3
4 = min(4, 1+1+1+1) => 1
...
9 = 9 => 1
...
12 = min(4+4+4, 1+1+1+9, 1+1...+1) = 4+4+4 => 3
13 = min(4+9, 4+4+4+1, 1+1+1+1+9, 1+1...+1) = 4+9 => 2

dp[n] = 1 ~ n 까지의 dp 중 가장 작은값 + 1
*/
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// perfect square number 1,4,9,16,25...
		if i*i < n+1 {
			dp[i*i] = 1
		}
		if dp[i] > 0 {
			continue
		}

		minVal := 1<<31 - 1
		for j := 1; j*j <= i; j++ {
			minVal = min(minVal, dp[i-(j*j)])
		}
		dp[i] = minVal + 1
	}
	return dp[n]
}

func main() {
	for i := 1; i < 15; i++ {
		fmt.Println(numSquares(i))
	}
}
