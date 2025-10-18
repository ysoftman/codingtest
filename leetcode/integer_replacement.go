/*
https://leetcode.com/problems/integer-replacement
397. Integer Replacement
Medium
Given a positive integer n, you can apply one of the following operations:

If n is even, replace n with n / 2.
If n is odd, replace n with either n + 1 or n - 1.
Return the minimum number of operations needed for n to become 1.

Example 1:
Input: n = 8
Output: 3
Explanation: 8 -> 4 -> 2 -> 1

Example 2:
Input: n = 7
Output: 4
Explanation: 7 -> 8 -> 4 -> 2 -> 1
or 7 -> 6 -> 3 -> 2 -> 1

Example 3:
Input: n = 4
Output: 2

Constraints:
1 <= n <= 231 - 1
*/
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// bottom-up
// n=200000000, memory limit exceeded
// func integerReplacement(n int) int {
// 	if n <= 1 {
// 		return 0
// 	}
// 	if n == 2 {
//      return 1
// 	}
// 	dp := make([]int, n+1)
// 	dp[0] = 0
// 	dp[1] = 0
// 	dp[2] = 1
// 	for i := 3; i <= n; i++ {
// 		if i%2 == 0 {
// 			dp[i] = dp[i/2] + 1
// 		} else {
// 			// dp[n+1] 을 파악해되는데, dp[n+1] = dp[(n+1)/2]+1 과 같다.
// 			dp[i] = min(dp[i-1]+1, dp[(i+1)/2]+2)
// 		}
// 	}
// 	return dp[n]
// }

// top-down
func integerReplacement(n int) int {
	if n <= 1 {
		return 0
	}
	r := 0
	if n%2 == 0 {
		n /= 2
		r = integerReplacement(n) + 1
	} else {
		r = min(integerReplacement(n-1), integerReplacement(n+1)) + 1
	}
	return r
}

func main() {
	fmt.Println(integerReplacement(8))
	fmt.Println(integerReplacement(7))
	fmt.Println(integerReplacement(4))
	fmt.Println(integerReplacement(1))
	fmt.Println(integerReplacement(2))
	fmt.Println(integerReplacement(30))
	fmt.Println(integerReplacement(200000000))
}
