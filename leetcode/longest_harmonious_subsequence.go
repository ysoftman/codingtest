/*
https://leetcode.com/problems/longest-harmonious-subsequence/
594. Longest Harmonious Subsequence
Easy

We define a harmonious array as an array where the difference between its maximum value and its minimum value is exactly 1.

Given an integer array nums, return the length of its longest harmonious subsequence among all its possible subsequences.

A subsequence of array is a sequence that can be derived from the array by deleting some or no elements without changing the order of the remaining elements.

Example 1:
Input: nums = [1,3,2,2,5,2,3,7]
Output: 5
Explanation: The longest harmonious subsequence is [3,2,2,2,3].

Example 2:
Input: nums = [1,2,3,4]
Output: 2

Example 3:
Input: nums = [1,1,1,1]
Output: 0

Constraints:
1 <= nums.length <= 2 * 104
-109 <= nums[i] <= 109
*/
package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func findLHS(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	r := 0
	for k, v := range m {
		// 현재숫자의 cnt 와 현재숫자+1 의 cnt 를 합(길이)... 중 가장 큰 길이
		if cnt, ok := m[k+1]; ok {
			r = max(r, v+cnt)
		}
	}
	return r
}

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	fmt.Println(findLHS([]int{1, 2, 3, 4}))
	fmt.Println(findLHS([]int{1, 1, 1, 1}))
}
