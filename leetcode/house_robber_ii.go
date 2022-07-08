/*
https://leetcode.com/problems/house-robber-ii/
213. House Robber II
Medium
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

Example 1:
Input: nums = [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.

Example 2:
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.

Example 3:
Input: nums = [1,2,3]
Output: 3

Constraints:
1 <= nums.length <= 100
0 <= nums[i] <= 1000
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
2,3,2
2(0번) 를 훔치면 더이상 훔칠 수 없다. 1번째, 2번째가 0번째와 인접해 있어서
case: 2
case: 3
case: 2
-> max money: 3

1,2,3,1
case: 1 + 3
case: 2 + 1
-> max money: 1+3 = 4

1,2,3
case: 1
case: 2
case: 3
-> max money: 3

점화식
0번째부터 시작하는 경우 dp0[i] = max(dp0[i-2] + nums[i], dp0[i])
1번째부터 시작하는 경우 dp1[i] = max(dp1[i-2] + nums[i], dp1[i])
max(dp0[len(dp0)-2], dp1[len(dp1)-1])
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	if len(nums) == 3 {
		return max(max(nums[0], nums[1]), nums[2])
	}
	dp0 := make([]int, len(nums))
	dp0[0] = nums[0]
	dp0[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums)-1; i++ {
		dp0[i] = max(dp0[i-1], dp0[i-2]+nums[i])
	}

	dp1 := make([]int, len(nums))
	dp1[1] = nums[1]
	dp1[2] = max(nums[1], nums[2])
	for i := 3; i < len(nums); i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i])
	}

	// fmt.Println(dp0)
	// fmt.Println(dp1)
	return max(dp0[len(dp0)-2], dp1[len(dp1)-1])
}

func main() {
	fmt.Println(rob([]int{0}))
	fmt.Println(rob([]int{1, 2}))
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{1, 2, 3}))
	fmt.Println(rob([]int{12, 3, 5, 32, 3, 4, 3, 6, 12, 6, 76, 7, 4, 6}))
	fmt.Println(rob([]int{12, 3, 5, 32, 3, 4, 3, 6, 12, 76, 7, 4, 6}))
}
