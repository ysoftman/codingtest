/*
https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
81. Search in Rotated Sorted Array II
Medium
There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).
Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].
Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.
You must decrease the overall operation steps as much as possible.

Example 1:
Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true

Example 2:
Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false

Constraints:
1 <= nums.length <= 5000
-104 <= nums[i] <= 104
nums is guaranteed to be rotated at some pivot.
-104 <= target <= 104

Follow up: This problem is similar to Search in Rotated Sorted Array, but nums may contain duplicates. Would this affect the runtime complexity? How and why?
*/

package main

import "fmt"

func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
func findPivot(nums []int, left, right int) int {
	if left > right {
		return 0
	}
	mid := left + (right-left)/2
	if mid+1 < len(nums) && nums[mid] > nums[mid+1] {
		return mid + 1
	}
	if mid-1 >= 0 && nums[mid] < nums[mid-1] {
		return mid
	}
	if pivot := findPivot(nums, left, mid-1); pivot > 0 {
		return pivot
	}
	if pivot := findPivot(nums, mid+1, right); pivot > 0 {
		return pivot
	}
	return 0
}
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	pivot := findPivot(nums, left, right)
	// fmt.Println("pivot index:", pivot)

	if binarySearch(nums, 0, pivot-1, target) >= 0 {
		return true
	}
	if binarySearch(nums, pivot, len(nums)-1, target) >= 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	fmt.Println(search([]int{1, 0, 1, 1, 1}, 0))
	fmt.Println(search([]int{1}, 0))
	fmt.Println(search([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2))
}
