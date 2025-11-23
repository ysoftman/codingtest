/*
1262. Greatest Sum Divisible by Three
https://leetcode.com/problems/greatest-sum-divisible-by-three
Medium
Given an integer array nums, return the maximum possible sum of elements of the array such that it is divisible by three.

Example 1:
Input: nums = [3,6,5,1,8]
Output: 18
Explanation: Pick numbers 3, 6, 1 and 8 their sum is 18 (maximum sum divisible by 3).

Example 2:
Input: nums = [4]
Output: 0
Explanation: Since 4 is not divisible by 3, do not pick any number.

Example 3:
Input: nums = [1,2,3,4,4]
Output: 12
Explanation: Pick numbers 1, 3, 4 and 4 their sum is 12 (maximum sum divisible by 3).

Constraints:
1 <= nums.length <= 4 * 104
1 <= nums[i] <= 104
*/
package main

import "fmt"

/*
dynamic programming
i 위치에서 나머지가 0,1,2 케이스 각각 장 큰 값을 유지해간다. 0 이면 nums[i] 값을 포함하지 않는경우임
   3  6  5  1  8
0  3  9  9 15 18
1  0  0  0 10 22
2  0  0 14 14 23
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSumDivThree(nums []int) int {
	dp := make([][]int, len(nums))
	dp[0] = make([]int, 3)

	if nums[0]%3 == 0 {
		dp[0][0] = nums[0]
	} else if nums[0]%3 == 1 {
		dp[0][1] = nums[0]
	} else if nums[0]%3 == 2 {
		dp[0][2] = nums[0]
	}

	// fmt.Println(dp[0])
	for i := 1; i < len(nums); i++ {
		dp[i] = make([]int, 3) // modulo: 0, 1, 2
		dp[i][0] = dp[i-1][0]
		dp[i][1] = dp[i-1][1]
		dp[i][2] = dp[i-1][2]
		for j := range dp[i] {
			switch (nums[i] + dp[i-1][j]) % 3 {
			case 0:
				dp[i][0] = max(dp[i][0], nums[i]+dp[i-1][j])
			case 1:
				dp[i][1] = max(dp[i][1], nums[i]+dp[i-1][j])
			case 2:
				dp[i][2] = max(dp[i][2], nums[i]+dp[i-1][j])
			}
		}
		// fmt.Println(dp[i])
	}

	return dp[len(nums)-1][0]
}

func main() {
	fmt.Println(maxSumDivThree([]int{3, 6, 5, 1, 8}))
	fmt.Println(maxSumDivThree([]int{4}))
	fmt.Println(maxSumDivThree([]int{1, 2, 3, 4, 4}))
}
