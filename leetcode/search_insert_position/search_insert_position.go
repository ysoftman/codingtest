/*
https://leetcode.com/problems/search-insert-position/
35. Search Insert Position
Easy
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
You must write an algorithm with O(log n) runtime complexity.


Example 1:
Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:
Input: nums = [1,3,5,6], target = 2
Output: 1

Example 3:
Input: nums = [1,3,5,6], target = 7
Output: 4
*/

package main

import "fmt"

// O(n)
func searchInsert2(nums []int, target int) int {
	for i := range nums {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

// O(log n)
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	if nums[len(nums)-1] < target {
		return len(nums)
	}
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
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
	nums = []int{1, 3, 5, 6}
	target = 2
	fmt.Println(searchInsert(nums, target))
	nums = []int{1, 3, 5, 6}
	target = 7
	fmt.Println(searchInsert(nums, target))
}
