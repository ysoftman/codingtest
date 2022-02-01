/*
https://leetcode.com/problems/longest-increasing-subsequence/
300. Longest Increasing Subsequence
Given an integer array nums, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements. For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].

Example 1:
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:
Input: nums = [0,1,0,3,2,3]
Output: 4

Example 3:
Input: nums = [7,7,7,7,7,7,7]
Output: 1

*/
package main

import "fmt"

/*
i번째는 0~(i-1)=>j 를 확인하면서 j에서의 최대 lis 를 기록한다(dynamic programming/memoization)
*/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	// 0번째 element 는 1개, 즉 lis 길이 1
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		maxlis := 0
		for j := 0; j < i; j++ {
			// j들중 i 보다 작은 경우 중 가증 큰 lis 길이를 찾아
			if nums[j] < nums[i] && maxlis < dp[j] {
				maxlis = dp[j]
			}
		}
		// 현재 i lis 길이 (dp[i])를 업데이트 한다.
		dp[i] = maxlis + 1

	}
	// dp 중에서 가능 큰 lis 길이를 찾아야 한다.
	maxlis := dp[0]
	for i := 0; i < len(nums); i++ {
		if maxlis < dp[i] {
			maxlis = dp[i]
		}
	}
	return maxlis
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}
