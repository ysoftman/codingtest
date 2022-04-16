/*
https://leetcode.com/problems/unique-binary-search-trees/
96. Unique Binary Search Trees
Medium
Given an integer n, return the number of structurally unique BST's (binary search trees) which has exactly n nodes of unique values from 1 to n.

Example 1:
Input: n = 3
Output: 5

Example 2:
Input: n = 1
Output: 1

Constraints:
1 <= n <= 19
*/
package main

import "fmt"

/*
0 => 1
1 => (dp[1-1]*dp[1-1]) => (1) => 1
2 => (dp[1-1]*dp[2-1])+(dp[2-1]*dp[2-2]) => (1*1)+(1*1) => 2
3 => (dp[1-1]*dp[3-1])+(dp[2-1]*dp[3-2])+(dp[3-1]*dp[3-3]) => (1*2)+(1*1)+(2*1) => 5
*/
func numTrees(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			left := j - 1
			right := i - j
			dp[i] += dp[left] * dp[right]
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(4))
	fmt.Println(numTrees(5))
	fmt.Println(numTrees(6))
	fmt.Println(numTrees(7))
}
