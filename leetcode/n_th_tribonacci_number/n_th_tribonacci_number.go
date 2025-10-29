/*
https://leetcode.com/problems/n-th-tribonacci-number/
1137. N-th Tribonacci Number
Easy
The Tribonacci sequence Tn is defined as follows:
T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
Given n, return the value of Tn.

Example 1:
Input: n = 4
Output: 4
Explanation:
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4

Example 2:
Input: n = 25
Output: 1389537
*/
package main

import "fmt"

func tribonacci(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	if n >= 1 {
		dp[1] = 1
	}
	if n >= 2 {
		dp[2] = 1
	}
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]

}

func main() {
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(25))
}
