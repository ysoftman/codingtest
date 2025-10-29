/*
https://leetcode.com/problems/domino-and-tromino-tiling/
790. Domino and Tromino Tiling
Medium

You have two types of tiles: a 2 x 1 domino shape and a tromino shape. You may rotate these shapes.

Given an integer n, return the number of ways to tile an 2 x n board. Since the answer may be very large, return it modulo 109 + 7.

In a tiling, every square must be covered by a tile. Two tilings are different if and only if there are two 4-directionally adjacent cells on the board such that exactly one of the tilings has both squares occupied by a tile.

Example 1:
Input: n = 3
Output: 5
Explanation: The five different ways are show above.

Example 2:
Input: n = 1
Output: 1

Constraints:
1 <= n <= 1000
*/
package main

import "fmt"

// dynamic programming
// 1, 1, 2, 5, 11, 24, 53 ....
// f(n) = 2*f(n-1) + f(n-3)
func numTilings(n int) int {
	modulo := 1_000_000_000 + 7
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % modulo
	}
	return dp[n]
}

func main() {
	fmt.Println(numTilings(1))
	fmt.Println(numTilings(2))
	fmt.Println(numTilings(3))
	fmt.Println(numTilings(4))
	fmt.Println(numTilings(5))
}
