/*
https://leetcode.com/problems/minimum-path-sum/
64. Minimum Path Sum
Medium

Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.
Note: You can only move either down or right at any point in time.

Example 1:
Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.

Example 2:
Input: grid = [[1,2,3],[4,5,6]]
Output: 12

Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 100
*/
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i >= 1 && j >= 1 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			} else if i >= 1 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i == 0 && j >= 1 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}
