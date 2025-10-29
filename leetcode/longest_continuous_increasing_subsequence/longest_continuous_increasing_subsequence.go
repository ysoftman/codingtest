/*
https://leetcode.com/problems/longest-continuous-increasing-subsequence/
674. Longest Continuous Increasing Subsequence
Easy

Given an unsorted array of integers nums, return the length of the longest continuous increasing subsequence (i.e. subarray). The subsequence must be strictly increasing.

A continuous increasing subsequence is defined by two indices l and r (l < r) such that it is [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] and for each l <= i < r, nums[i] < nums[i + 1].

Example 1:
Input: nums = [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5] with length 3.
Even though [1,3,5,7] is an increasing subsequence, it is not continuous as elements 5 and 7 are separated by element
4.

Example 2:
Input: nums = [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2] with length 1. Note that it must be strictly
increasing.

Constraints:
1 <= nums.length <= 104
-109 <= nums[i] <= 109
*/
package main

import "fmt"

func findLengthOfLCIS2(nums []int) int {
	maxlen := 1
	cnt := 1
	for i := 0; i < len(nums)-1; i++ {
		// cis 길이 파악
		if nums[i] < nums[i+1] {
			cnt++
			if maxlen < cnt {
				maxlen = cnt
			}
		} else { // cis 1 부터 시작
			cnt = 1
		}
	}
	return maxlen
}

func findLengthOfLCIS(nums []int) int {
	maxlen := 1
	pre := 0
	for i := 1; i < len(nums); i++ {
		// 줄어드는 지점 파악
		if i > 0 && nums[i-1] >= nums[i] {
			pre = i
		}
		// 현재 - 이전(줄어드는 지점) = 하나의 cis 길이가 된다.
		if maxlen < (i-pre)+1 {
			maxlen = (i - pre) + 1
		}
	}
	return maxlen
}

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2, 2}))
	fmt.Println(findLengthOfLCIS([]int{1}))
	fmt.Println(findLengthOfLCIS([]int{2, 2, 3, 4, 1, 5, 8, 11, 22}))
	fmt.Println(findLengthOfLCIS([]int{2, 2, 3, 4, 1, 5, 8, 11, 22, 8}))
}
