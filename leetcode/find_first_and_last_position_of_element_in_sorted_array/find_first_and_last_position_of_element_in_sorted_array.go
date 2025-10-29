/*
https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
34. Find First and Last Position of Element in Sorted Array
Medium
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
If target is not found in the array, return [-1, -1].
You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Example 2:
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

Example 3:
Input: nums = [], target = 0
Output: [-1,-1]

Constraints:
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109
*/
package main

import "fmt"

func recursiveSearchRange(nums []int, target, left, right int, result []int) {
	if left < 0 || right >= len(nums) {
		return
	}
	// binary search
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if mid < result[0] {
				result[0] = mid
			}
			if mid > result[1] {
				result[1] = mid
			}
			// fmt.Println(left, mid, right)
			recursiveSearchRange(nums, target, left, mid-1, result)
			recursiveSearchRange(nums, target, mid+1, right, result)
			return
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}
}
func searchRange(nums []int, target int) []int {
	result := []int{1<<31 - 1, -1}
	recursiveSearchRange(nums, target, 0, len(nums)-1, result)

	if result[0] == 1<<31-1 {
		result[0] = -1
	}
	return result
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
}
