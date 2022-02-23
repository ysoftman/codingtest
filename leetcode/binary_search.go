/*
https://leetcode.com/problems/binary-search/
704. Binary Search
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.
Example 1:
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1

*/
package main

import "fmt"

func recurSearch(nums []int, left, right, target int) int {
	mid := left + (right-left)/2
	if left > right {
		return -1
	}
	if target == nums[mid] {
		return mid
	}
	if target > nums[mid] {
		return recurSearch(nums, mid+1, right, target)
	} else {
		return recurSearch(nums, left, mid-1, target)
	}
}
func search(nums []int, target int) int {
	return recurSearch(nums, 0, len(nums)-1, target)
}

func search2(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + ((j - i) / 2)
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println(search([]int{1, 2}, 2))
	fmt.Println(search([]int{5}, 5))
	fmt.Println("-----")
	fmt.Println(search2([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search2([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println(search2([]int{1, 2}, 2))
	fmt.Println(search2([]int{5}, 5))
}
