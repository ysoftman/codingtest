/*
https://leetcode.com/problems/search-in-rotated-sorted-array/
33. Search in Rotated Sorted Array
Medium
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Example 3:
Input: nums = [1], target = 0
Output: -1


Constraints:

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
All values of nums are unique.
nums is an ascending array that is possibly rotated.
-104 <= target <= 104
*/
package main

import "fmt"

// working, but it's not a solution! => O(N)
func search2(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return -1
}

func binarySearch(nums []int, target, left, right int) int {
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

// O(logN)
func search(nums []int, target int) int {
	rotatedIdx := 0
	// O(logN)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if mid+1 < len(nums) && nums[mid] > nums[mid+1] {
			rotatedIdx = mid
			break
		}
		if mid-1 >= 0 && nums[mid] < nums[mid-1] {
			rotatedIdx = mid - 1
			break
		}

		if nums[mid] > nums[0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// O(logN)
	fmt.Println("rotatedIdx:", rotatedIdx)
	if nums[0] <= target && target <= nums[rotatedIdx] {
		return binarySearch(nums, target, 0, rotatedIdx)
	} else if rotatedIdx+1 < len(nums) && nums[rotatedIdx+1] <= target && target <= nums[len(nums)-1] {
		return binarySearch(nums, target, rotatedIdx+1, len(nums)-1)
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 2))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 1))
	fmt.Println(search([]int{1}, 0))
	fmt.Println(search([]int{7, 8, 1, 2, 3, 4, 5, 6}, 2))
}
