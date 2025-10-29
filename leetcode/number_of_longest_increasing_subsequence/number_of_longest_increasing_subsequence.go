/*
https://leetcode.com/problems/number-of-longest-increasing-subsequence/
673. Number of Longest Increasing Subsequence
Medium
Given an integer array nums, return the number of longest increasing subsequences.

Notice that the sequence has to be strictly increasing.

Example 1:
Input: nums = [1,3,5,4,7]
Output: 2
Explanation: The two longest increasing subsequences are [1, 3, 4, 7] and [1, 3, 5, 7].

Example 2:
Input: nums = [2,2,2,2,2]
Output: 5
Explanation: The length of the longest increasing subsequence is 1, and there are 5 increasing subsequences of length 1, so output 5.

Constraints:
1 <= nums.length <= 2000
-106 <= nums[i] <= 106
*/
package main

import "fmt"

/*
1,3,5,4,7
dp[0] 1 = 1
dp[1] 3 = dp[0]+1, => number of LIC: 1
dp[n] N = dp[0]~dp[n-1] 중에서 N 보다 작은것들 중 가장 큰 dp 값(max) + 1, dp[0]~dp[n-1] 중 max 값의 개수 => number of LIC
*/
// dynamic programming, time complexity: O(n^2)
func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	cnt := make([]int, len(nums))
	// 각숫자는 자기 자신으로 최소 1개의 LIC 가 됨으로 모든 dp,cnt 1로 초기회
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		cnt[i] = 1
	}
	maxDP := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				// 가장 큰 dp 값 찾기
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					// 개수 기록 설정
					cnt[i] = cnt[j]
				} else if dp[i] == dp[j]+1 {
					// dp 값이 같은 경우 개수 추가
					cnt[i] += cnt[j]
				}
			}
		}
		// LIC 가 큰 dp의 값을 기록
		if maxDP < dp[i] {
			maxDP = dp[i]
		}
	}
	maxcnt := 0
	for i := 0; i < len(nums); i++ {
		// dp 중 maxDP 와 같은 dp의 cnt 값을 합산
		if dp[i] == maxDP {
			maxcnt += cnt[i]
		}
	}
	// fmt.Println(dp)

	return maxcnt
}

func main() {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
}
