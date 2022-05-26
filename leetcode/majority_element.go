/*
https://leetcode.com/problems/majority-element/
169. Majority Element
Given an array nums of size n, return the majority element.
The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

Example 1:
Input: nums = [3,2,3]
Output: 3

Example 2:
Input: nums = [2,2,1,1,1,2,2]
Output: 2

Constraints:
n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109

Follow-up: Could you solve the problem in linear time and in O(1) space?
*/

package main

import (
	"fmt"
	"sort"
)

// brute-force, space complexity O(N)
func majorityElement2(nums []int) int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}

// space complexity O(1)
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	pre := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			if 1+(i-pre) > len(nums)/2 {
				return nums[i]
			}
		} else {
			pre = i
		}
	}
	return 0
}

func main() {
	fmt.Println(majorityElement([]int{1}))
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
