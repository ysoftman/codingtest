/*
https://leetcode.com/problems/largest-number-at-least-twice-of-others
747. Largest Number At Least Twice of Others
Easy
You are given an integer array nums where the largest integer is unique.

Determine whether the largest element in the array is at least twice as much as every other number in the array. If it is, return the index of the largest element, or return -1 otherwise.

Example 1:
Input: nums = [3,6,1,0]
Output: 1
Explanation: 6 is the largest integer.
For every other number in the array x, 6 is at least twice as big as x.
The index of value 6 is 1, so we return 1.

Example 2:
Input: nums = [1,2,3,4]
Output: -1
Explanation: 4 is less than twice the value of 3, so we return -1.

Constraints:
2 <= nums.length <= 50
0 <= nums[i] <= 100
The largest element in nums is unique.
*/
package main

import (
	"fmt"
	"sort"
)

// time complexity O(NlogN)
func dominantIndex2(nums []int) int {
	orgNums := make([]int, len(nums))
	copy(orgNums, nums)

	sort.Slice(nums, func(a, b int) bool {
		return nums[a] > nums[b]
	})
	// fmt.Println(nums)
	if nums[0] >= nums[1]*2 {
		for i := range orgNums {
			if orgNums[i] == nums[0] {
				return i
			}
		}
	}
	return -1
}

// time complexity O(N)
func dominantIndex(nums []int) int {
	max := 0
	maxIdx := 0
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
			maxIdx = i
		}
	}
	for i := range nums {
		if nums[i]*2 > max && i != maxIdx {
			return -1
		}
	}
	return maxIdx
}

func main() {
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}))
	fmt.Println(dominantIndex([]int{1, 2, 3, 4}))
}
