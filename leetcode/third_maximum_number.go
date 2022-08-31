/*
https://leetcode.com/problems/third-maximum-number/
414. Third Maximum Number
Easy
Given an integer array nums, return the third distinct maximum number in this array. If the third maximum does not exist, return the maximum number.

Example 1:
Input: nums = [3,2,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2.
The third distinct maximum is 1.

Example 2:
Input: nums = [1,2]
Output: 2
Explanation:
The first distinct maximum is 2.
The second distinct maximum is 1.
The third distinct maximum does not exist, so the maximum (2) is returned instead.

Example 3:
Input: nums = [2,2,3,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2 (both 2's are counted together since they have the same value).
The third distinct maximum is 1.


Constraints:
1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1

Follow up: Can you find an O(n) solution?
*/
package main

import (
	"fmt"
	"sort"
)

// time complexity: O(nlogn)
func thirdMax2(nums []int) int {
	// descending sort
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] > nums[b]
	})
	pre := nums[0]
	rank := 1
	for i := 0; i < len(nums); i++ {
		// skip duplicates
		if nums[i] == pre {
			continue
		}
		pre = nums[i]
		rank++
		if rank == 3 {
			return nums[i]
		}
	}
	return nums[0]
}

// time complexity: O(n)
func thirdMax(nums []int) int {
	r := []int{-(1<<63 - 1), -(1<<63 - 1), -(1<<63 - 1)}
	for i := 0; i < len(nums); i++ {
		// 이미 r 에 있는 값이면 스킵
		if r[0] == nums[i] || r[1] == nums[i] || r[2] == nums[i] {
			continue
		}
		// r[0]~[2] 에 내림 차순으로 값 유지
		if r[0] < nums[i] {
			r[2] = r[1]
			r[1] = r[0]
			r[0] = nums[i]
		} else if r[1] < nums[i] {
			r[2] = r[1]
			r[1] = nums[i]
		} else if r[2] < nums[i] {
			r[2] = nums[i]
		}
	}
	// fmt.Println(r)
	// 3개 값이 없는 경우 최대 값 리턴
	if r[2] == -(1<<63 - 1) {
		return r[0]
	}
	return r[2]
}

func main() {
	fmt.Println(thirdMax([]int{3, 2, 1}))
	fmt.Println(thirdMax([]int{1, 2}))
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
	fmt.Println(thirdMax([]int{5, 2, 2}))
	fmt.Println(thirdMax([]int{1, 2, 2, 5, 3, 5}))
}
