/*
https://leetcode.com/problems/triangle/

120. Triangle
Medium

Given a triangle array, return the minimum path sum from top to bottom.
For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

Example 1:
Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:
   2
  3 4
 6 5 7
4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).


Example 2:
Input: triangle = [[-10]]
Output: -10
*/
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minimumTotal(triangle [][]int) int {
	lenRows := len(triangle)

	dp := make([][]int, lenRows)
	for i := 0; i < lenRows; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}

	for i := 0; i < lenRows; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			a, b, c := 0, 0, 0
			finda, findb := false, false
			if i-1 >= 0 && j-1 >= 0 {
				a = dp[i-1][j-1]
				finda = true
			}
			if i-1 >= 0 && j < len(triangle[i-1]) {
				b = dp[i-1][j]
				findb = true
			}
			if finda {
				c = a
			}
			if findb {
				c = b
			}
			if finda && findb {
				c = min(a, b)
			}
			dp[i][j] += c
			dp[i][j] += triangle[i][j]
			// fmt.Println(a, b, triangle[i][j], "..", dp[i][j])
		}
	}

	// find in last row's columns
	minSum := 1<<31 - 1
	for i := 0; i < len(dp[lenRows-1]); i++ {
		if minSum > dp[lenRows-1][i] {
			minSum = dp[lenRows-1][i]
		}
	}
	return minSum
}
func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
	fmt.Println(minimumTotal([][]int{{-10}}))
}
