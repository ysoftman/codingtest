/*
https://leetcode.com/problems/majority-element-ii/
229. Majority Element II
Medium
Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

Example 1:
Input: nums = [3,2,3]
Output: [3]

Example 2:
Input: nums = [1]
Output: [1]

Example 3:
Input: nums = [1,2]
Output: [1,2]

Constraints:
1 <= nums.length <= 5 * 104
-109 <= nums[i] <= 109

Follow up: Could you solve the problem in linear time and in O(1) space?
*/
package main

import (
	"fmt"
	"sort"
)

// brute-force, space complexity: O(N)
func majorityElement2(nums []int) []int {
	r := []int{}
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if v > len(nums)/3 {
			r = append(r, k)
		}
	}
	return r
}

// space complexity: O(1)
func majorityElement(nums []int) []int {
	r := []int{}
	if len(nums) == 1 {
		return append(r, nums[0])
	}
	sort.Ints(nums)
	pre := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if pre != nums[i] {
			if cnt > len(nums)/3 {
				r = append(r, pre)
			}
			cnt = 0
		}
		cnt++
		if i == len(nums)-1 && cnt > len(nums)/3 {
			r = append(r, nums[i])
		}
		pre = nums[i]
	}
	return r
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{1}))
	fmt.Println(majorityElement([]int{1, 2}))
	fmt.Println(majorityElement([]int{2, 2}))
}
