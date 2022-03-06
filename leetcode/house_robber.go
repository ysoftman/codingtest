/*
https://leetcode.com/problems/house-robber/
198. House Robber
Medium
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

Example 1:
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.

Example 2:
Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.
*/
package main

import "fmt"

/*
[2, 7, 9, 3, 1]
i i+2 i+3

2 9 1
2 3
7 3
7 1

backtracking => i += max(i-3, i-2)
ex) 3 += max(2, 7)
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func rob(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		a, b := 0, 0
		if i-3 >= 0 {
			a = dp[i-3]
		}
		if i-2 >= 0 {
			b = dp[i-2]
		}
		dp[i] += nums[i] + max(a, b)
	}
	if len(dp) == 1 {
		return dp[0]
	}
	if dp[len(dp)-2] > dp[len(dp)-1] {
		return dp[len(dp)-2]
	}
	return dp[len(dp)-1]
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{0}))
}
